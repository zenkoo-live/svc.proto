# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [svc.biz.account/account.proto](#svc-biz-account_account-proto)
    - [Account](#svc-biz-account-Account)
  
- [svc.infra.setting/setting.proto](#svc-infra-setting_setting-proto)
    - [SettingGreetingReq](#svc-infra-setting-SettingGreetingReq)
    - [SettingGreetingResp](#svc-infra-setting-SettingGreetingResp)
  
    - [Setting](#svc-infra-setting-Setting)
  
- [svc.web.viewer/viewer.proto](#svc-web-viewer_viewer-proto)
    - [Viewer](#svc-web-viewer-Viewer)
  
- [svc.biz.asset/asset.proto](#svc-biz-asset_asset-proto)
    - [AnchorCoinDetail](#svc-biz-asset-AnchorCoinDetail)
    - [AnchorCoinValue](#svc-biz-asset-AnchorCoinValue)
    - [DecrAnchorCoinReq](#svc-biz-asset-DecrAnchorCoinReq)
    - [DecrAnchorCoinResp](#svc-biz-asset-DecrAnchorCoinResp)
    - [DecrMerchantCoinReq](#svc-biz-asset-DecrMerchantCoinReq)
    - [DecrMerchantCoinResp](#svc-biz-asset-DecrMerchantCoinResp)
    - [DecrMerchantMoneyReq](#svc-biz-asset-DecrMerchantMoneyReq)
    - [DecrMerchantMoneyResp](#svc-biz-asset-DecrMerchantMoneyResp)
    - [DecrUnionCoinReq](#svc-biz-asset-DecrUnionCoinReq)
    - [DecrUnionCoinResp](#svc-biz-asset-DecrUnionCoinResp)
    - [DecrUserCoinReq](#svc-biz-asset-DecrUserCoinReq)
    - [DecrUserCoinResp](#svc-biz-asset-DecrUserCoinResp)
    - [DecrUserMoneyReq](#svc-biz-asset-DecrUserMoneyReq)
    - [DecrUserMoneyResp](#svc-biz-asset-DecrUserMoneyResp)
    - [GetAnchorCoinMultiReq](#svc-biz-asset-GetAnchorCoinMultiReq)
    - [GetAnchorCoinMultiResp](#svc-biz-asset-GetAnchorCoinMultiResp)
    - [GetAnchorCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetAnchorCoinMultiResp-ValueMapEntry)
    - [GetAnchorCoinReq](#svc-biz-asset-GetAnchorCoinReq)
    - [GetAnchorCoinResp](#svc-biz-asset-GetAnchorCoinResp)
    - [GetMerchantCoinMultiReq](#svc-biz-asset-GetMerchantCoinMultiReq)
    - [GetMerchantCoinMultiResp](#svc-biz-asset-GetMerchantCoinMultiResp)
    - [GetMerchantCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetMerchantCoinMultiResp-ValueMapEntry)
    - [GetMerchantCoinReq](#svc-biz-asset-GetMerchantCoinReq)
    - [GetMerchantCoinResp](#svc-biz-asset-GetMerchantCoinResp)
    - [GetMerchantMoneyMultiReq](#svc-biz-asset-GetMerchantMoneyMultiReq)
    - [GetMerchantMoneyMultiResp](#svc-biz-asset-GetMerchantMoneyMultiResp)
    - [GetMerchantMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetMerchantMoneyMultiResp-ValueMapEntry)
    - [GetMerchantMoneyReq](#svc-biz-asset-GetMerchantMoneyReq)
    - [GetMerchantMoneyResp](#svc-biz-asset-GetMerchantMoneyResp)
    - [GetUnionCoinMultiReq](#svc-biz-asset-GetUnionCoinMultiReq)
    - [GetUnionCoinMultiResp](#svc-biz-asset-GetUnionCoinMultiResp)
    - [GetUnionCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetUnionCoinMultiResp-ValueMapEntry)
    - [GetUnionCoinReq](#svc-biz-asset-GetUnionCoinReq)
    - [GetUnionCoinResp](#svc-biz-asset-GetUnionCoinResp)
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
    - [IncrAnchorCoinReq](#svc-biz-asset-IncrAnchorCoinReq)
    - [IncrAnchorCoinResp](#svc-biz-asset-IncrAnchorCoinResp)
    - [IncrMerchantCoinReq](#svc-biz-asset-IncrMerchantCoinReq)
    - [IncrMerchantCoinResp](#svc-biz-asset-IncrMerchantCoinResp)
    - [IncrMerchantMoneyReq](#svc-biz-asset-IncrMerchantMoneyReq)
    - [IncrMerchantMoneyResp](#svc-biz-asset-IncrMerchantMoneyResp)
    - [IncrUnionCoinReq](#svc-biz-asset-IncrUnionCoinReq)
    - [IncrUnionCoinResp](#svc-biz-asset-IncrUnionCoinResp)
    - [IncrUserCoinReq](#svc-biz-asset-IncrUserCoinReq)
    - [IncrUserCoinResp](#svc-biz-asset-IncrUserCoinResp)
    - [IncrUserMoneyReq](#svc-biz-asset-IncrUserMoneyReq)
    - [IncrUserMoneyResp](#svc-biz-asset-IncrUserMoneyResp)
    - [ListAnchorCoinDetailReq](#svc-biz-asset-ListAnchorCoinDetailReq)
    - [ListAnchorCoinDetailResp](#svc-biz-asset-ListAnchorCoinDetailResp)
    - [ListMerchantCoinDetailReq](#svc-biz-asset-ListMerchantCoinDetailReq)
    - [ListMerchantCoinDetailResp](#svc-biz-asset-ListMerchantCoinDetailResp)
    - [ListMerchantMoneyDetailReq](#svc-biz-asset-ListMerchantMoneyDetailReq)
    - [ListMerchantMoneyDetailResp](#svc-biz-asset-ListMerchantMoneyDetailResp)
    - [ListUnionCoinDetailReq](#svc-biz-asset-ListUnionCoinDetailReq)
    - [ListUnionCoinDetailResp](#svc-biz-asset-ListUnionCoinDetailResp)
    - [ListUserCoinDetailReq](#svc-biz-asset-ListUserCoinDetailReq)
    - [ListUserCoinDetailResp](#svc-biz-asset-ListUserCoinDetailResp)
    - [ListUserMoneyDetailReq](#svc-biz-asset-ListUserMoneyDetailReq)
    - [ListUserMoneyDetailResp](#svc-biz-asset-ListUserMoneyDetailResp)
    - [MerchantCoinDetail](#svc-biz-asset-MerchantCoinDetail)
    - [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue)
    - [MerchantMoneyDetail](#svc-biz-asset-MerchantMoneyDetail)
    - [UnionCoinDetail](#svc-biz-asset-UnionCoinDetail)
    - [UnionCoinValue](#svc-biz-asset-UnionCoinValue)
    - [UserCoinDetail](#svc-biz-asset-UserCoinDetail)
    - [UserCoinValue](#svc-biz-asset-UserCoinValue)
    - [UserMoneyDetail](#svc-biz-asset-UserMoneyDetail)
  
    - [Asset](#svc-biz-asset-Asset)
  
- [svc.web.dashboard/dashboard.proto](#svc-web-dashboard_dashboard-proto)
    - [Dashboard](#svc-web-dashboard-Dashboard)
  
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
    - [GiftListAdminReq](#svc-biz-gift-GiftListAdminReq)
    - [GiftListAdminResp](#svc-biz-gift-GiftListAdminResp)
    - [GiftListReq](#svc-biz-gift-GiftListReq)
    - [GiftListResp](#svc-biz-gift-GiftListResp)
    - [GiftSendRecord](#svc-biz-gift-GiftSendRecord)
    - [GiftSendRecordReq](#svc-biz-gift-GiftSendRecordReq)
    - [GiftSendRecordResp](#svc-biz-gift-GiftSendRecordResp)
    - [GiftSendReq](#svc-biz-gift-GiftSendReq)
    - [GiftSendResp](#svc-biz-gift-GiftSendResp)
    - [GiftUpdateReq](#svc-biz-gift-GiftUpdateReq)
    - [GiftUpdateResp](#svc-biz-gift-GiftUpdateResp)
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


 

 

 


<a name="svc-biz-account-Account"></a>

### Account


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="svc-infra-setting_setting-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.infra.setting/setting.proto



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





 

 

 


<a name="svc-infra-setting-Setting"></a>

### Setting


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
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



<a name="svc-biz-asset-AnchorCoinDetail"></a>

### AnchorCoinDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value | [int64](#int64) |  |  |
| value_before | [int64](#int64) |  |  |
| value_after | [int64](#int64) |  |  |
| limited_value | [int64](#int64) |  |  |
| limited_value_before | [int64](#int64) |  |  |
| limited_value_after | [int64](#int64) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-AnchorCoinValue"></a>

### AnchorCoinValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-DecrAnchorCoinReq"></a>

### DecrAnchorCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-DecrAnchorCoinResp"></a>

### DecrAnchorCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [AnchorCoinValue](#svc-biz-asset-AnchorCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrMerchantCoinReq"></a>

### DecrMerchantCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-DecrMerchantCoinResp"></a>

### DecrMerchantCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrMerchantMoneyReq"></a>

### DecrMerchantMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-DecrMerchantMoneyResp"></a>

### DecrMerchantMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| new_value | [int64](#int64) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrUnionCoinReq"></a>

### DecrUnionCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-DecrUnionCoinResp"></a>

### DecrUnionCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrUserCoinReq"></a>

### DecrUserCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-DecrUserCoinResp"></a>

### DecrUserCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [UserCoinValue](#svc-biz-asset-UserCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-DecrUserMoneyReq"></a>

### DecrUserMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-DecrUserMoneyResp"></a>

### DecrUserMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| new_value | [int64](#int64) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-GetAnchorCoinMultiReq"></a>

### GetAnchorCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [string](#string) | repeated |  |
| app_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetAnchorCoinMultiResp"></a>

### GetAnchorCoinMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetAnchorCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetAnchorCoinMultiResp-ValueMapEntry) | repeated | uid-&gt;value |






<a name="svc-biz-asset-GetAnchorCoinMultiResp-ValueMapEntry"></a>

### GetAnchorCoinMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [AnchorCoinValue](#svc-biz-asset-AnchorCoinValue) |  |  |






<a name="svc-biz-asset-GetAnchorCoinReq"></a>

### GetAnchorCoinReq
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetAnchorCoinResp"></a>

### GetAnchorCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-GetMerchantCoinMultiReq"></a>

### GetMerchantCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [string](#string) | repeated |  |
| app_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetMerchantCoinMultiResp"></a>

### GetMerchantCoinMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetMerchantCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetMerchantCoinMultiResp-ValueMapEntry) | repeated | uid-&gt;value |






<a name="svc-biz-asset-GetMerchantCoinMultiResp-ValueMapEntry"></a>

### GetMerchantCoinMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  |  |






<a name="svc-biz-asset-GetMerchantCoinReq"></a>

### GetMerchantCoinReq
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetMerchantCoinResp"></a>

### GetMerchantCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-GetMerchantMoneyMultiReq"></a>

### GetMerchantMoneyMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetMerchantMoneyMultiResp"></a>

### GetMerchantMoneyMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetMerchantMoneyMultiResp.ValueMapEntry](#svc-biz-asset-GetMerchantMoneyMultiResp-ValueMapEntry) | repeated | uid-&gt;value |






<a name="svc-biz-asset-GetMerchantMoneyMultiResp-ValueMapEntry"></a>

### GetMerchantMoneyMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-GetMerchantMoneyReq"></a>

### GetMerchantMoneyReq
余额 money-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |






<a name="svc-biz-asset-GetMerchantMoneyResp"></a>

### GetMerchantMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-GetUnionCoinMultiReq"></a>

### GetUnionCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [string](#string) | repeated |  |
| app_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetUnionCoinMultiResp"></a>

### GetUnionCoinMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value_map | [GetUnionCoinMultiResp.ValueMapEntry](#svc-biz-asset-GetUnionCoinMultiResp-ValueMapEntry) | repeated | uid-&gt;value |






<a name="svc-biz-asset-GetUnionCoinMultiResp-ValueMapEntry"></a>

### GetUnionCoinMultiResp.ValueMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  |  |






<a name="svc-biz-asset-GetUnionCoinReq"></a>

### GetUnionCoinReq
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetUnionCoinResp"></a>

### GetUnionCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-GetUserCoinMultiReq"></a>

### GetUserCoinMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [string](#string) | repeated |  |
| app_id | [string](#string) |  |  |






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
虚拟币 coin-----------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |






<a name="svc-biz-asset-GetUserCoinResp"></a>

### GetUserCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-GetUserMoneyMultiReq"></a>

### GetUserMoneyMultiReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [string](#string) | repeated |  |






<a name="svc-biz-asset-GetUserMoneyMultiResp"></a>

### GetUserMoneyMultiResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
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
| uid | [string](#string) |  |  |






<a name="svc-biz-asset-GetUserMoneyResp"></a>

### GetUserMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="svc-biz-asset-IncrAnchorCoinReq"></a>

### IncrAnchorCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-IncrAnchorCoinResp"></a>

### IncrAnchorCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [AnchorCoinValue](#svc-biz-asset-AnchorCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrMerchantCoinReq"></a>

### IncrMerchantCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-IncrMerchantCoinResp"></a>

### IncrMerchantCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [MerchantCoinValue](#svc-biz-asset-MerchantCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrMerchantMoneyReq"></a>

### IncrMerchantMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-IncrMerchantMoneyResp"></a>

### IncrMerchantMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| new_value | [int64](#int64) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrUnionCoinReq"></a>

### IncrUnionCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-IncrUnionCoinResp"></a>

### IncrUnionCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [UnionCoinValue](#svc-biz-asset-UnionCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrUserCoinReq"></a>

### IncrUserCoinReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别,增加普通余额或低权限余额按此字段自动判断 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-IncrUserCoinResp"></a>

### IncrUserCoinResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| new_value | [UserCoinValue](#svc-biz-asset-UserCoinValue) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-IncrUserMoneyReq"></a>

### IncrUserMoneyReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| value | [int64](#int64) |  |  |
| trans_type | [int64](#int64) |  | 交易类别 |
| trade_id | [string](#string) |  | 业务方交易id,业务方保证唯一,支持幂等 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-IncrUserMoneyResp"></a>

### IncrUserMoneyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| new_value | [int64](#int64) |  |  |
| trade_id | [string](#string) |  | 业务方交易id |
| detail_id | [string](#string) |  | 明细id |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |






<a name="svc-biz-asset-ListAnchorCoinDetailReq"></a>

### ListAnchorCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListAnchorCoinDetailResp"></a>

### ListAnchorCoinDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [AnchorCoinDetail](#svc-biz-asset-AnchorCoinDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListMerchantCoinDetailReq"></a>

### ListMerchantCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
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






<a name="svc-biz-asset-ListMerchantMoneyDetailReq"></a>

### ListMerchantMoneyDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| trans_direction | [int64](#int64) |  | 交易方向 0 不限 1 增加 2 减少 |
| trans_type | [int64](#int64) | repeated | 交易类别: 空表示不限 |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开始时间(含) |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间(不含) |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |






<a name="svc-biz-asset-ListMerchantMoneyDetailResp"></a>

### ListMerchantMoneyDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MerchantMoneyDetail](#svc-biz-asset-MerchantMoneyDetail) | repeated |  |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 单页条数 |
| total | [int64](#int64) |  | 符合筛选的总条数 |






<a name="svc-biz-asset-ListUnionCoinDetailReq"></a>

### ListUnionCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
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






<a name="svc-biz-asset-ListUserCoinDetailReq"></a>

### ListUserCoinDetailReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
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
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value | [int64](#int64) |  |  |
| value_before | [int64](#int64) |  |  |
| value_after | [int64](#int64) |  |  |
| limited_value | [int64](#int64) |  |  |
| limited_value_before | [int64](#int64) |  |  |
| limited_value_after | [int64](#int64) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-MerchantCoinValue"></a>

### MerchantCoinValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-MerchantMoneyDetail"></a>

### MerchantMoneyDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| uid | [string](#string) |  |  |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value | [int64](#int64) |  |  |
| value_before | [int64](#int64) |  |  |
| value_after | [int64](#int64) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-UnionCoinDetail"></a>

### UnionCoinDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value | [int64](#int64) |  |  |
| value_before | [int64](#int64) |  |  |
| value_after | [int64](#int64) |  |  |
| limited_value | [int64](#int64) |  |  |
| limited_value_before | [int64](#int64) |  |  |
| limited_value_after | [int64](#int64) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-UnionCoinValue"></a>

### UnionCoinValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-UserCoinDetail"></a>

### UserCoinDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| uid | [string](#string) |  |  |
| app_id | [string](#string) |  |  |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value | [int64](#int64) |  |  |
| value_before | [int64](#int64) |  |  |
| value_after | [int64](#int64) |  |  |
| limited_value | [int64](#int64) |  |  |
| limited_value_before | [int64](#int64) |  |  |
| limited_value_after | [int64](#int64) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |






<a name="svc-biz-asset-UserCoinValue"></a>

### UserCoinValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | 余额 |
| limited_value | [int64](#int64) |  | 低权限余额 |






<a name="svc-biz-asset-UserMoneyDetail"></a>

### UserMoneyDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| detail_id | [string](#string) |  | 明细id |
| trade_id | [string](#string) |  | 业务方交易id |
| uid | [string](#string) |  |  |
| serial_number | [int64](#int64) |  | 流水号,单账户连续自增 |
| trans_direction | [int64](#int64) |  | 交易方向 1 增加 2 减少 |
| trans_type | [int64](#int64) |  | 交易类别 |
| value | [int64](#int64) |  |  |
| value_before | [int64](#int64) |  |  |
| value_after | [int64](#int64) |  |  |
| trans_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 业务时间 |
| summary | [string](#string) |  | 摘要,json 实际会做格式和核心字段校验 |





 

 

 


<a name="svc-biz-asset-Asset"></a>

### Asset


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUserMoney | [GetUserMoneyReq](#svc-biz-asset-GetUserMoneyReq) | [GetUserMoneyResp](#svc-biz-asset-GetUserMoneyResp) | 余额 money |
| GetUserMoneyMulti | [GetUserMoneyMultiReq](#svc-biz-asset-GetUserMoneyMultiReq) | [GetUserMoneyMultiResp](#svc-biz-asset-GetUserMoneyMultiResp) |  |
| IncrUserMoney | [IncrUserMoneyReq](#svc-biz-asset-IncrUserMoneyReq) | [IncrUserMoneyResp](#svc-biz-asset-IncrUserMoneyResp) |  |
| DecrUserMoney | [DecrUserMoneyReq](#svc-biz-asset-DecrUserMoneyReq) | [DecrUserMoneyResp](#svc-biz-asset-DecrUserMoneyResp) |  |
| ListUserMoneyDetail | [ListUserMoneyDetailReq](#svc-biz-asset-ListUserMoneyDetailReq) | [ListUserMoneyDetailResp](#svc-biz-asset-ListUserMoneyDetailResp) |  |
| GetUserCoin | [GetUserCoinReq](#svc-biz-asset-GetUserCoinReq) | [GetUserCoinResp](#svc-biz-asset-GetUserCoinResp) | 虚拟币coin |
| GetUserCoinMulti | [GetUserCoinMultiReq](#svc-biz-asset-GetUserCoinMultiReq) | [GetUserCoinMultiResp](#svc-biz-asset-GetUserCoinMultiResp) |  |
| IncrUserCoin | [IncrUserCoinReq](#svc-biz-asset-IncrUserCoinReq) | [IncrUserCoinResp](#svc-biz-asset-IncrUserCoinResp) |  |
| DecrUserCoin | [DecrUserCoinReq](#svc-biz-asset-DecrUserCoinReq) | [DecrUserCoinResp](#svc-biz-asset-DecrUserCoinResp) |  |
| ListUserCoinDetail | [ListUserCoinDetailReq](#svc-biz-asset-ListUserCoinDetailReq) | [ListUserCoinDetailResp](#svc-biz-asset-ListUserCoinDetailResp) |  |
| GetAnchorCoin | [GetAnchorCoinReq](#svc-biz-asset-GetAnchorCoinReq) | [GetAnchorCoinResp](#svc-biz-asset-GetAnchorCoinResp) | 虚拟币coin |
| GetAnchorCoinMulti | [GetAnchorCoinMultiReq](#svc-biz-asset-GetAnchorCoinMultiReq) | [GetAnchorCoinMultiResp](#svc-biz-asset-GetAnchorCoinMultiResp) |  |
| IncrAnchorCoin | [IncrAnchorCoinReq](#svc-biz-asset-IncrAnchorCoinReq) | [IncrAnchorCoinResp](#svc-biz-asset-IncrAnchorCoinResp) |  |
| DecrAnchorCoin | [DecrAnchorCoinReq](#svc-biz-asset-DecrAnchorCoinReq) | [DecrAnchorCoinResp](#svc-biz-asset-DecrAnchorCoinResp) |  |
| ListAnchorCoinDetail | [ListAnchorCoinDetailReq](#svc-biz-asset-ListAnchorCoinDetailReq) | [ListAnchorCoinDetailResp](#svc-biz-asset-ListAnchorCoinDetailResp) |  |
| GetUnionCoin | [GetUnionCoinReq](#svc-biz-asset-GetUnionCoinReq) | [GetUnionCoinResp](#svc-biz-asset-GetUnionCoinResp) | 虚拟币coin |
| GetUnionCoinMulti | [GetUnionCoinMultiReq](#svc-biz-asset-GetUnionCoinMultiReq) | [GetUnionCoinMultiResp](#svc-biz-asset-GetUnionCoinMultiResp) |  |
| IncrUnionCoin | [IncrUnionCoinReq](#svc-biz-asset-IncrUnionCoinReq) | [IncrUnionCoinResp](#svc-biz-asset-IncrUnionCoinResp) |  |
| DecrUnionCoin | [DecrUnionCoinReq](#svc-biz-asset-DecrUnionCoinReq) | [DecrUnionCoinResp](#svc-biz-asset-DecrUnionCoinResp) |  |
| ListUnionCoinDetail | [ListUnionCoinDetailReq](#svc-biz-asset-ListUnionCoinDetailReq) | [ListUnionCoinDetailResp](#svc-biz-asset-ListUnionCoinDetailResp) |  |
| GetMerchantMoney | [GetMerchantMoneyReq](#svc-biz-asset-GetMerchantMoneyReq) | [GetMerchantMoneyResp](#svc-biz-asset-GetMerchantMoneyResp) | 余额 money |
| GetMerchantMoneyMulti | [GetMerchantMoneyMultiReq](#svc-biz-asset-GetMerchantMoneyMultiReq) | [GetMerchantMoneyMultiResp](#svc-biz-asset-GetMerchantMoneyMultiResp) |  |
| IncrMerchantMoney | [IncrMerchantMoneyReq](#svc-biz-asset-IncrMerchantMoneyReq) | [IncrMerchantMoneyResp](#svc-biz-asset-IncrMerchantMoneyResp) |  |
| DecrMerchantMoney | [DecrMerchantMoneyReq](#svc-biz-asset-DecrMerchantMoneyReq) | [DecrMerchantMoneyResp](#svc-biz-asset-DecrMerchantMoneyResp) |  |
| ListMerchantMoneyDetail | [ListMerchantMoneyDetailReq](#svc-biz-asset-ListMerchantMoneyDetailReq) | [ListMerchantMoneyDetailResp](#svc-biz-asset-ListMerchantMoneyDetailResp) |  |
| GetMerchantCoin | [GetMerchantCoinReq](#svc-biz-asset-GetMerchantCoinReq) | [GetMerchantCoinResp](#svc-biz-asset-GetMerchantCoinResp) | 虚拟币coin |
| GetMerchantCoinMulti | [GetMerchantCoinMultiReq](#svc-biz-asset-GetMerchantCoinMultiReq) | [GetMerchantCoinMultiResp](#svc-biz-asset-GetMerchantCoinMultiResp) |  |
| IncrMerchantCoin | [IncrMerchantCoinReq](#svc-biz-asset-IncrMerchantCoinReq) | [IncrMerchantCoinResp](#svc-biz-asset-IncrMerchantCoinResp) |  |
| DecrMerchantCoin | [DecrMerchantCoinReq](#svc-biz-asset-DecrMerchantCoinReq) | [DecrMerchantCoinResp](#svc-biz-asset-DecrMerchantCoinResp) |  |
| ListMerchantCoinDetail | [ListMerchantCoinDetailReq](#svc-biz-asset-ListMerchantCoinDetailReq) | [ListMerchantCoinDetailResp](#svc-biz-asset-ListMerchantCoinDetailResp) |  |

 



<a name="svc-web-dashboard_dashboard-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.web.dashboard/dashboard.proto


 

 

 


<a name="svc-web-dashboard-Dashboard"></a>

### Dashboard


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



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
| gift_id | [int64](#int64) |  |  |






<a name="svc-biz-gift-GiftGetRecordReq"></a>

### GiftGetRecordReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pageno | [int32](#int32) |  | 第几页 |
| pagenum | [int32](#int32) |  | 每页几条数据 |
| room_id | [int64](#int64) |  | 房间id |
| live_id | [int64](#int64) |  | 直播id |
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
| gift_id | [int64](#int64) |  |  |






<a name="svc-biz-gift-GiftGetResp"></a>

### GiftGetResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift | [GiftInfo](#svc-biz-gift-GiftInfo) |  |  |






<a name="svc-biz-gift-GiftInfo"></a>

### GiftInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [int64](#int64) |  | 礼物id |
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






<a name="svc-biz-gift-GiftListAdminReq"></a>

### GiftListAdminReq
礼物列表（后台使用）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| return_count | [bool](#bool) |  | 是否返回总数 |
| pageno | [int32](#int32) |  | 第几页 |
| pagenum | [int32](#int32) |  | 每页几条数据 |
| type | [GiftType](#svc-biz-gift-GiftType) |  | 礼物类型 |
| status | [GiftStatus](#svc-biz-gift-GiftStatus) |  | 礼物状态 |
| keyword | [string](#string) |  | 关键字 |






<a name="svc-biz-gift-GiftListAdminResp"></a>

### GiftListAdminResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [GiftInfo](#svc-biz-gift-GiftInfo) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="svc-biz-gift-GiftListReq"></a>

### GiftListReq
礼物列表（房间使用）






<a name="svc-biz-gift-GiftListResp"></a>

### GiftListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [GiftInfo](#svc-biz-gift-GiftInfo) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="svc-biz-gift-GiftSendRecord"></a>

### GiftSendRecord



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gift_id | [int64](#int64) |  | 礼物id |
| gift_name | [string](#string) |  | 礼物名 |
| num | [int32](#int32) |  | 数量 |
| price | [int32](#int32) |  | 单价 |
| total_price | [int32](#int32) |  | 总价 |
| from_uid | [int64](#int64) |  | 送礼uid |
| from_nickname | [string](#string) |  | 送礼人昵称 |
| to_uid | [int64](#int64) |  | 接收者uid |
| to_nickname | [string](#string) |  | 接收者昵称 |
| room_id | [int64](#int64) |  | 房间id |
| live_id | [int64](#int64) |  | 直播id |
| create_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |






<a name="svc-biz-gift-GiftSendRecordReq"></a>

### GiftSendRecordReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pageno | [int32](#int32) |  | 第几页 |
| pagenum | [int32](#int32) |  | 每页几条数据 |
| uid | [int64](#int64) |  | 用户id |
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
| uuid | [int64](#int64) |  | 唯一标识 |
| order_id | [int64](#int64) |  | 支付订单id |
| gift_id | [int64](#int64) |  | 礼物id |
| from_uid | [int64](#int64) |  | 赠送者uid |
| to_uid | [int64](#int64) |  | 接收者uid |
| num | [int32](#int32) |  | 数量 |
| room_id | [int64](#int64) |  | 房间id |
| live_id | [int64](#int64) |  | 直播id |






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
| gift_id | [int64](#int64) |  |  |
| gift | [GiftInfo](#svc-biz-gift-GiftInfo) |  |  |
| update_filed | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="svc-biz-gift-GiftUpdateResp"></a>

### GiftUpdateResp







<a name="svc-biz-gift-LiveStatReq"></a>

### LiveStatReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room_id | [int64](#int64) |  |  |
| live_id | [int64](#int64) |  |  |






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
| Add | [GiftAddReq](#svc-biz-gift-GiftAddReq) | [GiftAddResp](#svc-biz-gift-GiftAddResp) |  |
| Get | [GiftGetReq](#svc-biz-gift-GiftGetReq) | [GiftGetResp](#svc-biz-gift-GiftGetResp) |  |
| Update | [GiftUpdateReq](#svc-biz-gift-GiftUpdateReq) | [GiftUpdateResp](#svc-biz-gift-GiftUpdateResp) |  |
| ListAdmin | [GiftListAdminReq](#svc-biz-gift-GiftListAdminReq) | [GiftListAdminResp](#svc-biz-gift-GiftListAdminResp) |  |
| List | [GiftListReq](#svc-biz-gift-GiftListReq) | [GiftListResp](#svc-biz-gift-GiftListResp) |  |
| Send | [GiftSendReq](#svc-biz-gift-GiftSendReq) | [GiftSendResp](#svc-biz-gift-GiftSendResp) |  |
| SendRecord | [GiftSendRecordReq](#svc-biz-gift-GiftSendRecordReq) | [GiftSendRecordResp](#svc-biz-gift-GiftSendRecordResp) |  |
| GetRecord | [GiftGetRecordReq](#svc-biz-gift-GiftGetRecordReq) | [GiftGetRecordResp](#svc-biz-gift-GiftGetRecordResp) |  |
| LiveStat | [LiveStatReq](#svc-biz-gift-LiveStatReq) | [LiveStatResp](#svc-biz-gift-LiveStatResp) |  |

 



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

