package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func main() {
	//Lire le fichier words.txt
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	words := string(content)

	//Creer une liste de mots et sa longueur
	var list []string
	list = strings.Split(words, "\n")
	//fmt.Println(list)
	lenlist := len(list)
	//fmt.Println(lenlist)

	//Choisir un mot au hasard
	randlist := rand.Intn(lenlist)
	//fmt.Println(randlist)
	hidenword := list[randlist-1]
	//fmt.Println(hidenword)

	//Creer une liste de lettres du mot choisi
	var hidenstr []string
	hidenstr = strings.Split(hidenword, "")
	//fmt.Println(hidenstr)

	//Creer un affichage du mot choisi avec le nombre de lettres
	var hidenline []string
	for i := 0; i < len(hidenstr); i++ {
		hidenline = append(hidenline, "_")
	}
	fmt.Println(hidenline)

	//Hangman

	//Definir variable erreur, un message victoire et un message defaite
	errcount := 0
	win := "Bravo, vous avez gagné !"
	lose := "Dommage, vous avez perdu !"

	//Definir une variable avec les lettres deja essayees
	var tried []string

	for nimput := 1; nimput <= 26; nimput++ {
		fmt.Println("Vous avez deja essayé les lettres suivantes: ", tried)
		fmt.Println("Vous avez encore le droit a ", 6-errcount, " erreur(s)")
		fmt.Println("Veuillez entrer une lettre: ")
		var input string
		fmt.Scanln(&input)
		tried = append(tried, input)
		//Si la lettre est dans le mot, afficher le mot avec la lettre
		if strings.Contains(hidenword, input) {
			for i := 0; i < len(hidenstr); i++ {
				if hidenstr[i] == input {
					hidenline[i] = input
				}
			}
			fmt.Println(hidenline)
			//Si le mot est complet, afficher le message victoire et arreter le programme
			if strings.Join(hidenline, "") == hidenword {
				fmt.Println(win)
				break
			}
		} else {
			//Si la lettre n'est pas dans le mot, afficher l'erreur et augmenter errcount
			fmt.Println("Erreur, cette lettre n'est pas dans le mot")
			errcount++
			PrintHangman(errcount)
			//Si errcount est egal a 6, afficher le message defaite et arreter le programme
			if errcount == 6 {
				fmt.Println(lose)
				break
			}
		}
	}
}

func PrintHangman(errcount int) {
	hangman := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}

	switch errcount {
	case 1:
		hangman[2] = "  O   |"
	case 2:
		hangman[2] = "  O   |"
		hangman[3] = "  |   |"
	case 3:
		hangman[2] = "  O   |"
		hangman[3] = " /|   |"
	case 4:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
	case 5:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = " /    |"
	case 6:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = " / \\  |"
	}

	for _, line := range hangman {
		fmt.Println(line)
	}
}
