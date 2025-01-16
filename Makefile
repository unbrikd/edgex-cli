# The version is set by the VERSION file. If the file is not present, it defaults to 0.0.0.
VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)

# The git SHA is set by the current git commit SHA.
GIT_SHA=$(shell git rev-parse HEAD)

# Go variables for the current project
GO_OUT?="bin"
GO_APP="edgexcli"
GO_TESTFLAGS?=-race
GO_MODULE="github.com/unbrikd/edgex-cli"
GO_LDFLAGS_EXTRA=-X $(GO_MODULE).Version=$(VERSION) -X $(GO_MODULE).Revision=$(GIT_SHA)
GO_BUILDFLAGS_EXTRA=-trimpath -mod=readonly -buildmode=pie

gems-linux-cross-compile:
	@mkdir -p $(GO_OUT)
	@docker buildx build \
		--platform=$(GOOS)/$(GOARCH) \
		--build-arg GOOS=$(GOOS) \
		--build-arg GOARCH=$(GOARCH) \
		--build-arg GO_LDFLAGS_EXTRA="$(GO_LDFLAGS_EXTRA)" \
		--build-arg GO_BUILDFLAGS_EXTRA="$(GO_BUILDFLAGS_EXTRA)" \
		--build-arg GO_APP=$(GO_APP) \
		--output type=docker -t $(GO_APP) \
		-f ./docker/Dockerfile .

	@docker create --platform=$(GOOS)/$(GOARCH) --name temp-container $(GO_APP)
	@docker cp temp-container:/app/$(GO_APP) $(GO_OUT)/$(GO_APP)_$(GOOS)-$(GOARCH)
	@docker remove temp-container
	@docker rmi $(GO_APP)

# Build a binary for the edge gateway and VM running GEMS EdgeX Linux.
# The linked libraries are set to be compatible with the GEMS Linux configuration.
edge-gw-bin:
	@$(MAKE) gems-linux-cross-compile GOOS="linux" GOARCH="arm64" GO_LDFLAGS_EXTRA="-I /lib/ld-linux-aarch64.so.1 $(GO_LDFLAGS_EXTRA)"
	