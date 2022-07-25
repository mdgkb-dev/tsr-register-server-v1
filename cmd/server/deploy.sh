#!/bin/bash

DEPLOY_BRANCH="master"
BIN_PATH=./cmd/server
PROCESS_NAME=$1
PIDFILE=${BIN_PATH}/${PROCESS_NAME}.pid
PROCESS_FILE=${BIN_PATH}/${PROCESS_NAME}

#git reset --hard && \
#git pull --all && \
#git checkout $branch && \

echo $PIDFILE
if [ -f "$PIDFILE" ]; then
    echo "$FILE exists."
    echo `cat ${PIDFILE}`
    kill -9 `cat ${PIDFILE}` && rm -f "${PIDFILE}"
fi
go build -o $PROCESS_FILE ./cmd/server/main.go

nohup $PROCESS_FILE & exit
