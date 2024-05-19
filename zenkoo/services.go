// - GENERATED BY generate.go
// - DO NOT MODIFY THIS FILE

package zenkoo

import (
	cltGrpc "github.com/go-micro/plugins/v4/client/grpc"

	svcBizAccount "github.com/zenkoo-live/svc.proto/svc.biz.account"
	svcBizGift "github.com/zenkoo-live/svc.proto/svc.biz.gift"
	svcBizRoom "github.com/zenkoo-live/svc.proto/svc.biz.room"
	svcInfraSetting "github.com/zenkoo-live/svc.proto/svc.infra.setting"
	svcWebDashboard "github.com/zenkoo-live/svc.proto/svc.web.dashboard"
	svcWebStreamer "github.com/zenkoo-live/svc.proto/svc.web.streamer"
	svcWebViewer "github.com/zenkoo-live/svc.proto/svc.web.viewer"
)

const (
	AppName = "zenkoo"

	SvcBizAccount   = "svc.biz.account"
	SvcWebDashboard = "svc.web.dashboard"
	SvcBizGift      = "svc.biz.gift"
	SvcBizRoom      = "svc.biz.room"
	SvcInfraSetting = "svc.infra.setting"
	SvcWebStreamer  = "svc.web.streamer"
	SvcWebViewer    = "svc.web.viewer"
)

var (
	clt = cltGrpc.NewClient()

	SvcBizAccountAccount     = svcBizAccount.NewAccountService(AppName+"::"+SvcBizAccount, clt)
	SvcWebDashboardDashboard = svcWebDashboard.NewDashboardService(AppName+"::"+SvcWebDashboard, clt)
	SvcBizGiftGift           = svcBizGift.NewGiftService(AppName+"::"+SvcBizGift, clt)
	SvcBizRoomRoom           = svcBizRoom.NewCategoryService(AppName+"::"+SvcBizRoom, clt)
	SvcInfraSettingSetting   = svcInfraSetting.NewSettingService(AppName+"::"+SvcInfraSetting, clt)
	SvcWebStreamerStreamer   = svcWebStreamer.NewStreamerService(AppName+"::"+SvcWebStreamer, clt)
	SvcWebViewerViewer       = svcWebViewer.NewViewerService(AppName+"::"+SvcWebViewer, clt)
)
