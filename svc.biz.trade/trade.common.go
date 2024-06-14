package trade

// TransState 交易状态
type TransState int8

const (
	TransStateInit                TransState = 0
	TransStateConfirmed           TransState = 1
	TransStateSettlePrepared      TransState = 2
	TransStateSettling            TransState = 3
	TransStateSuccess             TransState = 4
	TransStateFailed              TransState = 5
	TransStateFailedPartially     TransState = 6
	TransStateConfirmedFailed     TransState = 7
	TransStateSettlePrepareFailed TransState = 8
)

// PayeeReceiptState 收款方收款状态
type PayeeReceiptState int8

const (
	PayeeReceiptStateInit    PayeeReceiptState = 0
	PayeeReceiptStateSuccess PayeeReceiptState = 1
	PayeeReceiptStateFailed  PayeeReceiptState = 2
)

// TransEvent 交易事件
type TransEvent string

const (
	TransEventPayConfirm    TransEvent = "pay_confirm"
	TransEventSettlePrepare TransEvent = "settle_prepare"
	TransEventSettle        TransEvent = "settle"
)
