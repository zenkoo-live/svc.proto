package asset

const (
	DecrCoinRuleNormalOnly        = 1
	DecrCoinRuleLimitedOnly       = 2
	DecrCoinRuleNormalThenLimited = 3
	DecrCoinRuleLimitedThenNormal = 4
)

// TransType 交易类别
type TransType int64

const (
	TransTypeSendGiftInLive    TransType = 1
	TransTypeReceiveGiftInLive TransType = 2
	TransTypeGiftCommission    TransType = 3

	TransTypeBuyLiveTicket        TransType = 4
	TransTypeReceiveLiveTicket    TransType = 5
	TransTypeLiveTicketCommission TransType = 6

	TransTypePayLiveDurationFee        TransType = 7
	TransTypeReceiveLiveDurationFee    TransType = 8
	TransTypeLiveDurationFeeCommission TransType = 9

	TransTypeJoinAnchorFansGroup       TransType = 10
	TransTypeReceiveAnchorFansGroup    TransType = 11
	TransTypeAnchorFansGroupCommission TransType = 12

	TransTypePayBulletChat        TransType = 13
	TransTypeReceiveBulletChat    TransType = 14
	TransTypeBulletChatCommission TransType = 15

	TransTypeVipActivate TransType = 16
	TransTypeVipExtend   TransType = 17
	TransTypeVipUpgrade  TransType = 18

	TransTypeBuyRide    TransType = 19
	TransTypeBuyLuckyId TransType = 20

	TransTypeMoneyRecharge     TransType = 21
	TransTypeMoneyWithdraw     TransType = 22
	TransTypeMoneyExchangeCoin TransType = 23
)

// AssetType 账户类型
type AssetType int

const (
	//AssetTypeNone          AssetType = 0
	AssetTypeUserMoney     AssetType = 1
	AssetTypeUserCoin      AssetType = 2
	AssetTypeStreamerMoney AssetType = 3
	AssetTypeStreamerCoin  AssetType = 4
	AssetTypeUnionMoney    AssetType = 5
	AssetTypeUnionCoin     AssetType = 6
	//AssetTypeMerchantMoney AssetType = 7
	AssetTypeMerchantCoin    AssetType = 8
	AssetTypeUserFiatAccount AssetType = 9
)
