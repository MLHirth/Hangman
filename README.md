# Hangman Game in Go

## Introduction
This is a simple Hangman game implemented in Go. The game allows players to guess letters of a randomly selected word until they either correctly guess the word or run out of attempts.

## Features
- Single-player Hangman gameplay
- Words are embedded into the binary using Go's `embed` package (no external file needed)
- Case-insensitive letter guessing
- Proper input validation to prevent duplicate guesses and invalid characters
- Console-based UI with clear instructions

## Installation
### Clone the Repository
```sh
git clone https://github.com/MLHirth/Hangman.git
cd Hangman
```

### Build the Executable
```sh
go build Hangman.go
```

### Run the Game
```sh
./Hangman
```

## How to Play
1. Start the game by running `./hangman`.
2. You will be prompted to choose whether to start a game.
3. The game will either let you provide a word or select a random word.
4. Guess letters one by one. Correct guesses will reveal the letters in the word.
5. The game ends when you either guess the word correctly or run out of guesses.

## File Structure
```
HangmanGo/
│── go.mod
│── Hangman.go  # Main game logic
│── WordList.txt  # Embedded word list (only needed for modifications)
```

## Embedding WordList.txt
This project embeds `WordList.txt` directly into the Go binary. To modify the word list:
1. Edit `WordList.txt`
2. Rebuild the binary:
   ```sh
   go build Hangman.go
   ```

## Troubleshooting
**Problem:** "No words found" error.
- Ensure `WordList.txt` is embedded properly using `//go:embed`.
- Try rebuilding the program: `go build Hangman.go`

**Problem:** Game does not recognize guessed letters.
- Ensure you are entering only **one letter at a time**.
- Check that the letter is present in the word list.

## Contact
For issues or questions, create an issue on GitHub or contact the repository owner.

