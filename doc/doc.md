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
    - [BuyLuckyIdReq](#svc-biz-trade-BuyLuckyIdReq)
    - [BuyLuckyIdResp](#svc-biz-trade-BuyLuckyIdResp)
    - [BuyRideReq](#svc-biz-trade-BuyRideReq)
    - [BuyRideResp](#svc-biz-trade-BuyRideResp)
    - [GiftInfo](#svc-biz-trade-GiftInfo)
    - [JoinAnchorFansGroupReq](#svc-biz-trade-JoinAnchorFansGroupReq)
    - [JoinAnchorFansGroupResp](#svc-biz-trade-JoinAnchorFansGroupResp)
    - [LiveInfo](#svc-biz-trade-LiveInfo)
    - [MoneyExchangeCoinReq](#svc-biz-trade-MoneyExchangeCoinReq)
    - [MoneyExchangeCoinResp](#svc-biz-trade-MoneyExchangeCoinResp)
    - [MoneyRechargeReq](#svc-biz-trade-MoneyRechargeReq)
    - [MoneyRechargeReq.PayInfo](#svc-biz-trade-MoneyRechargeReq-PayInfo)
    - [MoneyRechargeResp](#svc-biz-trade-MoneyRechargeResp)
    - [MoneyWithdrawReq](#svc-biz-trade-MoneyWithdrawReq)
    - [MoneyWithdrawResp](#svc-biz-trade-MoneyWithdrawResp)
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
    - [VipActivateReq](#svc-biz-trade-VipActivateReq)
    - [VipActivateResp](#svc-biz-trade-VipActivateResp)
    - [VipExtendReq](#svc-biz-trade-VipExtendReq)
    - [VipExtendResp](#svc-biz-trade-VipExtendResp)
    - [VipUpgradeReq](#svc-biz-trade-VipUpgradeReq)
    - [VipUpgradeResp](#svc-biz-trade-VipUpgradeResp)
  
    - [Trade](#svc-biz-trade-Trade)
  
- [svc.infra.static/static.proto](#svc-infra-static_static-proto)
    - [InitDBResp](#svc-infra-static-InitDBResp)
    - [StreamRequestInfo](#svc-infra-static-StreamRequestInfo)
    - [UploadRequestMessage](#svc-infra-static-UploadRequestMessage)
    - [UploadResponseMessage](#svc-infra-static-UploadResponseMessage)
    - [UploadStreamRequestMessage](#svc-infra-static-UploadStreamRequestMessage)
  
    - [Static](#svc-infra-static-Static)
  
- [svc.biz.log/log.proto](#svc-biz-log_log-proto)
    - [AddLogReq](#svc-biz-log-AddLogReq)
    - [AddLogResp](#svc-biz-log-AddLogResp)
    - [LogInfo](#svc-biz-log-LogInfo)
    - [MGetLastLogReq](#svc-biz-log-MGetLastLogReq)
    - [MGetLastLogResp](#svc-biz-log-MGetLastLogResp)
    - [MGetLastLogResp.ItemsEntry](#svc-biz-log-MGetLastLogResp-ItemsEntry)
  
    - [Log](#svc-biz-log-Log)
  
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
    - [MGetCategoryReq](#svc-biz-room-MGetCategoryReq)
    - [MGetCategoryResp](#svc-biz-room-MGetCategoryResp)
    - [MGetCategoryResp.ItemsEntry](#svc-biz-room-MGetCategoryResp-ItemsEntry)
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
    - [ForbidRoomReq](#svc-biz-room-ForbidRoomReq)
    - [ForbidRoomResp](#svc-biz-room-ForbidRoomResp)
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
    - [ResumeRoomReq](#svc-biz-room-ResumeRoomReq)
    - [ResumeRoomResp](#svc-biz-room-ResumeRoomResp)
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
  
- [svc.biz.relation/relation.proto](#svc-biz-relation_relation-proto)
    - [GetRelationListReq](#svc-biz-relation-GetRelationListReq)
    - [GetRelationListResp](#svc-biz-relation-GetRelationListResp)
    - [RelationAddReq](#svc-biz-relation-RelationAddReq)
    - [RelationCheckReq](#svc-biz-relation-RelationCheckReq)
    - [RelationCheckResp](#svc-biz-relation-RelationCheckResp)
    - [RelationDelReq](#svc-biz-relation-RelationDelReq)
    - [RelationGetReq](#svc-biz-relation-RelationGetReq)
    - [RelationGetResp](#svc-biz-relation-RelationGetResp)
    - [RelationInfo](#svc-biz-relation-RelationInfo)
  
    - [RelationType](#svc-biz-relation-RelationType)
  
    - [Relation](#svc-biz-relation-Relation)
  
- [svc.web.dashboard/dashboard.proto](#svc-web-dashboard_dashboard-proto)
    - [Dashboard](#svc-web-dashboard-Dashboard)
  
- [svc.biz.org/org.proto](#svc-biz-org_org-proto)
    - [AddDepartmentReq](#svc-biz-org-AddDepartmentReq)
    - [AddDepartmentResp](#svc-biz-org-AddDepartmentResp)
    - [AddMerchantReq](#svc-biz-org-AddMerchantReq)
    - [AddMerchantResp](#svc-biz-org-AddMerchantResp)
    - [AddUnionReq](#svc-biz-org-AddUnionReq)
    - [AddUnionResp](#svc-biz-org-AddUnionResp)
    - [DeleteDepartmentReq](#svc-biz-org-DeleteDepartmentReq)
    - [DeleteDepartmentResp](#svc-biz-org-DeleteDepartmentResp)
    - [DeleteMerchantReq](#svc-biz-org-DeleteMerchantReq)
    - [DeleteMerchantResp](#svc-biz-org-DeleteMerchantResp)
    - [DeleteUnionReq](#svc-biz-org-DeleteUnionReq)
    - [DeleteUnionResp](#svc-biz-org-DeleteUnionResp)
    - [Department](#svc-biz-org-Department)
    - [GetDepartmentReq](#svc-biz-org-GetDepartmentReq)
    - [GetDepartmentResp](#svc-biz-org-GetDepartmentResp)
    - [GetMerchantReq](#svc-biz-org-GetMerchantReq)
    - [GetMerchantResp](#svc-biz-org-GetMerchantResp)
    - [GetUnionReq](#svc-biz-org-GetUnionReq)
    - [GetUnionResp](#svc-biz-org-GetUnionResp)
    - [InitDBResp](#svc-biz-org-InitDBResp)
    - [ListDepartmentsReq](#svc-biz-org-ListDepartmentsReq)
    - [ListDepartmentsResp](#svc-biz-org-ListDepartmentsResp)
    - [ListMerchantsReq](#svc-biz-org-ListMerchantsReq)
    - [ListMerchantsResp](#svc-biz-org-ListMerchantsResp)
    - [ListUnionsReq](#svc-biz-org-ListUnionsReq)
    - [ListUnionsResp](#svc-biz-org-ListUnionsResp)
    - [Merchant](#svc-biz-org-Merchant)
    - [Union](#svc-biz-org-Union)
    - [UpdateDepartmentReq](#svc-biz-org-UpdateDepartmentReq)
    - [UpdateDepartmentResp](#svc-biz-org-UpdateDepartmentResp)
    - [UpdateMerchantReq](#svc-biz-org-UpdateMerchantReq)
    - [UpdateMerchantResp](#svc-biz-org-UpdateMerchantResp)
    - [UpdateUnionReq](#svc-biz-org-UpdateUnionReq)
    - [UpdateUnionResp](#svc-biz-org-UpdateUnionResp)
  
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
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






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
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






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
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






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
| additions | [string](#string) |  | 扩展属性 |






<a name="svc-biz-account-Streamer"></a>

### Streamer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| display_id | [string](#string) |  | 用于显示和索引的数字ID |
| name | [string](#string) |  | 用户名 |
| nickname | [string](#string) |  | 昵称 |
| avatar | [string](#string) |  | 头像 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| password | [string](#string) | optional | 密码 |
| salt | [string](#string) | optional | 加密混淆 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| additions | [string](#string) |  | 扩展属性 |






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
| additions | [string](#string) |  | 扩展属性 |






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
| avatar | [string](#string) |  | 头像 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| device_ident | [string](#string) |  | 设备号 / 指纹 |
| password | [string](#string) | optional | 密码 |
| salt | [string](#string) | optional | 加密混淆 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| additions | [string](#string) |  | 扩展属性 |





 

 

 


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






<a name="svc-biz-trade-BuyLiveTicketResp"></a>

### BuyLiveTicketResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-BuyLuckyIdReq"></a>

### BuyLuckyIdReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |






<a name="svc-biz-trade-BuyLuckyIdResp"></a>

### BuyLuckyIdResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-BuyRideReq"></a>

### BuyRideReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |






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
| price | [int32](#int32) |  |  |
| num | [string](#string) |  |  |
| icon | [string](#string) |  |  |






<a name="svc-biz-trade-JoinAnchorFansGroupReq"></a>

### JoinAnchorFansGroupReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| live_info | [LiveInfo](#svc-biz-trade-LiveInfo) |  |  |






<a name="svc-biz-trade-JoinAnchorFansGroupResp"></a>

### JoinAnchorFansGroupResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






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






<a name="svc-biz-trade-MoneyExchangeCoinReq"></a>

### MoneyExchangeCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| user | [UserInfo](#svc-biz-trade-UserInfo) |  |  |






<a name="svc-biz-trade-MoneyExchangeCoinResp"></a>

### MoneyExchangeCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






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
| amount | [string](#string) |  | 币种金额 |






<a name="svc-biz-trade-MoneyRechargeResp"></a>

### MoneyRechargeResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






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
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-PayBulletChatReq"></a>

### PayBulletChatReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |
| live_info | [LiveInfo](#svc-biz-trade-LiveInfo) |  |  |






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






<a name="svc-biz-trade-VipActivateReq"></a>

### VipActivateReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |






<a name="svc-biz-trade-VipActivateResp"></a>

### VipActivateResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-VipExtendReq"></a>

### VipExtendReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |






<a name="svc-biz-trade-VipExtendResp"></a>

### VipExtendResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_result | [TradeResult](#svc-biz-trade-TradeResult) |  |  |






<a name="svc-biz-trade-VipUpgradeReq"></a>

### VipUpgradeReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trade_info | [TradeInfo](#svc-biz-trade-TradeInfo) |  |  |
| buyer | [UserInfo](#svc-biz-trade-UserInfo) |  |  |






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






<a name="svc-infra-static-StreamRequestInfo"></a>

### StreamRequestInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 为空将会自动生成 |
| bucket | [string](#string) |  | 自定义桶 |
| user_id | [string](#string) |  | 用户id |
| merchant_id | [string](#string) |  | 商户id |






<a name="svc-infra-static-UploadRequestMessage"></a>

### UploadRequestMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 为空将会自动生成 |
| binary | [bytes](#bytes) |  | 文件内容 |
| is_pre_mode | [bool](#bool) |  | 预写模式, 返回token客户端根据这个直接上传 |
| bucket | [string](#string) |  | 自定义桶 |
| user_id | [string](#string) |  | 用户id |
| merchant_id | [string](#string) |  | 商户id |






<a name="svc-infra-static-UploadResponseMessage"></a>

### UploadResponseMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | 返回文件路径 ，如果是预写模式该值就是token 或者地址 |
| domain | [string](#string) |  | 域 |
| provider | [string](#string) |  | oss 提供商 |






<a name="svc-infra-static-UploadStreamRequestMessage"></a>

### UploadStreamRequestMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| info | [StreamRequestInfo](#svc-infra-static-StreamRequestInfo) |  |  |
| binary | [bytes](#bytes) |  |  |





 

 

 


<a name="svc-infra-static-Static"></a>

### Static


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InitDB | [.google.protobuf.Empty](#google-protobuf-Empty) | [InitDBResp](#svc-infra-static-InitDBResp) | 初始化数据库 |
| UploadAvatar | [UploadRequestMessage](#svc-infra-static-UploadRequestMessage) | [UploadResponseMessage](#svc-infra-static-UploadResponseMessage) | 上传头像 |
| UploadCover | [UploadRequestMessage](#svc-infra-static-UploadRequestMessage) | [UploadResponseMessage](#svc-infra-static-UploadResponseMessage) | 上传封面 |
| UploadVideo | [UploadRequestMessage](#svc-infra-static-UploadRequestMessage) | [UploadResponseMessage](#svc-infra-static-UploadResponseMessage) | 上传视频 |
| UploadImage | [UploadRequestMessage](#svc-infra-static-UploadRequestMessage) | [UploadResponseMessage](#svc-infra-static-UploadResponseMessage) | 上传图片 |
| UploadFile | [UploadRequestMessage](#svc-infra-static-UploadRequestMessage) | [UploadResponseMessage](#svc-infra-static-UploadResponseMessage) | 上传文件 |
| UploadStreamFile | [UploadStreamRequestMessage](#svc-infra-static-UploadStreamRequestMessage) stream | [UploadResponseMessage](#svc-infra-static-UploadResponseMessage) | 流式上传文件 |

 



<a name="svc-biz-log_log-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.log/log.proto



<a name="svc-biz-log-AddLogReq"></a>

### AddLogReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| log_info | [LogInfo](#svc-biz-log-LogInfo) |  |  |






<a name="svc-biz-log-AddLogResp"></a>

### AddLogResp







<a name="svc-biz-log-LogInfo"></a>

### LogInfo
LogInfo 日志详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| log_id | [string](#string) |  |  |
| object | [string](#string) |  | 操作对象 |
| object_id | [string](#string) |  | 操作对象uuid |
| action | [string](#string) |  | 操作行为 |
| operator | [string](#string) |  | 操作人 |
| operate_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 操作时间 |
| extra | [string](#string) |  | 扩展信息,爱存啥存啥 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="svc-biz-log-MGetLastLogReq"></a>

### MGetLastLogReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [string](#string) |  | 操作对象 |
| object_ids | [string](#string) | repeated | 操作对象uuid |






<a name="svc-biz-log-MGetLastLogResp"></a>

### MGetLastLogResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MGetLastLogResp.ItemsEntry](#svc-biz-log-MGetLastLogResp-ItemsEntry) | repeated |  |






<a name="svc-biz-log-MGetLastLogResp-ItemsEntry"></a>

### MGetLastLogResp.ItemsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [LogInfo](#svc-biz-log-LogInfo) |  |  |





 

 

 


<a name="svc-biz-log-Log"></a>

### Log
分类

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddLog | [AddLogReq](#svc-biz-log-AddLogReq) | [AddLogResp](#svc-biz-log-AddLogResp) | AddLog 记录日志 |
| MGetLastLog | [MGetLastLogReq](#svc-biz-log-MGetLastLogReq) | [MGetLastLogResp](#svc-biz-log-MGetLastLogResp) | MGetLastLog 批量获取最近一次操作 |

 



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
| page | [int32](#int32) |  | 页数 |
| limit | [int32](#int32) |  | 条数 |
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






<a name="svc-biz-room-MGetCategoryReq"></a>

### MGetCategoryReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_ids | [string](#string) | repeated |  |






<a name="svc-biz-room-MGetCategoryResp"></a>

### MGetCategoryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MGetCategoryResp.ItemsEntry](#svc-biz-room-MGetCategoryResp-ItemsEntry) | repeated |  |






<a name="svc-biz-room-MGetCategoryResp-ItemsEntry"></a>

### MGetCategoryResp.ItemsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [CategoryInfo](#svc-biz-room-CategoryInfo) |  |  |






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
| MGetCategory | [MGetCategoryReq](#svc-biz-room-MGetCategoryReq) | [MGetCategoryResp](#svc-biz-room-MGetCategoryResp) | 获取分类 |
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






<a name="svc-biz-room-ForbidRoomReq"></a>

### ForbidRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  | 主播id |
| forbid_expire | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 封禁过期时间 |
| reason | [string](#string) |  | 封禁原因 |






<a name="svc-biz-room-ForbidRoomResp"></a>

### ForbidRoomResp







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






<a name="svc-biz-room-ResumeRoomReq"></a>

### ResumeRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  | 主播id |






<a name="svc-biz-room-ResumeRoomResp"></a>

### ResumeRoomResp







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
| forbid_expire | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 封禁到期时间 |
| forbid_reason | [string](#string) |  | 封禁原因 |
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
| ForbidRoom | [ForbidRoomReq](#svc-biz-room-ForbidRoomReq) | [ForbidRoomResp](#svc-biz-room-ForbidRoomResp) | ForbidRoom 封禁直播间 |
| ResumeRoom | [ResumeRoomReq](#svc-biz-room-ResumeRoomReq) | [ResumeRoomResp](#svc-biz-room-ResumeRoomResp) | ResumeRoom 解封直播间 |
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
| status | [string](#string) |  | 状态 |
| manager_id | [string](#string) |  | 操作账号 |
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
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-ChangeStreamerCoinResp"></a>

### ChangeStreamerCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  |  |
| new_value | [StreamerCoinValue](#svc-biz-asset-StreamerCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






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
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






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
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






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

 



<a name="svc-biz-relation_relation-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.relation/relation.proto



<a name="svc-biz-relation-GetRelationListReq"></a>

### GetRelationListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation_type | [RelationType](#svc-biz-relation-RelationType) |  |  |
| member_id | [string](#string) |  |  |
| page | [int64](#int64) |  |  |
| limit | [int64](#int64) |  |  |






<a name="svc-biz-relation-GetRelationListResp"></a>

### GetRelationListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [RelationInfo](#svc-biz-relation-RelationInfo) | repeated |  |






<a name="svc-biz-relation-RelationAddReq"></a>

### RelationAddReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation_info | [RelationInfo](#svc-biz-relation-RelationInfo) |  |  |






<a name="svc-biz-relation-RelationCheckReq"></a>

### RelationCheckReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation_type | [RelationType](#svc-biz-relation-RelationType) |  |  |
| member_id | [string](#string) |  | 成员 |
| r_member_id | [string](#string) |  | 产生关系的成员 |






<a name="svc-biz-relation-RelationCheckResp"></a>

### RelationCheckResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |






<a name="svc-biz-relation-RelationDelReq"></a>

### RelationDelReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation_type | [RelationType](#svc-biz-relation-RelationType) |  |  |
| member_id | [string](#string) |  | 成员 |
| r_member_id | [string](#string) |  | 产生关系的成员 |






<a name="svc-biz-relation-RelationGetReq"></a>

### RelationGetReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation_type | [RelationType](#svc-biz-relation-RelationType) |  |  |
| member_id | [string](#string) |  | 成员 |
| r_member_id | [string](#string) |  | 产生关系的成员 |






<a name="svc-biz-relation-RelationGetResp"></a>

### RelationGetResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation_info | [RelationInfo](#svc-biz-relation-RelationInfo) |  |  |






<a name="svc-biz-relation-RelationInfo"></a>

### RelationInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| relation_type | [RelationType](#svc-biz-relation-RelationType) |  | 关系类型 |
| member_id | [string](#string) |  | 成员（名单属于谁） |
| r_member_id | [string](#string) |  | 产生关系的成员（名单内有谁；当为ip或者设备号时不是uuid格式） |
| build_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 建立关系时间 |
| expire_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 过期时间（可无，为空则永久有效） |
| operator_id | [string](#string) |  | 操作人 |
| remark | [string](#string) |  | 备注 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |





 


<a name="svc-biz-relation-RelationType"></a>

### RelationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| RelationTypeUnknown | 0 | 未知 |
| RelationTypeFollow | 1 | 关注主播 |
| RelationTypeHistory | 2 | 观看历史 |
| RelationTypeMuzzle | 3 | 禁言 |
| RelationTypeBlacklistIP | 4 | ip黑名单 |
| RelationTypeBlacklistDevice | 5 | 设备黑名单 |


 

 


<a name="svc-biz-relation-Relation"></a>

### Relation


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RelationAdd | [RelationAddReq](#svc-biz-relation-RelationAddReq) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| RelationGet | [RelationGetReq](#svc-biz-relation-RelationGetReq) | [RelationGetResp](#svc-biz-relation-RelationGetResp) |  |
| RelationDel | [RelationDelReq](#svc-biz-relation-RelationDelReq) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| RelationCheck | [RelationCheckReq](#svc-biz-relation-RelationCheckReq) | [RelationCheckResp](#svc-biz-relation-RelationCheckResp) |  |
| GetRelationList | [GetRelationListReq](#svc-biz-relation-GetRelationListReq) | [GetRelationListResp](#svc-biz-relation-GetRelationListResp) |  |

 



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



<a name="svc-biz-org-AddDepartmentReq"></a>

### AddDepartmentReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| department | [Department](#svc-biz-org-Department) |  |  |






<a name="svc-biz-org-AddDepartmentResp"></a>

### AddDepartmentResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| department | [Department](#svc-biz-org-Department) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-biz-org-AddMerchantReq"></a>

### AddMerchantReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant | [Merchant](#svc-biz-org-Merchant) |  |  |






<a name="svc-biz-org-AddMerchantResp"></a>

### AddMerchantResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant | [Merchant](#svc-biz-org-Merchant) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-biz-org-AddUnionReq"></a>

### AddUnionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union | [Union](#svc-biz-org-Union) |  |  |






<a name="svc-biz-org-AddUnionResp"></a>

### AddUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union | [Union](#svc-biz-org-Union) |  |  |
| result | [bool](#bool) |  |  |






<a name="svc-biz-org-DeleteDepartmentReq"></a>

### DeleteDepartmentReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Department](#svc-biz-org-Department) |  |  |






<a name="svc-biz-org-DeleteDepartmentResp"></a>

### DeleteDepartmentResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-biz-org-DeleteMerchantReq"></a>

### DeleteMerchantReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Merchant](#svc-biz-org-Merchant) |  |  |






<a name="svc-biz-org-DeleteMerchantResp"></a>

### DeleteMerchantResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-biz-org-DeleteUnionReq"></a>

### DeleteUnionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-org-Union) |  |  |






<a name="svc-biz-org-DeleteUnionResp"></a>

### DeleteUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| deleted | [int64](#int64) |  |  |






<a name="svc-biz-org-Department"></a>

### Department
Models


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| name | [string](#string) |  | 名字 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| additions | [string](#string) |  | 扩展属性 |






<a name="svc-biz-org-GetDepartmentReq"></a>

### GetDepartmentReq
{{{ [Department]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Department](#svc-biz-org-Department) |  |  |






<a name="svc-biz-org-GetDepartmentResp"></a>

### GetDepartmentResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| department | [Department](#svc-biz-org-Department) |  |  |






<a name="svc-biz-org-GetMerchantReq"></a>

### GetMerchantReq
{{{ [Merchant]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Merchant](#svc-biz-org-Merchant) |  |  |






<a name="svc-biz-org-GetMerchantResp"></a>

### GetMerchantResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchant | [Merchant](#svc-biz-org-Merchant) |  |  |






<a name="svc-biz-org-GetUnionReq"></a>

### GetUnionReq
{{{ [Union]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-org-Union) |  |  |






<a name="svc-biz-org-GetUnionResp"></a>

### GetUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| union | [Union](#svc-biz-org-Union) |  |  |






<a name="svc-biz-org-InitDBResp"></a>

### InitDBResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [bool](#bool) |  |  |






<a name="svc-biz-org-ListDepartmentsReq"></a>

### ListDepartmentsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Department](#svc-biz-org-Department) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="svc-biz-org-ListDepartmentsResp"></a>

### ListDepartmentsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| departments | [Department](#svc-biz-org-Department) | repeated |  |






<a name="svc-biz-org-ListMerchantsReq"></a>

### ListMerchantsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Merchant](#svc-biz-org-Merchant) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="svc-biz-org-ListMerchantsResp"></a>

### ListMerchantsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| merchants | [Merchant](#svc-biz-org-Merchant) | repeated |  |






<a name="svc-biz-org-ListUnionsReq"></a>

### ListUnionsReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-org-Union) |  |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="svc-biz-org-ListUnionsResp"></a>

### ListUnionsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unions | [Union](#svc-biz-org-Union) | repeated |  |






<a name="svc-biz-org-Merchant"></a>

### Merchant



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| name | [string](#string) |  | 名字 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| additions | [string](#string) |  | 扩展属性 |






<a name="svc-biz-org-Union"></a>

### Union



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| name | [string](#string) |  | 名字 |
| merchant_id | [string](#string) |  | 商户ID |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| additions | [string](#string) |  | 扩展属性 |






<a name="svc-biz-org-UpdateDepartmentReq"></a>

### UpdateDepartmentReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Department](#svc-biz-org-Department) |  |  |






<a name="svc-biz-org-UpdateDepartmentResp"></a>

### UpdateDepartmentResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |






<a name="svc-biz-org-UpdateMerchantReq"></a>

### UpdateMerchantReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Merchant](#svc-biz-org-Merchant) |  |  |






<a name="svc-biz-org-UpdateMerchantResp"></a>

### UpdateMerchantResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |






<a name="svc-biz-org-UpdateUnionReq"></a>

### UpdateUnionReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Union](#svc-biz-org-Union) |  |  |






<a name="svc-biz-org-UpdateUnionResp"></a>

### UpdateUnionResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated | [int64](#int64) |  |  |





 

 

 


<a name="svc-biz-org-Org"></a>

### Org


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InitDB | [.google.protobuf.Empty](#google-protobuf-Empty) | [InitDBResp](#svc-biz-org-InitDBResp) | 初始化数据库 |
| GetDepartment | [GetDepartmentReq](#svc-biz-org-GetDepartmentReq) | [GetDepartmentResp](#svc-biz-org-GetDepartmentResp) | 获取部门 |
| ListDepartments | [ListDepartmentsReq](#svc-biz-org-ListDepartmentsReq) | [ListDepartmentsResp](#svc-biz-org-ListDepartmentsResp) | 获取部门列表 |
| AddDepartment | [AddDepartmentReq](#svc-biz-org-AddDepartmentReq) | [AddDepartmentResp](#svc-biz-org-AddDepartmentResp) | 添加部门 |
| UpdateDepartment | [UpdateDepartmentReq](#svc-biz-org-UpdateDepartmentReq) | [UpdateDepartmentResp](#svc-biz-org-UpdateDepartmentResp) | 更新部门 |
| DeleteDepartment | [DeleteDepartmentReq](#svc-biz-org-DeleteDepartmentReq) | [DeleteDepartmentResp](#svc-biz-org-DeleteDepartmentResp) | 删除部门 |
| GetMerchant | [GetMerchantReq](#svc-biz-org-GetMerchantReq) | [GetMerchantResp](#svc-biz-org-GetMerchantResp) | 获取商户 |
| ListMerchants | [ListMerchantsReq](#svc-biz-org-ListMerchantsReq) | [ListMerchantsResp](#svc-biz-org-ListMerchantsResp) | 获取商户列表 |
| AddMerchant | [AddMerchantReq](#svc-biz-org-AddMerchantReq) | [AddMerchantResp](#svc-biz-org-AddMerchantResp) | 添加商户 |
| UpdateMerchant | [UpdateMerchantReq](#svc-biz-org-UpdateMerchantReq) | [UpdateMerchantResp](#svc-biz-org-UpdateMerchantResp) | 更新商户 |
| DeleteMerchant | [DeleteMerchantReq](#svc-biz-org-DeleteMerchantReq) | [DeleteMerchantResp](#svc-biz-org-DeleteMerchantResp) | 删除商户 |
| GetUnion | [GetUnionReq](#svc-biz-org-GetUnionReq) | [GetUnionResp](#svc-biz-org-GetUnionResp) | 获取工会 |
| ListUnions | [ListUnionsReq](#svc-biz-org-ListUnionsReq) | [ListUnionsResp](#svc-biz-org-ListUnionsResp) | 获取工会列表 |
| AddUnion | [AddUnionReq](#svc-biz-org-AddUnionReq) | [AddUnionResp](#svc-biz-org-AddUnionResp) | 添加工会 |
| UpdateUnion | [UpdateUnionReq](#svc-biz-org-UpdateUnionReq) | [UpdateUnionResp](#svc-biz-org-UpdateUnionResp) | 更新工会 |
| DeleteUnion | [DeleteUnionReq](#svc-biz-org-DeleteUnionReq) | [DeleteUnionResp](#svc-biz-org-DeleteUnionResp) | 删除工会 |

 



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

