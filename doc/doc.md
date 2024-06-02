# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [svc.biz.account/account.proto](#svc-biz-account_account-proto)
    - [AddManagerReq](#svc-biz-account-AddManagerReq)
    - [AddManagerResp](#svc-biz-account-AddManagerResp)
    - [AddStreamerReq](#svc-biz-account-AddStreamerReq)
    - [AddStreamerResp](#svc-biz-account-AddStreamerResp)
    - [AddUnionReq](#svc-biz-account-AddUnionReq)
    - [AddUnionResp](#svc-biz-account-AddUnionResp)
    - [AddViewerReq](#svc-biz-account-AddViewerReq)
    - [AddViewerResp](#svc-biz-account-AddViewerResp)
    - [DeleteManagerReq](#svc-biz-account-DeleteManagerReq)
    - [DeleteManagerResp](#svc-biz-account-DeleteManagerResp)
    - [DeleteStreamerReq](#svc-biz-account-DeleteStreamerReq)
    - [DeleteStreamerResp](#svc-biz-account-DeleteStreamerResp)
    - [DeleteUnionReq](#svc-biz-account-DeleteUnionReq)
    - [DeleteUnionResp](#svc-biz-account-DeleteUnionResp)
    - [DeleteViewerReq](#svc-biz-account-DeleteViewerReq)
    - [DeleteViewerResp](#svc-biz-account-DeleteViewerResp)
    - [GetManagerReq](#svc-biz-account-GetManagerReq)
    - [GetManagerResp](#svc-biz-account-GetManagerResp)
    - [GetStreamerReq](#svc-biz-account-GetStreamerReq)
    - [GetStreamerResp](#svc-biz-account-GetStreamerResp)
    - [GetUnionReq](#svc-biz-account-GetUnionReq)
    - [GetUnionResp](#svc-biz-account-GetUnionResp)
    - [GetViewerReq](#svc-biz-account-GetViewerReq)
    - [GetViewerResp](#svc-biz-account-GetViewerResp)
    - [InitDBResp](#svc-biz-account-InitDBResp)
    - [ListManagersReq](#svc-biz-account-ListManagersReq)
    - [ListManagersResp](#svc-biz-account-ListManagersResp)
    - [ListStreamersReq](#svc-biz-account-ListStreamersReq)
    - [ListStreamersResp](#svc-biz-account-ListStreamersResp)
    - [ListUnionsReq](#svc-biz-account-ListUnionsReq)
    - [ListUnionsResp](#svc-biz-account-ListUnionsResp)
    - [ListViewersReq](#svc-biz-account-ListViewersReq)
    - [ListViewersResp](#svc-biz-account-ListViewersResp)
    - [Manager](#svc-biz-account-Manager)
    - [Streamer](#svc-biz-account-Streamer)
    - [Union](#svc-biz-account-Union)
    - [UpdateManagerReq](#svc-biz-account-UpdateManagerReq)
    - [UpdateManagerResp](#svc-biz-account-UpdateManagerResp)
    - [UpdateStreamerReq](#svc-biz-account-UpdateStreamerReq)
    - [UpdateStreamerResp](#svc-biz-account-UpdateStreamerResp)
    - [UpdateUnionReq](#svc-biz-account-UpdateUnionReq)
    - [UpdateUnionResp](#svc-biz-account-UpdateUnionResp)
    - [UpdateViewerReq](#svc-biz-account-UpdateViewerReq)
    - [UpdateViewerResp](#svc-biz-account-UpdateViewerResp)
    - [Viewer](#svc-biz-account-Viewer)
  
    - [Account](#svc-biz-account-Account)
  
- [svc.biz.trade/trade.proto](#svc-biz-trade_trade-proto)
    - [BuyLiveTicketReq](#svc-biz-trade-BuyLiveTicketReq)
    - [BuyLiveTicketResp](#svc-biz-trade-BuyLiveTicketResp)
    - [BuyLuckyIdInfo](#svc-biz-trade-BuyLuckyIdInfo)
    - [BuyLuckyIdReq](#svc-biz-trade-BuyLuckyIdReq)
    - [BuyLuckyIdResp](#svc-biz-trade-BuyLuckyIdResp)
    - [BuyRideInfo](#svc-biz-trade-BuyRideInfo)
    - [BuyRideReq](#svc-biz-trade-BuyRideReq)
    - [BuyRideResp](#svc-biz-trade-BuyRideResp)
    - [GiftInfo](#svc-biz-trade-GiftInfo)
    - [JoinAnchorFansGroupInfo](#svc-biz-trade-JoinAnchorFansGroupInfo)
    - [JoinAnchorFansGroupReq](#svc-biz-trade-JoinAnchorFansGroupReq)
    - [JoinAnchorFansGroupResp](#svc-biz-trade-JoinAnchorFansGroupResp)
    - [LiveDurationFeeInfo](#svc-biz-trade-LiveDurationFeeInfo)
    - [LiveInfo](#svc-biz-trade-LiveInfo)
    - [LiveTicketInfo](#svc-biz-trade-LiveTicketInfo)
    - [MoneyExchangeCoinInfo](#svc-biz-trade-MoneyExchangeCoinInfo)
    - [MoneyExchangeCoinReq](#svc-biz-trade-MoneyExchangeCoinReq)
    - [MoneyExchangeCoinResp](#svc-biz-trade-MoneyExchangeCoinResp)
    - [MoneyRechargeReq](#svc-biz-trade-MoneyRechargeReq)
    - [MoneyRechargeReq.PayInfo](#svc-biz-trade-MoneyRechargeReq-PayInfo)
    - [MoneyRechargeResp](#svc-biz-trade-MoneyRechargeResp)
    - [MoneyResult](#svc-biz-trade-MoneyResult)
    - [MoneyWithdrawReq](#svc-biz-trade-MoneyWithdrawReq)
    - [MoneyWithdrawResp](#svc-biz-trade-MoneyWithdrawResp)
    - [PayBulletChatInfo](#svc-biz-trade-PayBulletChatInfo)
    - [PayBulletChatReq](#svc-biz-trade-PayBulletChatReq)
    - [PayBulletChatResp](#svc-biz-trade-PayBulletChatResp)
    - [PayLiveDurationFeeReq](#svc-biz-trade-PayLiveDurationFeeReq)
    - [PayLiveDurationFeeResp](#svc-biz-trade-PayLiveDurationFeeResp)
    - [SendGiftInLiveReq](#svc-biz-trade-SendGiftInLiveReq)
    - [SendGiftInLiveResp](#svc-biz-trade-SendGiftInLiveResp)
    - [TradeInfo](#svc-biz-trade-TradeInfo)
    - [TradeResult](#svc-biz-trade-TradeResult)
    - [UserCoinValue](#svc-biz-trade-UserCoinValue)
    - [UserInfo](#svc-biz-trade-UserInfo)
    - [VipActivateInfo](#svc-biz-trade-VipActivateInfo)
    - [VipActivateReq](#svc-biz-trade-VipActivateReq)
    - [VipActivateResp](#svc-biz-trade-VipActivateResp)
    - [VipExtendInfo](#svc-biz-trade-VipExtendInfo)
    - [VipExtendReq](#svc-biz-trade-VipExtendReq)
    - [VipExtendResp](#svc-biz-trade-VipExtendResp)
    - [VipUpgradeInfo](#svc-biz-trade-VipUpgradeInfo)
    - [VipUpgradeReq](#svc-biz-trade-VipUpgradeReq)
    - [VipUpgradeResp](#svc-biz-trade-VipUpgradeResp)
  
    - [Trade](#svc-biz-trade-Trade)
  
- [svc.infra.static/static.proto](#svc-infra-static_static-proto)
    - [InitDBResp](#svc-infra-static-InitDBResp)
  
    - [Static](#svc-infra-static-Static)
  
- [svc.biz.room/category.proto](#svc-biz-room_category-proto)
    - [CategoryInfo](#svc-biz-room-CategoryInfo)
    - [CreateCategoryReq](#svc-biz-room-CreateCategoryReq)
    - [CreateCategoryResp](#svc-biz-room-CreateCategoryResp)
    - [DeleteCategoryReq](#svc-biz-room-DeleteCategoryReq)
    - [GetCategoryReq](#svc-biz-room-GetCategoryReq)
    - [GetCategoryResp](#svc-biz-room-GetCategoryResp)
    - [ListCategoryReq](#svc-biz-room-ListCategoryReq)
    - [ListCategoryResp](#svc-biz-room-ListCategoryResp)
    - [ListCategoryTreeReq](#svc-biz-room-ListCategoryTreeReq)
    - [ListCategoryTreeResp](#svc-biz-room-ListCategoryTreeResp)
    - [UpdateCategoryReq](#svc-biz-room-UpdateCategoryReq)
  
    - [Category](#svc-biz-room-Category)
  
- [svc.biz.room/live.proto](#svc-biz-room_live-proto)
    - [GetLiveReq](#svc-biz-room-GetLiveReq)
    - [GetLiveResp](#svc-biz-room-GetLiveResp)
    - [ListLiveReq](#svc-biz-room-ListLiveReq)
    - [ListLiveResp](#svc-biz-room-ListLiveResp)
    - [LiveInfo](#svc-biz-room-LiveInfo)
    - [MGetLiveReq](#svc-biz-room-MGetLiveReq)
    - [MGetLiveResp](#svc-biz-room-MGetLiveResp)
    - [MGetLiveResp.ItemsEntry](#svc-biz-room-MGetLiveResp-ItemsEntry)
  
    - [Live](#svc-biz-room-Live)
  
- [svc.biz.room/room.proto](#svc-biz-room_room-proto)
    - [CreateRoomReq](#svc-biz-room-CreateRoomReq)
    - [CreateRoomResp](#svc-biz-room-CreateRoomResp)
    - [GetOnlineRoomListReq](#svc-biz-room-GetOnlineRoomListReq)
    - [GetOnlineRoomListResp](#svc-biz-room-GetOnlineRoomListResp)
    - [GetRoomByStreamerIDReq](#svc-biz-room-GetRoomByStreamerIDReq)
    - [GetRoomByStreamerIDResp](#svc-biz-room-GetRoomByStreamerIDResp)
    - [GetRoomListReq](#svc-biz-room-GetRoomListReq)
    - [GetRoomListResp](#svc-biz-room-GetRoomListResp)
    - [GetRoomReq](#svc-biz-room-GetRoomReq)
    - [GetRoomResp](#svc-biz-room-GetRoomResp)
    - [MGetRoomsByStreamerIDsReq](#svc-biz-room-MGetRoomsByStreamerIDsReq)
    - [MGetRoomsByStreamerIDsResp](#svc-biz-room-MGetRoomsByStreamerIDsResp)
    - [MGetRoomsByStreamerIDsResp.ItemsEntry](#svc-biz-room-MGetRoomsByStreamerIDsResp-ItemsEntry)
    - [MGetRoomsReq](#svc-biz-room-MGetRoomsReq)
    - [MGetRoomsResp](#svc-biz-room-MGetRoomsResp)
    - [MGetRoomsResp.ItemsEntry](#svc-biz-room-MGetRoomsResp-ItemsEntry)
    - [RoomInfo](#svc-biz-room-RoomInfo)
    - [StartLiveReq](#svc-biz-room-StartLiveReq)
    - [StartLiveResp](#svc-biz-room-StartLiveResp)
    - [StopLiveReq](#svc-biz-room-StopLiveReq)
    - [StopLiveResp](#svc-biz-room-StopLiveResp)
    - [Stream](#svc-biz-room-Stream)
    - [StreamPull](#svc-biz-room-StreamPull)
    - [StreamPush](#svc-biz-room-StreamPush)
    - [UpdateRoomReq](#svc-biz-room-UpdateRoomReq)
  
    - [LiveDisplayType](#svc-biz-room-LiveDisplayType)
    - [LiveStatus](#svc-biz-room-LiveStatus)
    - [SortType](#svc-biz-room-SortType)
  
    - [Room](#svc-biz-room-Room)
  
- [svc.infra.setting/setting.proto](#svc-infra-setting_setting-proto)
    - [AddConfigurationReq](#svc-infra-setting-AddConfigurationReq)
    - [AddConfigurationResp](#svc-infra-setting-AddConfigurationResp)
    - [Configuration](#svc-infra-setting-Configuration)
    - [DeleteConfigurationReq](#svc-infra-setting-DeleteConfigurationReq)
    - [DeleteConfigurationResp](#svc-infra-setting-DeleteConfigurationResp)
    - [GetConfigurationReq](#svc-infra-setting-GetConfigurationReq)
    - [GetConfigurationResp](#svc-infra-setting-GetConfigurationResp)
    - [InitDBResp](#svc-infra-setting-InitDBResp)
    - [ListConfigurationsReq](#svc-infra-setting-ListConfigurationsReq)
    - [ListConfigurationsResp](#svc-infra-setting-ListConfigurationsResp)
    - [SettingGreetingReq](#svc-infra-setting-SettingGreetingReq)
    - [SettingGreetingResp](#svc-infra-setting-SettingGreetingResp)
    - [UpdateConfigurationReq](#svc-infra-setting-UpdateConfigurationReq)
    - [UpdateConfigurationResp](#svc-infra-setting-UpdateConfigurationResp)
  
    - [Setting](#svc-infra-setting-Setting)
  
- [svc.web.viewer/viewer.proto](#svc-web-viewer_viewer-proto)
    - [Viewer](#svc-web-viewer-Viewer)
  
- [svc.biz.asset/asset.proto](#svc-biz-asset_asset-proto)
    - [ChangeMerchantCoinResp](#svc-biz-asset-ChangeMerchantCoinResp)
    - [ChangeStreamerCoinResp](#svc-biz-asset-ChangeStreamerCoinResp)
    - [ChangeStreamerMoneyResp](#svc-biz-asset-ChangeStreamerMoneyResp)
    - [ChangeUnionCoinResp](#svc-biz-asset-ChangeUnionCoinResp)
    - [ChangeUnionMoneyResp](#svc-biz-asset-ChangeUnionMoneyResp)
    - [ChangeUserCoinResp](#svc-biz-asset-ChangeUserCoinResp)
    - [ChangeUserMoneyResp](#svc-biz-asset-ChangeUserMoneyResp)
    - [DecrMerchantCoinReq](#svc-biz-asset-DecrMerchantCoinReq)
    - [DecrStreamerCoinReq](#svc-biz-asset-DecrStreamerCoinReq)
    - [DecrStreamerMoneyReq](#svc-biz-asset-DecrStreamerMoneyReq)
    - [DecrUnionCoinReq](#svc-biz-asset-DecrUnionCoinReq)
    - [DecrUnionMoneyReq](#svc-biz-asset-DecrUnionMoneyReq)
    - [DecrUserCoinReq](#svc-biz-asset-DecrUserCoinReq)
    - [DecrUserMoneyReq](#svc-biz-asset-DecrUserMoneyReq)
    - [GetMerchantCoinMultiReq](#svc-biz-asset-GetMerchantCoinMultiReq)
    - [GetMerchantCoinMultiResp](#svc-biz-asset-GetMerchantCoinMultiResp)
    - [GetMerchantCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetMerchantCoinMultiResp-ValueMapEntry)
    - [GetMerchantCoinReq](#svc-biz-asset-GetMerchantCoinReq)
    - [GetMerchantCoinResp](#svc-biz-asset-GetMerchantCoinResp)
    - [GetStreamerCoinMultiReq](#svc-biz-asset-GetStreamerCoinMultiReq)
    - [GetStreamerCoinMultiResp](#svc-biz-asset-GetStreamerCoinMultiResp)
    - [GetStreamerCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetStreamerCoinMultiResp-ValueMapEntry)
    - [GetStreamerCoinReq](#svc-biz-asset-GetStreamerCoinReq)
    - [GetStreamerCoinResp](#svc-biz-asset-GetStreamerCoinResp)
    - [GetStreamerMoneyMultiReq](#svc-biz-asset-GetStreamerMoneyMultiReq)
    - [GetStreamerMoneyMultiResp](#svc-biz-asset-GetStreamerMoneyMultiResp)
    - [GetStreamerMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetStreamerMoneyMultiResp-ValueMapEntry)
    - [GetStreamerMoneyReq](#svc-biz-asset-GetStreamerMoneyReq)
    - [GetStreamerMoneyResp](#svc-biz-asset-GetStreamerMoneyResp)
    - [GetUnionCoinMultiReq](#svc-biz-asset-GetUnionCoinMultiReq)
    - [GetUnionCoinMultiResp](#svc-biz-asset-GetUnionCoinMultiResp)
    - [GetUnionCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetUnionCoinMultiResp-ValueMapEntry)
    - [GetUnionCoinReq](#svc-biz-asset-GetUnionCoinReq)
    - [GetUnionCoinResp](#svc-biz-asset-GetUnionCoinResp)
    - [GetUnionMoneyMultiReq](#svc-biz-asset-GetUnionMoneyMultiReq)
    - [GetUnionMoneyMultiResp](#svc-biz-asset-GetUnionMoneyMultiResp)
    - [GetUnionMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetUnionMoneyMultiResp-ValueMapEntry)
    - [GetUnionMoneyReq](#svc-biz-asset-GetUnionMoneyReq)
    - [GetUnionMoneyResp](#svc-biz-asset-GetUnionMoneyResp)
    - [GetUserCoinMultiReq](#svc-biz-asset-GetUserCoinMultiReq)
    - [GetUserCoinMultiResp](#svc-biz-asset-GetUserCoinMultiResp)
    - [GetUserCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetUserCoinMultiResp-ValueMapEntry)
    - [GetUserCoinReq](#svc-biz-asset-GetUserCoinReq)
    - [GetUserCoinResp](#svc-biz-asset-GetUserCoinResp)
    - [GetUserMoneyMultiReq](#svc-biz-asset-GetUserMoneyMultiReq)
    - [GetUserMoneyMultiResp](#svc-biz-asset-GetUserMoneyMultiResp)
    - [GetUserMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetUserMoneyMultiResp-ValueMapEntry)
    - [GetUserMoneyReq](#svc-biz-asset-GetUserMoneyReq)
    - [GetUserMoneyResp](#svc-biz-asset-GetUserMoneyResp)
    - [IncrMerchantCoinReq](#svc-biz-asset-IncrMerchantCoinReq)
    - [IncrStreamerCoinReq](#svc-biz-asset-IncrStreamerCoinReq)
    - [IncrStreamerMoneyReq](#svc-biz-asset-IncrStreamerMoneyReq)
    - [IncrUnionCoinReq](#svc-biz-asset-IncrUnionCoinReq)
    - [IncrUnionMoneyReq](#svc-biz-asset-IncrUnionMoneyReq)
    - [IncrUserCoinReq](#svc-biz-asset-IncrUserCoinReq)
    - [IncrUserMoneyReq](#svc-biz-asset-IncrUserMoneyReq)
    - [ListMerchantCoinDetailReq](#svc-biz-asset-ListMerchantCoinDetailReq)
    - [ListMerchantCoinDetailResp](#svc-biz-asset-ListMerchantCoinDetailResp)
    - [ListStreamerCoinDetailReq](#svc-biz-asset-ListStreamerCoinDetailReq)
    - [ListStreamerCoinDetailResp](#svc-biz-asset-ListStreamerCoinDetailResp)
    - [ListStreamerMoneyDetailReq](#svc-biz-asset-ListStreamerMoneyDetailReq)
    - [ListStreamerMoneyDetailResp](#svc-biz-asset-ListStreamerMoneyDetailResp)
    - [ListUnionCoinDetailReq](#svc-biz-asset-ListUnionCoinDetailReq)
    - [ListUnionCoinDetailResp](#svc-biz-asset-ListUnionCoinDetailResp)
    - [ListUnionMoneyDetailReq](#svc-biz-asset-ListUnionMoneyDetailReq)
    - [ListUnionMoneyDetailResp](#svc-biz-asset-ListUnionMoneyDetailResp)
    - [ListUserCoinDetailReq](#svc-biz-asset-ListUserCoinDetailReq)
    - [ListUserCoinDetailResp](#svc-biz-asset-ListUserCoinDetailResp)
    - [ListUserMoneyDetailReq](#svc-biz-asset-ListUserMoneyDetailReq)
    - [ListUserMoneyDetailResp](#svc-biz-asset-ListUserMoneyDetailResp)
    - [MerchantCoinDetail](#svc-biz-asset-MerchantCoinDetail)
    - [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue)
    - [StreamerCoinDetail](#svc-biz-asset-StreamerCoinDetail)
    - [StreamerCoinValue](#svc-biz-asset-StreamerCoinValue)
    - [StreamerMoneyDetail](#svc-biz-asset-StreamerMoneyDetail)
    - [UnionCoinDetail](#svc-biz-asset-UnionCoinDetail)
    - [UnionCoinValue](#svc-biz-asset-UnionCoinValue)
    - [UnionMoneyDetail](#svc-biz-asset-UnionMoneyDetail)
    - [UserCoinDetail](#svc-biz-asset-UserCoinDetail)
    - [UserCoinValue](#svc-biz-asset-UserCoinValue)
    - [UserMoneyDetail](#svc-biz-asset-UserMoneyDetail)
    - [ValueChange](#svc-biz-asset-ValueChange)
  
    - [Asset](#svc-biz-asset-Asset)
  
- [third_party/buf/validate/priv/private.proto](#third_party_buf_validate_priv_private-proto)
    - [Constraint](#buf-validate-priv-Constraint)
    - [FieldConstraints](#buf-validate-priv-FieldConstraints)
  
    - [File-level Extensions](#third_party_buf_validate_priv_private-proto-extensions)
  
- [third_party/buf/validate/expression.proto](#third_party_buf_validate_expression-proto)
    - [Constraint](#buf-validate-Constraint)
    - [Violation](#buf-validate-Violation)
    - [Violations](#buf-validate-Violations)
  
- [third_party/buf/validate/validate.proto](#third_party_buf_validate_validate-proto)
    - [AnyRules](#buf-validate-AnyRules)
    - [BoolRules](#buf-validate-BoolRules)
    - [BytesRules](#buf-validate-BytesRules)
    - [DoubleRules](#buf-validate-DoubleRules)
    - [DurationRules](#buf-validate-DurationRules)
    - [EnumRules](#buf-validate-EnumRules)
    - [FieldConstraints](#buf-validate-FieldConstraints)
    - [Fixed32Rules](#buf-validate-Fixed32Rules)
    - [Fixed64Rules](#buf-validate-Fixed64Rules)
    - [FloatRules](#buf-validate-FloatRules)
    - [Int32Rules](#buf-validate-Int32Rules)
    - [Int64Rules](#buf-validate-Int64Rules)
    - [MapRules](#buf-validate-MapRules)
    - [MessageConstraints](#buf-validate-MessageConstraints)
    - [OneofConstraints](#buf-validate-OneofConstraints)
    - [RepeatedRules](#buf-validate-RepeatedRules)
    - [SFixed32Rules](#buf-validate-SFixed32Rules)
    - [SFixed64Rules](#buf-validate-SFixed64Rules)
    - [SInt32Rules](#buf-validate-SInt32Rules)
    - [SInt64Rules](#buf-validate-SInt64Rules)
    - [StringRules](#buf-validate-StringRules)
    - [TimestampRules](#buf-validate-TimestampRules)
    - [UInt32Rules](#buf-validate-UInt32Rules)
    - [UInt64Rules](#buf-validate-UInt64Rules)
  
    - [Ignore](#buf-validate-Ignore)
    - [KnownRegex](#buf-validate-KnownRegex)
  
    - [File-level Extensions](#third_party_buf_validate_validate-proto-extensions)
    - [File-level Extensions](#third_party_buf_validate_validate-proto-extensions)
    - [File-level Extensions](#third_party_buf_validate_validate-proto-extensions)
  
- [svc.web.dashboard/dashboard.proto](#svc-web-dashboard_dashboard-proto)
    - [Dashboard](#svc-web-dashboard-Dashboard)
  
- [svc.biz.org/org.proto](#svc-biz-org_org-proto)
    - [InitDBResp](#svc-biz-org-InitDBResp)
  
    - [Org](#svc-biz-org-Org)
  
- [svc.infra.notifier/notifier.proto](#svc-infra-notifier_notifier-proto)
    - [InitDBResp](#svc-infra-notifier-InitDBResp)
  
    - [Notifier](#svc-infra-notifier-Notifier)
  
- [svc.web.streamer/streamer.proto](#svc-web-streamer_streamer-proto)
    - [Streamer](#svc-web-streamer-Streamer)
  
- [svc.biz.gift/gift.proto](#svc-biz-gift_gift-proto)
    - [GiftAddReq](#svc-biz-gift-GiftAddReq)
    - [GiftAddResp](#svc-biz-gift-GiftAddResp)
    - [GiftGetRecordReq](#svc-biz-gift-GiftGetRecordReq)
    - [GiftGetRecordResp](#svc-biz-gift-GiftGetRecordResp)
    - [GiftGetReq](#svc-biz-gift-GiftGetReq)
    - [GiftGetResp](#svc-biz-gift-GiftGetResp)
    - [GiftInfo](#svc-biz-gift-GiftInfo)
    - [GiftSendRecord](#svc-biz-gift-GiftSendRecord)
    - [GiftSendRecordReq](#svc-biz-gift-GiftSendRecordReq)
    - [GiftSendRecordResp](#svc-biz-gift-GiftSendRecordResp)
    - [GiftSendReq](#svc-biz-gift-GiftSendReq)
    - [GiftSendResp](#svc-biz-gift-GiftSendResp)
    - [GiftUpdateReq](#svc-biz-gift-GiftUpdateReq)
    - [GiftUpdateResp](#svc-biz-gift-GiftUpdateResp)
    - [ListAdminReq](#svc-biz-gift-ListAdminReq)
    - [ListAdminResp](#svc-biz-gift-ListAdminResp)
    - [ListOnlineAllReq](#svc-biz-gift-ListOnlineAllReq)
    - [ListOnlineByTypeReq](#svc-biz-gift-ListOnlineByTypeReq)
    - [ListOnlineResp](#svc-biz-gift-ListOnlineResp)
    - [LiveStatReq](#svc-biz-gift-LiveStatReq)
    - [LiveStatResp](#svc-biz-gift-LiveStatResp)
  
    - [GiftRecommend](#svc-biz-gift-GiftRecommend)
    - [GiftStatus](#svc-biz-gift-GiftStatus)
    - [GiftType](#svc-biz-gift-GiftType)
  
    - [Gift](#svc-biz-gift-Gift)
  
- [Scalar Value Types](#scalar-value-types)



<a name="svc-biz-account_account-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.account/account.proto



<a name="svc-biz-account-AddManagerReq"></a>

### AddManagerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| manager | [Manager](#svc-biz-account-Manager) |  |  |






<a name="svc-biz-account-AddManagerResp"></a>

### AddManagerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| manager | [Manager](#svc-biz-account-Manager) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-biz-account-AddStreamerReq"></a>

### AddStreamerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer | [Streamer](#svc-biz-account-Streamer) |  |  |






<a name="svc-biz-account-AddStreamerResp"></a>

### AddStreamerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer | [Streamer](#svc-biz-account-Streamer) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-biz-account-AddUnionReq"></a>

### AddUnionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union | [Union](#svc-biz-account-Union) |  |  |






<a name="svc-biz-account-AddUnionResp"></a>

### AddUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union | [Union](#svc-biz-account-Union) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-biz-account-AddViewerReq"></a>

### AddViewerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| viewer | [Viewer](#svc-biz-account-Viewer) |  |  |






<a name="svc-biz-account-AddViewerResp"></a>

### AddViewerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| viewer | [Viewer](#svc-biz-account-Viewer) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-biz-account-DeleteManagerReq"></a>

### DeleteManagerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Manager](#svc-biz-account-Manager) |  |  |






<a name="svc-biz-account-DeleteManagerResp"></a>

### DeleteManagerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-biz-account-DeleteStreamerReq"></a>

### DeleteStreamerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Streamer](#svc-biz-account-Streamer) |  |  |






<a name="svc-biz-account-DeleteStreamerResp"></a>

### DeleteStreamerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-biz-account-DeleteUnionReq"></a>

### DeleteUnionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-account-Union) |  |  |






<a name="svc-biz-account-DeleteUnionResp"></a>

### DeleteUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-biz-account-DeleteViewerReq"></a>

### DeleteViewerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Viewer](#svc-biz-account-Viewer) |  |  |






<a name="svc-biz-account-DeleteViewerResp"></a>

### DeleteViewerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-biz-account-GetManagerReq"></a>

### GetManagerReq
{{{ [Manager]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Manager](#svc-biz-account-Manager) |  |  |






<a name="svc-biz-account-GetManagerResp"></a>

### GetManagerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| manager | [Manager](#svc-biz-account-Manager) |  |  |






<a name="svc-biz-account-GetStreamerReq"></a>

### GetStreamerReq
{{{ [Streamer]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Streamer](#svc-biz-account-Streamer) |  |  |






<a name="svc-biz-account-GetStreamerResp"></a>

### GetStreamerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer | [Streamer](#svc-biz-account-Streamer) |  |  |






<a name="svc-biz-account-GetUnionReq"></a>

### GetUnionReq
{{{ [Union]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-account-Union) |  |  |






<a name="svc-biz-account-GetUnionResp"></a>

### GetUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union | [Union](#svc-biz-account-Union) |  |  |






<a name="svc-biz-account-GetViewerReq"></a>

### GetViewerReq
{{{ [Viewer]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Viewer](#svc-biz-account-Viewer) |  |  |






<a name="svc-biz-account-GetViewerResp"></a>

### GetViewerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| viewer | [Viewer](#svc-biz-account-Viewer) |  |  |






<a name="svc-biz-account-InitDBResp"></a>

### InitDBResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |






<a name="svc-biz-account-ListManagersReq"></a>

### ListManagersReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Manager](#svc-biz-account-Manager) |  |  |






<a name="svc-biz-account-ListManagersResp"></a>

### ListManagersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| managers | [Manager](#svc-biz-account-Manager) | repeated |  |






<a name="svc-biz-account-ListStreamersReq"></a>

### ListStreamersReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Streamer](#svc-biz-account-Streamer) |  |  |






<a name="svc-biz-account-ListStreamersResp"></a>

### ListStreamersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamers | [Streamer](#svc-biz-account-Streamer) | repeated |  |






<a name="svc-biz-account-ListUnionsReq"></a>

### ListUnionsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-account-Union) |  |  |






<a name="svc-biz-account-ListUnionsResp"></a>

### ListUnionsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unions | [Union](#svc-biz-account-Union) | repeated |  |






<a name="svc-biz-account-ListViewersReq"></a>

### ListViewersReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Viewer](#svc-biz-account-Viewer) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="svc-biz-account-ListViewersResp"></a>

### ListViewersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| viewers | [Viewer](#svc-biz-account-Viewer) | repeated |  |






<a name="svc-biz-account-Manager"></a>

### Manager



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| name | [string](#string) |  | 名字 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| password | [string](#string) | optional | 密码 |
| salt | [string](#string) | optional | 加密混淆 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |






<a name="svc-biz-account-Streamer"></a>

### Streamer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| display_id | [string](#string) |  | 用于显示和索引的数字ID |
| name | [string](#string) |  | 用户名 |
| nickname | [string](#string) |  | 昵称 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| password | [string](#string) | optional | 密码 |
| salt | [string](#string) | optional | 加密混淆 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |






<a name="svc-biz-account-Union"></a>

### Union



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| name | [string](#string) |  | 名字 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| password | [string](#string) | optional | 密码 |
| salt | [string](#string) | optional | 加密混淆 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |






<a name="svc-biz-account-UpdateManagerReq"></a>

### UpdateManagerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Manager](#svc-biz-account-Manager) |  |  |






<a name="svc-biz-account-UpdateManagerResp"></a>

### UpdateManagerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |






<a name="svc-biz-account-UpdateStreamerReq"></a>

### UpdateStreamerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Streamer](#svc-biz-account-Streamer) |  |  |






<a name="svc-biz-account-UpdateStreamerResp"></a>

### UpdateStreamerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |






<a name="svc-biz-account-UpdateUnionReq"></a>

### UpdateUnionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-account-Union) |  |  |






<a name="svc-biz-account-UpdateUnionResp"></a>

### UpdateUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |






<a name="svc-biz-account-UpdateViewerReq"></a>

### UpdateViewerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Viewer](#svc-biz-account-Viewer) |  |  |






<a name="svc-biz-account-UpdateViewerResp"></a>

### UpdateViewerResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |






<a name="svc-biz-account-Viewer"></a>

### Viewer
Models


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| display_id | [string](#string) |  | 用于显示和索引的数字ID |
| name | [string](#string) |  | 用户名 |
| nickname | [string](#string) |  | 昵称 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| device_ident | [string](#string) |  | 设备号 / 指纹 |
| password | [string](#string) | optional | 密码 |
| salt | [string](#string) | optional | 加密混淆 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |





 

 

 


<a name="svc-biz-account-Account"></a>

### Account


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InitDB | [.google.protobuf.Empty](#google-protobuf-Empty) | [InitDBResp](#svc-biz-account-InitDBResp) | 初始化数据库 |
| GetViewer | [GetViewerReq](#svc-biz-account-GetViewerReq) | [GetViewerResp](#svc-biz-account-GetViewerResp) | 获取普通账号 |
| ListViewers | [ListViewersReq](#svc-biz-account-ListViewersReq) | [ListViewersResp](#svc-biz-account-ListViewersResp) | 获取普通账号列表 |
| AddViewer | [AddViewerReq](#svc-biz-account-AddViewerReq) | [AddViewerResp](#svc-biz-account-AddViewerResp) | 添加普通账号 |
| UpdateViewer | [UpdateViewerReq](#svc-biz-account-UpdateViewerReq) | [UpdateViewerResp](#svc-biz-account-UpdateViewerResp) | 更新普通账号 |
| DeleteViewer | [DeleteViewerReq](#svc-biz-account-DeleteViewerReq) | [DeleteViewerResp](#svc-biz-account-DeleteViewerResp) | 删除普通账号 |
| GetStreamer | [GetStreamerReq](#svc-biz-account-GetStreamerReq) | [GetStreamerResp](#svc-biz-account-GetStreamerResp) | 获取主播账号 |
| ListStreamers | [ListStreamersReq](#svc-biz-account-ListStreamersReq) | [ListStreamersResp](#svc-biz-account-ListStreamersResp) | 获取主播账号列表 |
| AddStreamer | [AddStreamerReq](#svc-biz-account-AddStreamerReq) | [AddStreamerResp](#svc-biz-account-AddStreamerResp) | 添加普通账号 |
| UpdateStreamer | [UpdateStreamerReq](#svc-biz-account-UpdateStreamerReq) | [UpdateStreamerResp](#svc-biz-account-UpdateStreamerResp) | 更新主播账号 |
| DeleteStreamer | [DeleteStreamerReq](#svc-biz-account-DeleteStreamerReq) | [DeleteStreamerResp](#svc-biz-account-DeleteStreamerResp) | 删除主播账号 |
| GetManager | [GetManagerReq](#svc-biz-account-GetManagerReq) | [GetManagerResp](#svc-biz-account-GetManagerResp) | 获取管理账号 |
| ListManagers | [ListManagersReq](#svc-biz-account-ListManagersReq) | [ListManagersResp](#svc-biz-account-ListManagersResp) | 获取管理账号列表 |
| AddManager | [AddManagerReq](#svc-biz-account-AddManagerReq) | [AddManagerResp](#svc-biz-account-AddManagerResp) | 添加管理账号 |
| UpdateManager | [UpdateManagerReq](#svc-biz-account-UpdateManagerReq) | [UpdateManagerResp](#svc-biz-account-UpdateManagerResp) | 更新管理账号 |
| DeleteManager | [DeleteManagerReq](#svc-biz-account-DeleteManagerReq) | [DeleteManagerResp](#svc-biz-account-DeleteManagerResp) | 删除管理账号 |
| GetUnion | [GetUnionReq](#svc-biz-account-GetUnionReq) | [GetUnionResp](#svc-biz-account-GetUnionResp) | 获取工会账号 |
| ListUnions | [ListUnionsReq](#svc-biz-account-ListUnionsReq) | [ListUnionsResp](#svc-biz-account-ListUnionsResp) | 获取工会账号列表 |
| AddUnion | [AddUnionReq](#svc-biz-account-AddUnionReq) | [AddUnionResp](#svc-biz-account-AddUnionResp) | 添加工会账号 |
| UpdateUnion | [UpdateUnionReq](#svc-biz-account-UpdateUnionReq) | [UpdateUnionResp](#svc-biz-account-UpdateUnionResp) | 更新工会账号 |
| DeleteUnion | [DeleteUnionReq](#svc-biz-account-DeleteUnionReq) | [DeleteUnionResp](#svc-biz-account-DeleteUnionResp) | 删除工会账号 |

 



<a name="svc-biz-trade_trade-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.trade/trade.proto



<a name="svc-biz-trade-BuyLiveTicketReq"></a>

### BuyLiveTicketReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| live_info | [LiveInfo](#svc-biz-trade-LiveInfo) |  |  |
| ticket | [LiveTicketInfo](#svc-biz-trade-LiveTicketInfo) |  |  |






<a name="svc-biz-trade-BuyLiveTicketResp"></a>

### BuyLiveTicketResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-BuyLuckyIdInfo"></a>

### BuyLuckyIdInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-BuyLuckyIdReq"></a>

### BuyLuckyIdReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| lucky_id | [BuyLuckyIdInfo](#svc-biz-trade-BuyLuckyIdInfo) |  |  |






<a name="svc-biz-trade-BuyLuckyIdResp"></a>

### BuyLuckyIdResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-BuyRideInfo"></a>

### BuyRideInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-BuyRideReq"></a>

### BuyRideReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| ride | [BuyRideInfo](#svc-biz-trade-BuyRideInfo) |  |  |






<a name="svc-biz-trade-BuyRideResp"></a>

### BuyRideResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-GiftInfo"></a>

### GiftInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [string](#string) |  |  |
| gift_name | [string](#string) |  |  |
| gift_type | [int64](#int64) |  |  |
| price | [int64](#int64) |  |  |
| num | [int64](#int64) |  |  |
| icon | [string](#string) |  |  |






<a name="svc-biz-trade-JoinAnchorFansGroupInfo"></a>

### JoinAnchorFansGroupInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-JoinAnchorFansGroupReq"></a>

### JoinAnchorFansGroupReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| live_info | [LiveInfo](#svc-biz-trade-LiveInfo) |  |  |
| fans_group | [JoinAnchorFansGroupInfo](#svc-biz-trade-JoinAnchorFansGroupInfo) |  |  |






<a name="svc-biz-trade-JoinAnchorFansGroupResp"></a>

### JoinAnchorFansGroupResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-LiveDurationFeeInfo"></a>

### LiveDurationFeeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-LiveInfo"></a>

### LiveInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| room_id | [string](#string) |  |  |
| anchor_id | [string](#string) |  |  |
| live_id | [string](#string) |  |  |
| live_type | [int64](#int64) |  | 直播类型 |
| live_category | [int64](#int64) |  | 直播分类 |






<a name="svc-biz-trade-LiveTicketInfo"></a>

### LiveTicketInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-MoneyExchangeCoinInfo"></a>

### MoneyExchangeCoinInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| money_amount | [int64](#int64) |  |  |
| coin_amount | [int64](#int64) |  |  |
| exchange_rate | [int64](#int64) |  |  |






<a name="svc-biz-trade-MoneyExchangeCoinReq"></a>

### MoneyExchangeCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| user | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| exchange_info | [MoneyExchangeCoinInfo](#svc-biz-trade-MoneyExchangeCoinInfo) |  |  |






<a name="svc-biz-trade-MoneyExchangeCoinResp"></a>

### MoneyExchangeCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| money_result | [MoneyResult](#svc-biz-trade-MoneyResult) |  |  |






<a name="svc-biz-trade-MoneyRechargeReq"></a>

### MoneyRechargeReq
待讨论


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| user | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| pay_info | [MoneyRechargeReq.PayInfo](#svc-biz-trade-MoneyRechargeReq-PayInfo) |  |  |
| recharge_amount | [int64](#int64) |  | 充值余额金额 |
| attach | [string](#string) |  | 透传的附加信息,等待约定 |






<a name="svc-biz-trade-MoneyRechargeReq-PayInfo"></a>

### MoneyRechargeReq.PayInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pay_channel | [string](#string) |  |  |
| pay_method | [string](#string) |  |  |
| currency_code | [string](#string) |  | 币种代码 CNY |
| amount | [int64](#int64) |  | 币种金额 |






<a name="svc-biz-trade-MoneyRechargeResp"></a>

### MoneyRechargeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| money_result | [MoneyResult](#svc-biz-trade-MoneyResult) |  |  |






<a name="svc-biz-trade-MoneyResult"></a>

### MoneyResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |
| serial_number | [int64](#int64) |  |  |






<a name="svc-biz-trade-MoneyWithdrawReq"></a>

### MoneyWithdrawReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| user | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| withdraw_amount | [int64](#int64) |  | 提现金额 |






<a name="svc-biz-trade-MoneyWithdrawResp"></a>

### MoneyWithdrawResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| money_result | [MoneyResult](#svc-biz-trade-MoneyResult) |  |  |






<a name="svc-biz-trade-PayBulletChatInfo"></a>

### PayBulletChatInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-PayBulletChatReq"></a>

### PayBulletChatReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| live_info | [LiveInfo](#svc-biz-trade-LiveInfo) |  |  |
| bullet_chat | [PayBulletChatInfo](#svc-biz-trade-PayBulletChatInfo) |  |  |






<a name="svc-biz-trade-PayBulletChatResp"></a>

### PayBulletChatResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-PayLiveDurationFeeReq"></a>

### PayLiveDurationFeeReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| live_info | [LiveInfo](#svc-biz-trade-LiveInfo) |  |  |
| duration_fee | [LiveDurationFeeInfo](#svc-biz-trade-LiveDurationFeeInfo) |  |  |






<a name="svc-biz-trade-PayLiveDurationFeeResp"></a>

### PayLiveDurationFeeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-SendGiftInLiveReq"></a>

### SendGiftInLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| sender | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| receiver | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| gift | [GiftInfo](#svc-biz-trade-GiftInfo) |  |  |
| live_info | [LiveInfo](#svc-biz-trade-LiveInfo) |  |  |






<a name="svc-biz-trade-SendGiftInLiveResp"></a>

### SendGiftInLiveResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-TradeInfo"></a>

### TradeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  | 交易发生的商户 |
| app_id | [string](#string) |  | 交易发生的app |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-trade-TradeResult"></a>

### TradeResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [UserCoinValue](#svc-biz-trade-UserCoinValue) |  |  |
| serial_number | [int64](#int64) |  |  |






<a name="svc-biz-trade-UserCoinValue"></a>

### UserCoinValue
总余额=value&#43;limited_value


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-trade-UserInfo"></a>

### UserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| uid | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| avatar | [string](#string) |  |  |
| union_id | [string](#string) |  |  |






<a name="svc-biz-trade-VipActivateInfo"></a>

### VipActivateInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-VipActivateReq"></a>

### VipActivateReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| vip | [VipActivateInfo](#svc-biz-trade-VipActivateInfo) |  |  |






<a name="svc-biz-trade-VipActivateResp"></a>

### VipActivateResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-VipExtendInfo"></a>

### VipExtendInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-VipExtendReq"></a>

### VipExtendReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| vip | [VipExtendInfo](#svc-biz-trade-VipExtendInfo) |  |  |






<a name="svc-biz-trade-VipExtendResp"></a>

### VipExtendResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-VipUpgradeInfo"></a>

### VipUpgradeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [int64](#int64) |  |  |






<a name="svc-biz-trade-VipUpgradeReq"></a>

### VipUpgradeReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| vip | [VipUpgradeInfo](#svc-biz-trade-VipUpgradeInfo) |  |  |






<a name="svc-biz-trade-VipUpgradeResp"></a>

### VipUpgradeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |





 

 

 


<a name="svc-biz-trade-Trade"></a>

### Trade


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendGiftInLive | [SendGiftInLiveReq](#svc-biz-trade-SendGiftInLiveReq) | [SendGiftInLiveResp](#svc-biz-trade-SendGiftInLiveResp) |  |
| BuyLiveTicket | [BuyLiveTicketReq](#svc-biz-trade-BuyLiveTicketReq) | [BuyLiveTicketResp](#svc-biz-trade-BuyLiveTicketResp) |  |
| PayLiveDurationFee | [PayLiveDurationFeeReq](#svc-biz-trade-PayLiveDurationFeeReq) | [PayLiveDurationFeeResp](#svc-biz-trade-PayLiveDurationFeeResp) |  |
| JoinAnchorFansGroup | [JoinAnchorFansGroupReq](#svc-biz-trade-JoinAnchorFansGroupReq) | [JoinAnchorFansGroupResp](#svc-biz-trade-JoinAnchorFansGroupResp) |  |
| PayBulletChat | [PayBulletChatReq](#svc-biz-trade-PayBulletChatReq) | [PayBulletChatResp](#svc-biz-trade-PayBulletChatResp) |  |
| VipActivate | [VipActivateReq](#svc-biz-trade-VipActivateReq) | [VipActivateResp](#svc-biz-trade-VipActivateResp) |  |
| VipExtend | [VipExtendReq](#svc-biz-trade-VipExtendReq) | [VipExtendResp](#svc-biz-trade-VipExtendResp) |  |
| VipUpgrade | [VipUpgradeReq](#svc-biz-trade-VipUpgradeReq) | [VipUpgradeResp](#svc-biz-trade-VipUpgradeResp) |  |
| BuyRide | [BuyRideReq](#svc-biz-trade-BuyRideReq) | [BuyRideResp](#svc-biz-trade-BuyRideResp) |  |
| BuyLuckyId | [BuyLuckyIdReq](#svc-biz-trade-BuyLuckyIdReq) | [BuyLuckyIdResp](#svc-biz-trade-BuyLuckyIdResp) |  |
| MoneyRecharge | [MoneyRechargeReq](#svc-biz-trade-MoneyRechargeReq) | [MoneyRechargeResp](#svc-biz-trade-MoneyRechargeResp) |  |
| MoneyWithdraw | [MoneyWithdrawReq](#svc-biz-trade-MoneyWithdrawReq) | [MoneyWithdrawResp](#svc-biz-trade-MoneyWithdrawResp) |  |
| MoneyExchangeCoin | [MoneyExchangeCoinReq](#svc-biz-trade-MoneyExchangeCoinReq) | [MoneyExchangeCoinResp](#svc-biz-trade-MoneyExchangeCoinResp) |  |

 



<a name="svc-infra-static_static-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.infra.static/static.proto



<a name="svc-infra-static-InitDBResp"></a>

### InitDBResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |





 

 

 


<a name="svc-infra-static-Static"></a>

### Static


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InitDB | [.google.protobuf.Empty](#google-protobuf-Empty) | [InitDBResp](#svc-infra-static-InitDBResp) | 初始化数据库 |

 



<a name="svc-biz-room_category-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.room/category.proto



<a name="svc-biz-room-CategoryInfo"></a>

### CategoryInfo
CategoryInfo 分类详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [string](#string) |  |  |
| parent_category_id | [string](#string) |  | 父级ID |
| category_code | [string](#string) |  | 代号（唯一，预留） |
| category_name | [string](#string) |  | 名称 |
| sort | [int32](#int32) |  | 排序 |
| childrens | [CategoryInfo](#svc-biz-room-CategoryInfo) | repeated |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="svc-biz-room-CreateCategoryReq"></a>

### CreateCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [CategoryInfo](#svc-biz-room-CategoryInfo) |  |  |






<a name="svc-biz-room-CreateCategoryResp"></a>

### CreateCategoryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [CategoryInfo](#svc-biz-room-CategoryInfo) |  |  |






<a name="svc-biz-room-DeleteCategoryReq"></a>

### DeleteCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [string](#string) |  |  |






<a name="svc-biz-room-GetCategoryReq"></a>

### GetCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [string](#string) |  |  |






<a name="svc-biz-room-GetCategoryResp"></a>

### GetCategoryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [CategoryInfo](#svc-biz-room-CategoryInfo) |  |  |






<a name="svc-biz-room-ListCategoryReq"></a>

### ListCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| return_count | [bool](#bool) |  | 是否返回总数 |
| level | [int32](#int32) |  | 查询标识（0查询所有，1=查询一级分类；2=查询二级分类） |
| parent_category_id | [string](#string) |  | 父级ID |
| category_name | [string](#string) |  | 分类名 |






<a name="svc-biz-room-ListCategoryResp"></a>

### ListCategoryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [CategoryInfo](#svc-biz-room-CategoryInfo) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="svc-biz-room-ListCategoryTreeReq"></a>

### ListCategoryTreeReq
获取全部板块分类（分类及子分类树结构）






<a name="svc-biz-room-ListCategoryTreeResp"></a>

### ListCategoryTreeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [CategoryInfo](#svc-biz-room-CategoryInfo) | repeated |  |






<a name="svc-biz-room-UpdateCategoryReq"></a>

### UpdateCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [string](#string) |  |  |
| category | [CategoryInfo](#svc-biz-room-CategoryInfo) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |





 

 

 


<a name="svc-biz-room-Category"></a>

### Category
分类

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCategory | [GetCategoryReq](#svc-biz-room-GetCategoryReq) | [GetCategoryResp](#svc-biz-room-GetCategoryResp) | 获取分类 |
| CreateCategory | [CreateCategoryReq](#svc-biz-room-CreateCategoryReq) | [CreateCategoryResp](#svc-biz-room-CreateCategoryResp) | 创建分类 |
| UpdateCategory | [UpdateCategoryReq](#svc-biz-room-UpdateCategoryReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新某个分类信息 |
| DeleteCategory | [DeleteCategoryReq](#svc-biz-room-DeleteCategoryReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除一个分类信息 |
| ListCategory | [ListCategoryReq](#svc-biz-room-ListCategoryReq) | [ListCategoryResp](#svc-biz-room-ListCategoryResp) | 获取分类，返回子级集合 |
| ListCategoryTree | [ListCategoryTreeReq](#svc-biz-room-ListCategoryTreeReq) | [ListCategoryTreeResp](#svc-biz-room-ListCategoryTreeResp) | 获取全部板块分类（分类及子分类树结构） |

 



<a name="svc-biz-room_live-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.room/live.proto



<a name="svc-biz-room-GetLiveReq"></a>

### GetLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live_id | [string](#string) |  | 直播id |






<a name="svc-biz-room-GetLiveResp"></a>

### GetLiveResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live | [LiveInfo](#svc-biz-room-LiveInfo) |  | 直播信息 |






<a name="svc-biz-room-ListLiveReq"></a>

### ListLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  | 页数 |
| limit | [int32](#int32) |  | 条数 |
| streamer_id | [string](#string) |  | 主播id |






<a name="svc-biz-room-ListLiveResp"></a>

### ListLiveResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [LiveInfo](#svc-biz-room-LiveInfo) | repeated | 直播信息 |






<a name="svc-biz-room-LiveInfo"></a>

### LiveInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live_id | [string](#string) |  | id |
| streamer_id | [string](#string) |  | 主播id |
| room_id | [string](#string) |  | 房间id |
| category_id | [string](#string) |  | 分类id |
| start_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开播时间 |
| end_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="svc-biz-room-MGetLiveReq"></a>

### MGetLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live_ids | [string](#string) | repeated | 直播id列表 |






<a name="svc-biz-room-MGetLiveResp"></a>

### MGetLiveResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MGetLiveResp.ItemsEntry](#svc-biz-room-MGetLiveResp-ItemsEntry) | repeated | 直播信息 |






<a name="svc-biz-room-MGetLiveResp-ItemsEntry"></a>

### MGetLiveResp.ItemsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [LiveInfo](#svc-biz-room-LiveInfo) |  |  |





 

 

 


<a name="svc-biz-room-Live"></a>

### Live
直播

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetLive | [GetLiveReq](#svc-biz-room-GetLiveReq) | [GetLiveResp](#svc-biz-room-GetLiveResp) | 查询直播间信息 |
| MGetLive | [MGetLiveReq](#svc-biz-room-MGetLiveReq) | [MGetLiveResp](#svc-biz-room-MGetLiveResp) | 批量获取直播间信息 |
| ListLive | [ListLiveReq](#svc-biz-room-ListLiveReq) | [ListLiveResp](#svc-biz-room-ListLiveResp) | 获取在播直播间列表 |

 



<a name="svc-biz-room_room-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.room/room.proto



<a name="svc-biz-room-CreateRoomReq"></a>

### CreateRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |






<a name="svc-biz-room-CreateRoomResp"></a>

### CreateRoomResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |






<a name="svc-biz-room-GetOnlineRoomListReq"></a>

### GetOnlineRoomListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  | 页数 |
| limit | [int32](#int32) |  | 条数（建议固定20） |
| sort_type | [SortType](#svc-biz-room-SortType) |  | 排序类型 |
| category_id | [string](#string) |  | 分类id |
| bind_tags | [string](#string) | repeated | 标签 |
| merchants | [string](#string) | repeated | 商户id |






<a name="svc-biz-room-GetOnlineRoomListResp"></a>

### GetOnlineRoomListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [RoomInfo](#svc-biz-room-RoomInfo) | repeated | 房间列表 |






<a name="svc-biz-room-GetRoomByStreamerIDReq"></a>

### GetRoomByStreamerIDReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  | 主播id |
| with_stream | [bool](#bool) |  | 是否带流信息 |






<a name="svc-biz-room-GetRoomByStreamerIDResp"></a>

### GetRoomByStreamerIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |
| stream | [Stream](#svc-biz-room-Stream) |  | 流信息 |






<a name="svc-biz-room-GetRoomListReq"></a>

### GetRoomListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  | 页数 |
| limit | [int32](#int32) |  | 条数 |
| streamer_id | [string](#string) |  | 主播id |






<a name="svc-biz-room-GetRoomListResp"></a>

### GetRoomListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [RoomInfo](#svc-biz-room-RoomInfo) | repeated | 房间列表 |






<a name="svc-biz-room-GetRoomReq"></a>

### GetRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [string](#string) |  | id |
| with_stream | [bool](#bool) |  | 是否带流信息 |






<a name="svc-biz-room-GetRoomResp"></a>

### GetRoomResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |
| stream | [Stream](#svc-biz-room-Stream) |  | 流信息 |






<a name="svc-biz-room-MGetRoomsByStreamerIDsReq"></a>

### MGetRoomsByStreamerIDsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_ids | [string](#string) | repeated | 主播id |






<a name="svc-biz-room-MGetRoomsByStreamerIDsResp"></a>

### MGetRoomsByStreamerIDsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MGetRoomsByStreamerIDsResp.ItemsEntry](#svc-biz-room-MGetRoomsByStreamerIDsResp-ItemsEntry) | repeated | 房间信息 map streamer_id -&gt; room |






<a name="svc-biz-room-MGetRoomsByStreamerIDsResp-ItemsEntry"></a>

### MGetRoomsByStreamerIDsResp.ItemsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [RoomInfo](#svc-biz-room-RoomInfo) |  |  |






<a name="svc-biz-room-MGetRoomsReq"></a>

### MGetRoomsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_ids | [string](#string) | repeated | id |






<a name="svc-biz-room-MGetRoomsResp"></a>

### MGetRoomsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MGetRoomsResp.ItemsEntry](#svc-biz-room-MGetRoomsResp-ItemsEntry) | repeated | 房间信息 map room_id -&gt; room |






<a name="svc-biz-room-MGetRoomsResp-ItemsEntry"></a>

### MGetRoomsResp.ItemsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [RoomInfo](#svc-biz-room-RoomInfo) |  |  |






<a name="svc-biz-room-RoomInfo"></a>

### RoomInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [string](#string) |  | id |
| display_id | [string](#string) |  | 显示id |
| streamer_id | [string](#string) |  | 主播id |
| category_id | [string](#string) |  | 分类id |
| title | [string](#string) |  | 标题 |
| intro | [string](#string) |  | 简介 |
| ban_expire | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 封禁到期时间 |
| hidden | [bool](#bool) |  | 是否隐藏 |
| merchants | [string](#string) | repeated | 商户id |
| bind_tags | [string](#string) | repeated | 标签 |
| live_id | [string](#string) |  | 直播id（开播状态时才会有，关播时清空） |
| live_region | [int32](#int32) |  | 直播区域（开播状态时才会有，关播时清空） |
| live_status | [LiveStatus](#svc-biz-room-LiveStatus) |  | 房间状态：1关播，2开播 |
| live_display_type | [LiveDisplayType](#svc-biz-room-LiveDisplayType) |  | 横竖屏类型（开播状态时才会有，关播时清空） |
| live_start_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开播时间（开播状态时才会有，关播时清空） |
| score_recommend | [int32](#int32) |  | 推荐分数（后台配置） |
| score_search | [int32](#int32) |  | 搜索分数（后台配置） |
| score_subscribe | [int32](#int32) |  | 关注分数（后台配置） |
| score_hot | [int32](#int32) |  | 热度分数（后台配置） |
| score_glamour | [int32](#int32) |  | 魅力值 |
| score_online | [int32](#int32) |  | 在线人数 |
| score_selected_gift | [int32](#int32) |  | 精选礼物最后赠送时间戳 |
| score_gift | [int32](#int32) |  | 纯礼物收益 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="svc-biz-room-StartLiveReq"></a>

### StartLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  | 主播id |
| display_type | [LiveDisplayType](#svc-biz-room-LiveDisplayType) |  | 横竖屏类型 |






<a name="svc-biz-room-StartLiveResp"></a>

### StartLiveResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |
| stream | [Stream](#svc-biz-room-Stream) |  | 流信息 |






<a name="svc-biz-room-StopLiveReq"></a>

### StopLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  | 主播id |






<a name="svc-biz-room-StopLiveResp"></a>

### StopLiveResp







<a name="svc-biz-room-Stream"></a>

### Stream
流信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pull | [StreamPull](#svc-biz-room-StreamPull) |  | 推流信息 |
| push | [StreamPush](#svc-biz-room-StreamPush) |  | 推流信息 |






<a name="svc-biz-room-StreamPull"></a>

### StreamPull
拉流地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rtmp | [string](#string) |  | RTMP 拉流地址 |
| flv | [string](#string) |  | FLV 拉流地址 |
| m3u8 | [string](#string) |  | M3U8 拉流地址 |
| udp | [string](#string) |  | UDP 拉流地址 |






<a name="svc-biz-room-StreamPush"></a>

### StreamPush
推流地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rtmp | [string](#string) |  | RTMP 推流地址 |
| web_rtc | [string](#string) |  | WebRTC 推流地址 |
| srt | [string](#string) |  | SRT 推流地址 |
| rmtp_over_srt | [string](#string) |  | RTMP over SRT 推流地址 |






<a name="svc-biz-room-UpdateRoomReq"></a>

### UpdateRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [string](#string) |  | id |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |





 


<a name="svc-biz-room-LiveDisplayType"></a>

### LiveDisplayType


| Name | Number | Description |
| ---- | ------ | ----------- |
| LiveDisplayTypeUnknown | 0 | 未知 |
| LiveDisplayTypeHorizontal | 1 | 横屏 |
| LiveDisplayTypeVertical | 2 | 竖屏 |



<a name="svc-biz-room-LiveStatus"></a>

### LiveStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| LiveStatusUnknown | 0 | 未知 |
| LiveStatusOnline | 1 | 上线 |
| LiveStatusOffline | 2 | 下线 |



<a name="svc-biz-room-SortType"></a>

### SortType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SortTypeUnknown | 0 | 未知 |
| SortTypeHot | 1 | 热度值（首页除关注之外排序都是使用热度值[推荐tab&#43;分类tab]；后台热度值 &gt; 魅力值 &gt;在线观众 &gt; 开播时间） |
| SortTypeSearch | 2 | 搜索（搜索页面的推荐列表排序；后台搜索推荐重 &gt; 魅力值 &gt;在线观众 &gt; 开播时间） |
| SortTypeSubscribe | 3 | 关注（搜索页面的推荐列表排序；后台关注推荐重 &gt; 魅力值 &gt;在线观众 &gt; 开播时间） |
| SortTypeRecommend | 4 | 推荐（房间详情页的”更多直播“推荐列表排序；后台推荐重 &gt; 魅力值 &gt;在线观众 &gt; 开播时间） |
| SortTypeSelected | 5 | 精选（精选礼物最后赠送时间 &gt; 礼物纯收益） |


 

 


<a name="svc-biz-room-Room"></a>

### Room
Room 房间

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateRoom | [CreateRoomReq](#svc-biz-room-CreateRoomReq) | [CreateRoomResp](#svc-biz-room-CreateRoomResp) | CreateRoom 创建房间 |
| UpdateRoom | [UpdateRoomReq](#svc-biz-room-UpdateRoomReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | UpdateRoom 更新房间 |
| GetRoom | [GetRoomReq](#svc-biz-room-GetRoomReq) | [GetRoomResp](#svc-biz-room-GetRoomResp) | GetRoom 查询房间 |
| GetRoomByStreamerID | [GetRoomByStreamerIDReq](#svc-biz-room-GetRoomByStreamerIDReq) | [GetRoomByStreamerIDResp](#svc-biz-room-GetRoomByStreamerIDResp) | GetRoomByStreamerID 查询房间 |
| MGetRooms | [MGetRoomsReq](#svc-biz-room-MGetRoomsReq) | [MGetRoomsResp](#svc-biz-room-MGetRoomsResp) | MGetRooms 查询房间 |
| MGetRoomsByStreamerIDs | [MGetRoomsByStreamerIDsReq](#svc-biz-room-MGetRoomsByStreamerIDsReq) | [MGetRoomsByStreamerIDsResp](#svc-biz-room-MGetRoomsByStreamerIDsResp) | MGetRoomByStreamerIDs 批量查询房间 |
| GetRoomList | [GetRoomListReq](#svc-biz-room-GetRoomListReq) | [GetRoomListResp](#svc-biz-room-GetRoomListResp) | GetRoomList 查询房间列表（后台使用此接口） |
| GetOnlineRoomList | [GetOnlineRoomListReq](#svc-biz-room-GetOnlineRoomListReq) | [GetOnlineRoomListResp](#svc-biz-room-GetOnlineRoomListResp) | GetOnlineRoomList 查询在线房间列表（用户端列表使用此接口） |
| StartLive | [StartLiveReq](#svc-biz-room-StartLiveReq) | [StartLiveResp](#svc-biz-room-StartLiveResp) | StartLive 开始直播 |
| StopLive | [StopLiveReq](#svc-biz-room-StopLiveReq) | [StopLiveResp](#svc-biz-room-StopLiveResp) | StopLive 关闭直播 |

 



<a name="svc-infra-setting_setting-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.infra.setting/setting.proto



<a name="svc-infra-setting-AddConfigurationReq"></a>

### AddConfigurationReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [Configuration](#svc-infra-setting-Configuration) |  |  |






<a name="svc-infra-setting-AddConfigurationResp"></a>

### AddConfigurationResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [Configuration](#svc-infra-setting-Configuration) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-infra-setting-Configuration"></a>

### Configuration



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| type | [string](#string) |  | 配置类型 |
| data | [string](#string) |  | 配置数据 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |






<a name="svc-infra-setting-DeleteConfigurationReq"></a>

### DeleteConfigurationReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Configuration](#svc-infra-setting-Configuration) |  |  |






<a name="svc-infra-setting-DeleteConfigurationResp"></a>

### DeleteConfigurationResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-infra-setting-GetConfigurationReq"></a>

### GetConfigurationReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Configuration](#svc-infra-setting-Configuration) |  |  |






<a name="svc-infra-setting-GetConfigurationResp"></a>

### GetConfigurationResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [Configuration](#svc-infra-setting-Configuration) |  |  |






<a name="svc-infra-setting-InitDBResp"></a>

### InitDBResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |






<a name="svc-infra-setting-ListConfigurationsReq"></a>

### ListConfigurationsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Configuration](#svc-infra-setting-Configuration) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="svc-infra-setting-ListConfigurationsResp"></a>

### ListConfigurationsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configurations | [Configuration](#svc-infra-setting-Configuration) | repeated |  |






<a name="svc-infra-setting-SettingGreetingReq"></a>

### SettingGreetingReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="svc-infra-setting-SettingGreetingResp"></a>

### SettingGreetingResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="svc-infra-setting-UpdateConfigurationReq"></a>

### UpdateConfigurationReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Configuration](#svc-infra-setting-Configuration) |  |  |






<a name="svc-infra-setting-UpdateConfigurationResp"></a>

### UpdateConfigurationResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |





 

 

 


<a name="svc-infra-setting-Setting"></a>

### Setting


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InitDB | [.google.protobuf.Empty](#google-protobuf-Empty) | [InitDBResp](#svc-infra-setting-InitDBResp) | 初始化数据库 |
| GetConfiguration | [GetConfigurationReq](#svc-infra-setting-GetConfigurationReq) | [GetConfigurationResp](#svc-infra-setting-GetConfigurationResp) | 获取配置 |
| ListConfigurations | [ListConfigurationsReq](#svc-infra-setting-ListConfigurationsReq) | [ListConfigurationsResp](#svc-infra-setting-ListConfigurationsResp) | 获取配置列表 |
| AddConfiguration | [AddConfigurationReq](#svc-infra-setting-AddConfigurationReq) | [AddConfigurationResp](#svc-infra-setting-AddConfigurationResp) | 添加配置 |
| UpdateConfiguration | [UpdateConfigurationReq](#svc-infra-setting-UpdateConfigurationReq) | [UpdateConfigurationResp](#svc-infra-setting-UpdateConfigurationResp) | 更新配置 |
| DeleteConfiguration | [DeleteConfigurationReq](#svc-infra-setting-DeleteConfigurationReq) | [DeleteConfigurationResp](#svc-infra-setting-DeleteConfigurationResp) | 删除配置 |
| Greeting | [SettingGreetingReq](#svc-infra-setting-SettingGreetingReq) | [SettingGreetingResp](#svc-infra-setting-SettingGreetingResp) |  |

 



<a name="svc-web-viewer_viewer-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.web.viewer/viewer.proto


 

 

 


<a name="svc-web-viewer-Viewer"></a>

### Viewer


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="svc-biz-asset_asset-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.asset/asset.proto



<a name="svc-biz-asset-ChangeMerchantCoinResp"></a>

### ChangeMerchantCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| new_value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  |  |
| change_value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |






<a name="svc-biz-asset-ChangeStreamerCoinResp"></a>

### ChangeStreamerCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| new_value | [StreamerCoinValue](#svc-biz-asset-StreamerCoinValue) |  |  |
| change_value | [StreamerCoinValue](#svc-biz-asset-StreamerCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |






<a name="svc-biz-asset-ChangeStreamerMoneyResp"></a>

### ChangeStreamerMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| new_value | [int64](#int64) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-ChangeUnionCoinResp"></a>

### ChangeUnionCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| new_value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  |  |
| change_value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |






<a name="svc-biz-asset-ChangeUnionMoneyResp"></a>

### ChangeUnionMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| new_value | [int64](#int64) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-ChangeUserCoinResp"></a>

### ChangeUserCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| new_value | [UserCoinValue](#svc-biz-asset-UserCoinValue) |  |  |
| change_value | [UserCoinValue](#svc-biz-asset-UserCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |






<a name="svc-biz-asset-ChangeUserMoneyResp"></a>

### ChangeUserMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| uid | [string](#string) |  |  |
| new_value | [int64](#int64) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrMerchantCoinReq"></a>

### DecrMerchantCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| rule | [int64](#int64) |  | 扣减规则: 0 仅普通余额 1 仅低权限余额 |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrStreamerCoinReq"></a>

### DecrStreamerCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| rule | [int64](#int64) |  | 扣减规则: 0 仅普通余额 1 仅低权限余额 |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrStreamerMoneyReq"></a>

### DecrStreamerMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrUnionCoinReq"></a>

### DecrUnionCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| rule | [int64](#int64) |  | 扣减规则: 0 仅普通余额 1 仅低权限余额 |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrUnionMoneyReq"></a>

### DecrUnionMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrUserCoinReq"></a>

### DecrUserCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| rule | [int64](#int64) |  | 扣减规则: 0 仅普通余额 1 仅低权限余额 2 先扣普通后扣低权限 3 先扣低权限后扣普通 |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrUserMoneyReq"></a>

### DecrUserMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| uid | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-GetMerchantCoinMultiReq"></a>

### GetMerchantCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_ids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetMerchantCoinMultiResp"></a>

### GetMerchantCoinMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetMerchantCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetMerchantCoinMultiResp-ValueMapEntry) | repeated | merchant_id-&gt;value |






<a name="svc-biz-asset-GetMerchantCoinMultiResp-ValueMapEntry"></a>

### GetMerchantCoinMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  |  |






<a name="svc-biz-asset-GetMerchantCoinReq"></a>

### GetMerchantCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetMerchantCoinResp"></a>

### GetMerchantCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  | 余额 |






<a name="svc-biz-asset-GetStreamerCoinMultiReq"></a>

### GetStreamerCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_ids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetStreamerCoinMultiResp"></a>

### GetStreamerCoinMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetStreamerCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetStreamerCoinMultiResp-ValueMapEntry) | repeated | streamer_id-&gt;value |






<a name="svc-biz-asset-GetStreamerCoinMultiResp-ValueMapEntry"></a>

### GetStreamerCoinMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [StreamerCoinValue](#svc-biz-asset-StreamerCoinValue) |  |  |






<a name="svc-biz-asset-GetStreamerCoinReq"></a>

### GetStreamerCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetStreamerCoinResp"></a>

### GetStreamerCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [StreamerCoinValue](#svc-biz-asset-StreamerCoinValue) |  | 余额 |






<a name="svc-biz-asset-GetStreamerMoneyMultiReq"></a>

### GetStreamerMoneyMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_ids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetStreamerMoneyMultiResp"></a>

### GetStreamerMoneyMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetStreamerMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetStreamerMoneyMultiResp-ValueMapEntry) | repeated | streamer_id-&gt;value |






<a name="svc-biz-asset-GetStreamerMoneyMultiResp-ValueMapEntry"></a>

### GetStreamerMoneyMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-GetStreamerMoneyReq"></a>

### GetStreamerMoneyReq
余额 money-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetStreamerMoneyResp"></a>

### GetStreamerMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-GetUnionCoinMultiReq"></a>

### GetUnionCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_ids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetUnionCoinMultiResp"></a>

### GetUnionCoinMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetUnionCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetUnionCoinMultiResp-ValueMapEntry) | repeated | union_id-&gt;value |






<a name="svc-biz-asset-GetUnionCoinMultiResp-ValueMapEntry"></a>

### GetUnionCoinMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  |  |






<a name="svc-biz-asset-GetUnionCoinReq"></a>

### GetUnionCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetUnionCoinResp"></a>

### GetUnionCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  | 余额 |






<a name="svc-biz-asset-GetUnionMoneyMultiReq"></a>

### GetUnionMoneyMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_ids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetUnionMoneyMultiResp"></a>

### GetUnionMoneyMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetUnionMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetUnionMoneyMultiResp-ValueMapEntry) | repeated | union_id-&gt;value |






<a name="svc-biz-asset-GetUnionMoneyMultiResp-ValueMapEntry"></a>

### GetUnionMoneyMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-GetUnionMoneyReq"></a>

### GetUnionMoneyReq
余额 money-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetUnionMoneyResp"></a>

### GetUnionMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-GetUserCoinMultiReq"></a>

### GetUserCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetUserCoinMultiResp"></a>

### GetUserCoinMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetUserCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetUserCoinMultiResp-ValueMapEntry) | repeated | uid-&gt;value |






<a name="svc-biz-asset-GetUserCoinMultiResp-ValueMapEntry"></a>

### GetUserCoinMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [UserCoinValue](#svc-biz-asset-UserCoinValue) |  |  |






<a name="svc-biz-asset-GetUserCoinReq"></a>

### GetUserCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |






<a name="svc-biz-asset-GetUserCoinResp"></a>

### GetUserCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [UserCoinValue](#svc-biz-asset-UserCoinValue) |  | 余额 |






<a name="svc-biz-asset-GetUserMoneyMultiReq"></a>

### GetUserMoneyMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| uids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetUserMoneyMultiResp"></a>

### GetUserMoneyMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| value_map | [GetUserMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetUserMoneyMultiResp-ValueMapEntry) | repeated | uid-&gt;value |






<a name="svc-biz-asset-GetUserMoneyMultiResp-ValueMapEntry"></a>

### GetUserMoneyMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-GetUserMoneyReq"></a>

### GetUserMoneyReq
余额 money-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| uid | [string](#string) |  |  |






<a name="svc-biz-asset-GetUserMoneyResp"></a>

### GetUserMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-IncrMerchantCoinReq"></a>

### IncrMerchantCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrStreamerCoinReq"></a>

### IncrStreamerCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| value | [StreamerCoinValue](#svc-biz-asset-StreamerCoinValue) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrStreamerMoneyReq"></a>

### IncrStreamerMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrUnionCoinReq"></a>

### IncrUnionCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrUnionMoneyReq"></a>

### IncrUnionMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrUserCoinReq"></a>

### IncrUserCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| value | [UserCoinValue](#svc-biz-asset-UserCoinValue) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrUserMoneyReq"></a>

### IncrUserMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| uid | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-ListMerchantCoinDetailReq"></a>

### ListMerchantCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListMerchantCoinDetailResp"></a>

### ListMerchantCoinDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MerchantCoinDetail](#svc-biz-asset-MerchantCoinDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListStreamerCoinDetailReq"></a>

### ListStreamerCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListStreamerCoinDetailResp"></a>

### ListStreamerCoinDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [StreamerCoinDetail](#svc-biz-asset-StreamerCoinDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListStreamerMoneyDetailReq"></a>

### ListStreamerMoneyDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListStreamerMoneyDetailResp"></a>

### ListStreamerMoneyDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [StreamerMoneyDetail](#svc-biz-asset-StreamerMoneyDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListUnionCoinDetailReq"></a>

### ListUnionCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListUnionCoinDetailResp"></a>

### ListUnionCoinDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UnionCoinDetail](#svc-biz-asset-UnionCoinDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListUnionMoneyDetailReq"></a>

### ListUnionMoneyDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListUnionMoneyDetailResp"></a>

### ListUnionMoneyDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UnionMoneyDetail](#svc-biz-asset-UnionMoneyDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListUserCoinDetailReq"></a>

### ListUserCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListUserCoinDetailResp"></a>

### ListUserCoinDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UserCoinDetail](#svc-biz-asset-UserCoinDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListUserMoneyDetailReq"></a>

### ListUserMoneyDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| uid | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListUserMoneyDetailResp"></a>

### ListUserMoneyDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [UserMoneyDetail](#svc-biz-asset-UserMoneyDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-MerchantCoinDetail"></a>

### MerchantCoinDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| limited_value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-MerchantCoinValue"></a>

### MerchantCoinValue
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-StreamerCoinDetail"></a>

### StreamerCoinDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| limited_value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-StreamerCoinValue"></a>

### StreamerCoinValue
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-StreamerMoneyDetail"></a>

### StreamerMoneyDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value_change | [ValueChange](#svc-biz-asset-ValueChange) |  | 交易类别 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-UnionCoinDetail"></a>

### UnionCoinDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| limited_value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-UnionCoinValue"></a>

### UnionCoinValue
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-UnionMoneyDetail"></a>

### UnionMoneyDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union_id | [string](#string) |  |  |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value_change | [ValueChange](#svc-biz-asset-ValueChange) |  | 交易类别 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-UserCoinDetail"></a>

### UserCoinDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| limited_value_change | [ValueChange](#svc-biz-asset-ValueChange) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-UserCoinValue"></a>

### UserCoinValue
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-UserMoneyDetail"></a>

### UserMoneyDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant_id | [string](#string) |  |  |
| uid | [string](#string) |  |  |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value_change | [ValueChange](#svc-biz-asset-ValueChange) |  | 交易类别 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-ValueChange"></a>

### ValueChange



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |
| value_before | [int64](#int64) |  |  |
| value_after | [int64](#int64) |  |  |





 

 

 


<a name="svc-biz-asset-Asset"></a>

### Asset


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUserMoney | [GetUserMoneyReq](#svc-biz-asset-GetUserMoneyReq) | [GetUserMoneyResp](#svc-biz-asset-GetUserMoneyResp) | 余额 money |
| GetUserMoneyMulti | [GetUserMoneyMultiReq](#svc-biz-asset-GetUserMoneyMultiReq) | [GetUserMoneyMultiResp](#svc-biz-asset-GetUserMoneyMultiResp) |  |
| IncrUserMoney | [IncrUserMoneyReq](#svc-biz-asset-IncrUserMoneyReq) | [ChangeUserMoneyResp](#svc-biz-asset-ChangeUserMoneyResp) |  |
| DecrUserMoney | [DecrUserMoneyReq](#svc-biz-asset-DecrUserMoneyReq) | [ChangeUserMoneyResp](#svc-biz-asset-ChangeUserMoneyResp) |  |
| ListUserMoneyDetail | [ListUserMoneyDetailReq](#svc-biz-asset-ListUserMoneyDetailReq) | [ListUserMoneyDetailResp](#svc-biz-asset-ListUserMoneyDetailResp) |  |
| GetUserCoin | [GetUserCoinReq](#svc-biz-asset-GetUserCoinReq) | [GetUserCoinResp](#svc-biz-asset-GetUserCoinResp) | 虚拟币coin |
| GetUserCoinMulti | [GetUserCoinMultiReq](#svc-biz-asset-GetUserCoinMultiReq) | [GetUserCoinMultiResp](#svc-biz-asset-GetUserCoinMultiResp) |  |
| IncrUserCoin | [IncrUserCoinReq](#svc-biz-asset-IncrUserCoinReq) | [ChangeUserCoinResp](#svc-biz-asset-ChangeUserCoinResp) |  |
| DecrUserCoin | [DecrUserCoinReq](#svc-biz-asset-DecrUserCoinReq) | [ChangeUserCoinResp](#svc-biz-asset-ChangeUserCoinResp) |  |
| ListUserCoinDetail | [ListUserCoinDetailReq](#svc-biz-asset-ListUserCoinDetailReq) | [ListUserCoinDetailResp](#svc-biz-asset-ListUserCoinDetailResp) |  |
| GetStreamerMoney | [GetStreamerMoneyReq](#svc-biz-asset-GetStreamerMoneyReq) | [GetStreamerMoneyResp](#svc-biz-asset-GetStreamerMoneyResp) | ---------------Streamer主播资产--------------- 余额 money |
| GetStreamerMoneyMulti | [GetStreamerMoneyMultiReq](#svc-biz-asset-GetStreamerMoneyMultiReq) | [GetStreamerMoneyMultiResp](#svc-biz-asset-GetStreamerMoneyMultiResp) |  |
| IncrStreamerMoney | [IncrStreamerMoneyReq](#svc-biz-asset-IncrStreamerMoneyReq) | [ChangeStreamerMoneyResp](#svc-biz-asset-ChangeStreamerMoneyResp) |  |
| DecrStreamerMoney | [DecrStreamerMoneyReq](#svc-biz-asset-DecrStreamerMoneyReq) | [ChangeStreamerMoneyResp](#svc-biz-asset-ChangeStreamerMoneyResp) |  |
| ListStreamerMoneyDetail | [ListStreamerMoneyDetailReq](#svc-biz-asset-ListStreamerMoneyDetailReq) | [ListStreamerMoneyDetailResp](#svc-biz-asset-ListStreamerMoneyDetailResp) |  |
| GetStreamerCoin | [GetStreamerCoinReq](#svc-biz-asset-GetStreamerCoinReq) | [GetStreamerCoinResp](#svc-biz-asset-GetStreamerCoinResp) | 虚拟币coin |
| GetStreamerCoinMulti | [GetStreamerCoinMultiReq](#svc-biz-asset-GetStreamerCoinMultiReq) | [GetStreamerCoinMultiResp](#svc-biz-asset-GetStreamerCoinMultiResp) |  |
| IncrStreamerCoin | [IncrStreamerCoinReq](#svc-biz-asset-IncrStreamerCoinReq) | [ChangeStreamerCoinResp](#svc-biz-asset-ChangeStreamerCoinResp) |  |
| DecrStreamerCoin | [DecrStreamerCoinReq](#svc-biz-asset-DecrStreamerCoinReq) | [ChangeStreamerCoinResp](#svc-biz-asset-ChangeStreamerCoinResp) |  |
| ListStreamerCoinDetail | [ListStreamerCoinDetailReq](#svc-biz-asset-ListStreamerCoinDetailReq) | [ListStreamerCoinDetailResp](#svc-biz-asset-ListStreamerCoinDetailResp) |  |
| GetUnionMoney | [GetUnionMoneyReq](#svc-biz-asset-GetUnionMoneyReq) | [GetUnionMoneyResp](#svc-biz-asset-GetUnionMoneyResp) | ---------------Union工会资产--------------- 余额 money -- |
| GetUnionMoneyMulti | [GetUnionMoneyMultiReq](#svc-biz-asset-GetUnionMoneyMultiReq) | [GetUnionMoneyMultiResp](#svc-biz-asset-GetUnionMoneyMultiResp) |  |
| IncrUnionMoney | [IncrUnionMoneyReq](#svc-biz-asset-IncrUnionMoneyReq) | [ChangeUnionMoneyResp](#svc-biz-asset-ChangeUnionMoneyResp) |  |
| DecrUnionMoney | [DecrUnionMoneyReq](#svc-biz-asset-DecrUnionMoneyReq) | [ChangeUnionMoneyResp](#svc-biz-asset-ChangeUnionMoneyResp) |  |
| ListUnionMoneyDetail | [ListUnionMoneyDetailReq](#svc-biz-asset-ListUnionMoneyDetailReq) | [ListUnionMoneyDetailResp](#svc-biz-asset-ListUnionMoneyDetailResp) |  |
| GetUnionCoin | [GetUnionCoinReq](#svc-biz-asset-GetUnionCoinReq) | [GetUnionCoinResp](#svc-biz-asset-GetUnionCoinResp) | 虚拟币coin |
| GetUnionCoinMulti | [GetUnionCoinMultiReq](#svc-biz-asset-GetUnionCoinMultiReq) | [GetUnionCoinMultiResp](#svc-biz-asset-GetUnionCoinMultiResp) |  |
| IncrUnionCoin | [IncrUnionCoinReq](#svc-biz-asset-IncrUnionCoinReq) | [ChangeUnionCoinResp](#svc-biz-asset-ChangeUnionCoinResp) |  |
| DecrUnionCoin | [DecrUnionCoinReq](#svc-biz-asset-DecrUnionCoinReq) | [ChangeUnionCoinResp](#svc-biz-asset-ChangeUnionCoinResp) |  |
| ListUnionCoinDetail | [ListUnionCoinDetailReq](#svc-biz-asset-ListUnionCoinDetailReq) | [ListUnionCoinDetailResp](#svc-biz-asset-ListUnionCoinDetailResp) |  |
| GetMerchantCoin | [GetMerchantCoinReq](#svc-biz-asset-GetMerchantCoinReq) | [GetMerchantCoinResp](#svc-biz-asset-GetMerchantCoinResp) | 虚拟币coin |
| GetMerchantCoinMulti | [GetMerchantCoinMultiReq](#svc-biz-asset-GetMerchantCoinMultiReq) | [GetMerchantCoinMultiResp](#svc-biz-asset-GetMerchantCoinMultiResp) |  |
| IncrMerchantCoin | [IncrMerchantCoinReq](#svc-biz-asset-IncrMerchantCoinReq) | [ChangeMerchantCoinResp](#svc-biz-asset-ChangeMerchantCoinResp) |  |
| DecrMerchantCoin | [DecrMerchantCoinReq](#svc-biz-asset-DecrMerchantCoinReq) | [ChangeMerchantCoinResp](#svc-biz-asset-ChangeMerchantCoinResp) |  |
| ListMerchantCoinDetail | [ListMerchantCoinDetailReq](#svc-biz-asset-ListMerchantCoinDetailReq) | [ListMerchantCoinDetailResp](#svc-biz-asset-ListMerchantCoinDetailResp) |  |

 



<a name="third_party_buf_validate_priv_private-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## third_party/buf/validate/priv/private.proto



<a name="buf-validate-priv-Constraint"></a>

### Constraint
Do not use. Internal to protovalidate library


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| message | [string](#string) |  |  |
| expression | [string](#string) |  |  |






<a name="buf-validate-priv-FieldConstraints"></a>

### FieldConstraints
Do not use. Internal to protovalidate library


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cel | [Constraint](#buf-validate-priv-Constraint) | repeated |  |





 

 


<a name="third_party_buf_validate_priv_private-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| field | FieldConstraints | .google.protobuf.FieldOptions | 1160 | Do not use. Internal to protovalidate library |

 

 



<a name="third_party_buf_validate_expression-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## third_party/buf/validate/expression.proto



<a name="buf-validate-Constraint"></a>

### Constraint
`Constraint` represents a validation rule written in the Common Expression
Language (CEL) syntax. Each Constraint includes a unique identifier, an
optional error message, and the CEL expression to evaluate. For more
information on CEL, [see our documentation](https://github.com/bufbuild/protovalidate/blob/main/docs/cel.md).

```proto
message Foo {
  option (buf.validate.message).cel = {
    id: &#34;foo.bar&#34;
    message: &#34;bar must be greater than 0&#34;
    expression: &#34;this.bar &gt; 0&#34;
  };
  int32 bar = 1;
}
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | `id` is a string that serves as a machine-readable name for this Constraint. It should be unique within its scope, which could be either a message or a field. |
| message | [string](#string) |  | `message` is an optional field that provides a human-readable error message for this Constraint when the CEL expression evaluates to false. If a non-empty message is provided, any strings resulting from the CEL expression evaluation are ignored. |
| expression | [string](#string) |  | `expression` is the actual CEL expression that will be evaluated for validation. This string must resolve to either a boolean or a string value. If the expression evaluates to false or a non-empty string, the validation is considered failed, and the message is rejected. |






<a name="buf-validate-Violation"></a>

### Violation
`Violation` represents a single instance where a validation rule, expressed
as a `Constraint`, was not met. It provides information about the field that
caused the violation, the specific constraint that wasn&#39;t fulfilled, and a
human-readable error message.

```json
{
  &#34;fieldPath&#34;: &#34;bar&#34;,
  &#34;constraintId&#34;: &#34;foo.bar&#34;,
  &#34;message&#34;: &#34;bar must be greater than 0&#34;
}
```


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| field_path | [string](#string) |  | `field_path` is a machine-readable identifier that points to the specific field that failed the validation. This could be a nested field, in which case the path will include all the parent fields leading to the actual field that caused the violation. |
| constraint_id | [string](#string) |  | `constraint_id` is the unique identifier of the `Constraint` that was not fulfilled. This is the same `id` that was specified in the `Constraint` message, allowing easy tracing of which rule was violated. |
| message | [string](#string) |  | `message` is a human-readable error message that describes the nature of the violation. This can be the default error message from the violated `Constraint`, or it can be a custom message that gives more context about the violation. |
| for_key | [bool](#bool) |  | `for_key` indicates whether the violation was caused by a map key, rather than a value. |






<a name="buf-validate-Violations"></a>

### Violations
`Violations` is a collection of `Violation` messages. This message type is returned by
protovalidate when a proto message fails to meet the requirements set by the `Constraint` validation rules.
Each individual violation is represented by a `Violation` message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| violations | [Violation](#buf-validate-Violation) | repeated | `violations` is a repeated field that contains all the `Violation` messages corresponding to the violations detected. |





 

 

 

 



<a name="third_party_buf_validate_validate-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## third_party/buf/validate/validate.proto



<a name="buf-validate-AnyRules"></a>

### AnyRules
AnyRules describe constraints applied exclusively to the `google.protobuf.Any` well-known type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| in | [string](#string) | repeated | `in` requires the field&#39;s `type_url` to be equal to one of the specified values. If it doesn&#39;t match any of the specified values, an error message is generated.

```proto message MyAny { // The `value` field must have a `type_url` equal to one of the specified values. google.protobuf.Any value = 1 [(buf.validate.field).any.in = [&#34;type.googleapis.com/MyType1&#34;, &#34;type.googleapis.com/MyType2&#34;]]; } ``` |
| not_in | [string](#string) | repeated | requires the field&#39;s type_url to be not equal to any of the specified values. If it matches any of the specified values, an error message is generated.

```proto message MyAny { // The field `value` must not have a `type_url` equal to any of the specified values. google.protobuf.Any value = 1 [(buf.validate.field).any.not_in = [&#34;type.googleapis.com/ForbiddenType1&#34;, &#34;type.googleapis.com/ForbiddenType2&#34;]]; } ``` |






<a name="buf-validate-BoolRules"></a>

### BoolRules
BoolRules describes the constraints applied to `bool` values. These rules
may also be applied to the `google.protobuf.BoolValue` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [bool](#bool) | optional | `const` requires the field value to exactly match the specified boolean value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyBool { // value must equal true bool value = 1 [(buf.validate.field).bool.const = true]; } ``` |






<a name="buf-validate-BytesRules"></a>

### BytesRules
BytesRules describe the constraints applied to `bytes` values. These rules
may also be applied to the `google.protobuf.BytesValue` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [bytes](#bytes) | optional | `const` requires the field value to exactly match the specified bytes value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyBytes { // value must be &#34;\x01\x02\x03\x04&#34; bytes value = 1 [(buf.validate.field).bytes.const = &#34;\x01\x02\x03\x04&#34;]; } ``` |
| len | [uint64](#uint64) | optional | `len` requires the field value to have the specified length in bytes. If the field value doesn&#39;t match, an error message is generated.

```proto message MyBytes { // value length must be 4 bytes. optional bytes value = 1 [(buf.validate.field).bytes.len = 4]; } ``` |
| min_len | [uint64](#uint64) | optional | `min_len` requires the field value to have at least the specified minimum length in bytes. If the field value doesn&#39;t meet the requirement, an error message is generated.

```proto message MyBytes { // value length must be at least 2 bytes. optional bytes value = 1 [(buf.validate.field).bytes.min_len = 2]; } ``` |
| max_len | [uint64](#uint64) | optional | `max_len` requires the field value to have at most the specified maximum length in bytes. If the field value exceeds the requirement, an error message is generated.

```proto message MyBytes { // value must be at most 6 bytes. optional bytes value = 1 [(buf.validate.field).bytes.max_len = 6]; } ``` |
| pattern | [string](#string) | optional | `pattern` requires the field value to match the specified regular expression ([RE2 syntax](https://github.com/google/re2/wiki/Syntax)). The value of the field must be valid UTF-8 or validation will fail with a runtime error. If the field value doesn&#39;t match the pattern, an error message is generated.

```proto message MyBytes { // value must match regex pattern &#34;^[a-zA-Z0-9]&#43;$&#34;. optional bytes value = 1 [(buf.validate.field).bytes.pattern = &#34;^[a-zA-Z0-9]&#43;$&#34;]; } ``` |
| prefix | [bytes](#bytes) | optional | `prefix` requires the field value to have the specified bytes at the beginning of the string. If the field value doesn&#39;t meet the requirement, an error message is generated.

```proto message MyBytes { // value does not have prefix \x01\x02 optional bytes value = 1 [(buf.validate.field).bytes.prefix = &#34;\x01\x02&#34;]; } ``` |
| suffix | [bytes](#bytes) | optional | `suffix` requires the field value to have the specified bytes at the end of the string. If the field value doesn&#39;t meet the requirement, an error message is generated.

```proto message MyBytes { // value does not have suffix \x03\x04 optional bytes value = 1 [(buf.validate.field).bytes.suffix = &#34;\x03\x04&#34;]; } ``` |
| contains | [bytes](#bytes) | optional | `contains` requires the field value to have the specified bytes anywhere in the string. If the field value doesn&#39;t meet the requirement, an error message is generated.

```protobuf message MyBytes { // value does not contain \x02\x03 optional bytes value = 1 [(buf.validate.field).bytes.contains = &#34;\x02\x03&#34;]; } ``` |
| in | [bytes](#bytes) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value doesn&#39;t match any of the specified values, an error message is generated.

```protobuf message MyBytes { // value must in [&#34;\x01\x02&#34;, &#34;\x02\x03&#34;, &#34;\x03\x04&#34;] optional bytes value = 1 [(buf.validate.field).bytes.in = {&#34;\x01\x02&#34;, &#34;\x02\x03&#34;, &#34;\x03\x04&#34;}]; } ``` |
| not_in | [bytes](#bytes) | repeated | `not_in` requires the field value to be not equal to any of the specified values. If the field value matches any of the specified values, an error message is generated.

```proto message MyBytes { // value must not in [&#34;\x01\x02&#34;, &#34;\x02\x03&#34;, &#34;\x03\x04&#34;] optional bytes value = 1 [(buf.validate.field).bytes.not_in = {&#34;\x01\x02&#34;, &#34;\x02\x03&#34;, &#34;\x03\x04&#34;}]; } ``` |
| ip | [bool](#bool) |  | `ip` ensures that the field `value` is a valid IP address (v4 or v6) in byte format. If the field value doesn&#39;t meet this constraint, an error message is generated.

```proto message MyBytes { // value must be a valid IP address optional bytes value = 1 [(buf.validate.field).bytes.ip = true]; } ``` |
| ipv4 | [bool](#bool) |  | `ipv4` ensures that the field `value` is a valid IPv4 address in byte format. If the field value doesn&#39;t meet this constraint, an error message is generated.

```proto message MyBytes { // value must be a valid IPv4 address optional bytes value = 1 [(buf.validate.field).bytes.ipv4 = true]; } ``` |
| ipv6 | [bool](#bool) |  | `ipv6` ensures that the field `value` is a valid IPv6 address in byte format. If the field value doesn&#39;t meet this constraint, an error message is generated. ```proto message MyBytes { // value must be a valid IPv6 address optional bytes value = 1 [(buf.validate.field).bytes.ipv6 = true]; } ``` |






<a name="buf-validate-DoubleRules"></a>

### DoubleRules
DoubleRules describes the constraints applied to `double` values. These
rules may also be applied to the `google.protobuf.DoubleValue` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [double](#double) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyDouble { // value must equal 42.0 double value = 1 [(buf.validate.field).double.const = 42.0]; } ``` |
| lt | [double](#double) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyDouble { // value must be less than 10.0 double value = 1 [(buf.validate.field).double.lt = 10.0]; } ``` |
| lte | [double](#double) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyDouble { // value must be less than or equal to 10.0 double value = 1 [(buf.validate.field).double.lte = 10.0]; } ``` |
| gt | [double](#double) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyDouble { // value must be greater than 5.0 [double.gt] double value = 1 [(buf.validate.field).double.gt = 5.0];

 // value must be greater than 5 and less than 10.0 [double.gt_lt] double other_value = 2 [(buf.validate.field).double = { gt: 5.0, lt: 10.0 }];

 // value must be greater than 10 or less than 5.0 [double.gt_lt_exclusive] double another_value = 3 [(buf.validate.field).double = { gt: 10.0, lt: 5.0 }]; } ``` |
| gte | [double](#double) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyDouble { // value must be greater than or equal to 5.0 [double.gte] double value = 1 [(buf.validate.field).double.gte = 5.0];

 // value must be greater than or equal to 5.0 and less than 10.0 [double.gte_lt] double other_value = 2 [(buf.validate.field).double = { gte: 5.0, lt: 10.0 }];

 // value must be greater than or equal to 10.0 or less than 5.0 [double.gte_lt_exclusive] double another_value = 3 [(buf.validate.field).double = { gte: 10.0, lt: 5.0 }]; } ``` |
| in | [double](#double) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyDouble { // value must be in list [1.0, 2.0, 3.0] repeated double value = 1 (buf.validate.field).double = { in: [1.0, 2.0, 3.0] }; } ``` |
| not_in | [double](#double) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyDouble { // value must not be in list [1.0, 2.0, 3.0] repeated double value = 1 (buf.validate.field).double = { not_in: [1.0, 2.0, 3.0] }; } ``` |
| finite | [bool](#bool) |  | `finite` requires the field value to be finite. If the field value is infinite or NaN, an error message is generated. |






<a name="buf-validate-DurationRules"></a>

### DurationRules
DurationRules describe the constraints applied exclusively to the `google.protobuf.Duration` well-known type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [google.protobuf.Duration](#google-protobuf-Duration) | optional | `const` dictates that the field must match the specified value of the `google.protobuf.Duration` type exactly. If the field&#39;s value deviates from the specified value, an error message will be generated.

```proto message MyDuration { // value must equal 5s google.protobuf.Duration value = 1 [(buf.validate.field).duration.const = &#34;5s&#34;]; } ``` |
| lt | [google.protobuf.Duration](#google-protobuf-Duration) |  | `lt` stipulates that the field must be less than the specified value of the `google.protobuf.Duration` type, exclusive. If the field&#39;s value is greater than or equal to the specified value, an error message will be generated.

```proto message MyDuration { // value must be less than 5s google.protobuf.Duration value = 1 [(buf.validate.field).duration.lt = &#34;5s&#34;]; } ``` |
| lte | [google.protobuf.Duration](#google-protobuf-Duration) |  | `lte` indicates that the field must be less than or equal to the specified value of the `google.protobuf.Duration` type, inclusive. If the field&#39;s value is greater than the specified value, an error message will be generated.

```proto message MyDuration { // value must be less than or equal to 10s google.protobuf.Duration value = 1 [(buf.validate.field).duration.lte = &#34;10s&#34;]; } ``` |
| gt | [google.protobuf.Duration](#google-protobuf-Duration) |  | `gt` requires the duration field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyDuration { // duration must be greater than 5s [duration.gt] google.protobuf.Duration value = 1 [(buf.validate.field).duration.gt = { seconds: 5 }];

 // duration must be greater than 5s and less than 10s [duration.gt_lt] google.protobuf.Duration another_value = 2 [(buf.validate.field).duration = { gt: { seconds: 5 }, lt: { seconds: 10 } }];

 // duration must be greater than 10s or less than 5s [duration.gt_lt_exclusive] google.protobuf.Duration other_value = 3 [(buf.validate.field).duration = { gt: { seconds: 10 }, lt: { seconds: 5 } }]; } ``` |
| gte | [google.protobuf.Duration](#google-protobuf-Duration) |  | `gte` requires the duration field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyDuration { // duration must be greater than or equal to 5s [duration.gte] google.protobuf.Duration value = 1 [(buf.validate.field).duration.gte = { seconds: 5 }];

 // duration must be greater than or equal to 5s and less than 10s [duration.gte_lt] google.protobuf.Duration another_value = 2 [(buf.validate.field).duration = { gte: { seconds: 5 }, lt: { seconds: 10 } }];

 // duration must be greater than or equal to 10s or less than 5s [duration.gte_lt_exclusive] google.protobuf.Duration other_value = 3 [(buf.validate.field).duration = { gte: { seconds: 10 }, lt: { seconds: 5 } }]; } ``` |
| in | [google.protobuf.Duration](#google-protobuf-Duration) | repeated | `in` asserts that the field must be equal to one of the specified values of the `google.protobuf.Duration` type. If the field&#39;s value doesn&#39;t correspond to any of the specified values, an error message will be generated.

```proto message MyDuration { // value must be in list [1s, 2s, 3s] google.protobuf.Duration value = 1 [(buf.validate.field).duration.in = [&#34;1s&#34;, &#34;2s&#34;, &#34;3s&#34;]]; } ``` |
| not_in | [google.protobuf.Duration](#google-protobuf-Duration) | repeated | `not_in` denotes that the field must not be equal to any of the specified values of the `google.protobuf.Duration` type. If the field&#39;s value matches any of these values, an error message will be generated.

```proto message MyDuration { // value must not be in list [1s, 2s, 3s] google.protobuf.Duration value = 1 [(buf.validate.field).duration.not_in = [&#34;1s&#34;, &#34;2s&#34;, &#34;3s&#34;]]; } ``` |






<a name="buf-validate-EnumRules"></a>

### EnumRules
EnumRules describe the constraints applied to `enum` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [int32](#int32) | optional | `const` requires the field value to exactly match the specified enum value. If the field value doesn&#39;t match, an error message is generated.

```proto enum MyEnum { MY_ENUM_UNSPECIFIED = 0; MY_ENUM_VALUE1 = 1; MY_ENUM_VALUE2 = 2; }

message MyMessage { // The field `value` must be exactly MY_ENUM_VALUE1. MyEnum value = 1 [(buf.validate.field).enum.const = 1]; } ``` |
| defined_only | [bool](#bool) | optional | `defined_only` requires the field value to be one of the defined values for this enum, failing on any undefined value.

```proto enum MyEnum { MY_ENUM_UNSPECIFIED = 0; MY_ENUM_VALUE1 = 1; MY_ENUM_VALUE2 = 2; }

message MyMessage { // The field `value` must be a defined value of MyEnum. MyEnum value = 1 [(buf.validate.field).enum.defined_only = true]; } ``` |
| in | [int32](#int32) | repeated | `in` requires the field value to be equal to one of the specified enum values. If the field value doesn&#39;t match any of the specified values, an error message is generated.

```proto enum MyEnum { MY_ENUM_UNSPECIFIED = 0; MY_ENUM_VALUE1 = 1; MY_ENUM_VALUE2 = 2; }

message MyMessage { // The field `value` must be equal to one of the specified values. MyEnum value = 1 [(buf.validate.field).enum = { in: [1, 2]}]; } ``` |
| not_in | [int32](#int32) | repeated | `not_in` requires the field value to be not equal to any of the specified enum values. If the field value matches one of the specified values, an error message is generated.

```proto enum MyEnum { MY_ENUM_UNSPECIFIED = 0; MY_ENUM_VALUE1 = 1; MY_ENUM_VALUE2 = 2; }

message MyMessage { // The field `value` must not be equal to any of the specified values. MyEnum value = 1 [(buf.validate.field).enum = { not_in: [1, 2]}]; } ``` |






<a name="buf-validate-FieldConstraints"></a>

### FieldConstraints
FieldConstraints encapsulates the rules for each type of field. Depending on
the field, the correct set should be used to ensure proper validations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cel | [Constraint](#buf-validate-Constraint) | repeated | `cel` is a repeated field used to represent a textual expression in the Common Expression Language (CEL) syntax. For more information on CEL, [see our documentation](https://github.com/bufbuild/protovalidate/blob/main/docs/cel.md).

```proto message MyMessage { // The field `value` must be greater than 42. optional int32 value = 1 [(buf.validate.field).cel = { id: &#34;my_message.value&#34;, message: &#34;value must be greater than 42&#34;, expression: &#34;this &gt; 42&#34;, }]; } ``` |
| required | [bool](#bool) |  | If `required` is true, the field must be populated. A populated field can be described as &#34;serialized in the wire format,&#34; which includes:

- the following &#34;nullable&#34; fields must be explicitly set to be considered populated: - singular message fields (whose fields may be unpopulated/default values) - member fields of a oneof (may be their default value) - proto3 optional fields (may be their default value) - proto2 scalar fields (both optional and required) - proto3 scalar fields must be non-zero to be considered populated - repeated and map fields must be non-empty to be considered populated

```proto message MyMessage { // The field `value` must be set to a non-null value. optional MyOtherMessage value = 1 [(buf.validate.field).required = true]; } ``` |
| ignore | [Ignore](#buf-validate-Ignore) |  | Skip validation on the field if its value matches the specified criteria. See Ignore enum for details.

```proto message UpdateRequest { // The uri rule only applies if the field is populated and not an empty // string. optional string url = 1 [ (buf.validate.field).ignore = IGNORE_IF_DEFAULT_VALUE, (buf.validate.field).string.uri = true, ]; } ``` |
| float | [FloatRules](#buf-validate-FloatRules) |  | Scalar Field Types |
| double | [DoubleRules](#buf-validate-DoubleRules) |  |  |
| int32 | [Int32Rules](#buf-validate-Int32Rules) |  |  |
| int64 | [Int64Rules](#buf-validate-Int64Rules) |  |  |
| uint32 | [UInt32Rules](#buf-validate-UInt32Rules) |  |  |
| uint64 | [UInt64Rules](#buf-validate-UInt64Rules) |  |  |
| sint32 | [SInt32Rules](#buf-validate-SInt32Rules) |  |  |
| sint64 | [SInt64Rules](#buf-validate-SInt64Rules) |  |  |
| fixed32 | [Fixed32Rules](#buf-validate-Fixed32Rules) |  |  |
| fixed64 | [Fixed64Rules](#buf-validate-Fixed64Rules) |  |  |
| sfixed32 | [SFixed32Rules](#buf-validate-SFixed32Rules) |  |  |
| sfixed64 | [SFixed64Rules](#buf-validate-SFixed64Rules) |  |  |
| bool | [BoolRules](#buf-validate-BoolRules) |  |  |
| string | [StringRules](#buf-validate-StringRules) |  |  |
| bytes | [BytesRules](#buf-validate-BytesRules) |  |  |
| enum | [EnumRules](#buf-validate-EnumRules) |  | Complex Field Types |
| repeated | [RepeatedRules](#buf-validate-RepeatedRules) |  |  |
| map | [MapRules](#buf-validate-MapRules) |  |  |
| any | [AnyRules](#buf-validate-AnyRules) |  | Well-Known Field Types |
| duration | [DurationRules](#buf-validate-DurationRules) |  |  |
| timestamp | [TimestampRules](#buf-validate-TimestampRules) |  |  |
| skipped | [bool](#bool) |  | **Deprecated.** DEPRECATED: use ignore=IGNORE_ALWAYS instead. TODO: remove this field pre-v1. |
| ignore_empty | [bool](#bool) |  | **Deprecated.** DEPRECATED: use ignore=IGNORE_IF_UNPOPULATED instead. TODO: remove this field pre-v1. |






<a name="buf-validate-Fixed32Rules"></a>

### Fixed32Rules
Fixed32Rules describes the constraints applied to `fixed32` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [fixed32](#fixed32) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyFixed32 { // value must equal 42 fixed32 value = 1 [(buf.validate.field).fixed32.const = 42]; } ``` |
| lt | [fixed32](#fixed32) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyFixed32 { // value must be less than 10 fixed32 value = 1 [(buf.validate.field).fixed32.lt = 10]; } ``` |
| lte | [fixed32](#fixed32) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyFixed32 { // value must be less than or equal to 10 fixed32 value = 1 [(buf.validate.field).fixed32.lte = 10]; } ``` |
| gt | [fixed32](#fixed32) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyFixed32 { // value must be greater than 5 [fixed32.gt] fixed32 value = 1 [(buf.validate.field).fixed32.gt = 5];

 // value must be greater than 5 and less than 10 [fixed32.gt_lt] fixed32 other_value = 2 [(buf.validate.field).fixed32 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [fixed32.gt_lt_exclusive] fixed32 another_value = 3 [(buf.validate.field).fixed32 = { gt: 10, lt: 5 }]; } ``` |
| gte | [fixed32](#fixed32) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyFixed32 { // value must be greater than or equal to 5 [fixed32.gte] fixed32 value = 1 [(buf.validate.field).fixed32.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [fixed32.gte_lt] fixed32 other_value = 2 [(buf.validate.field).fixed32 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [fixed32.gte_lt_exclusive] fixed32 another_value = 3 [(buf.validate.field).fixed32 = { gte: 10, lt: 5 }]; } ``` |
| in | [fixed32](#fixed32) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyFixed32 { // value must be in list [1, 2, 3] repeated fixed32 value = 1 (buf.validate.field).fixed32 = { in: [1, 2, 3] }; } ``` |
| not_in | [fixed32](#fixed32) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyFixed32 { // value must not be in list [1, 2, 3] repeated fixed32 value = 1 (buf.validate.field).fixed32 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-Fixed64Rules"></a>

### Fixed64Rules
Fixed64Rules describes the constraints applied to `fixed64` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [fixed64](#fixed64) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyFixed64 { // value must equal 42 fixed64 value = 1 [(buf.validate.field).fixed64.const = 42]; } ``` |
| lt | [fixed64](#fixed64) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyFixed64 { // value must be less than 10 fixed64 value = 1 [(buf.validate.field).fixed64.lt = 10]; } ``` |
| lte | [fixed64](#fixed64) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyFixed64 { // value must be less than or equal to 10 fixed64 value = 1 [(buf.validate.field).fixed64.lte = 10]; } ``` |
| gt | [fixed64](#fixed64) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyFixed64 { // value must be greater than 5 [fixed64.gt] fixed64 value = 1 [(buf.validate.field).fixed64.gt = 5];

 // value must be greater than 5 and less than 10 [fixed64.gt_lt] fixed64 other_value = 2 [(buf.validate.field).fixed64 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [fixed64.gt_lt_exclusive] fixed64 another_value = 3 [(buf.validate.field).fixed64 = { gt: 10, lt: 5 }]; } ``` |
| gte | [fixed64](#fixed64) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyFixed64 { // value must be greater than or equal to 5 [fixed64.gte] fixed64 value = 1 [(buf.validate.field).fixed64.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [fixed64.gte_lt] fixed64 other_value = 2 [(buf.validate.field).fixed64 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [fixed64.gte_lt_exclusive] fixed64 another_value = 3 [(buf.validate.field).fixed64 = { gte: 10, lt: 5 }]; } ``` |
| in | [fixed64](#fixed64) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyFixed64 { // value must be in list [1, 2, 3] repeated fixed64 value = 1 (buf.validate.field).fixed64 = { in: [1, 2, 3] }; } ``` |
| not_in | [fixed64](#fixed64) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyFixed64 { // value must not be in list [1, 2, 3] repeated fixed64 value = 1 (buf.validate.field).fixed64 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-FloatRules"></a>

### FloatRules
FloatRules describes the constraints applied to `float` values. These
rules may also be applied to the `google.protobuf.FloatValue` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [float](#float) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyFloat { // value must equal 42.0 float value = 1 [(buf.validate.field).float.const = 42.0]; } ``` |
| lt | [float](#float) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyFloat { // value must be less than 10.0 float value = 1 [(buf.validate.field).float.lt = 10.0]; } ``` |
| lte | [float](#float) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyFloat { // value must be less than or equal to 10.0 float value = 1 [(buf.validate.field).float.lte = 10.0]; } ``` |
| gt | [float](#float) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyFloat { // value must be greater than 5.0 [float.gt] float value = 1 [(buf.validate.field).float.gt = 5.0];

 // value must be greater than 5 and less than 10.0 [float.gt_lt] float other_value = 2 [(buf.validate.field).float = { gt: 5.0, lt: 10.0 }];

 // value must be greater than 10 or less than 5.0 [float.gt_lt_exclusive] float another_value = 3 [(buf.validate.field).float = { gt: 10.0, lt: 5.0 }]; } ``` |
| gte | [float](#float) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyFloat { // value must be greater than or equal to 5.0 [float.gte] float value = 1 [(buf.validate.field).float.gte = 5.0];

 // value must be greater than or equal to 5.0 and less than 10.0 [float.gte_lt] float other_value = 2 [(buf.validate.field).float = { gte: 5.0, lt: 10.0 }];

 // value must be greater than or equal to 10.0 or less than 5.0 [float.gte_lt_exclusive] float another_value = 3 [(buf.validate.field).float = { gte: 10.0, lt: 5.0 }]; } ``` |
| in | [float](#float) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyFloat { // value must be in list [1.0, 2.0, 3.0] repeated float value = 1 (buf.validate.field).float = { in: [1.0, 2.0, 3.0] }; } ``` |
| not_in | [float](#float) | repeated | `in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyFloat { // value must not be in list [1.0, 2.0, 3.0] repeated float value = 1 (buf.validate.field).float = { not_in: [1.0, 2.0, 3.0] }; } ``` |
| finite | [bool](#bool) |  | `finite` requires the field value to be finite. If the field value is infinite or NaN, an error message is generated. |






<a name="buf-validate-Int32Rules"></a>

### Int32Rules
Int32Rules describes the constraints applied to `int32` values. These
rules may also be applied to the `google.protobuf.Int32Value` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [int32](#int32) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyInt32 { // value must equal 42 int32 value = 1 [(buf.validate.field).int32.const = 42]; } ``` |
| lt | [int32](#int32) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyInt32 { // value must be less than 10 int32 value = 1 [(buf.validate.field).int32.lt = 10]; } ``` |
| lte | [int32](#int32) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyInt32 { // value must be less than or equal to 10 int32 value = 1 [(buf.validate.field).int32.lte = 10]; } ``` |
| gt | [int32](#int32) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyInt32 { // value must be greater than 5 [int32.gt] int32 value = 1 [(buf.validate.field).int32.gt = 5];

 // value must be greater than 5 and less than 10 [int32.gt_lt] int32 other_value = 2 [(buf.validate.field).int32 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [int32.gt_lt_exclusive] int32 another_value = 3 [(buf.validate.field).int32 = { gt: 10, lt: 5 }]; } ``` |
| gte | [int32](#int32) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyInt32 { // value must be greater than or equal to 5 [int32.gte] int32 value = 1 [(buf.validate.field).int32.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [int32.gte_lt] int32 other_value = 2 [(buf.validate.field).int32 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [int32.gte_lt_exclusive] int32 another_value = 3 [(buf.validate.field).int32 = { gte: 10, lt: 5 }]; } ``` |
| in | [int32](#int32) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyInt32 { // value must be in list [1, 2, 3] repeated int32 value = 1 (buf.validate.field).int32 = { in: [1, 2, 3] }; } ``` |
| not_in | [int32](#int32) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyInt32 { // value must not be in list [1, 2, 3] repeated int32 value = 1 (buf.validate.field).int32 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-Int64Rules"></a>

### Int64Rules
Int64Rules describes the constraints applied to `int64` values. These
rules may also be applied to the `google.protobuf.Int64Value` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [int64](#int64) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyInt64 { // value must equal 42 int64 value = 1 [(buf.validate.field).int64.const = 42]; } ``` |
| lt | [int64](#int64) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyInt64 { // value must be less than 10 int64 value = 1 [(buf.validate.field).int64.lt = 10]; } ``` |
| lte | [int64](#int64) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyInt64 { // value must be less than or equal to 10 int64 value = 1 [(buf.validate.field).int64.lte = 10]; } ``` |
| gt | [int64](#int64) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyInt64 { // value must be greater than 5 [int64.gt] int64 value = 1 [(buf.validate.field).int64.gt = 5];

 // value must be greater than 5 and less than 10 [int64.gt_lt] int64 other_value = 2 [(buf.validate.field).int64 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [int64.gt_lt_exclusive] int64 another_value = 3 [(buf.validate.field).int64 = { gt: 10, lt: 5 }]; } ``` |
| gte | [int64](#int64) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyInt64 { // value must be greater than or equal to 5 [int64.gte] int64 value = 1 [(buf.validate.field).int64.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [int64.gte_lt] int64 other_value = 2 [(buf.validate.field).int64 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [int64.gte_lt_exclusive] int64 another_value = 3 [(buf.validate.field).int64 = { gte: 10, lt: 5 }]; } ``` |
| in | [int64](#int64) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyInt64 { // value must be in list [1, 2, 3] repeated int64 value = 1 (buf.validate.field).int64 = { in: [1, 2, 3] }; } ``` |
| not_in | [int64](#int64) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyInt64 { // value must not be in list [1, 2, 3] repeated int64 value = 1 (buf.validate.field).int64 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-MapRules"></a>

### MapRules
MapRules describe the constraints applied to `map` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_pairs | [uint64](#uint64) | optional | Specifies the minimum number of key-value pairs allowed. If the field has fewer key-value pairs than specified, an error message is generated.

```proto message MyMap { // The field `value` must have at least 2 key-value pairs. map&lt;string, string&gt; value = 1 [(buf.validate.field).map.min_pairs = 2]; } ``` |
| max_pairs | [uint64](#uint64) | optional | Specifies the maximum number of key-value pairs allowed. If the field has more key-value pairs than specified, an error message is generated.

```proto message MyMap { // The field `value` must have at most 3 key-value pairs. map&lt;string, string&gt; value = 1 [(buf.validate.field).map.max_pairs = 3]; } ``` |
| keys | [FieldConstraints](#buf-validate-FieldConstraints) | optional | Specifies the constraints to be applied to each key in the field.

```proto message MyMap { // The keys in the field `value` must follow the specified constraints. map&lt;string, string&gt; value = 1 [(buf.validate.field).map.keys = { string: { min_len: 3 max_len: 10 } }]; } ``` |
| values | [FieldConstraints](#buf-validate-FieldConstraints) | optional | Specifies the constraints to be applied to the value of each key in the field. Message values will still have their validations evaluated unless skip is specified here.

```proto message MyMap { // The values in the field `value` must follow the specified constraints. map&lt;string, string&gt; value = 1 [(buf.validate.field).map.values = { string: { min_len: 5 max_len: 20 } }]; } ``` |






<a name="buf-validate-MessageConstraints"></a>

### MessageConstraints
MessageConstraints represents validation rules that are applied to the entire message.
It includes disabling options and a list of Constraint messages representing Common Expression Language (CEL) validation rules.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disabled | [bool](#bool) | optional | `disabled` is a boolean flag that, when set to true, nullifies any validation rules for this message. This includes any fields within the message that would otherwise support validation.

```proto message MyMessage { // validation will be bypassed for this message option (buf.validate.message).disabled = true; } ``` |
| cel | [Constraint](#buf-validate-Constraint) | repeated | `cel` is a repeated field of type Constraint. Each Constraint specifies a validation rule to be applied to this message. These constraints are written in Common Expression Language (CEL) syntax. For more information on CEL, [see our documentation](https://github.com/bufbuild/protovalidate/blob/main/docs/cel.md).

```proto message MyMessage { // The field `foo` must be greater than 42. option (buf.validate.message).cel = { id: &#34;my_message.value&#34;, message: &#34;value must be greater than 42&#34;, expression: &#34;this.foo &gt; 42&#34;, }; optional int32 foo = 1; } ``` |






<a name="buf-validate-OneofConstraints"></a>

### OneofConstraints
The `OneofConstraints` message type enables you to manage constraints for
oneof fields in your protobuf messages.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| required | [bool](#bool) | optional | If `required` is true, exactly one field of the oneof must be present. A validation error is returned if no fields in the oneof are present. The field itself may still be a default value; further constraints should be placed on the fields themselves to ensure they are valid values, such as `min_len` or `gt`.

```proto message MyMessage { oneof value { // Either `a` or `b` must be set. If `a` is set, it must also be // non-empty; whereas if `b` is set, it can still be an empty string. option (buf.validate.oneof).required = true; string a = 1 [(buf.validate.field).string.min_len = 1]; string b = 2; } } ``` |






<a name="buf-validate-RepeatedRules"></a>

### RepeatedRules
RepeatedRules describe the constraints applied to `repeated` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_items | [uint64](#uint64) | optional | `min_items` requires that this field must contain at least the specified minimum number of items.

Note that `min_items = 1` is equivalent to setting a field as `required`.

```proto message MyRepeated { // value must contain at least 2 items repeated string value = 1 [(buf.validate.field).repeated.min_items = 2]; } ``` |
| max_items | [uint64](#uint64) | optional | `max_items` denotes that this field must not exceed a certain number of items as the upper limit. If the field contains more items than specified, an error message will be generated, requiring the field to maintain no more than the specified number of items.

```proto message MyRepeated { // value must contain no more than 3 item(s) repeated string value = 1 [(buf.validate.field).repeated.max_items = 3]; } ``` |
| unique | [bool](#bool) | optional | `unique` indicates that all elements in this field must be unique. This constraint is strictly applicable to scalar and enum types, with message types not being supported.

```proto message MyRepeated { // repeated value must contain unique items repeated string value = 1 [(buf.validate.field).repeated.unique = true]; } ``` |
| items | [FieldConstraints](#buf-validate-FieldConstraints) | optional | `items` details the constraints to be applied to each item in the field. Even for repeated message fields, validation is executed against each item unless skip is explicitly specified.

```proto message MyRepeated { // The items in the field `value` must follow the specified constraints. repeated string value = 1 [(buf.validate.field).repeated.items = { string: { min_len: 3 max_len: 10 } }]; } ``` |






<a name="buf-validate-SFixed32Rules"></a>

### SFixed32Rules
SFixed32Rules describes the constraints applied to `fixed32` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sfixed32](#sfixed32) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MySFixed32 { // value must equal 42 sfixed32 value = 1 [(buf.validate.field).sfixed32.const = 42]; } ``` |
| lt | [sfixed32](#sfixed32) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MySFixed32 { // value must be less than 10 sfixed32 value = 1 [(buf.validate.field).sfixed32.lt = 10]; } ``` |
| lte | [sfixed32](#sfixed32) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MySFixed32 { // value must be less than or equal to 10 sfixed32 value = 1 [(buf.validate.field).sfixed32.lte = 10]; } ``` |
| gt | [sfixed32](#sfixed32) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySFixed32 { // value must be greater than 5 [sfixed32.gt] sfixed32 value = 1 [(buf.validate.field).sfixed32.gt = 5];

 // value must be greater than 5 and less than 10 [sfixed32.gt_lt] sfixed32 other_value = 2 [(buf.validate.field).sfixed32 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [sfixed32.gt_lt_exclusive] sfixed32 another_value = 3 [(buf.validate.field).sfixed32 = { gt: 10, lt: 5 }]; } ``` |
| gte | [sfixed32](#sfixed32) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySFixed32 { // value must be greater than or equal to 5 [sfixed32.gte] sfixed32 value = 1 [(buf.validate.field).sfixed32.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [sfixed32.gte_lt] sfixed32 other_value = 2 [(buf.validate.field).sfixed32 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [sfixed32.gte_lt_exclusive] sfixed32 another_value = 3 [(buf.validate.field).sfixed32 = { gte: 10, lt: 5 }]; } ``` |
| in | [sfixed32](#sfixed32) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MySFixed32 { // value must be in list [1, 2, 3] repeated sfixed32 value = 1 (buf.validate.field).sfixed32 = { in: [1, 2, 3] }; } ``` |
| not_in | [sfixed32](#sfixed32) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MySFixed32 { // value must not be in list [1, 2, 3] repeated sfixed32 value = 1 (buf.validate.field).sfixed32 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-SFixed64Rules"></a>

### SFixed64Rules
SFixed64Rules describes the constraints applied to `fixed64` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sfixed64](#sfixed64) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MySFixed64 { // value must equal 42 sfixed64 value = 1 [(buf.validate.field).sfixed64.const = 42]; } ``` |
| lt | [sfixed64](#sfixed64) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MySFixed64 { // value must be less than 10 sfixed64 value = 1 [(buf.validate.field).sfixed64.lt = 10]; } ``` |
| lte | [sfixed64](#sfixed64) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MySFixed64 { // value must be less than or equal to 10 sfixed64 value = 1 [(buf.validate.field).sfixed64.lte = 10]; } ``` |
| gt | [sfixed64](#sfixed64) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySFixed64 { // value must be greater than 5 [sfixed64.gt] sfixed64 value = 1 [(buf.validate.field).sfixed64.gt = 5];

 // value must be greater than 5 and less than 10 [sfixed64.gt_lt] sfixed64 other_value = 2 [(buf.validate.field).sfixed64 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [sfixed64.gt_lt_exclusive] sfixed64 another_value = 3 [(buf.validate.field).sfixed64 = { gt: 10, lt: 5 }]; } ``` |
| gte | [sfixed64](#sfixed64) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySFixed64 { // value must be greater than or equal to 5 [sfixed64.gte] sfixed64 value = 1 [(buf.validate.field).sfixed64.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [sfixed64.gte_lt] sfixed64 other_value = 2 [(buf.validate.field).sfixed64 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [sfixed64.gte_lt_exclusive] sfixed64 another_value = 3 [(buf.validate.field).sfixed64 = { gte: 10, lt: 5 }]; } ``` |
| in | [sfixed64](#sfixed64) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MySFixed64 { // value must be in list [1, 2, 3] repeated sfixed64 value = 1 (buf.validate.field).sfixed64 = { in: [1, 2, 3] }; } ``` |
| not_in | [sfixed64](#sfixed64) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MySFixed64 { // value must not be in list [1, 2, 3] repeated sfixed64 value = 1 (buf.validate.field).sfixed64 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-SInt32Rules"></a>

### SInt32Rules
SInt32Rules describes the constraints applied to `sint32` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sint32](#sint32) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MySInt32 { // value must equal 42 sint32 value = 1 [(buf.validate.field).sint32.const = 42]; } ``` |
| lt | [sint32](#sint32) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MySInt32 { // value must be less than 10 sint32 value = 1 [(buf.validate.field).sint32.lt = 10]; } ``` |
| lte | [sint32](#sint32) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MySInt32 { // value must be less than or equal to 10 sint32 value = 1 [(buf.validate.field).sint32.lte = 10]; } ``` |
| gt | [sint32](#sint32) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySInt32 { // value must be greater than 5 [sint32.gt] sint32 value = 1 [(buf.validate.field).sint32.gt = 5];

 // value must be greater than 5 and less than 10 [sint32.gt_lt] sint32 other_value = 2 [(buf.validate.field).sint32 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [sint32.gt_lt_exclusive] sint32 another_value = 3 [(buf.validate.field).sint32 = { gt: 10, lt: 5 }]; } ``` |
| gte | [sint32](#sint32) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySInt32 { // value must be greater than or equal to 5 [sint32.gte] sint32 value = 1 [(buf.validate.field).sint32.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [sint32.gte_lt] sint32 other_value = 2 [(buf.validate.field).sint32 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [sint32.gte_lt_exclusive] sint32 another_value = 3 [(buf.validate.field).sint32 = { gte: 10, lt: 5 }]; } ``` |
| in | [sint32](#sint32) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MySInt32 { // value must be in list [1, 2, 3] repeated sint32 value = 1 (buf.validate.field).sint32 = { in: [1, 2, 3] }; } ``` |
| not_in | [sint32](#sint32) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MySInt32 { // value must not be in list [1, 2, 3] repeated sint32 value = 1 (buf.validate.field).sint32 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-SInt64Rules"></a>

### SInt64Rules
SInt64Rules describes the constraints applied to `sint64` values.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sint64](#sint64) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MySInt64 { // value must equal 42 sint64 value = 1 [(buf.validate.field).sint64.const = 42]; } ``` |
| lt | [sint64](#sint64) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MySInt64 { // value must be less than 10 sint64 value = 1 [(buf.validate.field).sint64.lt = 10]; } ``` |
| lte | [sint64](#sint64) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MySInt64 { // value must be less than or equal to 10 sint64 value = 1 [(buf.validate.field).sint64.lte = 10]; } ``` |
| gt | [sint64](#sint64) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySInt64 { // value must be greater than 5 [sint64.gt] sint64 value = 1 [(buf.validate.field).sint64.gt = 5];

 // value must be greater than 5 and less than 10 [sint64.gt_lt] sint64 other_value = 2 [(buf.validate.field).sint64 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [sint64.gt_lt_exclusive] sint64 another_value = 3 [(buf.validate.field).sint64 = { gt: 10, lt: 5 }]; } ``` |
| gte | [sint64](#sint64) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MySInt64 { // value must be greater than or equal to 5 [sint64.gte] sint64 value = 1 [(buf.validate.field).sint64.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [sint64.gte_lt] sint64 other_value = 2 [(buf.validate.field).sint64 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [sint64.gte_lt_exclusive] sint64 another_value = 3 [(buf.validate.field).sint64 = { gte: 10, lt: 5 }]; } ``` |
| in | [sint64](#sint64) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MySInt64 { // value must be in list [1, 2, 3] repeated sint64 value = 1 (buf.validate.field).sint64 = { in: [1, 2, 3] }; } ``` |
| not_in | [sint64](#sint64) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MySInt64 { // value must not be in list [1, 2, 3] repeated sint64 value = 1 (buf.validate.field).sint64 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-StringRules"></a>

### StringRules
StringRules describes the constraints applied to `string` values These
rules may also be applied to the `google.protobuf.StringValue` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [string](#string) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyString { // value must equal `hello` string value = 1 [(buf.validate.field).string.const = &#34;hello&#34;]; } ``` |
| len | [uint64](#uint64) | optional | `len` dictates that the field value must have the specified number of characters (Unicode code points), which may differ from the number of bytes in the string. If the field value does not meet the specified length, an error message will be generated.

```proto message MyString { // value length must be 5 characters string value = 1 [(buf.validate.field).string.len = 5]; } ``` |
| min_len | [uint64](#uint64) | optional | `min_len` specifies that the field value must have at least the specified number of characters (Unicode code points), which may differ from the number of bytes in the string. If the field value contains fewer characters, an error message will be generated.

```proto message MyString { // value length must be at least 3 characters string value = 1 [(buf.validate.field).string.min_len = 3]; } ``` |
| max_len | [uint64](#uint64) | optional | `max_len` specifies that the field value must have no more than the specified number of characters (Unicode code points), which may differ from the number of bytes in the string. If the field value contains more characters, an error message will be generated.

```proto message MyString { // value length must be at most 10 characters string value = 1 [(buf.validate.field).string.max_len = 10]; } ``` |
| len_bytes | [uint64](#uint64) | optional | `len_bytes` dictates that the field value must have the specified number of bytes. If the field value does not match the specified length in bytes, an error message will be generated.

```proto message MyString { // value length must be 6 bytes string value = 1 [(buf.validate.field).string.len_bytes = 6]; } ``` |
| min_bytes | [uint64](#uint64) | optional | `min_bytes` specifies that the field value must have at least the specified number of bytes. If the field value contains fewer bytes, an error message will be generated.

```proto message MyString { // value length must be at least 4 bytes string value = 1 [(buf.validate.field).string.min_bytes = 4]; }

``` |
| max_bytes | [uint64](#uint64) | optional | `max_bytes` specifies that the field value must have no more than the specified number of bytes. If the field value contains more bytes, an error message will be generated.

```proto message MyString { // value length must be at most 8 bytes string value = 1 [(buf.validate.field).string.max_bytes = 8]; } ``` |
| pattern | [string](#string) | optional | `pattern` specifies that the field value must match the specified regular expression (RE2 syntax), with the expression provided without any delimiters. If the field value doesn&#39;t match the regular expression, an error message will be generated.

```proto message MyString { // value does not match regex pattern `^[a-zA-Z]//$` string value = 1 [(buf.validate.field).string.pattern = &#34;^[a-zA-Z]//$&#34;]; } ``` |
| prefix | [string](#string) | optional | `prefix` specifies that the field value must have the specified substring at the beginning of the string. If the field value doesn&#39;t start with the specified prefix, an error message will be generated.

```proto message MyString { // value does not have prefix `pre` string value = 1 [(buf.validate.field).string.prefix = &#34;pre&#34;]; } ``` |
| suffix | [string](#string) | optional | `suffix` specifies that the field value must have the specified substring at the end of the string. If the field value doesn&#39;t end with the specified suffix, an error message will be generated.

```proto message MyString { // value does not have suffix `post` string value = 1 [(buf.validate.field).string.suffix = &#34;post&#34;]; } ``` |
| contains | [string](#string) | optional | `contains` specifies that the field value must have the specified substring anywhere in the string. If the field value doesn&#39;t contain the specified substring, an error message will be generated.

```proto message MyString { // value does not contain substring `inside`. string value = 1 [(buf.validate.field).string.contains = &#34;inside&#34;]; } ``` |
| not_contains | [string](#string) | optional | `not_contains` specifies that the field value must not have the specified substring anywhere in the string. If the field value contains the specified substring, an error message will be generated.

```proto message MyString { // value contains substring `inside`. string value = 1 [(buf.validate.field).string.not_contains = &#34;inside&#34;]; } ``` |
| in | [string](#string) | repeated | `in` specifies that the field value must be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message will be generated.

```proto message MyString { // value must be in list [&#34;apple&#34;, &#34;banana&#34;] repeated string value = 1 [(buf.validate.field).string.in = &#34;apple&#34;, (buf.validate.field).string.in = &#34;banana&#34;]; } ``` |
| not_in | [string](#string) | repeated | `not_in` specifies that the field value cannot be equal to any of the specified values. If the field value is one of the specified values, an error message will be generated. ```proto message MyString { // value must not be in list [&#34;orange&#34;, &#34;grape&#34;] repeated string value = 1 [(buf.validate.field).string.not_in = &#34;orange&#34;, (buf.validate.field).string.not_in = &#34;grape&#34;]; } ``` |
| email | [bool](#bool) |  | `email` specifies that the field value must be a valid email address (addr-spec only) as defined by [RFC 5322](https://tools.ietf.org/html/rfc5322#section-3.4.1). If the field value isn&#39;t a valid email address, an error message will be generated.

```proto message MyString { // value must be a valid email address string value = 1 [(buf.validate.field).string.email = true]; } ``` |
| hostname | [bool](#bool) |  | `hostname` specifies that the field value must be a valid hostname as defined by [RFC 1034](https://tools.ietf.org/html/rfc1034#section-3.5). This constraint doesn&#39;t support internationalized domain names (IDNs). If the field value isn&#39;t a valid hostname, an error message will be generated.

```proto message MyString { // value must be a valid hostname string value = 1 [(buf.validate.field).string.hostname = true]; } ``` |
| ip | [bool](#bool) |  | `ip` specifies that the field value must be a valid IP (v4 or v6) address, without surrounding square brackets for IPv6 addresses. If the field value isn&#39;t a valid IP address, an error message will be generated.

```proto message MyString { // value must be a valid IP address string value = 1 [(buf.validate.field).string.ip = true]; } ``` |
| ipv4 | [bool](#bool) |  | `ipv4` specifies that the field value must be a valid IPv4 address. If the field value isn&#39;t a valid IPv4 address, an error message will be generated.

```proto message MyString { // value must be a valid IPv4 address string value = 1 [(buf.validate.field).string.ipv4 = true]; } ``` |
| ipv6 | [bool](#bool) |  | `ipv6` specifies that the field value must be a valid IPv6 address, without surrounding square brackets. If the field value is not a valid IPv6 address, an error message will be generated.

```proto message MyString { // value must be a valid IPv6 address string value = 1 [(buf.validate.field).string.ipv6 = true]; } ``` |
| uri | [bool](#bool) |  | `uri` specifies that the field value must be a valid, absolute URI as defined by [RFC 3986](https://tools.ietf.org/html/rfc3986#section-3). If the field value isn&#39;t a valid, absolute URI, an error message will be generated.

```proto message MyString { // value must be a valid URI string value = 1 [(buf.validate.field).string.uri = true]; } ``` |
| uri_ref | [bool](#bool) |  | `uri_ref` specifies that the field value must be a valid URI as defined by [RFC 3986](https://tools.ietf.org/html/rfc3986#section-3) and may be either relative or absolute. If the field value isn&#39;t a valid URI, an error message will be generated.

```proto message MyString { // value must be a valid URI string value = 1 [(buf.validate.field).string.uri_ref = true]; } ``` |
| address | [bool](#bool) |  | `address` specifies that the field value must be either a valid hostname as defined by [RFC 1034](https://tools.ietf.org/html/rfc1034#section-3.5) (which doesn&#39;t support internationalized domain names or IDNs) or a valid IP (v4 or v6). If the field value isn&#39;t a valid hostname or IP, an error message will be generated.

```proto message MyString { // value must be a valid hostname, or ip address string value = 1 [(buf.validate.field).string.address = true]; } ``` |
| uuid | [bool](#bool) |  | `uuid` specifies that the field value must be a valid UUID as defined by [RFC 4122](https://tools.ietf.org/html/rfc4122#section-4.1.2). If the field value isn&#39;t a valid UUID, an error message will be generated.

```proto message MyString { // value must be a valid UUID string value = 1 [(buf.validate.field).string.uuid = true]; } ``` |
| tuuid | [bool](#bool) |  | `tuuid` (trimmed UUID) specifies that the field value must be a valid UUID as defined by [RFC 4122](https://tools.ietf.org/html/rfc4122#section-4.1.2) with all dashes omitted. If the field value isn&#39;t a valid UUID without dashes, an error message will be generated.

```proto message MyString { // value must be a valid trimmed UUID string value = 1 [(buf.validate.field).string.tuuid = true]; } ``` |
| ip_with_prefixlen | [bool](#bool) |  | `ip_with_prefixlen` specifies that the field value must be a valid IP (v4 or v6) address with prefix length. If the field value isn&#39;t a valid IP with prefix length, an error message will be generated.

```proto message MyString { // value must be a valid IP with prefix length string value = 1 [(buf.validate.field).string.ip_with_prefixlen = true]; } ``` |
| ipv4_with_prefixlen | [bool](#bool) |  | `ipv4_with_prefixlen` specifies that the field value must be a valid IPv4 address with prefix. If the field value isn&#39;t a valid IPv4 address with prefix length, an error message will be generated.

```proto message MyString { // value must be a valid IPv4 address with prefix length string value = 1 [(buf.validate.field).string.ipv4_with_prefixlen = true]; } ``` |
| ipv6_with_prefixlen | [bool](#bool) |  | `ipv6_with_prefixlen` specifies that the field value must be a valid IPv6 address with prefix length. If the field value is not a valid IPv6 address with prefix length, an error message will be generated.

```proto message MyString { // value must be a valid IPv6 address prefix length string value = 1 [(buf.validate.field).string.ipv6_with_prefixlen = true]; } ``` |
| ip_prefix | [bool](#bool) |  | `ip_prefix` specifies that the field value must be a valid IP (v4 or v6) prefix. If the field value isn&#39;t a valid IP prefix, an error message will be generated. The prefix must have all zeros for the masked bits of the prefix (e.g., `127.0.0.0/16`, not `127.0.0.1/16`).

```proto message MyString { // value must be a valid IP prefix string value = 1 [(buf.validate.field).string.ip_prefix = true]; } ``` |
| ipv4_prefix | [bool](#bool) |  | `ipv4_prefix` specifies that the field value must be a valid IPv4 prefix. If the field value isn&#39;t a valid IPv4 prefix, an error message will be generated. The prefix must have all zeros for the masked bits of the prefix (e.g., `127.0.0.0/16`, not `127.0.0.1/16`).

```proto message MyString { // value must be a valid IPv4 prefix string value = 1 [(buf.validate.field).string.ipv4_prefix = true]; } ``` |
| ipv6_prefix | [bool](#bool) |  | `ipv6_prefix` specifies that the field value must be a valid IPv6 prefix. If the field value is not a valid IPv6 prefix, an error message will be generated. The prefix must have all zeros for the masked bits of the prefix (e.g., `2001:db8::/48`, not `2001:db8::1/48`).

```proto message MyString { // value must be a valid IPv6 prefix string value = 1 [(buf.validate.field).string.ipv6_prefix = true]; } ``` |
| host_and_port | [bool](#bool) |  | `host_and_port` specifies the field value must be a valid host and port pair. The host must be a valid hostname or IP address while the port must be in the range of 0-65535, inclusive. IPv6 addresses must be delimited with square brackets (e.g., `[::1]:1234`). |
| well_known_regex | [KnownRegex](#buf-validate-KnownRegex) |  | `well_known_regex` specifies a common well-known pattern defined as a regex. If the field value doesn&#39;t match the well-known regex, an error message will be generated.

```proto message MyString { // value must be a valid HTTP header value string value = 1 [(buf.validate.field).string.well_known_regex = KNOWN_REGEX_HTTP_HEADER_VALUE]; } ```

#### KnownRegex

`well_known_regex` contains some well-known patterns.

| Name | Number | Description | |-------------------------------|--------|-------------------------------------------| | KNOWN_REGEX_UNSPECIFIED | 0 | | | KNOWN_REGEX_HTTP_HEADER_NAME | 1 | HTTP header name as defined by [RFC 7230](https://tools.ietf.org/html/rfc7230#section-3.2) | | KNOWN_REGEX_HTTP_HEADER_VALUE | 2 | HTTP header value as defined by [RFC 7230](https://tools.ietf.org/html/rfc7230#section-3.2.4) | |
| strict | [bool](#bool) | optional | This applies to regexes `HTTP_HEADER_NAME` and `HTTP_HEADER_VALUE` to enable strict header validation. By default, this is true, and HTTP header validations are [RFC-compliant](https://tools.ietf.org/html/rfc7230#section-3). Setting to false will enable looser validations that only disallow `\r\n\0` characters, which can be used to bypass header matching rules.

```proto message MyString { // The field `value` must have be a valid HTTP headers, but not enforced with strict rules. string value = 1 [(buf.validate.field).string.strict = false]; } ``` |






<a name="buf-validate-TimestampRules"></a>

### TimestampRules
TimestampRules describe the constraints applied exclusively to the `google.protobuf.Timestamp` well-known type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional | `const` dictates that this field, of the `google.protobuf.Timestamp` type, must exactly match the specified value. If the field value doesn&#39;t correspond to the specified timestamp, an error message will be generated.

```proto message MyTimestamp { // value must equal 2023-05-03T10:00:00Z google.protobuf.Timestamp created_at = 1 [(buf.validate.field).timestamp.const = {seconds: 1727998800}]; } ``` |
| lt | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | requires the duration field value to be less than the specified value (field &lt; value). If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyDuration { // duration must be less than &#39;P3D&#39; [duration.lt] google.protobuf.Duration value = 1 [(buf.validate.field).duration.lt = { seconds: 259200 }]; } ``` |
| lte | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | requires the timestamp field value to be less than or equal to the specified value (field &lt;= value). If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyTimestamp { // timestamp must be less than or equal to &#39;2023-05-14T00:00:00Z&#39; [timestamp.lte] google.protobuf.Timestamp value = 1 [(buf.validate.field).timestamp.lte = { seconds: 1678867200 }]; } ``` |
| lt_now | [bool](#bool) |  | `lt_now` specifies that this field, of the `google.protobuf.Timestamp` type, must be less than the current time. `lt_now` can only be used with the `within` rule.

```proto message MyTimestamp { // value must be less than now google.protobuf.Timestamp created_at = 1 [(buf.validate.field).timestamp.lt_now = true]; } ``` |
| gt | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | `gt` requires the timestamp field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyTimestamp { // timestamp must be greater than &#39;2023-01-01T00:00:00Z&#39; [timestamp.gt] google.protobuf.Timestamp value = 1 [(buf.validate.field).timestamp.gt = { seconds: 1672444800 }];

 // timestamp must be greater than &#39;2023-01-01T00:00:00Z&#39; and less than &#39;2023-01-02T00:00:00Z&#39; [timestamp.gt_lt] google.protobuf.Timestamp another_value = 2 [(buf.validate.field).timestamp = { gt: { seconds: 1672444800 }, lt: { seconds: 1672531200 } }];

 // timestamp must be greater than &#39;2023-01-02T00:00:00Z&#39; or less than &#39;2023-01-01T00:00:00Z&#39; [timestamp.gt_lt_exclusive] google.protobuf.Timestamp other_value = 3 [(buf.validate.field).timestamp = { gt: { seconds: 1672531200 }, lt: { seconds: 1672444800 } }]; } ``` |
| gte | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | `gte` requires the timestamp field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyTimestamp { // timestamp must be greater than or equal to &#39;2023-01-01T00:00:00Z&#39; [timestamp.gte] google.protobuf.Timestamp value = 1 [(buf.validate.field).timestamp.gte = { seconds: 1672444800 }];

 // timestamp must be greater than or equal to &#39;2023-01-01T00:00:00Z&#39; and less than &#39;2023-01-02T00:00:00Z&#39; [timestamp.gte_lt] google.protobuf.Timestamp another_value = 2 [(buf.validate.field).timestamp = { gte: { seconds: 1672444800 }, lt: { seconds: 1672531200 } }];

 // timestamp must be greater than or equal to &#39;2023-01-02T00:00:00Z&#39; or less than &#39;2023-01-01T00:00:00Z&#39; [timestamp.gte_lt_exclusive] google.protobuf.Timestamp other_value = 3 [(buf.validate.field).timestamp = { gte: { seconds: 1672531200 }, lt: { seconds: 1672444800 } }]; } ``` |
| gt_now | [bool](#bool) |  | `gt_now` specifies that this field, of the `google.protobuf.Timestamp` type, must be greater than the current time. `gt_now` can only be used with the `within` rule.

```proto message MyTimestamp { // value must be greater than now google.protobuf.Timestamp created_at = 1 [(buf.validate.field).timestamp.gt_now = true]; } ``` |
| within | [google.protobuf.Duration](#google-protobuf-Duration) | optional | `within` specifies that this field, of the `google.protobuf.Timestamp` type, must be within the specified duration of the current time. If the field value isn&#39;t within the duration, an error message is generated.

```proto message MyTimestamp { // value must be within 1 hour of now google.protobuf.Timestamp created_at = 1 [(buf.validate.field).timestamp.within = {seconds: 3600}]; } ``` |






<a name="buf-validate-UInt32Rules"></a>

### UInt32Rules
UInt32Rules describes the constraints applied to `uint32` values. These
rules may also be applied to the `google.protobuf.UInt32Value` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [uint32](#uint32) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyUInt32 { // value must equal 42 uint32 value = 1 [(buf.validate.field).uint32.const = 42]; } ``` |
| lt | [uint32](#uint32) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyUInt32 { // value must be less than 10 uint32 value = 1 [(buf.validate.field).uint32.lt = 10]; } ``` |
| lte | [uint32](#uint32) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyUInt32 { // value must be less than or equal to 10 uint32 value = 1 [(buf.validate.field).uint32.lte = 10]; } ``` |
| gt | [uint32](#uint32) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyUInt32 { // value must be greater than 5 [uint32.gt] uint32 value = 1 [(buf.validate.field).uint32.gt = 5];

 // value must be greater than 5 and less than 10 [uint32.gt_lt] uint32 other_value = 2 [(buf.validate.field).uint32 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [uint32.gt_lt_exclusive] uint32 another_value = 3 [(buf.validate.field).uint32 = { gt: 10, lt: 5 }]; } ``` |
| gte | [uint32](#uint32) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyUInt32 { // value must be greater than or equal to 5 [uint32.gte] uint32 value = 1 [(buf.validate.field).uint32.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [uint32.gte_lt] uint32 other_value = 2 [(buf.validate.field).uint32 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [uint32.gte_lt_exclusive] uint32 another_value = 3 [(buf.validate.field).uint32 = { gte: 10, lt: 5 }]; } ``` |
| in | [uint32](#uint32) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyUInt32 { // value must be in list [1, 2, 3] repeated uint32 value = 1 (buf.validate.field).uint32 = { in: [1, 2, 3] }; } ``` |
| not_in | [uint32](#uint32) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyUInt32 { // value must not be in list [1, 2, 3] repeated uint32 value = 1 (buf.validate.field).uint32 = { not_in: [1, 2, 3] }; } ``` |






<a name="buf-validate-UInt64Rules"></a>

### UInt64Rules
UInt64Rules describes the constraints applied to `uint64` values. These
rules may also be applied to the `google.protobuf.UInt64Value` Well-Known-Type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [uint64](#uint64) | optional | `const` requires the field value to exactly match the specified value. If the field value doesn&#39;t match, an error message is generated.

```proto message MyUInt64 { // value must equal 42 uint64 value = 1 [(buf.validate.field).uint64.const = 42]; } ``` |
| lt | [uint64](#uint64) |  | `lt` requires the field value to be less than the specified value (field &lt; value). If the field value is equal to or greater than the specified value, an error message is generated.

```proto message MyUInt64 { // value must be less than 10 uint64 value = 1 [(buf.validate.field).uint64.lt = 10]; } ``` |
| lte | [uint64](#uint64) |  | `lte` requires the field value to be less than or equal to the specified value (field &lt;= value). If the field value is greater than the specified value, an error message is generated.

```proto message MyUInt64 { // value must be less than or equal to 10 uint64 value = 1 [(buf.validate.field).uint64.lte = 10]; } ``` |
| gt | [uint64](#uint64) |  | `gt` requires the field value to be greater than the specified value (exclusive). If the value of `gt` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyUInt64 { // value must be greater than 5 [uint64.gt] uint64 value = 1 [(buf.validate.field).uint64.gt = 5];

 // value must be greater than 5 and less than 10 [uint64.gt_lt] uint64 other_value = 2 [(buf.validate.field).uint64 = { gt: 5, lt: 10 }];

 // value must be greater than 10 or less than 5 [uint64.gt_lt_exclusive] uint64 another_value = 3 [(buf.validate.field).uint64 = { gt: 10, lt: 5 }]; } ``` |
| gte | [uint64](#uint64) |  | `gte` requires the field value to be greater than or equal to the specified value (exclusive). If the value of `gte` is larger than a specified `lt` or `lte`, the range is reversed, and the field value must be outside the specified range. If the field value doesn&#39;t meet the required conditions, an error message is generated.

```proto message MyUInt64 { // value must be greater than or equal to 5 [uint64.gte] uint64 value = 1 [(buf.validate.field).uint64.gte = 5];

 // value must be greater than or equal to 5 and less than 10 [uint64.gte_lt] uint64 other_value = 2 [(buf.validate.field).uint64 = { gte: 5, lt: 10 }];

 // value must be greater than or equal to 10 or less than 5 [uint64.gte_lt_exclusive] uint64 another_value = 3 [(buf.validate.field).uint64 = { gte: 10, lt: 5 }]; } ``` |
| in | [uint64](#uint64) | repeated | `in` requires the field value to be equal to one of the specified values. If the field value isn&#39;t one of the specified values, an error message is generated.

```proto message MyUInt64 { // value must be in list [1, 2, 3] repeated uint64 value = 1 (buf.validate.field).uint64 = { in: [1, 2, 3] }; } ``` |
| not_in | [uint64](#uint64) | repeated | `not_in` requires the field value to not be equal to any of the specified values. If the field value is one of the specified values, an error message is generated.

```proto message MyUInt64 { // value must not be in list [1, 2, 3] repeated uint64 value = 1 (buf.validate.field).uint64 = { not_in: [1, 2, 3] }; } ``` |





 


<a name="buf-validate-Ignore"></a>

### Ignore
Specifies how FieldConstraints.ignore behaves. See the documentation for
FieldConstraints.required for definitions of &#34;populated&#34; and &#34;nullable&#34;.

| Name | Number | Description |
| ---- | ------ | ----------- |
| IGNORE_UNSPECIFIED | 0 | Validation is only skipped if it&#39;s an unpopulated nullable fields.

```proto syntax=&#34;proto3&#34;;

message Request { // The uri rule applies to any value, including the empty string. string foo = 1 [ (buf.validate.field).string.uri = true ];

 // The uri rule only applies if the field is set, including if it&#39;s // set to the empty string. optional string bar = 2 [ (buf.validate.field).string.uri = true ];

 // The min_items rule always applies, even if the list is empty. repeated string baz = 3 [ (buf.validate.field).repeated.min_items = 3 ];

 // The custom CEL rule applies only if the field is set, including if // it&#39;s the &#34;zero&#34; value of that message. SomeMessage quux = 4 [ (buf.validate.field).cel = {/* ... */} ]; } ``` |
| IGNORE_IF_UNPOPULATED | 1 | Validation is skipped if the field is unpopulated. This rule is redundant if the field is already nullable. This value is equivalent behavior to the deprecated ignore_empty rule.

```proto syntax=&#34;proto3

message Request { // The uri rule applies only if the value is not the empty string. string foo = 1 [ (buf.validate.field).string.uri = true, (buf.validate.field).ignore = IGNORE_IF_UNPOPULATED ];

 // IGNORE_IF_UNPOPULATED is equivalent to IGNORE_UNSPECIFIED in this // case: the uri rule only applies if the field is set, including if // it&#39;s set to the empty string. optional string bar = 2 [ (buf.validate.field).string.uri = true, (buf.validate.field).ignore = IGNORE_IF_UNPOPULATED ];

 // The min_items rule only applies if the list has at least one item. repeated string baz = 3 [ (buf.validate.field).repeated.min_items = 3, (buf.validate.field).ignore = IGNORE_IF_UNPOPULATED ];

 // IGNORE_IF_UNPOPULATED is equivalent to IGNORE_UNSPECIFIED in this // case: the custom CEL rule applies only if the field is set, including // if it&#39;s the &#34;zero&#34; value of that message. SomeMessage quux = 4 [ (buf.validate.field).cel = {/* ... */}, (buf.validate.field).ignore = IGNORE_IF_UNPOPULATED ]; } ``` |
| IGNORE_IF_DEFAULT_VALUE | 2 | Validation is skipped if the field is unpopulated or if it is a nullable field populated with its default value. This is typically the zero or empty value, but proto2 scalars support custom defaults. For messages, the default is a non-null message with all its fields unpopulated.

```proto syntax=&#34;proto3

message Request { // IGNORE_IF_DEFAULT_VALUE is equivalent to IGNORE_IF_UNPOPULATED in // this case; the uri rule applies only if the value is not the empty // string. string foo = 1 [ (buf.validate.field).string.uri = true, (buf.validate.field).ignore = IGNORE_IF_DEFAULT_VALUE ];

 // The uri rule only applies if the field is set to a value other than // the empty string. optional string bar = 2 [ (buf.validate.field).string.uri = true, (buf.validate.field).ignore = IGNORE_IF_DEFAULT_VALUE ];

 // IGNORE_IF_DEFAULT_VALUE is equivalent to IGNORE_IF_UNPOPULATED in // this case; the min_items rule only applies if the list has at least // one item. repeated string baz = 3 [ (buf.validate.field).repeated.min_items = 3, (buf.validate.field).ignore = IGNORE_IF_DEFAULT_VALUE ];

 // The custom CEL rule only applies if the field is set to a value other // than an empty message (i.e., fields are unpopulated). SomeMessage quux = 4 [ (buf.validate.field).cel = {/* ... */}, (buf.validate.field).ignore = IGNORE_IF_DEFAULT_VALUE ]; } ```

This rule is affected by proto2 custom default values:

```proto syntax=&#34;proto2&#34;;

message Request { // The gt rule only applies if the field is set and it&#39;s value is not the default (i.e., not -42). The rule even applies if the field is set to zero since the default value differs. optional int32 value = 1 [ default = -42, (buf.validate.field).int32.gt = 0, (buf.validate.field).ignore = IGNORE_IF_DEFAULT_VALUE ]; } |
| IGNORE_ALWAYS | 3 | The validation rules of this field will be skipped and not evaluated. This is useful for situations that necessitate turning off the rules of a field containing a message that may not make sense in the current context, or to temporarily disable constraints during development.

```proto message MyMessage { // The field&#39;s rules will always be ignored, including any validation&#39;s // on value&#39;s fields. MyOtherMessage value = 1 [ (buf.validate.field).ignore = IGNORE_ALWAYS]; } ``` |
| IGNORE_EMPTY | 1 | Deprecated: Use IGNORE_IF_UNPOPULATED instead. TODO: Remove this value pre-v1. |
| IGNORE_DEFAULT | 2 | Deprecated: Use IGNORE_IF_DEFAULT_VALUE. TODO: Remove this value pre-v1. |



<a name="buf-validate-KnownRegex"></a>

### KnownRegex
WellKnownRegex contain some well-known patterns.

| Name | Number | Description |
| ---- | ------ | ----------- |
| KNOWN_REGEX_UNSPECIFIED | 0 |  |
| KNOWN_REGEX_HTTP_HEADER_NAME | 1 | HTTP header name as defined by [RFC 7230](https://tools.ietf.org/html/rfc7230#section-3.2). |
| KNOWN_REGEX_HTTP_HEADER_VALUE | 2 | HTTP header value as defined by [RFC 7230](https://tools.ietf.org/html/rfc7230#section-3.2.4). |


 


<a name="third_party_buf_validate_validate-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| field | FieldConstraints | .google.protobuf.FieldOptions | 1159 | Rules specify the validations to be performed on this field. By default, no validation is performed against a field. |
| message | MessageConstraints | .google.protobuf.MessageOptions | 1159 | Rules specify the validations to be performed on this message. By default, no validation is performed against a message. |
| oneof | OneofConstraints | .google.protobuf.OneofOptions | 1159 | Rules specify the validations to be performed on this oneof. By default, no validation is performed against a oneof. |

 

 



<a name="svc-web-dashboard_dashboard-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.web.dashboard/dashboard.proto


 

 

 


<a name="svc-web-dashboard-Dashboard"></a>

### Dashboard


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="svc-biz-org_org-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.org/org.proto



<a name="svc-biz-org-InitDBResp"></a>

### InitDBResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |





 

 

 


<a name="svc-biz-org-Org"></a>

### Org


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InitDB | [.google.protobuf.Empty](#google-protobuf-Empty) | [InitDBResp](#svc-biz-org-InitDBResp) | 初始化数据库 |

 



<a name="svc-infra-notifier_notifier-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.infra.notifier/notifier.proto



<a name="svc-infra-notifier-InitDBResp"></a>

### InitDBResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |





 

 

 


<a name="svc-infra-notifier-Notifier"></a>

### Notifier


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InitDB | [.google.protobuf.Empty](#google-protobuf-Empty) | [InitDBResp](#svc-infra-notifier-InitDBResp) | 初始化数据库 |

 



<a name="svc-web-streamer_streamer-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.web.streamer/streamer.proto


 

 

 


<a name="svc-web-streamer-Streamer"></a>

### Streamer


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="svc-biz-gift_gift-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.gift/gift.proto



<a name="svc-biz-gift-GiftAddReq"></a>

### GiftAddReq
添加礼物


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift | [GiftInfo](#svc-biz-gift-GiftInfo) |  |  |






<a name="svc-biz-gift-GiftAddResp"></a>

### GiftAddResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [string](#string) |  |  |






<a name="svc-biz-gift-GiftGetRecordReq"></a>

### GiftGetRecordReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pageno | [int64](#int64) |  | 第几页 |
| pagenum | [int64](#int64) |  | 每页几条数据 |
| room_id | [string](#string) |  | 房间id |
| live_id | [string](#string) |  | 直播id |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间 |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间 |






<a name="svc-biz-gift-GiftGetRecordResp"></a>

### GiftGetRecordResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [GiftSendRecord](#svc-biz-gift-GiftSendRecord) | repeated |  |






<a name="svc-biz-gift-GiftGetReq"></a>

### GiftGetReq
获取礼物


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [string](#string) |  |  |






<a name="svc-biz-gift-GiftGetResp"></a>

### GiftGetResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift | [GiftInfo](#svc-biz-gift-GiftInfo) |  |  |






<a name="svc-biz-gift-GiftInfo"></a>

### GiftInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [string](#string) |  | 礼物id |
| gift_name | [string](#string) |  | 礼物名称 |
| type | [GiftType](#svc-biz-gift-GiftType) |  | 礼物类型 |
| status | [GiftStatus](#svc-biz-gift-GiftStatus) |  | 礼物状态 |
| recommend | [GiftRecommend](#svc-biz-gift-GiftRecommend) |  | 礼物推荐状态 |
| price | [int32](#int32) |  | 价格 |
| desc | [string](#string) |  | 描述 |
| combo_timeout | [int32](#int32) |  | combo触发间隔时间 |
| combo_showtime | [int32](#int32) |  | combo效果展示时间 |
| prize | [string](#string) |  | 奖励(json字符串:{&#34;user_exp&#34;:100, &#34;anchor_exp&#34;:20}) |
| pack | [string](#string) |  | 批量包(json字符串:[{pack:&#34;20&#34;,desc:&#34;&#34;},{pack:&#34;99&#34;,desc:&#34;&#34;}]) |
| pic | [string](#string) |  | 图片资源(json字符串:{&#34;icon&#34;:&#34;&#34;, &#34;icon_gif&#34;:&#34;&#34;, &#34;chat_icon&#34;:&#34;&#34;, &#34;combo_bg&#34;:&#34;&#34;, &#34;combo_icon&#34;:&#34;&#34;}) |






<a name="svc-biz-gift-GiftSendRecord"></a>

### GiftSendRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [string](#string) |  | 礼物id |
| gift_name | [string](#string) |  | 礼物名 |
| num | [int32](#int32) |  | 数量 |
| total_price | [int32](#int32) |  | 总价 |
| from_uid | [string](#string) |  | 送礼uid |
| room_id | [string](#string) |  | 房间id |
| live_id | [string](#string) |  | 直播id |
| send_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |






<a name="svc-biz-gift-GiftSendRecordReq"></a>

### GiftSendRecordReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pageno | [int64](#int64) |  | 第几页 |
| pagenum | [int64](#int64) |  | 每页几条数据 |
| from_uid | [string](#string) |  | 用户id |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间 |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间 |






<a name="svc-biz-gift-GiftSendRecordResp"></a>

### GiftSendRecordResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [GiftSendRecord](#svc-biz-gift-GiftSendRecord) | repeated |  |






<a name="svc-biz-gift-GiftSendReq"></a>

### GiftSendReq
赠送礼物


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | 唯一标识 |
| order_id | [string](#string) |  | 支付订单id |
| gift_id | [string](#string) |  | 礼物id |
| num | [int32](#int32) |  | 数量 |
| prize | [int32](#int32) |  | 礼物单价（主要用做验证） |
| from_uid | [string](#string) |  | 赠送者uid |
| to_uid | [string](#string) |  | 接收者uid |
| room_id | [string](#string) |  | 房间id |
| live_id | [string](#string) |  | 直播id |






<a name="svc-biz-gift-GiftSendResp"></a>

### GiftSendResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sendret | [bool](#bool) |  | 赠送结果 |
| combo_num | [int32](#int32) |  | 连击数 |
| combo_group | [string](#string) |  | 连击组 |






<a name="svc-biz-gift-GiftUpdateReq"></a>

### GiftUpdateReq
更新礼物


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [string](#string) |  |  |
| gift | [GiftInfo](#svc-biz-gift-GiftInfo) |  |  |
| update_filed | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="svc-biz-gift-GiftUpdateResp"></a>

### GiftUpdateResp







<a name="svc-biz-gift-ListAdminReq"></a>

### ListAdminReq
礼物列表（后台使用）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| return_count | [bool](#bool) |  | 是否返回总数 |
| pageno | [int64](#int64) |  | 第几页 |
| pagenum | [int64](#int64) |  | 每页几条数据 |
| type | [GiftType](#svc-biz-gift-GiftType) |  | 礼物类型 |
| status | [GiftStatus](#svc-biz-gift-GiftStatus) |  | 礼物状态 |
| keyword | [string](#string) |  | 关键字 |






<a name="svc-biz-gift-ListAdminResp"></a>

### ListAdminResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [GiftInfo](#svc-biz-gift-GiftInfo) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="svc-biz-gift-ListOnlineAllReq"></a>

### ListOnlineAllReq







<a name="svc-biz-gift-ListOnlineByTypeReq"></a>

### ListOnlineByTypeReq
礼物列表（房间使用）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [GiftType](#svc-biz-gift-GiftType) |  | 礼物类型 |






<a name="svc-biz-gift-ListOnlineResp"></a>

### ListOnlineResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [GiftInfo](#svc-biz-gift-GiftInfo) | repeated |  |






<a name="svc-biz-gift-LiveStatReq"></a>

### LiveStatReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [string](#string) |  |  |
| live_id | [string](#string) |  |  |






<a name="svc-biz-gift-LiveStatResp"></a>

### LiveStatResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_cost | [int32](#int32) |  |  |





 


<a name="svc-biz-gift-GiftRecommend"></a>

### GiftRecommend


| Name | Number | Description |
| ---- | ------ | ----------- |
| GiftRecommendUnknown | 0 | 未知 |
| GiftRecommendYes | 1 | 推荐 |
| GiftRecommendNo | 2 | 未推荐 |



<a name="svc-biz-gift-GiftStatus"></a>

### GiftStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| GiftStatusUnknown | 0 | 未知 |
| GiftStatusOnline | 1 | 上线 |
| GiftStatusOffline | 2 | 下线 |



<a name="svc-biz-gift-GiftType"></a>

### GiftType


| Name | Number | Description |
| ---- | ------ | ----------- |
| GiftTypeUnknown | 0 | 未知 |
| GiftTypeCommon | 1 | 普通礼物 |
| GiftTypeFan | 2 | 粉丝礼物 |
| GiftTypeAristocrat | 3 | 贵族礼物 |


 

 


<a name="svc-biz-gift-Gift"></a>

### Gift


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Add | [GiftAddReq](#svc-biz-gift-GiftAddReq) | [GiftAddResp](#svc-biz-gift-GiftAddResp) | Add 添加礼物 |
| Get | [GiftGetReq](#svc-biz-gift-GiftGetReq) | [GiftGetResp](#svc-biz-gift-GiftGetResp) | Get 查询礼物 |
| Update | [GiftUpdateReq](#svc-biz-gift-GiftUpdateReq) | [GiftUpdateResp](#svc-biz-gift-GiftUpdateResp) | Update 更新礼物 |
| ListAdmin | [ListAdminReq](#svc-biz-gift-ListAdminReq) | [ListAdminResp](#svc-biz-gift-ListAdminResp) | ListAdmin 后台查询礼物列表接口 |
| ListOnlineByType | [ListOnlineByTypeReq](#svc-biz-gift-ListOnlineByTypeReq) | [ListOnlineResp](#svc-biz-gift-ListOnlineResp) | ListOnlineByType 前台房间礼物查询接口 |
| ListOnlineAll | [ListOnlineAllReq](#svc-biz-gift-ListOnlineAllReq) | [ListOnlineResp](#svc-biz-gift-ListOnlineResp) | ListOnlineAll 所有礼物的缓存接口 |
| Send | [GiftSendReq](#svc-biz-gift-GiftSendReq) | [GiftSendResp](#svc-biz-gift-GiftSendResp) | Send 送礼物接口 |
| SendRecord | [GiftSendRecordReq](#svc-biz-gift-GiftSendRecordReq) | [GiftSendRecordResp](#svc-biz-gift-GiftSendRecordResp) | SendRecord 送礼记录 |
| GetRecord | [GiftGetRecordReq](#svc-biz-gift-GiftGetRecordReq) | [GiftGetRecordResp](#svc-biz-gift-GiftGetRecordResp) | GetRecord 收礼记录 |
| LiveStat | [LiveStatReq](#svc-biz-gift-LiveStatReq) | [LiveStatResp](#svc-biz-gift-LiveStatResp) | LiveStat 直播统计 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

