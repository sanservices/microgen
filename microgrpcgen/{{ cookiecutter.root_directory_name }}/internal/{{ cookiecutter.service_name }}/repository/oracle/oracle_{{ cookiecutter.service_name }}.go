package oracle

{% if cookiecutter.use_database == 'y' %}
import (
	"context"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/entity"
)

// GetUser stands for getting a thing from db
func (m *Oracle) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	thing := &entity.User{}

	//TODO: Implement mysql calls
	return thing, nil
}

func (m *Oracle) SaveUser(ctx context.Context, thing *entity.User) error {

	//TODO: Implement mysql calls
	return nil
}
{% endif %}
