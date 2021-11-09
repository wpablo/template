package model

import (
	"gitlab.ar.bsch/santander-tecnologia/santander-go-kit/types/time"
)

type MyEntity struct {
	Name string `json:"name"`

	ExampleTime time.ShortTime `json:"time"`
}
