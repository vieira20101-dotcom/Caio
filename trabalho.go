package main

import (
	"errors"
	"fmt"
)

// --- Interfaces ---

// Interface básica
type Animal interface {
	Falar() string
	Identificar() string
}

// Interface adicional
type Vivo interface {
	EstaVivo() bool
}

// Interface composta
type SerVivo interface {
	Animal
	Vivo
}

// --- Struct Pessoa ---

type Pessoa struct {
	Nome  string
	Idade int
	Vivo  bool
}

func (p Pessoa) Falar() string {
	return "Oi, meu nome é " + p.Nome
}

func (p *Pessoa) MudarNome(novoNome string) {
	p.Nome = novoNome
}

func (p Pessoa) Identificar() string {
	return fmt.Sprintf("Pessoa: %s, %d anos", p.Nome, p.Idade)
}

func (p Pessoa) EstaVivo() bool {
	return p.Vivo
}

func (p Pessoa) VerificarMaioridade() error {
	if p.Idade < 18 {
		return errors.New("pessoa é menor de idade")
	}
	return nil
}

// --- Struct Cachorro ---

type Cachorro struct {
	Nome  string
	Idade int
}

func (c Cachorro) Falar() string {
	return "Au au! Eu sou " + c.Nome
}

func (c Cachorro) Identificar() string {
	return fmt.Sprintf("Cachorro: %s, %d anos", c.Nome, c.Idade)
}

func (c Cachorro) EstaVivo() bool {
	return true
}

// --- Função que usa interfaces ---

func Apresentar(s SerVivo) {
	fmt.Println(s.Falar())
	fmt.Println(s.Identificar())
	fmt.Println("Está vivo?", s.EstaVivo())
	fmt.Println("---")
}

// --- Função principal ---

func main() {
	// Pessoa
	p := Pessoa{Nome: "João", Idade: 17, Vivo: true}

	fmt.Println("Endereço de memória de Pessoa:", &p)
	err := p.VerificarMaioridade()
	if err != nil {
		fmt.Println("Erro:", err)
	}

	// Alterando nome via ponteiro
	p.MudarNome("Carlos")
	fmt.Println("Nome atualizado:", p.Nome)

	// Cachorro
	c := Cachorro{Nome: "Rex", Idade: 5}

	// Usando interface composta com ponteiros
	var seres []SerVivo
	seres = append(seres, &p, &c)

	// Apresentar todos
	for _, s := range seres {
		Apresentar(s)
	}

	// Mapa de SerVivo com ponteiros
	seresMap := map[string]SerVivo{
		"pessoa":   &p,
		"cachorro": &c,
	}

	fmt.Println("🔍 Buscando no mapa:")
	for chave, valor := range seresMap {
		fmt.Printf("[%s] -> %s\n", chave, valor.Identificar())
	}
}
