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
