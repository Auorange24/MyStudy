# kafka

## kafka install

### install kafka with docker

```docker
docker pull apache/kafka:4.0.0
docker run -p 9092:9092 -name my-kafka apache/kafka:4.0.0
docker exec -it my-kafka /bin/bash
```
