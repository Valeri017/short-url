package link

import (
	"short-url/pkg/db"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(datebase *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: datebase,
	}
}

func (repo LinkRepository) Create(link *Link)
