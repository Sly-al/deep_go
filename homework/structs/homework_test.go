package main

import (
	"math"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		person.name = name
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x = int32(x)
		person.y = int32(y)
		person.z = int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.mana = uint16(mana)
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.health = uint16(health)
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.respect = uint8(respect)
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.strength = uint8(strength)
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.exp = uint8(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.lvl = uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.mask |= House
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.mask |= Weapon
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.mask |= Famaly
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.pType = personType
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

const (
	House = 1 << iota
	Weapon
	Famaly
)

type GamePerson struct {
	name                              string
	pType                             int
	x, y, z                           int32
	gold                              uint32
	mana, health                      uint16
	respect, strength, exp, lvl, mask uint8
}

func NewGamePerson(options ...Option) GamePerson {
	person := &GamePerson{}

	for _, opt := range options {
		opt(person)
	}

	return *person
}

func (p *GamePerson) Name() string {
	return p.name
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.mana)
}

func (p *GamePerson) Health() int {
	return int(p.health)
}

func (p *GamePerson) Respect() int {
	return int(p.respect)
}

func (p *GamePerson) Strength() int {
	return int(p.strength)
}

func (p *GamePerson) Experience() int {
	return int(p.exp)
}

func (p *GamePerson) Level() int {
	return int(p.lvl)
}

func (p *GamePerson) HasHouse() bool {
	return p.mask&House != 0
}

func (p *GamePerson) HasGun() bool {
	return p.mask&Weapon != 0
}

func (p *GamePerson) HasFamilty() bool {
	return p.mask&Famaly != 0
}

func (p *GamePerson) Type() int {
	return p.pType
}

func TestGamePerson(t *testing.T) {
	assert.LessOrEqual(t, unsafe.Sizeof(GamePerson{}), uintptr(64))

	const x, y, z = math.MinInt32, math.MaxInt32, 0
	const name = "aaaaaaaaaaaaa_bbbbbbbbbbbbb_cccccccccccccc"
	const personType = BuilderGamePersonType
	const gold = math.MaxInt32
	const mana = 1000
	const health = 1000
	const respect = 10
	const strength = 10
	const experience = 10
	const level = 10

	options := []Option{
		WithName(name),
		WithCoordinates(x, y, z),
		WithGold(gold),
		WithMana(mana),
		WithHealth(health),
		WithRespect(respect),
		WithStrength(strength),
		WithExperience(experience),
		WithLevel(level),
		WithHouse(),
		WithFamily(),
		WithType(personType),
	}

	person := NewGamePerson(options...)
	assert.Equal(t, name, person.Name())
	assert.Equal(t, x, person.X())
	assert.Equal(t, y, person.Y())
	assert.Equal(t, z, person.Z())
	assert.Equal(t, gold, person.Gold())
	assert.Equal(t, mana, person.Mana())
	assert.Equal(t, health, person.Health())
	assert.Equal(t, respect, person.Respect())
	assert.Equal(t, strength, person.Strength())
	assert.Equal(t, experience, person.Experience())
	assert.Equal(t, level, person.Level())
	assert.True(t, person.HasHouse())
	assert.True(t, person.HasFamilty())
	assert.False(t, person.HasGun())
	assert.Equal(t, personType, person.Type())
}
