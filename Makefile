create-kafka-topic:
	docker exec -it kafka kafka-topics --create --topic greencity-topic --partitions 4 --replication-factor 1 --bootstrap-server kafka:9092

delete-kafka-topic:
	docker exec -it kafka kafka-topics --delete --topic greencity-topic --bootstrap-server kafka:9092

describe-kafka-topic:
	docker exec -it kafka kafka-topics --describe --topic greencity-topic --bootstrap-server kafka:9092

consume-kafka-topic:
	docker exec -it kafka kafka-console-consumer --topic greencity-topic --bootstrap-server kafka:9092

build:
	docker build -t greencity:latest .

run-compose:
	docker-compose up -d $(name)

down-compose:
	docker-compose down $(name)

clean:
	docker system prune -a --volumes --force
