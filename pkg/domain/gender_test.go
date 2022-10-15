package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenderString(t *testing.T) {
	male := Male
	female := Female
	assert.Equal(t, "MALE", male.String())
	assert.Equal(t, "FEMALE", female.String())
}

func TestGenderFromString(t *testing.T) {
	male, _ := FromString("MALE")
	female, _ := FromString("FEMALE")

	assert.Equal(t, Male, male)
	assert.Equal(t, Female, female)
}
