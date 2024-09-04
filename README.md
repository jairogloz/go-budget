## How to run with Docker

1. Build your Docker image with:

```bash
docker build -t go-budget .
```

2. Run your container with:

```bash
docker run -d --name go-budget-container --env-file .env -p 8080:8080 go-budget
```