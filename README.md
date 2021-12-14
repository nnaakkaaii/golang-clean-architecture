# Clean Architecture by Golang

```
.
├── README.md
├── app
│   ├── adapters
│   │   ├── controller
│   │   │   ├── controller.go
│   │   │   ├── ds_controller.go
│   │   │   └── errors.go
│   │   ├── presenter
│   │   │   ├── ds_view.go
│   │   │   ├── i_view.go
│   │   │   └── presenter.go
│   │   └── repository
│   │       ├── ds_database.go
│   │       ├── i_database.go
│   │       └── repository.go
│   ├── domains
│   │   └── entity
│   │       └── entity.go
│   ├── interfaces
│   │   ├── database
│   │   │   └── database.go
│   │   ├── server
│   │   │   └── server.go
│   │   └── view
│   │       └── view.go
│   └── usecases
│       └── interactor
│           ├── ds_presenter.go
│           ├── i_presenter.go
│           ├── i_repository.go
│           └── interactor.go
├── config
│   └── config.yaml
├── fmt.sh
├── go.mod
├── go.sum
└── main.go
```