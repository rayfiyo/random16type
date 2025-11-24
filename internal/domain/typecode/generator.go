package typecode

// 4文字のタイプコード（例: INTJ）を表現
type Code string

// Code を生成
type Generator interface {
	Generate() Code
}
