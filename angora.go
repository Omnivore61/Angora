// angora project angora.go
package angora

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	glfw "github.com/go-gl/glfw3"
	"runtime"
	"unsafe"
)

type GraphicsContext interface {
	// TODO: provide methods
}

type App struct {
	Width, Height    int
	window           *glfw.Window
	OnUpdateCallback func(*App)
	OnRenderCallback func(*App)
}

func pump(app *App) {
	window := app.window
	for !window.ShouldClose() {
		updateAndRender(app)
		do(func() {
			glfw.PollEvents()
		})
	}
	do(func() {
		glfw.Terminate()
	})
}

func (app *App) Run() {
	go pump(app)
	runtime.LockOSThread()
	for f := range mainfunc {
		f()
	}
	err := recover()
	if err != nil {
		glfw.Terminate()
		panic(err) // continue the panic
	}
}

var mainfunc = make(chan func())

func do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}

func glfwErrorCallback(err glfw.ErrorCode, desc string) {
	panic(fmt.Errorf("%v: %v\n", err, desc))
}

func Init(title string, width, height int) (app *App, err error) {
	runtime.LockOSThread()
	app = &App{}
	glfw.SetErrorCallback(glfwErrorCallback)
	if !glfw.Init() {
		err = fmt.Errorf("GLFW3 initialization failure")
		return
	}
	glfw.WindowHint(glfw.Resizable, 0)
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return
	}
	window.MakeContextCurrent()
	if err = gl.Init(); err != nil {
		return
	}
	glfw.SwapInterval(1)
	app.window = window
	app.Width, app.Height = window.GetSize()
	window.SetUserPointer(unsafe.Pointer(app))
	window.SetCharacterCallback(characterCallbackFn)
	window.SetCloseCallback(closeCallbackFn)
	window.SetCursorEnterCallback(cursorEnterCallbackFn)
	window.SetCursorPositionCallback(cursorPositionCallbackFn)
	window.SetFocusCallback(focusCallbackFn)
	window.SetFramebufferSizeCallback(framebufferSizeCallbackFn)
	window.SetIconifyCallback(iconifyCallbackFn)
	window.SetKeyCallback(keyCallbackFn)
	window.SetMouseButtonCallback(mouseButtonCallbackFn)
	window.SetPositionCallback(positionCallbackFn)
	window.SetRefreshCallback(refreshCallbackFn)
	window.SetScrollCallback(scrollCallbackFn)
	window.SetSizeCallback(sizeCallbackFn)

	runtime.UnlockOSThread()
	return
}

func updateAndRender(app *App) {
	app.OnUpdateCallback(app)
	app.OnRenderCallback(app)
	do(func() {
		app.window.SwapBuffers()
	})
}

func characterCallbackFn(w *glfw.Window, char uint) {
	var _, _ = w, char
}

func closeCallbackFn(w *glfw.Window) {
	var _ = w
}

func cursorEnterCallbackFn(w *glfw.Window, entered bool) {
	var _, _ = w, entered
}

func cursorPositionCallbackFn(w *glfw.Window, xpos float64, ypos float64) {
	var _, _, _ = w, xpos, ypos
}

func focusCallbackFn(w *glfw.Window, focused bool) {
	var _, _ = w, focused
}

func framebufferSizeCallbackFn(w *glfw.Window, width int, height int) {
	var _, _, _ = w, width, height
}

func iconifyCallbackFn(w *glfw.Window, iconified bool) {
	var _, _ = w, iconified
}

func keyCallbackFn(w *glfw.Window, key glfw.Key,
	scancode int, action glfw.Action, mods glfw.ModifierKey) {
	var _, _, _, _, _ = w, key, scancode, action, mods
}

func mouseButtonCallbackFn(w *glfw.Window, button glfw.MouseButton,
	action glfw.Action, mod glfw.ModifierKey) {
	var _, _, _, _ = w, button, action, mod
}

func positionCallbackFn(w *glfw.Window, xpos int, ypos int) {
	var _, _, _ = w, xpos, ypos
}

func refreshCallbackFn(w *glfw.Window) {
	var _ = w
}

func scrollCallbackFn(w *glfw.Window, xoff float64, yoff float64) {
	var _, _, _ = w, xoff, yoff
}

func sizeCallbackFn(w *glfw.Window, width int, height int) {
	var _, _, _ = w, width, height
}
