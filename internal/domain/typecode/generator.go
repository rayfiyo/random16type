package typecode

// Code represents a 4-letter type string (e.g. INTJ).
type Code string

// Generator produces a Code.
type Generator interface {
	Generate() Code
}
