package funcs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tarefa struct {
	Nome      string
	Concluida bool
}

var tarefas []Tarefa
var reader = bufio.NewReader(os.Stdin)

func ListarOpcoes() {
	fmt.Println("Bem vindo ao Gerenciador de Tarefas!\nEscolha uma opção:\n" +
		"1. Adicionar Tarefa\n" +
		"2. Listar Tarefas\n" +
		"3. Concluir Tarefa\n" +
		"4. Remover Tarefa\n" +
		"5. Sair")
}

func AdicionarTarefa() {
	fmt.Print("Digite o nome da tarefa: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)
	fmt.Printf("Tarefa '%s' adicionada com sucesso!\n", nome)
	tarefa := Tarefa{Nome: nome, Concluida: false}
	tarefas = append(tarefas, tarefa)
}

func ListarTarefas() {
	if len(tarefas) == 0 {
		fmt.Println("Nenhuma tarefa cadastrada.")
		return
	}

	fmt.Println("Tarefas:")
	for i, tarefa := range tarefas {
		fmt.Printf("%d. ", i+1)
		if tarefa.Concluida {
			fmt.Println("[X] " + tarefa.Nome)
		} else {
			fmt.Println("[ ] " + tarefa.Nome)
		}
	}
}

func ConcluirTarefa() error {
	fmt.Print("Digite o nome da tarefa a ser concluída: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	for i, tarefa := range tarefas {
		if tarefa.Nome == nome {
			tarefas[i].Concluida = true
			fmt.Printf("Tarefa '%s' concluída com sucesso!\n", nome)
			return nil
		}
	}
	err := fmt.Errorf("tarefa '%s' não encontrada", nome)
	fmt.Println(err)
	return err
}

func RemoverTarefa() error {
	fmt.Print("Digite o nome da tarefa a ser removida: ")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	for i, tarefa := range tarefas {
		if tarefa.Nome == nome {
			indiceARemover := i

			// Forma simples de remover sem criar um novo slice
			// parte1 := tarefas[:indiceARemover]
			// parte2 := tarefas[indiceARemover+1:]
			// tarefasAtt := append(parte1, parte2...)

			tarefas = append(tarefas[:indiceARemover], tarefas[indiceARemover+1:]...)
			fmt.Printf("Tarefa '%s' removida com sucesso!\n", nome)
			return nil
		}
	}

	err := fmt.Errorf("tarefa '%s' não encontrada", nome)
	fmt.Println(err)
	return err
}

func Sair() {
	fmt.Println("Saindo do gerenciador de tarefas. Até mais!")
}
