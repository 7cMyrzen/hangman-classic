package main

//Imports necessaires au bon fonctionnement du programme
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

// Fonction pour reset le terminal
func resetTerminal() {
	fmt.Print("\033c")
}

// Fonction principale
func main() {
	//Definition des couleurs
	m := color.New(color.FgHiMagenta)
	hc := color.New(color.FgHiCyan)
	r := color.New(color.FgHiRed)

	//Reset du terminal
	resetTerminal()

	//Definition dU choix de mode de jeu
	var choice string
	m.Println("Voulez-vous lancer le pendu de maniere automatique [a] ou voulez vous choisir le mot a faire deviner manuellement [m] ?")
	fmt.Scanln(&choice)

	//Si le choix est automatique
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

		//Lancement du jeu avec les parametres automatiques
		playHangman(word, win, lose, tryNumber)
		return
	}

	//Si le choix est manuel
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
		hc.Println("	[4]		4 tentives")
		hc.Println("	[6]		6 tentives")
		hc.Println("	[10]		10 tentatives")
		fmt.Scanln(&tryChoice)

		//Lancement du jeu avec les parametres manuels
		if tryChoice == "4" {
			tryNumber := 4
			playHangman(word, win, lose, tryNumber)
			return
		}
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

		//Si le choix n'existe pas
	} else {
		resetTerminal()
		r.Println("Erreur : ", choice)
		r.Println("Ce choix n'existe pas.")
		time.Sleep(3 * time.Second)
		resetTerminal()
		return
	}
}

// Fonction pour lancer le jeu
func playHangman(word string, win string, lose string, tryNumber int) {
	//Definition des couleurs
	m := color.New(color.FgHiMagenta)
	hc := color.New(color.FgHiCyan)
	r := color.New(color.FgHiRed)
	y := color.New(color.FgHiYellow)
	//Definition des variables
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

	//Boucle pour lancer le jeu
	for nimput := 1; nimput <= 26; nimput++ {

		//Reset du terminal a chaque tour de boucle
		resetTerminal()

		//Affichage du pendu en fonction du nombre de tentatives
		if maxTry == 4 {
			printHangmanHard(errcount)
		}
		if maxTry == 6 {
			printHangmanNormal(errcount)
		}
		if maxTry == 10 {
			printHangmanEasy(errcount)
		}

		//Affichage du mot cache
		hc.Println(hiddenline)

		//Affichage des instructions
		m.Println("Vous avez deja essayé les lettres suivantes: ", tried)
		m.Println("Vous avez encore le droit a ", maxTry-errcount, " erreur(s)")
		m.Println("Veuillez entrer une lettre: ")

		//Definition de la variable input et ajout de la lettre essayee dans la liste des lettres essayees
		var input string
		fmt.Scanln(&input)
		tried = append(tried, input)

		//Verification de la presence de la lettre dans le mot et affichage
		if strings.Contains(word, input) {
			for i := 0; i < len(hiddenword); i++ {
				if hiddenword[i] == input {
					hiddenline[i] = input
				}
			}

			//Verification de la victoire
			if strings.Join(hiddenline, "") == word {
				resetTerminal()
				y.Println(hiddenline)
				y.Println(win)
				time.Sleep(2 * time.Second)

				//Demande de recommencer
				retryHangman()
				break
			}

			//Si la lettre n'est pas dans le mot
		} else {

			//Affichage d'une erreur
			resetTerminal()
			r.Println("Erreur, cette lettre n'est pas dans le mot")
			errcount++
			time.Sleep(1 * time.Second)

			//Verification si le nombre d'erreur est egal au nombre de tentatives
			if maxTry == 4 {
				if errcount == 4 {

					//Gestion de la defaite
					resetTerminal()
					r.Println(lose)
					r.Println("Le mot etait", word)
					time.Sleep(3 * time.Second)
					retryHangman()
					break
				}
			}
			if maxTry == 6 {
				if errcount == 6 {

					//Gestion de la defaite
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

					//Gestion de la defaite
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

// Fonction pour afficher le pendu {niveau difficile}
func printHangmanHard(errcount int) {
	//Creation du pendu vide
	hangman := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}

	//Modification du pendu en fonction du nombre d'erreur
	switch errcount {
	case 1:
		hangman[2] = "  O   |"
	case 2:
		hangman[2] = "  O   |"
		hangman[3] = "  |   |"
	case 3:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
	case 4:
		hangman[2] = "  O   |"
		hangman[3] = " /|\\  |"
		hangman[4] = " / \\  |"
	}

	//Affichage du pendu
	for _, line := range hangman {
		fmt.Println(line)
	}
}

// Fonction pour afficher le pendu {niveau normal}
func printHangmanNormal(errcount int) {
	//Creation du pendu vide
	hangman := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}

	//Modification du pendu en fonction du nombre d'erreur
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

	//Affichage du pendu
	for _, line := range hangman {
		fmt.Println(line)
	}
}

// Fonction pour afficher le pendu {niveau facile}
func printHangmanEasy(errcount int) {
	//Creation du pendu vide
	hangman := []string{
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
		"       ",
		"         ",
	}

	//Modification du pendu en fonction du nombre d'erreur
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

	//Affichage du pendu
	for _, line := range hangman {
		fmt.Println(line)
	}
}

// Fonction pour recommencer le jeu
func retryHangman() {
	//Definition des couleurs
	hc := color.New(color.FgHiCyan)
	r := color.New(color.FgHiRed)
	resetTerminal()

	//Demande de recommencer
	hc.Println("Voulez-vous recommencer le pendu [o/n] ?")
	var retry string
	fmt.Scanln(&retry)

	//Si oui
	if retry == "o" {

		//Relancement du jeu
		main()
		return
	}

	//Si non
	if retry == "n" {

		//Reset du terminal et fin du programme
		resetTerminal()
		return
	}

	//Si le choix n'est pas valide
	if retry != "n" && retry != "o" {

		//Affichage d'une erreur
		resetTerminal()
		r.Println("Erreur : [o] pour oui et [n] pour non.")
		time.Sleep(2 * time.Second)
		resetTerminal()
		return
	}
}
