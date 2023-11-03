# critik8s

monitoring tool for k8s

## k3d setup

### Requirements

- Docker, Docker compose
- Helm
- k3d
- kubectl
- Go (1.20+)

## Create the cluster

You create the `k3d-critik8s` cluster with:

```shell
make k3d-start
```

To specify the number of agents (worker nodes):

```shell
make k3d-start AGENT_NODES=2
```

> **WARNING**: the command updates your `.kube/config` with the credentials of
> the just created `k3d-critik8s` cluster and sets its context as default.

### Delete the cluster

```shell
make k3d-delete
```
