package iltb

import (
	"fmt"
	"strings"

	"github.com/Henry-Sarabia/article"
	"github.com/pkg/errors"
)

type stage struct {
	base     string
	value    float64
	class    *class
	material *category
	content  *category
}

func appraise(s *stage) (float64, error) {
	return s.value * s.material.Multiplier, nil
}

func compose(s *stage) (string, error) {
	tok, err := tokenize(s.class.Format)
	if err != nil {
		return "", err
	}

	p, err := parse(s, tok)
	if err != nil {
		return "", err
	}

	return p, nil
}

func tokenize(format string) ([]string, error) {
	tok := strings.Fields(format)
	if tok[len(tok)-1] == "<article>" {
		return nil, errors.New("article token cannot be the last token in a format")
	}
	return tok, nil
}

func parse(s *stage, tok []string) (string, error) {
	for i := len(tok) - 1; i >= 0; i-- {
		switch tok[i] {
		case "<article>":
			tok[i] = article.Indefinite(tok[i+1])

		case "<material>":
			m, err := s.material.random()
			if err != nil {
				return "", errors.Wrap(err, "cannot retrieve random material type")
			}
			tok[i] = m

		case "<base>":
			tok[i] = s.base

		case "<verb>":
			v, err := s.class.randomVerb()
			if err != nil {
				return "", errors.Wrap(err, "cannot retrieve random verb")
			}
			tok[i] = v

		case "<content>":
			c, err := s.content.random()
			if err != nil {
				return "", errors.Wrap(err, "cannot retrieve random content type")
			}
			tok[i] = c

		default:
			return "", fmt.Errorf("unexpected token '%v' in format", tok[i])
		}
	}

	return strings.Join(tok, " "), nil
}
