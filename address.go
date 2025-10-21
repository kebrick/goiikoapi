package goiikoapi

import "context"

// Address содержит методы для работы с адресами
type Address struct {
	client *Client
}

// Проверяем, что Address реализует IAddress
var _ IAddress = (*Address)(nil)

// Regions реплицирует Address.regions
func (a *Address) Regions(ctx context.Context, organizationIDs []string) (*BaseRegionsModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	body, status, err := a.client.post(ctx, "/api/1/regions", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseRegionsModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// Cities реплицирует Address.cities
func (a *Address) Cities(ctx context.Context, organizationIDs []string) (*BaseCitiesModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	body, status, err := a.client.post(ctx, "/api/1/cities", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseCitiesModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// StreetsByCity реплицирует Address.by_city
func (a *Address) StreetsByCity(ctx context.Context, organizationID, cityID string) (*BaseStreetByCityModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
		"cityId": cityID,
	}
	body, status, err := a.client.post(ctx, "/api/1/streets/by_city", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseStreetByCityModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
