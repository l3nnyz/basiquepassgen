package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cfg := parseFlags()

	if err := cfg.validate(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur:", err)
		os.Exit(1)
	}

	gen := NewGenerator(cfg)

	for i := 0; i < cfg.count; i++ {
		password, err := gen.Generate(cfg.length)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Erreur lors de la génération:", err)
			os.Exit(1)
		}

		if cfg.count > 1 {
			fmt.Printf("%2d. %s\n", i+1, password)
		} else {
			fmt.Println(password)
		}
	}

	showStrength(cfg)
}

func parseFlags() *config {
	cfg := &config{}

	flag.IntVar(&cfg.length, "length", defaultLength, "Longueur du mot de passe")
	flag.IntVar(&cfg.length, "l", defaultLength, "Longueur du mot de passe (raccourci)")
	flag.IntVar(&cfg.count, "count", defaultCount, "Nombre de mots de passe à générer")
	flag.IntVar(&cfg.count, "n", defaultCount, "Nombre de mots de passe à générer (raccourci)")
	flag.BoolVar(&cfg.upper, "upper", true, "Inclure des lettres majuscules")
	flag.BoolVar(&cfg.upper, "u", true, "Inclure des lettres majuscules (raccourci)")
	flag.BoolVar(&cfg.numbers, "numbers", true, "Inclure des chiffres")
	flag.BoolVar(&cfg.numbers, "d", true, "Inclure des chiffres (raccourci)")
	flag.BoolVar(&cfg.symbols, "symbols", false, "Inclure des symboles")
	flag.BoolVar(&cfg.symbols, "s", false, "Inclure des symboles (raccourci)")
	flag.StringVar(&cfg.exclude, "exclude", "", "Caractères à exclure")
	flag.StringVar(&cfg.exclude, "e", "", "Caractères à exclure (raccourci)")
	flag.BoolVar(&cfg.noClip, "no-clip", true, "Désactiver la copie dans le presse-papiers (réservé pour usage futur)")

	flag.Usage = customUsage
	flag.Parse()

	return cfg
}

func customUsage() {
	fmt.Fprintln(os.Stderr, "passgen - Générateur de mots de passe sécurisés")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Usage: passgen [options]")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Exemples:")
	fmt.Fprintln(os.Stderr, "  passgen -l 20 -s              Génère un mot de passe de 20 caractères avec symboles")
	fmt.Fprintln(os.Stderr, "  passgen -n 5 -d=true -s=true  Génère 5 mots de passe avec chiffres et symboles")
	fmt.Fprintln(os.Stderr, "  passgen -l 32 -u=false        Génère un mot de passe de 32 chars sans majuscules")
	fmt.Fprintln(os.Stderr, "  passgen -e \"oO0lI1|\"           Exclut les caractères ambigus o, O, 0, l, I, 1, |")
}
