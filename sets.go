package mtgjson

import (
	"fmt"
	"sort"

	"github.com/reanisz/mtgjson-go/models"
)

func (mtgjson *Database) FindSetsByCodes(sets []string) (map[string]*models.Set, error) {
	query := fmt.Sprintf(`SELECT *
	  FROM sets
	  WHERE sets.code in ( %v )
	`, generate_placeholders(len(sets)))

	var args []interface{}
	for i := range sets {
		args = append(args, &sets[i])
	}

	rows, err := mtgjson.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	res := make(map[string]*models.Set)

	for rows.Next() {
		set := new(models.Set)
		rows.StructScan(set)

		res[set.Code] = set

		mtgjson.setsCache[set.Code] = set
	}

	return res, nil
}

func (mtgjson *Database) PrepareSetCache(sets []string) error {
	var codes []string

	for _, code := range sets {
		if _, ok := mtgjson.setsCache[code]; !ok {
			codes = append(codes, code)
		}
	}

	if len(codes) == 0 {
		return nil
	}

	_, err := mtgjson.FindSetsByCodes(codes)

	return err
}

func (mtgjson *Database) GetSetByCode(code string) (*models.Set, error) {
	if v, ok := mtgjson.setsCache[code]; ok {
		return v, nil
	}

	sets, err := mtgjson.FindSetsByCodes([]string{code})

	if err != nil {
		return nil, err
	}

	if 0 == len(sets) {
		return nil, fmt.Errorf("set not found. (code=%v)", code)
	}

	for _, v := range sets {
		return v, nil
	}

	return nil, fmt.Errorf("InvalidOperation")
}

func (mtgjson *Database) IsPremierSet(set *models.Set) bool {
	if set.Type.String == "promo" {
		return true
	}

	return false
}

func (mtgjson *Database) CompareAsBestSet(lhs, rhs *models.Set) (bool, bool) {
	if eq, cmp := compInt(lhs.Isonlineonly, rhs.Isonlineonly); !eq {
		return false, !cmp
	}
	if eq, cmp := compBool(mtgjson.IsPremierSet(lhs), mtgjson.IsPremierSet(rhs)); !eq {
		return false, !cmp
	}
	if eq, cmp := compInt(lhs.Isforeignonly, rhs.Isforeignonly); !eq {
		return false, !cmp
	}
	if eq, cmp := compTime(lhs.Releasedate.Time, rhs.Releasedate.Time); !eq {
		return false, cmp
	}

	return true, false
}

func (mtgjson *Database) GetBestSet(sets []*models.Set) *models.Set {
	sort.Slice(sets, func(i, j int) bool {
		_, cmp := mtgjson.CompareAsBestSet(sets[i], sets[j])
		return cmp
	})

	return sets[len(sets)-1]
}
