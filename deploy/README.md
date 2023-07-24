# OpenKF Deploy

Using the OpenKF deploy tools, you can deploy OpenKF to docker clusters.

## Build images

### Build images for web

Build and run web image

```bash
# pwd = /deploy/docker
docker build -f Dockerfile.web -t openkf-web ../../.

# optional: run web image
docker run -d -p 8080:80 openkf-web
```

### Build images for server

```bash
# build openkf-server
docker build -f docker/Dockerfile.server -t openkf-server:latest ../.
```

You can check the images by running the following commands:
```bash
$ docker images

REPOSITORY                                            TAG                 IMAGE ID       CREATED              SIZE
openkf-server                                         latest              f1676becc5ce   About a minute ago   60.5MB
...
```
## Run images

### Use docker-compose to run backend without OpenIMServer

```bash
docker-compose -f docker-compose/docker-compose-dev.yaml up -d
```

You can check the containers by running the following commands:
```bash
$ docker-compose -f docker-compose/docker-compose-dev.yaml ps

NAME                IMAGE                  COMMAND                  SERVICE             CREATED             STATUS                         PORTS
openkf-minio        minio/minio            "/usr/bin/docker-ent…"   minio               8 minutes ago       Up 27 seconds                  0.0.0.0:9100->9000/tcp, :::9100->9000/tcp, 0.0.0.0:9190->9090/tcp, :::9190->9090/tcp
openkf-mysql        mysql:latest           "docker-entrypoint.s…"   mysql               47 minutes ago      Restarting (1) 7 seconds ago   
openkf-redis        redis:latest           "docker-entrypoint.s…"   redis               47 minutes ago      Up 29 seconds                  0.0.0.0:6479->6379/tcp, :::6479->6379/tcp
openkf-server       openkf-server:latest   "./docker_init_serve…"   server              30 seconds ago      Up 5 seconds                   0.0.0.0:10010->10010/tcp, :::10010->10010/tcp                                                                             3.2s
```

Check the logs of the server container:
```bash
$ docker-compose -f docker-compose/docker-compose-dev.yaml logs
```
