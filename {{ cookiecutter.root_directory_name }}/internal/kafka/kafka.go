package kafka

import(
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"
	"github.com/sanservices/kit/kafkalistener"
)

type Kafka struct {
	service {{ cookiecutter.service_name }}.Service
	mb      *kafkalistener.MessageBroker
}

func New(service {{ cookiecutter.service_name }}.Service, mb *kafkalistener.MessageBroker) *Kafka {
	return &Kafka{
		service: service,
		mb:      mb,
	}
}
