DockerfilePROJECT:=kdb-sidecar

MYSQL_SIDECAR_IMAGE_NAME ?= kdbdeveloper/mysql-sidecar:v0.0.1

.PHONY: mysql
mysql:
	GOOS=linux go build -ldflags="-w -s" -a -installsuffix "" -o hack/docker/manager cmd/sidecar/main.go

.PHONY: mysql-docker
mysql-docker: mysql
	docker build -f hack/docker/Dockerfile -t $(MYSQL_SIDECAR_IMAGE_NAME) .
	docker push  $(MYSQL_SIDECAR_IMAGE_NAME)









