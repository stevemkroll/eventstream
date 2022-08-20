# eventstream

```bash
# run kafka
docker compose up zookeeper kafka
```

```bash
# build and run test image
docker compose build test
docker compose up test

# or run with
docker compose run test
```

```bash
# run the consumer
go run main.go consumer --run
```

```bash
# send a message with the producer
go run main.go producer --message 'this is a message'
```
