# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [svc.biz.account/account.proto](#svc-biz-account_account-proto)
    - [AddManagerReq](#svc-biz-account-AddManagerReq)
    - [AddManagerResp](#svc-biz-account-AddManagerResp)
    - [AddStreamerReq](#svc-biz-account-AddStreamerReq)
    - [AddStreamerResp](#svc-biz-account-AddStreamerResp)
    - [AddViewerReq](#svc-biz-account-AddViewerReq)
    - [AddViewerResp](#svc-biz-account-AddViewerResp)
    - [DeleteManagerReq](#svc-biz-account-DeleteManagerReq)
    - [DeleteManagerResp](#svc-biz-account-DeleteManagerResp)
    - [DeleteStreamerReq](#svc-biz-account-DeleteStreamerReq)
    - [DeleteStreamerResp](#svc-biz-account-DeleteStreamerResp)
    - [DeleteViewerReq](#svc-biz-account-DeleteViewerReq)
    - [DeleteViewerResp](#svc-biz-account-DeleteViewerResp)
    - [GetManagerReq](#svc-biz-account-GetManagerReq)
    - [GetManagerResp](#svc-biz-account-GetManagerResp)
    - [GetStreamerReq](#svc-biz-account-GetStreamerReq)
    - [GetStreamerResp](#svc-biz-account-GetStreamerResp)
    - [GetViewerReq](#svc-biz-account-GetViewerReq)
    - [GetViewerResp](#svc-biz-account-GetViewerResp)
    - [InitDBResp](#svc-biz-account-InitDBResp)
    - [ListManagersReq](#svc-biz-account-ListManagersReq)
    - [ListManagersResp](#svc-biz-account-ListManagersResp)
    - [ListStreamersReq](#svc-biz-account-ListStreamersReq)
    - [ListStreamersResp](#svc-biz-account-ListStreamersResp)
    - [ListViewersReq](#svc-biz-account-ListViewersReq)
    - [ListViewersResp](#svc-biz-account-ListViewersResp)
    - [Manager](#svc-biz-account-Manager)
    - [Streamer](#svc-biz-account-Streamer)
    - [UpdateManagerReq](#svc-biz-account-UpdateManagerReq)
    - [UpdateManagerResp](#svc-biz-account-UpdateManagerResp)
    - [UpdateStreamerReq](#svc-biz-account-UpdateStreamerReq)
    - [UpdateStreamerResp](#svc-biz-account-UpdateStreamerResp)
    - [UpdateViewerReq](#svc-biz-account-UpdateViewerReq)
    - [UpdateViewerResp](#svc-biz-account-UpdateViewerResp)
    - [Viewer](#svc-biz-account-Viewer)
  
    - [ManagerTypes](#svc-biz-account-ManagerTypes)
  
    - [Account](#svc-biz-account-Account)
  
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
    - [GetLiveInfoReq](#svc-biz-room-GetLiveInfoReq)
    - [GetLiveInfoResp](#svc-biz-room-GetLiveInfoResp)
    - [LiveInfo](#svc-biz-room-LiveInfo)
    - [MGetLiveInfoReq](#svc-biz-room-MGetLiveInfoReq)
    - [MGetLiveInfoResp](#svc-biz-room-MGetLiveInfoResp)
    - [MGetLiveInfoResp.ItemsEntry](#svc-biz-room-MGetLiveInfoResp-ItemsEntry)
    - [OnlineLiveListReq](#svc-biz-room-OnlineLiveListReq)
    - [OnlineLiveListResp](#svc-biz-room-OnlineLiveListResp)
    - [StartLiveReq](#svc-biz-room-StartLiveReq)
    - [StartLiveResp](#svc-biz-room-StartLiveResp)
    - [StopLiveReq](#svc-biz-room-StopLiveReq)
    - [StopLiveResp](#svc-biz-room-StopLiveResp)
    - [UpdateLiveReq](#svc-biz-room-UpdateLiveReq)
  
    - [LiveDisplayType](#svc-biz-room-LiveDisplayType)
    - [LiveStatus](#svc-biz-room-LiveStatus)
  
    - [Live](#svc-biz-room-Live)
  
- [svc.biz.room/room.proto](#svc-biz-room_room-proto)
    - [CreateRoomReq](#svc-biz-room-CreateRoomReq)
    - [CreateRoomResp](#svc-biz-room-CreateRoomResp)
    - [GetRoomListReq](#svc-biz-room-GetRoomListReq)
    - [GetRoomListResp](#svc-biz-room-GetRoomListResp)
    - [GetRoomReq](#svc-biz-room-GetRoomReq)
    - [GetRoomResp](#svc-biz-room-GetRoomResp)
    - [RoomInfo](#svc-biz-room-RoomInfo)
    - [UpdateRoomReq](#svc-biz-room-UpdateRoomReq)
  
    - [Room](#svc-biz-room-Room)
  
- [svc.infra.setting/setting.proto](#svc-infra-setting_setting-proto)
    - [SettingGreetingReq](#svc-infra-setting-SettingGreetingReq)
    - [SettingGreetingResp](#svc-infra-setting-SettingGreetingResp)
  
    - [Setting](#svc-infra-setting-Setting)
  
- [svc.web.viewer/viewer.proto](#svc-web-viewer_viewer-proto)
    - [Viewer](#svc-web-viewer-Viewer)
  
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






<a name="svc-biz-account-ListViewersReq"></a>

### ListViewersReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [Viewer](#svc-biz-account-Viewer) |  |  |






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
| type | [ManagerTypes](#svc-biz-account-ManagerTypes) |  | 类型：平台 / 商户 / 工会 |
| name | [string](#string) |  | 用户名 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| password | [string](#string) | optional | 密码 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |






<a name="svc-biz-account-Streamer"></a>

### Streamer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 内部ID |
| display_id | [string](#string) |  | 用于显示和索引的数字ID |
| name | [string](#string) |  | 用户名 |
| mobile | [string](#string) |  | 手机 |
| email | [string](#string) |  | 邮箱 |
| password | [string](#string) | optional | 密码 |
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
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |





 


<a name="svc-biz-account-ManagerTypes"></a>

### ManagerTypes


| Name | Number | Description |
| ---- | ------ | ----------- |
| ManagerTypeUnknown | 0 | 未知 |
| ManagerTypePlatform | 1 | 平台管理者 |
| ManagerTypeMerchant | 2 | 商户管理者 |
| ManagerTypeUnion | 3 | 工会 |


 

 


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
| GetManager | [GetManagerReq](#svc-biz-account-GetManagerReq) | [GetManagerResp](#svc-biz-account-GetManagerResp) | 获取主播账号 |
| ListManagers | [ListManagersReq](#svc-biz-account-ListManagersReq) | [ListManagersResp](#svc-biz-account-ListManagersResp) | 获取主播账号列表 |
| AddManager | [AddManagerReq](#svc-biz-account-AddManagerReq) | [AddManagerResp](#svc-biz-account-AddManagerResp) | 添加普通账号 |
| UpdateManager | [UpdateManagerReq](#svc-biz-account-UpdateManagerReq) | [UpdateManagerResp](#svc-biz-account-UpdateManagerResp) | 更新主播账号 |
| DeleteManager | [DeleteManagerReq](#svc-biz-account-DeleteManagerReq) | [DeleteManagerResp](#svc-biz-account-DeleteManagerResp) | 删除主播账号 |

 



<a name="svc-biz-room_category-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## svc.biz.room/category.proto



<a name="svc-biz-room-CategoryInfo"></a>

### CategoryInfo
CategoryInfo 分类详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| parent_id | [string](#string) |  | 父级ID |
| category_code | [string](#string) |  | 代号（唯一，预留） |
| category_name | [string](#string) |  | 名称 |
| sort | [int32](#int32) |  | 排序 |
| childrens | [CategoryInfo](#svc-biz-room-CategoryInfo) | repeated |  |






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
| parent_id | [string](#string) |  | 父级ID，跟级分类传1 |
| return_count | [bool](#bool) |  | 是否返回总数 |






<a name="svc-biz-room-ListCategoryResp"></a>

### ListCategoryResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [CategoryInfo](#svc-biz-room-CategoryInfo) | repeated |  |






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



<a name="svc-biz-room-GetLiveInfoReq"></a>

### GetLiveInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live_id | [string](#string) |  | 直播id |






<a name="svc-biz-room-GetLiveInfoResp"></a>

### GetLiveInfoResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live | [LiveInfo](#svc-biz-room-LiveInfo) |  | 直播信息 |






<a name="svc-biz-room-LiveInfo"></a>

### LiveInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id |
| streamer_id | [string](#string) |  | 主播id |
| category_id | [string](#string) |  | 分类id |
| status | [LiveStatus](#svc-biz-room-LiveStatus) |  | 状态：1表示开播，2表示关播 |
| display_type | [LiveDisplayType](#svc-biz-room-LiveDisplayType) |  | 横竖屏类型 |
| start_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 开播时间 |
| end_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 结束时间 |
| score_recommend | [int32](#int32) |  | 推荐分数 |
| score_search | [int32](#int32) |  | 搜索分数 |
| score_subscribe | [int32](#int32) |  | 关注分数 |
| score_hot | [int32](#int32) |  | 热度分数 |
| score_glamour | [int32](#int32) |  | 魅力分数 |
| score_online | [int32](#int32) |  | 在线人数 |
| score_gift | [int32](#int32) |  | 礼物分数 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="svc-biz-room-MGetLiveInfoReq"></a>

### MGetLiveInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live_ids | [string](#string) | repeated | 直播id列表 |






<a name="svc-biz-room-MGetLiveInfoResp"></a>

### MGetLiveInfoResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [MGetLiveInfoResp.ItemsEntry](#svc-biz-room-MGetLiveInfoResp-ItemsEntry) | repeated | 直播信息 |






<a name="svc-biz-room-MGetLiveInfoResp-ItemsEntry"></a>

### MGetLiveInfoResp.ItemsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [LiveInfo](#svc-biz-room-LiveInfo) |  |  |






<a name="svc-biz-room-OnlineLiveListReq"></a>

### OnlineLiveListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  | 页数 |
| limit | [int32](#int32) |  | 条数 |
| category_id | [string](#string) |  | 分类id |






<a name="svc-biz-room-OnlineLiveListResp"></a>

### OnlineLiveListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [LiveInfo](#svc-biz-room-LiveInfo) | repeated | 直播信息 |






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
| live_id | [string](#string) |  | 直播id |






<a name="svc-biz-room-StopLiveReq"></a>

### StopLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streamer_id | [string](#string) |  | 主播id |






<a name="svc-biz-room-StopLiveResp"></a>

### StopLiveResp







<a name="svc-biz-room-UpdateLiveReq"></a>

### UpdateLiveReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| live_id | [string](#string) |  | 直播id |
| live | [LiveInfo](#svc-biz-room-LiveInfo) |  | 直播信息 |
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


 

 


<a name="svc-biz-room-Live"></a>

### Live
房间

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartLive | [StartLiveReq](#svc-biz-room-StartLiveReq) | [StartLiveResp](#svc-biz-room-StartLiveResp) | 开始直播 |
| StopLive | [StopLiveReq](#svc-biz-room-StopLiveReq) | [StopLiveResp](#svc-biz-room-StopLiveResp) | 关闭直播 |
| UpdateLive | [UpdateLiveReq](#svc-biz-room-UpdateLiveReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新直播 |
| GetLiveInfo | [GetLiveInfoReq](#svc-biz-room-GetLiveInfoReq) | [GetLiveInfoResp](#svc-biz-room-GetLiveInfoResp) | 查询直播间信息 |
| MGetLiveInfo | [MGetLiveInfoReq](#svc-biz-room-MGetLiveInfoReq) | [MGetLiveInfoResp](#svc-biz-room-MGetLiveInfoResp) | 批量获取直播间信息 |
| OnlineLiveList | [OnlineLiveListReq](#svc-biz-room-OnlineLiveListReq) | [OnlineLiveListResp](#svc-biz-room-OnlineLiveListResp) | 获取在播直播间列表 |

 



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
| id | [string](#string) |  | id |






<a name="svc-biz-room-GetRoomResp"></a>

### GetRoomResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |






<a name="svc-biz-room-RoomInfo"></a>

### RoomInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id |
| display_id | [string](#string) |  | 显示id |
| streamer_id | [string](#string) |  | 主播id |
| title | [string](#string) |  | 标题 |
| intro | [string](#string) |  | 简介 |
| status | [LiveStatus](#svc-biz-room-LiveStatus) |  | 房间状态：1关播，2开播 |
| last_category_id | [int64](#int64) |  | 最后一次板块标识 |
| score_recommend | [int32](#int32) |  | 推荐分数 |
| score_search | [int32](#int32) |  | 搜索分数 |
| score_subscribe | [int32](#int32) |  | 关注分数 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="svc-biz-room-UpdateRoomReq"></a>

### UpdateRoomReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id |
| room | [RoomInfo](#svc-biz-room-RoomInfo) |  | 房间信息 |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |





 

 

 


<a name="svc-biz-room-Room"></a>

### Room
房间

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateRoom | [CreateRoomReq](#svc-biz-room-CreateRoomReq) | [CreateRoomResp](#svc-biz-room-CreateRoomResp) | 创建房间 |
| UpdateRoom | [UpdateRoomReq](#svc-biz-room-UpdateRoomReq) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新房间 |
| GetRoom | [GetRoomReq](#svc-biz-room-GetRoomReq) | [GetRoomResp](#svc-biz-room-GetRoomResp) | 查询房间 |
| GetRoomList | [GetRoomListReq](#svc-biz-room-GetRoomListReq) | [GetRoomListResp](#svc-biz-room-GetRoomListResp) | 查询房间列表，直读mysql |

 



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

