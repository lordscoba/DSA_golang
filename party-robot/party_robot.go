package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {

	var text string
	text += "Welcome to my party, " + name + "!\n"
	text += "You have been assigned to table " + fmt.Sprintf("%03d", table) + ". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here.\n"
	text += "You will be sitting next to " + neighbor + "."

	return text
}
