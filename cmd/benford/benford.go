package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/CMiksche/benford"
	"github.com/Selahattin/benford/pkg/chart"
	readcsv "github.com/Selahattin/benford/pkg/readCsv"

	"github.com/Selahattin/benford/pkg/version"
)

var (
	versionFlag = flag.Bool("version", false, "Show version information.")
	csvFile     = flag.String("csv", "./test.csv", "for use csv file")
	colIndex    = flag.Int("colIndex", 5, "index at spesific col")
)

func main() {

	flag.Parse()

	// Show version information
	if *versionFlag {
		fmt.Fprintln(os.Stdout, version.Print("benford"))
		os.Exit(0)
	}
	datas, err := readcsv.ReadCsv(*csvFile, *colIndex)
	if err != nil {
		log.Fatal(err)
	}
	benfordResult := benford.CalcBenfords(datas)
	chiSquared := benfordResult.ChiSquared
	fmt.Println(chiSquared)
	dist := benfordResult.Dist
	err = chart.GenerateChart(dist)
	if err != nil {
		log.Fatal(err)

	}
	/*fmt.Println(int((dist[0]) * 100))
	fmt.Println(int((dist[1]) * 100))
	fmt.Println(int((dist[2]) * 100))
	fmt.Println(int((dist[3]) * 100))
	fmt.Println(int((dist[4]) * 100))
	fmt.Println(int((dist[5]) * 100))
	fmt.Println(int((dist[6]) * 100))
	fmt.Println(int((dist[7]) * 100))
	fmt.Println(int((dist[8]) * 100))*/
}
