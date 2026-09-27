package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	lowerCharset  = "abcdefghijklmnopqrstuvwxyz"
	upperCharset  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberCharset = "0123456789"
	symbolCharset = "!@#$%^&*()-_=+[]{};:,.<>?/~"
)

type Generator struct {
	Charset       string
	RequiredSets  []string
	Exclude       string
}

func NewGenerator(c *config) *Generator {
	g := &Generator{
		Exclude: c.exclude,
	}
	g.Charset = buildCharset(c)
	g.RequiredSets = ensureAtLeastOneOfEach(c)
	return g
}

func (g *Generator) Generate(length int) (string, error) {
	password := make([]byte, length)

	for i := 0; i < len(g.RequiredSets) && i < length; i++ {
		ch, err := randomChar(g.RequiredSets[i])
		if err != nil {
			return "", err
		}
		password[i] = ch
	}

	for i := len(g.RequiredSets); i < length; i++ {
		ch, err := randomChar(g.Charset)
		if err != nil {
			return "", err
		}
		password[i] = ch
	}

	if err := shuffle(password); err != nil {
		return "", err
	}

	return string(password), nil
}

func buildCharset(c *config) string {
	charset := lowerCharset
	if c.upper {
		charset += upperCharset
	}
	if c.numbers {
		charset += numberCharset
	}
	if c.symbols {
		charset += symbolCharset
	}

	if len(c.exclude) > 0 {
		for _, ch := range c.exclude {
			charset = strings.ReplaceAll(charset, string(ch), "")
		}
	}

	return charset
}

func ensureAtLeastOneOfEach(c *config) []string {
	var required []string
	lower := lowerCharset
	upper := upperCharset
	numbers := numberCharset
	symbols := symbolCharset

	if len(c.exclude) > 0 {
		for _, ch := range c.exclude {
			lower = strings.ReplaceAll(lower, string(ch), "")
			upper = strings.ReplaceAll(upper, string(ch), "")
			numbers = strings.ReplaceAll(numbers, string(ch), "")
			symbols = strings.ReplaceAll(symbols, string(ch), "")
		}
	}

	if len(lower) > 0 {
		required = append(required, lower)
	}
	if c.upper && len(upper) > 0 {
		required = append(required, upper)
	}
	if c.numbers && len(numbers) > 0 {
		required = append(required, numbers)
	}
	if c.symbols && len(symbols) > 0 {
		required = append(required, symbols)
	}

	return required
}

func randomChar(charset string) (byte, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charset[n.Int64()], nil
}

func shuffle(data []byte) error {
	n := len(data)
	for i := n - 1; i > 0; i-- {
		max := big.NewInt(int64(i + 1))
		jBig, err := rand.Int(rand.Reader, max)
		if err != nil {
			return err
		}
		j := int(jBig.Int64())
		data[i], data[j] = data[j], data[i]
	}
	return nil
}

func charsetSize(c *config) int {
	size := len(lowerCharset)
	if c.upper {
		size += len(upperCharset)
	}
	if c.numbers {
		size += len(numberCharset)
	}
	if c.symbols {
		size += len(symbolCharset)
	}
	if len(c.exclude) > 0 {
		size -= len(c.exclude)
	}
	return size
}
