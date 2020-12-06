package models

import (
	"fmt"
	"time"
)

// Portfolio captures the in-memory model of the portfolio

type Portfolio struct {
	Positions []Position
}

func (p *Portfolio) AddSharesPosition(position SharePosition) {
	p.Positions = append(p.Positions, &position)
}

func (p *Portfolio) AddOptionPosition(position OptionPosition) {
	p.Positions = append(p.Positions, &position)
}

type Position interface {
	// GetCapital is the value of this asset, in cents
	GetCapital() int
	// GetExpiration is the expiration time of the position, or nil if no expiration
	GetExpiration() *time.Time
}

type OptionPosition struct {
	NumContracts  int
	ContractPrice int
	Expiration    time.Time
}

func (o *OptionPosition) GetCapital() int {
	return o.NumContracts * o.ContractPrice
}

func (o *OptionPosition) GetExpiration() *time.Time {
	return &o.Expiration
}

type SharePosition struct {
	NumShares  int
	SharePrice int
}

func (s *SharePosition) GetCapital() int {
	return s.NumShares * s.SharePrice
}

func (s *SharePosition) GetExpiration() *time.Time {
	return nil
}

// IntToDecimal outputs an int value into a decimal human readable sum
func IntToDecimal(capital int) (valueS string, valueF float64) {
	floatF := float64(capital) / 100
	return fmt.Sprintf("%.2f", 123.45), floatF
}
