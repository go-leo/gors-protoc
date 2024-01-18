COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# generate api handler; api-handler-{api version}
.PHONY: api
api:
	@echo "generate grpc..."
	@protoc --proto_path=. \
           --proto_path=../third_party \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           $(shell find . -name "*.proto")
	@echo "generate grpc finsh."

	@echo "generate http..."
	@protoc --proto_path=. \
           --proto_path=../third_party \
           --go_out=paths=source_relative:. \
           --go-gorsp_out=paths=source_relative:. \
           $(shell find . -name "*.proto") && \
		   protoc-go-inject-tagp -input="*.pb.go"
	@echo "generate http finsh."

	@echo "generate vaildator..."
	@protoc \
		-I . \
		-I ../third_party \
		--go_out=paths=source_relative:. \
		--validate_out="lang=go,paths=source_relative:."  \
		$(shell find . -name "*.proto")
	@echo "generate vaildator finsh"

	@echo "generate swagger..."
	@protoc --proto_path=. \
           --proto_path=../third_party \
           --swagger_out=logtostderr=true,allow_delete_body=true:./ \
           $(shell find . -name "*.proto")
	@echo "generate swagger finsh"
