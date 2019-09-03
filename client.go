package main

import (
	"fmt"
	"net"
"time"

	"github.com/bmmcginty/gumble/gumble"
	"github.com/bmmcginty/gumble/gumbleopenal"
	"github.com/bmmcginty/gumble/gumbleutil"
)

func (b *Barnard) start() {
	b.Config.Attach(gumbleutil.AutoBitrate)
	b.Config.Attach(b)
	b.Config.Address = b.Address
	// test Audio
	_, err := gumbleopenal.New(b.Client,b.UserConfig.GetInputDevice(),b.UserConfig.GetOutputDevice(),true)
if err != nil {
b.exitWithError(err)
return
}
//connect, not reconnect
b.connect(false)
}

func (b *Barnard) exitWithError(err error) {
b.Ui.Close()
b.exitStatus=1
b.exitMessage=err.Error()
}

func (b *Barnard) connect(reconnect bool) bool {
	var err error
	_, err = gumble.DialWithDialer(new(net.Dialer), b.Config, &b.TLSConfig)
	if err != nil {
if(reconnect) {
b.Log(err.Error())
} else {
b.exitWithError(err)
}
return false
	}

	stream, err := gumbleopenal.New(b.Client,b.UserConfig.GetInputDevice(),b.UserConfig.GetOutputDevice(),false)
if err != nil {
b.exitWithError(err)
return false
}
b.Stream=stream
		b.Stream.AttachStream(b.Client)
b.Connected=true
return true
}

func (b *Barnard) OnConnect(e *gumble.ConnectEvent) {
	b.Client = e.Client

	b.Ui.SetActive(uiViewInput)
	b.UiTree.Rebuild()
	b.Ui.Refresh()

for _,u := range b.Client.Users {
b.UserConfig.UpdateUser(u)
}

	b.UpdateInputStatus(fmt.Sprintf("[%s]", e.Client.Self.Channel.Name))
	b.AddOutputLine(fmt.Sprintf("Connected to %s", b.Client.Conn.RemoteAddr()))
wmsg := ""
	if e.WelcomeMessage != nil {
wmsg = esc(*e.WelcomeMessage)
}
b.Notify("connect","me",wmsg)
if wmsg != "" {
		b.AddOutputLine(fmt.Sprintf("Welcome message: %s", wmsg))
	}
	b.Ui.Refresh()
}

func (b *Barnard) OnDisconnect(e *gumble.DisconnectEvent) {
	var reason string
	switch e.Type {
	case gumble.DisconnectError:
		reason = "connection error"
	}
b.Notify("disconnect","me",reason)
	if reason == "" {
		b.AddOutputLine("Disconnected")
	} else {
		b.AddOutputLine("Disconnected: " + reason)
	}
b.Tx=false
b.Connected=false
	b.UiTree.Rebuild()
	b.Ui.Refresh()
go b.reconnectGoroutine()
}

func (b *Barnard) reconnectGoroutine() {
for {
res := b.connect(true)
if res==true {
break
}
time.Sleep(15 * time.Second)
}
}

func (b *Barnard) Log(s string) {
	b.AddOutputMessage(nil, s)
}

func (b *Barnard) OnTextMessage(e *gumble.TextMessageEvent) {
	var public = false
	for _, c := range e.Channels {
		if c.Name == b.Client.Self.Channel.Name {
			public = true
			break
		}
	}
	if public {
b.Notify("msg",e.Sender.Name,e.Message)
		b.AddOutputMessage(e.Sender, e.Message)
	} else {
b.Notify("pm",e.Sender.Name,e.Message)
		b.AddOutputPrivateMessage(e.Sender, b.Client.Self, e.Message)
	}
}

func (b *Barnard) OnUserChange(e *gumble.UserChangeEvent) {
if e.User != nil {
b.UserConfig.UpdateUser(e.User)
}
	var s = "unknown"
 var t = "unknown"
	if e.Type.Has(gumble.UserChangeConnected) {
		s = "joined"
t="join"
	}
	if e.Type.Has(gumble.UserChangeDisconnected) {
		s = "left"
t="leave"
		if e.User == b.selectedUser {
			b.SetSelectedUser(nil)
		}
	}
	if e.User.Channel.Name == b.Client.Self.Channel.Name {
b.Notify(t,e.User.Name,e.User.Channel.Name)
		b.AddOutputLine(fmt.Sprintf("%s %s %s", e.User.Name, s, e.User.Channel.Name))
	}
	if e.Type.Has(gumble.UserChangeChannel) && e.User == b.Client.Self {
		b.UpdateInputStatus(fmt.Sprintf("[%s]", e.User.Channel.Name))
	}
	b.UiTree.Rebuild()
	b.Ui.Refresh()
}

func (b *Barnard) OnChannelChange(e *gumble.ChannelChangeEvent) {
		b.UpdateInputStatus(fmt.Sprintf("[%s]", e.Channel.Name))
	b.UiTree.Rebuild()
	b.Ui.Refresh()
}

func (b *Barnard) OnPermissionDenied(e *gumble.PermissionDeniedEvent) {
	var info string
	switch e.Type {
	case gumble.PermissionDeniedOther:
		info = e.String
	case gumble.PermissionDeniedPermission:
		info = "insufficient permissions"
	case gumble.PermissionDeniedSuperUser:
		info = "cannot modify SuperUser"
	case gumble.PermissionDeniedInvalidChannelName:
		info = "invalid channel name"
	case gumble.PermissionDeniedTextTooLong:
		info = "text too long"
	case gumble.PermissionDeniedTemporaryChannel:
		info = "temporary channel"
	case gumble.PermissionDeniedMissingCertificate:
		info = "missing certificate"
	case gumble.PermissionDeniedInvalidUserName:
		info = "invalid user name"
	case gumble.PermissionDeniedChannelFull:
		info = "channel full"
	case gumble.PermissionDeniedNestingLimit:
		info = "nesting limit"
	}
	b.AddOutputLine(fmt.Sprintf("Permission denied: %s", info))
}

func (b *Barnard) OnUserList(e *gumble.UserListEvent) {
//for _,u := range e.UserList {
//b.UserConfig.UpdateUser(u)
//}
}

func (b *Barnard) OnACL(e *gumble.ACLEvent) {
}

func (b *Barnard) OnBanList(e *gumble.BanListEvent) {
}

func (b *Barnard) OnContextActionChange(e *gumble.ContextActionChangeEvent) {
}

func (b *Barnard) OnServerConfig(e *gumble.ServerConfigEvent) {
}
