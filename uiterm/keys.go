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
KeyAltAltF1 Key = KeyAltF1 + (1<<16)
KeyAltAltF2 Key = KeyAltF2 + (1<<16)
KeyAltAltF3 Key = KeyAltF3 + (1<<16)
KeyAltAltF4 Key = KeyAltF4 + (1<<16)
KeyAltAltF5 Key = KeyAltF5 + (1<<16)
KeyAltAltF6 Key = KeyAltF6 + (1<<16)
KeyAltAltF7 Key = KeyAltF7 + (1<<16)
KeyAltAltF8 Key = KeyAltF8 + (1<<16)
KeyAltAltF9 Key = KeyAltF9 + (1<<16)
KeyAltAltF10 Key = KeyAltF10 + (1<<16)
KeyAltAltF11 Key = KeyAltF11 + (1<<16)
KeyAltAltF12 Key = KeyAltF12 + (1<<16)
KeyAltAltInsert Key = KeyAltInsert + (1<<16)
KeyAltAltDelete Key = KeyAltDelete + (1<<16)
KeyAltAltHome Key = KeyAltHome + (1<<16)
KeyAltAltEnd Key = KeyAltEnd + (1<<16)
KeyAltAltPgup Key = KeyAltPgup + (1<<16)
KeyAltAltPgdn Key = KeyAltPgdn + (1<<16)
KeyAltAltArrowUp Key = KeyAltArrowUp + (1<<16)
KeyAltAltArrowDown Key = KeyAltArrowDown + (1<<16)
KeyAltAltArrowLeft Key = KeyAltArrowLeft + (1<<16)
KeyAltAltArrowRight Key = KeyAltArrowRight + (1<<16)
KeyAltAltCtrlTilde Key = KeyAltCtrlTilde + (1<<16)
KeyAltAltCtrl2 Key = KeyAltCtrl2 + (1<<16)
KeyAltAltCtrlSpace Key = KeyAltCtrlSpace + (1<<16)
KeyAltAltCtrlA Key = KeyAltCtrlA + (1<<16)
KeyAltAltCtrlB Key = KeyAltCtrlB + (1<<16)
KeyAltAltCtrlC Key = KeyAltCtrlC + (1<<16)
KeyAltAltCtrlD Key = KeyAltCtrlD + (1<<16)
KeyAltAltCtrlE Key = KeyAltCtrlE + (1<<16)
KeyAltAltCtrlF Key = KeyAltCtrlF + (1<<16)
KeyAltAltCtrlG Key = KeyAltCtrlG + (1<<16)
KeyAltAltBackspace Key = KeyAltBackspace + (1<<16)
KeyAltAltCtrlH Key = KeyAltCtrlH + (1<<16)
KeyAltAltTab Key = KeyAltTab + (1<<16)
KeyAltAltCtrlI Key = KeyAltCtrlI + (1<<16)
KeyAltAltCtrlJ Key = KeyAltCtrlJ + (1<<16)
KeyAltAltCtrlK Key = KeyAltCtrlK + (1<<16)
KeyAltAltCtrlL Key = KeyAltCtrlL + (1<<16)
KeyAltAltEnter Key = KeyAltEnter + (1<<16)
KeyAltAltCtrlM Key = KeyAltCtrlM + (1<<16)
KeyAltAltCtrlN Key = KeyAltCtrlN + (1<<16)
KeyAltAltCtrlO Key = KeyAltCtrlO + (1<<16)
KeyAltAltCtrlP Key = KeyAltCtrlP + (1<<16)
KeyAltAltCtrlQ Key = KeyAltCtrlQ + (1<<16)
KeyAltAltCtrlR Key = KeyAltCtrlR + (1<<16)
KeyAltAltCtrlS Key = KeyAltCtrlS + (1<<16)
KeyAltAltCtrlT Key = KeyAltCtrlT + (1<<16)
KeyAltAltCtrlU Key = KeyAltCtrlU + (1<<16)
KeyAltAltCtrlV Key = KeyAltCtrlV + (1<<16)
KeyAltAltCtrlW Key = KeyAltCtrlW + (1<<16)
KeyAltAltCtrlX Key = KeyAltCtrlX + (1<<16)
KeyAltAltCtrlY Key = KeyAltCtrlY + (1<<16)
KeyAltAltCtrlZ Key = KeyAltCtrlZ + (1<<16)
KeyAltAltEsc Key = KeyAltEsc + (1<<16)
KeyAltAltCtrlLsqBracket Key = KeyAltCtrlLsqBracket + (1<<16)
KeyAltAltCtrl3 Key = KeyAltCtrl3 + (1<<16)
KeyAltAltCtrl4 Key = KeyAltCtrl4 + (1<<16)
KeyAltAltCtrlBackslash Key = KeyAltCtrlBackslash + (1<<16)
KeyAltAltCtrl5 Key = KeyAltCtrl5 + (1<<16)
KeyAltAltCtrlRsqBracket Key = KeyAltCtrlRsqBracket + (1<<16)
KeyAltAltCtrl6 Key = KeyAltCtrl6 + (1<<16)
KeyAltAltCtrl7 Key = KeyAltCtrl7 + (1<<16)
KeyAltAltCtrlSlash Key = KeyAltCtrlSlash + (1<<16)
KeyAltAltCtrlUnderscore Key = KeyAltCtrlUnderscore + (1<<16)
KeyAltAltSpace Key = KeyAltSpace + (1<<16)
KeyAltAltBackspace2 Key = KeyAltBackspace2 + (1<<16)
KeyAltAltCtrl8 Key = KeyAltCtrl8 + (1<<16)
KeyAltAltAltF1 Key = KeyAltAltF1 + (1<<16)
KeyAltAltAltF2 Key = KeyAltAltF2 + (1<<16)
KeyAltAltAltF3 Key = KeyAltAltF3 + (1<<16)
KeyAltAltAltF4 Key = KeyAltAltF4 + (1<<16)
KeyAltAltAltF5 Key = KeyAltAltF5 + (1<<16)
KeyAltAltAltF6 Key = KeyAltAltF6 + (1<<16)
KeyAltAltAltF7 Key = KeyAltAltF7 + (1<<16)
KeyAltAltAltF8 Key = KeyAltAltF8 + (1<<16)
KeyAltAltAltF9 Key = KeyAltAltF9 + (1<<16)
KeyAltAltAltF10 Key = KeyAltAltF10 + (1<<16)
KeyAltAltAltF11 Key = KeyAltAltF11 + (1<<16)
KeyAltAltAltF12 Key = KeyAltAltF12 + (1<<16)
KeyAltAltAltInsert Key = KeyAltAltInsert + (1<<16)
KeyAltAltAltDelete Key = KeyAltAltDelete + (1<<16)
KeyAltAltAltHome Key = KeyAltAltHome + (1<<16)
KeyAltAltAltEnd Key = KeyAltAltEnd + (1<<16)
KeyAltAltAltPgup Key = KeyAltAltPgup + (1<<16)
KeyAltAltAltPgdn Key = KeyAltAltPgdn + (1<<16)
KeyAltAltAltArrowUp Key = KeyAltAltArrowUp + (1<<16)
KeyAltAltAltArrowDown Key = KeyAltAltArrowDown + (1<<16)
KeyAltAltAltArrowLeft Key = KeyAltAltArrowLeft + (1<<16)
KeyAltAltAltArrowRight Key = KeyAltAltArrowRight + (1<<16)
KeyAltAltAltCtrlTilde Key = KeyAltAltCtrlTilde + (1<<16)
KeyAltAltAltCtrl2 Key = KeyAltAltCtrl2 + (1<<16)
KeyAltAltAltCtrlSpace Key = KeyAltAltCtrlSpace + (1<<16)
KeyAltAltAltCtrlA Key = KeyAltAltCtrlA + (1<<16)
KeyAltAltAltCtrlB Key = KeyAltAltCtrlB + (1<<16)
KeyAltAltAltCtrlC Key = KeyAltAltCtrlC + (1<<16)
KeyAltAltAltCtrlD Key = KeyAltAltCtrlD + (1<<16)
KeyAltAltAltCtrlE Key = KeyAltAltCtrlE + (1<<16)
KeyAltAltAltCtrlF Key = KeyAltAltCtrlF + (1<<16)
KeyAltAltAltCtrlG Key = KeyAltAltCtrlG + (1<<16)
KeyAltAltAltBackspace Key = KeyAltAltBackspace + (1<<16)
KeyAltAltAltCtrlH Key = KeyAltAltCtrlH + (1<<16)
KeyAltAltAltTab Key = KeyAltAltTab + (1<<16)
KeyAltAltAltCtrlI Key = KeyAltAltCtrlI + (1<<16)
KeyAltAltAltCtrlJ Key = KeyAltAltCtrlJ + (1<<16)
KeyAltAltAltCtrlK Key = KeyAltAltCtrlK + (1<<16)
KeyAltAltAltCtrlL Key = KeyAltAltCtrlL + (1<<16)
KeyAltAltAltEnter Key = KeyAltAltEnter + (1<<16)
KeyAltAltAltCtrlM Key = KeyAltAltCtrlM + (1<<16)
KeyAltAltAltCtrlN Key = KeyAltAltCtrlN + (1<<16)
KeyAltAltAltCtrlO Key = KeyAltAltCtrlO + (1<<16)
KeyAltAltAltCtrlP Key = KeyAltAltCtrlP + (1<<16)
KeyAltAltAltCtrlQ Key = KeyAltAltCtrlQ + (1<<16)
KeyAltAltAltCtrlR Key = KeyAltAltCtrlR + (1<<16)
KeyAltAltAltCtrlS Key = KeyAltAltCtrlS + (1<<16)
KeyAltAltAltCtrlT Key = KeyAltAltCtrlT + (1<<16)
KeyAltAltAltCtrlU Key = KeyAltAltCtrlU + (1<<16)
KeyAltAltAltCtrlV Key = KeyAltAltCtrlV + (1<<16)
KeyAltAltAltCtrlW Key = KeyAltAltCtrlW + (1<<16)
KeyAltAltAltCtrlX Key = KeyAltAltCtrlX + (1<<16)
KeyAltAltAltCtrlY Key = KeyAltAltCtrlY + (1<<16)
KeyAltAltAltCtrlZ Key = KeyAltAltCtrlZ + (1<<16)
KeyAltAltAltEsc Key = KeyAltAltEsc + (1<<16)
KeyAltAltAltCtrlLsqBracket Key = KeyAltAltCtrlLsqBracket + (1<<16)
KeyAltAltAltCtrl3 Key = KeyAltAltCtrl3 + (1<<16)
KeyAltAltAltCtrl4 Key = KeyAltAltCtrl4 + (1<<16)
KeyAltAltAltCtrlBackslash Key = KeyAltAltCtrlBackslash + (1<<16)
KeyAltAltAltCtrl5 Key = KeyAltAltCtrl5 + (1<<16)
KeyAltAltAltCtrlRsqBracket Key = KeyAltAltCtrlRsqBracket + (1<<16)
KeyAltAltAltCtrl6 Key = KeyAltAltCtrl6 + (1<<16)
KeyAltAltAltCtrl7 Key = KeyAltAltCtrl7 + (1<<16)
KeyAltAltAltCtrlSlash Key = KeyAltAltCtrlSlash + (1<<16)
KeyAltAltAltCtrlUnderscore Key = KeyAltAltCtrlUnderscore + (1<<16)
KeyAltAltAltSpace Key = KeyAltAltSpace + (1<<16)
KeyAltAltAltBackspace2 Key = KeyAltAltBackspace2 + (1<<16)
KeyAltAltAltCtrl8 Key = KeyAltAltCtrl8 + (1<<16)
KeyAltAltAltA Key = KeyAltAltA + (1<<16)
KeyAltAltAltB Key = KeyAltAltB + (1<<16)
KeyAltAltAltC Key = KeyAltAltC + (1<<16)
KeyAltAltAltD Key = KeyAltAltD + (1<<16)
KeyAltAltAltE Key = KeyAltAltE + (1<<16)
KeyAltAltAltF Key = KeyAltAltF + (1<<16)
KeyAltAltAltG Key = KeyAltAltG + (1<<16)
KeyAltAltAltH Key = KeyAltAltH + (1<<16)
KeyAltAltAltI Key = KeyAltAltI + (1<<16)
KeyAltAltAltJ Key = KeyAltAltJ + (1<<16)
KeyAltAltAltK Key = KeyAltAltK + (1<<16)
KeyAltAltAltL Key = KeyAltAltL + (1<<16)
KeyAltAltAltM Key = KeyAltAltM + (1<<16)
KeyAltAltAltN Key = KeyAltAltN + (1<<16)
KeyAltAltAltO Key = KeyAltAltO + (1<<16)
KeyAltAltAltP Key = KeyAltAltP + (1<<16)
KeyAltAltAltQ Key = KeyAltAltQ + (1<<16)
KeyAltAltAltR Key = KeyAltAltR + (1<<16)
KeyAltAltAltS Key = KeyAltAltS + (1<<16)
KeyAltAltAltT Key = KeyAltAltT + (1<<16)
KeyAltAltAltU Key = KeyAltAltU + (1<<16)
KeyAltAltAltV Key = KeyAltAltV + (1<<16)
KeyAltAltAltW Key = KeyAltAltW + (1<<16)
KeyAltAltAltX Key = KeyAltAltX + (1<<16)
KeyAltAltAltY Key = KeyAltAltY + (1<<16)
KeyAltAltAltZ Key = KeyAltAltZ + (1<<16)
KeyAltAltA Key = KeyAltA + (1<<16)
KeyAltAltB Key = KeyAltB + (1<<16)
KeyAltAltC Key = KeyAltC + (1<<16)
KeyAltAltD Key = KeyAltD + (1<<16)
KeyAltAltE Key = KeyAltE + (1<<16)
KeyAltAltF Key = KeyAltF + (1<<16)
KeyAltAltG Key = KeyAltG + (1<<16)
KeyAltAltH Key = KeyAltH + (1<<16)
KeyAltAltI Key = KeyAltI + (1<<16)
KeyAltAltJ Key = KeyAltJ + (1<<16)
KeyAltAltK Key = KeyAltK + (1<<16)
KeyAltAltL Key = KeyAltL + (1<<16)
KeyAltAltM Key = KeyAltM + (1<<16)
KeyAltAltN Key = KeyAltN + (1<<16)
KeyAltAltO Key = KeyAltO + (1<<16)
KeyAltAltP Key = KeyAltP + (1<<16)
KeyAltAltQ Key = KeyAltQ + (1<<16)
KeyAltAltR Key = KeyAltR + (1<<16)
KeyAltAltS Key = KeyAltS + (1<<16)
KeyAltAltT Key = KeyAltT + (1<<16)
KeyAltAltU Key = KeyAltU + (1<<16)
KeyAltAltV Key = KeyAltV + (1<<16)
KeyAltAltW Key = KeyAltW + (1<<16)
KeyAltAltX Key = KeyAltX + (1<<16)
KeyAltAltY Key = KeyAltY + (1<<16)
KeyAltAltZ Key = KeyAltZ + (1<<16)
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
