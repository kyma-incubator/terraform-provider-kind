.PHONY: build
build: 
	./before-commit.sh

.PHONY: ci-build
ci-build:
	./before-commit.sh ci

.PHONY: ci-pr
ci-pr: build

.PHONY: ci-master
ci-master: build

.PHONY: ci-release
ci-release: build

.PHONY: clean
clean:
	rm -rf bin

ci-testacc:
	./before-commit.sh ci testacc

testacc:
	./before-commit.sh testacc

install: GOOS=$(shell go env GOOS)
install: GOARCH=$(shell go env GOARCH)
ifeq ($(OS),Windows_NT)  # is Windows_NT on XP, 2000, 7, Vista, 10...
install: DESTINATION=$(APPDATA)/terraform.d/plugins/$(GOOS)_$(GOARCH)
else
install: DESTINATION=$(HOME)/.terraform.d/plugins/$(GOOS)_$(GOARCH)
endif
install: build
	@echo "==> Installing plugin to $(DESTINATION)"
	@mkdir -p $(DESTINATION)
	@cp ./bin/terraform-provider-kind $(DESTINATION)
