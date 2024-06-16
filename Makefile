PROJECT=zenkoo::svc.proto
PREFIX=$(shell pwd)
VERSION=$(shell git describe --match 'v[0-9]*'  --always)
DEFAULT_BRANCH=$(shell git symbolic-ref --short -q HEAD)

ifndef BRANCH
	BRANCH=$(DEFAULT_BRANCH)
endif

ifdef CI_COMMIT_REF_SLUG
	BRANCH=$(CI_COMMIT_REF_SLUG)
endif

ifndef GO
	GO=go
endif

ifndef GOFMT
	GOFMT=gofmt
endif

ifndef PROTOC
	PROTOC=protoc
endif

ifndef GIT
	GIT=git
endif

SOURCE_DIR=$(PREFIX)
DOCS_DIR=$(PREFIX)/docs
PROTOS=$(wildcard $(SOURCE_DIR)/*.proto)

.PHONY: all summary micro build doc
.DEFAULT: all

# Targets
all: summary build doc

summary:
	@printf "\033[1;37m  == \033[1;32m$(PROJECT) \033[1;33m$(VERSION) \033[1;37m==\033[0m\n"
	@printf "    Go          : \033[1;37m`$(GO) version`\033[0m\n"
	@printf "    Git         : \033[1;37m$(shell git version)\033[0m\n"
	@printf "    Branch      : \033[1;37m$(BRANCH)\033[0m\n"
	@echo

build:
	@printf "\033[1;36m  Compiling protos ...\033[0m\n"
	@for f in $(shell find . -name '*.proto' ! -path '*/third_party/*') ; do \
		printf "    Compiling <$${f}>\n"  && \
		$(PROTOC) --go_out=. --go-grpc_out=. --micro_out=. $${f} ; \
	done
	@echo
	@printf "\033[1;36m  Generating services ...\033[0m\n"
	@$(GO) run generate.go
	@echo

micro:
	@$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@$(GO) install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	@$(GO) install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	@$(GO) install github.com/envoyproxy/protoc-gen-validate@latest

doc:
	@printf "\033[1;36m  Generating document ...\033[0m\n"
	@mkdir -p $(PREFIX)/doc
	@$(PROTOC) --doc_opt=markdown,doc.md --doc_out=./doc $(shell find . -name '*.proto')
	@echo
