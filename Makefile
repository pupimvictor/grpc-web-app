.PHONY: install build serve clean pack deploy ship

TAG?=$(shell date +%s)

export TAG

###################### server ######################
vendor:
	go mod vendor

build:
	GOOS=linux go build -o grpc-web-app-server ./cmd

serve: build
	./service1

clean:
	rm ./service1

pack:
	docker build -t us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server:$(TAG) .
	docker tag us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server:$(TAG) us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server:latest

upload:
	docker push us.gcr.io/nyt-adtech-dev/grpc-web-app/events-server

deploy:
	kubectl delete deployment events-server
	kubectl apply -f kube.yaml

ship: install build pack upload deploy

###################### client ######################

protoc:
	protoc event.proto --js_out=import_style=commonjs:./client/generated --grpc-web_out=import_style=commonjs,mode=grpcwebtext,out=event_grpc_pb.js:./client/generated --plugin=protoc-gen-grpc-web=/Users/208323/git/grpc-web/javascript/net/grpc/web/protoc-gen-grpc-web

