package models

type CoinMarketCap struct {
	Data struct {
		List []struct {
			ID        int    `json:"id"`
			Name      string `json:"n"`
			UrlFormat string `json:"uf"`
			Domain    string `json:"dn"`
			AddrUrl   string `json:"addrUrl"`
			ChainId   int    `json:"chId"`
			CryptoId  int    `json:"cid"`
		} `json:"list"`
	} `json:"data"`
}