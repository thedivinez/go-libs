# Increment patch version by default
VERSION?=$(shell git describe --tags --abbrev=0 2>/dev/null | awk -F . '{OFS="."; $$NF+=1; print}')

ifeq ($(VERSION),)
	VERSION:=v0.1.0
endif

version-patch:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F . '{OFS="."; $$NF+=1; print}'))

version-minor:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F . '{OFS="."; $(NF-1)+=1; $$NF=0; print}'))

version-major:
	$(eval VERSION=$(shell git describe --tags --abbrev=0 | awk -F . '{$($NF-2)+=1; $(NF-1)=0; $$NF=0; print}'))

generate:
	templ generate
	go generate ./...

public:
	echo $(VERSION)
	git push
	git tag $(VERSION)
	git push --tags
	GOPROXY=proxy.golang.org go list -m github.com/thedivinez/go-libs@$(VERSION)

.PHONY: generate public