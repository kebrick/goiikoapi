package goiikoapi

import (
	"encoding/json"
)

// BaseResponseModel соответствует полю correlationId во всех ответах
type BaseResponseModel struct {
	CorrelationID string `json:"correlationId,omitempty"`
}

// ErrorModel базовая ошибка API iiko
type ErrorModel struct {
	BaseResponseModel
	ErrorDescription string `json:"errorDescription,omitempty"`
	Error            string `json:"error,omitempty"`
}

// CustomErrorModel с добавлением HTTP статуса
type CustomErrorModel struct {
	ErrorModel
	StatusCode int `json:"-"`
}

// IdNameModel универсальная модель id+name
type IdNameModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrganizationModel соответствует Organization из /api/1/organizations
type OrganizationModel struct {
	ID                                string   `json:"id"`
	Name                              string   `json:"name"`
	Country                           *string  `json:"country,omitempty"`
	RestaurantAddress                 *string  `json:"restaurantAddress,omitempty"`
	Latitude                          *float64 `json:"latitude,omitempty"`
	Longitude                         *float64 `json:"longitude,omitempty"`
	UseUaeAddressingSystem            *bool    `json:"useUaeAddressingSystem,omitempty"`
	Version                           *string  `json:"version,omitempty"`
	CurrencyIsoName                   *string  `json:"currencyIsoName,omitempty"`
	CurrencyMinimumDenomination       *float64 `json:"currencyMinimumDenomination,omitempty"`
	CountryPhoneCode                  *string  `json:"countryPhoneCode,omitempty"`
	MarketingSourceRequiredInDelivery *bool    `json:"marketingSourceRequiredInDelivery,omitempty"`
	DefaultDeliveryCityID             *string  `json:"defaultDeliveryCityId,omitempty"`
	DeliveryCityIDs                   []string `json:"deliveryCityIds,omitempty"`
	DeliveryServiceType               *string  `json:"deliveryServiceType,omitempty"`
	DefaultCallCenterPaymentTypeID    *string  `json:"defaultCallCenterPaymentTypeId,omitempty"`
	OrderItemCommentEnabled           *bool    `json:"orderItemCommentEnabled,omitempty"`
	Inn                               *string  `json:"inn,omitempty"`
	AddressFormatType                 *string  `json:"addressFormatType,omitempty"`
	IsConfirmationEnabled             *bool    `json:"isConfirmationEnabled,omitempty"`
	ConfirmAllowedIntervalInMinutes   *int     `json:"confirmAllowedIntervalInMinutes,omitempty"`
	ResponseType                      *string  `json:"responseType,omitempty"`
}

// BaseOrganizationsModel корневой ответ для organizations
type BaseOrganizationsModel struct {
	BaseResponseModel
	Organizations []OrganizationModel `json:"organizations"`
}

// ListIDs утилита для получения списка id организаций
func (m BaseOrganizationsModel) ListIDs() []string {
	ids := make([]string, 0, len(m.Organizations))
	for _, o := range m.Organizations {
		ids = append(ids, o.ID)
	}
	return ids
}

// internalErrorEnvelope минимальный конверт для проверки errorDescription в ответе
type internalErrorEnvelope struct {
	ErrorDescription string `json:"errorDescription,omitempty"`
}

// DetectAPIError пытается извлечь errorDescription из произвольного ответа
func DetectAPIError(body []byte) (string, bool) {
	var env internalErrorEnvelope
	if err := json.Unmarshal(body, &env); err != nil {
		return "", false
	}
	if env.ErrorDescription != "" {
		return env.ErrorDescription, true
	}
	return "", false
}

// OrderTypeModel для order_types
type OrderTypeModel struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	OrderServiceType string `json:"orderServiceType"`
	IsDeleted        bool   `json:"isDeleted"`
	ExternalRevision *int   `json:"externalRevision,omitempty"`
}

type OrderTypeOrganizationModel struct {
	OrganizationID string           `json:"organizationId"`
	Items          []OrderTypeModel `json:"items"`
}

type BaseOrderTypesModel struct {
	BaseResponseModel
	OrderTypes []OrderTypeOrganizationModel `json:"orderTypes"`
}

// PaymentTypeModel для payment_types
type PaymentTypeModel struct {
	ID                           string        `json:"id"`
	Name                         string        `json:"name"`
	Code                         *string       `json:"code,omitempty"`
	Comment                      *string       `json:"comment,omitempty"`
	Combinable                   bool          `json:"combinable"`
	ExternalRevision             *int          `json:"externalRevision,omitempty"`
	ApplicableMarketingCampaigns []string      `json:"applicableMarketingCampaigns,omitempty"`
	IsDeleted                    bool          `json:"isDeleted"`
	PrintCheque                  bool          `json:"printCheque"`
	PaymentProcessingType        *string       `json:"paymentProcessingType,omitempty"`
	PaymentTypeKind              *string       `json:"paymentTypeKind,omitempty"`
	TerminalGroups               []IdNameModel `json:"terminalGroups,omitempty"`
}

type BasePaymentTypesModel struct {
	BaseResponseModel
	PaymentTypes []PaymentTypeModel `json:"paymentTypes"`
}

// DiscountModel для discounts
type DiscountProductCategoryModel struct {
	CategoryID   string  `json:"categoryId"`
	CategoryName *string `json:"categoryName,omitempty"`
	Percent      float64 `json:"percent"`
}

type DiscountItemModel struct {
	ID                       string                         `json:"id"`
	Name                     string                         `json:"name"`
	Percent                  float64                        `json:"percent"`
	IsCategorisedDiscount    bool                           `json:"isCategorisedDiscount"`
	ProductCategoryDiscounts []DiscountProductCategoryModel `json:"productCategoryDiscounts"`
	Comment                  *string                        `json:"comment,omitempty"`
	CanBeAppliedSelectively  string                         `json:"canBeAppliedSelectively"`
	MinOrderSum              *float64                       `json:"minOrderSum,omitempty"`
	Mode                     string                         `json:"mode"`
	Sum                      float64                        `json:"sum"`
	CanApplyByCardNumber     bool                           `json:"canApplyByCardNumber"`
	IsManual                 bool                           `json:"isManual"`
	IsCard                   bool                           `json:"isCard"`
	IsAutomatic              bool                           `json:"isAutomatic"`
	IsDeleted                bool                           `json:"isDeleted"`
}

type DiscountOrganizationModel struct {
	OrganizationID string              `json:"organizationId"`
	Items          []DiscountItemModel `json:"items"`
}

type BaseDiscountsModel struct {
	BaseResponseModel
	Discounts []DiscountOrganizationModel `json:"discounts"`
}

// CancelCauseModel для cancel_causes
type CancelCauseModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"isDeleted"`
}

type BaseCancelCausesModel struct {
	BaseResponseModel
	CancelCauses []CancelCauseModel `json:"cancelCauses"`
}

// RemovalTypeModel для removal_types
type RemovalTypeModel struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Comment             *string `json:"comment,omitempty"`
	CanWriteoffToCafe   bool    `json:"canWriteoffToCafe"`
	CanWriteoffToWaiter bool    `json:"canWriteoffToWaiter"`
	CanWriteoffToUser   bool    `json:"canWriteoffToUser"`
	ReasonRequired      bool    `json:"reasonRequired"`
	Manual              bool    `json:"manual"`
	IsDeleted           bool    `json:"isDeleted"`
}

type BaseRemovalTypesModel struct {
	BaseResponseModel
	RemovalTypes []RemovalTypeModel `json:"removalTypes"`
}

// TipsTypeModel для tips_types
type TipsTypeModel struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	OrganizationIDs   []string `json:"organizationIds"`
	OrderServiceTypes []string `json:"orderServiceTypes"`
	PaymentTypesIDs   []string `json:"paymentTypesIds"`
}

type BaseTipsTypesModel struct {
	BaseResponseModel
	TipsTypes []TipsTypeModel `json:"tipsTypes"`
}

// Menu models для nomenclature, menu, menu_by_id
type NomenclatureGroupModel struct {
	ImageLinks       []string `json:"imageLinks"`
	ParentGroup      *string  `json:"parentGroup,omitempty"`
	Order            int      `json:"order"`
	IsIncludedInMenu bool     `json:"isIncludedInMenu"`
	IsGroupModifier  bool     `json:"isGroupModifier"`
	ID               string   `json:"id"`
	Code             *string  `json:"code,omitempty"`
	Name             string   `json:"name"`
	Description      *string  `json:"description,omitempty"`
	AdditionalInfo   *string  `json:"additionalInfo,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	IsDeleted        *bool    `json:"isDeleted,omitempty"`
	SeoDescription   *string  `json:"seoDescription,omitempty"`
	SeoText          *string  `json:"seoText,omitempty"`
	SeoKeywords      *string  `json:"seoKeywords,omitempty"`
	SeoTitle         *string  `json:"seoTitle,omitempty"`
}

type ProductCategoryModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"isDeleted"`
}

type SizeModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Priority  *int   `json:"priority,omitempty"`
	IsDefault *bool  `json:"isDefault,omitempty"`
}

type SizePriceModel struct {
	CurrentPrice       float64  `json:"currentPrice"`
	IsIncludedInMenu   bool     `json:"isIncludedInMenu"`
	NextPrice          *float64 `json:"nextPrice,omitempty"`
	NextIncludedInMenu bool     `json:"nextIncludedInMenu"`
	NextDatePrice      *string  `json:"nextDatePrice,omitempty"`
}

type SizePriceItemModel struct {
	SizeID *string        `json:"sizeId,omitempty"`
	Price  SizePriceModel `json:"price"`
}

type ModifierModel struct {
	ID                  string `json:"id"`
	DefaultAmount       *int   `json:"defaultAmount,omitempty"`
	MinAmount           int    `json:"minAmount"`
	MaxAmount           int    `json:"maxAmount"`
	Required            *bool  `json:"required,omitempty"`
	HideIfDefaultAmount *bool  `json:"hideIfDefaultAmount,omitempty"`
	Splittable          *bool  `json:"splittable,omitempty"`
	FreeOfChargeAmount  *int   `json:"freeOfChargeAmount,omitempty"`
}

type GroupModifierModel struct {
	ID                                   string          `json:"id"`
	MinAmount                            int             `json:"minAmount"`
	MaxAmount                            int             `json:"maxAmount"`
	Required                             bool            `json:"required"`
	ChildModifiersHaveMinMaxRestrictions *bool           `json:"childModifiersHaveMinMaxRestrictions,omitempty"`
	ChildModifiers                       []ModifierModel `json:"childModifiers"`
	HideIfDefaultAmount                  *bool           `json:"hideIfDefaultAmount,omitempty"`
	DefaultAmount                        *int            `json:"defaultAmount,omitempty"`
	Splittable                           *bool           `json:"splittable,omitempty"`
	FreeOfChargeAmount                   *int            `json:"freeOfChargeAmount,omitempty"`
}

type ProductModel struct {
	FatAmount               *float64              `json:"fatAmount,omitempty"`
	ProteinsAmount          *float64              `json:"proteinsAmount,omitempty"`
	CarbohydratesAmount     *float64              `json:"carbohydratesAmount,omitempty"`
	EnergyAmount            *float64              `json:"energyAmount,omitempty"`
	FatFullAmount           *float64              `json:"fatFullAmount,omitempty"`
	ProteinsFullAmount      *float64              `json:"proteinsFullAmount,omitempty"`
	CarbohydratesFullAmount *float64              `json:"carbohydratesFullAmount,omitempty"`
	EnergyFullAmount        *float64              `json:"energyFullAmount,omitempty"`
	Weight                  *float64              `json:"weight,omitempty"`
	GroupID                 *string               `json:"groupId,omitempty"`
	ProductCategoryID       *string               `json:"productCategoryId,omitempty"`
	Type                    *string               `json:"type,omitempty"`
	OrderItemType           string                `json:"orderItemType"`
	ModifierSchemaID        *string               `json:"modifierSchemaId,omitempty"`
	ModifierSchemaName      *string               `json:"modifierSchemaName,omitempty"`
	Splittable              bool                  `json:"splittable"`
	MeasureUnit             string                `json:"measureUnit"`
	SizePrices              []SizePriceItemModel  `json:"sizePrices"`
	Modifiers               []ModifierModel       `json:"modifiers"`
	GroupModifiers          []*GroupModifierModel `json:"groupModifiers"`
	ImageLinks              []string              `json:"imageLinks"`
	DoNotPrintInCheque      bool                  `json:"doNotPrintInCheque"`
	ParentGroup             *string               `json:"parentGroup,omitempty"`
	Order                   int                   `json:"order"`
	FullNameEnglish         *string               `json:"fullNameEnglish,omitempty"`
	UseBalanceForSell       bool                  `json:"useBalanceForSell"`
	CanSetOpenPrice         bool                  `json:"canSetOpenPrice"`
	ID                      string                `json:"id"`
	Code                    *string               `json:"code,omitempty"`
	Name                    string                `json:"name"`
	Description             *string               `json:"description,omitempty"`
	AdditionalInfo          *string               `json:"additionalInfo,omitempty"`
	Tags                    []string              `json:"tags,omitempty"`
	IsDeleted               *bool                 `json:"isDeleted,omitempty"`
	SeoDescription          *string               `json:"seoDescription,omitempty"`
	SeoText                 *string               `json:"seoText,omitempty"`
	SeoKeywords             *string               `json:"seoKeywords,omitempty"`
	SeoTitle                *string               `json:"seoTitle,omitempty"`
}

type BaseNomenclatureModel struct {
	BaseResponseModel
	Groups            []NomenclatureGroupModel `json:"groups"`
	ProductCategories []ProductCategoryModel   `json:"productCategories"`
	Products          []ProductModel           `json:"products"`
	Sizes             []SizeModel              `json:"sizes"`
	Revision          int                      `json:"revision"`
}

// Menu models для /api/2/menu
type BaseMenuModel struct {
	BaseResponseModel
	ExternalMenus   []IdNameModel `json:"externalMenus,omitempty"`
	PriceCategories []IdNameModel `json:"priceCategories,omitempty"`
}

// Menu by ID models для /api/2/menu/by_id
type AllergenGroupModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type TaxCategoryModel struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Percentage float64 `json:"percentage"`
}

type PriceModel struct {
	OrganizationID string  `json:"organizationId"`
	Price          float64 `json:"price"`
}

type RestrictionModel struct {
	MinQuantity  int `json:"minQuantity"`
	MaxQuantity  int `json:"maxQuantity"`
	FreeQuantity int `json:"freeQuantity"`
	ByDefault    int `json:"byDefault"`
}

type ItemModifierGroupItemModel struct {
	Prices                   []PriceModel         `json:"prices"`
	SKU                      string               `json:"sku"`
	Name                     string               `json:"name"`
	Description              string               `json:"description"`
	ButtonImage              string               `json:"buttonImage"`
	Restrictions             RestrictionModel     `json:"restrictions"`
	AllergenGroups           []AllergenGroupModel `json:"allergenGroups"`
	NutritionPerHundredGrams map[string]any       `json:"nutritionPerHundredGrams"`
	PortionWeightGrams       float64              `json:"portionWeightGrams"`
	Tags                     []IdNameModel        `json:"tags"`
	ItemID                   string               `json:"itemId"`
}

type ItemModifierGroupModel struct {
	Items                                []ItemModifierGroupItemModel `json:"items"`
	Name                                 string                       `json:"name"`
	Description                          string                       `json:"description"`
	Restrictions                         RestrictionModel             `json:"restrictions"`
	CanBeDivided                         bool                         `json:"canBeDivided"`
	ItemGroupID                          string                       `json:"itemGroupId"`
	ChildModifiersHaveMinMaxRestrictions bool                         `json:"childModifiersHaveMinMaxRestrictions"`
	SKU                                  string                       `json:"sku"`
}

type ItemSizeModel struct {
	Prices                   PriceModel               `json:"prices"`
	ItemModifierGroups       []ItemModifierGroupModel `json:"itemModifierGroups"`
	SKU                      string                   `json:"sku"`
	SizeCode                 string                   `json:"sizeCode"`
	SizeName                 string                   `json:"sizeName"`
	IsDefault                *bool                    `json:"isDefault,omitempty"`
	PortionWeightGrams       float64                  `json:"portionWeightGrams"`
	SizeID                   string                   `json:"sizeId"`
	NutritionPerHundredGrams map[string]any           `json:"nutritionPerHundredGrams"`
	ButtonImageURL           string                   `json:"buttonImageUrl"`
	ButtonImageCroppedURL    string                   `json:"buttonImageCroppedUrl"`
}

type MenuItemModel struct {
	SKU              string               `json:"sku"`
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	AllergenGroups   []AllergenGroupModel `json:"allergenGroups"`
	ItemID           string               `json:"itemId"`
	ModifierSchemaID string               `json:"modofierSchemaId"`
	TaxCategory      TaxCategoryModel     `json:"taxCategory"`
	OrderItemType    string               `json:"orderItemType"`
	ItemSizes        []ItemSizeModel      `json:"itemSizes"`
}

type MenuItemCategoryModel struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	ButtonImageURL string          `json:"buttonImageUrl"`
	HeaderImageURL string          `json:"headerImageUrl"`
	Items          []MenuItemModel `json:"items"`
}

type BaseMenuByIdModel struct {
	ID             string                  `json:"id"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description"`
	ItemCategories []MenuItemCategoryModel `json:"itemCategories"`
}

// Orders/Deliveries models
type CustomerModel struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Surname         *string `json:"surname,omitempty"`
	Comment         *string `json:"comment,omitempty"`
	Gender          string  `json:"gender"`
	InBlacklist     bool    `json:"inBlacklist"`
	BlacklistReason *string `json:"blacklistReason,omitempty"`
	Birthdate       *string `json:"birthdate,omitempty"`
	Type            string  `json:"type"`
}

type CauseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CancelInfoModel struct {
	WhenCancelled string     `json:"whenCancelled"`
	Cause         CauseModel `json:"cause"`
	Comment       *string    `json:"comment,omitempty"`
}

type EmployeeModel struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Phone *string `json:"phone,omitempty"`
}

type CourierInfoModel struct {
	Courier                   EmployeeModel `json:"courier"`
	IsCourierSelectedManually bool          `json:"isCourierSelectedManually"`
}

type ProblemOrderModel struct {
	HasProblem  bool    `json:"hasProblem"`
	Description *string `json:"description,omitempty"`
}

type MarketingSourceOrderModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExternalCourierServiceOrderModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ConceptionOrderModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type GuestsInfoOrderModel struct {
	Count               int  `json:"count"`
	SplitBetweenPersons bool `json:"splitBetweenPersons"`
}

type CombosItemOrderModel struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Amount   int     `json:"amount"`
	Price    float64 `json:"price"`
	SourceID string  `json:"sourceId"`
}

type PaymentTypeOrderModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type PaymentItemOrderModel struct {
	PaymentType            PaymentTypeOrderModel `json:"paymentType"`
	Sum                    float64               `json:"sum"`
	IsPreliminary          bool                  `json:"isPreliminary"`
	IsExternal             bool                  `json:"isExternal"`
	IsProcessedExternally  bool                  `json:"isProcessedExternally"`
	IsFiscalizedExternally *bool                 `json:"isFiscalizedExternally,omitempty"`
}

type TipsTypeOrderModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TipsItemOrderModel struct {
	TipsType               TipsTypeOrderModel    `json:"tipsType"`
	PaymentType            PaymentTypeOrderModel `json:"paymentType"`
	Sum                    float64               `json:"sum"`
	IsPreliminary          bool                  `json:"isPreliminary"`
	IsExternal             bool                  `json:"isExternal"`
	IsProcessedExternally  bool                  `json:"isProcessedExternally"`
	IsFiscalizedExternally *bool                 `json:"isFiscalizedExternally,omitempty"`
}

type DiscountTypeModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DiscountsItemOrderModel struct {
	DiscountType       DiscountTypeModel `json:"discountType"`
	Sum                float64           `json:"sum"`
	SelectivePositions []string          `json:"selectivePositions,omitempty"`
}

type OrderItemDeletionMethodModel struct {
	ID          string       `json:"id"`
	Comment     *string      `json:"comment,omitempty"`
	RemovalType *IdNameModel `json:"removalType,omitempty"`
}

type OrderItemDeletedModel struct {
	DeletionMethod *OrderItemDeletionMethodModel `json:"deletionMethod,omitempty"`
}

type OrderTypeOrderModel struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	OrderServiceType string `json:"orderServiceType"`
}

type OrderItemComboInformationModel struct {
	ComboID       string `json:"comboId"`
	ComboSourceID string `json:"comboSourceId"`
	GroupID       string `json:"groupId"`
}

type OrderProductItemModel struct {
	Product          IdNameModel                     `json:"product"`
	Modifiers        []OrderProductItemModel         `json:"modifiers,omitempty"`
	Price            *float64                        `json:"price,omitempty"`
	Cost             float64                         `json:"cost"`
	PricePredefined  bool                            `json:"pricePredefined"`
	PositionID       *string                         `json:"positionId,omitempty"`
	TaxPercent       *float64                        `json:"taxPercent,omitempty"`
	Type             string                          `json:"type"`
	Status           string                          `json:"status"`
	Deleted          *OrderItemDeletedModel          `json:"deleted,omitempty"`
	Amount           float64                         `json:"amount"`
	Comment          *string                         `json:"comment,omitempty"`
	WhenPrinted      *string                         `json:"whenPrinted,omitempty"`
	Size             *IdNameModel                    `json:"size,omitempty"`
	ComboInformation *OrderItemComboInformationModel `json:"comboInformation,omitempty"`
}

type CreatedDeliveryOrderModel struct {
	ParentDeliveryID         *string                           `json:"parentDeliveryId,omitempty"`
	Customer                 *CustomerModel                    `json:"customer,omitempty"`
	Phone                    string                            `json:"phone"`
	DeliveryPoint            map[string]any                    `json:"deliveryPoint,omitempty"`
	Status                   string                            `json:"status"`
	CancelInfo               *CancelInfoModel                  `json:"cancelInfo,omitempty"`
	CourierInfo              *CourierInfoModel                 `json:"courierInfo,omitempty"`
	CompleteBefore           string                            `json:"completeBefore"`
	WhenCreated              string                            `json:"whenCreated"`
	WhenConfirmed            *string                           `json:"whenConfirmed,omitempty"`
	WhenPrinted              *string                           `json:"whenPrinted,omitempty"`
	WhenSended               *string                           `json:"whenSended,omitempty"`
	WhenDelivered            *string                           `json:"whenDelivered,omitempty"`
	Comment                  *string                           `json:"comment,omitempty"`
	Problem                  *ProblemOrderModel                `json:"problem,omitempty"`
	Operator                 *EmployeeModel                    `json:"operator,omitempty"`
	MarketingSource          *MarketingSourceOrderModel        `json:"marketingSource,omitempty"`
	DeliveryDuration         *int                              `json:"deliveryDuration,omitempty"`
	IndexInCourierRoute      *int                              `json:"indexInCourierRoute,omitempty"`
	CookingStartTime         string                            `json:"cookingStartTime"`
	IsDeleted                *bool                             `json:"isDeleted,omitempty"`
	WhenReceivedByAPI        *string                           `json:"whenReceivedByApi,omitempty"`
	WhenReceivedFromFront    *string                           `json:"whenReceivedFromFront,omitempty"`
	MovedFromDeliveryID      *string                           `json:"movedFromDeliveryId,omitempty"`
	MovedFromTerminalGroupID *string                           `json:"movedFromTerminalGroupId,omitempty"`
	MovedFromOrganizationID  *string                           `json:"movedFromOrganizationId,omitempty"`
	ExternalCourierService   *ExternalCourierServiceOrderModel `json:"externalCourierService,omitempty"`
	Sum                      float64                           `json:"sum"`
	Number                   int                               `json:"number"`
	SourceKey                *string                           `json:"sourceKey,omitempty"`
	WhenBillPrinted          *string                           `json:"whenBillPrinted,omitempty"`
	WhenClosed               *string                           `json:"whenClosed,omitempty"`
	Conception               *ConceptionOrderModel             `json:"conception,omitempty"`
	GuestsInfo               GuestsInfoOrderModel              `json:"guestsInfo"`
	Items                    []OrderProductItemModel           `json:"items"`
	Combos                   []CombosItemOrderModel            `json:"combos,omitempty"`
	Payments                 []PaymentItemOrderModel           `json:"payments,omitempty"`
	Tips                     []TipsItemOrderModel              `json:"tips,omitempty"`
	Discounts                []DiscountsItemOrderModel         `json:"discounts,omitempty"`
	OrderType                *OrderTypeModel                   `json:"orderType,omitempty"`
	TerminalGroupID          string                            `json:"terminalGroupId"`
	ProcessedPaymentsSum     *int                              `json:"processedPaymentsSum,omitempty"`
}

type ErrorInfoModel struct {
	Code           string      `json:"code"`
	Message        *string     `json:"message,omitempty"`
	Description    *string     `json:"description,omitempty"`
	AdditionalData interface{} `json:"additionalData,omitempty"`
}

type ByOrderItemModel struct {
	ID             string                     `json:"id"`
	ExternalNumber *string                    `json:"externalNumber,omitempty"`
	OrganizationID string                     `json:"organizationId"`
	Timestamp      int64                      `json:"timestamp"`
	CreationStatus *string                    `json:"creationStatus,omitempty"`
	ErrorInfo      *ErrorInfoModel            `json:"errorInfo,omitempty"`
	Order          *CreatedDeliveryOrderModel `json:"order,omitempty"`
}

type ByIdModel struct {
	BaseResponseModel
	Orders []ByOrderItemModel `json:"orders,omitempty"`
}

type OrdersByOrganizationsModel struct {
	OrganizationID string             `json:"organizationId"`
	Orders         []ByOrderItemModel `json:"orders,omitempty"`
}

type ByDeliveryDateAndStatusModel struct {
	BaseResponseModel
	MaxRevision           int                          `json:"maxRevision"`
	OrdersByOrganizations []OrdersByOrganizationsModel `json:"ordersByOrganizations,omitempty"`
}

type ByDeliveryDateAndSourceKeyAndFilter struct {
	ByDeliveryDateAndStatusModel
}

type CreatedOrderInfoModel struct {
	ID             string                     `json:"id"`
	ExternalNumber *string                    `json:"externalNumber,omitempty"`
	OrganizationID string                     `json:"organizationId"`
	Timestamp      int64                      `json:"timestamp"`
	CreationStatus *string                    `json:"creationStatus,omitempty"`
	ErrorInfo      *ErrorInfoModel            `json:"errorInfo,omitempty"`
	Order          *CreatedDeliveryOrderModel `json:"order,omitempty"`
}

type BaseCreatedOrderInfoModel struct {
	BaseResponseModel
	OrderInfo CreatedOrderInfoModel `json:"orderInfo"`
}

type BaseCreatedDeliveryOrderInfoModel struct {
	BaseResponseModel
	OrderInfo *CreatedOrderInfoModel `json:"orderInfo,omitempty"`
}

// Address/TerminalGroup models
type RegionsItemModel struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ExternalRevision *int   `json:"externalRevision,omitempty"`
	IsDeleted        bool   `json:"isDeleted"`
}

type RegionsModel struct {
	OrganizationID string             `json:"organizationId"`
	Items          []RegionsItemModel `json:"items,omitempty"`
}

type BaseRegionsModel struct {
	BaseResponseModel
	Regions []RegionsModel `json:"regions,omitempty"`
}

type CitiesItemModel struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	ExternalRevision *int    `json:"externalRevision,omitempty"`
	IsDeleted        bool    `json:"isDeleted"`
	ClassifierID     *string `json:"classifierId,omitempty"`
	AdditionalInfo   *string `json:"additionalInfo,omitempty"`
}

type CitiesModel struct {
	OrganizationID string            `json:"organizationId"`
	Items          []CitiesItemModel `json:"items,omitempty"`
}

type BaseCitiesModel struct {
	BaseResponseModel
	Cities []CitiesModel `json:"cities,omitempty"`
}

type StreetsItemModel struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	ExternalRevision *int    `json:"externalRevision,omitempty"`
	ClassifierID     *string `json:"classifierId,omitempty"`
	IsDeleted        bool    `json:"isDeleted"`
}

type BaseStreetByCityModel struct {
	BaseResponseModel
	Streets []StreetsItemModel `json:"streets,omitempty"`
}

type TerminalGroupItemModel struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	OrganizationID string  `json:"organizationId"`
	Address        *string `json:"address,omitempty"`
}

type TerminalGroupsModel struct {
	OrganizationID string                   `json:"organizationId"`
	Items          []TerminalGroupItemModel `json:"items,omitempty"`
}

type BaseTerminalGroupsModel struct {
	BaseResponseModel
	TerminalGroups []TerminalGroupsModel `json:"terminalGroups,omitempty"`
}

type TGIsAliveItemModel struct {
	IsAlive         bool   `json:"isAlive"`
	TerminalGroupID string `json:"terminalGroupId"`
	OrganizationID  string `json:"organizationId"`
}

type BaseTGIsAliveModel struct {
	BaseResponseModel
	IsAliveStatus []TGIsAliveItemModel `json:"isAliveStatus,omitempty"`
}

// Customers models
type CardCIModel struct {
	ID          string  `json:"id"`
	Track       string  `json:"track"`
	Number      string  `json:"number"`
	ValidToDate *string `json:"validToDate,omitempty"`
}

type CategoriesCIModel struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	IsActive              bool   `json:"isActive"`
	IsDefaultForNewGuests bool   `json:"isDefaultForNewGuests"`
}

type WalletBalanceCIModel struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Type    int     `json:"type"`
	Balance float64 `json:"balance"`
}

type CustomerInfoModel struct {
	ID                            string                 `json:"id"`
	ReferrerID                    *string                `json:"referrerId,omitempty"`
	Name                          *string                `json:"name,omitempty"`
	Surname                       *string                `json:"surname,omitempty"`
	MiddleName                    *string                `json:"middleName,omitempty"`
	Comment                       *string                `json:"comment,omitempty"`
	Phone                         *string                `json:"phone,omitempty"`
	CultureName                   *string                `json:"cultureName,omitempty"`
	Birthday                      *string                `json:"birthday,omitempty"`
	Email                         *string                `json:"email,omitempty"`
	Sex                           int                    `json:"sex"`
	ConsentStatus                 int                    `json:"consentStatus"`
	Anonymized                    bool                   `json:"anonymized"`
	Cards                         []CardCIModel          `json:"cards,omitempty"`
	Categories                    []CategoriesCIModel    `json:"categories,omitempty"`
	WalletBalances                []WalletBalanceCIModel `json:"walletBalances,omitempty"`
	UserData                      *string                `json:"userData,omitempty"`
	ShouldReceivePromoActionsInfo *bool                  `json:"shouldReceivePromoActionsInfo,omitempty"`
	ShouldReceiveLoyaltyInfo      *bool                  `json:"shouldReceiveLoyaltyInfo,omitempty"`
	ShouldReceiveOrderStatusInfo  *bool                  `json:"shouldReceiveOrderStatusInfo,omitempty"`
	PersonalDataConsentFrom       *string                `json:"personalDataConsentFrom,omitempty"`
	PersonalDataConsentTo         *string                `json:"personalDataConsentTo,omitempty"`
	PersonalDataProcessingFrom    *string                `json:"personalDataProcessingFrom,omitempty"`
	PersonalDataProcessingTo      *string                `json:"personalDataProcessingTo,omitempty"`
	IsDeleted                     *bool                  `json:"isDeleted,omitempty"`
}

type CustomerCreateOrUpdateModel struct {
	ID string `json:"id"`
}

type CustomerProgramAddResponse struct {
	UserWalletID string `json:"userWalletId"`
}

type WalletHoldResponse struct {
	TransactionID string `json:"transactionId"`
}

// TypeRCI enum для типов идентификаторов клиентов
type TypeRCI string

const (
	TypeRCIPhone      TypeRCI = "phone"
	TypeRCICardTrack  TypeRCI = "cardTrack"
	TypeRCICardNumber TypeRCI = "cardNumber"
	TypeRCIEmail      TypeRCI = "email"
	TypeRCIID         TypeRCI = "id"
)

// Notifications/Commands/WebHook models
type BaseStatusModel struct {
	State     string                 `json:"state"`
	Exception *BaseStatusExceptModel `json:"exception,omitempty"`
}

type BaseStatusExceptModel struct {
	Message *string `json:"message,omitempty"`
}

// WebHook models
type ExternalDataModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LoyaltyInfoModel struct {
	Coupon                  *string  `json:"coupon,omitempty"`
	AppliedManualConditions []string `json:"appliedManualConditions,omitempty"`
}

type WHDeliveryOrderModel struct {
	CreatedDeliveryOrderModel
	WhenCookingCompleted   *string             `json:"whenCookingCompleted,omitempty"`
	MovedToDeliveryID      *string             `json:"movedToDeliveryId,omitempty"`
	MovedToTerminalGroupID *string             `json:"movedToTerminalGroupId,omitempty"`
	MovedToOrganizationID  *string             `json:"movedToOrganizationId,omitempty"`
	MenuID                 *string             `json:"menuId,omitempty"`
	DeliveryZone           *string             `json:"deliveryZone,omitempty"`
	EstimatedTime          *string             `json:"estimatedTime,omitempty"`
	IsAsap                 *bool               `json:"isAsap,omitempty"`
	WhenPacked             *string             `json:"whenPacked,omitempty"`
	LoyaltyInfo            *LoyaltyInfoModel   `json:"loyaltyInfo,omitempty"`
	ExternalData           []ExternalDataModel `json:"externalData,omitempty"`
}

type EventInfoModel struct {
	ID             string                `json:"id"`
	PosID          *string               `json:"posId,omitempty"`
	ExternalNumber *string               `json:"externalNumber,omitempty"`
	OrganizationID string                `json:"organizationId"`
	Timestamp      int64                 `json:"timestamp"`
	CreationStatus string                `json:"creationStatus"`
	ErrorInfo      *ErrorInfoModel       `json:"errorInfo,omitempty"`
	Order          *WHDeliveryOrderModel `json:"order,omitempty"`
}

type WebHookDeliveryOrderEventInfoModel struct {
	EventType      string          `json:"eventType"`
	EventTime      *string         `json:"eventTime,omitempty"`
	OrganizationID string          `json:"organizationId"`
	CorrelationID  string          `json:"correlationId"`
	EventInfo      *EventInfoModel `json:"eventInfo,omitempty"`
}

// Employees models
type CourierItemModel struct {
	Employee EmployeeModel `json:"employee"`
	Phone    *string       `json:"phone,omitempty"`
}

type CouriersByOrganizationModel struct {
	OrganizationID string             `json:"organizationId"`
	Items          []CourierItemModel `json:"items"`
}

type BaseCouriersModel struct {
	BaseResponseModel
	Employees []CouriersByOrganizationModel `json:"employees,omitempty"`
}

type EmployeeInfoModel struct {
	Employee EmployeeModel `json:"employee"`
}

type BaseEmployeeInfoModel struct {
	BaseResponseModel
	EmployeeInfo EmployeeInfoModel `json:"employeeInfo"`
}

type EmployeeTerminalItemModel struct {
	TerminalGroupID string `json:"terminalGroupId"`
	IsOpen          bool   `json:"isOpen"`
}

type BaseEmployeeTerminalModel struct {
	BaseResponseModel
	EmployeeID string                      `json:"employeeId"`
	Terminals  []EmployeeTerminalItemModel `json:"terminals"`
}
