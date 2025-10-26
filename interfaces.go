package goiikoapi

import "context"

// IDictionaries интерфейс для работы со словарями
type IDictionaries interface {
	OrderTypes(ctx context.Context, organizationIDs []string) (*BaseOrderTypesModel, *CustomErrorModel, error)
	PaymentTypes(ctx context.Context, organizationIDs []string) (*BasePaymentTypesModel, *CustomErrorModel, error)
	Discounts(ctx context.Context, organizationIDs []string) (*BaseDiscountsModel, *CustomErrorModel, error)
	CancelCauses(ctx context.Context, organizationIDs []string) (*BaseCancelCausesModel, *CustomErrorModel, error)
	RemovalTypes(ctx context.Context, organizationIDs []string) (*BaseRemovalTypesModel, *CustomErrorModel, error)
	TipsTypes(ctx context.Context) (*BaseTipsTypesModel, *CustomErrorModel, error)
}

// IMenu интерфейс для работы с меню
type IMenu interface {
	Nomenclature(ctx context.Context, organizationID string, startRevision *int) (*BaseNomenclatureModel, *CustomErrorModel, error)
	Menu(ctx context.Context) (*BaseMenuModel, *CustomErrorModel, error)
	MenuByID(ctx context.Context, externalMenuID string, organizationIDs []string, priceCategoryID *string) (*BaseMenuByIdModel, *CustomErrorModel, error)
}

// IOrders интерфейс для работы с заказами
type IOrders interface {
	OrderCreate(ctx context.Context, organizationID, terminalGroupID string, order map[string]any, createOrderSettings *int) (*BaseCreatedOrderInfoModel, *CustomErrorModel, error)
	OrderByID(ctx context.Context, organizationIDs []string, orderIDs, posOrderIDs, returnExternalDataKeys, sourceKeys []string) (*ByIdModel, *CustomErrorModel, error)
}

// IDeliveries интерфейс для работы с доставкой
type IDeliveries interface {
	DeliveryCreate(ctx context.Context, organizationID string, order map[string]any, terminalGroupID *string, createOrderSettings *int) (*BaseCreatedDeliveryOrderInfoModel, *CustomErrorModel, error)
	UpdateOrderDeliveryStatus(ctx context.Context, organizationIDs []string, orderID, deliveryStatus, deliveryDate string) (*BaseResponseModel, *CustomErrorModel, error)
	Confirm(ctx context.Context, organizationIDs []string, orderID string) (*BaseResponseModel, *CustomErrorModel, error)
	CancelConfirmation(ctx context.Context, organizationIDs []string, orderID string) (*BaseResponseModel, *CustomErrorModel, error)
	ByDeliveryDateAndStatus(ctx context.Context, organizationIDs []string, deliveryDateFrom, deliveryDateTo string, statuses, sourceKeys []string) (*ByDeliveryDateAndStatusModel, *CustomErrorModel, error)
}

// IAddress интерфейс для работы с адресами
type IAddress interface {
	Regions(ctx context.Context, organizationIDs []string) (*BaseRegionsModel, *CustomErrorModel, error)
	Cities(ctx context.Context, organizationIDs []string) (*BaseCitiesModel, *CustomErrorModel, error)
	StreetsByCity(ctx context.Context, organizationID, cityID string) (*BaseStreetByCityModel, *CustomErrorModel, error)
}

// ITerminalGroup интерфейс для работы с группами терминалов
type ITerminalGroup interface {
	TerminalGroups(ctx context.Context, organizationIDs []string, includeDisabled bool) (*BaseTerminalGroupsModel, *CustomErrorModel, error)
	IsAlive(ctx context.Context, organizationIDs, terminalGroupIDs []string) (*BaseTGIsAliveModel, *CustomErrorModel, error)
}

// ICustomers интерфейс для работы с клиентами
type ICustomers interface {
	CustomerInfo(ctx context.Context, organizationID, identifier string, identifierType TypeRCI) (*CustomerInfoModel, *CustomErrorModel, error)
	CustomerCreateOrUpdate(ctx context.Context, organizationID string, opts ...CustomerCreateOrUpdateOption) (*CustomerCreateOrUpdateModel, *CustomErrorModel, error)
	CustomerProgramAdd(ctx context.Context, customerID, programID, organizationID string) (*CustomerProgramAddResponse, *CustomErrorModel, error)
	CustomerCardAdd(ctx context.Context, customerID, cardTrack, cardNumber, organizationID string) (*BaseResponseModel, *CustomErrorModel, error)
	CustomerCardDelete(ctx context.Context, customerID, cardTrack, organizationID string) (*BaseResponseModel, *CustomErrorModel, error)
	CustomerWalletHold(ctx context.Context, customerID, walletID, organizationID string, sum float64, transactionID, comment *string) (*WalletHoldResponse, *CustomErrorModel, error)
	CustomerWalletCancelHold(ctx context.Context, organizationID, transactionID string) (*BaseResponseModel, *CustomErrorModel, error)
	CustomerWalletTopup(ctx context.Context, customerID, walletID, organizationID string, sum float64, comment *string) (*BaseResponseModel, *CustomErrorModel, error)
	CustomerWalletChargeoff(ctx context.Context, customerID, walletID, organizationID string, sum float64, comment *string) (*BaseResponseModel, *CustomErrorModel, error)
}

// INotifications интерфейс для работы с уведомлениями
type INotifications interface {
	Send(ctx context.Context, orderSource, orderID, additionalInfo, organizationID, messageType string) (*BaseResponseModel, *CustomErrorModel, error)
}

// ICommands интерфейс для работы с командами
type ICommands interface {
	Status(ctx context.Context, organizationID, correlationID string) (*BaseStatusModel, *CustomErrorModel, error)
}

// IWebHook интерфейс для работы с webhook'ами
type IWebHook interface {
	ParseWebhookOrder(data []map[string]any) ([]WebHookDeliveryOrderEventInfoModel, error)
	ParseWebhookReserve(data []map[string]any) ([]WebHookDeliveryOrderEventInfoModel, error)
}

// IEmployees интерфейс для работы с сотрудниками
type IEmployees interface {
	Couriers(ctx context.Context, organizationIDs []string) (*BaseCouriersModel, *CustomErrorModel, error)
	EmployeeInfo(ctx context.Context, organizationID, id string) (*BaseEmployeeInfoModel, *CustomErrorModel, error)
	ShiftClockin(ctx context.Context, organizationID, terminalGroupID, employeeID string, roleID *string) (*BaseResponseModel, *CustomErrorModel, error)
	ShiftClockout(ctx context.Context, organizationID, terminalGroupID, employeeID string) (*BaseResponseModel, *CustomErrorModel, error)
	ShiftIsOpen(ctx context.Context, organizationID, terminalGroupID, employeeID string) (*BaseEmployeeInfoModel, *CustomErrorModel, error)
	ShiftByCourier(ctx context.Context, employeeID string) (*BaseEmployeeTerminalModel, *CustomErrorModel, error)
}

// IClient основной интерфейс клиента
type IClient interface {
	// Базовые методы
	Organizations(ctx context.Context, organizationIDs []string, returnAdditionalInfo, includeDisabled *bool) (*BaseOrganizationsModel, *CustomErrorModel, error)
	LastDataRaw() []byte

	// Группы методов
	GetDictionaries() IDictionaries
	GetMenu() IMenu
	GetOrders() IOrders
	GetDeliveries() IDeliveries
	GetAddress() IAddress
	GetTerminalGroup() ITerminalGroup
	GetCustomers() ICustomers
	GetNotifications() INotifications
	GetCommands() ICommands
	GetWebHook() IWebHook
	GetEmployees() IEmployees
}
