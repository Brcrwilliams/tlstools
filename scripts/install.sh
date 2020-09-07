#!/bin/sh

usage() {
    echo "Usage: $0 [--target darwin|linux|windows|all] [--tool readpem|x509meta]"
    echo "Build binaries for the given tools and targets and place them in /usr/local/bin/"
    echo
    echo "  --tool       The tool to compile. If empty, compiles all tools."
}

while [ -n "$1" ]; do
    case "$1" in
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

cd "$(dirname "$0")/.." || exit 1

if [ -n "$TOOL" ]; then
    echo "Installing: $TOOL"
    go build -o "/usr/local/bin/${TOOL}" "./cmd/${TOOL}" || exit 1
    exit 0
fi

for dir in ./cmd/*; do
    t="$(basename "$dir")"
    echo "Installing: $t"
    go build -o "/usr/local/bin/${t}" "${dir}"
done
