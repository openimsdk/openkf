HUGO_VERSION      = $(shell grep ^HUGO_VERSION netlify.toml | tail -n 1 | cut -d '=' -f 2 | tr -d " \"\n")
NODE_BIN          = node_modules/.bin
NETLIFY_FUNC      = $(NODE_BIN)/netlify-lambda

# The CONTAINER_ENGINE variable is used for specifying the container engine. By default 'docker' is used
# but this can be overridden when calling make, e.g.
# CONTAINER_ENGINE=podman make container-image
CONTAINER_ENGINE ?= docker
IMAGE_REGISTRY ?= gcr.io/k8s-staging-sig-docs
IMAGE_VERSION=$(shell scripts/hash-files.sh Dockerfile Makefile | cut -c 1-12)
CONTAINER_IMAGE   = $(IMAGE_REGISTRY)/k8s-website-hugo:v$(HUGO_VERSION)-$(IMAGE_VERSION)
# Mount read-only to allow use with tools like Podman in SELinux mode
# Container targets don't need to write into /src
CONTAINER_RUN     = "$(CONTAINER_ENGINE)" run --rm --interactive --tty --volume "$(CURDIR):/src:ro,Z"

CCRED=\033[0;31m
CCEND=\033[0m

# Docker buildx related settings for multi-arch images
DOCKER_BUILDX ?= docker buildx

.PHONY: all build build-preview help serve

REPO_ROOT:=${CURDIR}/..
# setup go for managing hugo
# go1.9+ can autodetect GOROOT, but if some other tool sets it ...
GOROOT:=
# enable modules
GO111MODULE=on
# disable CGO by default for static binaries
CGO_ENABLED=0
DOCKER       = docker
HUGO_VERSION = 0.60.0
DOCKER_IMAGE = jojomi/hugo:$(HUGO_VERSION)
DOCKER_RUN   = $(DOCKER) run --rm --interactive --tty --volume $(realpath $(CURDIR)/..):/src -p 1313:1313 --workdir /src/site --entrypoint=hugo --platform linux/amd64 $(DOCKER_IMAGE)

HUGO_BIN:=$(REPO_ROOT)/bin/hugo

$(HUGO_BIN):
	go build -o $(HUGO_BIN) github.com/gohugoio/hugo

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

module-check: ## Check if all of the required submodules are correctly initialized.
	@git submodule status --recursive | awk '/^[+-]/ {err = 1; printf "\033[31mWARNING\033[0m Submodule not initialized: \033[34m%s\033[0m\n",$$2} END { if (err != 0) print "You need to run \033[32mmake module-init\033[0m to initialize missing modules first"; exit err }' 1>&2

module-init: ## Initialize required submodules.
	@echo "Initializing submodules..." 1>&2
	@git submodule update --init --recursive --depth 1

all: build ## Build site with production settings and put deliverables in ./public

build: module-check ## Build site with non-production settings and put deliverables in ./public
	$(HUGO_BIN) --cleanDestinationDir --minify --environment development

build-preview: module-check ## Build site with drafts and future posts enabled
	$(HUGO_BIN) --cleanDestinationDir --buildDrafts --buildFuture --environment preview

deploy-preview: ## Deploy preview site via netlify
	GOMAXPROCS=1 hugo --cleanDestinationDir --enableGitInfo --buildFuture --environment preview -b $(DEPLOY_PRIME_URL)

functions-build: ## Build functions
	$(NETLIFY_FUNC) build functions-src

hugo: $(HUGO_BIN) ## Build site with production settings and put deliverables in ./public
	$(HUGO_BIN) --cleanDestinationDir --minify --environment production

serve: hugo ## Serve site at http://localhost:1313/
	$(HUGO_BIN) server --bind="0.0.0.0" \
	--ignoreCache \
	--buildFuture \
	--disableFastRender

.PHONY: build serve hugo module-check module-init