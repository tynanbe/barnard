package main

import (
	"fmt"
	"net"
	"os"

	"github.com/bmmcginty/gumble/gumble"
	"github.com/bmmcginty/gumble/gumbleopenal"
	"github.com/bmmcginty/gumble/gumbleutil"
)

func (b *Barnard) start() {
	b.Config.Attach(gumbleutil.AutoBitrate)
	b.Config.Attach(b)
	b.Config.Address = b.Address

	var err error
	_, err = gumble.DialWithDialer(new(net.Dialer), b.Config, &b.TLSConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Audio
	if os.Getenv("ALSOFT_LOGLEVEL") == "" {
		os.Setenv("ALSOFT_LOGLEVEL", "0")
	}
	if stream, err := gumbleopenal.New(b.Client); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	} else {
		b.Stream = stream
	}
}

func (b *Barnard) OnConnect(e *gumble.ConnectEvent) {
	b.Client = e.Client

	b.Ui.SetActive(uiViewInput)
	b.UiTree.Rebuild()
	b.Ui.Refresh()

	b.UpdateInputStatus(fmt.Sprintf("[%s]", e.Client.Self.Channel.Name))
	b.AddOutputLine(fmt.Sprintf("Connected to %s", b.Client.Conn.RemoteAddr()))
	if e.WelcomeMessage != nil {
		b.AddOutputLine(fmt.Sprintf("Welcome message: %s", esc(*e.WelcomeMessage)))
	}
}

func (b *Barnard) OnDisconnect(e *gumble.DisconnectEvent) {
	var reason string
	switch e.Type {
	case gumble.DisconnectError:
		reason = "connection error"
	}
	if reason == "" {
		b.AddOutputLine("Disconnected")
	} else {
		b.AddOutputLine("Disconnected: " + reason)
	}
	b.UiTree.Rebuild()
	b.Ui.Refresh()
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
		b.AddOutputMessage(e.Sender, e.Message)
	} else {
		b.AddOutputPrivateMessage(e.Sender, b.Client.Self, e.Message)
	}
}

func (b *Barnard) OnUserChange(e *gumble.UserChangeEvent) {
	var s = "unknown"
	if e.Type.Has(gumble.UserChangeConnected) {
		s = "joined"
	}
	if e.Type.Has(gumble.UserChangeDisconnected) {
		s = "left"
		if e.User == b.selectedUser {
			b.SetSelectedUser(nil)
		}
	}
	if e.User.Channel.Name == b.Client.Self.Channel.Name {
		b.AddOutputLine(fmt.Sprintf("%s %s %s", e.User.Name, s, e.User.Channel.Name))
	}
	if e.Type.Has(gumble.UserChangeChannel) && e.User == b.Client.Self {
		b.UpdateInputStatus(fmt.Sprintf("[%s]", e.User.Channel.Name))
	}
	b.UiTree.Rebuild()
	b.Ui.Refresh()
}

func (b *Barnard) OnChannelChange(e *gumble.ChannelChangeEvent) {
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
}

func (b *Barnard) OnACL(e *gumble.ACLEvent) {
}

func (b *Barnard) OnBanList(e *gumble.BanListEvent) {
}

func (b *Barnard) OnContextActionChange(e *gumble.ContextActionChangeEvent) {
}

func (b *Barnard) OnServerConfig(e *gumble.ServerConfigEvent) {
}
