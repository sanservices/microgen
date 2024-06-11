package sqlite

{% if cookiecutter.use_database == 'y' %}
import (
	"context"

	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}/entity"
)

func (s *Sqlite) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	thing := &entity.User{}
	err := s.db.GetContext(ctx, thing, "SELECT * FROM THINGS WHERE id = ?", id)

	return thing, err
}

func (s *Sqlite) SaveUser(ctx context.Context, thing *entity.User) error {
	const query string = "INSERT INTO THINGS (id, name) VALUES (?,?)"
	_, err := s.db.ExecContext(ctx, query, thing.ID, thing.FirstName)

	return err
}
{% endif %}