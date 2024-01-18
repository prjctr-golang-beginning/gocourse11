package main

import "fmt"

// ToothCondition - тип для опису стану зуба
type ToothCondition struct {
	DecayLevel int
}

// FillingMixture - суміш для пломбування
type FillingMixture struct {
	Material string
	Hardness int
}

// FillingMixtureBuilder - інтерфейс будівельника
type FillingMixtureBuilder interface {
	SetMaterial(material string) FillingMixtureBuilder
	SetHardness(hardness int) FillingMixtureBuilder
	Build() FillingMixture
}

// ConcreteMixtureBuilder - конкретний будівельник суміші
type ConcreteMixtureBuilder struct {
	material string
	hardness int
}

func (b *ConcreteMixtureBuilder) SetMaterial(material string) FillingMixtureBuilder {
	b.material = material
	return b
}

func (b *ConcreteMixtureBuilder) SetHardness(hardness int) FillingMixtureBuilder {
	b.hardness = hardness
	return b
}

func (b *ConcreteMixtureBuilder) Build() FillingMixture {
	return FillingMixture{Material: b.material, Hardness: b.hardness}
}

// FillingStation - станція для приготування суміші (стара версія)
type FillingStation struct{}

func (station *FillingStation) PrepareMixture(condition ToothCondition) FillingMixture {
	// Логіка створення суміші для старої версії станції
	builder := &ConcreteMixtureBuilder{}
	return builder.SetMaterial("Standard").SetHardness(condition.DecayLevel).Build()
}

// NewFillingStation - нова версія станції
type NewFillingStation struct{}

func (station *NewFillingStation) NewPrepareMixture(material string, hardness int) FillingMixture {
	// Нова логіка створення суміші
	return FillingMixture{Material: material, Hardness: hardness}
}

// FillingStationAdapter - адаптер для нової версії станції
type FillingStationAdapter struct {
	NewStation *NewFillingStation
}

func (adapter *FillingStationAdapter) PrepareMixture(condition ToothCondition) FillingMixture {
	// Адаптація до інтерфейсу старої станції
	material := "Enhanced"
	hardness := condition.DecayLevel
	return adapter.NewStation.NewPrepareMixture(material, hardness)
}

func main() {
	condition := ToothCondition{DecayLevel: 3}

	// Створення суміші за допомогою старої станції
	oldStation := &FillingStation{}
	mixture1 := oldStation.PrepareMixture(condition)
	fmt.Printf("Old Station Mixture: %+v\n", mixture1)

	// Створення суміші за допомогою нової станції з адаптером
	newStation := &NewFillingStation{}
	adapter := &FillingStationAdapter{NewStation: newStation}
	mixture2 := adapter.PrepareMixture(condition)
	fmt.Printf("New Station Mixture (via Adapter): %+v\n", mixture2)
}
