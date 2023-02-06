set -e

go run genoverlay.go > /tmp/layla-overlay.json
env GOOS=linux GOARCH=arm64 CGO_ENABLED=1 \
    CGO_CFLAGS='-fno-common -fno-short-enums -ffunction-sections -fdata-sections -fPIC -g -O3' \
    go build -buildmode=c-archive -overlay=/tmp/layla-overlay.json -o=layla.a -tags=nintendosdk
gcc -o layla main.c layla.a -lpthread 
