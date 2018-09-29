#!/bin/bash

APP_NAME=jvm-monitor
APP_VERSION=1.0-beta1

DIST_DIR=dist

gobuild() {
	ext=""
	if [ "$1" == "windows" ]; then
		ext=".exe"
	fi
	GOOS=$1 GOARCH=$2 go build -o $DIST_DIR/$1/$2/${APP_NAME}${ext} ${FLAGS} \
	-ldflags \
	"\
	-X 'main.appName=${APP_NAME}' \
	-X 'main.appVersion=${APP_VERSION}' \
	" \
	.
}

cd "$( dirname "$0" )"
rm -rf $DIST_DIR
export CGO_ENABLED=0

gobuild darwin amd64
gobuild linux amd64
gobuild linux 386
# gobuild windows amd64
# gobuild windows 386
