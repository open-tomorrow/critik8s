# Helm Charts

This directory contains helm charts to install `critik8s`.

## RabbitMq Operator

This chart will install the RabbitMq operator to deploy and manage RabbitMQ clusters. See <https://github.com/rabbitmq/cluster-operator>

  ```bash
  helm install rabbitmq-operator charts/rabbitmq-operator/ -n rabbitmq --create-namespace
  ```

## RabbitMq Cluster

This chart will install the RabbitMq Custer, which provides the messaging services inside your cluster.

  ```bash
  helm install rabbitmq-cluster charts/rabbitmq-cluster/ --values settings.yaml -n rabbitmq --create-namespace
  ```

## Test RabbitMq Cluster

- Wait until the RabbitMq cluster is `ready`.

  ```bash
  kubectl -n rabbitmq get rabbitmqclusters.rabbitmq.com
  ```

- Run `perf-test` pod

    ```bash
    username="$(kubectl -n rabbitmq get secret rabbitmq-cluster-default-user -o jsonpath='{.data.username}' | base64 --decode)"
    password="$(kubectl -n rabbitmq get secret rabbitmq-cluster-default-user -o jsonpath='{.data.password}' | base64 --decode)"
    service="$(kubectl -n rabbitmq get service rabbitmq-cluster -o jsonpath='{.spec.clusterIP}')"
    kubectl -n rabbitmq run perf-test --image=pivotalrabbitmq/perf-test -- --uri amqp://$username:$password@$service
    ```

- Check the logs to see new messages

    ```bash
    kubectl -n rabbitmq logs --follow perf-test
    ```

### References

<https://www.rabbitmq.com/kubernetes/operator/quickstart-operator.html>
