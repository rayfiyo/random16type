package typecode

import "github.com/rayfiyo/random16type/internal/domain/typecode"

// ドメインのジェネレータをアプリケーション層から利用するためのラッパー
type GeneratorUsecase struct {
	generator typecode.Generator
}

func NewGeneratorUsecase(generator typecode.Generator) *GeneratorUsecase {
	return &GeneratorUsecase{generator: generator}
}

func (u *GeneratorUsecase) Generate() typecode.Code {
	return u.generator.Generate()
}
