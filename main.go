package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// CandlestickData 表示蜡烛图的数据
type CandlestickData struct {
	Date  string
	Open  float64
	Close float64
	Low   float64
	High  float64
}

func main() {
	// 创建蜡烛图实例
	kline := charts.NewKLine()

	// 设置全局配置
	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Candlestick Chart with Markers",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Date",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Price",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Trigger:   "axis",
			TriggerOn: "mousemove",
		}),
	)

	// 生成模拟数据
	data := generateCandlestickData()

	// 提取数据
	dates := make([]string, 0)
	values := make([]opts.KlineData, 0)
	for _, d := range data {
		dates = append(dates, d.Date)
		values = append(values, opts.KlineData{Value: [4]float64{d.Open, d.Close, d.Low, d.High}})
	}

	// 设置 X 轴和 Y 轴数据
	kline.SetXAxis(dates).AddSeries("kline", values)

	// 标记特定图形
	markPoints := []opts.MarkPoint{
		{
			Name:       "Hammer",
			Coord:      []interface{}{dates[2], values[2].Value[1]}, // 标记锤头线
			Value:      "Hammer",
			Symbol:     "pin",
			SymbolSize: 20,
		},
		{
			Name:       "Doji",
			Coord:      []interface{}{dates[5], values[5].Value[1]}, // 标记十字星
			Value:      "Doji",
			Symbol:     "circle",
			SymbolSize: 20,
		},
	}

	kline.SetSeriesOptions(
		charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
			Label: &opts.Label{
				Show: true,
			},
		}),
		charts.WithMarkPointNameCoordItemOpts(markPoints...),
	)

	// 生成 HTML 文件
	f, _ := os.Create("candlestick.html")
	kline.Render(f)
}

// generateCandlestickData 生成模拟的蜡烛图数据
func generateCandlestickData() []CandlestickData {
	data := make([]CandlestickData, 0)
	for i := 0; i < 10; i++ {
		open := 100 + rand.Float64()*10
		close := open + rand.Float64()*5 - 2.5
		low := math.Min(open, close) - rand.Float64()*2
		high := math.Max(open, close) + rand.Float64()*2
		data = append(data, CandlestickData{
			Date:  "2023-10-0" + strconv.Itoa(i+1),
			Open:  open,
			Close: close,
			Low:   low,
			High:  high,
		})
	}
	return data
}
