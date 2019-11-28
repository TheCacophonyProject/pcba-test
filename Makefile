.PHONY: install-packr
install-packr:
	go get github.com/gobuffalo/packr/packr

.PHONY: release
release: install-packr
	curl -sL https://git.io/goreleaser | bash
