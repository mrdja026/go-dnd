package dnddata

import (
	"encoding/json"
	"io"
	"net/http"
)

// CHAT GPT GENERATED CODE:
type Class struct {
	Index                    string              `json:"index"`
	Name                     string              `json:"name"`
	HitDie                   int                 `json:"hit_die"`
	ProficiencyChoices       []ProficiencyChoice `json:"proficiency_choices"`
	Proficiencies            []Proficiency       `json:"proficiencies"`
	SavingThrows             []SavingThrow       `json:"saving_throws"`
	StartingEquipment        []StartingEquipment `json:"starting_equipment"`
	StartingEquipmentOptions []EquipmentOption   `json:"starting_equipment_options"`
	ClassLevels              string              `json:"class_levels"`
	MultiClassing            MultiClassing       `json:"multi_classing"`
	Subclasses               []Subclass          `json:"subclasses"`
	URL                      string              `json:"url"`
}

type ProficiencyChoice struct {
	Desc   string       `json:"desc"`
	Choose int          `json:"choose"`
	Type   string       `json:"type"`
	From   ChoiceSource `json:"from"`
}

type ChoiceSource struct {
	OptionSetType string         `json:"option_set_type"`
	Options       []ChoiceOption `json:"options"`
}

type ChoiceOption struct {
	OptionType string         `json:"option_type"`
	Item       *ItemReference `json:"item,omitempty"`
	Choice     *NestedChoice  `json:"choice,omitempty"`
	Count      *int           `json:"count,omitempty"`
	Of         *ItemReference `json:"of,omitempty"`
}

type ItemReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type NestedChoice struct {
	Desc   string       `json:"desc"`
	Type   string       `json:"type"`
	Choose int          `json:"choose"`
	From   ChoiceSource `json:"from"`
}

type Proficiency struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type SavingThrow struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type StartingEquipment struct {
	Equipment ItemReference `json:"equipment"`
	Quantity  int           `json:"quantity"`
}

type EquipmentOption struct {
	Desc   string       `json:"desc"`
	Choose int          `json:"choose"`
	Type   string       `json:"type"`
	From   ChoiceSource `json:"from"`
}

type MultiClassing struct {
	Prerequisites []Prerequisite `json:"prerequisites"`
	Proficiencies []Proficiency  `json:"proficiencies"`
}

type Prerequisite struct {
	AbilityScore AbilityScore `json:"ability_score"`
	MinimumScore int          `json:"minimum_score"`
}

type AbilityScore struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Subclass struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type GeneralInfo struct { // for classes and races same struct
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type GeneralInfoResult struct {
	Count   int           `json:"count"`
	Results []GeneralInfo `json:"results"`
}

func GetClasses() GeneralInfoResult {
	response, error := http.Get("https://www.dnd5eapi.co/api/classes/")
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	responseData, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	var responseObject GeneralInfoResult
	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		panic(err)
	}

	return responseObject
}

func GetRaces() GeneralInfoResult {
	response, error := http.Get("https://www.dnd5eapi.co/api/races/")
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	responseData, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	var responseObject GeneralInfoResult
	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		panic(err)
	}

	return responseObject
}
