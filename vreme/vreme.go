// Package [vreme] provides a simple structure for keeping track of stations and their temperatures
//
// The package implements a map (Slovar) for named stations where each can store a custom number of
//  temperature measurements
// It includes functions for adding and displaying both, where IzpisPorocila also shows their statistics
// 	and prescribes a group into which the station falls into
//
//
// Example usage:
//
//	vreme.DodajPostajo("ID0", "station name", "location")
//	vreme.DodajMeritev(vreme.Slovar, "ID0", 20.0)
//
//	vreme.DodajPostajo("ID1", "another station name", "location")
//	vreme.DodajMeritev(vreme.Slovar, "ID1", 13.2)
//	vreme.DodajMeritev(vreme.Slovar, "ID1", 19.8)
//	vreme.DodajMeritev(vreme.Slovar, "ID1", 24.4)
//
//	vreme.IzpisMeritev(vreme.Slovar)
//	vreme.IzpisPorocila(vreme.Slovar)
//
//
// The default values for constraints:
// 	- minTemp = -50°C  |  minimum temperature limit,
// 	- maxTemp = 50°C   |  maximum temperature limit,
// 	- stMeritev = 3    |  minimum number of measurements needed for displaying the station (used in IzpisPorocil)

package vreme

import (
	"fmt"
	"slices"
)

// Postaja represents a single station with it's name, location and measurements
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// Slovar contains all made stations and maps them with unique keys (IDs)
var Slovar map[string]Postaja

// default values for constraints
var min_temp int = -50
var max_temp int = 50
var st_meritev int = 3

// Constraints changes the default values
func Constraints(minTemp int, maxTemp int, stMeritev int) {
	min_temp = minTemp
	max_temp = maxTemp
	st_meritev = stMeritev
}

// DodajPostajo creates a new station
func DodajPostajo(oznaka string, ime string, kraj string) {
	Slovar[oznaka] = Postaja{ime, kraj, []float64{}}
}

// DodajMeritev adds a measurement to the specified station if it follows all constraints
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64) {
	if temperatura > float64(max_temp) || temperatura < float64(min_temp) {
		fmt.Printf("error: temperatura %.3f je preveč ekstremna\n", temperatura)
		return
	}
	if vnos, ok := postaje[oznaka]; ok {
		vnos.meritve = append(vnos.meritve, temperatura)
		postaje[oznaka] = vnos
		fmt.Printf("info: temperatura %.2f dodana\n", temperatura)
	} else {
		fmt.Printf("error: postaja %s še ni narejena", oznaka)
	}
}

// IzpisMeritev prints all present stations and their measurements
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Printf("\nmeritve:\n")
	for i := range postaje {
		fmt.Printf("> %s | %s ( %s ) | %v\n", i, postaje[i].ime, postaje[i].kraj, postaje[i].meritve)
	}
}

// IzpisPorocila prints all present stations, their statistiscs and groups
// possible groups: "vroče" = hot, "toplo" = warm, "hladno" = cold, "mrzlo" = freezing
func IzpisPorocila(postaje map[string]Postaja) {
	fmt.Printf("\nporočila:\n")
	for i := range postaje {
		mini, maxi, povprecje, ok := statistika(postaje, i)
		if !ok {
			fmt.Printf("> %s | ni dovolj popularen/a (premalo meritev)\n", postaje[i].ime)
		} else {
			var izpis string
			switch {
			case povprecje >= 25:
				izpis = "vroče"
			case povprecje < 25 && povprecje >= 15:
				izpis = "toplo"
			case povprecje < 15 && povprecje >= 5:
				izpis = "hladno"
			case povprecje < 5:
				izpis = "mrzlo"
			}
			fmt.Printf("> %s | min: %.2f °C | max: %.2f °C | povprečje: %.3f °C | => %s\n", postaje[i].ime, mini, maxi, povprecje, izpis)

		}
	}
}

// statistika is a helper function called by IzpisPorocila
// it calculates and prints the minumum, maximum and avarage values
func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	if vnos, ok := postaje[oznaka]; ok {
		if len(vnos.meritve) < st_meritev {
			fmt.Printf("error: postaja ni dovolj popularna (premalo meritev)\n")
			return -1, -1, -1, false
		}
		var povprecje float64 = 0
		for i := range vnos.meritve {
			povprecje += vnos.meritve[i]
		}
		fmt.Printf("info: statistika za %s ", oznaka)
		return slices.Min(vnos.meritve), slices.Max(vnos.meritve), (povprecje / float64(len(vnos.meritve))), true

	} else {
		fmt.Printf("error: postaja %s še ni narejena\n", oznaka)
	}
	return -1, -1, -1, false
}
