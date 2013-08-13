// angora project events.go
package angora

import (
	. "github.com/arbaal/mathgl"
	glfw "github.com/go-gl/glfw3"
)

// Input Events
//
type TextEvent chan rune

// Control Events
//
type ControlEventType int

const (
	CET_NONE = ControlEventType(iota)
	CET_ON_CLOSE
	CET_KEY
	CET_MOUSE_BUTTON
	CET_MOUSE_MOVE
	CET_MOUSE_WHEEL
	CET_RESIZE
)

type ControlInfo struct {
	Ctype  ControlEventType
	Key    glfw.Key
	Button glfw.MouseButton
	Action glfw.Action
	Mods   glfw.ModifierKey
	*Vec2
}
type ControlEvent chan ControlInfo

// Output Events
//
type Presenter interface {
	ID() int
	Paint(context *GraphicsContext) error
}
type PresentEvent chan Presenter

type RemoveInfo struct {
	ID  int
	err error
}
type RemoveEvent chan RemoveInfo
