package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/wailsapp/wails/v3/internal/capabilities"

	"github.com/wailsapp/wails/v3/pkg/icons"

	"github.com/samber/lo"

	"github.com/wailsapp/wails/v2/pkg/assetserver"
	"github.com/wailsapp/wails/v2/pkg/assetserver/webview"
	assetserveroptions "github.com/wailsapp/wails/v2/pkg/options/assetserver"

	wailsruntime "github.com/wailsapp/wails/v3/internal/runtime"
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/logger"
)

var globalApplication *App

// isDebugMode is true if the application is running in debug mode
var isDebugMode func() bool

func init() {
	runtime.LockOSThread()
}

type EventListener struct {
	callback func()
}

func New(appOptions Options) *App {
	if globalApplication != nil {
		return globalApplication
	}

	mergeApplicationDefaults(&appOptions)

	result := &App{
		options:                   appOptions,
		applicationEventListeners: make(map[uint][]*EventListener),
		windows:                   make(map[uint]*WebviewWindow),
		systemTrays:               make(map[uint]*SystemTray),
		log:                       logger.New(appOptions.Logger.CustomLoggers...),
		contextMenus:              make(map[string]*Menu),
		pid:                       os.Getpid(),
	}
	globalApplication = result

	if !appOptions.Logger.Silent {
		result.log.AddOutput(&logger.Console{})
	}

	// Patch isDebug if we aren't in prod mode
	if isDebugMode == nil {
		isDebugMode = func() bool {
			return true
		}
	}

	result.Events = NewWailsEventProcessor(result.dispatchEventToWindows)

	opts := assetserveroptions.Options{
		Assets:     appOptions.Assets.FS,
		Handler:    appOptions.Assets.Handler,
		Middleware: assetserveroptions.Middleware(appOptions.Assets.Middleware),
	}

	// TODO ServingFrom disk?
	srv, err := assetserver.NewAssetServer("", opts, false, nil, wailsruntime.RuntimeAssetsBundle)
	if err != nil {
		result.fatal(err.Error())
	}

	// Pass through the capabilities
	srv.GetCapabilities = func() []byte {
		return globalApplication.capabilities.AsBytes()
	}

	srv.GetFlags = func() []byte {
		updatedOptions := result.impl.GetFlags(appOptions)
		flags, err := json.Marshal(updatedOptions)
		if err != nil {
			log.Fatal("Invalid flags provided to application: ", err.Error())
		}
		return flags
	}

	srv.UseRuntimeHandler(NewMessageProcessor())
	result.assets = srv

	result.bindings, err = NewBindings(appOptions.Bind)
	if err != nil {
		println("Fatal error in application initialisation: ", err.Error())
		os.Exit(1)
	}

	result.plugins = NewPluginManager(appOptions.Plugins, srv)
	err = result.plugins.Init()
	if err != nil {
		result.Quit()
		os.Exit(1)
	}

	err = result.bindings.AddPlugins(appOptions.Plugins)
	if err != nil {
		println("Fatal error in application initialisation: ", err.Error())
		os.Exit(1)
	}

	return result
}

func mergeApplicationDefaults(o *Options) {
	if o.Name == "" {
		o.Name = "My Wails Application"
	}
	if o.Description == "" {
		o.Description = "An application written using Wails"
	}
	if o.Icon == nil {
		o.Icon = icons.ApplicationLightMode256
	}

}

type (
	platformApp interface {
		run() error
		destroy()
		setApplicationMenu(menu *Menu)
		name() string
		getCurrentWindowID() uint
		showAboutDialog(name string, description string, icon []byte)
		setIcon(icon []byte)
		on(id uint)
		dispatchOnMainThread(id uint)
		hide()
		show()
		getPrimaryScreen() (*Screen, error)
		getScreens() ([]*Screen, error)
		GetFlags(options Options) map[string]any
		isOnMainThread() bool
	}

	runnable interface {
		run()
	}
)

func processPanic(value any) {
	if value == nil {
		value = fmt.Errorf("unknown error")
	}
	if globalApplication.options.PanicHandler != nil {
		globalApplication.options.PanicHandler(value)
		return
	}
	// Print the panic details
	fmt.Printf("Panic occurred: %v", value)

	// Print the stack trace
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	fmt.Println("Stack trace:")
	fmt.Println(string(buf))
	os.Exit(1)
}

// Messages sent from javascript get routed here
type windowMessage struct {
	windowId uint
	message  string
}

var windowMessageBuffer = make(chan *windowMessage)

type dragAndDropMessage struct {
	windowId  uint
	filenames []string
}

var windowDragAndDropBuffer = make(chan *dragAndDropMessage)

var _ webview.Request = &webViewAssetRequest{}

const webViewRequestHeaderWindowId = "x-wails-window-id"
const webViewRequestHeaderWindowName = "x-wails-window-name"

type webViewAssetRequest struct {
	webview.Request
	windowId   uint
	windowName string
}

func (r *webViewAssetRequest) Header() (http.Header, error) {
	h, err := r.Request.Header()
	if err != nil {
		return nil, err
	}

	hh := h.Clone()
	hh.Set(webViewRequestHeaderWindowId, strconv.FormatUint(uint64(r.windowId), 10))
	return hh, nil
}

var webviewRequests = make(chan *webViewAssetRequest)

type App struct {
	options                       Options
	applicationEventListeners     map[uint][]*EventListener
	applicationEventListenersLock sync.RWMutex

	// Windows
	windows     map[uint]*WebviewWindow
	windowsLock sync.Mutex

	// System Trays
	systemTrays      map[uint]*SystemTray
	systemTraysLock  sync.Mutex
	systemTrayID     uint
	systemTrayIDLock sync.RWMutex

	// MenuItems
	menuItems     map[uint]*MenuItem
	menuItemsLock sync.Mutex

	// Running
	running    bool
	runLock    sync.Mutex
	pendingRun []runnable

	bindings *Bindings
	plugins  *PluginManager

	// platform app
	impl platformApp

	// The main application menu
	ApplicationMenu *Menu

	clipboard *Clipboard
	Events    *EventProcessor
	log       *logger.Logger

	contextMenus     map[string]*Menu
	contextMenusLock sync.Mutex

	assets   *assetserver.AssetServer
	startURL string

	// Hooks
	windowCreatedCallbacks []func(window *WebviewWindow)
	pid                    int

	// Capabilities
	capabilities capabilities.Capabilities
}

func (a *App) getSystemTrayID() uint {
	a.systemTrayIDLock.Lock()
	defer a.systemTrayIDLock.Unlock()
	a.systemTrayID++
	return a.systemTrayID
}

func (a *App) getWindowForID(id uint) *WebviewWindow {
	a.windowsLock.Lock()
	defer a.windowsLock.Unlock()
	return a.windows[id]
}

func (a *App) deleteWindowByID(id uint) {
	a.windowsLock.Lock()
	defer a.windowsLock.Unlock()
	delete(a.windows, id)
}

func (a *App) Capabilities() capabilities.Capabilities {
	return a.capabilities
}

func (a *App) On(eventType events.ApplicationEventType, callback func()) func() {
	eventID := uint(eventType)
	a.applicationEventListenersLock.Lock()
	defer a.applicationEventListenersLock.Unlock()
	listener := &EventListener{
		callback: callback,
	}
	a.applicationEventListeners[eventID] = append(a.applicationEventListeners[eventID], listener)
	if a.impl != nil {
		go a.impl.on(eventID)
	}

	return func() {
		// lock the map
		a.applicationEventListenersLock.Lock()
		defer a.applicationEventListenersLock.Unlock()
		// Remove listener
		a.applicationEventListeners[eventID] = lo.Without(a.applicationEventListeners[eventID], listener)
	}
}
func (a *App) NewWebviewWindow() *WebviewWindow {
	return a.NewWebviewWindowWithOptions(WebviewWindowOptions{})
}

func (a *App) GetPID() int {
	return a.pid
}

func (a *App) info(message string, args ...any) {
	a.Log(&logger.Message{
		Level:   "INFO",
		Message: message,
		Data:    args,
		Sender:  "Wails",
	})
}

func (a *App) fatal(message string, args ...any) {
	msg := "************** FATAL **************\n"
	msg += message
	msg += "***********************************\n"

	a.Log(&logger.Message{
		Level:   "FATAL",
		Message: msg,
		Data:    args,
		Sender:  "Wails",
	})

	a.log.Flush()
	os.Exit(1)
}

func (a *App) error(message string, args ...any) {
	a.Log(&logger.Message{
		Level:   "ERROR",
		Message: message,
		Data:    args,
		Sender:  "Wails",
	})
}

func (a *App) NewWebviewWindowWithOptions(windowOptions WebviewWindowOptions) *WebviewWindow {
	newWindow := NewWindow(windowOptions)
	id := newWindow.id

	a.windowsLock.Lock()
	a.windows[id] = newWindow
	a.windowsLock.Unlock()

	// Call hooks
	for _, hook := range a.windowCreatedCallbacks {
		hook(newWindow)
	}

	a.runOrDeferToAppRun(newWindow)

	return newWindow
}

func (a *App) NewSystemTray() *SystemTray {
	id := a.getSystemTrayID()
	newSystemTray := NewSystemTray(id)

	a.systemTraysLock.Lock()
	a.systemTrays[id] = newSystemTray
	a.systemTraysLock.Unlock()

	a.runOrDeferToAppRun(newSystemTray)

	return newSystemTray
}

func (a *App) Run() error {
	a.info("Starting application")

	// Setup panic handler
	defer func() {
		if err := recover(); err != nil {
			processPanic(err)
		}
	}()

	a.impl = newPlatformApp(a)
	go func() {
		for {
			event := <-applicationEvents
			a.handleApplicationEvent(event)
		}
	}()
	go func() {
		for {
			event := <-windowEvents
			a.handleWindowEvent(event)
		}
	}()
	go func() {
		for {
			request := <-webviewRequests
			a.handleWebViewRequest(request)
		}
	}()
	go func() {
		for {
			event := <-windowMessageBuffer
			a.handleWindowMessage(event)
		}
	}()
	go func() {
		for {
			dragAndDropMessage := <-windowDragAndDropBuffer
			a.handleDragAndDropMessage(dragAndDropMessage)
		}
	}()

	go func() {
		for {
			menuItemID := <-menuItemClicked
			a.handleMenuItemClicked(menuItemID)
		}
	}()

	a.runLock.Lock()
	a.running = true

	for _, systray := range a.pendingRun {
		go systray.run()
	}
	a.pendingRun = nil

	a.runLock.Unlock()

	// set the application menu
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		a.impl.setApplicationMenu(a.ApplicationMenu)
	}
	a.impl.setIcon(a.options.Icon)

	err := a.impl.run()
	if err != nil {
		return err
	}

	a.plugins.Shutdown()

	return nil
}

func (a *App) handleApplicationEvent(event uint) {
	a.applicationEventListenersLock.RLock()
	listeners, ok := a.applicationEventListeners[event]
	a.applicationEventListenersLock.RUnlock()
	if !ok {
		return
	}
	for _, listener := range listeners {
		go listener.callback()
	}
}

func (a *App) handleDragAndDropMessage(event *dragAndDropMessage) {
	// Get window from window map
	a.windowsLock.Lock()
	window, ok := a.windows[event.windowId]
	a.windowsLock.Unlock()
	if !ok {
		log.Printf("WebviewWindow #%d not found", event.windowId)
		return
	}
	// Get callback from window
	window.handleDragAndDropMessage(event)
}

func (a *App) handleWindowMessage(event *windowMessage) {
	// Get window from window map
	a.windowsLock.Lock()
	window, ok := a.windows[event.windowId]
	a.windowsLock.Unlock()
	if !ok {
		log.Printf("WebviewWindow #%d not found", event.windowId)
		return
	}
	// Get callback from window
	window.handleMessage(event.message)
}

func (a *App) handleWebViewRequest(request *webViewAssetRequest) {
	// Get window from window map
	url, _ := request.URL()
	a.info("Window: '%s', Request: %s", request.windowName, url)
	a.assets.ServeWebViewRequest(request)
}

func (a *App) handleWindowEvent(event *WindowEvent) {
	// Get window from window map
	a.windowsLock.Lock()
	window, ok := a.windows[event.WindowID]
	a.windowsLock.Unlock()
	if !ok {
		log.Printf("WebviewWindow #%d not found", event.WindowID)
		return
	}
	window.handleWindowEvent(event.EventID)
}

func (a *App) handleMenuItemClicked(menuItemID uint) {
	menuItem := getMenuItemByID(menuItemID)
	if menuItem == nil {
		log.Printf("MenuItem #%d not found", menuItemID)
		return
	}
	menuItem.handleClick()
}

func (a *App) CurrentWindow() *WebviewWindow {
	if a.impl == nil {
		return nil
	}
	id := a.impl.getCurrentWindowID()
	a.windowsLock.Lock()
	defer a.windowsLock.Unlock()
	return a.windows[id]
}

func (a *App) Quit() {
	invokeSync(func() {
		a.windowsLock.Lock()
		for _, window := range a.windows {
			window.Destroy()
		}
		a.windowsLock.Unlock()
		a.systemTraysLock.Lock()
		for _, systray := range a.systemTrays {
			systray.Destroy()
		}
		a.systemTraysLock.Unlock()
		if a.impl != nil {
			a.impl.destroy()
		}
	})
}

func (a *App) SetMenu(menu *Menu) {
	a.ApplicationMenu = menu
	if a.impl != nil {
		a.impl.setApplicationMenu(menu)
	}
}
func (a *App) ShowAboutDialog() {
	if a.impl != nil {
		a.impl.showAboutDialog(a.options.Name, a.options.Description, a.options.Icon)
	}
}

func InfoDialog() *MessageDialog {
	return newMessageDialog(InfoDialogType)
}

func QuestionDialog() *MessageDialog {
	return newMessageDialog(QuestionDialogType)
}

func WarningDialog() *MessageDialog {
	return newMessageDialog(WarningDialogType)
}

func ErrorDialog() *MessageDialog {
	return newMessageDialog(ErrorDialogType)
}

func OpenDirectoryDialog() *MessageDialog {
	return newMessageDialog(OpenDirectoryDialogType)
}

func OpenFileDialog() *OpenFileDialogStruct {
	return newOpenFileDialog()
}

func SaveFileDialog() *SaveFileDialogStruct {
	return newSaveFileDialog()
}

func (a *App) GetPrimaryScreen() (*Screen, error) {
	return a.impl.getPrimaryScreen()
}

func (a *App) GetScreens() ([]*Screen, error) {
	return a.impl.getScreens()
}

func (a *App) Clipboard() *Clipboard {
	if a.clipboard == nil {
		a.clipboard = newClipboard()
	}
	return a.clipboard
}

func (a *App) dispatchOnMainThread(fn func()) {
	// If we are on the main thread, just call the function
	if a.impl.isOnMainThread() {
		fn()
		return
	}

	mainThreadFunctionStoreLock.Lock()
	id := generateFunctionStoreID()
	mainThreadFunctionStore[id] = fn
	mainThreadFunctionStoreLock.Unlock()
	// Call platform specific dispatch function
	a.impl.dispatchOnMainThread(id)
}

func OpenFileDialogWithOptions(options *OpenFileDialogOptions) *OpenFileDialogStruct {
	result := OpenFileDialog()
	result.SetOptions(options)
	return result
}

func SaveFileDialogWithOptions(s *SaveFileDialogOptions) *SaveFileDialogStruct {
	result := SaveFileDialog()
	result.SetOptions(s)
	return result
}

func (a *App) dispatchEventToWindows(event *WailsEvent) {
	for _, window := range a.windows {
		window.dispatchWailsEvent(event)
	}
}

func (a *App) Hide() {
	if a.impl != nil {
		a.impl.hide()
	}
}

func (a *App) Show() {
	if a.impl != nil {
		a.impl.show()
	}
}

func (a *App) Log(message *logger.Message) {
	a.log.Log(message)
}

func (a *App) RegisterContextMenu(name string, menu *Menu) {
	a.contextMenusLock.Lock()
	defer a.contextMenusLock.Unlock()
	a.contextMenus[name] = menu
}

func (a *App) getContextMenu(name string) (*Menu, bool) {
	a.contextMenusLock.Lock()
	defer a.contextMenusLock.Unlock()
	menu, ok := a.contextMenus[name]
	return menu, ok

}

func (a *App) OnWindowCreation(callback func(window *WebviewWindow)) {
	a.windowCreatedCallbacks = append(a.windowCreatedCallbacks, callback)
}

func (a *App) GetWindowByName(name string) *WebviewWindow {
	a.windowsLock.Lock()
	defer a.windowsLock.Unlock()
	for _, window := range a.windows {
		if window.Name() == name {
			return window
		}
	}
	return nil
}

func (a *App) runOrDeferToAppRun(r runnable) {
	a.runLock.Lock()
	running := a.running
	if !running {
		a.pendingRun = append(a.pendingRun, r)
	}
	a.runLock.Unlock()

	if running {
		r.run()
	}
}

func invokeSync(fn func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	globalApplication.dispatchOnMainThread(func() {
		defer func() {
			if err := recover(); err != nil {
				processPanic(err)
			}
		}()
		fn()
		wg.Done()
	})
	wg.Wait()
}

func invokeSyncWithResult[T any](fn func() T) (res T) {
	var wg sync.WaitGroup
	wg.Add(1)
	globalApplication.dispatchOnMainThread(func() {
		defer func() {
			if err := recover(); err != nil {
				processPanic(err)
			}
		}()
		res = fn()
		wg.Done()
	})
	wg.Wait()
	return res
}

func invokeSyncWithError(fn func() error) (err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	globalApplication.dispatchOnMainThread(func() {
		defer func() {
			if err := recover(); err != nil {
				processPanic(err)
			}
		}()
		err = fn()
		wg.Done()
	})
	wg.Wait()
	return
}

func invokeSyncWithResultAndError[T any](fn func() (T, error)) (res T, err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	globalApplication.dispatchOnMainThread(func() {
		defer func() {
			if err := recover(); err != nil {
				processPanic(err)
			}
		}()
		res, err = fn()
		wg.Done()
	})
	wg.Wait()
	return res, err
}
