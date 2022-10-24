redis:
	docker pull redis
	docker run --name simplebank-redis -d -p 7779:6379 --network bank-network redis --requirepass "ff1639579565"

clean:
	docker stop simplebank-redis
	docker rm simplebank-redis
	docker rmi redis:latest