
src := cmd/svg-to-jsx/main.go

# Strip debug
GO_FLAGS += "-ldflags=-s -w"
# No build path in executable
GO_FLAGS += -trimpath

all: \
	bin/svg-to-jsx-linux-arm64 \
	bin/svg-to-jsx-linux-amd64 \
	bin/svg-to-jsx-darwin-arm64 \
	bin/svg-to-jsx-darwin-amd64

bin:
	mkdir -p bin

$(src): 

bin/svg-to-jsx-linux-arm64: bin $(src)
		$(MAKE) GOOS=linux GOARCH=arm64 binary

bin/svg-to-jsx-linux-amd64: bin $(src)
		$(MAKE) GOOS=linux GOARCH=amd64 binary

bin/svg-to-jsx-darwin-amd64: bin $(src)
		$(MAKE) GOOS=darwin GOARCH=amd64 binary

bin/svg-to-jsx-darwin-arm64: bin $(src)
		$(MAKE) GOOS=darwin GOARCH=arm64 binary



binary: $(src)
	CGO_ENABLED=0 GOOS="$(GOOS)" GOARCH="$(GOARCH)" go build $(GO_FLAGS) -o "bin/svg-to-jsx-$(GOOS)-$(GOARCH)" ./cmd/svg-to-jsx

clean:
	rm -f bin/*


.PHONY: test
test:
	go run ./cmd/svg-to-jsx -o ./sample -v ./sample