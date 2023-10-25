run:
  REDIS_CONNECTION_STRING="redis://localhost:6379/10" go run server.go

spin-redis:
    docker run -d -p 6379:6379 redis

spin:
    docker compose up --build