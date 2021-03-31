package cmd

import (
	"os"

	"github.com/MihaiBlebea/coffee-shop-command/conn"
	"github.com/MihaiBlebea/coffee-shop-command/trans"
	"github.com/MihaiBlebea/coffee-shop-command/user"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the write db.",
	Long:  "Migrate the write db.",
	RunE: func(cmd *cobra.Command, args []string) error {

		l := logrus.New()

		l.SetFormatter(&logrus.JSONFormatter{})
		l.SetOutput(os.Stdout)
		l.SetLevel(logrus.InfoLevel)

		db, err := conn.ConnectSQL()
		if err != nil {
			return err
		}

		us := user.New(db, nil, nil, nil)

		ts := trans.New(db)

		if err := ts.Migrate(); err != nil {
			return err
		}

		if err := us.Migrate(); err != nil {
			return err
		}

		return nil
	},
}
