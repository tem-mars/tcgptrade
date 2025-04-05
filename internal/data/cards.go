package data

import (
	"time"
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
	//Genres    []string
	//Year      int32
	//Title     string
	//Runtime     int32
}
