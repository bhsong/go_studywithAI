package main

import "fmt"

type Character struct {
	Name string
	// Level int
	HP int
}

func takeDamage(c *Character, damage int) {
	// c는 포인터지만 (*c).HP 라고 안 써도 됨
	// Go가 알아서 처리해줌
	c.HP = c.HP - damage
}

func main() {
	// 방법 1: 필드명 명시
	hero := Character{
		Name: "Gopher",
		//Level: 1,
		HP: 100,
	}

	// 데이터 접근 (점 . 사용)
	fmt.Println(hero.Name)
	fmt.Println("Before damage, HP:", hero.HP)
	takeDamage(&hero, 10)
	fmt.Println("After damage, HP:", hero.HP)
}
