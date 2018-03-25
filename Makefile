PKG = github.com/yuuki/mftracer
COMMIT = $$(git describe --tags --always)
DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)

.PHONY: build
build:
	go generate ./...
	find cmd -maxdepth 1 -mindepth 1 | xargs -n1 -I{} go build -ldflags="$(BUILD_LDFLAGS)" $(PKG)/{}

.PHONY: test
test:
	go test -v ./...

.PHONY: deps
deps:
	go get github.com/go-bindata/go-bindata/...