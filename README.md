# Training session receiver Service

## Introduction

A training session receiver service is a service that receives training session data from user's devices and browser form. The service stores it in a database and provides an API that allows to manage training session data.

## Features

- Receive training session data from user's devices and browser form.
- Store training session data in a database.
- Provide an API to manage training session data.

## Configuration

The service is configured with a [config.ini](config.ini) file found at `/app/config.ini` or another file specified using the --config flag.

## Installation

To install the service, you need to have Go installed on your machine. You can download and install Go from the official website: [https://golang.org/](https://golang.org/).

After installing Go, clone the repository and build the service using the following commands:

```bash
git clone https://github.com/dredfort42/rupl_training_session_receiver.git
cd rupl_training_session_receiver
go build -o training_session_receiver ./cmd/training_session_receiver/main.go
```

## Usage

### Running the service

To start the service, run the following command:

```bash
./training_session_receiver
```

The service will start and listen on the host and port specified in the configuration file.

### Running the service in Debug mode

To run the service in Debug mode, set the DEBUG environment variable to true before starting the service:

```bash
env DEBUG=true ./training_session_receiver
```

The service will start in Debug mode and print additional information to the console.

### Running the service with a specific configuration file

To run the service with a specific configuration file, set the --config flag with the path to the configuration file while starting the service:

```bash
./training_session_receiver --config /path/to/my_config.ini
```

### Running the service in Docker

To run the service in Docker, build the Docker image using the following command:

```bash
docker build -t training_session_receiver .
```

After building the Docker image, run the service using the following command:

```bash
docker run -p 4242:4242 training_session_receiver
```

The service will start and listen on port 4242.

### Running the Service with Docker Hub Images

To run the service using Docker Hub images, use the following command:

```bash
docker run -p 4242:4242 dredfort/rupl_training_session_receiver:latest
```

This will download the service from Docker Hub and start it, listening on port 4242.

## API

The API endpoints are described in the [openapi.yaml](/api/openapi.yml) file.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
