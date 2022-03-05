package main

import (
  "errors"
  "fmt"
  "io/ioutil"
  "math/rand"
  "strings"
  "time"
)

const WORDS = "assets/words.txt"
const GUESSES = 6
const B = "⬛"
const G = "🟩"
const Y = "🟨"

type State struct {
  board []string
  word  string
  words []string
}

func Init() State {
  words := read()
  return State{
    board: []string{},
    word:  word(words),
    words: words,
  }
}

func (self *State) Handle(guess string) error {
  if !self.Valid(guess) {
    return errors.New("Invalid guess.")
  }

  line := ""
  for i, char := range guess {
    if self.word[i] == byte(char) {
      line += G
    } else if strings.Contains(self.word, string(char)) {
      line += Y
    } else {
      line += B
    }
  }

  self.board = append(self.board, line)
  return nil
}

func (self *State) Line() string {
  if len(self.board) != 0 {
    return self.board[len(self.board)-1]
  }
  return ""
}

func (self *State) Contains(word string) bool {
  for _, have := range self.words {
    if have == word {
      return true
    }
  }
  return false
}

func (self *State) Print() {
  for _, line := range self.board {
    fmt.Println(line)
  }
}

func (self *State) Valid(guess string) bool {
  return len(guess) == len(self.word) && self.Contains(guess)
}

func read() []string {
  r, e := ioutil.ReadFile(WORDS)
  if e != nil {
    panic(e)
  }
  return strings.Split(string(r), "\n")
}

func word(words []string) string {
  rand.Seed(time.Now().UnixNano())
  return words[rand.Intn(len(words))]
}

func prompt() string {
  fmt.Print("> ")
  var input string
  fmt.Scanln(&input)
  return input
}

func main() {
  state := Init()

  i := GUESSES
  for true {
    state.Print()

    if i == 0 {
      fmt.Printf("word: %s\n", state.word)
      break
    }

    guess := prompt()
    for true {
      if err := state.Handle(guess); err != nil {
        fmt.Println(err)
        guess = prompt()
      } else {
        break
      }
    }

    if guess == state.word {
      state.Print()
      fmt.Printf("word: %s\n", state.word)
      break
    }

    i -= 1
  }
}
