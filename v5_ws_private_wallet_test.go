package bybit

import (
	"encoding/json"
	"testing"

	"github.com/hirokisan/bybit/v2/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestV5WebsocketPrivate_Wallet(t *testing.T) {
	t.Run("for Unifield", func(t *testing.T) {
		respBody := map[string]interface{}{
			"topic":        "wallet",
			"id":           "A7EB626A-2B97-4806-BE16-BC9699B052D5",
			"creationTime": 1742176663485,
			"data": []map[string]interface{}{
				{
					"accountIMRate":          "0",
					"accountMMRate":          "0",
					"totalEquity":            "199.97795197",
					"totalWalletBalance":     "199.97785602",
					"totalMarginBalance":     "199.97785602",
					"totalAvailableBalance":  "148.14361026",
					"totalPerpUPL":           "0",
					"totalInitialMargin":     "0",
					"totalMaintenanceMargin": "0",
					"accountType":            "UNIFIED",
					"coin": []map[string]interface{}{
						{
							"coin":                "USDT",
							"equity":              "200.00005603",
							"usdValue":            "199.97785602",
							"walletBalance":       "200.00005603",
							"availableToWithdraw": "",
							"availableToBorrow":   "",
							"borrowAmount":        "0",
							"accruedInterest":     "0",
							"totalOrderIM":        "0",
							"totalPositionIM":     "0",
							"totalPositionMM":     "0",
							"unrealisedPnl":       "0",
							"cumRealisedPnl":      "-0.38465646",
							"bonus":               "0",
							"collateralSwitch":    true,
							"marginCollateral":    true,
							"locked":              "51.84",
							"spotHedgingQty":      "0",
						},
					},
				},
			},
		}
		bytesBody, err := json.Marshal(respBody)
		require.NoError(t, err)

		server, teardown := testhelper.NewWebsocketServer(
			testhelper.WithWebsocketHandlerOption(V5WebsocketPrivatePath, bytesBody),
		)
		defer teardown()

		wsClient := NewTestWebsocketClient().
			WithBaseURL(server.URL).
			WithAuth("test", "test")

		svc, err := wsClient.V5().Private()
		require.NoError(t, err)

		require.NoError(t, svc.Subscribe())

		{
			_, err := svc.SubscribeWallet(func(response V5WebsocketPrivateWalletResponse) error {
				testhelper.Compare(t, respBody, response)
				return nil
			})
			require.NoError(t, err)
		}

		assert.NoError(t, svc.Run())
		assert.NoError(t, svc.Ping())
		assert.NoError(t, svc.Close())
	})
}
