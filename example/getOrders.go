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
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	cli, err := goiikoapi.NewClient(
		"API_LOGIN",
		goiikoapi.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	orgs, apiErr, err := cli.Organizations(ctx, nil, nil, nil)
	if err != nil {
		log.Fatal("Transport error:", err)
	}
	if apiErr != nil {
		log.Fatal("API error:", apiErr.ErrorDescription)
	}

	// Форматируем даты в формат "2006-01-02 15:04:05.000"
	deliveryDateFrom := startOfDay.Format("2006-01-02 15:04:05.000")
	deliveryDateTo := endOfDay.Format("2006-01-02 15:04:05.000")

	// Статусы для фильтрации
	statuses := []string{"CookingCompleted", "Waiting", "OnWay"}

	// Получаем заказы через метод ByDeliveryDateAndStatus
	deliveries := cli.GetDeliveries()
	result, apiErr, err := deliveries.ByDeliveryDateAndStatus(
		ctx,
		orgs.ListIDs(),
		deliveryDateFrom,
		deliveryDateTo,
		statuses,
		nil, // sourceKeys
	)
	if err != nil {
		log.Fatal("Transport error:", err)
	}
	if apiErr != nil {
		log.Fatal("API error:", apiErr.ErrorDescription)
	}

	// Выводим результаты
	fmt.Printf("MaxRevision: %d\n", result.MaxRevision)
	fmt.Printf("Orders by organizations count: %d\n", len(result.OrdersByOrganizations))

	for _, orgOrders := range result.OrdersByOrganizations {
		fmt.Printf("\nOrganization ID: %s\n", orgOrders.OrganizationID)
		fmt.Printf("Orders count: %d\n", len(orgOrders.Orders))

		for _, orderItem := range orgOrders.Orders {
			if orderItem.Order != nil {
				fmt.Printf("  Order ID: %s, Number: %d, Status: %s, Sum: %.2f\n",
					orderItem.ID,
					orderItem.Order.Number,
					orderItem.Order.Status,
					orderItem.Order.Sum,
				)
				if orderItem.Order.CourierInfo != nil {
					fmt.Printf("    Courier: %s\n", orderItem.Order.CourierInfo.Courier.Name)
				}
			}
		}
	}


}
