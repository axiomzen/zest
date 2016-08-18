# Zest
Tests with a twist!

Zest is to assist with docker-driven workflow in both a local dev environment and in build servers.

## Requirements
Zest requires only itself (`zest`), `_zester`, and `docker` to be installed in your PATH. The Docker daemon must be running.

Each component to be built and tested with zest need to have a `Zestfile` and a `Dockerfile`

## Commands
Command | Result
--------|-------
build   | Runs the `Build()` script inside the build container
enter   | Launch an interactive shell inside the build container
test    | Runs the `Test()` script inside the test container
bundle  | Build the final container with the provided Dockerfile and tag with both version and latest
version | Execute `Version()` to see what the container will be tagged as
