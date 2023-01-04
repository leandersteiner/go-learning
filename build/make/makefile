SHELL := /bin/bash

run:
	go run main.go

build:
	go build -ldflags "-X main.build=local"

VERSION := 1.0

service:
	docker build \
		-f infra/docker/Dockerfile \
		-t service-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

minikube-up:
	minikube start

minikube-down:
	minikube stop

minikube-load:
	minikube image load service-amd64:$(VERSION)

minikube-apply:
	cat infra/k8s/base/service-pod/base-service.yaml | kubectl apply -f -

minikube-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

minikube-status-service:
	kubectl get pods -o wide --watch --namespace=service-system

minikube-logs:
	kubectl logs -l app=service --all-containers=true -f --tail=100 --namespace=service-system

minikube-restart:
	kubectl rollout restart deployment service-pod --namespace=service-system
