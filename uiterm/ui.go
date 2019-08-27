package uiterm

import (
	"errors"
"strings"
	"sync/atomic"

	"github.com/nsf/termbox-go"
)

type KeyListener func(ui *Ui, key Key)
type CommandListener func(ui *Ui, cmd string)

type UiManager interface {
	OnUiInitialize(ui *Ui)
	OnUiDoneInitialize(ui *Ui)
	OnUiResize(ui *Ui, width, height int)
}

type Ui struct {
	Fg, Bg Attribute

	close   chan bool
	manager UiManager

	drawCount     int32
	elements      map[string]*uiElement
	activeElement *uiElement

	keyListeners map[Key][]KeyListener
	commandListeners map[string][]CommandListener
}

type uiElement struct {
	Name           string
	X0, Y0, X1, Y1 int
	View           View
}

func New(manager UiManager) *Ui {
	ui := &Ui{
		close:        make(chan bool, 10),
		elements:     make(map[string]*uiElement),
		manager:      manager,
		keyListeners: make(map[Key][]KeyListener),
		commandListeners: make(map[string][]CommandListener),
	}
	return ui
}

func (ui *Ui) Close() {
	if termbox.IsInit {
		ui.close <- true
	}
}

func (ui *Ui) Refresh() {
	if termbox.IsInit {
		ui.beginDraw()
		defer ui.endDraw()

		termbox.Clear(termbox.Attribute(ui.Fg), termbox.Attribute(ui.Bg))
		termbox.HideCursor()
		for _, element := range ui.elements {
			element.View.uiDraw()
		}
	}
}

func (ui *Ui) beginDraw() {
	atomic.AddInt32(&ui.drawCount, 1)
}

func (ui *Ui) endDraw() {
	if count := atomic.AddInt32(&ui.drawCount, -1); count == 0 {
		termbox.Flush()
	}
}

func (ui *Ui) Active() string {
	return ui.activeElement.Name
}

func (ui *Ui) SetActive(name string) {
	element, _ := ui.elements[name]
	if ui.activeElement != nil {
		ui.activeElement.View.uiSetActive(false)
	}
	ui.activeElement = element
	if element != nil {
		element.View.uiSetActive(true)
	}
}

func (ui *Ui) Run(cmds chan string) error {
	if termbox.IsInit {
		return nil
	}
	if err := termbox.Init(); err != nil {
		return nil
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt)

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	ui.manager.OnUiInitialize(ui)
	ui.manager.OnUiDoneInitialize(ui)
	width, height := termbox.Size()
	ui.manager.OnUiResize(ui, width, height)
	ui.Refresh()

	for {
		select {
		case <-ui.close:
			return nil
  case cmd := <-cmds:
ui.onCommandEvent(cmd)
		case event := <-events:
			switch event.Type {
			case termbox.EventResize:
				ui.manager.OnUiResize(ui, event.Width, event.Height)
				ui.Refresh()
			case termbox.EventKey:
var k = uint32(event.Key)
				if event.Ch != 0 && event.Mod!=0 {
k=uint32(event.Ch)
}
				if event.Ch != 0 && event.Mod==0 {
					ui.onCharacterEvent(event.Ch)
				} else {
if event.Mod > 0 {
k = k + (uint32(event.Mod) << 16)
}
					ui.onKeyEvent(Key(k))
				}
			}
		}
	}
}

func (ui *Ui) onCharacterEvent(ch rune) {
	if ui.activeElement != nil {
		ui.activeElement.View.uiCharacterEvent(ch)
	}
}

func (ui *Ui) onKeyEvent(key Key) {
	if ui.keyListeners[key] != nil {
		for _, listener := range ui.keyListeners[key] {
			listener(ui, key)
		}
	}
	if ui.activeElement != nil {
		ui.activeElement.View.uiKeyEvent(key)
	}
}

func (ui *Ui) onCommandEvent(cmd string) {
 ta := strings.SplitN(cmd," ",2)
t := ta[0]
rest := ""
if len(ta)==2 {
rest=ta[1]
}
	if ui.commandListeners[t] != nil {
		for _, listener := range ui.commandListeners[t] {
			listener(ui, rest)
		}
	}
	if ui.commandListeners["*"] != nil {
		for _, listener := range ui.commandListeners["*"] {
			listener(ui, cmd)
		}
	}
//	if ui.activeElement != nil {
//		ui.activeElement.View.uiKeyEvent(key)
//	}
}

func (ui *Ui) Add(name string, view View) error {
	if _, ok := ui.elements[name]; ok {
		return errors.New("view already exists")
	}
	ui.elements[name] = &uiElement{
		Name: name,
		View: view,
	}
	view.uiInitialize(ui)
	return nil
}

func (ui *Ui) SetBounds(name string, x0, y0, x1, y1 int) error {
	element, ok := ui.elements[name]
	if !ok {
		return errors.New("view does not exist")
	}
	element.X0, element.Y0, element.X1, element.Y1 = x0, y0, x1, y1
	element.View.uiSetBounds(x0, y0, x1, y1)
	return nil
}

func (ui *Ui) AddKeyListener(listener KeyListener, key *Key) {
if key!=nil {
	ui.keyListeners[*key] = append(ui.keyListeners[*key], listener)
}
}

func (ui *Ui) AddCommandListener(listener CommandListener, cmd string) {
//if cmd!=nil {
	ui.commandListeners[cmd] = append(ui.commandListeners[cmd], listener)
//}
}

