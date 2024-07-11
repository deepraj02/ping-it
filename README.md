# PingIt: A URL Shortener CLI Service

PingIt is a URL shortener service that provides both a command-line interface (CLI) and a web API, written in Go. This project is designed to offer a simple and efficient way to shorten URLs directly from your terminal or through HTTP requests.

## Features

- **CLI for URL Shortening**: Easily shorten URLs from the command line.
- **Web API**: Integrate URL shortening into your applications with our HTTP API.
- **Docker Support**: Run PingIt in a containerized environment with Docker support.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go
- Docker
- Task

### Installing

Clone the repository to your local machine:

```sh
git clone https://github.com/deepraj02/pingit.git
cd pingit
```

## Local Development Setup

### Running the Containers

To start the API containers, you can use the Task command. This will build and run the API service along with its dependencies, such as the database service, in Docker containers.

```sh
task api-run
```

This command is defined in the [Taskfile.yml](Taskfile.yml) and will execute the necessary Docker commands to get your API service up and running.

### Building and Running the CLI

#### Building the CLI

To build the CLI application, you can use the following Task command:

```sh
task build-cli
```

This command compiles the CLI application and places the executable in the [`cli/bin`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdeepraj%2FDeveloper%2FGo%2Fping-it%2Fcli%2Fbin%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/deepraj/Developer/Go/ping-it/cli/bin") directory. The build process is defined in the [`Taskfile.yml`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdeepraj%2FDeveloper%2FGo%2Fping-it%2FTaskfile.yml%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/deepraj/Developer/Go/ping-it/Taskfile.yml") file.

#### Running the CLI

After building the CLI, you can run it using the following Task command:

```sh
task run-cli
```

This command executes the CLI application, allowing you to interact with the PingIt service from your command line. The specifics of this command are also detailed in the [Taskfile.yml](Taskfile.yml) file.
