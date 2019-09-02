#!/bin/bash

connect() {
    play -n synth .05 pl 1050 pl 1100 remix - pad 0 .05 repeat
}

disconnect() {
    play "|sox -np synth .1 sin 350 sin 440 norm -4 pad 0 .1 repeat 10 remix -" fade p 0 1 .5
}

is_function() {
    LC_ALL=C type "$1" 2> /dev/null | grep -q "$1 is a function"
}

join() {
    play "|sox -np synth .04 sin 1400 sin 2060 sin 2450 sin 2600 norm -8 remix - pad 0 .02 repeat 25" fade p 0 .75 .5
}

leave() {
    play -n synth .5 sin 480 sin 620 remix - norm -8 pad 0 0.5 repeat
}

micdown() {
    play -qnV0 synth .25 sin G6:E5  norm -8
}

micup() {
    play -qnV0 synth .25 sin E5:G6  norm -8
}

msg() {
    play -n synth .3 sin 1290:1490 sin 1494:1294 remix - norm -8
}

pm() {
    play -n synth .5 sin 440 sin 480 remix - norm -8
}

if is_function "$1" ; then
    eval "$1" &> /dev/null
else
    echo "The given barnard event has not yet been added."
fi

is_function "$1"
exit 0
