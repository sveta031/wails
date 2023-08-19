package application

import (
	"net/http"
)

const (
	ApplicationHide = 0
	ApplicationShow = 1
	ApplicationQuit = 2
)

var applicationMethodNames = map[int]string{
	ApplicationQuit: "Quit",
	ApplicationHide: "Hide",
	ApplicationShow: "Show",
}

func (m *MessageProcessor) processApplicationMethod(method int, rw http.ResponseWriter, r *http.Request, window Window, params QueryParams) {
	switch method {
	case ApplicationQuit:
		globalApplication.Quit()
		m.ok(rw)
	case ApplicationHide:
		globalApplication.Hide()
		m.ok(rw)
	case ApplicationShow:
		globalApplication.Show()
		m.ok(rw)
	default:
		m.httpError(rw, "Unknown application method: %d", method)
	}

	m.Info("Runtime Call:", "method", "Application."+applicationMethodNames[method])

}
