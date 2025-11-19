package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kebrick/goiikoapi"
)

func main() {
	ctx := context.Background()

	cli, err := goiikoapi.NewClient(
		"API_LOGIN",
		goiikoapi.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Получить организации и сохранить ids внутри клиента
	orgs, apiErr, err := cli.Organizations(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal("Transport error:", err)
	}
	if apiErr != nil {
		log.Fatal("API error:", apiErr.ErrorDescription)
	}

	fmt.Println("Organizations count:", len(orgs.Organizations))

	employees, apiErr, err := cli.Employees.Couriers(ctx, orgs.ListIDs())
	if err != nil {
		log.Fatal("Transport error:", err)
	}
	if apiErr != nil {
		log.Fatal("API error:", apiErr.ErrorDescription)
	}
	fmt.Println("employees count:", len(employees.Employees))
	for _, employee := range employees.Employees {
		fmt.Println(employee.OrganizationID)
		for _, ie := range employee.Items {
		
			fmt.Println("\t", ie.DisplayName, ie.ID)
			

		}
	}

}
