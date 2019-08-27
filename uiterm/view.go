package uiterm

type View interface {
	uiInitialize(ui *Ui)
	uiSetActive(active bool)
	uiSetBounds(x0, y0, x1, y1 int)
	uiDraw()
	uiKeyEvent(key Key)
	uiCharacterEvent(ch rune)
// commandEvent(cmd string)
}
