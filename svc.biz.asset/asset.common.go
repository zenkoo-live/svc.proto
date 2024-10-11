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
	TransTypeSendGiftInRoom    TransType = 1
	TransTypeReceiveGiftInRoom TransType = 2
	TransTypeGiftCommission    TransType = 3

	TransTypeBuyRoomTicket        TransType = 4
	TransTypeReceiveRoomTicket    TransType = 5
	TransTypeRoomTicketCommission TransType = 6

	TransTypePayRoomDurationFee        TransType = 7
	TransTypeReceiveRoomDurationFee    TransType = 8
	TransTypeRoomDurationFeeCommission TransType = 9

	TransTypeJoinStreamerFansGroup       TransType = 10
	TransTypeReceiveStreamerFansGroup    TransType = 11
	TransTypeStreamerFansGroupCommission TransType = 12

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

	TransTypeCoinAdminRecharge TransType = 24
	TransTypeCoinStreamerWage  TransType = 25
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
