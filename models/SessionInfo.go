package models

type SessionInfo struct {
	Cookies           map[string]string `json:"cookies"`
	Proxy             string            `json:"proxy"`
	TargetHostCountry string            `json:"targetHostCountry"`
	DeliveryCountry   string            `json:"deliveryCountry"`
	UserAgent         string            `json:"user_agent"`
	Code              int               `json:"code"`
}
