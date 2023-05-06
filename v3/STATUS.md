# Status

Status of features in v3. Incomplete - please add as you see fit.

## Application

Application interface methods

| Method                                                        | Windows | Linux | Mac | Notes |
|---------------------------------------------------------------|---------|-------|-----|-------|
| run() error                                                   |         |       | Y   |       |
| destroy()                                                     |         |       | Y   |       |
| setApplicationMenu(menu *Menu)                                |         |       | Y   |       |
| name() string                                                 |         |       | Y   |       |
| getCurrentWindowID() uint                                     |         |       | Y   |       |
| showAboutDialog(name string, description string, icon []byte) |         |       | Y   |       |
| setIcon(icon []byte)                                          |         |       | Y   |       |
| on(id uint)                                                   |         |       | Y   |       |
| dispatchOnMainThread(fn func())                               | Y       |       | Y   |       |
| hide()                                                        | Y       |       | Y   |       |
| show()                                                        | Y       |       | Y   |       |
| getPrimaryScreen() (*Screen, error)                           |         |       | Y   |       |
| getScreens() ([]*Screen, error)                               |         |       | Y   |       |

## Webview Window

Webview Window Interface Methods

| Method                                             | Windows | Linux | Mac | Notes                                    |    
|----------------------------------------------------|---------|-------|-----|------------------------------------------|    
| center()                                           | Y       |       | Y   |                                          |
| close()                                            |         |       | Y   |                                          |
| destroy()                                          |         |       | Y   |                                          |
| execJS(js string)                                  |         |       | Y   |                                          |
| forceReload()                                      |         |       | Y   |                                          |
| fullscreen()                                       | Y       |       | Y   |                                          |
| getScreen() (*Screen, error)                       |         |       | Y   |                                          |
| getZoom() float64                                  |         |       | Y   |                                          |
| height() int                                       | Y       |       | Y   |                                          |
| hide()                                             | Y       |       | Y   |                                          |
| isFullscreen() bool                                | Y       |       | Y   |                                          |
| isMaximised() bool                                 | Y       |       | Y   |                                          |
| isMinimised() bool                                 | Y       |       | Y   |                                          |
| maximise()                                         | Y       |       | Y   |                                          |
| minimise()                                         | Y       |       | Y   |                                          |
| nativeWindowHandle() (uintptr, error)              | Y       |       |     |                                          |
| on(eventID uint)                                   |         |       | Y   |                                          |
| openContextMenu(menu *Menu, data *ContextMenuData) |         |       | Y   |                                          |
| position() (int, int)                              | Y       |       | Y   |                                          |
| reload()                                           |         |       | Y   |                                          |
| run()                                              | Y       |       | Y   |                                          |
| setAlwaysOnTop(alwaysOnTop bool)                   | Y       |       | Y   |                                          |
| setBackgroundColour(color RGBA)                    | Y       |       | Y   |                                          |
| setFrameless(bool)                                 |         |       | Y   |                                          |
| setFullscreenButtonEnabled(enabled bool)           | -       |       | Y   | There is no fullscreen button in Windows |
| setHTML(html string)                               |         |       | Y   |                                          |
| setMaxSize(width, height int)                      | Y       |       | Y   |                                          |
| setMinSize(width, height int)                      | Y       |       | Y   |                                          |
| setPosition(x int, y int)                          | Y       |       | Y   |                                          |
| setResizable(resizable bool)                       | Y       |       | Y   |                                          |
| setSize(width, height int)                         | Y       |       | Y   |                                          |
| setTitle(title string)                             | Y       |       | Y   |                                          |
| setURL(url string)                                 |         |       | Y   |                                          |
| setZoom(zoom float64)                              |         |       | Y   |                                          |
| show()                                             | Y       |       | Y   |                                          |
| size() (int, int)                                  | Y       |       | Y   |                                          |
| toggleDevTools()                                   |         |       | Y   |                                          |
| unfullscreen()                                     | Y       |       | Y   |                                          |
| unmaximise()                                       | Y       |       | Y   |                                          |
| unminimise()                                       | Y       |       | Y   |                                          |
| width() int                                        | Y       |       | Y   |                                          |
| zoom()                                             |         |       | Y   |                                          |
| zoomIn()                                           |         |       | Y   |                                          |
| zoomOut()                                          |         |       | Y   |                                          |
| zoomReset()                                        |         |       | Y   |                                          |

## Runtime

### Application

| Feature | Windows | Linux | Mac | Notes |
|---------|---------|-------|-----|-------|
| Quit    |         |       | Y   |       |
| Hide    | Y       |       | Y   |       |
| Show    | Y       |       | Y   |       |

### Dialogs

| Feature  | Windows | Linux | Mac | Notes |
|----------|---------|-------|-----|-------|
| Info     |         |       | Y   |       |
| Warning  |         |       | Y   |       |
| Error    |         |       | Y   |       |
| Question |         |       | Y   |       |
| OpenFile |         |       | Y   |       |
| SaveFile |         |       | Y   |       |

### Clipboard

| Feature | Windows | Linux | Mac | Notes |
|---------|---------|-------|-----|-------|
| SetText |         |       | Y   |       |
| Text    |         |       | Y   |       |

### ContextMenu

| Feature         | Windows | Linux | Mac | Notes |
|-----------------|---------|-------|-----|-------|
| OpenContextMenu |         |       | Y   |       |

### Screens

| Feature    | Windows | Linux | Mac | Notes |
|------------|---------|-------|-----|-------|
| GetAll     | Y       |       | Y   |       |
| GetPrimary |         |       | Y   |       |
| GetCurrent |         |       | Y   |       |

### Window

| Feature             | Windows | Linux | Mac | Notes                                                                                |
|---------------------|---------|-------|-----|--------------------------------------------------------------------------------------|
| SetTitle            |         |       | Y   |                                                                                      |
| SetSize             |         |       | Y   |                                                                                      |
| Size                |         |       | Y   |                                                                                      |
| SetPosition         |         |       | Y   |                                                                                      |
| Position            |         |       | Y   |                                                                                      |
| FullScreen          |         |       | Y   |                                                                                      |
| UnFullscreen        |         |       | Y   |                                                                                      |
| Minimise            |         |       | Y   |                                                                                      |
| UnMinimise          |         |       | Y   |                                                                                      |
| Maximise            |         |       | Y   |                                                                                      |
| UnMaximise          |         |       | Y   |                                                                                      |
| Show                |         |       | Y   |                                                                                      |
| Hide                |         |       | Y   |                                                                                      |
| Center              |         |       | Y   |                                                                                      |
| SetBackgroundColour |         |       | Y   | https://github.com/MicrosoftEdge/WebView2Feedback/issues/1621#issuecomment-938234294 |
| SetAlwaysOnTop      |         |       | Y   |                                                                                      |
| SetResizable        |         |       | Y   |                                                                                      |
| SetMinSize          |         |       | Y   |                                                                                      |
| SetMaxSize          |         |       | Y   |                                                                                      |
| Width               |         |       | Y   |                                                                                      |
| Height              |         |       | Y   |                                                                                      |
| ZoomIn              |         |       | Y   | Increase view scale                                                                  |
| ZoomOut             |         |       | Y   | Decrease view scale                                                                  |
| ZoomReset           |         |       | Y   | Reset view scale                                                                     |
| GetZoom             |         |       | Y   | Get current view scale                                                               |
| SetZoom             |         |       | Y   | Set view scale                                                                       |
| Screen              |         |       | Y   | Get screen for window                                                                |

### Window Options

A 'Y' in the table below indicates that the option has been tested and is applied when the window is created.
An 'X' indicates that the option is not supported by the platform.

| Feature                         | Windows | Linux | Mac | Notes                                      |
|---------------------------------|---------|-------|-----|--------------------------------------------|
| Name                            |         |       |     |                                            |
| Title                           | Y       |       |     |                                            |
| Width                           | Y       |       |     |                                            |
| Height                          | Y       |       |     |                                            |
| AlwaysOnTop                     | Y       |       |     |                                            |
| URL                             |         |       |     |                                            |
| DisableResize                   | Y       |       |     |                                            |
| Frameless                       |         |       |     |                                            |
| MinWidth                        | Y       |       |     |                                            |
| MinHeight                       | Y       |       |     |                                            |
| MaxWidth                        | Y       |       |     |                                            |
| MaxHeight                       | Y       |       |     |                                            |
| StartState                      | Y       |       |     |                                            |
| Mac                             | -       | -     |     |                                            |
| BackgroundType                  |         |       |     | Acrylic seems to work but the others don't |
| BackgroundColour                | Y       |       |     |                                            |
| HTML                            |         |       |     |                                            |
| JS                              |         |       |     |                                            |
| CSS                             |         |       |     |                                            |
| X                               |         |       |     |                                            |
| Y                               |         |       |     |                                            |
| HideOnClose                     |         |       |     |                                            |
| FullscreenButtonEnabled         |         |       |     |                                            |
| Hidden                          |         |       |     |                                            |
| EnableFraudulentWebsiteWarnings |         |       |     |                                            |
| Zoom                            |         |       |     |                                            |
| EnableDragAndDrop               |         |       |     |                                            |
| Windows                         |         | -     | -   |                                            |

### Log

To log or not to log? System logger vs custom logger.

## Menu

| Event                    | Windows | Linux | Mac | Notes |
|--------------------------|---------|-------|-----|-------|
| Default Application Menu |         |       | Y   |       |

## Tray Menus

| Feature            | Windows | Linux | Mac | Notes |
|--------------------|---------|-------|-----|-------|
| Icon               |         |       | Y   |       |
| Label              |         |       | Y   |       |
| Label (ANSI Codes) |         |       |     |       |
| Menu               |         |       | Y   |       |

## Cross Platform Events

Mapping native events to cross-platform events.

| Event                    | Windows | Linux | Mac             | Notes |
|--------------------------|---------|-------|-----------------|-------|
| WindowWillClose          |         |       | WindowWillClose |       |
| WindowDidClose           |         |       |                 |       |
| WindowDidResize          |         |       |                 |       |
| WindowDidHide            |         |       |                 |       |
| ApplicationWillTerminate |         |       |                 |       |

... Add more

## Bindings Generation

TBD

## Models Generation

TBD

## Task file

TBD

## Theme

| Plugin | Windows | Linux | Mac | Notes |
|--------|---------|-------|-----|-------|
| Dark   | Y       |       |     |       |
| Light  | Y       |       |     |       |
| System | Y       |       |     |       |

## NSIS Installer

TBD

## Templates

TBD

## Plugins

Built-in plugin support:

| Plugin          | Windows | Linux | Mac | Notes |
|-----------------|---------|-------|-----|-------|
| Browser         |         |       | Y   |       |
| KV Store        |         |       | Y   |       |
| Log             |         |       | Y   |       |
| Single Instance |         |       | Y   |       |
| SQLite          |         |       | Y   |       |
| Start at login  |         |       | Y   |       |
| Server          |         |       |     |       |

## Packaging

|                 | Windows | Linux | Mac | Notes | 
|-----------------|---------|-------|-----|-------| 
| Icon Generation |         |       | Y   |       | 
| Icon Embedding  |         |       | Y   |       | 
| Info.plist      |         |       | Y   |       | 
| NSIS Installer  |         |       | -   |       | 
| Mac bundle      |         |       | Y   |       | 
| Windows exe     |         |       | -   |       | 

## Frameless Windows

| Feature | Windows | Linux | Mac | Notes |
|---------|---------|-------|-----|-------|
| Resize  |         |       |     |       |
| Drag    |         |       |     |       |

## Mac Specific

- [x] Translucency

## Windows Specific

- [x] Translucency
- [x] Custom Themes

### Windows Options

| Feature                           | Default | Notes                                       |
|-----------------------------------|---------|---------------------------------------------|
| BackdropType                      |         |                                             |
| DisableIcon                       |         |                                             |
| Theme                             |         |                                             |
| CustomTheme                       |         |                                             |
| DisableFramelessWindowDecorations |         |                                             |
| WindowMask                        | nil     | Makes the window the contents of the bitmap |

	// Select the type of translucent backdrop. Requires Windows 11 22621 or later.
	BackdropType BackdropType
	// Disable the icon in the titlebar
	DisableIcon bool
	// Theme. Defaults to SystemDefault which will use whatever the system theme is. The application will follow system theme changes.
	Theme Theme
	// Custom colours for dark/light mode
	CustomTheme *ThemeSettings

	// Disable all window decorations in Frameless mode, which means no "Aero Shadow" and no "Rounded Corner" will be shown.
	// "Rounded Corners" are only available on Windows 11.
	DisableFramelessWindowDecorations bool

	// WindowMask is used to set the window shape. Use a PNG with an alpha channel to create a custom shape.
	WindowMask []byte

## Linux Specific