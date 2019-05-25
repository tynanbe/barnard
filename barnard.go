package main

import (
	"crypto/tls"

	"github.com/bmmcginty/barnard/uiterm"
	"github.com/bmmcginty/barnard/config"
	"github.com/bmmcginty/gumble/gumble"
	"github.com/bmmcginty/gumble/gumbleopenal"
)

type Barnard struct {
	Config     *gumble.Config
	UserConfig *config.Config
Hotkeys *config.Hotkeys
	Client     *gumble.Client

	Address   string
	TLSConfig tls.Config

	Stream *gumbleopenal.Stream

	Ui              *uiterm.Ui
	UiOutput        uiterm.Textview
	UiInput         uiterm.Textbox
	UiStatus        uiterm.Label
	UiTree          uiterm.Tree
	UiInputStatus   uiterm.Label
	SelectedChannel *gumble.Channel
	selectedUser    *gumble.User
}
