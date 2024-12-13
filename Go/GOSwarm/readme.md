# GOSwarm - Cluster Docker Swarm com Vagrant e Go 🚀🐳

**GOSwarm** é um projeto que utiliza **Vagrant**, **Docker** e **Go** para configurar automaticamente um cluster Docker Swarm local. O objetivo do projeto é automatizar a criação de máquinas virtuais e a configuração de um cluster Docker Swarm, proporcionando uma maneira fácil e eficiente de gerenciar clusters locais sem implementações manuais.

## Estrutura do Projeto 🏗️

A estrutura do projeto é organizada da seguinte forma:

    GOSwarm/
    ├── Docker/     
    │
    ├── vagrant/                       
    │   └── Vagrantfile                
    ├── go/                            
    │   ├── main.go                   
    │   ├── swarm/                     
    │   │   ├── manager.go            
    │   │   ├── node.go                
    │   └── utils/                     
    │       └── command.go             
    ├── README.md                     
    └── .gitignore                     


## Requisitos 📋

- **Vagrant**: Para criar e gerenciar as máquinas virtuais.
- **VirtualBox** ou outro provider compatível com o Vagrant 🧑‍💻.
- **Docker**: Para configuração e execução do cluster Docker Swarm 🐋.
- **Go (>=1.16)**: Para automatizar a configuração do cluster ⚙️.

## Como Usar 🛠️

### 1. Clonar o Repositório

Clone o repositório para sua máquina local:


git clone https://github.com/seu-usuario/GOSwarm.git
cd GOSwarm

2. Subir as Máquinas Virtuais com Vagrant 🚀

O primeiro passo é subir as máquinas virtuais definidas no Vagrantfile. Isso pode ser feito com o comando:

vagrant up

Isso criará 4 máquinas virtuais:

    master: nó manager do Docker Swarm ⚡.
    node01, node02, node03: nós worker do Docker Swarm 🛠️.

O Vagrant se encarregará de configurar o ambiente e instalar o Docker automaticamente.
3. Executar o Código Go ⚙️

Após as máquinas virtuais estarem em funcionamento, você pode executar o código Go para configurar o Docker Swarm. Execute o seguinte comando:

go run go/main.go

Esse comando irá:

    Inicializar o Docker Swarm no nó manager ⚡.
    Adicionar os nós workers ao cluster 🛠️.

4. Verificar o Cluster Docker Swarm 📊

Após a execução do código Go, você pode verificar o status do cluster Swarm com o comando:

docker node ls

Isso deverá mostrar todos os nós do Swarm, incluindo o nó manager e os nós workers.
Detalhamento do Código 🔍
Vagrantfile ⚙️

O Vagrantfile define 4 máquinas virtuais com Vagrant, sendo uma delas o nó manager do Docker Swarm e as outras três os nós worker. Cada máquina virtual é configurada com um IP fixo e o Docker é instalado automaticamente.
Código Go 💻

    main.go: Orquestra o processo de configuração do cluster Swarm, chamando as funções dos pacotes swarm e utils.
    swarm/manager.go: Contém a lógica para configurar o nó manager do Swarm ⚡.
    swarm/node.go: Contém a lógica para adicionar os nós workers ao cluster 🛠️.
    utils/command.go: Funções auxiliares para execução de comandos no shell, como a execução de comandos docker e vagrant.

Contribuindo 🤝

Sinta-se à vontade para contribuir com o projeto. Se você encontrar algum problema ou tiver sugestões de melhorias, por favor, abra um issue ou envie um pull request.
Passos para Contribuição 📝

    Faça um fork deste repositório.
    Crie uma nova branch (git checkout -b minha-branch).
    Faça suas alterações e commit (git commit -am 'Adiciona nova feature').
    Push para a branch (git push origin minha-branch).
    Abra um pull request.

Licença 📄

Este projeto é licenciado sob a MIT License.


