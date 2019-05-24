package main

import (
	"encoding/json"
	"github.com/bmmcginty/gumble/gumble"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
)

type userConfig struct {
	config *jsonConfig
}

type jsonConfig struct {
	servers []*jsonServer
}

type jsonServer struct {
	host  string
	port  int
	users []*jsonUser
}

type jsonUser struct {
	username string
	volume   float32
}

func (c *userConfig) LoadConfig() {
	var fn = "~/.barnard.json"
	if !fileExists(fn) {
		c.config = &jsonConfig{}
	} else {
		var jc jsonConfig
		var data []byte
		data = readFile(fn)
		if data == nil {
			c.config = &jsonConfig{}
			return
		}
		err := json.Unmarshal(data, jc)
		if err != nil {
			panic(err)
		}
		c.config = &jc
	}
	return
}

func (c *userConfig) findServer(address string) *jsonServer {
	if c.config.servers == nil {
		c.config.servers = make([]*jsonServer, 0)
	}
	host, port := makeHostPort(address)
	var t *jsonServer
	for _, s := range c.config.servers {
		if s.port == port && s.host == host {
			t = s
			break
		}
	}
	if t == nil {
		t = &jsonServer{
			host: host,
			port: port,
		}
		c.config.servers = append(c.config.servers, t)
	}
	return t
}

func (c *userConfig) findUser(address string, user string) *jsonUser {
	var s *jsonServer
	s = c.findServer(address)
	if s.users == nil {
		s.users = make([]*jsonUser, 0)
	}
	var t *jsonUser
	for _, u := range s.users {
		if u.username == user {
			t = u
			break
		}
	}
	if t == nil {
		t = &jsonUser{
			username: user,
		}
		s.users = append(s.users, t)
	}
	return t
}

func (c *userConfig) updateUser(u *gumble.User) {
	var j *jsonUser
	j = c.findUser(u.GetClient().Config.Address, u.Name)
	j.volume = u.AudioSource.GetGain()
}

func (b *Barnard) makeEmptyConfig() {
	var c *userConfig
	c = &userConfig{}
	c.config = &jsonConfig{}
	c.config.servers = make([]*jsonServer, 0)
	host, port := makeHostPort(b.Client.Config.Address)
	var s = &jsonServer{
		host: host,
		port: port,
	}
	c.config.servers = append(c.config.servers, s)
}

func readFile(path string) []byte {
	if !fileExists(path) {
		return nil
	}
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return dat
}

func fileExists(path string) bool {
	path = resolvePath(path)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func resolvePath(path string) string {
	if strings.HasPrefix(path, "~/") || strings.Contains(path, "$HOME") {
		usr, err := user.Current()
		if err != nil {
			panic(err)
		}
		var hd = usr.HomeDir
		if strings.Contains(path, "$HOME") {
			path = strings.Replace(path, "$HOME", hd, 1)
		} else {
			path = hd + path[2:]
		}
	}
	return path
}

func makeHostPort(addr string) (string, int) {
	parts := strings.Split(addr, ":")
	host := parts[0]
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return host, port
}
