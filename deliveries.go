package goiikoapi

import (
	"context"
	"encoding/json"
)

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
func (d *Deliveries) Confirm(ctx context.Context, organizationID string, orderID string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
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

// ByDeliveryDateAndSourceKeyAndFilter реплицирует Deliveries.by_delivery_date_and_source_key_and_filter
func (d *Deliveries) ByDeliveryDateAndSourceKeyAndFilter(ctx context.Context, organizationIDs []string, terminalGroupIDs []string, deliveryDateFrom, deliveryDateTo *string, statuses []string, hasProblem *bool, orderServiceType, searchText *string, timeToCookingErrorTimeout, cookingTimeout *int, sortProperty, sortDirection *string, rowsCount *int, sourceKeys, orderIDs []string) (*ByDeliveryDateAndSourceKeyAndFilter, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationIds": organizationIDs,
	}
	if len(terminalGroupIDs) > 0 {
		data["terminalGroupIds"] = terminalGroupIDs
	}
	if deliveryDateFrom != nil {
		data["deliveryDateFrom"] = *deliveryDateFrom
	}
	if deliveryDateTo != nil {
		data["deliveryDateTo"] = *deliveryDateTo
	}
	if len(statuses) > 0 {
		data["statuses"] = statuses
	}
	if hasProblem != nil {
		data["hasProblem"] = *hasProblem
	}
	if orderServiceType != nil {
		data["orderServiceType"] = *orderServiceType
	}
	if searchText != nil {
		data["searchText"] = *searchText
	}
	if timeToCookingErrorTimeout != nil {
		data["timeToCookingErrorTimeout"] = *timeToCookingErrorTimeout
	}
	if cookingTimeout != nil {
		data["cookingTimeout"] = *cookingTimeout
	}
	if sortProperty != nil {
		data["sortProperty"] = *sortProperty
	}
	if sortDirection != nil {
		data["sortDirection"] = *sortDirection
	}
	if rowsCount != nil {
		data["rowsCount"] = *rowsCount
	}
	if len(sourceKeys) > 0 {
		data["sourceKeys"] = sourceKeys
	}
	if len(orderIDs) > 0 {
		data["orderIds"] = orderIDs
	}
	
	body, status, err := d.client.post(ctx, "/api/1/deliveries/by_delivery_date_and_source_key_and_filter", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out ByDeliveryDateAndSourceKeyAndFilter
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
