package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os/exec"
	"strings"
	"time"

	"github.com/bmmcginty/barnard/uiterm"
	"github.com/bmmcginty/barnard/gumble/gumble"
	"github.com/kennygrant/sanitize"
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

func Beep() {
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

func(b *Barnard) Notify(event string, who string, what string) {
b.notifyChannel <- []string{event,who,what}
}

func(b *Barnard) Beep() {
Beep()
}

func (b *Barnard) SetSelectedUser(user *gumble.User) {
	b.selectedUser = user
	if user == nil {
		if len(b.UiInput.Text) > 0 {
		}
		b.UpdateInputStatus(fmt.Sprintf("[%s]", b.Client.Self.Channel.Name))
	} else {
		b.UpdateInputStatus(fmt.Sprintf("[@%s]", user.Name))
	}
}

func (b *Barnard) GetInputStatus() string {
	return b.UiInputStatus.Text
}

func (b *Barnard) UpdateInputStatus(status string) {
	if len(status) > 20 {
		status = status[:17] + "..." + "]"
	}
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

func (b *Barnard) UpdateGeneralStatus(text string, notice bool) {
//		if (b.UiStatus.Text == text) {
//return
//}
if(notice) {
		b.UiStatus.Fg = uiterm.ColorWhite | uiterm.AttrBold
		b.UiStatus.Bg = uiterm.ColorRed
} else {
		b.UiStatus.Fg = uiterm.ColorBlack
		b.UiStatus.Bg = uiterm.ColorWhite
}
		b.UiStatus.Text = text
	b.Ui.Refresh()
}

func (b *Barnard) OnVoiceToggle(ui *uiterm.Ui, key uiterm.Key) {
b.setTransmit(ui,2)
}

func (b *Barnard) CommandLog(ui *uiterm.Ui, cmd string) {
b.AddOutputLine("command "+cmd)
}

func (b *Barnard) CommandTalk(ui *uiterm.Ui, cmd string) {
b.setTransmit(ui,2)
}

func (b *Barnard) CommandMicUp(ui *uiterm.Ui, cmd string) {
b.setTransmit(ui,1)
}

func (b *Barnard) CommandMicDown(ui *uiterm.Ui, cmd string) {
b.setTransmit(ui,0)
}

func (b *Barnard) setTransmit(ui *uiterm.Ui, val int) {
if b.Tx && val==1 {
return
}
if b.Tx==false && val==0 {
return
}
	if (b.Tx) {
b.Notify("micdown","me","")
b.Tx=false
b.UpdateGeneralStatus(" Idle ",false)
		b.Stream.StopSource()
} else if b.Connected==false {
b.Notify("error","me","no tx while disconnected")
b.Tx=false
b.UpdateGeneralStatus("no tx while disconnected",true)
	} else {
 b.Tx=true
		err := b.Stream.StartSource(b.UserConfig.GetInputDevice())
		if err != nil {
b.Notify("error","me",err.Error())
			b.UpdateGeneralStatus(err.Error(),true)
} else {
b.Notify("micup","me","")
b.UpdateGeneralStatus(" Tx  ",true)
} //if error transmit
		} //not transmitting
	} //func

func (b *Barnard) OnMicVolumeDown(ui *uiterm.Ui, key uiterm.Key) {
b.Stream.SetMicVolume(-0.1,true)
b.UserConfig.SetMicVolume(b.Stream.GetMicVolume())
}

func (b *Barnard) OnMicVolumeUp(ui *uiterm.Ui, key uiterm.Key) {
b.Stream.SetMicVolume(0.1,true)
b.UserConfig.SetMicVolume(b.Stream.GetMicVolume())
}

func (b *Barnard) OnQuitPress(ui *uiterm.Ui, key uiterm.Key) {
	b.Client.Disconnect()
	b.Ui.Close()
}

func (b *Barnard) CommandExit(ui *uiterm.Ui, cmd string) {
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
		if b.selectedUser != nil {
			b.selectedUser.Send(text)
			b.AddOutputPrivateMessage(b.Client.Self, b.selectedUser, text)
		} else {
			b.Client.Self.Channel.Send(text, false)
			b.AddOutputMessage(b.Client.Self, text)
		}
	}
}

func (b *Barnard) GotoChat() {
	b.OnFocusPress(b.Ui, uiterm.KeyTab)
}

func (b *Barnard) OnUiDoneInitialize(ui *uiterm.Ui) {
b.start()
}

func (b *Barnard) OnUiInitialize(ui *uiterm.Ui) {
	ui.Add(uiViewLogo, &uiterm.Label{
		Text: "Barnard ",
		Fg:   uiterm.ColorWhite | uiterm.AttrBold,
		Bg:   uiterm.ColorMagenta,
	})

	//	ui.Add(uiViewTop, &uiterm.Label{
	//		Fg: uiterm.ColorWhite,
	//		Bg: uiterm.ColorBlue,
	//	})

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
		Generator:         b.TreeItemBuild,
		KeyListener:       b.TreeItemKeyPress,
		CharacterListener: b.TreeItemCharacter,
		Fg:                uiterm.ColorWhite,
		Bg:                uiterm.ColorBlack,
	}
	ui.Add(uiViewTree, &b.UiTree)

//add this to see what your commands are coming in as raw strings
//	b.Ui.AddCommandListener(b.CommandLog, "*")
	b.Ui.AddCommandListener(b.CommandMicUp, "micup")
	b.Ui.AddCommandListener(b.CommandMicDown, "micdown")
	b.Ui.AddCommandListener(b.CommandTalk, "toggle")
	b.Ui.AddCommandListener(b.CommandTalk, "talk")
	b.Ui.AddCommandListener(b.CommandExit, "exit")
	b.Ui.AddKeyListener(b.OnFocusPress, b.Hotkeys.SwitchViews)
	b.Ui.AddKeyListener(b.OnVoiceToggle, b.Hotkeys.Talk)
	b.Ui.AddKeyListener(b.OnTimestampToggle, b.Hotkeys.ToggleTimestamps)
	b.Ui.AddKeyListener(b.OnQuitPress, b.Hotkeys.Exit)
	b.Ui.AddKeyListener(b.OnScrollOutputUp, b.Hotkeys.ScrollUp)
	b.Ui.AddKeyListener(b.OnScrollOutputDown, b.Hotkeys.ScrollDown)
	b.Ui.AddKeyListener(b.OnScrollOutputTop, b.Hotkeys.ScrollToTop)
	b.Ui.AddKeyListener(b.OnScrollOutputBottom, b.Hotkeys.ScrollToBottom)
	b.Ui.SetActive(uiViewInput)
	b.UiTree.Rebuild()
	b.Ui.Refresh()
}

func (b *Barnard) OnUiResize(ui *uiterm.Ui, width, height int) {
	treeHeight := 0
	outputHeight := 0
	active := b.Ui.Active()
	if active == uiViewTree {
		treeHeight = height - 4
		outputHeight = 0
	} else {
		treeHeight = 0
		outputHeight = height - 4
	}
	//		ui.SetBounds(uiViewLogo, 0, 0, 9, 1)
	ui.SetBounds(uiViewOutput, 0, 1, width, outputHeight+1)
	ui.SetBounds(uiViewTree, 0, 1, width, treeHeight+1)
	ui.SetBounds(uiViewStatus, 0, height-2, width, height-1)
	ui.SetBounds(uiViewInputStatus, 0, height-1, len(b.GetInputStatus()), height)
	//setting this to inputStatus+1 will leave one space between inputStatus and input box
	//x starts at 0, so 10 chars of text will go from 0 to 9, there'll be a space at char 10, and we'll start at (10+1)=11
	ui.SetBounds(uiViewInput, len(b.GetInputStatus())+1, height-1, width, height)
}
