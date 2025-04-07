package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tem-mars/tcgptrade/internal/data"
	"github.com/tem-mars/tcgptrade/internal/validator"
)

func (app *application) createCardHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Setcode string   `json:"setcode"`
		Setname string   `json:"setname"`
		Number  int32    `json:"number"`
		Name    string   `json:"name"`
		Type    string   `json:"type"`
		Rarity  string   `json:"rarity"`
		Weakness string   `json:"weakness"`
		Artist  string   `json:"artist"`
		Packs   []string `json:"packs"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	card := &data.Card{
		Setcode: input.Setcode,
		Setname: input.Setname,
		Number:  input.Number,
		Name:    input.Name,
		Type:    input.Type,
		Rarity:  input.Rarity,
		Weakness: input.Weakness,
		Artist:  input.Artist,
		Packs:   input.Packs,
	}

	v := validator.New()

	if data.ValidateMovie(v, card); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showCardHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	card := data.Card{
		ID:        id,
		CreatedAt: time.Now(),
		Setcode:   "A1",
		Setname:   "Genetic Apex",
		Number:    1,
		Name:      "Bulbasaur",
		Type:      "Grass",
		Rarity:    "Common",
		Weakness:  "Fire",
		Artist:    "Narumi Sato",
		Packs:     []string{"Genetic Apex", "Promo-A"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"card": card}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
