package cmd

import (
	"fmt"
	"strconv"

	"project-app-todo-list-cli-nama/service"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [id]",
    Short: "Menghapus tugas berdasarkan ID",
    Run: func(cmd *cobra.Command, args []string) {

        if len(args) == 0 {
            fmt.Println("Masukkan ID tugas yang ingin dihapus!")
            return
        }

        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("ID harus berupa angka!")
            return
        }

        err = service.DeleteTask(id)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        // Info sukses
        fmt.Printf("Tugas dengan ID %d berhasil dihapus!\n\n", id)

        // Tampilkan tabel terbaru
        tasks, _ := service.LoadTasks()

        fmt.Println("=================================  Updated To-Do List  =================================")

        fmt.Println("+----+----------------------+----------------+------------+")
        fmt.Println("| ID | Task                 | Status         | Priority   |")
        fmt.Println("+----+----------------------+----------------+------------+")

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

            // Format tabel FIXED WIDTH
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
    rootCmd.AddCommand(deleteCmd)
}
