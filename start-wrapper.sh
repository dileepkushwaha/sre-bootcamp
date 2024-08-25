#!/bin/sh
export APP_PORT="${APP_PORT:-8080}"
/root/run-migrations.sh
./sre-bootcamp --port $APP_PORT
/root/start.sh