# Personal v1

Homepage containing links and short bio.

```bash
Available make commands:

templ       run hot-reloading templ dev server
tailwind    run hot-reloading tailwind css server
clean       delete generated files
build       build all project files for deployment
```

How to build and deploy new code changes:

```bash
make build

# set this to vm ip
export LINODE=
scp \
    docker-compose.yaml \
    personal-v1.tar \
    nginx.conf \
    nginx-challenge.conf \
    fetch-ssl-certs.sh \
    root@${LINODE}:/root/

ssh $LINODE
cat personal-v1.tar | docker load
```

Obtain or renew ssl certs:

```bash
./fetch-ssl-certs.sh
```

Commands for operation and monitoring:

```bash
docker compose up --detach nginx personal
docker compose down
docker compose ls
docker compose logs -f nginx
docker compose logs -f personal
```
