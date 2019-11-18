#!/bin/bash

dogo() {
       #wire ./internal/serviceLocator
       go fmt
       go generate       
       gotest -v ./internal/composeservice
}



dogo

inotifywait --exclude "[^g].$|[^o]$" -m -r -e close_write ./ |
    while read path action file; do
           dogo
    done
