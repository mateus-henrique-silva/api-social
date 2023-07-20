package slug

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

func RemoveAcentoEspacoCaracterEspecial(s string) string {
	// Remove acentos.
	t := transform(strings.ToLower(s), removeAccent)

	// Remove espaços.
	t = strings.ReplaceAll(t, " ", "")

	// Remove caracteres especiais.
	reg := regexp.MustCompile("[^a-z0-9]+")
	t = reg.ReplaceAllString(t, "")

	return t
}

func transform(s string, f func(r rune) rune) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, c := range s {
		if c >= utf8.RuneSelf {
			if c = f(c); c < utf8.RuneSelf {
				b.WriteRune(c)
			}
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func removeAccent(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r
	}

	var accentMap = map[rune]rune{
		'à': 'a', 'á': 'a', 'â': 'a', 'ã': 'a', 'ä': 'a', 'å': 'a',
		'è': 'e', 'é': 'e', 'ê': 'e', 'ë': 'e',
		'ì': 'i', 'í': 'i', 'î': 'i', 'ï': 'i',
		'ò': 'o', 'ó': 'o', 'ô': 'o', 'õ': 'o', 'ö': 'o', 'ø': 'o',
		'ù': 'u', 'ú': 'u', 'û': 'u', 'ü': 'u',
		'ç': 'c',
		'ñ': 'n',
	}

	if replacement, ok := accentMap[unicode.ToLower(r)]; ok {
		return replacement
	}
	return r
}
