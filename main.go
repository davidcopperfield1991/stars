package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/davidcopperfield1991/stars/internal"
	server "github.com/davidcopperfield1991/stars/server"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd.AddCommand(starCmd, addCmd, listCmd, doneCmd, deleteCmd, helpCmd, todayCmd, tomatoCmd, reportCmd, serverCmd, weekCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		server.Server()
	},
}

var rootCmd = &cobra.Command{
	Use: "",
}

func updateStar(taskname string, star int) {
	db := internal.Dbde()
	var records []internal.Star
	db.Where("taskname = ?", taskname).First(&records)
	oldstar := records[0].Stars
	give_ster := star + oldstar
	db.Model(&internal.Star{}).Where("taskname", taskname).Update("Stars", give_ster)

}

var starCmd = &cobra.Command{
	Use: "star",
	Run: func(cmd *cobra.Command, args []string) {
		dt := time.Now()
		fmt.Println(dt)
		title := args[0]
		star, _ := strconv.Atoi(args[1])
		db := internal.Dbdaily()
		db.Create((&internal.Dailystar{Taskname: title, Stars: star}))
		updateStar(title, star)
	},
}

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agha salam")
		for i, s := range args {
			fmt.Println(i, s)
		}
		db := internal.Dbde()
		db.Create(&internal.Star{
			Taskname: args[0],
		})
	},
}

var doneCmd = &cobra.Command{
	Use: "done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agha salam")
		db := internal.Dbde()
		title := args[0]
		db.Model(&internal.Star{}).Where("taskname", title).Update("status", true)
	},
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.Dbde()
		title := args[0]
		db.Model(&internal.Star{}).Where("taskname", title).Delete(&internal.Star{})
	},
}

func khateFasele(s int) string {
	b := ""
	for i := 1; i < s; i++ {
		b += "-"
	}
	return b
}

func beauty(n string, s bool, se int) int {
	ln := len(n)
	lz := 30 - ln
	mohtawa, _ := fmt.Printf("task: %v %v status : %v ------- stars: %v  \n", n, khateFasele(lz), s, se)
	return mohtawa
}

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.Dbde()
		var records []internal.Star
		db.Find(&records)
		for i := range records {
			beauty(records[i].Taskname, records[i].Status, records[i].Stars)
		}
	},
}

func calculateValues(rows []internal.Dailystar) map[string]int {
	valueTotals := make(map[string]int)

	for _, row := range rows {
		valueTotals[row.Taskname] += row.Stars
	}

	return valueTotals
}

var todayCmd = &cobra.Command{
	Use: "today",
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.Dbdaily()
		dt := time.Now()
		runes := []rune(dt.String())
		time := string(runes[0:10])
		az := time + " 00:00:00"
		ta := time + " 23:59:59"
		var records []internal.Dailystar

		db.Where("created_at BETWEEN ? AND ?", az, ta).Find(&records)

		counts := calculateValues(records)

		fmt.Println("task star count:")
		for key, count := range counts {
			fmt.Printf("%s: %d\n", key, count)
		}
		total := 0
		for i := range records {
			total += records[i].Stars
		}
		fmt.Println("********today star***********")
		fmt.Printf("today star : %v \n", total)
	},
}

func calculateValuesweek(rows []internal.Dailystar) map[int]int {
	valueTotals := make(map[int]int)

	for _, row := range rows {
		_, _, day := row.CreatedAt.Date()
		valueTotals[day] += row.Stars
	}

	return valueTotals
}

var weekCmd = &cobra.Command{
	Use: "week",
	Run: func(cmd *cobra.Command, args []string) {
		db := internal.Dbdaily()
		dt := time.Now()
		week := dt.Add(-168 * time.Hour)
		runes := []rune(week.String())
		mytime := string(runes[0:10])

		runes2 := []rune(dt.String())
		time2 := string(runes2[0:10])

		az := mytime + " 00:00:00"
		ta := time2 + " 23:59:59"
		var records []internal.Dailystar

		db.Where("created_at BETWEEN ? AND ?", az, ta).Find(&records)

		counts := calculateValuesweek(records)

		fmt.Println("task star count:")

		result := make([]struct{ Key, Value int }, 0, len(counts))
		for k, v := range counts {
			result = append(result, struct{ Key, Value int }{k, v})
		}

		// Sort the slice by keys
		sort.Slice(result, func(i, j int) bool {
			return result[i].Key < result[j].Key
		})

		// Print the sorted key-value pairs
		for _, pair := range result {
			fmt.Printf("Key: %d, Value: %d\n", pair.Key, pair.Value)
		}

		total := 0
		for i := range records {
			total += records[i].Stars
		}
		fmt.Println("********this week star***********")
		fmt.Printf("today star : %v \n", total)
	},
}

var tomatoCmd = &cobra.Command{
	Use: "tomato",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 1; i < 25; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Minute)
			dt := time.Now()
			fmt.Println(dt)
		}
		fmt.Println("tomato sauce is ready.")
	},
}

var reportCmd = &cobra.Command{
	Use: "report",
	Run: func(cmd *cobra.Command, args []string) {
		days, _ := strconv.Atoi(args[0])
		internal.Report(days)

	},
}

var helpCmd = &cobra.Command{
	Use: "help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("star : this command with task name and number of star you want to give add one star to task \nadd : this command with task name add task to DB  \ndone : this command with task name change statuse of task to true \nlist: this command show DB rows \ndelete: this command with task name delete task from DB")
	},
}
