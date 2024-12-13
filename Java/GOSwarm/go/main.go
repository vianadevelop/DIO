package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Função para executar comandos no terminal
func executeCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Função para obter o token de join do Docker Swarm
func getJoinToken() (string, error) {
	cmd := exec.Command("vagrant", "ssh", "master", "-c", "docker swarm join-token -q worker")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// Função para adicionar um nó ao cluster
func joinSwarm(nodeName, joinToken string) error {
	cmd := exec.Command("vagrant", "ssh", nodeName, "-c", fmt.Sprintf("docker swarm join --token %s 192.168.50.10:2377", joinToken))
	return cmd.Run()
}

func main() {
	// Passo 1: Subir as VMs com Vagrant
	fmt.Println("Subindo as máquinas virtuais...")
	err := executeCommand("vagrant up")
	if err != nil {
		fmt.Printf("Erro ao subir as máquinas: %v\n", err)
		return
	}

	// Passo 2: Obter o token do Docker Swarm
	fmt.Println("Obtendo o token de join para os nodes...")
	joinToken, err := getJoinToken()
	if err != nil {
		fmt.Printf("Erro ao obter o token: %v\n", err)
		return
	}

	// Passo 3: Adicionar os nodes ao Docker Swarm

	nodes := []string{"node01", "node02", "node03"}
	for _, node := range nodes {
		fmt.Printf("Adicionando %s ao cluster...\n", node)
		err = joinSwarm(node, joinToken)
		if err != nil {
			fmt.Printf("Erro ao adicionar %s ao cluster: %v\n", node, err)
			return
		}
	}

	// Passo 4: Verificar o estado do cluster
	fmt.Println("Verificando o estado do cluster...")
	err = executeCommand("vagrant ssh master -c 'docker node ls'")
	if err != nil {
		fmt.Printf("Erro ao verificar o estado do cluster: %v\n", err)
		return
	}

	fmt.Println("Cluster Docker Swarm configurado com sucesso!")
}
