.PHONY: run

run:
	export CONFIG="config/config-local" && go run ./cmd/app/main.go