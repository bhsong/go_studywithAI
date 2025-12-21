package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Combatant interface {
	Attack(opponent Combatant)
	TakeDamage(amount int)
	IsAlive() bool
	GetName() string
}

type Player struct {
	Name string
	HP   int
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) IsAlive() bool {
	return p.HP > 0 // 조건식의 결과 자체가 bool
}

func (p *Player) TakeDamage(damage int) {
	p.HP -= damage
	fmt.Println(p.Name, "takes", damage, "damage! Remaining HP:", p.HP)
}

func (p *Player) Attack(opponent Combatant) {
	damage := rand.Intn(11) + 10
	opponent.TakeDamage(damage)
	fmt.Println(p.Name, "attacks", opponent.GetName(), "for", damage, "damage!")
}

type Monster struct {
	Name string
	HP   int
}

func (m *Monster) GetName() string {
	return m.Name
}

func (m *Monster) IsAlive() bool {
	return m.HP > 0 // 조건식의 결과 자체가 bool
}

func (m *Monster) TakeDamage(damage int) {
	m.HP -= damage
	fmt.Println(m.Name, "takes", damage, "damage! Remaining HP:", m.HP)
}

func (m *Monster) Attack(opponent Combatant) {
	damage := rand.Intn(11) + 5
	opponent.TakeDamage(damage)
	fmt.Println(m.Name, "attacks", opponent.GetName(), "for", damage, "damage!")
}

func main() {
	// 랜덤 시드 설정 (매번 다른 랜덤값이 나오게 함)
	rand.Seed(time.Now().UnixNano())
	// 플레이어와 몬스터 생성
	var player Combatant = &Player{Name: "Hero", HP: 100}
	var monster Combatant = &Monster{Name: "Goblin", HP: 80}

	// 전투 시작
	fmt.Println("=== 전투 시작! ===")
	for {
		// 1. 플레이어 턴
		fmt.Print("\n[1] 공격 [2] 도망\n")
		var input int
		fmt.Scan(&input)

		if input == 1 {
			player.Attack(monster)
		} else {
			fmt.Println(player.GetName(), "도망쳤다!")
			break
		}

		if !monster.IsAlive() {
			fmt.Println(monster.GetName(), "가(이) 쓰러졌다! 승리!")
			break
		}

		// 2. 몬스터 턴
		monster.Attack(player)

		if !player.IsAlive() {
			fmt.Println(player.GetName(), "가(이) 쓰러졌다! 패배!")
			break
		}
	}
}
