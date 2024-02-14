```
hexagonal/
│
├── adapters/
│   ├── http_handler.go
│   ├── database_repository.go
│   └── email_sender.go (Example)
│
├── core
│   ├── employee
│   │   ├── employee_entities.go
│   │   ├── employee_repository.go
│   │   └── employee_usecase.go
│   │ 
│   └── order (Example)
│       ├── order_entities.go (Example)
│       ├── order_repository.go (Example)
│       ├── order_usecase.go (Example)
│       ├── notification.go (Example)
│       └── notification_service.go (Example)
│
├── cmd/
│   └── main.go
│
└── go.mod
```