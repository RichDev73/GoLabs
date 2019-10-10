package main

import (
	"fmt"
)

/*
    ////////////////////////////////////////
   ////////////////////////////////////////
  //// Entrainement sur les variables ////
 ////////////////////////////////////////
////////////////////////////////////////


*/

func main() {

	// Première facon de noter les variable
	// VARIABLES NUMERIQUES
	var nombreEntier int = 100          // Nombre entier +/-
	var nombreRune rune = 90            // Equivalent au int
	var nombreFlottant float32 = 100.25 // Nombre à virdule +/-
	var nombreNonSigne uint = 90        // Nombre toujours positif

	// VARIABLE TYPE CHAINE
	var monPrenom string = "Richard" // Variable caractère 1 type pour tout
	var monNom string = "Lalieu"

	// VARIABLE TYPE BOOLEEN V/F
	var myTrue bool = true
	var myFalse bool = false

	const pi float32 = 3.14
	const (
		a = iota // iota s'incremente à chaque appel
		b
		c = "Pouet"
		d
	)

	/* 2ème facon de noter les variables
	   const (
		   a = 1
		   b
		   c
		   d
	   )
	   var (
	   	// VARIABLES NUMERIQUES
	   	nombreEntier int = 100          // Nombre entier +/-
	   	nombreRune rune = 90			// Equivalent au int
	   	nombreFlottant float32 = 100.25 // Nombre à virdule +/-
	   	nombreNonSigne uint = 90        // Nombre toujours positif

	   	// VARIABLE TYPE CHAINE
	   	monPrenom string = "Richard" 	// Variable caractère 1 type pour tout
	   	monNom string = "Lalieu"

	   	// VARIABLE TYPE BOOLEEN V/F
	   	myTrue bool = 1					//VRAI
	   	myFalse bool = 0				//FAUX

	   )
	*/

	// TRAITEMENT
	fmt.Println("")
	fmt.Println("Variable numérique :")
	fmt.Println("")
	fmt.Println(nombreEntier, ": Nombre Entier")
	fmt.Println(nombreRune, ": Nombre Entier")
	fmt.Println(nombreFlottant, ": Nombre Flottant sur 32bit")
	fmt.Println(nombreNonSigne, ": Nombre non signé, toujours '+'")
	fmt.Println(pi, ": Costante de Pi -", a, b, c, d, ": Valeurs constante avec incrémentation avec différente variable const")

	fmt.Println("")
	fmt.Println("Variable type chaine :")
	fmt.Println("")
	fmt.Println(monNom)
	fmt.Println(monPrenom)

	fmt.Println("")
	fmt.Println("Variable type booléen")
	fmt.Println("")
	fmt.Println(myTrue, "Vrai")
	fmt.Println(myFalse, "Faux")
}
