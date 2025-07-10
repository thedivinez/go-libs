ifeq ($(VERSION),)
     VERSION:=$(shell git describe --tags --abbrev=0 | awk -F .   '{OFS="."; $$NF+=1; print}')
endif

public:
	echo $(VERSION)
	git push
	git tag $(VERSION)
	git push --tags
	GOPROXY=proxy.golang.org go list -m github.com/thedivinez/go-libs@$(VERSION)