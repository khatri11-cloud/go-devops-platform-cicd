run:
	go run ./cmd/server

build:
	go build -o server ./cmd/server

docker-build:
	docker build -t devops-platform:v1 .

docker-run:
	docker run -p 8080:8080 devops-platform:v1

clean:
	rm -f server