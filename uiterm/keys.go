package uiterm
//go:generate enumer -type=Key -trimprefix=Key -yaml -json -transform=snake

/*
 * Source: https://godoc.org/github.com/nsf/termbox-go
 */

type Key uint32

const (
	KeyF1 Key = 0xFFFF - iota
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgup
	KeyPgdn
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight

	MouseLeft
	MouseMiddle
	MouseRight
)

const (
	KeyCtrlTilde      Key = 0x00
	KeyCtrl2          Key = 0x00
	KeyCtrlSpace      Key = 0x00
	KeyCtrlA          Key = 0x01
	KeyCtrlB          Key = 0x02
	KeyCtrlC          Key = 0x03
	KeyCtrlD          Key = 0x04
	KeyCtrlE          Key = 0x05
	KeyCtrlF          Key = 0x06
	KeyCtrlG          Key = 0x07
	KeyBackspace      Key = 0x08
	KeyCtrlH          Key = 0x08
	KeyTab            Key = 0x09
	KeyCtrlI          Key = 0x09
	KeyCtrlJ          Key = 0x0A
	KeyCtrlK          Key = 0x0B
	KeyCtrlL          Key = 0x0C
	KeyEnter          Key = 0x0D
	KeyCtrlM          Key = 0x0D
	KeyCtrlN          Key = 0x0E
	KeyCtrlO          Key = 0x0F
	KeyCtrlP          Key = 0x10
	KeyCtrlQ          Key = 0x11
	KeyCtrlR          Key = 0x12
	KeyCtrlS          Key = 0x13
	KeyCtrlT          Key = 0x14
	KeyCtrlU          Key = 0x15
	KeyCtrlV          Key = 0x16
	KeyCtrlW          Key = 0x17
	KeyCtrlX          Key = 0x18
	KeyCtrlY          Key = 0x19
	KeyCtrlZ          Key = 0x1A
	KeyEsc            Key = 0x1B
	KeyCtrlLsqBracket Key = 0x1B
	KeyCtrl3          Key = 0x1B
	KeyCtrl4          Key = 0x1C
	KeyCtrlBackslash  Key = 0x1C
	KeyCtrl5          Key = 0x1D
	KeyCtrlRsqBracket Key = 0x1D
	KeyCtrl6          Key = 0x1E
	KeyCtrl7          Key = 0x1F
	KeyCtrlSlash      Key = 0x1F
	KeyCtrlUnderscore Key = 0x1F
	KeySpace          Key = 0x20
	KeyBackspace2     Key = 0x7F
	KeyCtrl8          Key = 0x7F
)
//##altkeys##

const(
KeyAltF1 Key = KeyF1 + (1<<16)
KeyAltF2 Key = KeyF2 + (1<<16)
KeyAltF3 Key = KeyF3 + (1<<16)
KeyAltF4 Key = KeyF4 + (1<<16)
KeyAltF5 Key = KeyF5 + (1<<16)
KeyAltF6 Key = KeyF6 + (1<<16)
KeyAltF7 Key = KeyF7 + (1<<16)
KeyAltF8 Key = KeyF8 + (1<<16)
KeyAltF9 Key = KeyF9 + (1<<16)
KeyAltF10 Key = KeyF10 + (1<<16)
KeyAltF11 Key = KeyF11 + (1<<16)
KeyAltF12 Key = KeyF12 + (1<<16)
KeyAltInsert Key = KeyInsert + (1<<16)
KeyAltDelete Key = KeyDelete + (1<<16)
KeyAltHome Key = KeyHome + (1<<16)
KeyAltEnd Key = KeyEnd + (1<<16)
KeyAltPgup Key = KeyPgup + (1<<16)
KeyAltPgdn Key = KeyPgdn + (1<<16)
KeyAltArrowUp Key = KeyArrowUp + (1<<16)
KeyAltArrowDown Key = KeyArrowDown + (1<<16)
KeyAltArrowLeft Key = KeyArrowLeft + (1<<16)
KeyAltArrowRight Key = KeyArrowRight + (1<<16)
KeyAltCtrlTilde Key = KeyCtrlTilde + (1<<16)
KeyAltCtrl2 Key = KeyCtrl2 + (1<<16)
KeyAltCtrlSpace Key = KeyCtrlSpace + (1<<16)
KeyAltCtrlA Key = KeyCtrlA + (1<<16)
KeyAltCtrlB Key = KeyCtrlB + (1<<16)
KeyAltCtrlC Key = KeyCtrlC + (1<<16)
KeyAltCtrlD Key = KeyCtrlD + (1<<16)
KeyAltCtrlE Key = KeyCtrlE + (1<<16)
KeyAltCtrlF Key = KeyCtrlF + (1<<16)
KeyAltCtrlG Key = KeyCtrlG + (1<<16)
KeyAltBackspace Key = KeyBackspace + (1<<16)
KeyAltCtrlH Key = KeyCtrlH + (1<<16)
KeyAltTab Key = KeyTab + (1<<16)
KeyAltCtrlI Key = KeyCtrlI + (1<<16)
KeyAltCtrlJ Key = KeyCtrlJ + (1<<16)
KeyAltCtrlK Key = KeyCtrlK + (1<<16)
KeyAltCtrlL Key = KeyCtrlL + (1<<16)
KeyAltEnter Key = KeyEnter + (1<<16)
KeyAltCtrlM Key = KeyCtrlM + (1<<16)
KeyAltCtrlN Key = KeyCtrlN + (1<<16)
KeyAltCtrlO Key = KeyCtrlO + (1<<16)
KeyAltCtrlP Key = KeyCtrlP + (1<<16)
KeyAltCtrlQ Key = KeyCtrlQ + (1<<16)
KeyAltCtrlR Key = KeyCtrlR + (1<<16)
KeyAltCtrlS Key = KeyCtrlS + (1<<16)
KeyAltCtrlT Key = KeyCtrlT + (1<<16)
KeyAltCtrlU Key = KeyCtrlU + (1<<16)
KeyAltCtrlV Key = KeyCtrlV + (1<<16)
KeyAltCtrlW Key = KeyCtrlW + (1<<16)
KeyAltCtrlX Key = KeyCtrlX + (1<<16)
KeyAltCtrlY Key = KeyCtrlY + (1<<16)
KeyAltCtrlZ Key = KeyCtrlZ + (1<<16)
KeyAltEsc Key = KeyEsc + (1<<16)
KeyAltCtrlLsqBracket Key = KeyCtrlLsqBracket + (1<<16)
KeyAltCtrl3 Key = KeyCtrl3 + (1<<16)
KeyAltCtrl4 Key = KeyCtrl4 + (1<<16)
KeyAltCtrlBackslash Key = KeyCtrlBackslash + (1<<16)
KeyAltCtrl5 Key = KeyCtrl5 + (1<<16)
KeyAltCtrlRsqBracket Key = KeyCtrlRsqBracket + (1<<16)
KeyAltCtrl6 Key = KeyCtrl6 + (1<<16)
KeyAltCtrl7 Key = KeyCtrl7 + (1<<16)
KeyAltCtrlSlash Key = KeyCtrlSlash + (1<<16)
KeyAltCtrlUnderscore Key = KeyCtrlUnderscore + (1<<16)
KeyAltSpace Key = KeySpace + (1<<16)
KeyAltBackspace2 Key = KeyBackspace2 + (1<<16)
KeyAltCtrl8 Key = KeyCtrl8 + (1<<16)
KeyAltA Key = 0x61 + (1<<16)
KeyAltB Key = 0x62 + (1<<16)
KeyAltC Key = 0x63 + (1<<16)
KeyAltD Key = 0x64 + (1<<16)
KeyAltE Key = 0x65 + (1<<16)
KeyAltF Key = 0x66 + (1<<16)
KeyAltG Key = 0x67 + (1<<16)
KeyAltH Key = 0x68 + (1<<16)
KeyAltI Key = 0x69 + (1<<16)
KeyAltJ Key = 0x6a + (1<<16)
KeyAltK Key = 0x6b + (1<<16)
KeyAltL Key = 0x6c + (1<<16)
KeyAltM Key = 0x6d + (1<<16)
KeyAltN Key = 0x6e + (1<<16)
KeyAltO Key = 0x6f + (1<<16)
KeyAltP Key = 0x70 + (1<<16)
KeyAltQ Key = 0x71 + (1<<16)
KeyAltR Key = 0x72 + (1<<16)
KeyAltS Key = 0x73 + (1<<16)
KeyAltT Key = 0x74 + (1<<16)
KeyAltU Key = 0x75 + (1<<16)
KeyAltV Key = 0x76 + (1<<16)
KeyAltW Key = 0x77 + (1<<16)
KeyAltX Key = 0x78 + (1<<16)
KeyAltY Key = 0x79 + (1<<16)
KeyAltZ Key = 0x7a + (1<<16)
)
