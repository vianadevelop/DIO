package swarm

import (
	"fmt"
	"goswarm/go/utils"
)

func SetupNode(nodeName string) error {
	fmt.Printf("Configurando o nó worker: %s...\n", nodeName)

	// Comando para adicionar o nó ao Swarm
	err := utils.RunCommand(fmt.Sprintf("docker swarm join --token <join-token> <master-ip>:2377"))
	if err != nil {
		return fmt.Errorf("erro ao configurar o nó worker %s: %v", nodeName, err)
	}
	return nil
}
