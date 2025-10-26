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
		"your-api-login",
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
	
	// Пример работы через интерфейсы
	if len(orgs.Organizations) > 0 {
		orgID := orgs.Organizations[0].ID

		// Получаем словари
		dict := cli.GetDictionaries()
		orderTypes, _, _ := dict.OrderTypes(ctx, []string{orgID})
		fmt.Println("Order types count:", len(orderTypes.OrderTypes))
		
		// Получаем меню
		menu := cli.GetMenu()
		nomenclature, _, _ := menu.Nomenclature(ctx, orgID, nil)
		fmt.Println("Nomenclature groups count:", len(nomenclature.Groups))

		employees,_,_ := cli.Employees.Couriers(ctx, orgs.ListIDs())
		fmt.Println("Employess orgs: ", len(employees.Employees))
		for _,employee := range employees.Employees {
			fmt.Println("\tEmployee active    items:",len(employee.GetActive()))
			fmt.Println("\tEmployee no active items:",len(employee.GetNoActive()))
			for _, item := range employee.GetActive(){
				fmt.Println("\t\t", item.DisplayName)
			}
		}
	}
}
