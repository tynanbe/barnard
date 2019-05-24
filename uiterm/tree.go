package uiterm

import (
	"strings"

	"github.com/nsf/termbox-go"
)

type TreeItem interface {
	TreeItemStyle(fg, bg Attribute, active bool) (Attribute, Attribute)
	String() string
}

type renderedTreeItem struct {
	//String string
	Level int
	Item  TreeItem
}

type Tree struct {
	Fg, Bg    Attribute
	Generator func(item TreeItem) []TreeItem
	KeyListener  func(ui *Ui, tree *Tree, item TreeItem, mod Modifier, key Key)
	CharacterListener  func(ui *Ui, tree *Tree, item TreeItem, chr rune)

	lines      []renderedTreeItem
	activeLine int

	ui             *Ui
	active         bool
	x0, y0, x1, y1 int
}

func bounded(i, lower, upper int) int {
	if i < lower {
		return lower
	}
	if i > upper {
		return upper
	}
	return i
}

func (t *Tree) uiInitialize(ui *Ui) {
	t.ui = ui
}

func (t *Tree) uiSetActive(active bool) {
	t.active = active
	t.uiDraw()
}

func (t *Tree) uiSetBounds(x0, y0, x1, y1 int) {
	t.x0 = x0
	t.y0 = y0
	t.x1 = x1
	t.y1 = y1
	t.uiDraw()
}

func (t *Tree) Rebuild() {
	if t.Generator == nil {
		t.lines = []renderedTreeItem{}
		return
	}

	lines := []renderedTreeItem{}
	for _, item := range t.Generator(nil) {
		children := t.rebuild_rec(item, 0)
		if children != nil {
			lines = append(lines, children...)
		}
	}
	t.lines = lines
	t.SetActiveLine(0,true)
	t.uiDraw()
}

func (t *Tree) rebuild_rec(parent TreeItem, level int) []renderedTreeItem {
	if parent == nil {
		return nil
	}
	lines := []renderedTreeItem{
		renderedTreeItem{
			Level: level,
			Item:  parent,
		},
	}
	for _, item := range t.Generator(parent) {
		children := t.rebuild_rec(item, level+1)
		if children != nil {
			lines = append(lines, children...)
		}
	}
	return lines
}


func (t *Tree) uiDraw() {
	t.ui.beginDraw()
	defer t.ui.endDraw()

	if t.lines == nil {
		t.Rebuild()
	}

if t.y1-t.y0 <= 0 {
return
}

var	line = t.activeLine
var height=t.y1-t.y0
var startline=0
var total = len(t.lines)
//I'd welcome a better algorithm for this; for that matter, I'd love a book or reference for all sorts of GUI algorithms.
//if (startline+height) < line {
for startline=0; (startline+height) <= line; startline+=height {
}
//}
if startline+height >= total {
var rem=(startline+height)-total
startline-=rem
}
if (startline < 0) {
startline = 0
}
line=startline
	for y := t.y0; y < t.y1; y++ {
		var reader *strings.Reader
		var item TreeItem
		level := 0
		if line < len(t.lines) {
			item = t.lines[line].Item
			level = t.lines[line].Level
			reader = strings.NewReader(item.String())
		}
		for x := t.x0; x < t.x1; x++ {
			var chr rune = ' '
			fg := t.Fg
			bg := t.Bg
			dx := x - t.x0
			dy := y - t.y0
			if reader != nil && level*2 <= dx {
				if ch, _, err := reader.ReadRune(); err == nil {
					chr = ch
					fg, bg = item.TreeItemStyle(fg, bg, t.active && t.activeLine == dy)
				}
			}
			termbox.SetCell(x, y, chr, termbox.Attribute(fg), termbox.Attribute(bg))
		}
		if t.activeLine == (line) {
			termbox.SetCursor(t.x0, y)
		}
		line++
	}
}

func (t *Tree) SetActiveLine(num int, relative bool) {
if relative {
		t.activeLine = bounded(t.activeLine+num, 0, len(t.lines)-1)
} else {
		t.activeLine = bounded(num, 0, len(t.lines)-1)
}
}

func (t *Tree) uiKeyEvent(mod Modifier, key Key) {
var runHandler=true
	switch key {
	case KeyArrowUp:
t.SetActiveLine(-1, true)
runHandler=false
	case KeyArrowDown:
t.SetActiveLine(1, true)
runHandler=false
}
if runHandler==true && t.KeyListener != nil {
			t.KeyListener(t.ui, t, t.lines[t.activeLine].Item, mod, key)
		}
	t.uiDraw()
}

func (t *Tree) uiCharacterEvent(ch rune) {
if (t.KeyListener!=nil) {
t.CharacterListener(t.ui, t, t.lines[t.activeLine].Item, ch)
}
}
