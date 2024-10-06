#!/usr/bin/env bash

docker compose up --detach challenge
docker compose run --rm \
    certbot certonly --webroot \
    --webroot-path /var/www/certbot/ \
    -d peteli.dev \
    || echo "Failed to obtain or renew ssl certs"
docker compose down challenge
