COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# generate api handler; api-handler-{api version}
.PHONY: api
api:
	@echo "generate api..."
	@protoc --proto_path=. \
           --proto_path=../third_party \
           --go_out=paths=source_relative:. \
           --go-gors_out=paths=source_relative:. \
           $(shell find . -name "*.proto") && \
		   protoc-go-inject-tag -input="*.pb.go"
	@echo "generate api finsh."