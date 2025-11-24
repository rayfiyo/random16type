package random

import (
	"math/rand"
	"time"

	"github.com/rayfiyo/random16type/internal/domain/typecode"
)

// Generator produces random type codes using math/rand.
type Generator struct {
	r *rand.Rand
}

func NewGenerator(seed int64) *Generator {
	return &Generator{r: rand.New(rand.NewSource(seed))}
}

func NewGeneratorFromTime() *Generator {
	return NewGenerator(time.Now().UnixNano())
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
		code[i] = p[g.r.Intn(len(p))]
	}
	return typecode.Code(string(code))
}
