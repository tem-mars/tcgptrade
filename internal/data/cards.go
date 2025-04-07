package data

import (
	"time"

	"github.com/tem-mars/tcgptrade/internal/validator"
)

type Card struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Setcode   string    `json:"setcode"`
	Setname   string    `json:"setname"`
	Number    int32     `json:"number,string"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Rarity    string    `json:"rarity"`
	Weakness  string    `json:"weakness"`
	Artist    string    `json:"artist"`
	Packs     []string  `json:"packs"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, card *Card) {
	v.Check(card.Setname != "", "setname", "must be provided")
	v.Check(len(card.Setname) <= 500, "setname", "must not be more than 500 bytes long")

	v.Check(card.Name != "", "Name", "must be provided")
	v.Check(len(card.Name) <= 500, "Name", "must not be more than 500 bytes long")

	v.Check(card.Rarity != "", "rarity", "must be provided")
	v.Check(len(card.Rarity) <= 500, "rarity", "must not be more than 500 bytes long")

	v.Check(card.Packs != nil, "packs", "must be provided")
	v.Check(len(card.Packs) >= 1, "packs", "must contain at least 1 genre")
	v.Check(len(card.Packs) <= 5, "packs", "must not contain more than 5 genres")

	v.Check(validator.Unique(card.Packs), "packs", "must not contain duplicate values")

}
