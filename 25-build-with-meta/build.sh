#!/bin/bash

build() {
  #version=`git tag | tail -n 1`
  version="v1.0"
  git_hash=`git rev-parse HEAD`
  build_time=`date +%Y%m%d-%H:%M:%S`
  go_version=`go version`

  cmd="GOOS=darwin GOARCH=amd64 go build -o=./main -ldflags \"
    -X '25-build-with-meta/meta/config.Version=$version' \
    -X '25-build-with-meta/meta/config.GitHash=$git_hash' \
    -X '25-build-with-meta/meta/config.BuildTime=$build_time' \
    -X '25-build-with-meta/meta/config.GoVersion=$go_version' \""

  echo $cmd
  eval $cmd
}

build
echo "build finish"
