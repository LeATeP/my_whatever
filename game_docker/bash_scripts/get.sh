#!/bin/env bash

psql_conn () {
    psql -h "$PSQL_HOST" -U "$PSQL_USER" -d "$PSQL_DB" -c "$1" 
}

