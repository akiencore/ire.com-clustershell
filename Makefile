.PHONY: all
all:
	@$(MAKE) --no-print-directory deps
	@$(MAKE) --no-print-directory clshsheduler
	@$(MAKE) --no-print-directory clshexecutor

.PHONY: deps
deps:
	go mod download

.PHONY: clshsheduler
clshsheduler:
	go build ./cmd/scheduler/clshscheduler.go

.PHONY: clshexecutor
clshexecutor:
	go build ./cmd/executor/clshexecutor.go

