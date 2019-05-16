package main

import "fmt"

func main() {
	/*
	   elements := make(map[string]string)
	   elements["H"] = "Hydrogen"
	   elements["He"] = "Helium"
	   elements["Li"] = "Lithium"
	   elements["Be"] = "Beryllium"
	   elements["B"] = "Boron"
	   elements["C"] = "Carbon"
	   elements["N"] = "Nitrogen"
	   elements["O"] = "Oxygen"
	   elements["F"] = "Flurine"
	   elements["Ne"] = "Neon"
	*/

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Flourine",
		"Ne": "Neon",
	}

	fmt.Println(elements["Li"])

	// does not exist
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
	// exists
	if name, ok := elements["N"]; ok {
		fmt.Println(name, ok)
	}

}
