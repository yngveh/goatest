

.PHONY: dep
dep:
	dep ensure -v

.PHONY: generate
generate:
	@goagen app     -d github.com/yngveh/goatest/design
	@goagen swagger -d github.com/yngveh/goatest/design
	@goagen schema  -d github.com/yngveh/goatest/design
	@goagen main    -d github.com/yngveh/goatest/design

.PHONY: build
build:
	go build

.PHONY: clean
clean:
	rm -fr app swagger schema

