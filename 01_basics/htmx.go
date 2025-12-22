package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 1. 메인 페이지 핸들러 (처음 접속 시)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 편의상 HTML을 문자열로 넣었습니다. 실제로는 파일로 분리합니다.
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go & HTMX System Monitor</title>
			<script src="https://unpkg.com/htmx.org@1.9.10"></script>
			<style>
				body { font-family: monospace; padding: 20px; background: #222; color: #eee; }
				.box { border: 1px solid #444; padding: 20px; margin-top: 10px; }
				button { cursor: pointer; padding: 10px; background: #007bff; color: white; border: none; }
			</style>
		</head>
		<body>
			<h2>Simple System Monitor</h2>
			
			<button hx-get="/update-time" hx-target="#server-output" hx-swap="innerHTML">
				Check Server Time
			</button>

			<div id="server-output" class="box">
				Waiting for server response...
			</div>
		</body>
		</html>
		`
		w.Write([]byte(html))
	})

	// 2. 데이터 처리 핸들러 (AJAX 요청 처리)
	// JSON이 아니라 'HTML 조각'을 리턴하는 것이 포인트입니다.
	http.HandleFunc("/update-time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC1123)

		// 복잡한 JSON 파싱 없이, 화면에 찍을 HTML 문자열을 그대로 포맷팅해서 던집니다.
		// C언어의 printf와 개념이 똑같습니다.
		fmt.Fprintf(w, "<p>Server Time: <strong style='color:#0f0'>%s</strong></p>", currentTime)

		// 리눅스 개발자시니, 터미널 로그로 요청 확인
		fmt.Println("Request received, sending HTML fragment...")
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
