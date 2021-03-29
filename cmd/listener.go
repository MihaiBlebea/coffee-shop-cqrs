package cmd

import (
	"fmt"
	"os"

	"github.com/MihaiBlebea/coffee-shop-cqrs/broker"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listenerCmd)
}

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Test task.",
	Long:  "Test task.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		b, err := broker.New()
		if err != nil {
			return err
		}

		err = b.Listen()
		if err != nil {
			return err
		}

		fmt.Println(b)

		return nil
	},
}
