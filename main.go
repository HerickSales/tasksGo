package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// ListarOpcoes()
	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		opcao, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrada inválida! Digite um valor inteiro entre 1 e 5.")
			continue
		}

		if opcao < 1 || opcao > 5 {
			fmt.Println("Opção inválida! Digite um número entre 1 e 5.")
			continue
		}

		// switch opcao {
		// case 1:
		// 	AdicionarTarefa()
		// case 2:
		// 	ListarTarefas()
		// case 3:
		// 	ConcluirTarefa()
		// case 4:
		// 	RemoverTarefa()
		// case 5:
		// 	Sair()
		// 	return
		// }
	}
}
