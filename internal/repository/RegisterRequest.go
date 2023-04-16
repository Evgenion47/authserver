package repository

import (
	"authserver/internal/app"
	"authserver/internal/models"
	"context"
)

func (r *repository) RegisterRequest(ctx context.Context, data models.RegisterRequest) (err error) {
	const query1 = `
		
`
	const query2 = `

`
	resp, err := r.pool.Query(ctx, query1, data.Email)
	if resp.Next() {
		err = app.AlreadyExistErr
		return
	}

	_, err = r.pool.Exec(ctx, query2, data.Email, data.Password)

	return
}
