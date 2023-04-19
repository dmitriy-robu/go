package requests

type StoreUserSteamTradeURL struct {
	URL string `json:"steam_trade_url" binding:"required,url,max=255"`
}
