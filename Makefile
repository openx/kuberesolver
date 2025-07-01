

build:
	docker build -t grpc-test-server:latest -f test/server/Dockerfile .
	docker build -t grpc-test-client:latest -f test/client/Dockerfile .

apply-all:
	kubectl apply -f test/deployments/server-deployment.yaml
	kubectl apply -f test/deployments/client-deployment.yaml

delete-dep:
	kubectl delete deployment server
	kubectl delete deployment client



