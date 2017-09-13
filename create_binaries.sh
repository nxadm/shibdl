#!/bin/bash -e
set -xv
mkdir -p binaries
APP=`basename $(pwd)`
BIN="binaries/$APP"
PLATFORMS=("windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "linux/amd64" "linux/386")

function build {
    GOOS=$1
    GOARCH=$2
    OUTPUT="${BIN}-${GOOS}-${GOARCH}"
    if [ $GOOS = "windows" ]; then
        OUTPUT+='.exe'
    fi
    GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT
    sha512sum $OUTPUT > $OUTPUT.sha512
}

for i in ${PLATFORMS[@]}; do
    PLATFORMS_SPLIT=(${i//\// })
    GOOS=${PLATFORMS_SPLIT[0]}
    GOARCH=${PLATFORMS_SPLIT[1]}
    build $GOOS $GOARCH
done
