package repo

type DAO interface {
  NewUserRepo() UserRepo
}

type dao struct{}

func NewDAO() DAO {
  d := &dao{}
  return d
}

func (d *dao) NewUserRepo() UserRepo {
  return &userRepo{}
}

// func (d *dao) NewArtistRepo() ArtistRepo {
//   return &artistRepo{}
// }
