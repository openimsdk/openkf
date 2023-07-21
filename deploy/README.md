# OpenKF Deploy

Using the OpenKF deploy tools, you can deploy OpenKF to docker clusters.

### 1. Deploy with single image

1. Build and run web image

```bash
# pwd = /deploy/docker
docker build -f Dockerfile.web -t openkf-web ../../.

docker run -d -p 8080:80 openkf-web
```

### 2. Deploy with docker-compose