.PHONY: help check 
help:
	@echo Try targets: hi, init, optional_field

check:
	@version=`protoc --version | cut -f2 -d' ' | cut -f1,2 -d\.`; \
		if [ "$$version" != "3.15" ]; then \
			echo Warning: Expected version 3.15, got $$version; \
		fi

example: check
	protoc -I=. --go_out=. example.proto

PKG=github.com/cgthayer/grpc-edu/examplepb

optional_fields: example
	GOPATH=`pwd`/$(PKG):$${GOPATH} go build optional_fields.go

hi:
	go build hi.go

.PHONY: init
init:
	@echo Protoc Version $(protoc --version)
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go

