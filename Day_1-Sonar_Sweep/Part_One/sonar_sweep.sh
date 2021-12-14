#!/bin/sh

let result=0
while read val; do
    if [[ $val > $last ]]; then
        (( result++ ))
    fi
    last=$val
done

echo $result
