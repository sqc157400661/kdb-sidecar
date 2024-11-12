FROM registry.access.redhat.com/ubi8-minimal

LABEL maintainers="kdbdeveloper"
LABEL description="kdb-sidecar"

RUN microdnf install -y vim iputils

RUN mkdir -p /home/admin/

COPY hack/docker/manager  /home/admin/

WORKDIR /home/admin/

RUN  chmod +x /home/admin/manager