package identify

import (
	v1 "github.com/LEVI-Tempest/Candle/pkg/proto"
	"github.com/LEVI-Tempest/Candle/pkg/utils"
)

type CandlestickWrapper struct {
	*v1.Candlestick
}

type Trend string

const (
	TrendYin     Trend = "Yin"
	TrendYang    Trend = "Yang"
	TrendMiddle  Trend = "Middle"
	TrendUnknown Trend = "Unknown"
)

// Body returns the body of the candlestick
// always >= 0
func (c *CandlestickWrapper) Body() float64 {
	return utils.Abs(c.Close - c.Open)
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

func (c *CandlestickWrapper) Yang() bool {
	return c.Close >= c.Open
}

func (c *CandlestickWrapper) Yin() bool {
	return !c.Yang()
}

// DetermineTrend returns true if the candlestick is a long term yang
func DetermineTrend(cs []CandlestickWrapper, days int) Trend {
	if len(cs) <= days {
		return TrendUnknown
	}

	// Yang:
	// Count the number of days when closing price is rising in the last days[]
	risingCount := 0
	for i := len(cs) - days; i < len(cs)-1; i++ {
		if cs[i+1].Close > cs[i].Close {
			risingCount++
		}
	}

	// Determine the trend based on the majority
	if risingCount > days*3/5 {
		return TrendYang
	}

	return TrendYin
}
