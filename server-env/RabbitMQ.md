#### RabbitMQ

```shell

docker pull rabbitmq:3.9.13-management

docker run -d --name my-rabbitmq -p 5672:5672 -p 15672:15672 -v /Users/cc/rabbitmq:/var/lib/rabbitmq --hostname my-rabbitmq-node1 -e RABBITMQ_DEFAULT_VHOST=my-rabbitmq-vhost -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin rabbitmq:3.9.13-management

```
