# passgen

Générateur de mots de passe sécurisés écrit en Go.

## Installation

```bash
go build -o passgen .
```

## Utilisation

```bash
passgen [options]
```

### Options

| Option | Raccourci | Description | Défaut |
|---|---|---|---|
| `--length <n>` | `-l <n>` | Longueur du mot de passe | 16 |
| `--count <n>` | `-n <n>` | Nombre de mots de passe à générer | 1 |
| `--upper` | `-u` | Inclure des lettres majuscules | true |
| `--numbers` | `-d` | Inclure des chiffres | true |
| `--symbols` | `-s` | Inclure des symboles | false |
| `--exclude <chars>` | `-e <chars>` | Caractères à exclure | (aucun) |

### Exemples

Mot de passe simple de 16 caractères (défaut) :
```bash
passgen
```

Mot de passe de 20 caractères avec symboles :
```bash
passgen -l 20 -s
```

5 mots de passe de 24 caractères avec chiffres et symboles :
```bash
passgen -n 5 -d=true -s=true -l 24
```

Mot de passe de 32 caractères sans majuscules :
```bash
passgen -l 32 -u=false
```

Exclure les caractères ambigus (o, O, 0, l, I, 1, |) :
```bash
passgen -e "oO0lI1|"
```

## Structure du projet

```
.
├── main.go        # CLI et flags
├── generator.go   # Logique de génération (crypto/rand)
├── strength.go    # Validation & indicateur de force
└── go.mod
```
