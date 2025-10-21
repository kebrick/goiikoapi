package goiikoapi

import "context"

// Notifications содержит методы для работы с уведомлениями
type Notifications struct {
	client *Client
}

// Проверяем, что Notifications реализует INotifications
var _ INotifications = (*Notifications)(nil)

// Send реплицирует Notifications.send
func (n *Notifications) Send(ctx context.Context, orderSource, orderID, additionalInfo, organizationID, messageType string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"orderSource": orderSource,
		"orderId": orderID,
		"additionalInfo": additionalInfo,
		"messageType": messageType,
		"organizationId": organizationID,
	}
	body, status, err := n.client.post(ctx, "/api/1/notifications/send", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	return &out, nil, nil
}
