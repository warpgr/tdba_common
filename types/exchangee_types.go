package types

// ------------------- Candlestick ---------------------
// Represents fetched candlestick from exchanges.
type Candlestick struct {
	OpenTime                 int64   `json:"openTime" bson:"openTime" `
	Open                     float64 `json:"open" bson:"open"`
	High                     float64 `json:"high" bson:"high"`
	Low                      float64 `json:"low" bson:"low"`
	Close                    float64 `json:"close" bson:"close"`
	Volume                   float64 `json:"volume" bson:"volume"`
	CloseTime                int64   `json:"closeTime" bson:"closeTime"`
	QuoteAssetVolume         float64 `json:"quoteAssetVolume" bson:"quoteAssetVolume"`
	TradeNum                 int64   `json:"tradeNum" bson:"tradeNum"`
	TakerBuyBaseAssetVolume  float64 `json:"takerBuyBaseAssetVolume" bson:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume float64 `json:"takerBuyQuoteAssetVolume" bson:"takerBuyQuoteAssetVolume"`
}

// -------------------- ExchangeOrder ---------------------
// Struct for holding and sharing order information. Can use by different exchange api's.
type ExchangeOrder struct {
	Symbol                   string          `json:"symbol" bson:"symbol"`
	OrderID                  int64           `json:"orderId" bson:"orderId"`
	ClientOrderID            string          `json:"clientOrderId" bson:"clientOrderId"`
	TransactTime             int64           `json:"transactTime" bson:"transactTime"`
	Price                    string          `json:"price" bson:"price"`
	OrigQuantity             string          `json:"origQty" bson:"origQty"`
	ExecutedQuantity         string          `json:"executedQty" bson:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty" bson:"cummulativeQuoteQty"`
	IsIsolated               bool            `json:"isIsolated" bson:"isIsolated"` // for isolated margin
	Status                   OrderStatusType `json:"status" bson:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce" bson:"timeInForce"`
	Type                     OrderType       `json:"type" bson:"type"`
	Side                     SideType        `json:"side" bson:"side"`
	// for order response is set to FULL
	Fills                 []*Fill `json:"fills" bson:"fills"`
	MarginBuyBorrowAmount string  `json:"marginBuyBorrowAmount" bson:"marginBuyBorrowAmount"` // for margin
	MarginBuyBorrowAsset  string  `json:"marginBuyBorrowAsset" bson:"marginBuyBorrowAsset"`
}

type Fill struct {
	TradeID         int64  `json:"tradeId" bson:"tradeId"`
	Price           string `json:"price" bson:"price"`
	Quantity        string `json:"qty" bson:"qty"`
	Commission      string `json:"commission" bson:"commission"`
	CommissionAsset string `json:"commissionAsset" bson:"commissionAsset"`
}

// Order status type.
type OrderStatusType string

// Order type.
type OrderType string

// Order side type.
type SideType string

// Time in force type.
type TimeInForceType string

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
