up:
	docker-compose up -d

down:
	docker-compose down

reset:
	docker-compose down
	rm -rf ./mysql/data
	docker-compose up -d

logs:
	docker-compose logs -f

mysql:
	docker exec -it mysql8 mysql -umyuser -pmypassword goosedb

.PHONY: up down reset logs mysql
