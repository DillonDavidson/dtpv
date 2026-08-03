BINDIR := bin
BINARIES := dtpv dtpvclean

.PHONY: all build clean install test

all: build

build:
	@mkdir -p $(BINDIR)
	go build -o $(BINDIR)/dtpv ./cmd/dtpv
	go build -o $(BINDIR)/dtpvclean ./cmd/dtpvclean

install:
	go install ./cmd/dtpv
	go install ./cmd/dtpvclean

test:
	go test ./...

clean:
	rm -rf $(BINDIR)
