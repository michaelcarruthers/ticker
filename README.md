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

## Issues

The Advantage API returns `200` regardless if the API key has been rate limited.
As a result, `ticker` will panic.