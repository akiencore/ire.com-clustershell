.PHONY: all
all:
	@$(MAKE) --no-print-directory deps
	@$(MAKE) --no-print-directory generatekeypairs
	@$(MAKE) --no-print-directory clshscheduler
	@$(MAKE) --no-print-directory clshexecutor

.PHONY: deps
deps:
	go mod download
	
.PHONY: generatekeypairs
generatekeypairs:
	go run ./cmd/keygen/generateKeyPairs.go > ./crypting/keys.go

.PHONY: clshscheduler
clshscheduler:
	go build ./cmd/scheduler/clshscheduler.go

.PHONY: clshexecutor
clshexecutor:
	go build ./cmd/executor/clshexecutor.go



