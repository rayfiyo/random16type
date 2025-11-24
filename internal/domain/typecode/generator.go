package typecode

// Code は4文字のタイプコード（例: INTJ）を表す。
type Code string

// Generator は Code を生成する。
type Generator interface {
	Generate() Code
}
