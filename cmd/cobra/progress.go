package cli

import (
	"math/rand"
	"time"

	"github.com/achristie/bubbletea-examples/pkg/progress"
	"github.com/spf13/cobra"
)

var progressCmd = &cobra.Command{
	Use:     "progress",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Run: func(cmd *cobra.Command, args []string) {
		p := progress.NewProgress()
		go func() {
			for i := 0; i < 50; i++ {
				r := rand.Intn(2)
				time.Sleep(time.Duration(r) * time.Second)
				p.Send(i)
			}
		}()
		p.Start()

	},
}

func init() {
	rootCmd.AddCommand(progressCmd)
}
