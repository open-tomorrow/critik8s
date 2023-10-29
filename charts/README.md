# Helm Charts

This directory contains helm charts to install `critik8s`.

## RabbitMq Operator

This chart will install the RabbitMq operator to deploy and manage RabbitMQ clusters. See https://github.com/rabbitmq/cluster-operator

  ```
  helm install rabbitmq-operator rabbitmq-operator/ -n rabbitmq --create-namespace
  ```

## RabbitMq Cluster

This chart will install the RabbitMq Custer, which provides the messaging services inside your cluster.

  ```
  helm install rabbitmq-cluster rabbitmq-cluster/ -n rabbitmq --create-namespace
  ```

## Test RabbitMq Cluster

- Wait until the RabbitMq cluster is `ready`.

  ```
  kubectl -n rabbitmq get rabbitmqclusters.rabbitmq.com
  ```

- Run `perf-test` pod

    ```
    username="$(kubectl -n rabbitmq get secret rabbitmq-cluster-default-user -o jsonpath='{.data.username}' | base64 --decode)"
    password="$(kubectl -n rabbitmq get secret rabbitmq-cluster-default-user -o jsonpath='{.data.password}' | base64 --decode)"
    service="$(kubectl -n rabbitmq get service rabbitmq-cluster -o jsonpath='{.spec.clusterIP}')"
    kubectl -n rabbitmq run perf-test --image=pivotalrabbitmq/perf-test -- --uri amqp://$username:$password@$service
    ```

- Check the logs to see new messages

    ```
    kubectl -n rabbitmq logs --follow perf-test
    ```

### References

https://www.rabbitmq.com/kubernetes/operator/quickstart-operator.html
