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

func (repo LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
