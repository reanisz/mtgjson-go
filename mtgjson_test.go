package mtgjson

import (
	"testing"
)

func SetupDatabase(t *testing.T) (*Database, func()) {
	db, err := Open("sqlite3", "./data/AllPrintings.sqlite")
	if err != nil {
		t.Fatal(err)
	}

	return db, func() { db.Close() }
}

func TestFindCardByName(t *testing.T) {
	db, teardown := SetupDatabase(t)
	defer teardown()

	type args struct {
		name  string
		found bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Opt",
			args: args{name: "Opt", found: true},
		},
		{
			name: "Opttt",
			args: args{name: "Optttt", found: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards, err := db.FindCardsByName(tt.args.name)
			found := 0 < len(cards)
			if err != nil {
				t.Fatal(err)
			}
			if found != tt.args.found {
				t.Fatalf("invalid found by \"%v\", expected %v", tt.args.name, tt.args.found)
			}
		})
	}

}

func TestFindCardByForiegnName(t *testing.T) {
	db, teardown := SetupDatabase(t)
	defer teardown()

	type args struct {
		name  string
		found bool
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "選択",
			args: args{name: "選択", found: true},
		},
		{
			name: "洗濯機",
			args: args{name: "洗濯機", found: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards, err := db.FindCardsByForeignName(tt.args.name)
			found := 0 < len(cards)
			if err != nil {
				t.Fatal(err)
			}
			if found != tt.args.found {
				t.Fatalf("invalid found by \"%v\", expected %v", tt.args.name, tt.args.found)
			}
			t.Logf("%v :%v", tt.args.name, len(cards))
			if 0 < len(cards) {
				t.Logf("%v", cards[0])
			}
		})
	}

}
