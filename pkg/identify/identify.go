package identify

// Umbrella
// 伞形线
// 1. No Upper Shadow(or supper tinny Upper Shadow)
// 2. Lower Shadow is greater than 2x～3x Volume
// Note: also indicates Bottom Support or Top Resistance
func Umbrella(cs []CandlestickWrapper) bool {
	c := cs[0]
	body := c.Body()
	if c.UpperShadow() > 0 || c.UpperShadow() > 0.2*body {
		return false
	}

	return c.LowerShadow() > 2.5*body
}

// Hammer
// 锤子线
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

// EngulfingPattern
// 抱线形态
func EngulfingPattern() bool {
	return false
}

// BullishEngulfingPattern
// 看涨抱线形态
func BullishEngulfingPattern() bool {
	if !EngulfingPattern() {
		return false
	}

	return false
}

// BearishEngulfingPattern
// 看跌抱线形态
func BearishEngulfingPattern() bool {
	if !EngulfingPattern() {
		return false
	}

	return false
}

// DarkCloudCover
// 乌云盖顶
func DarkCloudCover() bool {
	return false
}

// PiercingPattern
// 刺透形态
func PiercingPattern() bool {
	return false
}
