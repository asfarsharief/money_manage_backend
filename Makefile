#Makefile to run the application, run test cases, and build docker image 
#=======================================================================

dep:
	@go get ./...
	
run:
	@go run lib/main.go serve

test:
	@go test -coverprofile=.code_coverage.out ./...

show-code-coverage: test
	@mkdir -p .code_coverage
	@go tool cover -html=.code_coverage.out -o .code_coverage/index.html
	@cd .code_coverage && python -m SimpleHTTPServer 3001

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo -o bin/app lib/main.go

image: build
	docker build -t money_manage_backend:latest -f docker/Dockerfile .

run-docker-image: image
	docker run -p 3030:3030 money_manage_backend:latest