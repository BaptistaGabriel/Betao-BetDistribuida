package main

import (
	"fmt"
)

func menu() {
	var option int 
	
	for {
		fmt.Println("1. Criar conta")
		fmt.Println("2. Fazer login")
		fmt.Scanln(&option)

		switch option {
		case 1:
			// CRIAR CONTA
			createAccount()
			return
		case 2:
			// FAZER LOGIN
			fmt.Println("Fazendo login...")
			login()
			return
		default:
			// USUARIO ERROU
			fmt.Println("Errando...")
		}
	}
}

func createAccount() {
	fmt.Println("Aqui")
	return
}

func login() {
	//
}

func main() {
	menu()
}