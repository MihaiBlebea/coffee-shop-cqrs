package cmd

import (
	"os"

	"github.com/MihaiBlebea/coffee-shop-query/conn"
	"github.com/MihaiBlebea/coffee-shop-query/evstore"
	orderv "github.com/MihaiBlebea/coffee-shop-query/order_view"
	"github.com/MihaiBlebea/coffee-shop-query/server"
	"github.com/MihaiBlebea/coffee-shop-query/server/handler"
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

		c, err := conn.ConnectMongo()
		if err != nil {
			return err
		}
		orderViewStore := orderv.New(c)

		ev, err := evstore.New(orderViewStore)
		if err != nil {
			return err
		}

		ev.Listen()

		h := handler.New(l)
		s := server.New(h, l)

		s.Run()

		return nil
	},
}
