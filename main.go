package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomType() string {
	pairs := [][]rune{
		{'I', 'E'},
		{'N', 'S'},
		{'T', 'F'},
		{'J', 'P'},
	}

	code := make([]rune, len(pairs))
	for i, p := range pairs {
		code[i] = p[r.Intn(len(p))]
	}
	return string(code)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	value := randomType()

	page := fmt.Sprintf(`<!DOCTYPE html>
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
</html>`, value)

	_, err := w.Write([]byte(page))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	log.Printf("serving random16type at http://localhost%s/", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
