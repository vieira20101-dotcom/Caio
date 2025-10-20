package main

import "fmt"

func main() {
	valorOriginal := 77

	fmt.Println("--- Variável Original ---")
	fmt.Println("Valor de 'valorOriginal':", valorOriginal)

	enderecoMemoria := &valorOriginal

	fmt.Printf("Endereço de Memória de 'valorOriginal': %p\n", enderecoMemoria)
	fmt.Printf("Tipo do endereço: %T\n", enderecoMemoria)

	// --- Desreferenciação (acessando o valor) ---

	valorAcessado := *enderecoMemoria

	fmt.Println("\n--- Acessando através do Ponteiro ---")
	fmt.Println("Valor no endereço:", valorAcessado)

	// --- Modificação via Ponteiro ---

	*enderecoMemoria = 99

	fmt.Println("\n--- Após Modificação via Ponteiro ---")
	fmt.Println("Novo Valor de 'valorOriginal':", valorOriginal)
	fmt.Println("Novo Valor no endereço:", *enderecoMemoria)

	fmt.Printf("Endereço de Memória (ainda o mesmo): %p\n", enderecoMemoria)
}
