package domain

import (
	"encoding/json"
	"fmt"
)

type Gender uint8

const maleString = "MALE"
const femaleString = "FEMALE"

const (
	Male Gender = iota + 1
	Female
)

func (g Gender) String() string {
	switch g {
	case Male:
		return maleString
	case Female:
		return femaleString
	}
	panic("Unsupported gender type")
}

func FromString(s string) (Gender, error) {
	switch s {
	case maleString:
		return Male, nil
	case femaleString:
		return Female, nil
	}

	return Gender(0), fmt.Errorf("Unsupported string value")
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

func (g *Gender) UnmarshalJSON(data []byte) (err error) {
	var gender string
	if err := json.Unmarshal(data, &gender); err != nil {
		return err
	}
	if *g, err = FromString(gender); err != nil {
		return err
	}
	return nil
}
