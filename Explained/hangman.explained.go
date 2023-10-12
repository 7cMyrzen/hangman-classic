// package + import necessaire au fonctionnement du programme
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
	//Lire le document words.txt avec ioutil et prendre le contenu(content) et l'erreur en cas de probleme(err)

	if err != nil {
		//Si il y a une erreur

		fmt.Println(err)
		// Afficher l'erreur
	}
	words := string(content)
	//Stocker le contenu de words.txt dans la variable words

	//Creer une liste de mots et sa longueur

	var list []string
	//Creer une liste(pour creer une liste on ajoute [] avant d'ecrire le type) du nom de list vide pour mettre des chaines de caracteres(string)

	list = strings.Split(words, "\n")
	//Vu que tous les mots sont separes par un retour a la ligne(/n), on ajoute a la liste(list) tous les mots en les separants quand on trouve un /n(retour a la ligne)

	//fmt.Println(list)
	//Test pour voir la liste precedement cree

	lenlist := len(list)
	//Creer une variable lenlist dans laquelle on stocke la longueur de la liste precedement cree en utilisant len(list)

	//fmt.Println(lenlist)
	//Test pour voir la longueur de la liste(lenlist) precedement cree

	//Choisir un mot au hasard

	randlist := rand.Intn(lenlist)
	//Creer une variable randlist pour stocker un nombre aleatoire plus petit que la longueur de la liste en utilisant rand.Intn(lenlist) <- lenlist sert a dire a la fonction rand.Intn de ne pas prendre un nombre plus grand que la longueur de la liste

	//fmt.Println(randlist)
	//Test pour afficher le chiffre aleatoire precedement genere

	hidenword := list[randlist-1]
	//Creer une variable hidenword qui stocke le mot de la liste(list) a la position du chiffre aleatoire(randlist) <- on utilise randlist -1 car les index commencent  0 (exemple: si on veut la 15eme valeur de list on ecrit list[15-1])

	//fmt.Println(hidenword)
	//Test pour verifier si hidenword stocke bien un mot leatoire de list

	//Creer une liste de lettres du mot choisi

	var hidenstr []string
	//Creer une variable hidenstr qui stocke une liste de type texte(string)

	hidenstr = strings.Split(hidenword, "")
	//On ajoute a la variable hidensstr un a un les caracteres qui compose le mot stocke dans hidenword (exemple: pour le mot voiture hiden str sera egale a: [v o i t u r e])

	//fmt.Println(hidenstr)
	//Test pour verifier le contenu de hidenstr

	//Creer un affichage du mot choisi avec le nombre de lettres

	var hidenline []string
	//Creer une variable hidenline qui stocke une liste de type texte(string)

	for i := 0; i < len(hidenstr); i++ {
		//On cree une boucle qui se repete autant de fois qu'il y a de caractere dans hidenstr

		hidenline = append(hidenline, "_")
		//Pour chaque repetition de la boucle, on ajoute _ a hidenline (exemple: si hidenstr = [v o i t u r e] hidenline sera egale a: [_ _ _ _ _ _ _])
	}
	fmt.Println(hidenline)
	//Test pour verifier le contenu de hidenline

	//Hangman

	//Definir variable erreur, un message victoire et un message defaite

	errcount := 0
	//Creer une variable de type nombre entier errcount qui va compter les erreur pour arreter le programme si besoin

	win := "Bravo, vous avez gagné !"
	//Creer une variable avec le message de victoire

	lose := "Dommage, vous avez perdu !"
	//Crrer une variable avec le message de defaite

	//Definir une variable avec les lettres deja essayees

	var tried []string
	//Creer une variable tried de type texte pour ensuite stocker les lettres deja ecrite pour eviter a l'utilisateur de la mettre 2 fois

	for nimput := 1; nimput <= 26; nimput++ {
		//On creer une boucle qui commence a 1 et s'arrete a 26 car il y a que 26 lettres pour que sa stoppe si quelqu'un spam la meme lettre

		fmt.Println("Vous avez deja essayé les lettres suivantes: ", tried)
		//On affiche a l'utilisateur les lettres deja essaye

		fmt.Println("Vous avez encore le droit a ", 6-errcount, " erreur(s)")
		//On lui affiche ensuite le nombre d'erreur possible restant avent de finir pendu en faisant le calcul 6 moins erreur comise stocke dans errcount

		fmt.Println("Veuillez entrer une lettre: ")
		//On demande a l'utilisateur d'ecrire une lettre

		var input string
		//On creer une variable input de type texte pour stocker la lettre de l'utilisateur

		fmt.Scanln(&input)
		//On regarde la lettre tape par l'utilisateur avec Scanln et on l'ajoute a la variable input

		tried = append(tried, input)
		//On ajoute a la liste tried la lettre de l'utilisateur

		//Si la lettre est dans le mot, afficher le mot avec la lettre

		if strings.Contains(hidenword, input) {
			//On test si la lettre de l'utilisateur(input) est dans hidenstr avec la methode Contains de strings

			for i := 0; i < len(hidenstr); i++ {
				//Dans le cas ou elle est dans hidenstr on creer une boucle avec i qui va de 0 a la taille de hidenstr

				if hidenstr[i] == input {
					//On test si hidenstr de la position i est egale a la lettre de l'utilisateur(input)

					hidenline[i] = input
					//Si c'est le cas dans hidenline on remplace le _ a la position i par la lettre de l'utilisateur (exemple: si le mot etait parapluie est que la lettre de l'utilisateur etait a hidenline sera egale a: [_ a _ a _ _ _ _ _])
				}
			}
			fmt.Println(hidenline)
			//On affiche hidenline qui a ete mis a jour (pour l'exemple de la ligne dessus, on aura: [_ a _ a _ _ _ _ _])

			//Si le mot est complet, afficher le message victoire et arreter le programme

			if strings.Join(hidenline, "") == hidenword {
				//On test si hidenline est egale a hidenword ce qui veut dire que l'utilisateur a le bon mot

				fmt.Println(win)
				//Si c'est le cas on affiche le message de victoire

				break
				//On utilise break pour arreter le programme pour eviter qu'il continu
			}
		} else {
			//Dans le cas ou l'utilisateur a pas la bonne lettre

			//Si la lettre n'est pas dans le mot, afficher l'erreur et augmenter errcount

			fmt.Println("Erreur, cette lettre n'est pas dans le mot")
			//On affiche comme quoi il s'est trompe

			errcount++
			//On ajoute 1 au nombre d'erreur

			PrintHangman(errcount)
			//On dit a la fonction d'afficher le pendu pour le nombre errcount d'erreur

			//Si errcount est egal a 6, afficher le message defaite et arreter le programme

			if errcount == 6 {
				//On test si errcount est egal a 6 ce qui veut dire que l'utilisateur a perdu

				fmt.Println(lose)
				//On affiche le message de defaite

				break
				// On arrete le programme
			}
		}
	}
}

func PrintHangman(errcount int) {
	//Tableau du pendu vide
	hangman := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}

	//On modifie les lignes du tableau dans le cas ou errcount est egale a 1,2,3,4,5,6
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

	//On affiche le tableau
	for _, line := range hangman {
		fmt.Println(line)
	}
}
