package helpers

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// SafeOrderBy evita SQL injection em ORDER BY ao aceitar apenas colunas em whitelist.
func SafeOrderBy(requested string, allowed map[string]string, fallback string) string {
	key := strings.ToLower(strings.TrimSpace(requested))
	if safeColumn, ok := allowed[key]; ok {
		return safeColumn
	}
	return fallback
}

// SafeSortDirection evita SQL injection na direção do ORDER BY.
func SafeSortDirection(requested, fallback string) string {
	direction := strings.ToUpper(strings.TrimSpace(requested))
	if direction == "ASC" || direction == "DESC" {
		return direction
	}
	return fallback
}

// EscapeLikeTerm evita curingas livres em consultas ILIKE e reduz abuso de busca.
func EscapeLikeTerm(value string) string {
	replacer := strings.NewReplacer(
		"\\", "\\\\",
		"%", "\\%",
		"_", "\\_",
	)
	return replacer.Replace(strings.TrimSpace(value))
}

func HashPassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
