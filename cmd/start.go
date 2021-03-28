package cmd

import (
	"os"

	"github.com/MihaiBlebea/coffee-shop-cqrs/coffee"
	"github.com/MihaiBlebea/coffee-shop-cqrs/conn"
	"github.com/MihaiBlebea/coffee-shop-cqrs/server"
	"github.com/MihaiBlebea/coffee-shop-cqrs/server/handler"
	"github.com/MihaiBlebea/coffee-shop-cqrs/trans"
	"github.com/MihaiBlebea/coffee-shop-cqrs/user"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the write server api.",
	Long:  "Start the write server api.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		db, err := conn.ConnectSQL()
		if err != nil {
			return err
		}

		cs, err := coffee.New()
		if err != nil {
			return err
		}

		ts := trans.New(db)
		if err != nil {
			return err
		}

		us := user.New(db, cs, ts)

		h := handler.New(us, l)
		s := server.New(h, l)

		s.Run()

		return nil
	},
}
