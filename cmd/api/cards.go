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
		Setname string   `json:"setname"`
		Name    string   `json:"name"`
		Rarity  string   `json:"rarity"`
		Packs   []string `json:"packs"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	card := &data.Card{
		Setname: input.Setname,
		Name:    input.Name,
		Rarity:  input.Rarity,
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
