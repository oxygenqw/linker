start:
	docker-compose up -d

stop:
	docker-compose down
	
build: 
	docker-compose build --no-cache

rebuild: build start