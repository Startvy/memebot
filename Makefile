build_meme:
	mkdir -p build
	sudo docker-compose -f docker-compose.build.yml up
	sudo chmod +x build/memebot
docker:
	docker-compose up
docker-prod:
	docker-compose up -d