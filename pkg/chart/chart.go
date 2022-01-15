package chart

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

var (
	barWidth = 80
)

// GenerateChart generates output.png file in root directory. If fails returns error
func GenerateChart(data []float64) error {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent
	stackedBarChart := chart.StackedBarChart{
		Title:      "Bandford Charts",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 75,
			},
		},
		Width:      800,
		Height:     600,
		XAxis:      chart.StyleShow(),
		YAxis:      chart.StyleShow(),
		BarSpacing: 40,
		Bars: []chart.StackedBar{
			{
				Name:  "Actual Data",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "1",
						Value: data[0] * 100,
					},
					{
						Label: "2",
						Value: data[1] * 100,
					},
					{
						Label: "3",
						Value: data[2] * 100,
					},
					{
						Label: "4",
						Value: data[3] * 100,
					},
					{
						Label: "5",
						Value: data[4] * 100,
					},
					{
						Label: "6",
						Value: data[5] * 100,
					},
					{
						Label: "7",
						Value: data[6] * 100,
					},
					{
						Label: "8",
						Value: data[7] * 100,
					},
					{
						Label: "9",
						Value: data[8] * 100,
					},
				},
			},
			{
				Name:  "ExpectedData",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "1",
						Value: 30.1,
					},
					{
						Label: "2",
						Value: 17.6,
					},
					{
						Label: "3",
						Value: 12.5,
					},
					{
						Label: "4",
						Value: 9.7,
					},
					{
						Label: "5",
						Value: 7.9,
					},
					{
						Label: "6",
						Value: 6.7,
					},
					{
						Label: "7",
						Value: 5.8,
					},
					{
						Label: "8",
						Value: 5.1,
					},
					{
						Label: "9",
						Value: 4.6,
					},
				},
			},
		},
	}

	pngFile, err := os.Create("output.png")
	if err != nil {
		return err
	}
	err = stackedBarChart.Render(chart.PNG, pngFile)
	if err != nil {
		return err
	}
	if err := pngFile.Close(); err != nil {
		return err
	}
	return nil
}
