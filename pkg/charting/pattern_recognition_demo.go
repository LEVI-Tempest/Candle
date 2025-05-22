package charting

import (
	"fmt"
	"log"
	"time"

	v1 "github.com/LEVI-Tempest/Candle/pkg/proto"
)

func Identify() {
	fmt.Println("Candlestick Pattern Recognition Demo")
	fmt.Println("===================================")

	// Create sample candlestick data with various patterns
	// 创建包含各种形态的示例蜡烛图数据
	candleData := createSampleData()

	// Create enhanced kline chart
	// 创建增强K线图
	enhancedKline := NewEnhancedKline()

	// Load the data
	// 加载数据
	enhancedKline.LoadData(candleData)
	fmt.Printf("Loaded %d candlesticks\n", len(enhancedKline.Data))

	// Auto-detect patterns
	// 自动检测形态
	enhancedKline.AutoDetectPatterns()
	fmt.Printf("Detected %d patterns\n\n", len(enhancedKline.Patterns))

	// Print detected patterns with details
	// 打印检测到的形态及详细信息
	fmt.Println("Detected Patterns:")
	fmt.Println("------------------")
	for i, pattern := range enhancedKline.Patterns {
		fmt.Printf("%d. %s\n", i+1, pattern.Type)
		fmt.Printf("   Position: %d\n", pattern.Position)
		fmt.Printf("   Price: %.2f\n", pattern.Price)
		fmt.Printf("   Strength: %.2f\n", pattern.Strength)
		fmt.Printf("   Risk: %.2f\n", pattern.Risk)
		fmt.Printf("   Time: %s\n\n", pattern.Time)
	}

	// Get pattern summary
	// 获取形态摘要
	summary := enhancedKline.GetPatternSummary()
	fmt.Println("Pattern Summary:")
	fmt.Println("----------------")
	for patternType, count := range summary {
		fmt.Printf("%-20s: %d occurrences\n", patternType, count)
	}

	// Create chart with patterns
	// 创建带有形态的图表
	enhancedKline.CreateChart("Candlestick Chart with Automatic Pattern Recognition")

	// Render to HTML file
	// 渲染到HTML文件
	filename := "pattern_recognition_demo.html"
	err := enhancedKline.RenderToFile(filename)
	if err != nil {
		log.Fatalf("Failed to render chart: %v", err)
	}

	fmt.Printf("\nChart successfully rendered to %s\n", filename)
	fmt.Println("Open the file in your web browser to view the interactive chart with marked patterns!")
}

// createSampleData creates sample candlestick data with various patterns
// 创建包含各种形态的示例蜡烛图数据
func createSampleData() []*v1.Candlestick {
	baseTime := time.Now().AddDate(0, 0, -50) // 50 days ago

	return []*v1.Candlestick{
		// Uptrend with normal candles
		{Timestamp: baseTime.Unix(), Open: 100, High: 105, Low: 98, Close: 103, Volume: 1000},
		{Timestamp: baseTime.AddDate(0, 0, 1).Unix(), Open: 103, High: 108, Low: 101, Close: 106, Volume: 1200},
		{Timestamp: baseTime.AddDate(0, 0, 2).Unix(), Open: 106, High: 110, Low: 104, Close: 109, Volume: 1100},

		// Doji - market indecision (十字星 - 市场犹豫)
		{Timestamp: baseTime.AddDate(0, 0, 3).Unix(), Open: 109, High: 113, Low: 105, Close: 109.2, Volume: 800},

		// Hammer - potential bullish reversal (锤头线 - 潜在看涨反转)
		{Timestamp: baseTime.AddDate(0, 0, 4).Unix(), Open: 108, High: 110, Low: 95, Close: 107, Volume: 1500},

		// Bullish Marubozu - strong buying pressure (看涨光头光脚 - 强烈买压)
		{Timestamp: baseTime.AddDate(0, 0, 5).Unix(), Open: 107, High: 120, Low: 107, Close: 120, Volume: 2000},

		// Shooting Star - potential bearish reversal (流星线 - 潜在看跌反转)
		{Timestamp: baseTime.AddDate(0, 0, 6).Unix(), Open: 120, High: 135, Low: 119, Close: 122, Volume: 1800},

		// Bearish Engulfing Pattern (看跌吞噬形态)
		{Timestamp: baseTime.AddDate(0, 0, 7).Unix(), Open: 122, High: 126, Low: 121, Close: 125, Volume: 1300}, // First candle (bullish)
		{Timestamp: baseTime.AddDate(0, 0, 8).Unix(), Open: 127, High: 128, Low: 115, Close: 118, Volume: 2200}, // Second candle (bearish, engulfs)

		// Downtrend continues
		{Timestamp: baseTime.AddDate(0, 0, 9).Unix(), Open: 118, High: 120, Low: 112, Close: 114, Volume: 1600},

		// Bullish Engulfing Pattern (看涨吞噬形态)
		{Timestamp: baseTime.AddDate(0, 0, 10).Unix(), Open: 114, High: 116, Low: 110, Close: 111, Volume: 1400}, // First candle (bearish)
		{Timestamp: baseTime.AddDate(0, 0, 11).Unix(), Open: 109, High: 118, Low: 108, Close: 117, Volume: 2100}, // Second candle (bullish, engulfs)

		// Morning Star Pattern (启明星形态)
		{Timestamp: baseTime.AddDate(0, 0, 12).Unix(), Open: 117, High: 118, Low: 110, Close: 112, Volume: 1700}, // First candle (bearish)
		{Timestamp: baseTime.AddDate(0, 0, 13).Unix(), Open: 108, High: 109, Low: 106, Close: 107, Volume: 900},  // Second candle (small body, gap down)
		{Timestamp: baseTime.AddDate(0, 0, 14).Unix(), Open: 109, High: 120, Low: 108, Close: 118, Volume: 2300}, // Third candle (bullish)

		// Three White Soldiers Pattern (红三兵形态)
		{Timestamp: baseTime.AddDate(0, 0, 15).Unix(), Open: 118, High: 125, Low: 117, Close: 124, Volume: 1800}, // First soldier
		{Timestamp: baseTime.AddDate(0, 0, 16).Unix(), Open: 122, High: 130, Low: 121, Close: 129, Volume: 1900}, // Second soldier
		{Timestamp: baseTime.AddDate(0, 0, 17).Unix(), Open: 127, High: 135, Low: 126, Close: 134, Volume: 2000}, // Third soldier

		// Spinning Top - market indecision (陀螺线 - 市场犹豫)
		{Timestamp: baseTime.AddDate(0, 0, 18).Unix(), Open: 134, High: 140, Low: 128, Close: 135, Volume: 1400},

		// Evening Star Pattern (黄昏之星形态)
		{Timestamp: baseTime.AddDate(0, 0, 19).Unix(), Open: 135, High: 142, Low: 134, Close: 141, Volume: 1600},  // First candle (bullish)
		{Timestamp: baseTime.AddDate(0, 0, 20).Unix(), Open: 143, High: 144, Low: 142, Close: 143.5, Volume: 800}, // Second candle (small body, gap up)
		{Timestamp: baseTime.AddDate(0, 0, 21).Unix(), Open: 142, High: 143, Low: 132, Close: 135, Volume: 2100},  // Third candle (bearish)

		// Three Black Crows Pattern (黑三鸦形态)
		{Timestamp: baseTime.AddDate(0, 0, 22).Unix(), Open: 135, High: 136, Low: 128, Close: 130, Volume: 1700}, // First crow
		{Timestamp: baseTime.AddDate(0, 0, 23).Unix(), Open: 132, High: 133, Low: 125, Close: 127, Volume: 1800}, // Second crow
		{Timestamp: baseTime.AddDate(0, 0, 24).Unix(), Open: 129, High: 130, Low: 120, Close: 122, Volume: 1900}, // Third crow

		// Tweezer Bottoms - potential support (镊子底部 - 潜在支撑)
		{Timestamp: baseTime.AddDate(0, 0, 25).Unix(), Open: 122, High: 125, Low: 118, Close: 120, Volume: 1500},
		{Timestamp: baseTime.AddDate(0, 0, 26).Unix(), Open: 121, High: 124, Low: 118, Close: 123, Volume: 1600}, // Same low as previous

		// Recovery with normal candles
		{Timestamp: baseTime.AddDate(0, 0, 27).Unix(), Open: 123, High: 128, Low: 122, Close: 126, Volume: 1300},
		{Timestamp: baseTime.AddDate(0, 0, 28).Unix(), Open: 126, High: 130, Low: 124, Close: 129, Volume: 1400},
		{Timestamp: baseTime.AddDate(0, 0, 29).Unix(), Open: 129, High: 133, Low: 127, Close: 131, Volume: 1200},
		{Timestamp: baseTime.AddDate(0, 0, 30).Unix(), Open: 131, High: 135, Low: 129, Close: 133, Volume: 1100},
	}
}
