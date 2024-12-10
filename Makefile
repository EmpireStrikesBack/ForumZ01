IMAGE_NAME=forum_i
CONTAINER_NAME=forum_c
PORT=8081

all: build run

build:
	npx tsc
	sudo docker build -t $(IMAGE_NAME) .

run:
	sudo docker run -p $(PORT):$(PORT) --name $(CONTAINER_NAME) $(IMAGE_NAME)

stop:
	sudo docker stop $(CONTAINER_NAME)

rm:
	sudo docker rm $(CONTAINER_NAME)

rmi:
	sudo docker rmi $(IMAGE_NAME)

clean: stop rm rmi

re: clean all

.PHONY: all build run stop rm rmi clean re
