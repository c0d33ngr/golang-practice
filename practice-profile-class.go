package main

import "fmt"

type Profile struct {
	FirstName string
	LastName string
	Gender string
	DateOfBirth int
	MonthOfBirth int
	YearOfBirth int
	CountryOfOrigin string
}

func main() {
	var first, last, gender, country string
	var date, month, year int

	fmt.Println("Enter your first name:")
	fmt.Scanln(&first)

	fmt.Println("Enter your last name:")
	fmt.Scanln(&last)

	fmt.Println("Enter your gender:")
	fmt.Scanln(&gender)

	fmt.Println("Enter your date of birth:")
	fmt.Scanln(&date)

	fmt.Println("Enter the month you were born using number format. Example, 1 for january:")
	fmt.Scanln(&month)

	fmt.Println("Enter the year you were born:")
	fmt.Scanln(&year)

	fmt.Println("Enter your country of origin:")
	fmt.Scanln(&country)


	person := Profile{
		FirstName: first,
		LastName: last,
		Gender: gender,
		DateOfBirth: date,
		MonthOfBirth: month,
		YearOfBirth: year,
		CountryOfOrigin: country,
	}


	fmt.Println(person.FirstName, "is a", person.Gender)

}
