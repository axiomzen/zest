# Zest
Tests with a twist!

Zest is to assist with docker-driven workflow in both a local dev environment and in build servers.

## Requirements
Zest requires `zest`, `_zester`, `docker`, and `docker-compose` to be installed in your PATH. The Docker daemon must be running.

# Services and Projects
Zest works in 2 ways: operating on single services, and on multi-service projects.

## Service
A service is a single container containing a running binary. This could be a database, a webapp, etc.
Each service to be built and tested with zest need to have a `Zestfile` and a `Dockerfile`. The name of the service will be taken from the current directory name

### Commands
Command | Result
--------|-------
init-service | Creates the prerequisite files for the current directory to be a zestable service
build   | Runs the `Build()` script inside the build container
enter   | Launch an interactive shell inside the build container
test    | Runs the `Test()` script inside the test container
bundle  | Build the final container with the provided Dockerfile and tag with both version and latest
version | Execute `Version()` to see what the container will be tagged as


## Project
A project is a collection of services that run together to form an application stack.
Each project must container a `docker-compose.yml` that defines the services in the project, and a `Peelfile` for determining what services are the root of the project.
The `Peelfile` can also optinally specify where to find the relavent compose files.

### Commands
Command | Result
--------|--------
init-project | Creates the prerequisite files for the current directory to be a zestable project
integrate | Run integration tests on a project
run | Start the environment with docker-compose
stop | clean up the docker-compose environment
all | build, test, bundle all folders in the pwd that are services, then integrate

## Integration testing
Zest integration tests assume that your integration tests are written as a standalone service, and that service depends on the service to be tested in the `docker-compose` config.

To run integration tests, use `zest integrate`. This is equivilant to running `docker-compose up <integrate-service>` where `<integrate-service>` is the service tagged with `integrate:` in the project's `Peelfile`.
Zest will clean up the compose environment when the tests are done.
