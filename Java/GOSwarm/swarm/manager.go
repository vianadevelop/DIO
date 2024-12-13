package swarm

import (
	"fmt"
	"goswarm/go/utils"
)

func SetupManager() error {
	fmt.Println("Configurando o nó manager...")

	// Comando para inicializar o Swarm no nó manager
	err := utils.RunCommand("docker swarm init")
	if err != nil {
		return fmt.Errorf("erro ao configurar o nó manager: %v", err)
	}
	return nil
}
