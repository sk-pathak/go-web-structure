### Basic Structure
```shell
/project-root
├── cmd/
│   └── server/
│       └── main.go
├── configs/
│   └── config.go
├── internal/
│   ├── app/
│   │   ├── handler/
│   │   │   ├── user_handler.go
│   │   │   └── post_handler.go
│   │   ├── service/
│   │   │   ├── user_service.go
│   │   │   └── post_service.go
│   │   ├── model/
│   │   │   ├── user.go
│   │   │   └── post.go
│   │   └── repository/
│   │       ├── user_repository.go
│   │       └── post_repository.go
│   ├── db/
│   │   ├── db.go
│   │   └── migrations/
│   │       └── <migrations_files>
│   ├── middleware/
│   │   ├── auth.go
│   │   └── logging.go
│   ├── transport/
│   │   └── http.go
├── sqlc/
│   ├── queries/
│   │   └── query.sql
│   └── sqlc.yaml
├── scripts/
│   └── migrate.sh
├── go.mod
├── go.sum
└── README.md
```
### Basic flow
routes->handler->service->repo

### cURL as HTTP client:

- curl https://jsonplaceholder.typicode.com/posts -> GET
- curl -X POST -H "Content-Type: application/json" -d '{"title": "foo", "body": "bar", "userId": 1}' https://jsonplaceholder.typicode.com/posts
- curl -u username:password https://example.com/api -> Basic Auth
- curl -F "file=@path/to/file" https://example.com/upload
- curl -o filename https://example.com/file -> Download file
- curl -L https://example.com -> Follow redirects
- curl -b "name=value" https://example.com -> Send cookies
- curl -c cookies.txt https://example.com -> Save cookies

### docker
- local postgres -> ```docker build -t go-structure``` then ```docker run -it --name go-structure --network host go-structure:latest``` (network host so that it connects to localhost of host machine)

- postgres in docker -> ```docker compose up --build```

### concurrency
concurrency is managing multiple tasks at the same time. Parallelism is executing multiple tasks at the same time.