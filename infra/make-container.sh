#!/bin/bash
set -eox pipefail

export GOARCH=amd64
export GOOS=linux

script_dir=$(dirname "$0")
cd $script_dir

for i in $(cat .env | grep -v '^#'); do export $i; done

pushd .. >/dev/null
make
popd >/dev/null

docker-compose up --build -d

until docker exec za-postgres psql -U $POSTGRES_USER $POSTGRES_DB -c "\l"; do
  echo "Waiting for za-postgres to be ready..."
  sleep 2
done

docker exec -i za-postgres psql -U $POSTGRES_USER $POSTGRES_DB < ../scheme.sql
