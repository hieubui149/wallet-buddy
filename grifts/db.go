package grifts

import (
	"wallet-buddy/models"

	"github.com/gobuffalo/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add wallets
		wallets := []string{"Crypto", "Bank"}
		for _, name := range wallets {
			w := models.Wallet{Name: name}
			err := models.DB.Create(&w)
			if err != nil {
				panic(err)
			}
		}

		return nil
	})

})
