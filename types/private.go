package types

//Fund
type FundChangeResult struct {
	Amount       string `json:"amount"`
	Balance      string `json:"balance"`
	CreateTime   string `json:"createTime"`
	CurrencyId   string `json:"currencyId"`
	FreezeAmount string `json:"freezeAmount"`
	Id           string `json:"id"`
	ModifyTime   string `json:"modifyTime"`
	UserId       string `json:"userId"`
}

type FundBalanceResult struct {
	AllUnrealizedPnl string `json:"allUnrealizedPnl"`
	AllMargin        string `json:"allMargin"`
	Amount           string `json:"amount"`
	FreezeAmount     string `json:"freezeAmount"`
	CurrencyName     string `json:"currencyName"`
	CreateTime       string `json:"createTime"`
	Id               string `json:"id"`
	CurrencyId       string `json:"currencyId"`
	AccountBalance   string `json:"accountBalance"`
	UserId           string `json:"userId"`
}

type Bill struct {
	PageSize int                  `json:"pageSize"`
	PageNum  int                  `json:"pageNum"`
	List     []FundBillItemResult `json:"list"`
}

type FundBillItemResult struct {
	UserId             string `json:"userId"`
	CurrencyId         string `json:"currencyId"`
	FundId             string `json:"fundId"`
	FreezeId           string `json:"freezeId"`
	BillType           int    `json:"type"`
	ChangeAmount       string `json:"changeAmount"`
	FeeRate            string `json:"feeRate"`
	Fee                string `json:"fee"`
	OperatorId         string `json:"operatorId"`
	BeforeAmount       string `json:"beforeAmount"`
	BeforeFreezeAmount string `json:"beforeFreezeAmount"`
	MarketId           string `json:"marketId"`
	OutsideId          string `json:"outsideId"`
	Id                 string `json:"id"`
	IsIn               int    `json:"isIn"`
	Available          string `json:"available"`
	Unit               string `json:"unit"`
	CreateTime         string `json:"createTime"`
	ModifyTime         string `json:"modifyTime"`
	Extend             string `json:"extend"`
}

type AssetChangeResult struct {
	AccountBalance          string `json:"accountBalance"`
	AccountBalanceConvert   string `json:"accountBalanceConvert"`
	AllMargin               string `json:"allMargin"`
	AllMarginConvert        string `json:"allMarginConvert"`
	AllUnrealizedPnl        string `json:"allUnrealizedPnl"`
	AllUnrealizedPnlConvert string `json:"allUnrealizedPnlConvert"`
	Available               string `json:"available"`
	AvailableConvert        string `json:"availableConvert"`
	ConvertUnit             string `json:"convertUnit"`
	Freeze                  string `json:"freeze"`
	FreezeConvert           string `json:"freezeConvert"`
	FuturesAccountType      int    `json:"futuresAccountType"`
	Percent                 string `json:"percent"`
	Unit                    string `json:"unit"`
	UserId                  string `json:"userId"`
}

type AssetInfoResult struct {
	AccountBalanceConvert   string `json:"accountBalanceConvert"`
	AllMarginConvert        string `json:"allMarginConvert"`
	AvailableConvert        string `json:"availableConvert"`
	Available               string `json:"available"`
	AllUnrealizedPnlConvert string `json:"allUnrealizedPnlConvert"`
	Percent                 string `json:"percent"`
	UserId                  string `json:"userId"`
	AllMargin               string `json:"allMargin"`
	AllUnrealizedPnl        string `json:"allUnrealizedPnl"`
	Unit                    string `json:"unit"`
	ConvertUnit             string `json:"convertUnit"`
	Freeze                  string `json:"freeze"`
	FreezeConvert           string `json:"freezeConvert"`
	AccountBalance          string `json:"accountBalance"`
}

//Positions
type PositionsChangeResult struct {
	Amount           string `json:"amount"`
	AutoLightenRatio string `json:"autoLightenRatio"`
	AvgPrice         string `json:"avgPrice"`
	ContractType     int    `json:"contractType"`
	CreateTime       string `json:"createTime"`
	FreezeAmount     string `json:"freezeAmount"`
	Id               string `json:"id"`
	Leverage         string `json:"leverage"`
	LiquidateLevel   string `json:"liquidateLevel"`
	LiquidatePrice   string `json:"liquidatePrice"`
	MaintainMargin   string `json:"maintainMargin"`
	Margin           string `json:"margin"`
	MarginBalance    string `json:"marginBalance"`
	MarginMode       int    `json:"marginMode"`
	MarginRate       string `json:"marginRate"`
	MarketId         string `json:"marketId"`
	Symbol           string `json:"symbol"` //TODO:marketName
	NominalValue     string `json:"nominalValue"`
	OriginId         string `json:"originId"`
	PositionsMode    int    `json:"positionsMode"`
	ReturnRate       string `json:"returnRate"`
	Side             int    `json:"side"`
	Status           int    `json:"status"`
	UnrealizedPnl    string `json:"unrealizedPnl"`
	UserId           string `json:"userId"`
}
type GetPositionsResult struct {
	Leverage         int    `json:"leverage"`
	ReturnRate       string `json:"returnRate"`
	AvgPrice         string `json:"avgPrice"`
	ContractType     int    `json:"contractType"`
	MarginMode       int    `json:"marginMode"`
	MarketId         string `json:"marketId"`
	MarginRate       string `json:"marginRate"`
	FreezeAmount     string `json:"freezeAmount"`
	OriginId         string `json:"originId"`
	autoLightenRatio string `json:"autoLightenRatio"`
	Id               string `json:"id"`
	MarginBalance    string `json:"marginBalance"`
	Amount           string `json:"amount"`
	Margin           string `json:"margin"`
	Side             int    `json:"side"`
	LiquidatePrice   string `json:"liquidatePrice"`
	UserId           string `json:"userId"`
	Symbol           string `json:"symbol"` //TODO:marketName
	CreateTime       string `json:"createTime"`
	UnrealizedPnl    string `json:"unrealizedPnl"`
	LiquidateLevel   int    `json:"liquidateLevel"`
	PositionsMode    int    `json:"positionsMode"`
	MaintainMargin   string `json:"maintainMargin"`
	NominalValue     string `json:"nominalValue"`
	Status           int    `json:"status"`

	ModifyTime string `json:"modifyTime,omitempty"`
	Extend     string `json:"extend,omitempty"`
}

type MarginInfoResult struct {
	PositionsId    string `json:"positionsId"`
	MaxAdd         string `json:"maxAdd"`
	MaxSub         string `json:"maxSub"`
	LiquidatePrice string `json:"liquidatePrice"`
}

type UpdateMarginResult struct {
	Leverage        int    `json:"leverage"`
	AvgPrice        string `json:"avgPrice"`
	BankruptcyPrice string `json:"bankruptcyPrice"`
	MarginMode      int    `json:"marginMode"`
	MarketId        string `json:"marketId"`
	MarginRate      string `json:"marginRate"`
	FreezeAmount    string `json:"freezeAmount"`
	ModifyTime      string `json:"modifyTime"`
	OriginId        string `json:"originId"`
	Id              string `json:"id"`
	MarginBalance   string `json:"marginBalance"`
	Amount          string `json:"amount"`
	Margin          string `json:"margin"`
	Side            int    `json:"side"`
	LiquidatePrice  string `json:"liquidatePrice"`
	KeyMark         string `json:"keyMark"`
	UserId          string `json:"userId"`
	MarketName      string `json:"marketName"`
	CreateTime      string `json:"createTime"`
	UnrealizedPnl   string `json:"unrealizedPnl"`
	LiquidateLevel  int    `json:"liquidateLevel"`
	PositionsMode   int    `json:"positionsMode"`
	MaintainMargin  string `json:"maintainMargin"`
	NominalValue    string `json:"nominalValue"`
	Open            bool   `json:"open"`
	Status          int    `json:"status"`
}

type GetSettingResult struct {
	Leverage         int     `json:"leverage"`
	CreateTime       string  `json:"createTime"`
	PositionsMode    int     `json:"positionsMode"`
	Id               string  `json:"id"`
	MarginMode       int     `json:"marginMode"`
	KeyMark          string  `json:"keyMark"`
	UserId           string  `json:"userId"`
	MarketId         string  `json:"marketId"` //TODO:最好返回symbol，MarketID不直观
	EnableAutoAppend int     `json:"enableAutoAppend"`
	MaxAppendAmount  float64 `json:"maxAppendAmount,string"`
	MarginCoins      string  `json:"marginCoins"`
}

type PositionSettingResult struct {
	UserId           string  `json:"userId"`
	MarketId         string  `json:"marketId"`
	Leverage         int     `json:"leverage"`
	MarginMode       int     `json:"marginMode"`
	PositionsMode    int     `json:"positionsMode"`
	MaxAppendAmount  float64 `josn:"maxAppendAmount"`
	EnableAutoAppend int     `json:"enableAutoAppend"`
	MarginCoins      string  `json:"marginCoins"`

	Id         string `json:"id"`
	CreateTime string `json:"createTime"`
	ModifyTime string `json:"modifyTime"`
}

type GetNominalValueResult struct {
	Side                  int    `json:"side"`
	OpenOrderNominalValue string `json:"openOrderNominalValue"`
	NominalValue          string `json:"nominalValue"`
	MarketId              string `json:"marketId"` //TODO:最好返回symbol，MarketID不直观

	Symbol string `json:"symbol,omitempty"`
}

//Order
type OrderChangeResult struct {
	Action          int    `json:"action"`
	Amount          string `json:"amount"`
	AvailableAmount string `json:"availableAmount"`
	AvailableValue  string `json:"availableValue"`
	AvgPrice        string `json:"avgPrice"`
	CanCancel       bool   `json:"canCancel"`
	CancelStatus    int    `json:"cancelStatus"`
	CreateTime      string `json:"createTime"`
	EntrustType     int    `json:"entrustType"`
	Id              string `json:"id"`
	Leverage        int    `json:"leverage"`
	MarketId        string `json:"marketId"`
	Price           string `json:"price"`
	ShowStatus      int    `json:"showStatus"`
	Side            int    `json:"side"`
	SourceType      int    `json:"sourceType"`
	Status          int    `json:"status"`
	TradeAmount     string `json:"tradeAmount"`
	TradeValue      string `json:"tradeValue"`
	OrderType       int    `json:"type"`
	UserId          string `json:"userId"`
	Value           string `json:"value"`
}

type GetOrderResult struct {
	Leverage        int    `json:"leverage"`
	AvgPrice        string `json:"avgPrice"`
	CancelStatus    int    `json:"cancelStatus"`
	OrderType       int    `json:"type"`
	MarketId        string `json:"marketId"`
	ModifyTime      string `json:"modifyTime"`
	AvailableAmount string `json:"availableAmount"`
	Price           string `json:"price"`
	Action          int    `json:"action"`
	Id              string `json:"id"`
	Value           string `json:"value"`
	Amount          string `json:"amount"`
	Margin          string `json:"margin"`
	Side            int    `json:"side"`
	AvailableValue  string `json:"availableValue"`
	TradeValue      string `json:"tradeValue"`
	ShowStatus      int    `json:"showStatus"`
	UserId          string `json:"userId"`
	TradeAmount     string `json:"tradeAmount"`
	CreateTime      string `json:"createTime"`
	SourceType      int    `json:"sourceType"`
	OrderCode       string `json:"orderCode"`
	EntrustType     int    `json:"entrustType"`
	CanCancel       bool   `json:"canCancel"`
	Status          int    `json:"status"`
}

type GetAllOrderResult struct {
	PageSize int              `json:"pageSize"`
	PageNum  int              `json:"pageNum"`
	List     []GetOrderResult `json:"list"`
}

type BatchCancelOrderItem struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data string `json:"data"`
}

type TradeList struct {
	Amount      string `json:"amount"`
	CreateTime  string `json:"createTime"`
	FeeAmount   string `json:"feeAmount"`
	FeeCurrency string `json:"feeCurrency"`
	Maker       bool   `json:"maker"`
	OrderId     string `json:"orderId"`
	Price       string `json:"price"`
	RelizedPnl  string `json:"relizedPnl"`
	Side        int    `json:"side"`
	UserId      string `json:"userId"`
}

type TradeListResult struct {
	PageSize int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
	List     []TradeList `json:"list"`
}

type TradeHistory struct {
	FeeAmount   string `json:"feeAmount"`
	Amount      string `json:"amount"`
	Side        int    `json:"side"`
	CreateTime  string `json:"createTime"`
	OrderId     string `json:"orderId"`
	Price       string `json:"price"`
	FeeCurrency string `json:"feeCurrency"`
	Maker       bool   `json:"maker"`
	RelizedPnl  string `json:"relizedPnl"`
	UserId      string `json:"userId"`
}
type TradeHistoryResult struct {
	PageSize int            `json:"pageSize"`
	PageNum  int            `json:"pageNum"`
	List     []TradeHistory `json:"list"`
}

// /Server/api/v2/Fund/getAccount
type GetAccountResult struct {
	Account Account  `json:"account"`
	Assets  []Assets `json:"assets"`
}
type Account struct {
	AccountBalance          string `json:"accountBalance"`
	AllMargin               string `json:"allMargin"`
	Available               string `json:"available"`
	Freeze                  string `json:"freeze"`
	AllUnrealizedPnl        string `json:"allUnrealizedPnl"`
	AccountBalanceConvert   string `json:"accountBalanceConvert"`
	AllMarginConvert        string `json:"allMarginConvert"`
	AvailableConvert        string `json:"availableConvert"`
	FreezeConvert           string `json:"freezeConvert"`
	AllUnrealizedPnlConvert string `json:"allUnrealizedPnlConvert"`
	ConvertUnit             string `json:"convertUnit"`
	Unit                    string `json:"unit"`
	Percent                 string `json:"percent"`
}

type Assets struct {
	UserId           string `json:"userId"`
	CurrencyId       string `json:"currencyId"`
	Amount           string `json:"amount"`
	FreezeAmount     string `json:"freezeAmount"`
	Id               string `json:"id"`
	AccountBalance   string `json:"accountBalance"`
	AllUnrealizedPnl string `json:"allUnrealizedPnl"`
	AllMargin        string `json:"allMargin"`
	CreateTime       string `json:"createTime"`
	ModifyTime       string `json:"modifyTime"`
	Extend           string `json:"extend,omitempty"`
}

//下单的参数
type Order struct {
	Symbol        string  `json:"symbol"`
	Side          int     `json:"side"`
	Action        int     `json:"action"`
	EntrustType   int     `json:"entrustType"`
	Amount        float64 `json:"amount"`
	Price         float64 `json:"price"`
	ClientOrderId string  `json:"clientOrderId,omitempty"`
	MatchPrice    int     `json:"matchPrice,omitempty"`
}

type BatchOrderItemResult struct {
	Code int    `json:""sCode"`
	Msg  string `json:""sMsg"`

	OrderId       string `json:""orderId,omitempty"`
	ClientOrderId string `json:""clientOrderId,omitempty"`
}

type MarginHistoryResult struct {
	PageSize int                 `json:"pageSize"`
	PageNum  int                 `json:"pageNum"`
	List     []MarginHistoryItem `json:"list"`
}

type MarginHistoryItem struct {
	Symbol       string  `json:"symbol"`
	Asset        string  `json:"asset"`
	Amount       float64 `json:"amount",string`
	MarginType   int     `json:"type"`
	IsAuto       int     `json:"isAuto"`
	ContractType int     `json:"contractType"`
	PositionSide string  `json:"positionSide"`
	CreateTime   int     `json:"createTime,string"`
}

type SetMarginResult struct {
	Id               string  `json:"id"`
	PositionsMode    int     `json:"positionsMode"`
	UserId           string  `json:"userId"`
	KeyMark          string  `json:"keyMark"`
	Leverage         int     `json:"leverage"`
	MarginMode       int     `json:"marginMode"`
	MarketId         int     `json:"marketId,string"`
	EnableAutoAppend int     `json:"enableAutoAppend"`
	MaxAppendAmount  float64 `json:"maxAppendAmount,string"`
	MarginCoins      string  `json:"marginCoins"`
	CreateTime       int     `json:"createTime,string"`
	ModifyTime       int     `json:"modifyTime,string"`
}

type AlgoOrderResult struct {
	PageSize int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
	List     []AlgoOrder `json:"list"`
}

type AlgoOrder struct {
	Id            string  `json:"id"`
	MarketId      int     `json:"marketId,string"`
	UserId        string  `json:"userId"`
	Side          int     `json:"side"`
	OrderType     int     `json:"orderType"`
	Amount        float64 `json:"amount"`
	TriggerPrice  float64 `json:"triggerPrice"`
	AlgoPrice     float64 `json:"algoPrice"`
	TriggerTime   int64   `json:"triggerTime"`
	TradedAmount  float64 `json:"tradedAmount"`
	PriceType     int     `json:"priceType"`
	BizType       int     `json:"bizType"`
	Leverage      int     `json:"leverage"`
	Status        int     `json:"status"`
	SourceType    int     `json:"sourceType"`
	ClientOrderId string  `json:"clientOrderId"`
}
