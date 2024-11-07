#!/bin/bash

count=$1

if [ $count -gt 100 ]; then
    count=100
fi

for (( i=1; i<=$count; i++ )); do
    echo "$i times I've printed danilaristikartano"
done
