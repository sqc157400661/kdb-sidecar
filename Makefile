DockerfilePROJECT:=go-admin

PAAS_IMAGE_NAME ?=sqc157400661/mysql:v0.0.16

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -a -installsuffix "" -o go-admin .

.PHONY: build-app-linux
build-app-linux:
	GOOS=linux go build -ldflags="-w -s" -a -installsuffix "" -o hack/go-admin  .

.PHONY: build-docker
build-docker:build-app-linux
	docker build -f hack/docker/Dockerfile.tadmin.local -t $(PAAS_IMAGE_NAME) .
	docker tag $(PAAS_IMAGE_NAME) registry.cn-hangzhou.aliyuncs.com/sqcimg/image:latest
	docker push registry.cn-hangzhou.aliyuncs.com/sqcimg/image:latest

.PHONY: build-mysql-docker
build-mysql-docker:build-app-linux
	docker build -f hack/docker/Dockerfile.mysqld.local -t $(PAAS_IMAGE_NAME) .

.PHONY: buildsqc
buildsqc:
	go build -ldflags="-w -s" -a -installsuffix "" -o go-admin .

.PHONY: runsqc
runsqc:buildsqc
	./go-admin server -c=config/settings.sqc.local.yml >> access.log #2>&1 &

.PHONY: run-dev
run-dev:buildsqc
	./go-admin server -c=hack/config/settings.local.mac.yml








