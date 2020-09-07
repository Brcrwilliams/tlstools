#!/bin/sh

usage() {
    echo "Usage: $0 [--target darwin|linux|windows|all] [--tool readpem|x509meta]"
    echo "Build binaries for the given tools and targets and output to bin/"
    echo
    echo "  --target     The target OS to compile for. If empty, compiles for host OS."
    echo "  --tool       The tool to compile. If empty, compiles all tools."
}

while [ -n "$1" ]; do
    case "$1" in
        --target)
            shift
            TARGET="$1"
            ;;
        --tool)
            shift
            TOOL="$1"
            ;;
        *)
            echo "Unknown argument: $1"
            usage
            exit 1
            ;;
    esac
    shift
done

build() {
    local tool="$1"
    if [ -n "$tool" ]; then
        name="${tool}-${GOOS}-${GOARCH}"
        echo "Building: $name"
        go build -o "bin/${name}" "./cmd/${tool}"
        return 0
    fi

    for dir in ./cmd/*; do
        t="$(basename "$dir")"
        name="${t}-${GOOS}-${GOARCH}"
        echo "Building: $name"
        go build -o "bin/${name}" "${dir}"
    done
}

cd "$(dirname "$0")/.." || exit 1

mkdir -p bin
if [ -z "$TARGET" ]; then
    TARGET="$(go env GOOS)"
fi

if [ "$TARGET" != all ]; then
    GOOS="$TARGET" GOARCH=amd64 build "$tool"
    exit 0
fi

for target in darwin linux windows; do
    GOOS="$target" GOARCH=amd64 build "$tool"
done
