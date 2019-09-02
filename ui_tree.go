package main

import (
//"math"
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

func (b *Barnard) changeVolume(users []*gumble.User, change float32) {
for _,u := range users {
au := u.AudioSource
if au==nil {
continue
}
var boost uint16
var cv float32
var ng float32
var curboost float32
curboost = float32((u.Boost-1))/10
cv = au.GetGain() + curboost
ng = cv+change
boost = uint16(1)
//b.AddOutputLine(fmt.Sprintf("cv %.2f change %.2f ng %.2f",cv,change,ng))
if ng > 1.0 {
//1.0 will give volume of one and boost of 1
//1.1 will give volume of 1 and boost of 2
//b.AddOutputLine(fmt.Sprintf("partperc %.2f",(ng*10)))
perc := uint16((ng*10))-10
perc+=1
boost=perc
ng=1.0
}
if ng < 0 {
ng=0.0
}
//b.AddOutputLine(fmt.Sprintf("boost %d ng %.2f",boost,ng))
u.Boost=boost
				au.SetGain(ng)
b.UserConfig.UpdateConfig(u)
}
b.UserConfig.SaveConfig()
}

func makeUsersArray(users gumble.Users) []*gumble.User {
t := make([]*gumble.User,0,len(users))
for _,u := range users {
t=append(t,u)
}
return t
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
		if key == *b.Hotkeys.VolumeDown {
b.changeVolume(makeUsersArray(c.Users),-0.1)
}
		if key == *b.Hotkeys.VolumeUp {
b.changeVolume(makeUsersArray(c.Users),0.1)
}
		} //set volume
	if treeItem.User != nil {
		var u = treeItem.User
		if key == *b.Hotkeys.VolumeDown {
b.changeVolume([]*gumble.User{u},-0.1)
}
		if key == *b.Hotkeys.VolumeUp {
b.changeVolume([]*gumble.User{u},0.1)
}
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
