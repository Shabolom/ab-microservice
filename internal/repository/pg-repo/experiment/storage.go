package experiment

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

//	type inMemoryStorage interface {
//		Replace(newCash map[int64]dto.RawExperiment)
//	}
type Storage struct {
	conn *pgxpool.Pool
	//inMemoryStorage inMemoryStorage
}

func New(conn *pgxpool.Pool) *Storage {
	return &Storage{
		conn: conn,
		//inMemoryStorage: inMemoryStorage,
	}
}
