package main //main is a package and we need to runtime with main package

import (
	"booking-app/helper"
	"fmt"
	"strconv"
) //fmt is external function package

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUser()

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			varFirstNames := getFirstName()
			fmt.Printf("The first names of bookings are: %v\n", varFirstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

			fmt.Println("Your input data is invalid, try again")
		}
	}

	/*
		city := "London"
		switch city {
		case "New York":
			// execute code for booking New York conference tickets
		case "Singapore", "Hong Kong":
		// execute code for booking Singapore and Hong Kong conference tickets
		case "London", "Berlin":
		// execute code for booking London and Berlin conference tickets
		case "Mexico City":
		// execute code for booking Mexico city conference tickets
		default: // For data not match
			fmt.Print("No valid city selected")
		}*/
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total fo %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}           // slice(array in js)
	for _, booking := range bookings { // index for showing index in output ex. : index: 1 booking: Nana
		//var names = strings.Fields(booking) // strings.Fields() = split
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// var bookings = [50]string{"Nana", "Nicole", "Peter"} // limit array at 50 elements
	var firstName string
	// ask user for their name
	var lastName string
	var email string
	var userTickets uint
	// have to define data type when not assign value of variable
	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName) // pointer

	fmt.Print("Enter your last name :")
	fmt.Scan(&lastName) // pointer

	fmt.Print("Enter your first email address :")
	fmt.Scan(&email) // pointer

	fmt.Print("Enter number of tickets :")
	fmt.Scan(&userTickets) // pointer

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//bookings = append(bookings, firstName+" "+lastName) // append similar array.push in JS

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
