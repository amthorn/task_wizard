#!/bin/sh

if [ "${ENVIRONMENT}" = "dev" ]; then
    echo "Executing go run ."
    go run /service/src/services/project_service/src/
else
    echo "Executing /main"
    /main
fi