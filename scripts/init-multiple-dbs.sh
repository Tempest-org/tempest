#!/bin/bash
set -e

databases=${POSTGRES_MULTIPLE_DATABASES}

for db in $(echo $databases | tr ',' ' '); do
  echo "Creating database: $db"
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE $db;
EOSQL
done
