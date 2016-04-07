#!/bin/bash
PROJECT_ROOT=$(cd ${BASH_SOURCE[0]%/*} && pwd)

COMMAND=${1:-build}

case $COMMAND in
    build )
        export GOPATH=$PROJECT_ROOT

        go test algorithms/qsort
        go build algorithms/qsort
        go install algorithms/qsort

        go test algorithms/bubblesort
        go build algorithms/bubblesort
        go install algorithms/bubblesort

        go build -o $PROJECT_ROOT/bin/sorter sorter
    ;;
    clean )
        rm -rf $PROJECT_ROOT/bin/*
        rm -rf $PROJECT_ROOT/pkg/*
    ;;
    * )
        echo "Usage:
    ${BASH_SOURCE[0]} <build | clean>"
    ;;
esac