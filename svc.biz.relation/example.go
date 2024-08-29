package relation

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func example() {
	ctx := context.Background()
	relationClient := NewRelationClient(nil)
	// 关注
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeFollow,
			MemberId:     "用户id",
			RMemberId:    "被关注的主播id",
			BuildTime:    timestamppb.Now(),
		},
	})
	// 我关注的数量
	relationClient.GetRelationCount(ctx, &GetRelationCountReq{
		RelationType: RelationType_RelationTypeFollow,
		MemberId:     "用户id",
	})
	// 关注我得数量
	relationClient.GetRelationCount(ctx, &GetRelationCountReq{
		RelationType: RelationType_RelationTypeFollow,
		RMemberId:    "被关注的主播id",
	})

	// 观看历史
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeHistory,
			MemberId:     "用户id",
			RMemberId:    "观看的主播id",
			BuildTime:    timestamppb.Now(),
		},
	})

	// 全平台禁言
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeMuzzle,
			MemberId:     RelationMemberPlatform,
			RMemberId:    "被禁言的用户id",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 禁言24小时
			// ExpireTime:   nil, // 永久禁言
			BuildTime: timestamppb.Now(),
		},
	})

	// 商户禁言
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeMuzzle,
			MemberId:     "商户id",
			RMemberId:    "被禁言的用户id",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 禁言24小时
			// ExpireTime:   nil, // 永久禁言
			BuildTime: timestamppb.Now(),
		},
	})

	// 房间禁言
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeMuzzle,
			MemberId:     "主播id",
			RMemberId:    "被禁言的用户id",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 禁言24小时
			// ExpireTime:   nil, // 永久禁言
			BuildTime: timestamppb.Now(),
		},
	})

	// 用户黑名单
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeBlacklistViewer,
			MemberId:     RelationMemberPlatform,
			RMemberId:    "被拉黑的用户id",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 封禁24小时
			// ExpireTime:   nil, // 永久拉黑
			BuildTime: timestamppb.Now(),
		},
	})

	// 主播黑名单
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeBlacklistStreamer,
			MemberId:     RelationMemberPlatform,
			RMemberId:    "被拉黑的主播id",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 封禁24小时
			// ExpireTime:   nil, // 永久拉黑
			BuildTime: timestamppb.Now(),
		},
	})

	// ip黑名单
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeBlacklistIP,
			MemberId:     RelationMemberPlatform,
			RMemberId:    "被封禁的ip",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 封禁24小时
			// ExpireTime:   nil, // 永久封禁
			BuildTime: timestamppb.Now(),
		},
	})

	// 设备黑名单
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeBlacklistDevice,
			MemberId:     RelationMemberPlatform,
			RMemberId:    "被封禁的设备标识",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 封禁24小时
			// ExpireTime:   nil, // 永久封禁
			BuildTime: timestamppb.Now(),
		},
	})

	// 超级管理员
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeSuperManager,
			MemberId:     RelationMemberPlatform,
			RMemberId:    "超级管理员用户id",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 超管24个小时有效
			// ExpireTime:   nil, // 永久有效
			BuildTime: timestamppb.Now(),
		},
	})

	// 房间管路
	relationClient.RelationAdd(ctx, &RelationAddReq{
		RelationInfo: &RelationInfo{
			RelationType: RelationType_RelationTypeRoomManager,
			MemberId:     "主播id",
			RMemberId:    "房管用户id",
			ExpireTime:   timestamppb.New(time.Now().Add(24 * time.Hour)), // 房管24个小时有效
			// ExpireTime:   nil, // 永久有效
			BuildTime: timestamppb.Now(),
		},
	})
}
