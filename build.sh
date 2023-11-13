#!/bin/bash

VERSION="1.0.0"

PLATFORMS=(
	"darwin/amd64"
	"darwin/arm64"
	"linux/386"
	"linux/amd64"
	"linux/arm64"
	"windows/386"
	"windows/amd64"
	"windows/arm64"
)

mkdir -p build

for PLATFORM in "${PLATFORMS[@]}"; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}

    OUTPUT_NAME="build/rc_${VERSION}_${GOOS}_${GOARCH}"

    if [ $GOOS = "windows" ]; then
        OUTPUT_NAME+='.exe'
    fi

    echo "Compile for $GOOS $GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_NAME

    if [ $? -ne 0 ]; then
        echo "Error during compilation for $GOOS $GOARCH"
        exit 1
    fi
done

echo "Done!"
