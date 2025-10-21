package goiikoapi

import "context"

// Dictionaries содержит методы для работы со словарями
type Dictionaries struct {
	client *Client
}

// Проверяем, что Dictionaries реализует IDictionaries
var _ IDictionaries = (*Dictionaries)(nil)

// OrderTypes реплицирует Dictionaries.order_types
func (d *Dictionaries) OrderTypes(ctx context.Context, organizationIDs []string) (*BaseOrderTypesModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	body, status, err := d.client.post(ctx, "/api/1/deliveries/order_types", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseOrderTypesModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// PaymentTypes реплицирует Dictionaries.payment_types
func (d *Dictionaries) PaymentTypes(ctx context.Context, organizationIDs []string) (*BasePaymentTypesModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	body, status, err := d.client.post(ctx, "/api/1/payment_types", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BasePaymentTypesModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// Discounts реплицирует Dictionaries.discounts
func (d *Dictionaries) Discounts(ctx context.Context, organizationIDs []string) (*BaseDiscountsModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	body, status, err := d.client.post(ctx, "/api/1/discounts", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseDiscountsModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CancelCauses реплицирует Dictionaries.cancel_causes
func (d *Dictionaries) CancelCauses(ctx context.Context, organizationIDs []string) (*BaseCancelCausesModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	body, status, err := d.client.post(ctx, "/api/1/cancel_causes", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseCancelCausesModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// RemovalTypes реплицирует Dictionaries.removal_types
func (d *Dictionaries) RemovalTypes(ctx context.Context, organizationIDs []string) (*BaseRemovalTypesModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	body, status, err := d.client.post(ctx, "/api/1/removal_types", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseRemovalTypesModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// TipsTypes реплицирует Dictionaries.tips_types
func (d *Dictionaries) TipsTypes(ctx context.Context) (*BaseTipsTypesModel, *CustomErrorModel, error) {
	body, status, err := d.client.post(ctx, "/api/1/tips_types", map[string]any{})
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseTipsTypesModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
