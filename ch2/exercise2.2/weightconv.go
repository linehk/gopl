package main

import "fmt"

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }

// KToP converts Kilogram to Pound
func KToP(k Kilogram) Pound { return Pound(k * 2.2045) }

// PToK converts Pound to Kilogram
func PToK(p Pound) Kilogram { return Kilogram(p / 2.2045) }
