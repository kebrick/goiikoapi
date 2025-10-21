package goiikoapi

import "context"

// Deliveries содержит методы для работы с доставкой
type Deliveries struct {
	client *Client
}

// Проверяем, что Deliveries реализует IDeliveries
var _ IDeliveries = (*Deliveries)(nil)

// DeliveryCreate реплицирует Deliveries.delivery_create
func (d *Deliveries) DeliveryCreate(ctx context.Context, organizationID string, order map[string]any, terminalGroupID *string, createOrderSettings *int) (*BaseCreatedDeliveryOrderInfoModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
		"order": order,
	}
	if terminalGroupID != nil {
		data["terminalGroupId"] = *terminalGroupID
	}
	if createOrderSettings != nil {
		data["createOrderSettings"] = map[string]int{"transportToFrontTimeout": *createOrderSettings}
	}
	body, status, err := d.client.post(ctx, "/api/1/deliveries/create", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseCreatedDeliveryOrderInfoModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// UpdateOrderDeliveryStatus реплицирует Deliveries.update_order_delivery_status
func (d *Deliveries) UpdateOrderDeliveryStatus(ctx context.Context, organizationIDs []string, orderID, deliveryStatus, deliveryDate string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationIds": organizationIDs,
		"orderId": orderID,
		"deliveryStatus": deliveryStatus,
	}
	if deliveryStatus == "Delivered" && deliveryDate != "" {
		data["deliveryDate"] = deliveryDate
	}
	body, status, err := d.client.post(ctx, "/api/1/deliveries/update_order_delivery_status", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// Confirm реплицирует Deliveries.confirm
func (d *Deliveries) Confirm(ctx context.Context, organizationIDs []string, orderID string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationIds": organizationIDs,
		"orderId": orderID,
	}
	body, status, err := d.client.post(ctx, "/api/1/deliveries/confirm", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// CancelConfirmation реплицирует Deliveries.cancel_confirmation
func (d *Deliveries) CancelConfirmation(ctx context.Context, organizationIDs []string, orderID string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationIds": organizationIDs,
		"orderId": orderID,
	}
	body, status, err := d.client.post(ctx, "/api/1/deliveries/cancel_confirmation", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// ByDeliveryDateAndStatus реплицирует Deliveries.by_delivery_date_and_status
func (d *Deliveries) ByDeliveryDateAndStatus(ctx context.Context, organizationIDs []string, deliveryDateFrom, deliveryDateTo string, statuses, sourceKeys []string) (*ByDeliveryDateAndStatusModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationIds": organizationIDs,
		"deliveryDateFrom": deliveryDateFrom,
	}
	if deliveryDateTo != "" { data["deliveryDateTo"] = deliveryDateTo }
	if len(statuses) > 0 { data["statuses"] = statuses }
	if len(sourceKeys) > 0 { data["sourceKeys"] = sourceKeys }
	
	body, status, err := d.client.post(ctx, "/api/1/deliveries/by_delivery_date_and_status", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out ByDeliveryDateAndStatusModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
