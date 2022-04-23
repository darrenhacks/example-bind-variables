FROM postgres:14.2

# Create the example table and seed it on initialization.
COPY schema.sql /docker-entrypoint-initdb.d/01-schema.sql
COPY data.sql /docker-entrypoint-initdb.d/02-data.sql
