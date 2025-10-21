package goiikoapi

import (
	"context"
	"encoding/json"
)

// Menu содержит методы для работы с меню
type Menu struct {
	client *Client
}

// Проверяем, что Menu реализует IMenu
var _ IMenu = (*Menu)(nil)

// Nomenclature реплицирует Menu.nomenclature
func (m *Menu) Nomenclature(ctx context.Context, organizationID string, startRevision *int) (*BaseNomenclatureModel, *CustomErrorModel, error) {
	data := map[string]any{"organizationId": organizationID}
	if startRevision != nil {
		data["startRevision"] = *startRevision
	}
	body, status, err := m.client.post(ctx, "/api/1/nomenclature", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseNomenclatureModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// Menu реплицирует Menu.menu
func (m *Menu) Menu(ctx context.Context) (*BaseMenuModel, *CustomErrorModel, error) {
	body, status, err := m.client.post(ctx, "/api/2/menu", map[string]any{})
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseMenuModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// MenuByID реплицирует Menu.menu_by_id
func (m *Menu) MenuByID(ctx context.Context, externalMenuID string, organizationIDs []string, priceCategoryID *string) (*BaseMenuByIdModel, *CustomErrorModel, error) {
	data := map[string]any{
		"externalMenuId": externalMenuID,
		"organizationIds": organizationIDs,
	}
	if priceCategoryID != nil {
		data["priceCategoryId"] = *priceCategoryID
	}
	body, status, err := m.client.post(ctx, "/api/2/menu/by_id", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseMenuByIdModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
