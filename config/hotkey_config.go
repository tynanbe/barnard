package config

import (
	"github.com/BenOnions/barnard/uiterm"
)

type Hotkeys struct {
	Talk             *uiterm.Key
	VolumeDown       *uiterm.Key
	VolumeUp         *uiterm.Key
	Exit             *uiterm.Key
	ToggleTimestamps *uiterm.Key
	SwitchViews      *uiterm.Key
	ClearOutput      *uiterm.Key
	ScrollUp         *uiterm.Key
	ScrollDown       *uiterm.Key
	ScrollToTop      *uiterm.Key
	ScrollToBottom   *uiterm.Key
}
