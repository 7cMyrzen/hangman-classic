package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func resetTerminal() {
	fmt.Print("\033c")
}

func main() {
	resetTerminal()
	var choice string
	fmt.Println("Voulez-vous lancer le pendu de maniere automatique [a] ou voulez vous choisir le mot a faire deviner manuellement [m] ? ")
	fmt.Scanln(&choice)
	if choice == "a" {
		resetTerminal()
		autoHangman()
		return
	}
	if choice == "m" {
		resetTerminal()
		manualHangman()
		return
	} else {
		resetTerminal()
		fmt.Println("Pitie apprends a lire, il y avait que 2 lettres et t'arrives a en choisir une autre.")
	}
}

func autoHangman() {
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	words := string(content)

	var list []string
	list = strings.Split(words, "\n")
	lenlist := len(list)

	randlist := rand.Intn(lenlist)
	hiddenword := list[randlist-1]

	var hiddenstr []string
	hiddenstr = strings.Split(hiddenword, "")

	var hiddenline []string
	for i := 0; i < len(hiddenstr); i++ {
		hiddenline = append(hiddenline, "_")
	}

	errcount := 0
	win := "Bravo, vous avez gagné !"
	lose := "Dommage, vous avez perdu !"

	var tried []string

	for nimput := 1; nimput <= 26; nimput++ {
		resetTerminal()
		PrintHangman(errcount)
		fmt.Println(hiddenline)
		fmt.Println("Vous avez deja essayé les lettres suivantes: ", tried)
		fmt.Println("Vous avez encore le droit a ", 6-errcount, " erreur(s)")
		fmt.Println("Veuillez entrer une lettre: ")
		var input string
		fmt.Scanln(&input)
		tried = append(tried, input)

		if strings.Contains(hiddenword, input) {
			for i := 0; i < len(hiddenstr); i++ {
				if hiddenstr[i] == input {
					hiddenline[i] = input
				}
			}
			//fmt.Println(hiddenline)
			if strings.Join(hiddenline, "") == hiddenword {
				resetTerminal()
				fmt.Println(hiddenline)
				fmt.Println(win)
				time.Sleep(2 * time.Second)
				retryHangman()
				break
			}
		} else {
			resetTerminal()
			fmt.Println("Erreur, cette lettre n'est pas dans le mot")
			errcount++
			time.Sleep(1 * time.Second)

			if errcount == 6 {
				resetTerminal()
				fmt.Println(lose)
				fmt.Println("Le mot etait", hiddenword)
				time.Sleep(2 * time.Second)
				retryHangman()
				break
			}
		}
	}
}

func manualHangman() {
	resetTerminal()
	var hiddenword string
	fmt.Println("Entrez le mot a faire deviner (pitie fait pas de faute d'orthographe)")
	fmt.Scanln(&hiddenword)

	var hiddenstr []string
	hiddenstr = strings.Split(hiddenword, "")

	var hiddenline []string
	for i := 0; i < len(hiddenstr); i++ {
		hiddenline = append(hiddenline, "_")
	}

	errcount := 0
	win := "Bravo, vous avez gagné !"
	lose := "Dommage, vous avez perdu !"

	var tried []string

	for nimput := 1; nimput <= 26; nimput++ {
		resetTerminal()
		PrintHangman(errcount)
		fmt.Println(hiddenline)
		fmt.Println("Vous avez deja essayé les lettres suivantes: ", tried)
		fmt.Println("Vous avez encore le droit a ", 6-errcount, " erreur(s)")
		fmt.Println("Veuillez entrer une lettre: ")
		var input string
		fmt.Scanln(&input)
		tried = append(tried, input)

		if strings.Contains(hiddenword, input) {
			for i := 0; i < len(hiddenstr); i++ {
				if hiddenstr[i] == input {
					hiddenline[i] = input
				}
			}
			//fmt.Println(hiddenline)
			if strings.Join(hiddenline, "") == hiddenword {
				resetTerminal()
				fmt.Println(hiddenline)
				fmt.Println(win)
				time.Sleep(2 * time.Second)
				retryHangman()
				break
			}
		} else {
			resetTerminal()
			fmt.Println("Erreur, cette lettre n'est pas dans le mot")
			errcount++
			time.Sleep(1 * time.Second)

			if errcount == 6 {
				resetTerminal()
				fmt.Println(lose)
				fmt.Println("Le mot etait", hiddenword)
				time.Sleep(2 * time.Second)
				retryHangman()
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

func retryHangman() {
	resetTerminal()
	fmt.Println("Voulez-vous recommencer le pendu [o/n] ?")
	var retry string
	fmt.Scanln(&retry)
	if retry == "o" {
		main()
	}
	if retry == "n" {
		resetTerminal()
		return
	}
	if retry != "n" && retry != "o" {
		resetTerminal()
		fmt.Println("[o] pour oui et [n] pour non, c'est pourtant pas complique.")
	}
}
