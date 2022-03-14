module {{ cookiecutter.module_name }}

go 1.16

require (
	github.com/Shopify/sarama v1.30.0
	github.com/ThreeDotsLabs/watermill v1.1.1
	github.com/ThreeDotsLabs/watermill-kafka/v2 v2.2.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/labstack/echo-contrib v0.12.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/hamba/avro v1.6.2
	github.com/mattn/go-sqlite3 v1.14.9
	github.com/opentracing/opentracing-go v1.2.0
	github.com/sanservices/apicore v1.1.1
	github.com/sanservices/apilogger/v2 v2.1.0
	github.com/sanservices/kit v1.3.1
	go.uber.org/fx v1.17.0
	gopkg.in/yaml.v2 v2.4.0
)
