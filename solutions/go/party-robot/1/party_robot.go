package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %v!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %v! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	if table < 10 {
        return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table 00%v. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.", name, table, direction, distance, neighbor)
    } else if table < 100 {
        return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table 0%v. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.", name, table, direction, distance, neighbor)
    } else {
        return fmt.Sprintf("Welcome to my party, %v!\nYou have been assigned to table %v. Your table is %v, exactly %.1f meters from here.\nYou will be sitting next to %v.", name, table, direction, distance, neighbor)
    }
}
