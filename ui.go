package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os/exec"
	"strings"
	"time"

	"github.com/bmmcginty/barnard/uiterm"
	"github.com/kennygrant/sanitize"
	"github.com/bmmcginty/gumble/gumble"
)

const (
	uiViewLogo        = "logo"
	uiViewTop         = "top"
	uiViewStatus      = "status"
	uiViewInput       = "input"
	uiViewInputStatus = "inputstatus"
	uiViewOutput      = "output"
	uiViewTree        = "tree"
)

func beep() {
	cmd := exec.Command("beep")
	cmdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	if cmdout != nil {
	}
}

func esc(str string) string {
	return sanitize.HTML(str)
}

func (b *Barnard) SetSelectedUser(user *gumble.User) {
b.selectedUser=user
if user==nil {
if len(b.UiInput.Text)>0 {
}
	b.UpdateInputStatus(fmt.Sprintf("[%s]", b.Client.Self.Channel.Name))
} else {
	b.UpdateInputStatus(fmt.Sprintf("[@%s]", user.Name))
}
}

func (b *Barnard) UpdateInputStatus(status string) {
	b.UiInputStatus.Text = status
	b.UiTree.Rebuild()
	b.Ui.Refresh()
}

func (b *Barnard) AddOutputLine(line string) {
	now := time.Now()
	b.UiOutput.AddLine(fmt.Sprintf("[%02d:%02d:%02d] %s", now.Hour(), now.Minute(), now.Second(), line))
}

func (b *Barnard) AddOutputMessage(sender *gumble.User, message string) {
	if sender == nil {
		b.AddOutputLine(message)
	} else {
		b.AddOutputLine(fmt.Sprintf("%s: %s", sender.Name, strings.TrimSpace(esc(message))))
	}
}

func (b *Barnard) AddOutputPrivateMessage(source *gumble.User, dest *gumble.User, message string) {
		b.AddOutputLine(fmt.Sprintf("pm/%s/%s: %s", source.Name, dest.Name, strings.TrimSpace(esc(message))))
	}

func (b *Barnard) OnTimestampToggle(ui *uiterm.Ui, key uiterm.Key) {
	b.UiOutput.ToggleTimestamps()
}

func (b *Barnard) OnVoiceToggle(ui *uiterm.Ui, key uiterm.Key) {
	if b.UiStatus.Text != " Idle " {
		b.UiStatus.Text = " Idle "
		b.UiStatus.Fg = uiterm.ColorBlack
		b.UiStatus.Bg = uiterm.ColorWhite
		b.Stream.StopSource()
	} else {
		b.UiStatus.Fg = uiterm.ColorWhite | uiterm.AttrBold
		b.UiStatus.Bg = uiterm.ColorRed
		b.UiStatus.Text = "  Tx  "
		err := b.Stream.StartSource()
if err!=nil {
b.UiStatus.Text=err.Error()
}
	}
	ui.Refresh()
}

func (b *Barnard) OnQuitPress(ui *uiterm.Ui, key uiterm.Key) {
	b.Client.Disconnect()
	b.Ui.Close()
}

func (b *Barnard) OnClearPress(ui *uiterm.Ui, key uiterm.Key) {
	b.UiOutput.Clear()
}

func (b *Barnard) OnScrollOutputUp(ui *uiterm.Ui, key uiterm.Key) {
	b.UiOutput.ScrollUp()
}

func (b *Barnard) OnScrollOutputDown(ui *uiterm.Ui, key uiterm.Key) {
	b.UiOutput.ScrollDown()
}

func (b *Barnard) OnScrollOutputTop(ui *uiterm.Ui, key uiterm.Key) {
	b.UiOutput.ScrollTop()
}

func (b *Barnard) OnScrollOutputBottom(ui *uiterm.Ui, key uiterm.Key) {
	b.UiOutput.ScrollBottom()
}

func (b *Barnard) OnFocusPress(ui *uiterm.Ui, key uiterm.Key) {
	active := b.Ui.Active()
	if active == uiViewInput {
		b.Ui.SetActive(uiViewTree)
	} else if active == uiViewTree {
		b.Ui.SetActive(uiViewInput)
	}
	width, height := termbox.Size()
	b.OnUiResize(ui, width, height)
	ui.Refresh()
}

func (b *Barnard) OnTextInput(ui *uiterm.Ui, textbox *uiterm.Textbox, text string) {
	if text == "" {
		return
	}
	if b.Client != nil && b.Client.Self != nil {
if b.selectedUser!=nil {
b.selectedUser.Send(text)
b.AddOutputPrivateMessage(b.Client.Self,b.selectedUser,text)
} else {
		b.Client.Self.Channel.Send(text, false)
		b.AddOutputMessage(b.Client.Self, text)
}
	}
}

func (b *Barnard) GotoChat() {
b.OnFocusPress(b.Ui,uiterm.KeyTab)
}

func (b *Barnard) OnUiInitialize(ui *uiterm.Ui) {
	ui.Add(uiViewLogo, &uiterm.Label{
		Text: " barnard ",
		Fg:   uiterm.ColorWhite | uiterm.AttrBold,
		Bg:   uiterm.ColorMagenta,
	})

	ui.Add(uiViewTop, &uiterm.Label{
		Fg: uiterm.ColorWhite,
		Bg: uiterm.ColorBlue,
	})

	b.UiStatus = uiterm.Label{
		Text: " Idle ",
		Fg:   uiterm.ColorBlack,
		Bg:   uiterm.ColorWhite,
	}
	ui.Add(uiViewStatus, &b.UiStatus)

	b.UiInput = uiterm.Textbox{
		Fg:    uiterm.ColorWhite,
		Bg:    uiterm.ColorBlack,
		Input: b.OnTextInput,
	}
	ui.Add(uiViewInput, &b.UiInput)

	b.UiInputStatus = uiterm.Label{
		Fg: uiterm.ColorBlack,
		Bg: uiterm.ColorWhite,
	}
	ui.Add(uiViewInputStatus, &b.UiInputStatus)

	b.UiOutput = uiterm.Textview{
		Fg: uiterm.ColorWhite,
		Bg: uiterm.ColorBlack,
	}
	ui.Add(uiViewOutput, &b.UiOutput)

	b.UiTree = uiterm.Tree{
		Generator: b.TreeItemBuild,
		KeyListener:  b.TreeItemKeyPress,
		CharacterListener:  b.TreeItemCharacter,
		Fg:        uiterm.ColorWhite,
		Bg:        uiterm.ColorBlack,
	}
	ui.Add(uiViewTree, &b.UiTree)

	b.Ui.AddKeyListener(b.OnFocusPress, uiterm.KeyTab)
	b.Ui.AddKeyListener(b.OnVoiceToggle, uiterm.KeyF1)
	b.Ui.AddKeyListener(b.OnTimestampToggle, uiterm.KeyF3)
	b.Ui.AddKeyListener(b.OnQuitPress, uiterm.KeyF10)
	b.Ui.AddKeyListener(b.OnClearPress, uiterm.KeyCtrlL)
	b.Ui.AddKeyListener(b.OnScrollOutputUp, uiterm.KeyPgup)
	b.Ui.AddKeyListener(b.OnScrollOutputDown, uiterm.KeyPgdn)
	b.Ui.AddKeyListener(b.OnScrollOutputTop, uiterm.KeyHome)
	b.Ui.AddKeyListener(b.OnScrollOutputBottom, uiterm.KeyEnd)

	b.start()
}

func (b *Barnard) OnUiResize(ui *uiterm.Ui, width, height int) {
	treeHeight := 0
	outputHeight := 0
	active := b.Ui.Active()
	if active == uiViewTree {
		treeHeight = 10
		outputHeight = 0
	} else {
		treeHeight = 0
		outputHeight = height - 4
	}
	ui.SetBounds(uiViewOutput, 0, 1, width, outputHeight)
	//0, 1, width-20, height-2)
	ui.SetBounds(uiViewTree, 0, 1, width, treeHeight)
	//width-20, 1, width, height-2)
	//	ui.SetBounds(uiViewLogo, 0, 0, 9, 1)
	//	ui.SetBounds(uiViewTop, 9, 0, width-6, 1)
	ui.SetBounds(uiViewStatus, 0, height-2, width, height-1)
	//width-6, 0, width, 1)
	ui.SetBounds(uiViewInput, 12, height-1, width, height)
	//0, height-1, width, height)
	ui.SetBounds(uiViewInputStatus, 0, height-1, 20, height)
	//0, height-2, width, height-1)
}
