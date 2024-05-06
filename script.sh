docker compose up
docker compose exec kafka kafka-topics.sh --create --topic meu-topico --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092
docker compose exec kafka kafka-console-producer.sh --topic meu-topico --bootstrap-server localhost:9093
