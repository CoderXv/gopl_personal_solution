package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbSoluteZeroC Celsius = -273.15
	Freezing      Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g˚C", c) }
func (f Fahrenheit) String() string {return fmt.Sprintf("%g˚F", f) }
func (k Kelvins) String() string {return fmt.Sprintf("%gK", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvins { return Kelvins(c + 273.15) }
