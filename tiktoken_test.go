package tiktoken

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
    ass := assert.New(t)
    enc, err := EncodingForModel("gpt-3.5-turbo-16k")
    ass.Nil(err, "Encoding  init should not be nil")
    tokens := enc.Encode("hello world!你好，世界！", []string{"all"}, []string{"all"})
    // these tokens are converted from the original python code
    sourceTokens := []int{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
    ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

    tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, nil)
    sourceTokens = []int{15339, 220, 100257}
    ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

    tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, []string{"all"})
    sourceTokens = []int{15339, 220, 100257}
    ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

    ass.Panics(func() {
        tokens = enc.Encode("hello <|endoftext|><|endofprompt|>", []string{"<|endoftext|>"}, []string{"all"})
    })
    ass.Panics(func() {
        tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, []string{"<|endoftext|>"})
    })
}

func TestQwenEncoding(t *testing.T) {
    ass := assert.New(t)
    enc, err := EncodingForModel("qwen")
    ass.Nil(err, "Encoding  init should not be nil")
    tokens := enc.Encode("hello world!你好，世界！", []string{"all"}, []string{"all"})
    ass.Equal(7, len(tokens))
}

func TestDecoding(t *testing.T) {
    ass := assert.New(t)
    // enc, err := GetEncoding("cl100k_base")
    enc, err := GetEncoding(MODEL_CL100K_BASE)
    enc2, err2 := GetEncoding(MODEL_QWEN_BASE)
    ass.Nil(err2, "Encoding  init should not be nil")
    ass.Nil(err, "Encoding  init should not be nil")
    sourceTokens := []int{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
    sourceTokens2 := []int{14990, 1879, 0, 108386, 3837, 99489, 6313}
    txt := enc.Decode(sourceTokens)
    txt2 := enc2.Decode(sourceTokens2)
    ass.Equal("hello world!你好，世界！", txt, "Decoding should be equal")
    ass.Equal("hello world!你好，世界！", txt2, "Decoding should be equal")
}
