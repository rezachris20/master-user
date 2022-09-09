package share

import "github.com/go-pg/pg/v10"

type Repository struct {
	DBRead, DBWrite *pg.DB
}

func NewRepository(dbRead, dbWrite *pg.DB) *Repository {
	return &Repository{DBRead: dbRead, DBWrite: dbWrite}
}
