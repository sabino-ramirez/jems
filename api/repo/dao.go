package repo

import "database/sql"

type DAO interface {
	UserRepo(db *sql.DB) UserRepo
}

type dao struct{}

func NewDAO() DAO {
	d := &dao{}
	return d
}

func (d *dao) UserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

// func (d *dao) NewPlaylistRepo() PlaylistRepo {
//   return &playlistRepo{}
// }
