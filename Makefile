DockerfilePROJECT:=kdb-sidecar

MYSQL_SIDECAR_IMAGE_NAME ?= kdbdeveloper/mysql-sidecar:v0.0.11

.PHONY: build
build:
	GOOS=linux go build -ldflags="-w -s" -a -installsuffix "" -o hack/docker/manager cmd/sidecar/main.go

.PHONY: docker
docker: build
	docker build -f hack/docker/Dockerfile -t $(MYSQL_SIDECAR_IMAGE_NAME) .
	#docker push  $(MYSQL_SIDECAR_IMAGE_NAME)


.PHONY: mysql-sidecar-80
mysql-sidecar-80: build
	docker build -f hack/mysql/Dockerfile -t $(MYSQL_SIDECAR_IMAGE_NAME) .
	#docker push  $(MYSQL_SIDECAR_IMAGE_NAME)








