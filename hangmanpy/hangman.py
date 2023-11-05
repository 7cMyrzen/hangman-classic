import os
import random
import time
from colorama import Fore, init
init(autoreset=True)

def clearTerminal():
    os.system('cls')

def ListTextFiles():
    files = []
    for file in os.listdir():
        if file.endswith(".txt"):
            files.append(file)
    return files

def pickWord(file):
    with open(file, "r") as wordsFile:
        content = wordsFile.read() 
        wordsList = content.splitlines()
        randomWord = random.choice(wordsList)
        return randomWord
        

def startFunction():
    clearTerminal()
    print(Fore.YELLOW + "Choix du mode de jeu :")
    print(Fore.YELLOW + "   [a] Jouer en mode automatique.")
    print(Fore.YELLOW + "   [m] Jouer en mode manuel.")
    print('')
    userChoice = input(Fore.YELLOW + "Votre choix : ")
    if userChoice == 'a':
        clearTerminal()
        print(Fore.GREEN + "Mode automatique")
        time.sleep(1)
        clearTerminal()
        win = "Bien joué, vous avez gagné !"
        lose = "Dommage, vous avez perdu !"
        txtFiles = ListTextFiles()
        autoFile = random.choice(txtFiles)
        word = pickWord(autoFile)
        userTry = 6
    if userChoice == 'm':
        clearTerminal()
        print(Fore.GREEN + "Mode manuel")
        time.sleep(1)
        clearTerminal()
        print(Fore.YELLOW + "Choix du mot:")
        print(Fore.YELLOW + "   [1] Choisir un mot au hasard dans un fichier texte de votre choix.")
        print(Fore.YELLOW + "   [2] Choisir votre propre mot a faire deviner.")
        print('')
        userChoice = input(Fore.YELLOW + "Votre choix : ")
        if userChoice == '1':
            clearTerminal()
            print(Fore.YELLOW + "Choix du fichier texte :")
            print('')
            txtFiles = ListTextFiles()
            for file in txtFiles:
                print(Fore.YELLOW + file)
            print('')
            userChoice = input(Fore.YELLOW + "Votre choix : ")
            word = pickWord(userChoice)
        if userChoice == '2':
            clearTerminal()
            word = input(Fore.YELLOW + "Veuillez entrer un mot : ")
        clearTerminal()
        win = input(Fore.YELLOW + "Veuillez entrer un message de victoire : ")
        clearTerminal()
        lose = input(Fore.YELLOW + "Veuillez entrer un message de défaite : ")
        clearTerminal()
        print(Fore.YELLOW + "Veuillez choisir le nombre d'erreurs autorisées :")
        print(Fore.YELLOW + "   [4] 4 tentatives.")
        print(Fore.YELLOW + "   [6] 6 tentatives.")
        print(Fore.YELLOW + "   [10] 10 tentatives.")
        print('')
        userChoice = input(Fore.YELLOW + "Votre choix : ")
        if userChoice == '4':
            userTry = 4
        if userChoice == '6':
            userTry = 6
        if userChoice == '10':
            userTry = 10
    clearTerminal()
    playHangman(win, lose, word, userTry)

def playHangman(win, lose, word, userTry):
    maxTry = userTry
    errcount = 0
    triedLetters = []
    hiddenWord = []
    for i in range(len(word)):
        hiddenWord.append("_")
    hidden_word = list(word)

    helpchar = ["&", " ", ",", ".", "-", "_"]
    for elmt in helpchar:
        if elmt in hidden_word:
                    for i in range(len(hidden_word)):
                        if hidden_word[i] == elmt:
                            hiddenWord[i] = elmt
                            hidden_word[i] = elmt

    while errcount < maxTry:
        clearTerminal()
        
        #affichage du pendu
        if maxTry == 4:
            printHangmanHard(errcount)
        if maxTry == 6:
            printHangmanNormal(errcount)
        if maxTry == 10:
            printHangmanEasy(errcount)

        hidden_word_str = " ".join(hiddenWord)
        tried_letters_str = " ".join(triedLetters)
        print(Fore.CYAN + "Mot à deviner : " + Fore.RED + hidden_word_str)
        print(Fore.CYAN + "Vous avez déjà essayé les lettres suivantes : " + Fore.RED + tried_letters_str)
        print(Fore.CYAN + "Il vous reste " + Fore.RED + str(maxTry - errcount) + Fore.CYAN + " tentatives.")
        inputLetter = input(Fore.CYAN + "Veuillez entrer une lettre : ")

        if inputLetter in triedLetters:
            print(Fore.RED + "Vous avez déjà essayé cette lettre !")
            time.sleep(2)
            continue
        else:
            triedLetters.append(inputLetter)
            if inputLetter in hidden_word:
                for i in range(len(hidden_word)):
                    if hidden_word[i] == inputLetter:
                        hiddenWord[i] = inputLetter
                        hidden_word[i] = "_"
            else:
                errcount += 1
                print(Fore.RED + "Cette lettre ne fait pas partie du mot !")
                time.sleep(2)
                continue
        if "_" not in hiddenWord:
            clearTerminal()
            print(Fore.GREEN + win)
            time.sleep(2)
            restart()
    clearTerminal()
    print(Fore.RED + lose)
    time.sleep(2)
    return restart()

def printHangmanHard(errcount):

    hangman = [
        "  +---+",
        "  |   |",
        "      |",
        "      |",
        "      |",
        "      |",
        "=========",
    ]

    if errcount == 1:
        hangman[2] = "  O   |"
    elif errcount == 2:
        hangman[2] = "  O   |"
        hangman[3] = "  |   |"
    elif errcount == 3:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
    elif errcount == 4:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
        hangman[4] = " / \\  |"

    for line in hangman:
        print(Fore.CYAN + line)

def printHangmanNormal(errcount):

    hangman = [
        "  +---+",
        "  |   |",
        "      |",
        "      |",
        "      |",
        "      |",
        "=========",
    ]

    if errcount == 1:
        hangman[2] = "  O   |"
    elif errcount == 2:
        hangman[2] = "  O   |"
        hangman[3] = "  |   |"
    elif errcount == 3:
        hangman[2] = "  O   |"
        hangman[3] = " /|   |"
    elif errcount == 4:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
    elif errcount == 5:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
        hangman[4] = " /    |"
    elif errcount == 6:
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
        hangman[4] = " / \\  |"

    for line in hangman:
        print(Fore.CYAN + line)

def printHangmanEasy(errcount):

    hangman = [
        "       ",
        "       ",
        "       ",
        "       ",
        "       ",
        "       ",
        "         ",
    ]

    if errcount >= 1:
        hangman[6] = "========="
    if errcount >= 2:
        hangman[0] = "      +"
        hangman[1] = "      |"
        hangman[2] = "      |"
        hangman[3] = "      |"
        hangman[4] = "      |"
        hangman[5] = "      |"
        hangman[6] = "========="
    if errcount >= 3:
        hangman[0] = "  +---+"
    if errcount >= 4:
        hangman[1] = "      |"
    if errcount >= 5:
        hangman[2] = "      |"
    if errcount >= 6:
        hangman[3] = "      |"
    if errcount >= 7:
        hangman[4] = "      |"
    if errcount >= 8:
        hangman[5] = "      |"
    if errcount >= 9:
        hangman[0] = "  +---+"
    if errcount >= 10:
        hangman[1] = "  |   |"
        hangman[2] = "  O   |"
        hangman[3] = " /|\\  |"
        hangman[4] = " / \\  |"

    for line in hangman:
        print(Fore.CYAN + line)

def restart():
    clearTerminal()
    print(Fore.YELLOW + "Voulez-vous rejouer ?")
    print(Fore.YELLOW + "   [o] Oui")
    print(Fore.YELLOW + "   [n] Non")
    print('')
    userChoice = input(Fore.YELLOW + "Votre choix : ")
    if userChoice == 'o':
        startFunction()
    if userChoice == 'n':
        clearTerminal()
        print(Fore.GREEN + "Merci d'avoir joué !")
        time.sleep(2)
        clearTerminal()
        exit()
    else:
        restart()

startFunction()
