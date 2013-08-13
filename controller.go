// angora project filters.go
package angora

import ()

type ControlHandlerInfo struct {
	evt ControlInfo
	f   func(ControlInfo)
}

type TextHandlerInfo struct {
	evt rune
	f   func(rune)
}

type GuiController struct {
	Action chan ControlHandlerInfo
	Text   chan TextHandlerInfo
}

func (c *GuiController) Run() {
	for {
		select {
		case chi := <-c.Action:
			chi.f(chi.evt)
			if chi.evt.Ctype == CET_ON_CLOSE || chi.evt.Ctype == CET_NONE {
				c.Action = nil // nothing more to process
				c.Text = nil   // don't allow any more receives of text
			}
		case thi := <-c.Text:
			thi.f(thi.evt)
		}
	}
}

func (c *GuiController) ControlHandler(in ControlEvent,
	m ControlInfo, f func(evt ControlInfo)) {
	for evt := range in {
		x, y := evt.X, evt.Y
		evt.X, evt.Y = 0.0, 0.0
		if evt == m {
			evt.X, evt.Y = x, y
			c.Action <- ControlHandlerInfo{evt, f}
		}
	}
}

type ControlActionMap map[ControlInfo]bool

func (c *GuiController) ControlMultiHandler(in ControlEvent,
	m ControlActionMap, f func(evt ControlInfo)) {
	for evt := range in {
		x, y := evt.X, evt.Y
		evt.X, evt.Y = 0.0, 0.0
		if m[evt] {
			evt.X, evt.Y = x, y
			c.Action <- ControlHandlerInfo{evt, f}
		}
	}
}

type TextExcludeMap map[rune]bool

func (c *GuiController) TextHandler(in TextEvent,
	m TextExcludeMap, f func(rune)) {
	for evt := range in {
		if !m[evt] {
			c.Text <- TextHandlerInfo{evt, f}
		}
	}
}
