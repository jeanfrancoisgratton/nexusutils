#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY=nxrmutils
fi

go build -o /opt/bin/${BINARY} .
