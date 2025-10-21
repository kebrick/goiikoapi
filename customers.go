package goiikoapi

import "context"

// Customers содержит методы для работы с клиентами
type Customers struct {
	client *Client
}

// Проверяем, что Customers реализует ICustomers
var _ ICustomers = (*Customers)(nil)

// CustomerInfo реплицирует Customers.customer_info
func (c *Customers) CustomerInfo(ctx context.Context, organizationID, identifier string, identifierType TypeRCI) (*CustomerInfoModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
		"type": identifierType,
	}
	// добавляем идентификатор в зависимости от типа
	switch identifierType {
	case TypeRCIPhone:
		data["phone"] = identifier
	case TypeRCICardTrack:
		data["cardTrack"] = identifier
	case TypeRCICardNumber:
		data["cardNumber"] = identifier
	case TypeRCIEmail:
		data["email"] = identifier
	case TypeRCIID:
		data["id"] = identifier
	}
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/info", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out CustomerInfoModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerCreateOrUpdate реплицирует Customers.customer_create_or_update
func (c *Customers) CustomerCreateOrUpdate(ctx context.Context, organizationID string, opts ...CustomerCreateOrUpdateOption) (*CustomerCreateOrUpdateModel, *CustomErrorModel, error) {
	data := map[string]any{"organizationId": organizationID}
	
	// применяем опции
	for _, opt := range opts {
		opt(data)
	}
	
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/create_or_update", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out CustomerCreateOrUpdateModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerProgramAdd реплицирует Customers.customer_program_add
func (c *Customers) CustomerProgramAdd(ctx context.Context, customerID, programID, organizationID string) (*CustomerProgramAddResponse, *CustomErrorModel, error) {
	data := map[string]any{
		"customerId": customerID,
		"programId": programID,
		"organizationId": organizationID,
	}
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/program/add", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out CustomerProgramAddResponse
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerCardAdd реплицирует Customers.customer_card_add
func (c *Customers) CustomerCardAdd(ctx context.Context, customerID, cardTrack, cardNumber, organizationID string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"customerId": customerID,
		"cardTrack": cardTrack,
		"cardNumber": cardNumber,
		"organizationId": organizationID,
	}
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/card/add", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerCardDelete реплицирует Customers.customer_card_delete
func (c *Customers) CustomerCardDelete(ctx context.Context, customerID, cardTrack, organizationID string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"customerId": customerID,
		"cardTrack": cardTrack,
		"organizationId": organizationID,
	}
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/card/remove", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerWalletHold реплицирует Customers.customer_wallet_hold
func (c *Customers) CustomerWalletHold(ctx context.Context, customerID, walletID, organizationID string, sum float64, transactionID, comment *string) (*WalletHoldResponse, *CustomErrorModel, error) {
	data := map[string]any{
		"customerId": customerID,
		"walletId": walletID,
		"sum": sum,
		"organizationId": organizationID,
	}
	if transactionID != nil { data["transactionId"] = *transactionID }
	if comment != nil { data["comment"] = *comment }
	
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/wallet/hold", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out WalletHoldResponse
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerWalletCancelHold реплицирует Customers.customer_wallet_cancel_hold
func (c *Customers) CustomerWalletCancelHold(ctx context.Context, organizationID, transactionID string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
		"transactionId": transactionID,
	}
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/wallet/cancel_hold", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerWalletTopup реплицирует Customers.customer_wallet_topup
func (c *Customers) CustomerWalletTopup(ctx context.Context, customerID, walletID, organizationID string, sum float64, comment *string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"customerId": customerID,
		"walletId": walletID,
		"sum": sum,
		"organizationId": organizationID,
	}
	if comment != nil { data["comment"] = *comment }
	
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/wallet/topup", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CustomerWalletChargeoff реплицирует Customers.customer_wallet_chargeoff
func (c *Customers) CustomerWalletChargeoff(ctx context.Context, customerID, walletID, organizationID string, sum float64, comment *string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"customerId": customerID,
		"walletId": walletID,
		"sum": sum,
		"organizationId": organizationID,
	}
	if comment != nil { data["comment"] = *comment }
	
	body, status, err := c.client.post(ctx, "/api/1/loyalty/iiko/customer/wallet/chargeoff", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
