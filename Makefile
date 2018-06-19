ifeq ($(ENV_FILE_PATH),)
	ENV_FILE_PATH = $(CURDIR)/env/default.env
endif
include $(ENV_FILE_PATH)
export $$(shell sed 's/=.*//' $(ENV_FILE_PATH))

PROJECT = mftracer
PKG = github.com/yuuki/$(PROJECT)
COMMIT = $$(git describe --tags --always)
DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)
CREDITS = ./CREDITS

.PHONY: build
build:
	go generate ./...
	find cmd -maxdepth 1 -mindepth 1 | xargs -n1 -I{} go build -ldflags="$(BUILD_LDFLAGS)" $(PKG)/{}

.PHONY: test
test:
	go test -v ./...

.PHONY: credits
credits:
	go get github.com/go-bindata/go-bindata/...
	_tools/credits > $(CREDITS)
ifneq (,$(git status -s $(CREDITS)))
	go generate -x ./...
endif

.PHONY: deploy
deploy: package deploy-sam

.PHONY: package
package:
	_tools/package $(S3_SAM_ARTIFACTS_BUCKET_NAME) $(PROJECT)

.PHONY: deploy-sam
deploy-sam:
	_tools/deploy $(ENV_FILE_PATH)

.PHONY: destroy
destroy:
	aws cloudformation delete-stack --stack-name $(PROJECT)

# vpc for trial
.PHONY: vpc/create
vpc/create:
	cfstacker create vpc
.PHONY: vpc/update
vpc/update:
	cfstacker update vpc
.PHONY: vpc/destroy
vpc/destroy:
	cfstacker delete vpc
