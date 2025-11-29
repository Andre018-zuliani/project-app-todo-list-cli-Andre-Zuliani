package cmd

import (
	"fmt"
	"project-app-todo-list-cli-nama/service"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
    Use:   "search [keyword]",
    Short: "Mencari tugas berdasarkan keyword",
    Run: func(cmd *cobra.Command, args []string) {

        if len(args) == 0 {
            fmt.Println("Masukkan keyword pencarian!")
            return
        }

        keyword := args[0]
        tasks, _ := service.SearchTask(keyword)

        fmt.Println("=================================  Search Result  =================================")

        // HEADER TABEL
        fmt.Println("+----+----------------------+----------------+------------+")
        fmt.Println("| ID | Title                | Status         | Priority   |")
        fmt.Println("+----+----------------------+----------------+------------+")

        // ISI TABEL
        for _, t := range tasks {

            // warna status
            status := t.Status
            switch status {
            case "completed":
                status = color.BlueString("completed")
            case "progress":
                status = color.YellowString("progress")
            case "pending":
                status = color.RedString("pending")
            }

            // warna priority
            prio := t.Priority
            switch prio {
            case "high":
                prio = color.RedString("high")
            case "medium":
                prio = color.YellowString("medium")
            case "low":
                prio = color.GreenString("low")
            }

            // format tabel FIXED WIDTH (tidak akan berantakan)
            fmt.Printf(
                "| %-2d | %-20s | %-14s | %-10s |\n",
                t.ID,
                t.Title,
                status,
                prio,
            )
        }

        // FOOTER TABEL
        fmt.Println("+----+----------------------+----------------+------------+")
    },
}

func init() {
    rootCmd.AddCommand(searchCmd)
}
