package main

import (
"fmt"
"sort"
	"github.com/bmmcginty/barnard/uiterm"
	"github.com/bmmcginty/gumble/gumble"
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
		return ti.Channel.Name
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

func (b *Barnard) TreeItemKeyPress(ui *uiterm.Ui, tree *uiterm.Tree, item uiterm.TreeItem, mod uiterm.Modifier, key uiterm.Key) {
	treeItem := item.(TreeItem)
if(key==uiterm.KeyEnter) {
	if treeItem.Channel != nil {
		b.Client.Self.Move(treeItem.Channel)
b.SetSelectedUser(nil)
	}
	if treeItem.User != nil {
if b.selectedUser==treeItem.User {
b.SetSelectedUser(nil)
} else {
b.SetSelectedUser(treeItem.User)
} //select
	} //if user and not selected
} //if enter key
if b.selectedUser!=nil {
var au = b.selectedUser.AudioSource
var set_gain=false
if key==uiterm.KeyF5 {
set_gain=true
var mingain = au.GetMinGain()
var gain=au.GetGain()
gain-=0.1
if gain < mingain {
gain=mingain
}
au.SetGain(gain)
} //f5
if key==uiterm.KeyF6 {
var maxgain = au.GetMaxGain()
var gain=au.GetGain()
gain+=0.1
if gain > maxgain {
gain=maxgain
}
au.SetGain(gain)
} //f5
if set_gain {
b.Log(fmt.Sprintf("%s gain %.2f",b.selectedUser.Name,au.GetGain()))
} //if set gain
} //selectedUser
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
ul=append(ul,user)
var u = ul[len(ul)-1]
_=u
}
sort.Slice(ul, func(i,j int) bool {
return ul[i].Name < ul[j].Name
})
	for _,user := range ul {
		users = append(users, TreeItem{
			User: user,
		})
	}

	channels := []uiterm.TreeItem{}
cl := []*gumble.Channel{}
	for _, subchannel := range treeItem.Channel.Children {
cl=append(cl,subchannel)
}
sort.Slice(cl,func(i,j int) bool {
return cl[i].Name<cl[j].Name
})
for _, subchannel := range cl {
		channels = append(channels, TreeItem{
			Channel: subchannel,
		})
	}

	return append(users, channels...)
}
