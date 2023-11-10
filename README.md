# critik8s

critik8s is a monitoring/alert tool for Kubernetes.

You decide which events will be reported as critical within your cluster.

## Development setup

### Requirements

- Node
- Helm
- k3d
- kubectl

## Install Node dependencies

```shell
npm install
```

## Create the k3d cluster

Create a new cluster; wait until traefik and other cluster service are up & running:

```shell
npm run cluster:create
```

## Deploy rabbitmq server

Deploy rabbitmq operator and cluster; wait until rabbitmq is up & running

```shell
npm run critik8s:create:ns && npm run rabbitmq:install
```

## Deploy critik8s

Deploy critik8s components:

```shell
npm run critik8s:install
```
