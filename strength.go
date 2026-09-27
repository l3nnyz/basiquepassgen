package main

import (
	"fmt"
	"os"
)

const (
	defaultLength = 16
	defaultCount  = 1
	minLength     = 4
)

type config struct {
	length  int
	count   int
	upper   bool
	numbers bool
	symbols bool
	exclude string
	noClip  bool
}

func (c *config) validate() error {
	if c.length < minLength {
		return fmt.Errorf("la longueur doit être au minimum de %d caractères", minLength)
	}

	if len(c.exclude) >= charsetSize(c) {
		return fmt.Errorf("trop de caractères exclus, il n'en reste pas assez pour générer le mot de passe")
	}

	return nil
}

func showStrength(c *config) {
	size := charsetSize(c)
	if size < 1 {
		return
	}

	entropy := float64(c.length) * log2(float64(size))

	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintf(os.Stderr, "Force du mot de passe: %s\n", strengthLabel(entropy))
}

func log2(x float64) float64 {
	n := 0.0
	for x > 1 {
		x /= 2
		n++
	}
	return n
}

func strengthLabel(entropy float64) string {
	switch {
	case entropy < 28:
		return "TRES FAIBLE"
	case entropy < 36:
		return "FAIBLE"
	case entropy < 60:
		return "MOYENNE"
	case entropy < 80:
		return "FORTE"
	default:
		return "TRES FORTE"
	}
}
