#!/bin/bash

DEPLOY_BRANCH="master"
PROCESS_NAME=$1
PROCESS_NAME=./cmd/server
PIDFILE=${BIN_PATH}/${PROCESS_NAME}.pid
PROCESS_FILE=${BIN_PATH}/${PROCESS_NAME}

git reset --hard && \
git pull --all && \
git checkout $branch && \

go build -o $PROCESS_FILE ./cmd/server/main.go && \
if [ -f "$PIDFILE" ]; then
    echo "$FILE exists."
    kill -9 `cat ${PIDFILE}` && rm -f "${PIDFILE}"
fi
echo $PROCESS_FILE

nohup $PROCESS_FILE &
exit
