#!/bin/bash
# barnard-sound.sh
# Description: Sounds and notification script for barnard.
#
# Copyright 2019, F123 Consulting, <information@f123.org>
# Copyright 2019, Storm Dragon, <storm_dragon@linux-a11y.org>
#
# This is free software; you can redistribute it and/or modify it under the
# terms of the GNU General Public License as published by the Free
# Software Foundation; either version 3, or (at your option) any later
# version.
#
# This software is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
# General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this package; see the file COPYING.  If not, write to the Free
# Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
# 02110-1301, USA.
#
#--code--
                                                                                                                                                                
# 1 is off, 0 is on
notify=0
sound=0

# Select notify type, notify-send or fenrir
notifyType="fenrir"

connect() {
    [[ $sound ]] && play -n synth .05 pl 1050 pl 1100 remix - pad 0 .05 repeat
    [[ $notify ]] && notify "${1}ed."
}

disconnect() {
    [[ $sound ]] && play "|sox -np synth .1 sin 350 sin 440 norm -4 pad 0 .1 repeat 10 remix -" fade p 0 1 .5
    [[ $notify ]] && notify "${1}ed."
}

is_function() {
    LC_ALL=C type "$1" 2> /dev/null | grep -q "$1 is a function"
}

join() {
    [[ $sound ]] && play "|sox -np synth .04 sin 1400 sin 2060 sin 2450 sin 2600 norm -8 remix - pad 0 .02 repeat 25" fade p 0 .75 .5
    [[ $notify ]] && notify "$2 ${1}ed the channel."
}

leave() {
    [[ $sound ]] && play -n synth .5 sin 480 sin 620 remix - norm -8 pad 0 0.5 repeat
    [[ $notify ]] && notify "$2 left the channel."
}

micdown() {
    [[ $sound ]] && play -qnV0 synth .25 sin G6:E5  norm -8
    [[ $notify ]] && notify "You have stopped transmitting."
}

micup() {
    [[ $sound ]] && play -qnV0 synth .25 sin E5:G6  norm -8
    [[ $notify ]] && notify "You are now transmitting."
}

msg() {
    [[ $sound ]] && play -n synth .3 sin 1290:1490 sin 1494:1294 remix - norm -8
    [[ $notify ]] && notify "$1 from $2: $3"
}

notify() {
if [[ "$notifyType" == "notify-send" ]]; then
    notify-send "$@"
else
    if [[ -e "/tmp/fenrirscreenreader-deamon.sock" ]]; then
        socatFile="/tmp/fenrirscreenreader-deamon.sock"
    else
        socatFile="$(find /tmp/ -maxdepth 1 -type s -name 'fenrirscreenreader-*.sock' | head -1)"
    fi
    echo "command say $@" | socat - UNIX-CLIENT:$socatFile
fi
}

pm() {
    [[ $sound ]] && play -n synth .5 sin 440 sin 480 remix - norm -8
    [[ $notify ]] && notify "$1 from $2: $3"
}

if is_function "$1" ; then
    eval "$1" "$1" "$2" "$3" &> /dev/null
else
    echo "The given barnard event has not yet been added."
fi

exit 0
