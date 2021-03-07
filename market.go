package bybit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// MarketService :
type MarketService struct {
	Client *Client
}

// OrderBookResponse :
type OrderBookResponse struct {
	CommonResponse `json:",inline"`
	Result         []OrderBookResult `json:"result"`
}

// UnmarshalJSON :
func (r *OrderBookResponse) UnmarshalJSON(data []byte) error {
	fmt.Println(string(data))
	return nil
}

// OrderBookResult :
type OrderBookResult struct {
	Symbol Symbol  `json:"symbol"`
	Price  float64 `json:"price"`
	Size   float64 `json:"size"`
	Side   Side    `json:"side"`
}

// OrderBook :
func (s *MarketService) OrderBook(symbol Symbol) (*OrderBookResponse, error) {
	var res OrderBookResponse

	params := map[string]string{
		"symbol": string(symbol),
	}

	url, err := s.Client.BuildPublicURL("/v2/public/orderBook/L2", params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListKlineParam :
type ListKlineParam struct {
	Symbol   Symbol   `json:"symbol"`
	Interval Interval `json:"interval"`
	From     int      `json:"from"`

	Limit *int `json:"limit"`
}

func (p *ListKlineParam) build() map[string]string {
	result := map[string]string{
		"symbol":   string(p.Symbol),
		"interval": string(p.Interval),
		"from":     strconv.Itoa(p.From),
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	return result
}

// ListKlineResponse :
type ListKlineResponse struct {
	CommonResponse `json:",inline"`
	Result         []ListKlineResult `json:"result"`
}

// ListKlineResult :
type ListKlineResult struct {
	Symbol   Symbol `json:"symbol"`
	Interval string `json:"interval"`
	OpenTime int    `json:"open_time`
	Open     string `json:"open"`
	High     string `json:"high"`
	Low      string `json:"low"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
	Turnover string `json:"turnover"`
}

// ListKline :
func (s *MarketService) ListKline(param ListKlineParam) (*ListKlineResponse, error) {
	var res ListKlineResponse

	url, err := s.Client.BuildPublicURL("/v2/public/kline/list", param.build())
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TickersResponse :
type TickersResponse struct {
	CommonResponse `json:",inline"`
	Result         []TickersResult `json:"result"`
}

// TickersResult :
type TickersResult struct {
	Symbol               Symbol        `json:"symbol"`
	BidPrice             string        `json:"bid_price"`
	AskPrice             string        `json:"ask_price"`
	LastPrice            string        `json:"last_price"`
	LastTickDirection    TickDirection `json:"last_tick_direction"`
	PrevPrice24h         string        `json:"prev_price_24h"`
	Price24hPcnt         string        `json:"price_24h_pcnt"`
	HighPrice24h         string        `json:"high_price_24h"`
	LowPrice24h          string        `json:"low_price_24h"`
	PrevPrice1h          string        `json:"prev_price_1h"`
	Price1hPcnt          string        `json:"price_1h_pcnt"`
	MarkPrice            string        `json:"mark_price"`
	IndexPrice           string        `json:"index_price"`
	OpenInterest         float64       `json:"open_interest"`
	OpenValue            string        `json:"open_value"`
	TotalTurnover        string        `json:"total_turnover"`
	Turnover24h          string        `json:"turnover_24h"`
	TotalVolume          float64       `json:"total_volume"`
	Volume24h            float64       `json:"volume_24h"`
	FundingRate          string        `json:"funding_rate"`
	PredictedFundingRate string        `json:"predicted_funding_rate"`
	NextFundingTime      string        `json:"next_funding_time"`
	CountdownHour        float64       `json:"countdown_hour"`
}

// Tickers :
func (s *MarketService) Tickers(symbol Symbol) (*TickersResponse, error) {
	var res TickersResponse

	params := map[string]string{
		"symbol": string(symbol),
	}

	url, err := s.Client.BuildPublicURL("/v2/public/tickers", params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TradingRecordsParam :
type TradingRecordsParam struct {
	Symbol Symbol `json:"symbol"`

	From  *int `json:"from"`
	Limit *int `json:"limit"`
}

func (p *TradingRecordsParam) build() map[string]string {
	result := map[string]string{
		"symbol": string(p.Symbol),
	}
	if p.From != nil {
		result["from"] = strconv.Itoa(*p.From)
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	return result
}

// TradingRecordsResponse :
type TradingRecordsResponse struct {
	CommonResponse `json:",inline"`
	Result         []TradingRecordsResult `json:"result"`
}

// TradingRecordsResult :
type TradingRecordsResult struct {
	ID     float64 `json:"id"`
	Symbol Symbol  `json:"symbol"`
	Price  float64 `json:"price"`
	Qty    float64 `json:"qty"`
	Side   Side    `json:"side"`
	Time   string  `json:"time"`
}

// TradingRecords :
func (s *MarketService) TradingRecords(param TradingRecordsParam) (*TradingRecordsResponse, error) {
	var res TradingRecordsResponse

	url, err := s.Client.BuildPublicURL("/v2/public/trading-records", param.build())
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SymbolsResponse :
type SymbolsResponse struct {
	CommonResponse `json:",inline"`
	Result         []SymbolsResult `json:"result"`
}

// SymbolsResult :
type SymbolsResult struct {
	Name           string         `json:"name"`
	BaseCurrency   string         `json:"base_currency"`
	QuoteCurrency  string         `json:"quote_currency"`
	PriceScale     float64        `json:"price_scale"`
	TakerFee       string         `json:"taker_fee"`
	MakerFee       string         `json:"maker_fee"`
	LeverageFilter LeverageFilter `json:"leverage_filter"`
	PriceFilter    PriceFilter    `json:"price_filter"`
	LotSizeFilter  LotSizeFilter  `json:"lot_size_filter"`
}

// LeverageFilter :
type LeverageFilter struct {
	MinLeverage  float64 `json:"min_leverage"`
	MaxLeverage  float64 `json:"max_leverage"`
	LeverageStep string  `json:"leverage_step"`
}

// PriceFilter :
type PriceFilter struct {
	MinPrice string `json:"min_price"`
	MaxPrice string `json:"max_price"`
	TickSize string `json:"tick_size"`
}

// LotSizeFilter :
type LotSizeFilter struct {
	MaxTradingQty float64 `json:"max_trading_qty"`
	MinTradingQty float64 `json:"min_trading_qty"`
	QtyStep       float64 `json:"qty_step"`
}

// Symbols :
func (s *MarketService) Symbols() (*SymbolsResponse, error) {
	var res SymbolsResponse

	url, err := s.Client.BuildPublicURL("/v2/public/symbols", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// LiqRecordsResponse :
type LiqRecordsResponse struct {
	CommonResponse `json:",inline"`
	Result         []LiqRecordsResult `json:"result"`
}

// LiqRecordsResult :
type LiqRecordsResult struct {
	ID     float64 `json:"id"`
	Qty    float64 `json:"qty"`
	Side   Side    `json:"side"`
	Time   int     `json:"time"` // or float64
	Symbol Symbol  `json:"symbol"`
	Price  float64 `json:"price"`
}

// LiqRecordsParam :
type LiqRecordsParam struct {
	Symbol Symbol `json:"symbol"`

	From      *int `json:"from"`
	Limit     *int `json:"limit"`
	StartTime *int `json:"start_time"`
	EndTime   *int `json:"end_time"`
}

func (p *LiqRecordsParam) build() map[string]string {
	result := map[string]string{
		"symbol": string(p.Symbol),
	}
	if p.From != nil {
		result["from"] = strconv.Itoa(*p.From)
	}
	if p.Limit != nil {
		result["limit"] = strconv.Itoa(*p.Limit)
	}
	if p.StartTime != nil {
		result["start_time"] = strconv.Itoa(*p.StartTime)
	}
	if p.EndTime != nil {
		result["end_time"] = strconv.Itoa(*p.EndTime)
	}
	return result
}

// LiqRecords :
func (s *MarketService) LiqRecords(param LiqRecordsParam) (*LiqRecordsResponse, error) {
	var res LiqRecordsResponse

	url, err := s.Client.BuildPublicURL("/v2/public/liq-records", param.build())
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
