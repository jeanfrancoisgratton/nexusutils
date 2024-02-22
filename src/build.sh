#!/usr/bin/env sh



if [ "$#" -gt 0 ]; then
    BINARY=nexusutils
fi

go build -o /opt/bin/${BINARY} .
