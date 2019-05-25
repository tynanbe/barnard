# barnard

This copy of Barnard and it's associated Gumble library have been modified to support usage by blind users.
Our thanks go out to Tim Cooper for the massive amount of work put into this client, originally found at [github.com/layeh/barnard](https://github.com/layeh/barnard).

## Keystrokes

The below keystroeks will be automatically created in the $HOME/.barnard.yaml file, if it does not already exist.
They can be customized in that file when Barnard is not running.
If you edit the file when Barnard _is running, your changes will be overwritten.

Pressing tab inside the main window switches between the user/channel tree view and the message input box.
When in the message input box:
* left and right arrow keys move by character
* home/end moves to the beginning/end of the text respectively
* enter submits the entered message

When in the treeview, pressing:
* f5 or f6 on a channel changes the volume for all users in that channel
* f5 or f6 on a user changes the volume for that user.
* enter on de-selected user selects that user for PM mode.
* enter on selected user de-selects the user
* enter on a channel de-selects any selected users (if any) and moves you to the specified channel.

## Volume

If you set the volume for a user (using the F5/F6 keys by default), Barnard will remember that volume, and will keep that user at that volume.
The volume is set for a single user on a single server.
This means you may have to set a person to a custom volume multiple times, if you are both on multiple servers together.

If you set the volume of a channel, you are basically relatively adjusting each user's volume.
If Jim's volume is set to 0.1, and larry's volume is set to 0.9, lowering the channel by one increment will mute Jim,and set Larry to 0.8.

You can change the volume for a user once that user has spoken at least once during a session.
Attempts to change the volume of a user who has not spoken will be ignored.
If you are unable to hear a user speaking, you can edit the .barnard.yaml file in your home directory, after closing Barnard, and set the volume parameter to 1.0 for a particular user.

### Technical

The volume for each user is set via the audio session created for that user; no talking means no session means no settable volume.

## PM Mode

This mode sets the text to the left of your message entry box to [@username], where username is the name of the person you are PMing.
While a user is selected, both private messages and channel messages are displayed.
However, any messages you send will be delivered only to the selected user.
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

    go get -u github.com/bmmcginty/barnard

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

## Contributors

Tim Cooper (<tim.cooper@layeh.com>)
Brandon McGinty-Carroll (<bmmcginty.barnard@bmcginty.us>)
