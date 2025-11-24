package httpserver

import (
	"fmt"
	"net/http"

	"github.com/rayfiyo/random16type/internal/domain/typecode"
)

// ユースケースの Generate を提供するインタフェース
type Generator interface {
	Generate() typecode.Code
}

type PageHandler struct {
	generator Generator
}

func NewPageHandler(generator Generator) *PageHandler {
	return &PageHandler{generator: generator}
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	code := h.generator.Generate()

	page := fmt.Sprintf(pageTemplate, code)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(page))
}

const pageTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Random 16 Type</title>
  <style>
    body {
      margin: 0;
      height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      font-family: "Helvetica Neue", Arial, sans-serif;
      background: linear-gradient(135deg, #f5f7fa, #c3cfe2);
      color: #222;
    }
     .card {
       background: #fff;
       padding: 32px 48px;
       border-radius: 16px;
       box-shadow: 0 12px 30px rgba(0, 0, 0, 0.08);
       text-align: center;
     }
     h1 {
       margin: 12px 0 12px;
       letter-spacing: 0.08em;
       font-size: 48px;
       font-family:monospace;
     }
     p {
       margin: 12px 0;
       color: #555;
     }
  </style>
</head>
<body>
  <div class="card">
    <p>Today's personality</p>
    <h1>%s</h1>
  </div>
</body>
</html>`
