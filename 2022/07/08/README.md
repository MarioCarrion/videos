# Software Architecture: Claim Check Pattern

Run Kafka locally:

```
docker run \
  -d \
  --rm \
  -p 2181:2181 \
  wurstmeister/zookeeper:3.4.6
```

```
docker run \
  -d \
  --rm \
  -p 9092:9092 \
  -e "KAFKA_ZOOKEEPER_CONNECT=host.docker.internal:2181" \
  -e "KAFKA_ADVERTISED_HOST_NAME=localhost" \
  -e "KAFKA_ADVERTISED_PORT=9092" \
  -e "KAFKA_AUTO_CREATE_TOPICS_ENABLE=true" \
  -e "KAFKA_CREATE_TOPICS=files:1:1" \
  wurstmeister/kafka:2.13-2.7.0
```
