FROM mysql:8
COPY ./script/ /docker-entrypoint-initdb.d/
EXPOSE 3306