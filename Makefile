buf:
	buf generate --path proto/org

submodules:
	git submodule update --init --recursive

build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 \
		docker-compose up --rm --build