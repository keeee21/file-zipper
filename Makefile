build:
	cd docker && bash compose.sh build

up:
	cd docker && bash compose.sh up

front:
	docker exec -it file-zipper_frontend sh

backend:
	docker exec -it file-zipper_api sh

stop:
	cd docker && docker compose stop