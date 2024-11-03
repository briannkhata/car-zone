package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"noOfCylinders"`
	CarRange      int64     `json:"carRange"`
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement must be greater than zero")
	}
	return nil
}

func validateNoOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("NoOfCylinders must be greater than zero")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}
