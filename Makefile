BINARY_NAME=azutils
VERSION=$(shell git describe --tags --always)
COMMIT=$(shell git rev-parse HEAD)
DATE=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}"

.PHONY: build-all
build-all: clean build-windows build-linux build-mac

.PHONY: clean
clean:
	rm -rf dist/

.PHONY: build-windows
build-windows:
	mkdir -p dist/windows
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o dist/windows/${BINARY_NAME}.exe
	zip -j dist/${BINARY_NAME}_windows_amd64.zip dist/windows/${BINARY_NAME}.exe

.PHONY: build-linux
build-linux:
	mkdir -p dist/linux
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o dist/linux/${BINARY_NAME}
	tar -czf dist/${BINARY_NAME}_linux_amd64.tar.gz -C dist/linux ${BINARY_NAME}

.PHONY: build-mac
build-mac:
	mkdir -p dist/darwin
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o dist/darwin/${BINARY_NAME}
	tar -czf dist/${BINARY_NAME}_darwin_amd64.tar.gz -C dist/darwin ${BINARY_NAME}
