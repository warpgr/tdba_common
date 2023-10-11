package types

const (
	OrderStatusTypeNew             string = "NEW"
	OrderStatusTypePartiallyFilled string = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          string = "FILLED"
	OrderStatusTypeCanceled        string = "CANCELED"
	OrderStatusTypePendingCancel   string = "PENDING_CANCEL"
	OrderStatusTypeRejected        string = "REJECTED"
	OrderStatusTypeExpired         string = "EXPIRED"

	OrderTypeLimit           string = "LIMIT"
	OrderTypeMarket          string = "MARKET"
	OrderTypeLimitMaker      string = "LIMIT_MAKER"
	OrderTypeStopLoss        string = "STOP_LOSS"
	OrderTypeStopLossLimit   string = "STOP_LOSS_LIMIT"
	OrderTypeTakeProfit      string = "TAKE_PROFIT"
	OrderTypeTakeProfitLimit string = "TAKE_PROFIT_LIMIT"

	SideTypeBuy  string = "BUY"
	SideTypeSell string = "SELL"

	TimeInForceTypeGTC string = "GTC"
	TimeInForceTypeIOC string = "IOC"
	TimeInForceTypeFOK string = "FOK"
)

type ExchangeNameType string

const (
	BinanceSPOTExchangeName    string = "BINANCE"
	BinanceFuturesExchangeName string = "BINANCE_PEREP"
)

type WorkingRequirements struct {
	PricePrecision    int     `json:"pricePrecision" bson:"pricePrecision"`
	QuantityPrecision int     `json:"quantityPrecision" bson:"quantityPrecision"`
	StepSize          float64 `json:"stepSize" bson:"stepSize"`
}

type AlgorithmName string

const (
	Algorithm1 AlgorithmName = "ALGORITHM_1"
)

type StableCoinNameType string

const (
	USDTStableCoin StableCoinNameType = "USDT"
	BUSDStableCoin StableCoinNameType = "BUSD"
	USDCStableCoin StableCoinNameType = "USDC"
)
