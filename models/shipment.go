package models

import (
	"errors"
)

/*ShipmentRequest represents the data needed to create a shipping label using Shippo API*/
type ShipmentRequest struct {
	Length       float64 `json:"length"`
	Width        float64 `json:"width"`
	Height       float64 `json:"height"`
	DistanceUnit string  `json:"distanceUnit"`
}

/*Dimensions represent the data needed to get the correct sized shipment*/
type Dimensions struct {
	Quantity     uint64  `json:"quantity" binding:"required"`
	OzInBag      float64 `json:"ozInBag" binding:"required"`
	Length       float64 `json:"length" binding:"required"`
	Width        float64 `json:"width" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	DistanceUnit string  `json:"distanceUnit" binding:"required"`
	MassUnit     string  `json:"massUnit" binding:"required"`
}

/*NewDimensions creates a new Dimensions struct with the specified size and weight*/
func NewDimensions(quantity uint64, ozInBag float64, req *ShipmentRequest) (*Dimensions, error) {
	dUnit, ok := toDistanceUnit(req.DistanceUnit)
	if !ok {
		return nil, errors.New("Invalid distance unit")
	}
	mUnit, ok := toMassUnit("oz")
	if !ok {
		return nil, errors.New("Invalid mass unit")
	}
	return &Dimensions{
		Quantity:     quantity,
		OzInBag:      ozInBag,
		Length:       req.Length,
		Width:        req.Width,
		Height:       req.Height,
		DistanceUnit: string(dUnit),
		MassUnit:     string(mUnit),
	}, nil
}

func toDistanceUnit(s string) (DistanceUnit, bool) {
	switch s {
	case CENTIMETER:
		return CENTIMETER, true
	case MILLIMETER:
		return MILLIMETER, true
	case METER:
		return METER, true
	case INCH:
		return INCH, true
	case FOOT:
		return FOOT, true
	case YARD:
		return YARD, true
	default:
		return "", false
	}
}

/*DistanceUnit is an enum wrapper for valid content type*/
type DistanceUnit string

/*Valid distance units Shippo accepts*/
const (
	CENTIMETER = "cm"
	MILLIMETER = "mm"
	METER      = "m"
	INCH       = "in"
	FOOT       = "ft"
	YARD       = "yd"
)

func toMassUnit(s string) (MassUnit, bool) {
	switch s {
	case GRAM:
		return GRAM, true
	case OUNCE:
		return OUNCE, true
	case POUND:
		return POUND, true
	case KILOGRAM:
		return KILOGRAM, true
	default:
		return "", false
	}
}

/*DistanceUnit is an enum wrapper for valid content type*/
type MassUnit string

/*valid MassUnits that Shippo accepts*/
const (
	GRAM     = "g"
	OUNCE    = "oz"
	POUND    = "lb"
	KILOGRAM = "kg"
)
