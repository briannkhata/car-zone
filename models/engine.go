package models

import (
	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"noOfCylinders"`
	CarRange      int64     `json:"carRange"`
}
