package cmd

import (
	"fmt"
	"project-app-todo-list-cli-nama/service"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "Menampilkan semua tugas",
    Run: func(cmd *cobra.Command, args []string) {

        tasks, _ := service.LoadTasks()

        fmt.Println("=================================  To-Do List  =================================")

        fmt.Println("+----+----------------------+----------------+------------+")
        fmt.Println("| No | Task                 | Status         | Priority   |")
        fmt.Println("+----+----------------------+----------------+------------+")

        for _, t := range tasks {

            // status coloring
            status := t.Status
            switch status {
            case "completed":
                status = color.BlueString("completed")
            case "progress":
                status = color.YellowString("progress")
            case "pending":
                status = color.RedString("pending")
            }

            // priority coloring
            prio := t.Priority
            switch prio {
            case "high":
                prio = color.RedString("high")
            case "medium":
                prio = color.YellowString("medium")
            case "low":
                prio = color.GreenString("low")
            }

            // format fixed-width (tidak akan geser)
            fmt.Printf(
                "| %-2d | %-20s | %-14s | %-10s |\n",
                t.ID,
                t.Title,
                status,
                prio,
            )
        }

        fmt.Println("+----+----------------------+----------------+------------+")
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
