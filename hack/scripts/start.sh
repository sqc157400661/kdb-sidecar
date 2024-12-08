#!/bin/bash

# set default values
Engine="MYSQL"
ConfigFile="/etc/config/config.yaml"

# update configuration file path based on environment variables
if [ -n "${CONFIG_FILE_ENV}" ]; then
    ConfigFile="${CONFIG_FILE_ENV}"
fi

# update database engine based on environment variables
if [ -n "${ENGINE_ENV}" ]; then
    Engine="${ENGINE_ENV}"
fi

if [[ "${Engine}" == "MYSQL" || "${Engine}" == "mysql" || "${Engine}" == "MySQL" ]]; then
    /kdb/bin/manager MySQLSidecar -c "${ConfigFile}"
else
    /kdb/bin/manager PGSidecar -c "${ConfigFile}"
fi

sleep 30