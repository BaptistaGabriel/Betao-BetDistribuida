package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func menu(mapClients map[int][]string) (string, string) {
	var option int 
	
	for {
		fmt.Println("1. Criar conta")
		fmt.Println("2. Fazer login")
		fmt.Scanln(&option)

		switch option {
		case 1:
			// CRIAR CONTA
			createAccount(mapClients)
		case 2:
			// FAZER LOGIN
			fmt.Println("Fazendo login...")
			information := login(mapClients)
			return information[0], information[2]

		default:
			// USUARIO ERROU
			fmt.Println("Errando...")
		}
	}
}

func createAccount(mapClients map[int][]string) {

	// Informações
	var name string
	var age string
	var password string

	informations := []string{}

	fmt.Println("... Criar conta ...")
	fmt.Println("Nome: ")
	fmt.Scanln(&name)

	fmt.Println("Idade: ")
	fmt.Scanln(&age)

	fmt.Println("Senha: ")
	fmt.Scanln(&password)

	if age < "18" {
		fmt.Println("Site proibido para menores de idade.")
		return
	}

	// Criar carteira na blockchain

	// Diretório onde as contas são armazenadas
	folderKeys := "blockchain/keystore"

	ks := keystore.NewKeyStore(folderKeys, keystore.StandardScryptN, keystore.StandardScryptP)

	// Criando uma nova conta
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatalf("Erro ao criar conta: %v", err)
	}

	informations = append(informations, name)
	informations = append(informations, age)
	informations = append(informations, account.Address.Hex())

	number := len(mapClients) + 1
	mapClients[number] = informations
	fmt.Printf("Número da sua conta: %d\n\n", number)

	fmt.Printf("Conta criada com sucesso: %s\n", account.Address.Hex())

	return
}

func login(mapClients map[int][]string) []string {
	var id int

	fmt.Println("Número da conta: ")
	fmt.Scanln(&id)
	
	information := mapClients[id]

	return information
}

func menu2() {
	var option int
	
	for {
		fmt.Println("1. Criar eventos")
		fmt.Println("2. Ver eventos")
		fmt.Println("3. Sair")
		fmt.Scanln(&option)
	
		switch option {
		case 1:
			events()
		case 2:
			seeEvents()
		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Errou...")
		}
	}
}

func events() {
	//
}

func seeEvents() {
	//
}

func main() {

	mapClients := make(map[int][]string)

	name, account := menu(mapClients)

	fmt.Println("Bem vindo(a) ", name)
	fmt.Println("Número da sua carteira: ", account)

	menu2()
}