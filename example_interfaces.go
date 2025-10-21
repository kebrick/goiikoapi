package goiikoapi

import (
	"context"
	"fmt"
	"time"
)

// Пример использования через интерфейсы
func ExampleUsageWithInterfaces() {
	ctx := context.Background()
	
	// Создаем клиент
	client, err := NewClient("your-api-login")
	if err != nil {
		panic(err)
	}
	
	// Используем через интерфейсы
	var dict IDictionaries = client.Dictionaries()
	var menu IMenu = client.Menu()
	var orders IOrders = client.Orders()
	
	// Получаем организации
	orgs, _, _ := client.Organizations(ctx, nil, nil, nil)
	if len(orgs.Organizations) == 0 {
		return
	}
	orgID := orgs.Organizations[0].ID
	
	// Работаем со словарями
	orderTypes, _, _ := dict.OrderTypes(ctx, []string{orgID})
	fmt.Printf("Order types: %d\n", len(orderTypes.OrderTypes))
	
	// Работаем с меню
	nomenclature, _, _ := menu.Nomenclature(ctx, orgID, nil)
	fmt.Printf("Nomenclature groups: %d\n", len(nomenclature.Groups))
	
	// Работаем с заказами
	orderResp, _, _ := orders.OrderCreate(ctx, orgID, "terminal-group-id", map[string]any{
		"phone": "+79990000000",
		"items": []map[string]any{{
			"productId": "product-id",
			"amount":    1,
		}},
	}, nil)
	fmt.Printf("Created order: %s\n", orderResp.OrderInfo.ID)
}

// Пример dependency injection
type OrderService struct {
	orders IOrders
	menu   IMenu
}

func NewOrderService(orders IOrders, menu IMenu) *OrderService {
	return &OrderService{
		orders: orders,
		menu:   menu,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, orgID, terminalGroupID string, phone string) (*BaseCreatedOrderInfoModel, *CustomErrorModel, error) {
	// Получаем меню для валидации
	_, _, err := s.menu.Nomenclature(ctx, orgID, nil)
	if err != nil {
		return nil, nil, err
	}
	
	// Создаем заказ
	return s.orders.OrderCreate(ctx, orgID, terminalGroupID, map[string]any{
		"phone": phone,
		"items": []map[string]any{{
			"productId": "default-product",
			"amount":    1,
		}},
	}, nil)
}

// Пример мок-объекта для тестирования
type MockOrders struct{}

func (m *MockOrders) OrderCreate(ctx context.Context, organizationID, terminalGroupID string, order map[string]any, createOrderSettings *int) (*BaseCreatedOrderInfoModel, *CustomErrorModel, error) {
	return &BaseCreatedOrderInfoModel{
		OrderInfo: CreatedOrderInfoModel{
			ID: "mock-order-id",
		},
	}, nil, nil
}

func (m *MockOrders) OrderByID(ctx context.Context, organizationIDs []string, orderIDs, posOrderIDs, returnExternalDataKeys, sourceKeys []string) (*ByIdModel, *CustomErrorModel, error) {
	return &ByIdModel{}, nil, nil
}

// Проверяем, что MockOrders реализует IOrders
var _ IOrders = (*MockOrders)(nil)

// Пример тестирования с мок-объектом
func ExampleTestingWithMock() {
	ctx := context.Background()
	
	// Создаем мок-объект
	mockOrders := &MockOrders{}
	
	// Создаем сервис с мок-объектом
	service := NewOrderService(mockOrders, nil) // menu может быть nil для этого теста
	
	// Тестируем создание заказа
	order, _, err := service.CreateOrder(ctx, "org-id", "terminal-id", "+79990000000")
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("Mock order created: %s\n", order.OrderInfo.ID)
}

// Пример работы с несколькими интерфейсами
func ExampleMultipleInterfaces() {
	ctx := context.Background()
	client, _ := NewClient("api-login")
	
	// Получаем все нужные интерфейсы
	dict := client.Dictionaries()
	menu := client.Menu()
	orders := client.Orders()
	deliveries := client.Deliveries()
	customers := client.Customers()
	
	// Получаем организации
	orgs, _, _ := client.Organizations(ctx, nil, nil, nil)
	if len(orgs.Organizations) == 0 {
		return
	}
	orgID := orgs.Organizations[0].ID
	
	// Работаем с разными группами методов
	orderTypes, _, _ := dict.OrderTypes(ctx, []string{orgID})
	nomenclature, _, _ := menu.Nomenclature(ctx, orgID, nil)
	
	// Создаем заказ
	orderResp, _, _ := orders.OrderCreate(ctx, orgID, "terminal-id", map[string]any{
		"phone": "+79990000000",
		"items": []map[string]any{{
			"productId": "product-id",
			"amount":    1,
		}},
	}, nil)
	
	// Создаем доставку
	deliveryResp, _, _ := deliveries.DeliveryCreate(ctx, orgID, map[string]any{
		"phone": "+79990000000",
		"items": []map[string]any{{
			"productId": "product-id",
			"amount":    1,
		}},
	}, nil, nil)
	
	// Работаем с клиентами
	customer, _, _ := customers.CustomerInfo(ctx, orgID, "+79990000000", TypeRCIPhone)
	
	fmt.Printf("Order types: %d\n", len(orderTypes.OrderTypes))
	fmt.Printf("Nomenclature groups: %d\n", len(nomenclature.Groups))
	fmt.Printf("Created order: %s\n", orderResp.OrderInfo.ID)
	fmt.Printf("Created delivery: %s\n", deliveryResp.OrderInfo.ID)
	fmt.Printf("Customer found: %v\n", customer != nil)
}
