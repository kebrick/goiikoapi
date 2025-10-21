package goiikoapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"
)

// Client реализует базовый доступ к iiko cloud API
type Client struct {
	apiLogin    string
	baseURL     string
	returnDict  bool
	debug       bool
	client      *http.Client
	defaultTO   time.Duration
	headers     http.Header
	mu          sync.RWMutex
	token       string
	tokenAt     time.Time
	lastDataRaw []byte

	organizationsIDs []string

	// Группы методов как в Python проекте
	Dictionaries  *Dictionaries
	Menu          *Menu
	Orders        *Orders
	Deliveries    *Deliveries
	Address       *Address
	TerminalGroup *TerminalGroup
	Customers     *Customers
	Notifications *Notifications
	Commands      *Commands
	WebHook       *WebHook
}

// Проверяем, что Client реализует IClient
var _ IClient = (*Client)(nil)

const (
	defaultBaseURL = "https://api-ru.iiko.services"
	defaultTimeout = 15 * time.Second
	tokenTTL       = 15 * time.Minute
)

// NewClient создает новый клиент. Если workingToken задан, не запрашивает новый токен сразу
func NewClient(apiLogin string, opts ...Option) (*Client, error) {
	c := &Client{
		apiLogin:  apiLogin,
		baseURL:   defaultBaseURL,
		client:    &http.Client{Timeout: defaultTimeout},
		defaultTO: defaultTimeout,
		headers:   make(http.Header),
	}
	c.headers.Set("Content-Type", "application/json")
	// кастомные опции
	for _, opt := range opts {
		opt(c)
	}
	// если токена нет, получить
	if c.token == "" {
		if err := c.refreshToken(context.Background()); err != nil {
			return nil, err
		}
	}
	
	// Инициализация групп методов
	c.Dictionaries = &Dictionaries{client: c}
	c.Menu = &Menu{client: c}
	c.Orders = &Orders{client: c}
	c.Deliveries = &Deliveries{client: c}
	c.Address = &Address{client: c}
	c.TerminalGroup = &TerminalGroup{client: c}
	c.Customers = &Customers{client: c}
	c.Notifications = &Notifications{client: c}
	c.Commands = &Commands{client: c}
	c.WebHook = &WebHook{}
	
	return c, nil
}

// Option опции конфигурации клиента
type Option func(*Client)

func WithBaseURL(url string) Option { return func(c *Client) { c.baseURL = url } }
func WithHTTPClient(hc *http.Client) Option { return func(c *Client) { if hc != nil { c.client = hc } } }
func WithTimeout(d time.Duration) Option { return func(c *Client) { c.defaultTO = d; c.client.Timeout = d } }
func WithDebug(debug bool) Option { return func(c *Client) { c.debug = debug } }
func WithReturnDict(v bool) Option { return func(c *Client) { c.returnDict = v } }
func WithWorkingToken(token string) Option {
	return func(c *Client) {
		if token != "" {
			c.setTokenLocked(token)
		}
	}
}

func (c *Client) setTokenLocked(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.token = token
	c.tokenAt = time.Now()
	c.headers.Set("Authorization", "Bearer "+token)
}

func (c *Client) tokenExpired() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.token == "" { return true }
	return time.Since(c.tokenAt) >= tokenTTL
}

func (c *Client) refreshToken(ctx context.Context) error {
	data := map[string]string{"apiLogin": c.apiLogin}
	body, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/1/access_token", bytes.NewReader(body))
	if err != nil { return err }
	req.Header = cloneHeader(c.headers)
	resp, err := c.client.Do(req)
	if err != nil { return err }
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	var out struct{ Token string `json:"token"`; ErrorDescription string `json:"errorDescription"` }
	_ = json.Unmarshal(b, &out)
	if out.ErrorDescription != "" {
		return errors.New(out.ErrorDescription)
	}
	if out.Token == "" {
		return errors.New("empty token in access_token response")
	}
	c.setTokenLocked(out.Token)
	return nil
}

// post выполняет POST с повтором при 401 (обновляет токен и повторяет один раз)
func (c *Client) post(ctx context.Context, url string, payload any) ([]byte, int, error) {
	if c.tokenExpired() {
		if err := c.refreshToken(ctx); err != nil { return nil, 0, err }
	}
	body, err := json.Marshal(payload)
	if err != nil { return nil, 0, err }
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+url, bytes.NewReader(body))
	if err != nil { return nil, 0, err }
	req.Header = cloneHeader(c.headers)
	resp, err := c.client.Do(req)
	if err != nil { return nil, 0, err }
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	// 401 -> обновить токен и один повтор
	if resp.StatusCode == http.StatusUnauthorized {
		if err := c.refreshToken(ctx); err != nil { return nil, resp.StatusCode, err }
		req2, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+url, bytes.NewReader(body))
		if err != nil { return nil, 0, err }
		req2.Header = cloneHeader(c.headers)
		resp2, err := c.client.Do(req2)
		if err != nil { return nil, 0, err }
		defer resp2.Body.Close()
		b2, _ := io.ReadAll(resp2.Body)
		c.mu.Lock(); c.lastDataRaw = b2; c.mu.Unlock()
		return b2, resp2.StatusCode, nil
	}
	c.mu.Lock(); c.lastDataRaw = b; c.mu.Unlock()
	return b, resp.StatusCode, nil
}

func cloneHeader(h http.Header) http.Header {
	cl := make(http.Header, len(h))
	for k, v := range h {
		vv := make([]string, len(v))
		copy(vv, v)
		cl[k] = vv
	}
	return cl
}

// Organizations реплицирует BaseAPI.organizations из Python
func (c *Client) Organizations(ctx context.Context, organizationIDs []string, returnAdditionalInfo, includeDisabled *bool) (*BaseOrganizationsModel, *CustomErrorModel, error) {
	data := map[string]any{}
	if len(organizationIDs) > 0 { data["organizationIds"] = organizationIDs }
	if returnAdditionalInfo != nil { data["returnAdditionalInfo"] = *returnAdditionalInfo }
	if includeDisabled != nil { data["includeDisabled"] = *includeDisabled }

	body, status, err := c.post(ctx, "/api/1/organizations", data)
	if err != nil { return nil, nil, err }
	if msg, ok := DetectAPIError(body); ok {
		cerr := &CustomErrorModel{ErrorModel: ErrorModel{BaseResponseModel: BaseResponseModel{}, ErrorDescription: msg}, StatusCode: status}
		return nil, cerr, nil
	}
	var out BaseOrganizationsModel
	if err := json.Unmarshal(body, &out); err != nil { return nil, nil, err }
	// сохранить ids в клиенте
	c.mu.Lock()
	c.organizationsIDs = out.ListIDs()
	c.mu.Unlock()
	return &out, nil, nil
}

// LastDataRaw возвращает сырое тело последнего ответа
func (c *Client) LastDataRaw() []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastDataRaw
}

// Методы для доступа к группам через интерфейсы
func (c *Client) Dictionaries() IDictionaries { return c.Dictionaries }
func (c *Client) Menu() IMenu { return c.Menu }
func (c *Client) Orders() IOrders { return c.Orders }
func (c *Client) Deliveries() IDeliveries { return c.Deliveries }
func (c *Client) Address() IAddress { return c.Address }
func (c *Client) TerminalGroup() ITerminalGroup { return c.TerminalGroup }
func (c *Client) Customers() ICustomers { return c.Customers }
func (c *Client) Notifications() INotifications { return c.Notifications }
func (c *Client) Commands() ICommands { return c.Commands }
func (c *Client) WebHook() IWebHook { return c.WebHook }
