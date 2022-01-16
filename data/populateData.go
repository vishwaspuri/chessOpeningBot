package data

import (
	"errors"
	"github.com/anaskhan96/soup"
	"os"
)

//Opening Model
type Opening struct {
	Code  string `json:"code,required"`
	Name  string `json:"name,required"`
	Moves string `json:"opening,required"`
}

//Openings Model for capturing multiple openings
type Openings []Opening

var url = "https://www.chessgames.com/chessecohelp.html"

func GetAllOpenings() Openings {
	var result Openings
	table := getTable()
	for _, tr := range table {
		var opening Opening
		cols := tr.FindAll("td")
		data := cols[1]
		opening.Code = cols[0].Find("font").Text()
		opening.Name = data.Find("b").Text()
		opening.Moves = data.Find("font").Find("font").Text()
		result = append(result, opening)
	}
	return result
}

func GetOpeningByCode(code string) (Opening, error) {
	table := getTable()
	for _, tr := range table {
		cols := tr.FindAll("td")
		data := cols[1]
		if code == cols[0].Find("font").Text() {
			var opening Opening
			opening.Code = cols[0].Find("font").Text()
			opening.Name = data.Find("b").Text()
			opening.Moves = data.Find("font").Find("font").Text()
			return opening, nil
		}
	}
	return Opening{}, errors.New("opening code not found")
}

func getTable() []soup.Root {
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	return doc.Find("table").FindAll("tr")
}
