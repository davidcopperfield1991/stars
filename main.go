package main

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/cobra"
)

type Star struct {
	gorm.Model
	Taskname string
	Stars    int
	Status   bool
}

type Dailystar struct {
	gorm.Model
	Taskname string
	Stars    int
}

func dbde() gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=star port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Star{})
	return *db
}

func dbdaily() gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=star port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Dailystar{})
	return *db
}

func main() {
	rootCmd.AddCommand(starCmd, addCmd, listCmd, doneCmd, deleteCmd, helpCmd, todayCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use: "",
}

func updateStar(taskname string, star int) {
	db := dbde()
	var records []Star
	db.Where("taskname = ?", taskname).First(&records)
	oldstar := records[0].Stars
	give_ster := star + oldstar
	db.Model(&Star{}).Where("taskname", taskname).Update("Stars", give_ster)

}

var starCmd = &cobra.Command{
	Use: "star",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		star, _ := strconv.Atoi(args[1])
		db := dbdaily()
		db.Create((&Dailystar{Taskname: title, Stars: star}))
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
		db := dbde()
		db.Create(&Star{
			Taskname: args[0],
		})
	},
}

var doneCmd = &cobra.Command{
	Use: "done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agha salam")
		db := dbde()
		title := args[0]
		db.Model(&Star{}).Where("taskname", title).Update("status", true)
	},
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		db := dbde()
		title := args[0]
		db.Model(&Star{}).Where("taskname", title).Delete(&Star{})
	},
}

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		db := dbde()
		var records []Star
		db.Find(&records)
		for i := range records {
			fmt.Printf("task: %v -------- status : %v ------- stars: %v \n", records[i].Taskname, records[i].Status, records[i].Stars)
		}
	},
}

var todayCmd = &cobra.Command{
	Use: "today",
	Run: func(cmd *cobra.Command, args []string) {
		db := dbdaily()
		dt := time.Now()
		runes := []rune(dt.String())
		time := string(runes[0:10])
		az := time + " 00:00:00"
		ta := time + " 23:59:59"
		var motaghayer []Dailystar
		db.Where("created_at BETWEEN ? AND ?", az, ta).Find(&motaghayer)
		total := 0
		for i := range motaghayer {
			fmt.Printf("task: %v --------------- stars: %v \n", motaghayer[i].Taskname, motaghayer[i].Stars)
			total += motaghayer[i].Stars
		}
		fmt.Printf("today star : %v \n", total)
	},
}

var helpCmd = &cobra.Command{
	Use: "help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("star : this command with task name and number of star you want to give add one star to task \nadd : this command with task name add task to DB  \ndone : this command with task name change statuse of task to true \nlist: this command show DB rows \ndelete: this command with task name delete task from DB")
	},
}
