package identify

import "github.com/LEVI-Tempest/Candle/pkg/types"

func (c *types.CandlestickWrapper) IsYang() bool {
	return c.Close >= c.Open
}
func (c *CandlestickWrapper) UpperShadow() float64 {
	if c.Close >= c.Open {
		return c.High - c.Close
	}
	return c.High - c.Open
}
func (c *CandlestickWrapper) LowerShadow() float64 {
	if c.Close >= c.Open {
		return c.Open - c.Low
	}
	return c.Close - c.Low
}

// Umbrella
// 1. No Upper Shadow(or tinny Upper Shadow)
// 2. Lower Shadow is greater than 2x Volume

func Umbrella(c CandlestickWrapper) bool {
	if c.UpperShadow() > 0 {
		return false
	}

	return c.LowerShadow() > 2*c.Volume
}

func Hammer(cs []CandlestickWrapper) bool {
	return false
}
