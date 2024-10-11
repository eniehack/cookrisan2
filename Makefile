FLAGS =
GO = go
BINARIES = crawler
BINDIR = bin
SRC = $(shell find . -type f -name '*.go' -print)
.PHONY: clean pre-build

all: pre-build $(BINARIES)

pre-build:
	mkdir -p ./$(BINDIR)

$(BINARIES): pre-build $(SRC) 
	$(GO) $(FLAGS) build -o ./$(BINDIR)/$@ ./cmd/$@/main.go

clean:
	rm -rf ./$(BINDIR)