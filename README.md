# Ticker

## Prerequisites

* [Docker](https://www.docker.com/)
* [Task](https://taskfile.dev/)

## Build

🚨 Ensure the prerequisites have been installed before continuing

To build the `ticker` binary run the following command from the repository 
root directory:

```bash
task app:build
```

To create the container image run:

```bash
task container:build
```

To push the container image run:

⚠️ Ensure you have authenticated with the registry host before continuing

```bash
task: container:push
```

The entire process of building the application and container image, and
pushing the container image by running:

```bash
task
```

### Deploy

The deployment leverages Kubernetes, and creates the following resources:

* deployment
* namespace
* service

The Kubernetes manifest can be found within the `deploy/` directory.
Before deploying ensure you have updated the manifest to include:

* `APIKEY`: Alpha Vantage API key (default: REPLACE_ME)
* `NDAYS`: Number of days (default: 3)
* `SYMBOL`: Stock symbol (default: MSFT)

To deploy `ticker` run:

```
kubectl apply -f deploy ticker.k8s.yaml
```

Note: The Kubernetes manifest is produced via [cdk8s](https://cdk8s.io/). The
repository containing the code that generated the manifest can be found 
[here](https://github.com/michaelcarruthers/cdk8s-ticker).

## Issues

The Advantage API returns `200` regardless if the API key has been rate limited.
As a result, `ticker` will panic.

## Resilience

The resiliency of `ticker` can be improved for production through the 
addition of:

* Caching
* Contexts
* Logging
* Signal handling
* Tests
* Timeouts