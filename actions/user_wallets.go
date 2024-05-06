package actions

import (
	"net/http"
	"your_finance/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

// Returns all wallets associated with the current user
// as well as the available wallets for user to interact with.
func MyWallets(c buffalo.Context) error {
	wallets := []models.Wallet{}
	err := models.DB.Order("name desc").All(&wallets)
	if err != nil {
		c.Flash().Add("danger", "Something went wrong fetching wallets")
	}

	user_id := c.Session().Get("current_user_id")
	user := models.User{}
	err = models.DB.Find(&user, user_id)
	if err != nil {
		c.Flash().Add("danger", "User not found")
	}
	models.DB.Load(&user, "Wallets")

	c.Set("all_wallets", wallets)
	c.Set("user", user)
	c.Set("isWalletInUse", func(walletID uuid.UUID) bool {
		for _, wallet := range user.Wallets {
			if wallet.ID == walletID {
				return true
			}
		}
		return false
	})

	return c.Render(200, r.HTML("user_wallets/list.html"))
}

// Add a new wallet to user integrated wallets
func UserWalletsCreate(c buffalo.Context) error {
	default_redirect_url := "/users/me/wallets"
	tx := c.Value("tx").(*pop.Connection)

	user_id := c.Session().Get("current_user_id")

	user := models.User{}
	if err := models.DB.Find(&user, user_id); err != nil {
		c.Flash().Add("danger", "User not found")
		return c.Redirect(http.StatusFound, default_redirect_url)
	}
	models.DB.Load(&user, "Wallets")

	wallet_id := c.Param("wallet_id")
	wallet := models.Wallet{}
	if err := models.DB.Find(&wallet, wallet_id); err != nil {
		c.Flash().Add("warning", "Wallet not found")
		return c.Redirect(http.StatusFound, default_redirect_url)
	}

	// Check if wallet is already in user wallets
	for _, w := range user.Wallets {
		if w.ID == wallet.ID {
			c.Flash().Add("warning", "Wallet is already in use")
			return c.Redirect(http.StatusFound, default_redirect_url)
		}
	}

	// Add wallet to user wallets
	user_wallet := models.UsersWallet{
		UserID:   user.ID,
		WalletID: wallet.ID,
	}
	if err := tx.Create(&user_wallet); err != nil {
		c.Flash().Add("danger", "Something went wrong adding wallet")
		return c.Redirect(http.StatusFound, default_redirect_url)
	}

	c.Flash().Add("success", "Wallet was added successfully")
	return c.Redirect(http.StatusFound, default_redirect_url)
}

// Remove a wallet from user integrated wallets
func UserWalletsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	default_redirect_url := "/users/me/wallets"

	user_id := c.Session().Get("current_user_id")
	user := models.User{}
	if err := models.DB.Find(&user, user_id); err != nil {
		c.Flash().Add("danger", "User not found")
		return c.Redirect(http.StatusFound, default_redirect_url)
	}

	user_wallet := models.UsersWallet{}
	if err := tx.Where("user_id = ?", user.ID).Where("wallet_id = ?", c.Param("wallet_id")).First(&user_wallet); err != nil {
		c.Flash().Add("warning", "Wallet does not belong to user")
		return c.Redirect(http.StatusFound, default_redirect_url)
	}

	// Destroy the UserWallet
	if err := tx.Destroy(&user_wallet); err != nil {
		c.Flash().Add("danger", "Something went wrong removing wallet")
		return c.Redirect(http.StatusFound, default_redirect_url)
	}

	c.Flash().Add("success", "Wallet was removed successfully")
	return c.Redirect(http.StatusFound, default_redirect_url)
}
