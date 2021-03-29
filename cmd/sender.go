package cmd

import (
	"fmt"
	"os"

	"github.com/MihaiBlebea/coffee-shop-cqrs/broker"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(senderCmd)
}

var senderCmd = &cobra.Command{
	Use:   "sender",
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

		err = b.CreateChannel()
		if err != nil {
			return err
		}

		foo := struct {
			Message string `json:"message"`
		}{
			Message: "Hello lume",
		}

		err = b.PublishMessage(foo)
		if err != nil {
			return err
		}

		fmt.Println(b)

		return nil
	},
}
