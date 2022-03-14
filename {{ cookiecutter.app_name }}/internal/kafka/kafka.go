package kafka

import (
	"context"

	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}"
	"github.com/sanservices/kit/kafkalistener"
)

type Kafka struct {
	service {{ cookiecutter.main_domain }}.Service
	mb      *kafkalistener.MessageBroker
}

func New(service {{ cookiecutter.main_domain }}.Service, mb *kafkalistener.MessageBroker) *Kafka {
	return &Kafka{
		service: service,
		mb:      mb,
	}
}

func (k *Kafka) StartListener(ctx context.Context) {
	k.mb.Listen(ctx, k.handlers())
}

func (k *Kafka) StopListener() error {
	return k.mb.Stop()
}
