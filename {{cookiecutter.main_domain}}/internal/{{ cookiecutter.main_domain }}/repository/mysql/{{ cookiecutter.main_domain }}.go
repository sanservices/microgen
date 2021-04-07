package mysql

import (
	"context"
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}/entity"
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}/repository/errs"
)

// GetThing stands for getting a thing from db
func (m *Mysql) GetThing(ctx context.Context, id uint) (*entity.ThingRec, error) {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errs.ErrNotAbleToStartTransaction
	}
	var thing *entity.ThingRec
	err = tx.Get(&thing, "call prc.GetThing(?)", id)

	return thing, err
}
