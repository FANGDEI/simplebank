postgres:
	docker pull postgres:14-alpine
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

clean:
	docker stop postgres
	docker rm postgres
	docker rmi postgres:14-alpine