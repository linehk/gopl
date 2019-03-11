package main

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%gM", m) }
func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }

// FToM converts Meter to Foot
func FToM(f Foot) Meter { return Meter(f / 3.2808) }

// MToF converts Foot to Meter
func MToF(m Meter) Foot { return Foot(m * 3.2808) }
