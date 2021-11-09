package service

import (
	"context"

	"template/internal/model"

	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/types/time"
)

type MyService struct{}

func NewMyService() *MyService {
	return &MyService{}
}

func (cs *MyService) GetEntity(ctx context.Context, id string) (model.MyEntity, error) {
	var err error = nil
	return model.MyEntity{Name: "sample", ExampleTime: time.ShortTime{}}, err
}
