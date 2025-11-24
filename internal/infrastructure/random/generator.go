package random

import (
	"math/rand/v2"

	"github.com/rayfiyo/random16type/internal/domain/typecode"
)

// math/rand/v2 で typecode をランダム生成
type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate() typecode.Code {
	pairs := [][]rune{
		{'I', 'E'},
		{'N', 'S'},
		{'T', 'F'},
		{'J', 'P'},
	}

	code := make([]rune, len(pairs))
	for i, p := range pairs {
		code[i] = p[rand.IntN(len(p))]
	}
	return typecode.Code(string(code))
}
