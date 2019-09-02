package main

import (
//	"fmt"
	"github.com/bmmcginty/barnard/uiterm"
	"github.com/bmmcginty/gumble/gumble"
	"sort"
)

type TreeItem struct {
	User    *gumble.User
	Channel *gumble.Channel
}

func (ti TreeItem) String() string {
	if ti.User != nil {
		return ti.User.Name
	}
	if ti.Channel != nil {
		return "#"+ti.Channel.Name
	}
	return ""
}

func (ti TreeItem) TreeItemStyle(fg, bg uiterm.Attribute, active bool) (uiterm.Attribute, uiterm.Attribute) {
	if ti.Channel != nil {
		fg |= uiterm.AttrBold
	}
	if active {
		fg, bg = bg, fg
	}
	return fg, bg
}

func (b *Barnard) TreeItemCharacter(ui *uiterm.Ui, tree *uiterm.Tree, item uiterm.TreeItem, ch rune) {
}

func (b *Barnard) TreeItemKeyPress(ui *uiterm.Ui, tree *uiterm.Tree, item uiterm.TreeItem, key uiterm.Key) {
	treeItem := item.(TreeItem)
	if key == uiterm.KeyEnter {
		if treeItem.Channel != nil {
			b.Client.Self.Move(treeItem.Channel)
			b.SetSelectedUser(nil)
			b.GotoChat()
		}
		if treeItem.User != nil {
			if b.selectedUser == treeItem.User {
				b.SetSelectedUser(nil)
				b.GotoChat()
			} else {
				b.SetSelectedUser(treeItem.User)
				b.GotoChat()
			} //select
		} //if user and not selected
	} //if enter key
	if treeItem.Channel != nil {
		var c = treeItem.Channel
		var change = float32(0.0)
		var changeType = ""
		if key == *b.Hotkeys.VolumeDown {
			changeType = "volume"
			change = -0.1
		}
		if key == *b.Hotkeys.VolumeUp {
			changeType = "volume"
			change = 0.1
		}
		if changeType == "volume" {
			for _, u := range c.Users {
				var au = u.AudioSource
if au==nil {
continue
}
				var gain = au.GetGain()
				gain += change
				if gain < au.GetMinGain() {
					gain = au.GetMinGain()
				}
				if gain > au.GetMaxGain() {
					gain = au.GetMaxGain()
				}
				au.SetGain(gain)
b.UserConfig.UpdateConfig(u)
b.UserConfig.SaveConfig()
			} //each user
		} //set volume
	} //enter on channel
	if treeItem.User != nil {
		var u = treeItem.User
		var au = u.AudioSource
if au!=nil {
		var set_gain = false
var gain float32
		if key==*b.Hotkeys.BoostDown {
u.Boost-=1
if u.Boost<1 {
u.Boost=1
}
//		au.SetPitch(au.GetPitch()-0.1)
		}
		if key==*b.Hotkeys.BoostUp {
u.Boost+=1
//		au.SetPitch(au.GetPitch()+0.1)
		}
		if key == *b.Hotkeys.VolumeDown {
			set_gain = true
			var mingain = au.GetMinGain()
			gain = au.GetGain()
			gain -= 0.1
			if gain < mingain {
				gain = mingain
			}
		} //f5
		if key == *b.Hotkeys.VolumeUp {
set_gain=true
			var maxgain = au.GetMaxGain()
			gain = au.GetGain()
			gain += 0.1
			if gain > maxgain {
				gain = maxgain
			}
		} //f5
		if set_gain {
			au.SetGain(gain)
b.UserConfig.UpdateConfig(u)
b.UserConfig.SaveConfig()
			//b.Log(fmt.Sprintf("%s gain %.2f",u.Name,au.GetGain()))
		} //if set gain
} //user has audioSource
	} //user highlighted
} //func

func (b *Barnard) TreeItemBuild(item uiterm.TreeItem) []uiterm.TreeItem {
	if b.Client == nil {
		return nil
	}

	var treeItem TreeItem
	if ti, ok := item.(TreeItem); !ok {
		root := b.Client.Channels[0]
		if root == nil {
			return nil
		}
		return []uiterm.TreeItem{
			TreeItem{
				Channel: root,
			},
		}
	} else {
		treeItem = ti
	}

	if treeItem.User != nil {
		return nil
	}

	users := []uiterm.TreeItem{}
	ul := []*gumble.User{}
	for _, user := range treeItem.Channel.Users {
		ul = append(ul, user)
		var u = ul[len(ul)-1]
		_ = u
	}
	sort.Slice(ul, func(i, j int) bool {
		return ul[i].Name < ul[j].Name
	})
	for _, user := range ul {
		users = append(users, TreeItem{
			User: user,
		})
	}

	channels := []uiterm.TreeItem{}
	cl := []*gumble.Channel{}
	for _, subchannel := range treeItem.Channel.Children {
		cl = append(cl, subchannel)
	}
	sort.Slice(cl, func(i, j int) bool {
		return cl[i].Name < cl[j].Name
	})
	for _, subchannel := range cl {
		channels = append(channels, TreeItem{
			Channel: subchannel,
		})
	}

	return append(users, channels...)
}
