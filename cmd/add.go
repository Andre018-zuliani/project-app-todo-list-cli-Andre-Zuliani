package cmd

import (
	"fmt"
	"project-app-todo-list-cli-nama/service"

	"github.com/spf13/cobra"
)

var (
    title       string
    description string
    priority    string
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Menambahkan tugas baru",
    Run: func(cmd *cobra.Command, args []string) {
        if title == "" {
            fmt.Println("Judul tugas tidak boleh kosong! Gunakan --title \"nama tugas\"")
            return
        }

        // default priority
        if priority == "" {
            priority = "low"
        }

        err := service.AddTask(title, description, priority)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Println("Tugas berhasil ditambahkan!")
        }
    },
}

func init() {
    addCmd.Flags().StringVarP(&title, "title", "t", "", "Judul tugas")
    addCmd.Flags().StringVarP(&description, "desc", "d", "", "Deskripsi tugas (opsional)")
    addCmd.Flags().StringVarP(&priority, "priority", "p", "low", "Prioritas tugas: low | medium | high")

    rootCmd.AddCommand(addCmd)
}
