module {{ cookiecutter.module_name }}

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/labstack/echo/v4 v4.3.0
	github.com/labstack/gommon v0.3.0
	github.com/sanservices/apicore v1.1.1
	github.com/sanservices/apilogger/v2 v2.0.2
	github.com/spf13/viper v1.7.1
	go.uber.org/fx v1.13.1
)
