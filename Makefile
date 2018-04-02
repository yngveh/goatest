
.PHONY: dep
dep:
	dep ensure -v

gen:
	@goa gen github.com/yngveh/goatest/design

.PHONY: build
build:
	cd cmd/calcsvc && \
	go build

.PHONY: build-cli
build-cli:
	cd cmd/calccli && \
	go build

.PHONY: build-all
build-all: build build-cli

.PHONY: run
run: build
	./cmd/calcsvc/calcsvc



.PHONY: clean
clean:
	rm -fr gen

