#!/bin/sh

xdpyinfo -display $DISPLAY | grep 'dimensions:' | awk '{print $2}' | awk -F x '{print $2}'
