package mtgjson

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/reanisz/mtgjson-go/models"
)

type Database struct {
	db *sqlx.DB

	setsCache map[string]*models.Set
}

func Open(driverName string, dataSourceName string) (*Database, error) {
	db := new(Database)

	sql, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.db = sql
	db.setsCache = make(map[string]*models.Set)

	return db, nil
}

func (mtgjson *Database) Close() error {
	return mtgjson.db.Close()
}

func compBool(lhs bool, rhs bool) (bool, bool) {
	if lhs == rhs {
		return true, false
	}

	return false, !lhs && rhs
}
func compInt(lhs int, rhs int) (bool, bool) {
	if lhs == rhs {
		return true, false
	}

	return false, lhs < rhs
}
func compTime(lhs time.Time, rhs time.Time) (bool, bool) {
	if lhs == rhs {
		return true, false
	}

	return false, lhs.Before(rhs)
}

func compStr(lhs string, rhs string) (bool, bool) {
	if lhs == rhs {
		return true, false
	}

	return false, lhs < rhs
}
