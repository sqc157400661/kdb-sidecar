#FROM debian:bullseye-slim
FROM silentred/debian:slim-util


ENV MYSQL_MAJOR 8.0
ENV MYSQL_VERSION 8.0.37-1debian11

RUN groupadd -r mysql && useradd -r -g mysql mysql

# create mysql config
RUN mkdir /etc/mysql && \
    echo "!includedir /u01/etc/" > /etc/mysql/my.cnf

# install supervisor, tools are in image silentred/debian:slim-util
 RUN apt-get update && apt-get install -y supervisor wget procps \
    vim dnsutils curl inetutils-ping net-tools

COPY hack/deb/mysql/mysql-community-server-core_"$MYSQL_VERSION"_amd64.deb  /tmp/mysql-server.deb
COPY hack/deb/mysql/mysql-community-client-plugins_"$MYSQL_VERSION"_amd64.deb  /tmp/mysql-client-plugins.deb
COPY hack/deb/mysql/mysql-community-client-core_"$MYSQL_VERSION"_amd64.deb  /tmp/mysql-client-core.deb

RUN apt-get -f install /tmp/mysql-server.deb -y && \
        apt-get -f install /tmp/mysql-client-plugins.deb -y && \
        apt-get -f install /tmp/mysql-client-core.deb -y && \
        rm /tmp/mysql-*.deb


COPY hack/deb/mysql/server_audit.so  /usr/lib/mysql/plugin/server_audit.so


RUN rm -rf /var/lib/apt/lists/*

RUN mkdir -p /u01/socket/ && \
    mkdir -p /var/lib/mysql-files &&\
    chmod 777 -R /u01 &&\
    mkdir -p /u01/etc/ &&\
    mkdir -p /var/lib/mysql &&\
    mkdir -p /u01/app/

# copy config for supervisor
COPY hack/scripts/supervisor_mysql.conf /etc/supervisor/conf.d/supervisor_mysql.conf
COPY hack/scripts/init_run_mysqld.sh /init_run_mysqld.sh
COPY hack/scripts/run_supervisor.sh /run_supervisor.sh
COPY hack/scripts/mysqld_health.sh /mysqld_health.sh
COPY hack/scripts/start_app.sh /start_app.sh
COPY hack/scripts/mysqld.cnf /u01/etc/mysqld.cnf

COPY hack/sql/*  /u01/app/
COPY hack/config/*  /u01/app/
COPY hack/go-admin  /u01/app/go-admin

RUN chmod +x /init_run_mysqld.sh && chmod +x /run_supervisor.sh &&  chmod +x /mysqld_health.sh  &&  chmod +x /start_app.sh

# set timezone
RUN rm /etc/localtime && \
    ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENTRYPOINT ["/bin/bash", "-c", "/run_supervisor.sh"]
