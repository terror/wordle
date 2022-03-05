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

  assert.Nil(state.Handle("crane"))
  assert.Equal(state.Line(), "🟨⬛⬛⬛⬛", "should be equal")

  assert.Nil(state.Handle("bloke"))
  assert.Equal(state.Line(), "🟩🟩🟩🟨⬛", "should be equal")

  assert.Nil(state.Handle("block"))
  assert.Equal(state.Line(), "🟩🟩🟩🟩🟩", "should be equal")
}

func TestDuplicates(t *testing.T) {
  assert := assert.New(t)

  state := Init()
  state.word = "weeds"

  assert.Nil(state.Handle("lakes"))
  assert.Equal(state.Line(), "⬛⬛⬛🟨🟩", "should be equal")

  assert.Nil(state.Handle("zowee"))
  assert.Equal(state.Line(), "⬛⬛🟨🟨🟨", "should be equal")

  assert.Nil(state.Handle("wheel"))
  assert.Equal(state.Line(), "🟩⬛🟩🟨⬛", "should be equal")
}

func TestInvalid(t *testing.T) {
  assert := assert.New(t)

  state := Init()

  assert.NotNil(state.Handle("l"))
  assert.NotNil(state.Handle("llllll"))
  assert.NotNil(state.Handle("abcde"))
}
