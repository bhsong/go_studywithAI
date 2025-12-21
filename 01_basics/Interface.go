package main

import "fmt"

// Attack 메서드를 가진 녀석은 무조건 Attacker로 쳐주겠다는 선언
type Attacker interface {
	Attack()
}

type Player struct{}

// Player가 Attack 메서드를 가짐 -> 자동으로 Attacker 인터페이스 조건 충족!
func (p *Player) Attack() {
	fmt.Println("플레이어가 검을 휘두릅니다!")
}

type Monster struct{}

// Monster가 Attack 메서드를 가짐 -> 자동으로 Attacker 인터페이스 조건 충족!
func (m *Monster) Attack() {
	fmt.Println("몬스터가 포효하며 공격합니다!")
}

func DoAttack(a Attacker) {
	a.Attack()
}

func main() {
	player := &Player{}
	monster := &Monster{}

	DoAttack(player)
	DoAttack(monster)
}
