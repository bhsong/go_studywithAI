package main

import "fmt"

type Character struct {
	Name string
	// Level int
	HP int
}

// func와 함수 이름 사이에 (c *Character) 추가
// 뜻: "이 함수는 *Character 타입에 속한 기능이다"
func (c *Character) takeDamage(damage int) {
	// c는 포인터지만 (*c).HP 라고 안 써도 됨
	// Go가 알아서 처리해줌
	c.HP = c.HP - damage
}

func (c *Character) showInfo() {
	fmt.Printf("Name: %s, HP: %d\n", c.Name, c.HP)
}

func main() {
	// 방법 1: 필드명 명시
	hero := Character{
		Name: "Gopher",
		//Level: 1,
		HP: 100,
	}

	// 메소드(리시버) 사용 (점 . 사용)
	hero.showInfo()
	hero.takeDamage(10)
	hero.showInfo()
}
