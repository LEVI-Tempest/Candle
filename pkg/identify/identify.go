package identify

import (
	"math"

	v1 "github.com/LEVI-Tempest/Candle/pkg/proto"
)

/***
 * @author Tempest
 * @description identify
 * @date 2025/05/18
 * @The code in this section was generated by Google Gemini 2.0 Flash.
 * 	I am responsible for the output of this code.
 * @version 1.0.0
 */

func NewCandlestickWrapper(c *v1.Candlestick) CandlestickWrapper {
	return CandlestickWrapper{Candlestick: c}
}

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

// Doji
// 十字星
// 1. Small body (Open and Close are very close)
// 2. Upper and Lower Shadows are significant
func Doji(cs []CandlestickWrapper) bool {
	if len(cs) < 1 {
		return false
	}
	c := cs[0]
	body := c.Body()
	upperShadow := c.UpperShadow()
	lowerShadow := c.LowerShadow()

	// Body is small relative to the range
	if body > 0.1*math.Abs(c.High-c.Low) { // Adjust 0.1 as needed
		return false
	}

	// Upper and lower shadows are significant
	return upperShadow > 0.2*math.Abs(c.High-c.Low) && lowerShadow > 0.2*math.Abs(c.High-c.Low) // Adjust 0.2 as needed
}

// Marubozu (Bullish and Bearish)
// 光头光脚 (阳线和阴线)
// 1. No upper shadow or very small upper shadow
// 2. No lower shadow or very small lower shadow
// Note: Bullish Marubozu indicates strong buying pressure, Bearish Marubozu indicates strong selling pressure.
func Marubozu(cs []CandlestickWrapper) bool {
	if len(cs) < 1 {
		return false
	}
	c := cs[0]
	body := c.Body()
	upperShadow := c.UpperShadow()
	lowerShadow := c.LowerShadow()
	tolerance := 0.01 * body // Allow for very small shadows

	return upperShadow <= tolerance && lowerShadow <= tolerance
}

// Hammer
// 锤头线
// 1. Small body at the upper end of the trading range
// 2. Long lower shadow (at least 2 times the body)
// 3. Little or no upper shadow
// Note: Appears in a downtrend and suggests a potential bullish reversal.
func Hammer(cs []CandlestickWrapper) bool {
	if len(cs) < 1 {
		return false
	}
	c := cs[0]
	body := c.Body()
	lowerShadow := c.LowerShadow()
	upperShadow := c.UpperShadow()

	if body == 0 { // Avoid division by zero
		return false
	}

	return lowerShadow > 2*body && upperShadow < 0.1*body && c.Close > c.Low && c.Close > c.Open // Body at the upper end
}

// HangingMan
// 吊颈线
// 1. Small body at the upper end of the trading range
// 2. Long lower shadow (at least 2 times the body)
// 3. Little or no upper shadow
// Note: Appears in an uptrend and suggests a potential bearish reversal. The shape is the same as a Hammer, but its context is different.
func HangingMan(cs []CandlestickWrapper) bool {
	if len(cs) < 1 {
		return false
	}
	c := cs[0]
	body := c.Body()
	lowerShadow := c.LowerShadow()
	upperShadow := c.UpperShadow()

	if body == 0 { // Avoid division by zero
		return false
	}

	return lowerShadow > 2*body && upperShadow < 0.1*body && c.Close > c.Low && c.Close < c.Open // Body at the upper end
}

// InvertedHammer
// 倒锤头线
// 1. Small body at the lower end of the trading range
// 2. Long upper shadow (at least 2 times the body)
// 3. Little or no lower shadow
// Note: Appears in a downtrend and suggests a potential bullish reversal.
func InvertedHammer(cs []CandlestickWrapper) bool {
	if len(cs) < 1 {
		return false
	}
	c := cs[0]
	body := c.Body()
	upperShadow := c.UpperShadow()
	lowerShadow := c.LowerShadow()

	if body == 0 { // Avoid division by zero
		return false
	}

	return upperShadow > 2*body && lowerShadow < 0.1*body && c.Open > c.Low && c.Close > c.Open // Body at the lower end
}

// ShootingStar
// 流星线
// 1. Small body at the lower end of the trading range
// 2. Long upper shadow (at least 2 times the body)
// 3. Little or no lower shadow
// Note: Appears in an uptrend and suggests a potential bearish reversal. The shape is the same as an Inverted Hammer, but its context is different.
func ShootingStar(cs []CandlestickWrapper) bool {
	if len(cs) < 1 {
		return false
	}
	c := cs[0]
	body := c.Body()
	upperShadow := c.UpperShadow()
	lowerShadow := c.LowerShadow()

	if body == 0 { // Avoid division by zero
		return false
	}

	return upperShadow > 2*body && lowerShadow < 0.1*body && c.Open > c.Low && c.Close < c.Open // Body at the lower end
}

// BullishEngulfing EngulfingPattern (Bullish and Bearish)
// 吞噬形态 (看涨和看跌)
// 1. Two candlesticks
// 2. The body of the second candlestick completely engulfs the body of the first candlestick.
// 3. Bullish Engulfing: First is bearish, second is bullish.
// 4. Bearish Engulfing: First is bullish, second is bearish.
// Note: Suggests a potential reversal of the current trend.
func BullishEngulfing(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	return first.IsBearish() && second.IsBullish() &&
		second.Open <= first.Close && second.Close >= first.Open
}

// BearishEngulfing
// 看跌吞噬
// (Implements the Bearish Engulfing part of the Engulfing Pattern)
func BearishEngulfing(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	return first.IsBullish() && second.IsBearish() &&
		second.Open >= first.Close && second.Close <= first.Open
}

// PiercingLine
// 刺透形态
// 1. Two candlesticks
// 2. First candlestick is a long bearish candle.
// 3. Second candlestick is a bullish candle that opens below the low of the first and closes above the midpoint of the first.
// Note: Appears in a downtrend and suggests a potential bullish reversal.
func PiercingLine(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	if !first.IsBearish() || !second.IsBullish() {
		return false
	}

	midPoint := first.Open - first.Body()/2
	return second.Open < first.Low && second.Close > midPoint
}

// DarkCloudCover
// 乌云盖顶
// 1. Two candlesticks
// 2. First candlestick is a long bullish candle.
// 3. Second candlestick is a bearish candle that opens above the high of the first and closes below the midpoint of the first.
// Note: Appears in an uptrend and suggests a potential bearish reversal.
func DarkCloudCover(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	if !first.IsBullish() || !second.IsBearish() {
		return false
	}

	midPoint := first.Open + first.Body()/2
	return second.Open > first.High && second.Close < midPoint
}

// MorningStar
// 启明星
// 1. Three candlesticks
// 2. First is a long bearish candle.
// 3. Second is a small-bodied candle (bullish or bearish) that gaps down from the first.
// 4. Third is a bullish candle that closes well into the body of the first candle.
// Note: Appears in a downtrend and suggests a potential bullish reversal.
func MorningStar(cs []CandlestickWrapper) bool {
	if len(cs) < 3 {
		return false
	}
	first := cs[2]
	second := cs[1]
	third := cs[0]

	if !first.IsBearish() || math.Min(second.Open, second.Close) >= first.Close {
		return false // Second candle gaps down
	}

	if !third.IsBullish() || third.Close <= (first.Open+first.Close)/2 {
		return false // Third candle closes well into the first
	}

	bodySecond := math.Abs(second.Open - second.Close)
	rangeSecond := second.High - second.Low
	return bodySecond <= 0.3*rangeSecond // Second candle has a small body (adjust 0.3 as needed)
}

// EveningStar
// 黄昏之星
// 1. Three candlesticks
// 2. First is a long bullish candle.
// 3. Second is a small-bodied candle (bullish or bearish) that gaps up from the first.
// 4. Third is a bearish candle that closes well into the body of the first candle.
// Note: Appears in an uptrend and suggests a potential bearish reversal.
func EveningStar(cs []CandlestickWrapper) bool {
	if len(cs) < 3 {
		return false
	}
	first := cs[2]
	second := cs[1]
	third := cs[0]

	if !first.IsBullish() || math.Max(second.Open, second.Close) <= first.Close {
		return false // Second candle gaps up
	}

	if !third.IsBearish() || third.Close >= (first.Open+first.Close)/2 {
		return false // Third candle closes well into the first
	}

	bodySecond := math.Abs(second.Open - second.Close)
	rangeSecond := second.High - second.Low
	return bodySecond <= 0.3*rangeSecond // Second candle has a small body (adjust 0.3 as needed)
}

// ThreeWhiteSoldiers
// 红三兵
// 1. Three consecutive long bullish candles.
// 2. Each candle opens within the body of the previous candle.
// 3. Each candle closes above the high of the previous candle.
// Note: Appears in a downtrend and suggests a strong bullish reversal.
func ThreeWhiteSoldiers(cs []CandlestickWrapper) bool {
	if len(cs) < 3 {
		return false
	}
	first := cs[2]
	second := cs[1]
	third := cs[0]

	if !first.IsBullish() || !second.IsBullish() || !third.IsBullish() {
		return false
	}

	if second.Open > first.Close || third.Open > second.Close { // Opens within the previous body
		return false
	}

	if second.Close <= first.High || third.Close <= second.High { // Closes above the previous high
		return false
	}

	return true
}

// ThreeBlackCrows
// 黑三鸦
// 1. Three consecutive long bearish candles.
// 2. Each candle opens within the body of the previous candle.
// 3. Each candle closes below the low of the previous candle.
// Note: Appears in an uptrend and suggests a strong bearish reversal.
func ThreeBlackCrows(cs []CandlestickWrapper) bool {
	if len(cs) < 3 {
		return false
	}
	first := cs[2]
	second := cs[1]
	third := cs[0]

	if !first.IsBearish() || !second.IsBearish() || !third.IsBearish() {
		return false
	}

	if second.Open < first.Close || third.Open < second.Close { // Opens within the previous body
		return false
	}

	if second.Close >= first.Low || third.Close >= second.Low { // Closes below the previous low
		return false
	}

	return true
}

// TweezerBottoms
// 镊子底部
// 1. Two candlesticks with matching lows.
// 2. The first candle is usually bearish, and the second is usually bullish.
// 3. Occurs in a downtrend and suggests a potential bullish reversal.
func TweezerBottoms(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	return math.Abs(first.Low-second.Low) < 0.001*math.Min(first.Low, second.Low) // Lows are very close
}

// TweezerTops
// 镊子顶部
// 1. Two candlesticks with matching highs.
// 2. The first candle is usually bullish, and the second is usually bearish.
// 3. Occurs in an uptrend and suggests a potential bearish reversal.
func TweezerTops(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	return math.Abs(first.High-second.High) < 0.001*math.Min(first.High, second.High) // Highs are very close
}

// FallingWindow
// 下降窗口
// 1. Two bearish candlesticks separated by a gap down (the low of the second is below the high of the first).
// 2. The gap represents an area of selling pressure.
// Note: Suggests continuation of the downtrend or a potential resistance level.
func FallingWindow(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	return first.IsBearish() && second.IsBearish() && second.High < first.Low
}

// RisingWindow
// 上升窗口
// 1. Two bullish candlesticks separated by a gap up (the high of the second is above the low of the first).
// 2. The gap represents an area of buying pressure.
// Note: Suggests continuation of the uptrend or a potential support level.
func RisingWindow(cs []CandlestickWrapper) bool {
	if len(cs) < 2 {
		return false
	}
	first := cs[1]
	second := cs[0]

	return first.IsBullish() && second.IsBullish() && second.Low > first.High
}

// SpinningTop
// 纺锤线
// 1. Small real body (the difference between open and close is small).
// 2. Long upper and lower shadows of roughly equal length.
// Note: Represents indecision in the market.
func SpinningTop(cs []CandlestickWrapper) bool {
	if len(cs) < 1 {
		return false
	}
	c := cs[0]
	body := c.Body()
	upperShadow := c.UpperShadow()
	lowerShadow := c.LowerShadow()
	totalRange := c.High - c.Low

	if totalRange == 0 {
		return false // Avoid division by zero.
	}
	bodyRatio := body / totalRange
	upperShadowRatio := upperShadow / totalRange
	lowerShadowRatio := lowerShadow / totalRange

	// Check for small body and significant shadows
	return bodyRatio < 0.3 && upperShadowRatio > 0.2 && lowerShadowRatio > 0.2 && math.Abs(upperShadow-lowerShadow) < 0.2*totalRange
}
