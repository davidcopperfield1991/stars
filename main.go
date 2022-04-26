package main

import (
	"fmt"
	"strconv"

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

func dbde() gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=admin dbname=star port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Star{})
	return *db
}

func main() {
	rootCmd.AddCommand(starCmd, addCmd, listCmd, doneCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use: "",
}

var starCmd = &cobra.Command{
	Use: "star",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		var records []Star
		db := dbde()
		db.Where("taskname = ?", title).First(&records)
		oldstar := records[0].Stars
		newstar, _ := strconv.Atoi(args[1])
		give_ster := newstar + oldstar
		db.Model(&Star{}).Where("taskname", title).Update("Stars", give_ster)
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
