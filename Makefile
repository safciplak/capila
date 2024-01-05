-include ./scripts/Makefile

# Docker identifier
DOCKERNAME="capila"

## cli:build: Builds the CLI file
cli\:build:
	$(START)
	@$(RUN_DEBUG) go build -o capila ./src/cli/build/executable.go
	$(END)
