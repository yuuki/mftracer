PKG = github.com/yuuki/mkr-flow-tracer
COMMIT = $$(git describe --tags --always)
DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)

.PHONY: build
build:
	find cmd -d 1 | xargs -n1 -I{} go build -ldflags="$(BUILD_LDFLAGS)" $(PKG)/{}

.PHONY: test
test:
	go test -v ./...