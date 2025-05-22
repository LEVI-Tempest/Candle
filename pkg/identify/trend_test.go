package identify

import (
	"testing"
	"time"

	v1 "github.com/LEVI-Tempest/Candle/pkg/proto"
)

func TestDetermineLongTermTrend(t *testing.T) {
	// Create test data for different trend scenarios
	// 创建不同趋势场景的测试数据

	// Upward trend data (上升趋势数据)
	upTrendCandles := []*v1.Candlestick{
		{Timestamp: time.Now().AddDate(0, 0, -6).Unix(), Open: 100, High: 105, Low: 98, Close: 102},
		{Timestamp: time.Now().AddDate(0, 0, -5).Unix(), Open: 102, High: 108, Low: 101, Close: 106},
		{Timestamp: time.Now().AddDate(0, 0, -4).Unix(), Open: 106, High: 110, Low: 104, Close: 109},
		{Timestamp: time.Now().AddDate(0, 0, -3).Unix(), Open: 109, High: 115, Low: 108, Close: 113},
		{Timestamp: time.Now().AddDate(0, 0, -2).Unix(), Open: 113, High: 118, Low: 112, Close: 117},
		{Timestamp: time.Now().AddDate(0, 0, -1).Unix(), Open: 117, High: 122, Low: 116, Close: 120},
		{Timestamp: time.Now().Unix(), Open: 120, High: 125, Low: 119, Close: 124},
	}

	// Downward trend data (下降趋势数据)
	downTrendCandles := []*v1.Candlestick{
		{Timestamp: time.Now().AddDate(0, 0, -6).Unix(), Open: 100, High: 102, Low: 98, Close: 99},
		{Timestamp: time.Now().AddDate(0, 0, -5).Unix(), Open: 99, High: 100, Low: 95, Close: 96},
		{Timestamp: time.Now().AddDate(0, 0, -4).Unix(), Open: 96, High: 97, Low: 92, Close: 93},
		{Timestamp: time.Now().AddDate(0, 0, -3).Unix(), Open: 93, High: 94, Low: 89, Close: 90},
		{Timestamp: time.Now().AddDate(0, 0, -2).Unix(), Open: 90, High: 91, Low: 86, Close: 87},
		{Timestamp: time.Now().AddDate(0, 0, -1).Unix(), Open: 87, High: 88, Low: 83, Close: 84},
		{Timestamp: time.Now().Unix(), Open: 84, High: 85, Low: 80, Close: 81},
	}

	// Sideways/oscillating trend data (震荡趋势数据)
	sidewaysTrendCandles := []*v1.Candlestick{
		{Timestamp: time.Now().AddDate(0, 0, -6).Unix(), Open: 100, High: 102, Low: 98, Close: 101},
		{Timestamp: time.Now().AddDate(0, 0, -5).Unix(), Open: 101, High: 103, Low: 99, Close: 100},
		{Timestamp: time.Now().AddDate(0, 0, -4).Unix(), Open: 100, High: 102, Low: 98, Close: 102},
		{Timestamp: time.Now().AddDate(0, 0, -3).Unix(), Open: 102, High: 104, Low: 100, Close: 101},
		{Timestamp: time.Now().AddDate(0, 0, -2).Unix(), Open: 101, High: 103, Low: 99, Close: 102},
		{Timestamp: time.Now().AddDate(0, 0, -1).Unix(), Open: 102, High: 104, Low: 100, Close: 100},
		{Timestamp: time.Now().Unix(), Open: 100, High: 102, Low: 98, Close: 101},
	}

	// Test upward trend detection
	// 测试上升趋势检测
	t.Run("UpTrend", func(t *testing.T) {
		trend, err := DetermineLongTermTrend(upTrendCandles, 7)
		if err != nil {
			t.Fatalf("Failed to determine upward trend: %v (判断上升趋势失败)", err)
		}
		if trend != TrendYang {
			t.Errorf("Expected upward trend (TrendYang), got %v (预期上升趋势)", trend)
		}
	})

	// Test downward trend detection
	// 测试下降趋势检测
	t.Run("DownTrend", func(t *testing.T) {
		trend, err := DetermineLongTermTrend(downTrendCandles, 7)
		if err != nil {
			t.Fatalf("Failed to determine downward trend: %v (判断下降趋势失败)", err)
		}
		if trend != TrendYin {
			t.Errorf("Expected downward trend (TrendYin), got %v (预期下降趋势)", trend)
		}
	})

	// Test sideways/oscillating trend detection
	// 测试震荡趋势检测
	t.Run("SidewaysTrend", func(t *testing.T) {
		trend, err := DetermineLongTermTrend(sidewaysTrendCandles, 7)
		if err != nil {
			t.Fatalf("Failed to determine sideways trend: %v (判断震荡趋势失败)", err)
		}
		if trend != TrendMiddle {
			t.Errorf("Expected sideways trend (TrendMiddle), got %v (预期震荡趋势)", trend)
		}
	})

	// Test insufficient data case
	// 测试数据不足的情况
	t.Run("InsufficientData", func(t *testing.T) {
		_, err := DetermineLongTermTrend(upTrendCandles[:3], 7)
		if err == nil {
			t.Error("Expected error when data is insufficient, but got none (预期数据不足时返回错误)")
		}
	})

	// Test invalid days parameter
	// 测试无效天数参数
	t.Run("InvalidDays", func(t *testing.T) {
		_, err := DetermineLongTermTrend(upTrendCandles, 0)
		if err == nil {
			t.Error("Expected error when days parameter is invalid, but got none (预期天数无效时返回错误)")
		}
	})
}
