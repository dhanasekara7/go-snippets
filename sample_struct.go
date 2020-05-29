package main

import "fmt"

// Doctor 'D' upper case, so visible outside
type Doctor struct {
	// number 'n' is lowercase, so not visible outside
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	aDoctorWithOutExplicitNameAssignment := Doctor{
		31,
		"Jon Pertwee",
		[]string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctorWithOutExplicitNameAssignment)
}
