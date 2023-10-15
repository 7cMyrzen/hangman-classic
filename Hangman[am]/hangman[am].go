package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func resetTerminal() {
	fmt.Print("\033c")
}

func main() {
	m := color.New(color.FgHiMagenta)
	hc := color.New(color.FgHiCyan)
	r := color.New(color.FgHiRed)
	resetTerminal()
	var choice string
	m.Println("Voulez-vous lancer le pendu de maniere automatique [a] ou voulez vous choisir le mot a faire deviner manuellement [m] ?")
	fmt.Scanln(&choice)
	if choice == "a" {
		resetTerminal()
		win := "Bravo, vous avez gagné !"
		lose := "Dommage, vous avez perdu !"
		content, err := ioutil.ReadFile("words.txt")
		if err != nil {
			fmt.Println(err)
		}
		words := string(content)
		var list []string
		list = strings.Split(words, "\n")
		lenlist := len(list)
		randlist := rand.Intn(lenlist)
		word := list[randlist-1]
		tryNumber := 6
		playHangman(word, win, lose, tryNumber)
		return
	}
	if choice == "m" {
		resetTerminal()
		var word string
		var win string
		var lose string
		var tryChoice string
		reader := bufio.NewReader(os.Stdin)
		m.Println("Quel est le mot a faire deviner ?")
		fmt.Scanln(&word)
		resetTerminal()
		m.Println("Quel est le message de victoire ?")
		inputw, _ := reader.ReadString('\n')
		win = inputw
		resetTerminal()
		m.Println("Quel est le message de defaite ?")
		inputl, _ := reader.ReadString('\n')
		lose = inputl
		resetTerminal()
		m.Println("Combien de tentative ?")
		hc.Println("	[6]		6 tentives")
		hc.Println("	[10]		10 tentatives")
		fmt.Scanln(&tryChoice)
		if tryChoice == "6" {
			tryNumber := 6
			playHangman(word, win, lose, tryNumber)
			return
		}
		if tryChoice == "10" {
			tryNumber := 10
			playHangman(word, win, lose, tryNumber)
			return
		} else {
			resetTerminal()
			r.Println("Erreur : ", tryChoice)
			r.Println("Ce choix n'existe pas.")
			time.Sleep(1 * time.Second)
			resetTerminal()
			return

		}
	} else {
		resetTerminal()
		r.Println("Erreur : ", choice)
		r.Println("Ce choix n'existe pas.")
		time.Sleep(3 * time.Second)
		resetTerminal()
		return
	}
}

func playHangman(word string, win string, lose string, tryNumber int) {
	m := color.New(color.FgHiMagenta)
	hc := color.New(color.FgHiCyan)
	r := color.New(color.FgHiRed)
	y := color.New(color.FgHiYellow)
	var maxTry int
	maxTry = tryNumber
	var hiddenword []string
	hiddenword = strings.Split(word, "")
	var hiddenline []string
	for i := 0; i < len(hiddenword); i++ {
		hiddenline = append(hiddenline, "_")
	}
	errcount := 0
	var tried []string

	for nimput := 1; nimput <= 26; nimput++ {
		resetTerminal()
		if maxTry == 6 {
			printHangmanNormal(errcount)
		}
		if maxTry == 10 {
			printHangmanEasy(errcount)
		}
		hc.Println(hiddenline)
		m.Println("Vous avez deja essayé les lettres suivantes: ", tried)
		m.Println("Vous avez encore le droit a ", maxTry-errcount, " erreur(s)")
		m.Println("Veuillez entrer une lettre: ")
		var input string
		fmt.Scanln(&input)
		tried = append(tried, input)

		if strings.Contains(word, input) {
			for i := 0; i < len(hiddenword); i++ {
				if hiddenword[i] == input {
					hiddenline[i] = input
				}
			}
			if strings.Join(hiddenline, "") == word {
				resetTerminal()
				y.Println(hiddenline)
				y.Println(win)
				time.Sleep(2 * time.Second)
				retryHangman()
				break
			}
		} else {
			resetTerminal()
			r.Println("Erreur, cette lettre n'est pas dans le mot")
			errcount++
			time.Sleep(1 * time.Second)

			if maxTry == 6 {
				if errcount == 6 {
					resetTerminal()
					r.Println(lose)
					r.Println("Le mot etait", word)
					time.Sleep(3 * time.Second)
					retryHangman()
					break
				}
			}
			if maxTry == 10 {
				if errcount == 10 {
					resetTerminal()
					r.Println(lose)
					r.Println("Le mot etait", word)
					time.Sleep(3 * time.Second)
					retryHangman()
					break
				}
			}
		}
	}
}

func printHangmanNormal(errcount int) {
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

func printHangmanEasy(errcount int) {
	hangman := []string{
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
		"         ",
	}

	switch errcount {
	case 1:
		hangman[6] = "========="
	case 2:
		hangman[0] = "      +"
		hangman[1] = "      |"
		hangman[2] = "      |"
		hangman[3] = "      |"
		hangman[4] = "      |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 3:
		hangman[0] = "  +---+"
		hangman[1] = "      |"
		hangman[2] = "      |"
		hangman[3] = "      |"
		hangman[4] = "      |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 4:
		hangman[0] = "  +---+"
		hangman[1] = "  |   |"
		hangman[2] = "      |"
		hangman[3] = "      |"
		hangman[4] = "      |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 5:
		hangman[0] = "  +---+"
		hangman[1] = "  |   |"
		hangman[2] = "  O   |"
		hangman[3] = "      |"
		hangman[4] = "      |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 6:
		hangman[0] = "  +---+"
		hangman[1] = "  |   |"
		hangman[2] = "  O   |"
		hangman[3] = "  |   |"
		hangman[4] = "      |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 7:
		hangman[0] = "  +---+"
		hangman[1] = "  |   |"
		hangman[2] = "  O   |"
		hangman[3] = " /|   |"
		hangman[4] = "      |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 8:
		hangman[0] = "  +---+"
		hangman[1] = "  |   |"
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = "      |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 9:
		hangman[0] = "  +---+"
		hangman[1] = "  |   |"
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = " /    |"
		hangman[5] = "      |"
		hangman[6] = "========="
	case 10:
		hangman[0] = "  +---+"
		hangman[1] = "  |   |"
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = " / \\  |"
		hangman[5] = "      |"
		hangman[6] = "========="
	}

	for _, line := range hangman {
		fmt.Println(line)
	}
}

func retryHangman() {
	hc := color.New(color.FgHiCyan)
	r := color.New(color.FgHiRed)
	resetTerminal()
	hc.Println("Voulez-vous recommencer le pendu [o/n] ?")
	var retry string
	fmt.Scanln(&retry)
	if retry == "o" {
		main()
		return
	}
	if retry == "n" {
		resetTerminal()
		return
	}
	if retry != "n" && retry != "o" {
		resetTerminal()
		r.Println("Erreur : [o] pour oui et [n] pour non.")
		time.Sleep(2 * time.Second)
		resetTerminal()
		return
	}
}
