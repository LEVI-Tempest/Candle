package identify

// Umbrella
// 1. No Upper Shadow(or supper tinny Upper Shadow)
// 2. Lower Shadow is greater than 2x Volume
// Note: also indicates Bottom Support or Top Resistance
func Umbrella(cs []CandlestickWrapper) bool {
	c := cs[0]
	body := c.Body()
	if c.UpperShadow() > 0 || c.UpperShadow() > 0.2*body {
		return false
	}

	return c.LowerShadow() > 2*body
}

func Hammer(cs []CandlestickWrapper) bool {
	if !Umbrella(cs) {
		return false
	}

	// History is not enough
	if len(cs) < 2 {
		return false
	}
	return cs[1].Yang()
}
