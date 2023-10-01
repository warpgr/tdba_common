package types

// --------------- PairOrderBehavior -------------------
// Represents behavior of trading pair in exchange.
type PairOrderBehavior struct {
	Entry      float64  `json:"entry" bson:"entry"`           // Estimated Entry price.
	StopLoss   float64  `json:"stopLoss" bson:"stopLoss"`     // Estimated StopLoss price.
	TakeProfit float64  `json:"takeProfit" bson:"takeProfit"` // Estimated TakeProfit price.
	Side       SideType `json:"side" bson:"side"`             // Position type BUY/.SELL.
	Quantity   float64  `json:"quantity" bson:"quantity"`     // Quantity of order.
}

// ------------------- Estimation --------------
// Struct for sharing algorithm computation result by given configurations.
type Estimation struct {
	// Position credentials
	Behavior    []PairOrderBehavior `json:"behavior" bson:"workingPair"`
	DualPair    bool                `json:"dualPair" bson:"workingPair"`
	WorkingPair string              `json:"workingPair" bson:"workingPair"` // Pair can be dual or not
	// Estimator properties
	EstimatorWeight int              `json:"estimatorWeight" bson:"estimatorName"` // Estimator weight.
	TimeFrame       TimeFrame        `json:"timeFrame" bson:"estimatorName"`       // In which timestamp used for this estimation
	EstimatorName   string           `json:"estimatorName" bson:"estimatorName"`   // Algorithm name which estimated this
	EstimatedAt     int64            `json:"estimatedAt" bson:"estimatedAt"`       // Unix timestamp since epoch milliseconds
	ExpiresAt       int64            `json:"expiresAt" bson:"expiresAt"`           // Expiration time point unix milliseconds.
	Reason          string           `json:"reason" bson:"reason"`                 // Reason of estimation.
	Exchange        ExchangeNameType `json:"exchange" bson:"exchange"`             // On which exchange must execute this estimation order.
	HaveOwnHandler  bool             `json:"haveOwnHandler" bson:"haveOwnHandler"` // If have own handler for orders.
	HandlerName     string           `json:"handlerName" bson:"handlerName"`       // Handler name specifies handler function name.
}

type EstimationStatus string

const (
	EstimationStatusInvalid    EstimationStatus = "INVALID"
	EstimationStatusCreated    EstimationStatus = "CREATED"
	EstimationStatusProcessing EstimationStatus = "PROCESSING"
	EstimationStatusCanceled   EstimationStatus = "CANCELED"
	EstimationStatusFailed     EstimationStatus = "FAILED"
	EstimationStatusDone       EstimationStatus = "DONE"
)
