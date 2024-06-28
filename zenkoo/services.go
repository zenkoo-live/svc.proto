// - GENERATED BY generate.go
// - DO NOT MODIFY THIS FILE

package zenkoo

import (
	cltGrpc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/zenkoo-live/svc.base/runtime"

	svcBizAccount "github.com/zenkoo-live/svc.proto/svc.biz.account"
	svcBizAsset "github.com/zenkoo-live/svc.proto/svc.biz.asset"
	svcBizGift "github.com/zenkoo-live/svc.proto/svc.biz.gift"
	svcBizLog "github.com/zenkoo-live/svc.proto/svc.biz.log"
	svcBizOrg "github.com/zenkoo-live/svc.proto/svc.biz.org"
	svcBizRelation "github.com/zenkoo-live/svc.proto/svc.biz.relation"
	svcBizRoom "github.com/zenkoo-live/svc.proto/svc.biz.room"
	svcBizTrade "github.com/zenkoo-live/svc.proto/svc.biz.trade"
	svcBizVip "github.com/zenkoo-live/svc.proto/svc.biz.vip"
	svcInfraLink "github.com/zenkoo-live/svc.proto/svc.infra.link"
	svcInfraNotifier "github.com/zenkoo-live/svc.proto/svc.infra.notifier"
	svcInfraPay "github.com/zenkoo-live/svc.proto/svc.infra.pay"
	svcInfraSetting "github.com/zenkoo-live/svc.proto/svc.infra.setting"
	svcInfraStatic "github.com/zenkoo-live/svc.proto/svc.infra.static"
	svcWebDashboard "github.com/zenkoo-live/svc.proto/svc.web.dashboard"
	svcWebStreamer "github.com/zenkoo-live/svc.proto/svc.web.streamer"
	svcWebViewer "github.com/zenkoo-live/svc.proto/svc.web.viewer"
)

const (
	AppName = "zenkoo"

	SvcBizAccount    = "svc.biz.account"
	SvcBizAsset      = "svc.biz.asset"
	SvcWebDashboard  = "svc.web.dashboard"
	SvcBizGift       = "svc.biz.gift"
	SvcInfraLink     = "svc.infra.link"
	SvcBizLog        = "svc.biz.log"
	SvcInfraNotifier = "svc.infra.notifier"
	SvcBizOrg        = "svc.biz.org"
	SvcInfraPay      = "svc.infra.pay"
	SvcBizRelation   = "svc.biz.relation"
	SvcBizRoom       = "svc.biz.room"
	SvcInfraSetting  = "svc.infra.setting"
	SvcInfraStatic   = "svc.infra.static"
	SvcWebStreamer   = "svc.web.streamer"
	SvcBizTrade      = "svc.biz.trade"
	SvcWebViewer     = "svc.web.viewer"
	SvcBizVip        = "svc.biz.vip"
)

var (
	clt = cltGrpc.NewClient()

	SvcBizAccountAccount     = svcBizAccount.NewAccountService(AppName+"::"+SvcBizAccount+runtime.AppendEnv(), clt)
	SvcBizAssetAsset         = svcBizAsset.NewAssetService(AppName+"::"+SvcBizAsset+runtime.AppendEnv(), clt)
	SvcWebDashboardDashboard = svcWebDashboard.NewDashboardService(AppName+"::"+SvcWebDashboard+runtime.AppendEnv(), clt)
	SvcBizGiftGift           = svcBizGift.NewGiftService(AppName+"::"+SvcBizGift+runtime.AppendEnv(), clt)
	SvcBizGiftGift_record    = svcBizGift.NewGiftRecordService(AppName+"::"+SvcBizGift+runtime.AppendEnv(), clt)
	SvcInfraLinkGateway      = svcInfraLink.NewLinkGatewayService(AppName+"::"+SvcInfraLink+runtime.AppendEnv(), clt)
	SvcInfraLinkInstruction  = svcInfraLink.NewLinkInstructionService(AppName+"::"+SvcInfraLink+runtime.AppendEnv(), clt)
	SvcInfraLinkMessage      = svcInfraLink.NewLinkMessageService(AppName+"::"+SvcInfraLink+runtime.AppendEnv(), clt)
	SvcInfraLinkStat         = svcInfraLink.NewLinkStatService(AppName+"::"+SvcInfraLink+runtime.AppendEnv(), clt)
	SvcInfraLinkTrace        = svcInfraLink.NewLinkTraceService(AppName+"::"+SvcInfraLink+runtime.AppendEnv(), clt)
	SvcBizLogLog             = svcBizLog.NewLogService(AppName+"::"+SvcBizLog+runtime.AppendEnv(), clt)
	SvcInfraNotifierNotifier = svcInfraNotifier.NewNotifierService(AppName+"::"+SvcInfraNotifier+runtime.AppendEnv(), clt)
	SvcBizOrgOrg             = svcBizOrg.NewOrgService(AppName+"::"+SvcBizOrg+runtime.AppendEnv(), clt)
	SvcInfraPayPay           = svcInfraPay.NewPayService(AppName+"::"+SvcInfraPay+runtime.AppendEnv(), clt)
	SvcBizRelationRelation   = svcBizRelation.NewRelationService(AppName+"::"+SvcBizRelation+runtime.AppendEnv(), clt)
	SvcBizRoomCategory       = svcBizRoom.NewCategoryService(AppName+"::"+SvcBizRoom+runtime.AppendEnv(), clt)
	SvcBizRoomLive           = svcBizRoom.NewLiveService(AppName+"::"+SvcBizRoom+runtime.AppendEnv(), clt)
	SvcBizRoomRoom           = svcBizRoom.NewRoomService(AppName+"::"+SvcBizRoom+runtime.AppendEnv(), clt)
	SvcInfraSettingSetting   = svcInfraSetting.NewSettingService(AppName+"::"+SvcInfraSetting+runtime.AppendEnv(), clt)
	SvcInfraStaticStatic     = svcInfraStatic.NewStaticService(AppName+"::"+SvcInfraStatic+runtime.AppendEnv(), clt)
	SvcWebStreamerStreamer   = svcWebStreamer.NewStreamerService(AppName+"::"+SvcWebStreamer+runtime.AppendEnv(), clt)
	SvcBizTradeTrade         = svcBizTrade.NewTradeService(AppName+"::"+SvcBizTrade+runtime.AppendEnv(), clt)
	SvcWebViewerViewer       = svcWebViewer.NewViewerService(AppName+"::"+SvcWebViewer+runtime.AppendEnv(), clt)
	SvcBizVipFanbase         = svcBizVip.NewFanbaseService(AppName+"::"+SvcBizVip+runtime.AppendEnv(), clt)
	SvcBizVipFanbase_member  = svcBizVip.NewFanbaseMemberService(AppName+"::"+SvcBizVip+runtime.AppendEnv(), clt)
	SvcBizVipLevel           = svcBizVip.NewLevelService(AppName+"::"+SvcBizVip+runtime.AppendEnv(), clt)
	SvcBizVipNoble           = svcBizVip.NewNobleService(AppName+"::"+SvcBizVip+runtime.AppendEnv(), clt)
	SvcBizVipNoble_member    = svcBizVip.NewNobleMemberService(AppName+"::"+SvcBizVip+runtime.AppendEnv(), clt)
)
