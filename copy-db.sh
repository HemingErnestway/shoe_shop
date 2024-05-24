#!/bin/bash

# Set the database connection details
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgresql
DB_NAME=shoe_shop_db

# Set the Docker container name
CONTAINER_NAME=shoeshop_db_1

# Dump the existing database
pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER -W $DB_PASSWORD $DB_NAME > db.dump

# Copy the dump file into the Docker container
docker cp db.dump $CONTAINER_NAME:/tmp/db.dump

# Restore the database in the Docker container
docker exec -it $CONTAINER_NAME psql -U $DB_USER $DB_NAME < /tmp/db.dump

# Remove the temporary dump file
rm db.dump