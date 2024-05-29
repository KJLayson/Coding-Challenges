package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	mod := int(math.Floor((float64(score) - 10.0) / 2.0))

	return mod
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	lowest := 100
	sum := 0

	for i := 0; i < 4; i++ {
		roll := rand.Intn(6) + 1
		sum += roll
		if roll < lowest {
			lowest = roll
		}
	}

	return sum - lowest

}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	new_char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    0,
	}

	new_char.Hitpoints = 10 + Modifier(new_char.Constitution)

	return new_char
}
