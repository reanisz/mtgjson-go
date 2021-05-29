package mtgjson

import (
	"fmt"
	"sort"
	"strings"

	"github.com/reanisz/mtgjson-go/models"
)

func (mtgjson *Database) FindCardsByNames(cardNames []string) (map[string][]*models.Card, error) {
	query := fmt.Sprintf(`SELECT *
	  FROM cards
	  WHERE cards.name in ( %v )
	  ;
	`, generate_placeholders(len(cardNames)))

	var args []interface{}
	for i := range cardNames {
		args = append(args, &cardNames[i])
	}

	rows, err := mtgjson.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	res := make(map[string][]*models.Card)

	for rows.Next() {
		card := new(models.Card)
		rows.StructScan(card)

		res[card.Name.String] = append(res[card.Name.String], card)
	}

	return res, nil
}

func (mtgjson *Database) FindCardsByName(cardName string) ([]*models.Card, error) {
	found, err := mtgjson.FindCardsByNames([]string{cardName})

	if err != nil {
		return nil, err
	}

	ret, ok := found[cardName]

	if !ok {
		ret = []*models.Card{}
	}

	return ret, nil
}

func generate_placeholders(num int) string {
	res := ""

	for i := 0; i < num; i++ {
		if 0 < i {
			res += ", "
		}
		res += "?"
	}

	return res
}

func (mtgjson *Database) FindCardsByUUIDs(uuids []string) (map[string]*models.Card, error) {
	query := fmt.Sprintf(`SELECT *
	  FROM cards
	  WHERE cards.uuid in ( %v )
	  ;
	`, generate_placeholders(len(uuids)))

	var args []interface{}
	for i := range uuids {
		args = append(args, &uuids[i])
	}

	rows, err := mtgjson.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	res := make(map[string]*models.Card)

	for rows.Next() {
		card := new(models.Card)
		rows.StructScan(card)

		res[card.UUID] = card
	}

	return res, nil
}

type CardWithForeignData struct {
	Card        *models.Card         `db:"cards"`
	ForeignData *models.ForeignDatum `db:"foreign_data"`
}

type CardWithAllFaceData struct {
	Mainface  *CardWithForeignData
	OtherFace []*CardWithForeignData
}

func (mtgjson *Database) FindCardsByForeignNames(cardNames []string) (map[string][]*CardWithForeignData, error) {
	res := make(map[string][]*CardWithForeignData)

	query := fmt.Sprintf(`SELECT *
	   FROM foreign_data
       WHERE foreign_data.name in ( %v )
       `, generate_placeholders(len(cardNames)))

	var args []interface{}
	for i := range cardNames {
		args = append(args, &cardNames[i])
	}

	rows, err := mtgjson.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	var uuids []string
	var founded_foriegn []*models.ForeignDatum

	for rows.Next() {
		data := new(models.ForeignDatum)
		rows.StructScan(data)

		founded_foriegn = append(founded_foriegn, data)
		uuids = append(uuids, data.UUID.String)
	}

	cards, err := mtgjson.FindCardsByUUIDs(uuids)
	if err != nil {
		return nil, err
	}

	for _, foreign := range founded_foriegn {
		r := new(CardWithForeignData)
		r.ForeignData = foreign
		r.Card = cards[foreign.UUID.String]

		res[r.ForeignData.Name.String] = append(res[r.ForeignData.Name.String], r)
	}

	return res, nil
}
func (mtgjson *Database) FindCardsByForeignName(cardName string) ([]*CardWithForeignData, error) {
	found, err := mtgjson.FindCardsByForeignNames([]string{cardName})

	if err != nil {
		return nil, err
	}

	ret, ok := found[cardName]

	if !ok {
		ret = []*CardWithForeignData{}
	}

	return ret, nil
}

func (mtgjson *Database) HasPaperCard(card *models.Card) bool {
	return strings.Contains(card.Availability.String, "paper")
}

func (mtgjson *Database) CompareAsBestCard(lhs, rhs *models.Card) (bool, bool) {
	if eq, cmp := compBool(mtgjson.HasPaperCard(lhs), mtgjson.HasPaperCard(rhs)); !eq {
		return false, cmp
	}
	if eq, cmp := compStr(lhs.Side.String, rhs.Side.String); !eq {
		return false, !cmp
	}

	return true, false
}

// 一番良さそうなカードを選ぶ
func (mtgjson *Database) SelectBestMatchCardWithForeign(cards []*CardWithForeignData) *CardWithAllFaceData {
	if len(cards) == 0 {
		return nil
	}

	sort.Slice(cards, func(i, j int) bool {
		lhs := cards[i]
		rhs := cards[j]

		lhsSet, err := mtgjson.GetSetByCode(lhs.Card.Setcode.String)
		if err != nil {
			panic(err)
		}
		rhsSet, err := mtgjson.GetSetByCode(rhs.Card.Setcode.String)
		if err != nil {
			panic(err)
		}

		if eq, cmp := compBool(lhs.ForeignData == nil, rhs.ForeignData == nil); !eq {
			return cmp
		}
		if lhs.ForeignData != nil {
			if eq, cmp := compBool(len(lhs.ForeignData.Text.String) == 0, len(lhs.ForeignData.Text.String) == 0); !eq {
				return !cmp
			}
		}
		if eq, cmp := mtgjson.CompareAsBestSet(lhsSet, rhsSet); !eq {
			return cmp
		}
		if eq, cmp := mtgjson.CompareAsBestCard(lhs.Card, rhs.Card); !eq {
			return cmp
		}
		return false
	})

	best := cards[len(cards)-1]

	res := new(CardWithAllFaceData)
	res.Mainface = best

	if best.Card.Otherfaceids.String != "" {
		for _, uuid := range strings.Split(best.Card.Otherfaceids.String, ",") {
			for _, card := range cards {
				if card.Card.UUID == uuid {
					res.OtherFace = append(res.OtherFace, card)
					break
				}
			}
		}
	}

	return res
}

func (mtgjson *Database) SelectBestMatchCard(cards []*models.Card) *CardWithAllFaceData {
	if len(cards) == 0 {
		return nil
	}

	var list []*CardWithForeignData

	for _, v := range cards {
		cf := new(CardWithForeignData)
		cf.Card = v
		cf.ForeignData = nil
		list = append(list, cf)
	}

	res := mtgjson.SelectBestMatchCardWithForeign(list)

	return res
}

func (mtgjson *Database) FindForeignCardNameByEnglishName(english_name string, language string) (string, error) {
	rows, err := mtgjson.db.Query(`SELECT DISTINCT foreign_data.name
      FROM cards
        JOIN foreign_data on cards.uuid = foreign_data.uuid
      WHERE
        cards.name = ?
        AND foreign_data.language = ?`, english_name, language)

	if err != nil {
		return "", err
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)

		return name, nil
	}

	return "", fmt.Errorf("Foreign card name is not found, name=%v, language=%v", english_name, language)
}
