package service

import "context"

// Service -
type Service interface {
	// CheckMacValue
	GenerateCheckMacValue(ctx context.Context, params map[string]interface{}) string
}

type service struct {
}

// NewService -
func NewService() Service {
	return &service{}
}
