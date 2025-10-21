package goiikoapi

import (
	"encoding/json"
	"errors"
)

// WebHook содержит утилиты для работы с webhook'ами
type WebHook struct{}

// Проверяем, что WebHook реализует IWebHook
var _ IWebHook = (*WebHook)(nil)

// ParseWebhookOrder реплицирует WebHook.parse_webhook_order
func (wh *WebHook) ParseWebhookOrder(data []map[string]any) ([]WebHookDeliveryOrderEventInfoModel, error) {
	var result []WebHookDeliveryOrderEventInfoModel
	for _, item := range data {
		var wh WebHookDeliveryOrderEventInfoModel
		jsonData, err := json.Marshal(item)
		if err != nil { return nil, err }
		if err := json.Unmarshal(jsonData, &wh); err != nil { return nil, err }
		result = append(result, wh)
	}
	return result, nil
}

// ParseWebhookReserve реплицирует WebHook.parse_webhook_reserve (заглушка)
func (wh *WebHook) ParseWebhookReserve(data []map[string]any) ([]WebHookDeliveryOrderEventInfoModel, error) {
	// В Python версии это FutureWarning, поэтому возвращаем ошибку
	return nil, errors.New("parse_webhook_reserve: в разработке")
}

// Глобальные функции для совместимости с Python API
// ParseWebhookOrder реплицирует WebHook.parse_webhook_order
func ParseWebhookOrder(data []map[string]any) ([]WebHookDeliveryOrderEventInfoModel, error) {
	wh := &WebHook{}
	return wh.ParseWebhookOrder(data)
}

// ParseWebhookReserve реплицирует WebHook.parse_webhook_reserve (заглушка)
func ParseWebhookReserve(data []map[string]any) ([]WebHookDeliveryOrderEventInfoModel, error) {
	wh := &WebHook{}
	return wh.ParseWebhookReserve(data)
}
