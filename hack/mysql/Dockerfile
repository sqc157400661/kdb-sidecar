FROM silentred/debian:slim-util

ARG XBVersion=8.0.35

LABEL maintainers="kdbdeveloper"
LABEL description="kdb-sidecar"


# install supervisor, tools are in image silentred/debian:slim-util
RUN apt-get update && apt-get install -y supervisor wget procps \
    vim dnsutils curl inetutils-ping net-tools cron

# https://www.percona.com/downloads
COPY hack/deb/percona-xtrabackup-*_"$XBVersion"-*.bullseye_amd64.deb  /tmp/percona-xtrabackup.deb

RUN apt-get -f install /tmp/percona-xtrabackup.deb -y && \
        rm /tmp/*.deb

# set timezone
RUN rm /etc/localtime && \
    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime


RUN mkdir -p /kdb/bin/ && mkdir -p /kdb/conf/

COPY hack/docker/manager  /kdb/bin/
COPY hack/scripts/*  /kdb/bin/

WORKDIR /home/admin/

RUN  chmod +x /kdb/bin/*