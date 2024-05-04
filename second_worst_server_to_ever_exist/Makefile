.PHONY: build
build:
	@ go build 

.PHONY: run
run:
	@ PORT=8080 ./unicorn

.PHONY: build-prod
build-prod:
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w -s' 

.PHONY: docker-build
docker-build:
	@ docker build -t unicorn .

.PHONY: docker-run
docker-run:
	@ docker run unicorn