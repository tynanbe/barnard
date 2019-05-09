# barnard

This copy of Barnard and it's associated Gumble library have been modified to support usage by blind users.
Our thanks go out to Tim Cooper for the massive amount of work put into this client, originally found at [github.com/layeh/barnard](https://github.com/layeh/barnard).

## Keystrokes

Pressing tab inside the main window switches between the user/channel tree view and the message input box.
When in the message input box:
*left and right arrow keys move by character
*home/end moves to the beginning/end of the text respectively
*enter submits the entered message
When in the treeview:
*Pressing f5 or f6 on a user changes the volume for that user.
*pressing enter on a user selects that user for PM mode.
*Pressing enter again on that user de-selects that user.
*Pressing enter on a channel de-selects any selected users (if any) and moves you to the specified channel.

## PM Mode

While a user is selected, both private messages and channel messages are displayed.
However, any messages sent will be delivered only to the selected user.
Private messages between you and another party are shown as `pm/source/dest`, where source and dest are the sender and receiver of the message respectively.
When you are finished sending private messages to a user, press tab to reactivate the tree view, and de-select the user as above.

## Info

barnard is a terminal-based client for the [Mumble](https://mumble.info) voice
chat software.

![Screenshot](https://i.imgur.com/B8ldT5k.png)

## Installation

Requirements:

1. [Go](https://golang.org/)
2. [Git](https://git-scm.com/)
3. [Opus](https://opus-codec.org/) development headers
4. [OpenAL](http://kcat.strangesoft.net/openal.html) development headers

To fetch and build:

    go get -u layeh.com/barnard

After running the command above, `barnard` will be compiled as `$(go env GOPATH)/bin/barnard`.

## Manual

### Key bindings

- <kbd>F1</kbd>: toggle voice transmission
- <kbd>Ctrl+L</kbd>: clear chat log
- <kbd>Tab</kbd>: toggle focus between chat and user tree
- <kbd>Page Up</kbd>: scroll chat up
- <kbd>Page Down</kbd>: scroll chat down
- <kbd>Home</kbd>: scroll chat to the top
- <kbd>End</kbd>: scroll chat to the bottom
- <kbd>F10</kbd>: quit

## License

GPLv2

## Author

Tim Cooper (<tim.cooper@layeh.com>)
