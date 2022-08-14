#!/bin/sh
# testing variables
if [ "$1" == "-device" ]; then
    ./agent -s "$2"
fi
if [ "$1" == "-check" ]; then
    echo "Run check"
    while true; do
        my_string=./vpnService
        # shellcheck disable=SC2143
        # shellcheck disable=SC2009
        if [[ -z $(ps -ef | grep './[a]gent') ]]
        then
            echo "Выключен\n"
        else
            echo "Работает"
        fi
        sleep 1d
    done
fi