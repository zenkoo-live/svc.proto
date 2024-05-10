# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [svc.web.streamer/streamer.proto](#svc-web-streamer_streamer-proto)
    - [Streamer](#svc-web-streamer-Streamer)
  
- [svc.infra.setting/setting.proto](#svc-infra-setting_setting-proto)
    - [SettingGreetingReq](#svc-infra-setting-SettingGreetingReq)
    - [SettingGreetingResp](#svc-infra-setting-SettingGreetingResp)
  
    - [Setting](#svc-infra-setting-Setting)
  
- [svc.web.viewer/viewer.proto](#svc-web-viewer_viewer-proto)
    - [Viewer](#svc-web-viewer-Viewer)
  
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
  
- [svc.web.dashboard/dashboard.proto](#svc-web-dashboard_dashboard-proto)
    - [Dashboard](#svc-web-dashboard-Dashboard)
  
- [svc.biz.account/account.proto](#svc-biz-account_account-proto)
    - [Account](#svc-biz-account-Account)
  
- [Scalar Value Types](#scalar-value-types)



<a name="svc-web-streamer_streamer-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.web.streamer/streamer.proto


 

 

 


<a name="svc-web-streamer-Streamer"></a>

### Streamer


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
| pageno | [int32](#int32) |  |  |
| pagenum | [int32](#int32) |  |  |
| type | [GiftType](#svc-biz-gift-GiftType) |  |  |
| status | [GiftStatus](#svc-biz-gift-GiftStatus) |  |  |
| keyword | [string](#string) |  |  |






<a name="svc-biz-gift-GiftListAdminResp"></a>

### GiftListAdminResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [GiftInfo](#svc-biz-gift-GiftInfo) | repeated |  |
| total | [int32](#int32) |  |  |






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

 



<a name="svc-web-dashboard_dashboard-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.web.dashboard/dashboard.proto


 

 

 


<a name="svc-web-dashboard-Dashboard"></a>

### Dashboard


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="svc-biz-account_account-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.account/account.proto


 

 

 


<a name="svc-biz-account-Account"></a>

### Account


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



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

