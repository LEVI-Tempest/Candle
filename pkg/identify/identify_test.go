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

func TestLongTermTrend(t *testing.T) {
	// 创建一个上升趋势的数据集
	upTrendCandles := []v1.Candlestick{
		{Open: 10.0, High: 11.0, Low: 9.5, Close: 10.5},
		{Open: 10.5, High: 11.5, Low: 10.2, Close: 11.0},
		{Open: 11.0, High: 12.0, Low: 10.8, Close: 11.8},
		{Open: 11.8, High: 12.5, Low: 11.5, Close: 12.2},
		{Open: 12.2, High: 13.0, Low: 12.0, Close: 12.8},
		{Open: 12.8, High: 13.5, Low: 12.5, Close: 13.2},
		{Open: 13.2, High: 14.0, Low: 13.0, Close: 13.8},
	}

	// 创建一个下降趋势的数据集
	downTrendCandles := []v1.Candlestick{
		{Open: 20.0, High: 20.5, Low: 19.5, Close: 19.8},
		{Open: 19.8, High: 20.0, Low: 19.0, Close: 19.2},
		{Open: 19.2, High: 19.5, Low: 18.5, Close: 18.8},
		{Open: 18.8, High: 19.0, Low: 18.0, Close: 18.2},
		{Open: 18.2, High: 18.5, Low: 17.5, Close: 17.8},
		{Open: 17.8, High: 18.0, Low: 17.0, Close: 17.2},
		{Open: 17.2, High: 17.5, Low: 16.5, Close: 16.8},
	}

	// 创建一个震荡/盘整趋势的数据集
	sidewaysTrendCandles := []v1.Candlestick{
		{Open: 15.0, High: 15.5, Low: 14.5, Close: 15.2},
		{Open: 15.2, High: 15.8, Low: 14.8, Close: 15.0},
		{Open: 15.0, High: 15.6, Low: 14.6, Close: 15.3},
		{Open: 15.3, High: 15.9, Low: 14.9, Close: 15.1},
		{Open: 15.1, High: 15.7, Low: 14.7, Close: 15.4},
		{Open: 15.4, High: 16.0, Low: 15.0, Close: 15.2},
		{Open: 15.2, High: 15.8, Low: 14.8, Close: 15.3},
	}

	// 将 Candlestick 转换为 CandlestickWrapper
	var upTrendWrapped []CandlestickWrapper
	for _, c := range upTrendCandles {
		upTrendWrapped = append(upTrendWrapped, NewCandlestickWrapper(&c))
	}

	var downTrendWrapped []CandlestickWrapper
	for _, c := range downTrendCandles {
		downTrendWrapped = append(downTrendWrapped, NewCandlestickWrapper(&c))
	}

	var sidewaysTrendWrapped []CandlestickWrapper
	for _, c := range sidewaysTrendCandles {
		sidewaysTrendWrapped = append(sidewaysTrendWrapped, NewCandlestickWrapper(&c))
	}

	// 测试简单趋势判断函数
	t.Run("SimpleTrendDetection", func(t *testing.T) {
		upTrend := DetermineTrend(upTrendWrapped, 5)
		downTrend := DetermineTrend(downTrendWrapped, 5)

		t.Logf("Simple Up Trend: %s", upTrend)
		t.Logf("Simple Down Trend: %s", downTrend)

		if upTrend != TrendYang {
			t.Errorf("Expected up trend to be Yang, got %s", upTrend)
		}

		if downTrend != TrendYin {
			t.Errorf("Expected down trend to be Yin, got %s", downTrend)
		}
	})

	// 测试高级趋势分析函数
	t.Run("AdvancedTrendAnalysis", func(t *testing.T) {
		// 设置波动性阈值为3%
		volatilityThreshold := 3.0

		upTrend, upMetrics := AnalyzeLongTermTrend(upTrendWrapped, 7, volatilityThreshold)
		downTrend, downMetrics := AnalyzeLongTermTrend(downTrendWrapped, 7, volatilityThreshold)
		sidewaysTrend, sidewaysMetrics := AnalyzeLongTermTrend(sidewaysTrendWrapped, 7, volatilityThreshold)

		t.Logf("Advanced Up Trend: %s, Metrics: %+v", upTrend, upMetrics)
		t.Logf("Advanced Down Trend: %s, Metrics: %+v", downTrend, downMetrics)
		t.Logf("Advanced Sideways Trend: %s, Metrics: %+v", sidewaysTrend, sidewaysMetrics)

		if upTrend != TrendYang {
			t.Errorf("Expected up trend to be Yang, got %s", upTrend)
		}

		if downTrend != TrendYin {
			t.Errorf("Expected down trend to be Yin, got %s", downTrend)
		}

		if sidewaysTrend != TrendMiddle {
			t.Errorf("Expected sideways trend to be Middle, got %s", sidewaysTrend)
		}

		// 检查指标值
		if upMetrics["priceChangePercent"] <= 0 {
			t.Errorf("Expected positive price change for up trend, got %f", upMetrics["priceChangePercent"])
		}

		if downMetrics["priceChangePercent"] >= 0 {
			t.Errorf("Expected negative price change for down trend, got %f", downMetrics["priceChangePercent"])
		}

		if sidewaysMetrics["volatility"] > volatilityThreshold*2 {
			t.Errorf("Expected low volatility for sideways trend, got %f", sidewaysMetrics["volatility"])
		}
	})
}
