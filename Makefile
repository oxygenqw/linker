run:
	docker-compose up -d

stop:
	docker-compose down
	
build: 
	docker-compose build --no-cache

logs:
	docker-compose logs -f

logs-linker:
	docker-compose logs -f linker

rebuild: build run
	