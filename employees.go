package goiikoapi

import (
	"context"
	"encoding/json"
)

// Employees содержит методы для работы с сотрудниками
type Employees struct {
	client *Client
}

// Проверяем, что Employees реализует IEmployees
var _ IEmployees = (*Employees)(nil)

// Couriers реплицирует Employees.couriers
func (e *Employees) Couriers(ctx context.Context, organizationIDs []string) (*BaseCouriersModel, *CustomErrorModel, error) {
	if len(organizationIDs) == 0 {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: "пустой список id организаций"}}, nil
	}
	data := map[string]any{
		"organizationIds": organizationIDs,
	}
	body, status, err := e.client.post(ctx, "/api/1/employees/couriers", data)
	if err != nil {
		return nil, nil, err
	}
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseCouriersModel
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, nil, err
	}
	return &out, nil, nil
}

// EmployeeInfo реплицирует Employees.employees_info
func (e *Employees) EmployeeInfo(ctx context.Context, organizationID, id string) (*BaseEmployeeInfoModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId": organizationID,
		"id":             id,
	}
	body, status, err := e.client.post(ctx, "/api/1/employees/info", data)
	if err != nil {
		return nil, nil, err
	}
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseEmployeeInfoModel
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, nil, err
	}
	return &out, nil, nil
}

// ShiftClockin реплицирует Employees.employees_shift_clockin
// Open personal session.
// This method is a command. Use api/1/commands/status method to get the progress status.
// Restriction group: Employees: shifts.
func (e *Employees) ShiftClockin(ctx context.Context, organizationID, terminalGroupID, employeeID string, roleID *string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId":  organizationID,
		"terminalGroupId": terminalGroupID,
		"employeeId":      employeeID,
	}
	if roleID != nil {
		data["roleId"] = *roleID
	}
	body, status, err := e.client.post(ctx, "/api/1/employees/shift/clockin", data)
	if err != nil {
		return nil, nil, err
	}
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, nil, err
	}
	return &out, nil, nil
}

// ShiftClockout реплицирует Employees.employees_shift_clockout
// Close personal session.
// This method is a command. Use api/1/commands/status method to get the progress status.
// Restriction group: Employees: shifts.
func (e *Employees) ShiftClockout(ctx context.Context, organizationID, terminalGroupID, employeeID string) (*BaseResponseModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId":  organizationID,
		"terminalGroupId": terminalGroupID,
		"employeeId":      employeeID,
	}
	body, status, err := e.client.post(ctx, "/api/1/employees/shift/clockout", data)
	if err != nil {
		return nil, nil, err
	}
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseResponseModel
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, nil, err
	}
	return &out, nil, nil
}

// ShiftIsOpen реплицирует Employees.employees_shift_is_open
// Check if personal session is open.
func (e *Employees) ShiftIsOpen(ctx context.Context, organizationID, terminalGroupID, employeeID string) (*BaseEmployeeInfoModel, *CustomErrorModel, error) {
	data := map[string]any{
		"organizationId":  organizationID,
		"terminalGroupId": terminalGroupID,
		"employeeId":      employeeID,
	}
	body, status, err := e.client.post(ctx, "/api/1/employees/shift/is_open", data)
	if err != nil {
		return nil, nil, err
	}
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseEmployeeInfoModel
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, nil, err
	}
	return &out, nil, nil
}

// ShiftByCourier реплицирует Employees.employees_shift_by_courier
// Get terminal groups where employee session is opened.
func (e *Employees) ShiftByCourier(ctx context.Context, employeeID string) (*BaseEmployeeTerminalModel, *CustomErrorModel, error) {
	data := map[string]any{
		"employeeId": employeeID,
	}
	body, status, err := e.client.post(ctx, "/api/1/employees/shift/by_courier", data)
	if err != nil {
		return nil, nil, err
	}
	if msg, ok := DetectAPIError(body); ok {
		return nil, &CustomErrorModel{ErrorModel: ErrorModel{ErrorDescription: msg}, StatusCode: status}, nil
	}
	var out BaseEmployeeTerminalModel
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, nil, err
	}
	return &out, nil, nil
}
