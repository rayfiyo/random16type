package typecode

import (
	"time"

	"github.com/rayfiyo/random16type/internal/domain/typecode"
)

// Cache は日付文字列をキーにコードを保持する。
type Cache interface {
	Get(key string) (typecode.Code, bool)
	Set(key string, code typecode.Code)
}

// ドメインのジェネレータをアプリケーション層から利用するためのラッパー。
// 同じ日付ではキャッシュ済みの結果を返し、日付が変わると生成し直す。
type GeneratorUsecase struct {
	generator typecode.Generator
	cache     Cache
	now       func() time.Time
}

func NewGeneratorUsecase(generator typecode.Generator, cache Cache) *GeneratorUsecase {
	return &GeneratorUsecase{
		generator: generator,
		cache:     cache,
		now:       time.Now,
	}
}

func (u *GeneratorUsecase) Generate() typecode.Code {
	today := u.now().In(time.Local).Format("2006-01-02")
	if cached, ok := u.cache.Get(today); ok {
		return cached
	}

	code := u.generator.Generate()
	u.cache.Set(today, code)
	return code
}
