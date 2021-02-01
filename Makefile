NAME=wgrest
DESTDIR?=./dist
GOFMT_FILES?=$$(find . -name '*.go')

default: build

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(DESTDIR)/$(NAME)-linux-amd64 cmd/wgrest-server/main.go
linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o $(DESTDIR)/$(NAME)-linux-arm64 cmd/wgrest-server/main.go

build: validate
	go build -o $(DESTDIR)/$(NAME) cmd/wgrest-server/main.go

validate: fmtcheck vet lint

fmt:
	@echo "==> Applying gofmt requirements..."
	gofmt -s -w $(GOFMT_FILES)

vet:
	@echo "==> Checking that code complies with go vet requirements..."
	@go vet $$(go list ./... ) ; if [ $$? -gt 0 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

lint:
	@echo "==> Checking that code complies with golint requirements..."
	@GO111MODULE=off go get -u golang.org/x/lint/golint
	@if [ -n "$$(golint $$(go list ./...) | grep -v 'should have comment.*or be unexported' | tee /dev/stderr)" ]; then \
		echo ""; \
		echo "golint found style issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review."; \
    	exit 1; \
	fi

fmtcheck:
	@echo "==> Checking that code complies with gofmt requirements..."
	@if [ -n "$$(gofmt -s -l ${GOFMT_FILES})" ]; then \
		echo ""; \
		echo 'gofmt found format issues.'; \
		echo "Please execute the command: 'make fmt' to reformat code."; \
		exit 1; \
	fi

generate:
	@echo "==> Generating API code with go-swagger..."
	@GO111MODULE=off go get -u github.com/go-swagger/go-swagger
	@go-swagger generate server
	@if [ $$? -gt 0 ]; then \
		echo ""; \
		echo "go-swagger found issues. Please check the reported issues"; \
		echo "and fix them if necessary before generating the code for review."; \
    	exit 1; \
    else \
		echo ""; \
		echo "Please execute the command: 'make build' to build the code."; \
	fi

all: validate linux-amd64 linux-arm64 darwin

.PHONY: build vet fmt fmtcheck lint validate all generate
