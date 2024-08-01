# go-clean-architecture-template

# Description
This is an example of implementation of Clean Architecture in Golang.

# Install Tools
```shell
go install github.com/go-delve/delve/cmd/dlv@latest
go install github.com/cosmtrek/air@latest
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install github.com/go-task/task/v3/cmd/task@latest
```

# Project Structure
```shell
.
├── README.md
├── Taskfile.yml
├── cmd
│   └── api
│       ├── app
│       │   └── app.go
│       ├── handler
│       │   ├── app_api_handler.go
│       │   ├── auth_api_handler.go
│       │   ├── company_api_handler.go
│       │   └── swagger_api_handler.go
│       └── main.go
├── config
│   ├── config_api.yml
│   ├── local.env
│   └── permissions
│       ├── model.conf
│       └── policy.csv
├── db.sql
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   ├── config.go
│   │   └── loader.go
│   ├── constants
│   │   ├── constants.go
│   │   ├── errorcode
│   │   │   └── error_code.go
│   │   ├── messages
│   │   │   ├── code.go
│   │   │   └── message.go
│   │   └── permissions
│   │       ├── generate.go
│   │       └── permissions.go
│   ├── core
│   │   ├── api
│   │   │   ├── commons
│   │   │   │   └── commons.go
│   │   │   ├── handler.go
│   │   │   ├── middleware
│   │   │   │   ├── apikey_auth.go
│   │   │   │   ├── auth.go
│   │   │   │   ├── jwt_auth.go
│   │   │   │   ├── permissions_auth.go
│   │   │   │   └── request_time.go
│   │   │   ├── models
│   │   │   │   └── api_error_model.go
│   │   │   ├── server.go
│   │   │   ├── swagger
│   │   │   │   └── swagger.go
│   │   │   └── utils.go
│   │   ├── appcontext
│   │   │   └── request_context.go
│   │   ├── auth
│   │   │   └── jwt
│   │   │       └── jwt.go
│   │   ├── config
│   │   ├── constants
│   │   │   ├── constants.go
│   │   │   └── header.go
│   │   ├── errorcode
│   │   │   └── error_code.go
│   │   ├── errors
│   │   │   ├── error_code.go
│   │   │   ├── errors.go
│   │   │   ├── example.go
│   │   │   ├── http_errors.go
│   │   │   ├── messages.go
│   │   │   ├── multierr
│   │   │   │   └── multierr.go
│   │   │   └── types.go
│   │   ├── json
│   │   │   └── json.go
│   │   ├── logger
│   │   │   ├── logger.go
│   │   │   └── zap
│   │   │       ├── test
│   │   │       │   └── main.go
│   │   │       └── zap_logger.go
│   │   ├── models
│   │   │   ├── api_model.go
│   │   │   ├── auth_key_model.go
│   │   │   ├── cert_model.go
│   │   │   ├── client_model.go
│   │   │   ├── data_model.go
│   │   │   ├── date_model.go
│   │   │   ├── file_model.go
│   │   │   ├── message_model.go
│   │   │   ├── page_model.go
│   │   │   ├── resource_model.go
│   │   │   └── ws_model.go
│   │   ├── nullvalue
│   │   │   └── null.go
│   │   ├── template
│   │   │   ├── jet
│   │   │   │   ├── example
│   │   │   │   │   └── main.go
│   │   │   │   └── jet_template.go
│   │   │   └── template.go
│   │   └── utils
│   │       ├── download.go
│   │       ├── env.go
│   │       ├── file.go
│   │       ├── http.go
│   │       ├── id.go
│   │       ├── json.go
│   │       ├── page.go
│   │       ├── password.go
│   │       ├── query.go
│   │       ├── signal.go
│   │       ├── utils.go
│   │       ├── validate.go
│   │       └── zip.go
│   ├── domain
│   │   ├── company.go
│   │   └── repomodels
│   │       └── company_list_repo_model.go
│   ├── modules
│   │   ├── auth
│   │   │   ├── auth_usecase.go
│   │   │   ├── models
│   │   │   │   ├── auth_login_model.go
│   │   │   │   └── auth_refresh_token_model.go
│   │   │   └── usecase
│   │   │       ├── auth_login_usecase.go
│   │   │       ├── auth_refresh_token_usecase.go
│   │   │       └── auth_usecase.go
│   │   └── company
│   │       ├── company_usecase.go
│   │       ├── models
│   │       │   ├── company_common_models.go
│   │       │   ├── company_create_model.go
│   │       │   ├── company_delete_model.go
│   │       │   ├── company_example_db_transaction_model.go
│   │       │   ├── company_find_by_id_model.go
│   │       │   ├── company_list_model.go
│   │       │   └── company_update_model.go
│   │       └── usecase
│   │           ├── company_common_usecase.go
│   │           ├── company_create_usecase.go
│   │           ├── company_delete_usecase.go
│   │           ├── company_example_db_transaction_usecase.go
│   │           ├── company_find_by_id_usecase.go
│   │           ├── company_list_usecase.go
│   │           ├── company_update_usecase.go
│   │           └── company_usecase.go
│   └── repository
│       └── company_repository.go
└── logs
└── app.log
```

# Run
```shell
task run-api
```
