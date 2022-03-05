package main

import (
  "github.com/stretchr/testify/assert"
  _ "github.com/terror/wordle/testing"
  "testing"
)

func TestBasic(t *testing.T) {
  assert := assert.New(t)

  state := Init()
  state.word = "block"

  state.Handle("crane")
  assert.Equal(state.line, "🟨⬛⬛⬛⬛", "state.line != expected line")

  state.Handle("bloke")
  assert.Equal(state.line, "🟩🟩🟩🟨⬛", "state.line != expected line")

  state.Handle("block")
  assert.Equal(state.line, "🟩🟩🟩🟩🟩", "state.line != expected line")
}

func TestInvalid(t *testing.T) {
  assert := assert.New(t)

  state := Init()
  state.word = "block"

  state.Handle("l")
  assert.Equal(state.line, "", "state.line != expected line")

  state.Handle("llllll")
  assert.Equal(state.line, "", "state.line != expected line")

  state.Handle("abcde")
  assert.Equal(state.line, "", "state.line != expected line")
}
