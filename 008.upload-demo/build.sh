#!/bin/bash
PROJECT_ROOT=$(cd ${BASH_SOURCE[0]%/*} && pwd)

COMMAND=${1:-build}

case $COMMAND in
    build )
        export GOPATH=$PROJECT_ROOT

        go test tpl
        go build tpl
        go install tpl

        go build -o $PROJECT_ROOT/bin/server server
    ;;
    clean )
        rm -rf $PROJECT_ROOT/bin/*
        rm -rf $PROJECT_ROOT/pkg/*
        rm -rf $PROJECT_ROOT/data/*
    ;;
    * )
        echo "Usage:
    ${BASH_SOURCE[0]} <build | clean>"
    ;;
esac
