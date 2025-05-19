package identify

import (
	"fmt"
	"testing"

	v1 "github.com/LEVI-Tempest/Candle/pkg/proto"
)

func TestIdentify(t *testing.T) {
	t.Log("test identify start ...")
	candles := []v1.Candlestick{
		{Open: 10, High: 12, Low: 9, Close: 11, Volume: 100},
		{Open: 11, High: 11.5, Low: 10.5, Close: 11.2, Volume: 120},
		{Open: 11.2, High: 13, Low: 11, Close: 12.5, Volume: 150},
		{Open: 12.5, High: 12.8, Low: 12.3, Close: 12.6, Volume: 90},
		{Open: 12.6, High: 13.2, Low: 12.4, Close: 12.9, Volume: 110},
		{Open: 12.9, High: 14, Low: 12.7, Close: 13.5, Volume: 130},
		{Open: 13.5, High: 13.8, Low: 13.1, Close: 13.3, Volume: 100},
		{Open: 13.3, High: 13.6, Low: 12.9, Close: 13.2, Volume: 120},
		{Open: 13.2, High: 13.5, Low: 12.8, Close: 13.4, Volume: 140},
		{Open: 13.4, High: 14.2, Low: 13, Close: 14, Volume: 160},
		{Open: 20, High: 22, Low: 18, Close: 21, Volume: 200}, // Add example data for Marubozu
		{Open: 20, High: 20.1, Low: 10, Close: 20.1, Volume: 200},
		{Open: 20, High: 20.1, Low: 10, Close: 15, Volume: 200},
	}

	// 将 Candlestick 转换为 CandlestickWrapper
	var wrappedCandles []CandlestickWrapper
	for _, c := range candles {
		wrappedCandles = append(wrappedCandles, NewCandlestickWrapper(&c))
	}

	// 示例用法
	fmt.Println("Doji:", Doji(wrappedCandles[:1]))
	fmt.Println("Marubozu:", Marubozu(wrappedCandles[10:11])) // Use the new example data.
	fmt.Println("Hammer:", Hammer(wrappedCandles[:1]))
	fmt.Println("Hanging Man:", HangingMan(wrappedCandles[:1]))
	fmt.Println("Inverted Hammer:", InvertedHammer(wrappedCandles[:1]))
	fmt.Println("Shooting Star:", ShootingStar(wrappedCandles[:1]))
	fmt.Println("Bullish Engulfing:", BullishEngulfing(wrappedCandles[:2]))
	fmt.Println("Bearish Engulfing:", BearishEngulfing(wrappedCandles[:2]))
	fmt.Println("Piercing Line:", PiercingLine(wrappedCandles[:2]))
	fmt.Println("Dark Cloud Cover:", DarkCloudCover(wrappedCandles[:2]))
	fmt.Println("Morning Star:", MorningStar(wrappedCandles[:3]))
	fmt.Println("Evening Star:", EveningStar(wrappedCandles[:3]))
	fmt.Println("Three White Soldiers:", ThreeWhiteSoldiers(wrappedCandles[:3]))
	fmt.Println("Three Black Crows:", ThreeBlackCrows(wrappedCandles[:3]))
	fmt.Println("Tweezer Bottoms:", TweezerBottoms(wrappedCandles[:2]))
	fmt.Println("Tweezer Tops:", TweezerTops(wrappedCandles[:2]))
	fmt.Println("Falling Window:", FallingWindow(wrappedCandles[:2]))
	fmt.Println("Rising Window:", RisingWindow(wrappedCandles[:2]))
	fmt.Println("Spinning Top:", SpinningTop(wrappedCandles[:1]))
	fmt.Println("Umbrella:", Umbrella(wrappedCandles[:1]))
}
