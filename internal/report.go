package internal

import (
	"fmt"
	"time"

	"github.com/go-echarts/go-echarts/v2/opts"
)

func Report(days int) []Dailystar {
	db := Dbdaily()
	dt := time.Now().AddDate(0, 0, -days)
	dtt := time.Now()
	runes := []rune(dt.String())
	runest := []rune(dtt.String())
	time := string(runes[0:10])
	timet := string(runest[0:10])
	az := time + " 00:00:00"
	ta := timet + " 23:59:59"
	fmt.Println(az)
	fmt.Println(ta)
	var motaghayer []Dailystar
	db.Where("created_at BETWEEN ? AND ?", az, ta).Find(&motaghayer)
	total := 0
	// result := []string{}
	for i := range motaghayer {
		fmt.Printf("task: %v --------------- stars: %v \n", motaghayer[i].Taskname, motaghayer[i].Stars)
		total += motaghayer[i].Stars
	}
	// fmt.Printf("today star : %v \n", total)
	return motaghayer

}

func NewReport(days int) []Star {
	db := Dbde()
	dt := time.Now().AddDate(0, 0, -days)
	dtt := time.Now()
	runes := []rune(dt.String())
	runest := []rune(dtt.String())
	time := string(runes[0:10])
	timet := string(runest[0:10])
	az := time + " 00:00:00"
	ta := timet + " 23:59:59"
	fmt.Println(az)
	fmt.Println(ta)
	var motaghayer []Star
	db.Where("created_at BETWEEN ? AND ?", az, ta).Find(&motaghayer)
	total := 0
	// result := []string{}
	for i := range motaghayer {
		fmt.Printf("task: %v --------------- stars: %v \n", motaghayer[i].Taskname, motaghayer[i].Stars)
		total += motaghayer[i].Stars
	}
	// fmt.Printf("today star : %v \n", total)
	return motaghayer

}

func Render(day int) []string {
	value := NewReport(day)
	result := []string{}
	for i := range value {
		result = append(result, value[i].Taskname)
	}
	fmt.Println(value)
	return result
}

func RenderInt(day int) []opts.BarData {
	items := make([]opts.BarData, 0)
	value := NewReport(day)
	for i := range value {
		items = append(items, opts.BarData{Value: value[i].Stars})
	}
	fmt.Println(value)
	return items
	// value := Report(day)
	// result := []opts.LineData{}
	// for i := range value {
	// 	result = append(result, value[i].Stars)
	// }
	// return result
}
