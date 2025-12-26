package main

import (
	"fmt"
	"time"
)

func inputAction(bobchannel chan string) {
	var key string
	fmt.Scan(&key)
	bobchannel <- key
	close(bobchannel)
}

func main() {
	bobchannel := make(chan string)

	go inputAction(bobchannel)

	select {
	case action := <-bobchannel:
		if action == "code123" {
			fmt.Println("폭탄 해체 성공!")
		} else {
			fmt.Println("비밀번호 오류! 폭탄이 터집니다!")
		}
	case <-time.After(5 * time.Second):
		fmt.Println("시간 초과! 쾅!!")
	}
}
