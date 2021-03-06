#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

CREATE USER faceapp WITH PASSWORD 'secretPassword';

CREATE DATABASE faceapp;
GRANT ALL PRIVILEGES ON DATABASE faceapp TO faceapp;
EOSQL