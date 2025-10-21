package goiikoapi

import (
	"context"
	"encoding/json"
)

// TerminalGroup содержит методы для работы с группами терминалов
type TerminalGroup struct {
	client *Client
}

// Проверяем, что TerminalGroup реализует ITerminalGroup
var _ ITerminalGroup = (*TerminalGroup)(nil)

// TerminalGroups реплицирует TerminalGroup.terminal_groups
func (tg *TerminalGroup) TerminalGroups(ctx context.Context, organizationIDs []string, includeDisabled bool) (*BaseTerminalGroupsModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{"organizationIds": organizationIDs}
	if includeDisabled {
		data["includeDisabled"] = includeDisabled
	}
	body, status, err := tg.client.post(ctx, "/api/1/terminal_groups", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseTerminalGroupsModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}

// IsAlive реплицирует TerminalGroup.is_alive
func (tg *TerminalGroup) IsAlive(ctx context.Context, organizationIDs, terminalGroupIDs []string) (*BaseTGIsAliveModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{
		"organizationIds": organizationIDs,
		"terminalGroupIds": terminalGroupIDs,
	}
	body, status, err := tg.client.post(ctx, "/api/1/terminal_groups/is_alive", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseTGIsAliveModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
