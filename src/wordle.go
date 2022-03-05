package main

import (
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
  line  string
  word  string
  words []string
}

func Init() State {
  words := read()
  return State{
    board: []string{},
    line:  "",
    word:  word(words),
    words: words,
  }
}

func (self *State) Handle(guess string) {
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

  i := 0
  for i < GUESSES {
    state.Print()

    guess := prompt()
    for len(guess) != len(state.word) || !state.Contains(guess) {
      fmt.Println("Invalid guess.")
      guess = prompt()
    }

    state.Handle(guess)

    if guess == state.word {
      state.Print()
      fmt.Printf("word: %s\n", state.word)
      break
    }

    i -= 1
  }
}
