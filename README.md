## goiikoapi — Go-клиент для iiko Cloud API

Легковесная, типобезопасная обёртка над iiko Cloud API (iiko Transport) на Go. Повторяет функциональность Python-проекта `pyiikocloudapi`, сохраняя названия и семантику методов.

### Установка

```bash
go get github.com/kebrick/goiikoapi
```

### Быстрый старт

```go
ctx := context.Background()

cli, err := goiikoapi.NewClient(
    "<YOUR_API_LOGIN>",
    goiikoapi.WithTimeout(30*time.Second),
)
if err != nil { /* handle */ }

// Получить организации и сохранить ids внутри клиента
orgs, apiErr, err := cli.Organizations(ctx, nil, nil, nil)
if err != nil { /* transport error */ }
if apiErr != nil { /* iiko errorDescription */ }
fmt.Println("org count:", len(orgs.Organizations))
```

### Конфигурация клиента

- goiikoapi.WithBaseURL(url string)
- goiikoapi.WithHTTPClient(hc *http.Client)
- goiikoapi.WithTimeout(d time.Duration)
- goiikoapi.WithDebug(debug bool)
- goiikoapi.WithReturnDict(v bool) — зарезервировано (в Go обычно не требуется)
- goiikoapi.WithWorkingToken(token string) — использовать готовый токен

Клиент автоматически обновляет токен при его протухании и повторяет запрос один раз при 401.

### Обработка ошибок

Каждый метод возвращает три значения: (data, apiError, err)
- err — транспортная ошибка (HTTP/JSON/сеть)
- apiError — ошибка iiko (поле errorDescription из ответа) с `StatusCode`

Проверяйте сначала err, затем apiError.

### Примеры использования

#### Organizations

```go
orgs, apiErr, err := cli.Organizations(ctx, nil, nil, nil)
```

#### Использование через интерфейсы

```go
// Получаем интерфейсы
var dict IDictionaries = cli.Dictionaries()
var menu IMenu = cli.Menu()
var orders IOrders = cli.Orders()

// Работаем через интерфейсы
orderTypes, _, _ := dict.OrderTypes(ctx, []string{"orgId"})
nomenclature, _, _ := menu.Nomenclature(ctx, "orgId", nil)
orderResp, _, _ := orders.OrderCreate(ctx, "orgId", "terminalId", orderData, nil)
```

#### Dependency Injection

```go
type OrderService struct {
    orders IOrders
    menu   IMenu
}

func NewOrderService(orders IOrders, menu IMenu) *OrderService {
    return &OrderService{orders: orders, menu: menu}
}

func (s *OrderService) CreateOrder(ctx context.Context, orgID string) (*BaseCreatedOrderInfoModel, error) {
    // Валидация через меню
    _, _, err := s.menu.Nomenclature(ctx, orgID, nil)
    if err != nil { return nil, err }
    
    // Создание заказа
    order, _, err := s.orders.OrderCreate(ctx, orgID, "terminalId", orderData, nil)
    return order, err
}

// Использование
service := NewOrderService(cli.Orders(), cli.Menu())
order, err := service.CreateOrder(ctx, "orgId")
```

#### Тестирование с мок-объектами

```go
type MockOrders struct{}

func (m *MockOrders) OrderCreate(ctx context.Context, orgID, terminalID string, order map[string]any, settings *int) (*BaseCreatedOrderInfoModel, *CustomErrorModel, error) {
    return &BaseCreatedOrderInfoModel{
        OrderInfo: CreatedOrderInfoModel{ID: "mock-order-id"},
    }, nil, nil
}

func (m *MockOrders) OrderByID(ctx context.Context, orgIDs, orderIDs, posOrderIDs, returnKeys, sourceKeys []string) (*ByIdModel, *CustomErrorModel, error) {
    return &ByIdModel{}, nil, nil
}

// Проверяем реализацию интерфейса
var _ IOrders = (*MockOrders)(nil)

// Используем мок в тестах
mockOrders := &MockOrders{}
service := NewOrderService(mockOrders, nil)
order, err := service.CreateOrder(ctx, "orgId")
```

#### Dictionaries

```go
orderTypes, _, _ := cli.Dictionaries.OrderTypes(ctx, []string{"orgId"})
payments,   _, _ := cli.Dictionaries.PaymentTypes(ctx, []string{"orgId"})
discounts,  _, _ := cli.Dictionaries.Discounts(ctx, []string{"orgId"})
causes,     _, _ := cli.Dictionaries.CancelCauses(ctx, []string{"orgId"})
removals,   _, _ := cli.Dictionaries.RemovalTypes(ctx, []string{"orgId"})
tips,       _, _ := cli.Dictionaries.TipsTypes(ctx)
```

#### Menu

```go
nom, _, _ := cli.Menu.Nomenclature(ctx, "orgId", nil)
menu, _, _ := cli.Menu.Menu(ctx)
menuByID, _, _ := cli.Menu.MenuByID(ctx, "externalMenuId", []string{"orgId"}, nil)
```

#### Orders / Deliveries

```go
// Создать заказ на стол (Orders)
orderResp, apiErr, err := cli.Orders.OrderCreate(ctx, "orgId", "tgId", map[string]any{
    "phone": "+79990000000",
    // ... остальная структура заказа как в API iiko
}, nil)

// Доставка
deliveryResp, _, _ := cli.Deliveries.DeliveryCreate(ctx, "orgId", map[string]any{/* order */}, nil, nil)

// Получить заказы по id
byID, _, _ := cli.Orders.OrderByID(ctx, []string{"orgId"}, []string{"orderId"}, nil, nil, nil)

// Обновить статус доставки
_, _, _ = cli.Deliveries.UpdateOrderDeliveryStatus(ctx, []string{"orgId"}, "orderId", "Delivered", time.Now().Format("2006-01-02 15:04:05.000"))

// Подтвердить / отменить подтверждение
_, _, _ = cli.Deliveries.Confirm(ctx, []string{"orgId"}, "orderId")
_, _, _ = cli.Deliveries.CancelConfirmation(ctx, []string{"orgId"}, "orderId")

// По статусам и датам
byStatus, _, _ := cli.Deliveries.ByDeliveryDateAndStatus(ctx, []string{"orgId"}, "2024-01-01 00:00:00.000", "2024-01-02 00:00:00.000", []string{"Delivered"}, nil)
```

#### Address / Terminal groups

```go
regions, _, _ := cli.Address.Regions(ctx, []string{"orgId"})
cities,  _, _ := cli.Address.Cities(ctx,  []string{"orgId"})
streets, _, _ := cli.Address.StreetsByCity(ctx, "orgId", "cityId")

tgs, _, _ := cli.TerminalGroup.TerminalGroups(ctx, []string{"orgId"}, false)
alive, _, _ := cli.TerminalGroup.IsAlive(ctx, []string{"orgId"}, []string{"tgId"})
```

#### Customers (лояльность)

```go
// Получить клиента по телефону
cust, _, _ := cli.Customers.CustomerInfo(ctx, "orgId", "+79990000000", goiikoapi.TypeRCIPhone)

// Создать/обновить клиента
cu, _, _ := cli.Customers.CustomerCreateOrUpdate(ctx, "orgId",
    goiikoapi.WithCustomerPhone("+79990000000"),
    goiikoapi.WithCustomerName("Иван"),
)

// Программы/карты/кошельки
prog, _, _ := cli.Customers.CustomerProgramAdd(ctx, "customerId", "programId", "orgId")
_, _, _ = cli.Customers.CustomerCardAdd(ctx, "customerId", "track", "12345", "orgId")
_, _, _ = cli.Customers.CustomerCardDelete(ctx, "customerId", "track", "orgId")
hold, _, _ := cli.Customers.CustomerWalletHold(ctx, "customerId", "walletId", "orgId", 100, nil, nil)
_, _, _ = cli.Customers.CustomerWalletCancelHold(ctx, "orgId", hold.TransactionID)
_, _, _ = cli.Customers.CustomerWalletTopup(ctx, "customerId", "walletId", "orgId", 200, nil)
_, _, _ = cli.Customers.CustomerWalletChargeoff(ctx, "customerId", "walletId", "orgId", 50, nil)
```

#### Notifications / Commands

```go
// Уведомление
_, _, _ = cli.Notifications.Send(ctx, "external", "orderId", "Need attention", "orgId", "delivery_attention")

// Статус команды (по correlationId)
status, _, _ := cli.Commands.Status(ctx, "orgId", "correlationId")
```

#### WebHook (парсинг событий)

```go
events, err := cli.WebHook.ParseWebhookOrder([]map[string]any{{
    "eventType": "DeliveryOrderUpdated",
    // ... остальные поля webhook
}})
// Или глобальная функция
events, err := goiikoapi.ParseWebhookOrder([]map[string]any{...})
```

### Отладка

- `WithDebug(true)` включает подробный лог запросов/ответов (внутренний raw-body доступен через `LastDataRaw()`)
- При 401 клиент автоматически обновит токен и повторит запрос

### Соответствие Python-версии

Имена и структура методов соответствуют `pyiikocloudapi` и официальной документации iiko Cloud API. Возвращаемые структуры приведены к Go-типам со строгой типизацией и корректными json-тегами.

### Примечания

- Формат времени для полей дат: `"2006-01-02 15:04:05.000"` (локальное время терминала)
- Для отмены/контекста используйте `context.Context` во всех вызовах

