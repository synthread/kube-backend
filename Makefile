
build:
	mkdir -p out
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o out/kube-backend main.go

docker:
	docker build --arch arm64 -t ghcr.io/synthread/kube-backend:latest .
	docker push ghcr.io/synthread/kube-backend:latest