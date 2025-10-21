package goiikoapi

import (
	"context"
	"encoding/json"
)

// Orders содержит методы для работы с заказами
type Orders struct {
	client *Client
}

// Проверяем, что Orders реализует IOrders
var _ IOrders = (*Orders)(nil)

// OrderCreate реплицирует Orders.order_create
func (o *Orders) OrderCreate(ctx context.Context, organizationID, terminalGroupID string, order map[string]any, createOrderSettings *int) (*BaseCreatedOrderInfoModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
		"terminalGroupId": terminalGroupID,
		"order": order,
	}
	if createOrderSettings != nil {
		data["createOrderSettings"] = *createOrderSettings
	}
	body, status, err := o.client.post(ctx, "/api/1/order/create", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseCreatedOrderInfoModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// OrderByID реплицирует Orders.order_by_id
func (o *Orders) OrderByID(ctx context.Context, organizationIDs []string, orderIDs, posOrderIDs, returnExternalDataKeys, sourceKeys []string) (*ByIdModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationIds": organizationIDs,
		"orderIds": orderIDs,
	}
	if len(posOrderIDs) > 0 { data["posOrderIds"] = posOrderIDs }
	if len(returnExternalDataKeys) > 0 { data["returnExternalDataKeys"] = returnExternalDataKeys }
	if len(sourceKeys) > 0 { data["sourceKeys"] = sourceKeys }
	
	body, status, err := o.client.post(ctx, "/api/1/order/by_id", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out ByIdModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
