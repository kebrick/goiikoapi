package goiikoapi

import "context"

// Commands содержит методы для работы с командами
type Commands struct {
	client *Client
}

// Проверяем, что Commands реализует ICommands
var _ ICommands = (*Commands)(nil)

// Status реплицирует Commands.status
func (c *Commands) Status(ctx context.Context, organizationID, correlationID string) (*BaseStatusModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
		"correlationId": correlationID,
	}
	body, status, err := c.client.post(ctx, "/api/1/commands/status", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseStatusModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
