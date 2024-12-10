ifeq ($(VERSION),)
     VERSION:=$(shell git describe --tags --abbrev=0 | awk -F .   '{OFS="."; $$NF+=1; print}')
endif

public:
	echo $(VERSION)
	git push
	git tag $(VERSION)
	git push --tags
	GOPROXY=proxy.golang.org go list -m github.com/thedivinez/go-libs@$(VERSION)

proto-gen:
	protoc   \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=require_unimplemented_servers=false:. \
	--go-grpc_opt=paths=source_relative proto/messaging.proto

proto-clean:
	protoc-go-inject-tag -remove_tag_comment -input="proto/*.pb.go"

proto: proto-gen | proto-clean

.PHONY: proto