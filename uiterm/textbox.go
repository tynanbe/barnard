package uiterm

import (
	"strings"
	//	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

type Textbox struct {
	Text   string
	Fg, Bg Attribute

	Input func(ui *Ui, textbox *Textbox, text string)

	ui             *Ui
	active         bool
	x0, y0, x1, y1 int
	pos            int
}

func (t *Textbox) uiInitialize(ui *Ui) {
	t.ui = ui
	t.pos = 0
}

func (t *Textbox) uiSetActive(active bool) {
	t.active = active
	t.uiDraw()
}

func (t *Textbox) uiSetBounds(x0, y0, x1, y1 int) {
	t.x0 = x0
	t.y0 = y0
	t.x1 = x1
	t.y1 = y1
	t.uiDraw()
}

func (t *Textbox) uiDraw() {
	t.ui.beginDraw()
	defer t.ui.endDraw()

	reader := strings.NewReader(t.Text)
	if t.pos < 0 {
		t.pos = 0
	}
	if t.pos > len(t.Text) {
		t.pos = len(t.Text)
	}
	for y := t.y0; y < t.y1; y++ {
		for x := t.x0; x < t.x1; x++ {
			var chr rune
			if ch, _, err := reader.ReadRune(); err != nil {
				chr = ' '
			} else {
				chr = ch
			}
			termbox.SetCell(x, y, chr, termbox.Attribute(t.Fg), termbox.Attribute(t.Bg))
		}
	}
	if t.active {
		var x = 0
		var y = 0
		var idx = -1
		var flag = false
		for y = t.y0; y < t.y1; y++ {
			for x = t.x0; x < t.x1; x++ {
				idx += 1
				if idx == t.pos {
					flag = true
				}
				if flag == true {
					break
				}
			}
			if flag == true {
				break
			}
		}
		termbox.SetCursor(x, y)
	}
}

func (t *Textbox) uiKeyEvent(mod Modifier, key Key) {
	redraw := false
	switch key {
	case KeyArrowLeft:
		t.pos -= 1
		redraw = true
	case KeyArrowRight:
		t.pos += 1
		redraw = true
	case KeyCtrlC:
		t.Text = ""
		t.pos = 0
		redraw = true
	case KeyEnter:
		if t.Input != nil {
			t.Input(t.ui, t, t.Text)
		}
		t.Text = ""
		t.pos = 0
		redraw = true
	case KeySpace:
		t.uiCharacterEvent(' ')
	case KeyBackspace:
	case KeyBackspace2:
		if len(t.Text) > 0 {
			if t.pos > 0 {
				t.Text = t.Text[:t.pos-1] + t.Text[t.pos:]
				t.pos -= 1
			}
		}
		//			if r, size := utf8.DecodeLastRuneInString(t.Text); r != utf8.RuneError {
		//				t.Text = t.Text[:len(t.Text)-size]
		//t.pos-=size
		redraw = true
		//			}
		//		}
	}
	if redraw {
		t.uiDraw()
	}
}

func (t *Textbox) uiCharacterEvent(chr rune) {
	var s = string(chr)
	t.Text = t.Text[:t.pos] + s + t.Text[t.pos:]
	t.pos += len(s)
	t.uiDraw()
}
