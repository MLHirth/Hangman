package main

import (
	"bufio" //buffered reader to read in the data
	_ "embed"
	"fmt"       //used for printing, scanning, etc
	"math/rand" //random package to produce random number
	"os"        //os package to open the file
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

//go:embed WordList.txt
var WordList string

func main() {
	StartGame()
}

func StartGame() {
	fmt.Println("Starting Game")
	fmt.Println("Would you like to start a game?")
	var choice string
	fmt.Scan(&choice)
	if choice == "start" || choice == "yes" || choice == "y" {
		InitializeGame()
	}
}

func InitializeGame() {
	fmt.Println("Initializing Game")
	fmt.Println("Would you like to choose the word?")
	var choice string
	fmt.Scan(&choice)
	if choice == "start" || choice == "yes" || choice == "y" {
		fmt.Println("Please provide a word")
		var word string
		fmt.Scan(&word)
		Hangman(strings.ToLower(word))
	} else {
		fmt.Println("A random word will be selected")
		Hangman(RandomWord())
	}

}

func RandomWord() string {
	words := strings.Split(strings.TrimSpace(WordList), "\n") // ✅ Correctly process the embedded file
	if len(words) == 0 {
		fmt.Println("Error: No words found in embedded file.")
		panic("Word list is empty")
	}
	randomNumber := rand.Intn(len(words)) // Get a random index
	return strings.ToLower(words[randomNumber])
}

func Hangman(word string) {
	var CorrectGuesses = FillInitialArray(word)
	GuessedSet := make(map[string]bool)
	var guessCounter = 0

	for strings.Join(CorrectGuesses, "") != word {
		str := "[" + strings.Join(CorrectGuesses, ", ") + "]"
		fmt.Println(str + " Number of guesses: " + strconv.Itoa(guessCounter))
		fmt.Println("Enter a letter: ")
		var letter string
		reader := bufio.NewReader(os.Stdin) // Fix input handling
		letter, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		letter = strings.TrimSpace(letter) // Fix newline issue
		if (len(letter) < len(word)) && (!GuessedSet[strings.ToLower(letter)]) {
			if len(letter) == 1 {
				if strings.Contains(word, strings.ToLower(letter)) {
					for i := 0; i < len(word); i++ {
						if string(word[i]) == letter {
							fmt.Println(letter+" exists, and is at position:", i)
							CorrectGuesses[i] = letter
						}
					}
				} else {
					fmt.Println(letter + " doesn't exist")
				}
			} else {
				fmt.Println(letter + " is not a single letter")
			}
			GuessedSet[strings.ToLower(letter)] = true
		} else if strings.EqualFold(letter, "hint") {
			for i := 0; i < len(CorrectGuesses); i++ {
				if GuessedSet["_"] {
					fmt.Println("hint: try entering the letter: " + CorrectGuesses[i])
					break
				}
			}
		} else if strings.EqualFold(letter, word) {
			if guessCounter == 0 {
				fmt.Println("You guessed the word " + word + " on your first try!! Are you sure you didn't know it beforehand? ;)")
			} else {
				fmt.Println(word + " is correct, you have guessed it correctly, after " + strconv.Itoa(guessCounter) + " guesses")
			}
			os.Exit(0)
		} else {
			if len(letter) == 1 {
				fmt.Println("You have already guessed this letter")
			} else if len(letter) > len(word) {
				fmt.Println("The word you entered is too long!!")
			} else {
				fmt.Println("You have already guessed this combination of letters or the word!!")
			}
		}
		str2 := "[" + strings.Join(CorrectGuesses, ", ") + "]"
		fmt.Println(str2)
		guessCounter++
		ClearConsole()
	}
	fmt.Println("Yay you have guessed it correctly!! You got it in " + strconv.Itoa(guessCounter) + " guesses")
}

func FillInitialArray(word string) []string {
	var CorrectGuesses = make([]string, 0)
	for i := 0; i < len(word); i++ {
		CorrectGuesses = append(CorrectGuesses, "_")
	}
	return CorrectGuesses
}

func WordToList(word string) []string {
	var TargetWord []string
	for _, letter := range word {
		TargetWord = append(TargetWord, string(letter))
	}
	return TargetWord
}

func ClearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows command
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	default:
		cmd := exec.Command("clear") // Unix/Linux/macOS command
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}
