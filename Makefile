.PHONY: install build serve clean pack deploy ship

TAG?=$(shell date +%s)

export TAG

protoc:
	protoc event.proto --js_out=import_style=commonjs:./client/generated --grpc-web_out=import_style=commonjs,mode=grpcwebtext,out=event_grpc_pb.js:./client/generated --plugin=protoc-gen-grpc-web=/Users/208323/git/grpc-web/javascript/net/grpc/web/protoc-gen-grpc-web
	protoc --proto_path=$(GOPATH)/src --proto_path=. --gogo_out=plugins=grpc:./server event.proto

###################### server ######################
vendor:
	go mod vendor

server-build:
	GOOS=linux go build -o grpc-web-app-server ./cmd

server-pack:
	docker build -t us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server:$(TAG) ./server
	docker tag us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server:$(TAG) us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server:latest

server-upload:
	docker push us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server

server-deploy:
	kubectl delete deployment events-server
	kubectl apply -f ./server/kube.yaml

server-ship: server-build server-pack server-upload server-deploy

###################### client ######################

cli-install:
	cd client && npm install && npx webpack

cli-pack:
	docker build -t us.gcr.io/nyt-adtech-dev/grpc-web-app/events-cli:$(TAG) ./client
	docker tag us.gcr.io/nyt-adtech-dev/grpc-web-app/events-cli:$(TAG) us.gcr.io/nyt-adtech-dev/grpc-web-app/events-cli:latest

cli-upload:
	docker push us.gcr.io/nyt-adtech-dev/grpc-web-app/events-cli

cli-deploy:
	kubectl delete deployment events-cli
	kubectl apply -f ./cli/kube.yaml



