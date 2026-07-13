SHELL := /bin/bash
GO ?= go
GO_PACKAGES := ./...
UNIT_PACKAGES := ./internal/...
INTEGRATION_PACKAGES := ./test/integration/...
BIN_DIR := bin
COVERAGE_FILE ?= coverage.out
COVERAGE_THRESHOLD ?= 80.0
PACKAGE_FILE ?= classroom3.tar.gz

CMDS := 01_errors 02_arrays 03_structs 04_new_make 05_slices 06_loops 07_functions 08_panics

.PHONY: help deps-check mod-check fmt fmt-check vet test test-unit test-integration test-race coverage coverage-check build package clean run-all ci $(addprefix run-,$(CMDS))

help:
	@echo "Available commands:"
	@echo "  make run-all              - run all cmd examples"
	@echo "  make test-unit            - unit tests for internal packages"
	@echo "  make test-integration     - integration tests for cmd output"
	@echo "  make ci                   - full local CI"
	@echo "  make build                - build all cmd examples"

deps-check:
	$(GO) mod download
	$(GO) mod verify

mod-check:
	$(GO) mod tidy
	@if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then \
		git diff --exit-code -- go.mod; \
		if [ -f go.sum ]; then git diff --exit-code -- go.sum; fi; \
	else \
		echo "Skipping git diff because this directory is not a git repository"; \
	fi

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './$(BIN_DIR)/*')

fmt-check:
	@files="$$(gofmt -l $$(find . -name '*.go' -not -path './$(BIN_DIR)/*'))"; \
	if [ -n "$$files" ]; then echo "Go files are not formatted:"; echo "$$files"; exit 1; fi

vet:
	$(GO) vet $(GO_PACKAGES)

test: test-unit test-integration

test-unit:
	$(GO) test $(UNIT_PACKAGES)

test-integration:
	$(GO) test $(INTEGRATION_PACKAGES)

test-race:
	$(GO) test -race $(UNIT_PACKAGES)

coverage:
	$(GO) test $(UNIT_PACKAGES) -covermode=atomic -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -func=$(COVERAGE_FILE)

coverage-check: coverage
	@coverage="$$(go tool cover -func=$(COVERAGE_FILE) | awk '/^total:/ {gsub("%", "", $$3); print $$3}')"; \
	awk -v coverage="$$coverage" -v threshold="$(COVERAGE_THRESHOLD)" 'BEGIN { \
		if (coverage + 0 < threshold + 0) { printf "coverage %.1f%% is below threshold %.1f%%\n", coverage, threshold; exit 1 } \
		printf "coverage %.1f%% is enough; threshold %.1f%%\n", coverage, threshold; \
	}'

run-all:
	@for cmd in $(CMDS); do echo "== $$cmd =="; $(GO) run ./cmd/$$cmd; done

run-01_errors:
	$(GO) run ./cmd/01_errors
run-02_arrays:
	$(GO) run ./cmd/02_arrays
run-03_structs:
	$(GO) run ./cmd/03_structs
run-04_new_make:
	$(GO) run ./cmd/04_new_make
run-05_slices:
	$(GO) run ./cmd/05_slices
run-06_loops:
	$(GO) run ./cmd/06_loops
run-07_functions:
	$(GO) run ./cmd/07_functions
run-08_panics:
	$(GO) run ./cmd/08_panics

build:
	@mkdir -p $(BIN_DIR)
	@for cmd in $(CMDS); do $(GO) build -o $(BIN_DIR)/$$cmd ./cmd/$$cmd; done

package: build
	tar -czf $(PACKAGE_FILE) -C $(BIN_DIR) $(CMDS)

ci: deps-check mod-check fmt-check vet test-unit test-integration test-race coverage-check build package

clean:
	rm -rf $(BIN_DIR) $(COVERAGE_FILE) $(PACKAGE_FILE)
