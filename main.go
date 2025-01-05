package main

import (
	"fmt"
	"emailn/internal/domain/campaing" // Importe o pacote corretamente
)

func main() {
	// Exemplo de uso do pacote campaing
	name := "Teste de Campanha"
	content := "Esse é o conteúdo da campanha."
	emails := []string{"email1@example.com", "email2@example.com"}

	// Crie uma nova campanha usando a função NewCampaing
	newCampaing := campaing.NewCampaing(name, content, emails)

	// Exiba as informações da campanha
	fmt.Printf("ID: %s\n", newCampaing.ID)
	fmt.Printf("Name: %s\n", newCampaing.Name)
	fmt.Printf("Content: %s\n", newCampaing.Content)
	fmt.Printf("Created On: %s\n", newCampaing.CreatedOn)
	fmt.Println("Contacts:")
	for _, contact := range newCampaing.Contacts {
		fmt.Printf("- %s\n", contact.Email)
	}
}
