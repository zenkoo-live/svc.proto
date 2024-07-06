// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: svc.biz.asset/asset.proto

package asset

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Asset service

func NewAssetEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Asset service

type AssetService interface {
	// 余额 money
	GetUserMoney(ctx context.Context, in *GetUserMoneyReq, opts ...client.CallOption) (*GetUserMoneyResp, error)
	GetUserMoneyMulti(ctx context.Context, in *GetUserMoneyMultiReq, opts ...client.CallOption) (*GetUserMoneyMultiResp, error)
	IncrUserMoney(ctx context.Context, in *IncrUserMoneyReq, opts ...client.CallOption) (*ChangeUserMoneyResp, error)
	DecrUserMoney(ctx context.Context, in *DecrUserMoneyReq, opts ...client.CallOption) (*ChangeUserMoneyResp, error)
	ListUserMoneyDetail(ctx context.Context, in *ListUserMoneyDetailReq, opts ...client.CallOption) (*ListUserMoneyDetailResp, error)
	// 虚拟币coin
	GetUserCoin(ctx context.Context, in *GetUserCoinReq, opts ...client.CallOption) (*GetUserCoinResp, error)
	GetUserCoinMulti(ctx context.Context, in *GetUserCoinMultiReq, opts ...client.CallOption) (*GetUserCoinMultiResp, error)
	IncrUserCoin(ctx context.Context, in *IncrUserCoinReq, opts ...client.CallOption) (*ChangeUserCoinResp, error)
	DecrUserCoin(ctx context.Context, in *DecrUserCoinReq, opts ...client.CallOption) (*ChangeUserCoinResp, error)
	ListUserCoinDetail(ctx context.Context, in *ListUserCoinDetailReq, opts ...client.CallOption) (*ListUserCoinDetailResp, error)
	// ---------------Streamer主播资产---------------
	// 余额 money
	GetStreamerMoney(ctx context.Context, in *GetStreamerMoneyReq, opts ...client.CallOption) (*GetStreamerMoneyResp, error)
	GetStreamerMoneyMulti(ctx context.Context, in *GetStreamerMoneyMultiReq, opts ...client.CallOption) (*GetStreamerMoneyMultiResp, error)
	IncrStreamerMoney(ctx context.Context, in *IncrStreamerMoneyReq, opts ...client.CallOption) (*ChangeStreamerMoneyResp, error)
	DecrStreamerMoney(ctx context.Context, in *DecrStreamerMoneyReq, opts ...client.CallOption) (*ChangeStreamerMoneyResp, error)
	ListStreamerMoneyDetail(ctx context.Context, in *ListStreamerMoneyDetailReq, opts ...client.CallOption) (*ListStreamerMoneyDetailResp, error)
	// 虚拟币coin
	GetStreamerCoin(ctx context.Context, in *GetStreamerCoinReq, opts ...client.CallOption) (*GetStreamerCoinResp, error)
	GetStreamerCoinMulti(ctx context.Context, in *GetStreamerCoinMultiReq, opts ...client.CallOption) (*GetStreamerCoinMultiResp, error)
	IncrStreamerCoin(ctx context.Context, in *IncrStreamerCoinReq, opts ...client.CallOption) (*ChangeStreamerCoinResp, error)
	DecrStreamerCoin(ctx context.Context, in *DecrStreamerCoinReq, opts ...client.CallOption) (*ChangeStreamerCoinResp, error)
	ListStreamerCoinDetail(ctx context.Context, in *ListStreamerCoinDetailReq, opts ...client.CallOption) (*ListStreamerCoinDetailResp, error)
	GetStreamerCoinTotal(ctx context.Context, in *GetStreamerCoinTotalReq, opts ...client.CallOption) (*GetStreamerCoinTotalResp, error)
	// ---------------Union工会资产---------------
	// 余额 money --
	GetUnionMoney(ctx context.Context, in *GetUnionMoneyReq, opts ...client.CallOption) (*GetUnionMoneyResp, error)
	GetUnionMoneyMulti(ctx context.Context, in *GetUnionMoneyMultiReq, opts ...client.CallOption) (*GetUnionMoneyMultiResp, error)
	IncrUnionMoney(ctx context.Context, in *IncrUnionMoneyReq, opts ...client.CallOption) (*ChangeUnionMoneyResp, error)
	DecrUnionMoney(ctx context.Context, in *DecrUnionMoneyReq, opts ...client.CallOption) (*ChangeUnionMoneyResp, error)
	ListUnionMoneyDetail(ctx context.Context, in *ListUnionMoneyDetailReq, opts ...client.CallOption) (*ListUnionMoneyDetailResp, error)
	// 虚拟币coin
	GetUnionCoin(ctx context.Context, in *GetUnionCoinReq, opts ...client.CallOption) (*GetUnionCoinResp, error)
	GetUnionCoinMulti(ctx context.Context, in *GetUnionCoinMultiReq, opts ...client.CallOption) (*GetUnionCoinMultiResp, error)
	IncrUnionCoin(ctx context.Context, in *IncrUnionCoinReq, opts ...client.CallOption) (*ChangeUnionCoinResp, error)
	DecrUnionCoin(ctx context.Context, in *DecrUnionCoinReq, opts ...client.CallOption) (*ChangeUnionCoinResp, error)
	ListUnionCoinDetail(ctx context.Context, in *ListUnionCoinDetailReq, opts ...client.CallOption) (*ListUnionCoinDetailResp, error)
	// 虚拟币coin
	GetMerchantCoin(ctx context.Context, in *GetMerchantCoinReq, opts ...client.CallOption) (*GetMerchantCoinResp, error)
	GetMerchantCoinMulti(ctx context.Context, in *GetMerchantCoinMultiReq, opts ...client.CallOption) (*GetMerchantCoinMultiResp, error)
	IncrMerchantCoin(ctx context.Context, in *IncrMerchantCoinReq, opts ...client.CallOption) (*ChangeMerchantCoinResp, error)
	DecrMerchantCoin(ctx context.Context, in *DecrMerchantCoinReq, opts ...client.CallOption) (*ChangeMerchantCoinResp, error)
	ListMerchantCoinDetail(ctx context.Context, in *ListMerchantCoinDetailReq, opts ...client.CallOption) (*ListMerchantCoinDetailResp, error)
}

type assetService struct {
	c    client.Client
	name string
}

func NewAssetService(name string, c client.Client) AssetService {
	return &assetService{
		c:    c,
		name: name,
	}
}

func (c *assetService) GetUserMoney(ctx context.Context, in *GetUserMoneyReq, opts ...client.CallOption) (*GetUserMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUserMoney", in)
	out := new(GetUserMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetUserMoneyMulti(ctx context.Context, in *GetUserMoneyMultiReq, opts ...client.CallOption) (*GetUserMoneyMultiResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUserMoneyMulti", in)
	out := new(GetUserMoneyMultiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) IncrUserMoney(ctx context.Context, in *IncrUserMoneyReq, opts ...client.CallOption) (*ChangeUserMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.IncrUserMoney", in)
	out := new(ChangeUserMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) DecrUserMoney(ctx context.Context, in *DecrUserMoneyReq, opts ...client.CallOption) (*ChangeUserMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.DecrUserMoney", in)
	out := new(ChangeUserMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) ListUserMoneyDetail(ctx context.Context, in *ListUserMoneyDetailReq, opts ...client.CallOption) (*ListUserMoneyDetailResp, error) {
	req := c.c.NewRequest(c.name, "Asset.ListUserMoneyDetail", in)
	out := new(ListUserMoneyDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetUserCoin(ctx context.Context, in *GetUserCoinReq, opts ...client.CallOption) (*GetUserCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUserCoin", in)
	out := new(GetUserCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetUserCoinMulti(ctx context.Context, in *GetUserCoinMultiReq, opts ...client.CallOption) (*GetUserCoinMultiResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUserCoinMulti", in)
	out := new(GetUserCoinMultiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) IncrUserCoin(ctx context.Context, in *IncrUserCoinReq, opts ...client.CallOption) (*ChangeUserCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.IncrUserCoin", in)
	out := new(ChangeUserCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) DecrUserCoin(ctx context.Context, in *DecrUserCoinReq, opts ...client.CallOption) (*ChangeUserCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.DecrUserCoin", in)
	out := new(ChangeUserCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) ListUserCoinDetail(ctx context.Context, in *ListUserCoinDetailReq, opts ...client.CallOption) (*ListUserCoinDetailResp, error) {
	req := c.c.NewRequest(c.name, "Asset.ListUserCoinDetail", in)
	out := new(ListUserCoinDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetStreamerMoney(ctx context.Context, in *GetStreamerMoneyReq, opts ...client.CallOption) (*GetStreamerMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetStreamerMoney", in)
	out := new(GetStreamerMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetStreamerMoneyMulti(ctx context.Context, in *GetStreamerMoneyMultiReq, opts ...client.CallOption) (*GetStreamerMoneyMultiResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetStreamerMoneyMulti", in)
	out := new(GetStreamerMoneyMultiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) IncrStreamerMoney(ctx context.Context, in *IncrStreamerMoneyReq, opts ...client.CallOption) (*ChangeStreamerMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.IncrStreamerMoney", in)
	out := new(ChangeStreamerMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) DecrStreamerMoney(ctx context.Context, in *DecrStreamerMoneyReq, opts ...client.CallOption) (*ChangeStreamerMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.DecrStreamerMoney", in)
	out := new(ChangeStreamerMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) ListStreamerMoneyDetail(ctx context.Context, in *ListStreamerMoneyDetailReq, opts ...client.CallOption) (*ListStreamerMoneyDetailResp, error) {
	req := c.c.NewRequest(c.name, "Asset.ListStreamerMoneyDetail", in)
	out := new(ListStreamerMoneyDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetStreamerCoin(ctx context.Context, in *GetStreamerCoinReq, opts ...client.CallOption) (*GetStreamerCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetStreamerCoin", in)
	out := new(GetStreamerCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetStreamerCoinMulti(ctx context.Context, in *GetStreamerCoinMultiReq, opts ...client.CallOption) (*GetStreamerCoinMultiResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetStreamerCoinMulti", in)
	out := new(GetStreamerCoinMultiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) IncrStreamerCoin(ctx context.Context, in *IncrStreamerCoinReq, opts ...client.CallOption) (*ChangeStreamerCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.IncrStreamerCoin", in)
	out := new(ChangeStreamerCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) DecrStreamerCoin(ctx context.Context, in *DecrStreamerCoinReq, opts ...client.CallOption) (*ChangeStreamerCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.DecrStreamerCoin", in)
	out := new(ChangeStreamerCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) ListStreamerCoinDetail(ctx context.Context, in *ListStreamerCoinDetailReq, opts ...client.CallOption) (*ListStreamerCoinDetailResp, error) {
	req := c.c.NewRequest(c.name, "Asset.ListStreamerCoinDetail", in)
	out := new(ListStreamerCoinDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetStreamerCoinTotal(ctx context.Context, in *GetStreamerCoinTotalReq, opts ...client.CallOption) (*GetStreamerCoinTotalResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetStreamerCoinTotal", in)
	out := new(GetStreamerCoinTotalResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetUnionMoney(ctx context.Context, in *GetUnionMoneyReq, opts ...client.CallOption) (*GetUnionMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUnionMoney", in)
	out := new(GetUnionMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetUnionMoneyMulti(ctx context.Context, in *GetUnionMoneyMultiReq, opts ...client.CallOption) (*GetUnionMoneyMultiResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUnionMoneyMulti", in)
	out := new(GetUnionMoneyMultiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) IncrUnionMoney(ctx context.Context, in *IncrUnionMoneyReq, opts ...client.CallOption) (*ChangeUnionMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.IncrUnionMoney", in)
	out := new(ChangeUnionMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) DecrUnionMoney(ctx context.Context, in *DecrUnionMoneyReq, opts ...client.CallOption) (*ChangeUnionMoneyResp, error) {
	req := c.c.NewRequest(c.name, "Asset.DecrUnionMoney", in)
	out := new(ChangeUnionMoneyResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) ListUnionMoneyDetail(ctx context.Context, in *ListUnionMoneyDetailReq, opts ...client.CallOption) (*ListUnionMoneyDetailResp, error) {
	req := c.c.NewRequest(c.name, "Asset.ListUnionMoneyDetail", in)
	out := new(ListUnionMoneyDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetUnionCoin(ctx context.Context, in *GetUnionCoinReq, opts ...client.CallOption) (*GetUnionCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUnionCoin", in)
	out := new(GetUnionCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetUnionCoinMulti(ctx context.Context, in *GetUnionCoinMultiReq, opts ...client.CallOption) (*GetUnionCoinMultiResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetUnionCoinMulti", in)
	out := new(GetUnionCoinMultiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) IncrUnionCoin(ctx context.Context, in *IncrUnionCoinReq, opts ...client.CallOption) (*ChangeUnionCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.IncrUnionCoin", in)
	out := new(ChangeUnionCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) DecrUnionCoin(ctx context.Context, in *DecrUnionCoinReq, opts ...client.CallOption) (*ChangeUnionCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.DecrUnionCoin", in)
	out := new(ChangeUnionCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) ListUnionCoinDetail(ctx context.Context, in *ListUnionCoinDetailReq, opts ...client.CallOption) (*ListUnionCoinDetailResp, error) {
	req := c.c.NewRequest(c.name, "Asset.ListUnionCoinDetail", in)
	out := new(ListUnionCoinDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetMerchantCoin(ctx context.Context, in *GetMerchantCoinReq, opts ...client.CallOption) (*GetMerchantCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetMerchantCoin", in)
	out := new(GetMerchantCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetMerchantCoinMulti(ctx context.Context, in *GetMerchantCoinMultiReq, opts ...client.CallOption) (*GetMerchantCoinMultiResp, error) {
	req := c.c.NewRequest(c.name, "Asset.GetMerchantCoinMulti", in)
	out := new(GetMerchantCoinMultiResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) IncrMerchantCoin(ctx context.Context, in *IncrMerchantCoinReq, opts ...client.CallOption) (*ChangeMerchantCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.IncrMerchantCoin", in)
	out := new(ChangeMerchantCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) DecrMerchantCoin(ctx context.Context, in *DecrMerchantCoinReq, opts ...client.CallOption) (*ChangeMerchantCoinResp, error) {
	req := c.c.NewRequest(c.name, "Asset.DecrMerchantCoin", in)
	out := new(ChangeMerchantCoinResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) ListMerchantCoinDetail(ctx context.Context, in *ListMerchantCoinDetailReq, opts ...client.CallOption) (*ListMerchantCoinDetailResp, error) {
	req := c.c.NewRequest(c.name, "Asset.ListMerchantCoinDetail", in)
	out := new(ListMerchantCoinDetailResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Asset service

type AssetHandler interface {
	// 余额 money
	GetUserMoney(context.Context, *GetUserMoneyReq, *GetUserMoneyResp) error
	GetUserMoneyMulti(context.Context, *GetUserMoneyMultiReq, *GetUserMoneyMultiResp) error
	IncrUserMoney(context.Context, *IncrUserMoneyReq, *ChangeUserMoneyResp) error
	DecrUserMoney(context.Context, *DecrUserMoneyReq, *ChangeUserMoneyResp) error
	ListUserMoneyDetail(context.Context, *ListUserMoneyDetailReq, *ListUserMoneyDetailResp) error
	// 虚拟币coin
	GetUserCoin(context.Context, *GetUserCoinReq, *GetUserCoinResp) error
	GetUserCoinMulti(context.Context, *GetUserCoinMultiReq, *GetUserCoinMultiResp) error
	IncrUserCoin(context.Context, *IncrUserCoinReq, *ChangeUserCoinResp) error
	DecrUserCoin(context.Context, *DecrUserCoinReq, *ChangeUserCoinResp) error
	ListUserCoinDetail(context.Context, *ListUserCoinDetailReq, *ListUserCoinDetailResp) error
	// ---------------Streamer主播资产---------------
	// 余额 money
	GetStreamerMoney(context.Context, *GetStreamerMoneyReq, *GetStreamerMoneyResp) error
	GetStreamerMoneyMulti(context.Context, *GetStreamerMoneyMultiReq, *GetStreamerMoneyMultiResp) error
	IncrStreamerMoney(context.Context, *IncrStreamerMoneyReq, *ChangeStreamerMoneyResp) error
	DecrStreamerMoney(context.Context, *DecrStreamerMoneyReq, *ChangeStreamerMoneyResp) error
	ListStreamerMoneyDetail(context.Context, *ListStreamerMoneyDetailReq, *ListStreamerMoneyDetailResp) error
	// 虚拟币coin
	GetStreamerCoin(context.Context, *GetStreamerCoinReq, *GetStreamerCoinResp) error
	GetStreamerCoinMulti(context.Context, *GetStreamerCoinMultiReq, *GetStreamerCoinMultiResp) error
	IncrStreamerCoin(context.Context, *IncrStreamerCoinReq, *ChangeStreamerCoinResp) error
	DecrStreamerCoin(context.Context, *DecrStreamerCoinReq, *ChangeStreamerCoinResp) error
	ListStreamerCoinDetail(context.Context, *ListStreamerCoinDetailReq, *ListStreamerCoinDetailResp) error
	GetStreamerCoinTotal(context.Context, *GetStreamerCoinTotalReq, *GetStreamerCoinTotalResp) error
	// ---------------Union工会资产---------------
	// 余额 money --
	GetUnionMoney(context.Context, *GetUnionMoneyReq, *GetUnionMoneyResp) error
	GetUnionMoneyMulti(context.Context, *GetUnionMoneyMultiReq, *GetUnionMoneyMultiResp) error
	IncrUnionMoney(context.Context, *IncrUnionMoneyReq, *ChangeUnionMoneyResp) error
	DecrUnionMoney(context.Context, *DecrUnionMoneyReq, *ChangeUnionMoneyResp) error
	ListUnionMoneyDetail(context.Context, *ListUnionMoneyDetailReq, *ListUnionMoneyDetailResp) error
	// 虚拟币coin
	GetUnionCoin(context.Context, *GetUnionCoinReq, *GetUnionCoinResp) error
	GetUnionCoinMulti(context.Context, *GetUnionCoinMultiReq, *GetUnionCoinMultiResp) error
	IncrUnionCoin(context.Context, *IncrUnionCoinReq, *ChangeUnionCoinResp) error
	DecrUnionCoin(context.Context, *DecrUnionCoinReq, *ChangeUnionCoinResp) error
	ListUnionCoinDetail(context.Context, *ListUnionCoinDetailReq, *ListUnionCoinDetailResp) error
	// 虚拟币coin
	GetMerchantCoin(context.Context, *GetMerchantCoinReq, *GetMerchantCoinResp) error
	GetMerchantCoinMulti(context.Context, *GetMerchantCoinMultiReq, *GetMerchantCoinMultiResp) error
	IncrMerchantCoin(context.Context, *IncrMerchantCoinReq, *ChangeMerchantCoinResp) error
	DecrMerchantCoin(context.Context, *DecrMerchantCoinReq, *ChangeMerchantCoinResp) error
	ListMerchantCoinDetail(context.Context, *ListMerchantCoinDetailReq, *ListMerchantCoinDetailResp) error
}

func RegisterAssetHandler(s server.Server, hdlr AssetHandler, opts ...server.HandlerOption) error {
	type asset interface {
		GetUserMoney(ctx context.Context, in *GetUserMoneyReq, out *GetUserMoneyResp) error
		GetUserMoneyMulti(ctx context.Context, in *GetUserMoneyMultiReq, out *GetUserMoneyMultiResp) error
		IncrUserMoney(ctx context.Context, in *IncrUserMoneyReq, out *ChangeUserMoneyResp) error
		DecrUserMoney(ctx context.Context, in *DecrUserMoneyReq, out *ChangeUserMoneyResp) error
		ListUserMoneyDetail(ctx context.Context, in *ListUserMoneyDetailReq, out *ListUserMoneyDetailResp) error
		GetUserCoin(ctx context.Context, in *GetUserCoinReq, out *GetUserCoinResp) error
		GetUserCoinMulti(ctx context.Context, in *GetUserCoinMultiReq, out *GetUserCoinMultiResp) error
		IncrUserCoin(ctx context.Context, in *IncrUserCoinReq, out *ChangeUserCoinResp) error
		DecrUserCoin(ctx context.Context, in *DecrUserCoinReq, out *ChangeUserCoinResp) error
		ListUserCoinDetail(ctx context.Context, in *ListUserCoinDetailReq, out *ListUserCoinDetailResp) error
		GetStreamerMoney(ctx context.Context, in *GetStreamerMoneyReq, out *GetStreamerMoneyResp) error
		GetStreamerMoneyMulti(ctx context.Context, in *GetStreamerMoneyMultiReq, out *GetStreamerMoneyMultiResp) error
		IncrStreamerMoney(ctx context.Context, in *IncrStreamerMoneyReq, out *ChangeStreamerMoneyResp) error
		DecrStreamerMoney(ctx context.Context, in *DecrStreamerMoneyReq, out *ChangeStreamerMoneyResp) error
		ListStreamerMoneyDetail(ctx context.Context, in *ListStreamerMoneyDetailReq, out *ListStreamerMoneyDetailResp) error
		GetStreamerCoin(ctx context.Context, in *GetStreamerCoinReq, out *GetStreamerCoinResp) error
		GetStreamerCoinMulti(ctx context.Context, in *GetStreamerCoinMultiReq, out *GetStreamerCoinMultiResp) error
		IncrStreamerCoin(ctx context.Context, in *IncrStreamerCoinReq, out *ChangeStreamerCoinResp) error
		DecrStreamerCoin(ctx context.Context, in *DecrStreamerCoinReq, out *ChangeStreamerCoinResp) error
		ListStreamerCoinDetail(ctx context.Context, in *ListStreamerCoinDetailReq, out *ListStreamerCoinDetailResp) error
		GetStreamerCoinTotal(ctx context.Context, in *GetStreamerCoinTotalReq, out *GetStreamerCoinTotalResp) error
		GetUnionMoney(ctx context.Context, in *GetUnionMoneyReq, out *GetUnionMoneyResp) error
		GetUnionMoneyMulti(ctx context.Context, in *GetUnionMoneyMultiReq, out *GetUnionMoneyMultiResp) error
		IncrUnionMoney(ctx context.Context, in *IncrUnionMoneyReq, out *ChangeUnionMoneyResp) error
		DecrUnionMoney(ctx context.Context, in *DecrUnionMoneyReq, out *ChangeUnionMoneyResp) error
		ListUnionMoneyDetail(ctx context.Context, in *ListUnionMoneyDetailReq, out *ListUnionMoneyDetailResp) error
		GetUnionCoin(ctx context.Context, in *GetUnionCoinReq, out *GetUnionCoinResp) error
		GetUnionCoinMulti(ctx context.Context, in *GetUnionCoinMultiReq, out *GetUnionCoinMultiResp) error
		IncrUnionCoin(ctx context.Context, in *IncrUnionCoinReq, out *ChangeUnionCoinResp) error
		DecrUnionCoin(ctx context.Context, in *DecrUnionCoinReq, out *ChangeUnionCoinResp) error
		ListUnionCoinDetail(ctx context.Context, in *ListUnionCoinDetailReq, out *ListUnionCoinDetailResp) error
		GetMerchantCoin(ctx context.Context, in *GetMerchantCoinReq, out *GetMerchantCoinResp) error
		GetMerchantCoinMulti(ctx context.Context, in *GetMerchantCoinMultiReq, out *GetMerchantCoinMultiResp) error
		IncrMerchantCoin(ctx context.Context, in *IncrMerchantCoinReq, out *ChangeMerchantCoinResp) error
		DecrMerchantCoin(ctx context.Context, in *DecrMerchantCoinReq, out *ChangeMerchantCoinResp) error
		ListMerchantCoinDetail(ctx context.Context, in *ListMerchantCoinDetailReq, out *ListMerchantCoinDetailResp) error
	}
	type Asset struct {
		asset
	}
	h := &assetHandler{hdlr}
	return s.Handle(s.NewHandler(&Asset{h}, opts...))
}

type assetHandler struct {
	AssetHandler
}

func (h *assetHandler) GetUserMoney(ctx context.Context, in *GetUserMoneyReq, out *GetUserMoneyResp) error {
	return h.AssetHandler.GetUserMoney(ctx, in, out)
}

func (h *assetHandler) GetUserMoneyMulti(ctx context.Context, in *GetUserMoneyMultiReq, out *GetUserMoneyMultiResp) error {
	return h.AssetHandler.GetUserMoneyMulti(ctx, in, out)
}

func (h *assetHandler) IncrUserMoney(ctx context.Context, in *IncrUserMoneyReq, out *ChangeUserMoneyResp) error {
	return h.AssetHandler.IncrUserMoney(ctx, in, out)
}

func (h *assetHandler) DecrUserMoney(ctx context.Context, in *DecrUserMoneyReq, out *ChangeUserMoneyResp) error {
	return h.AssetHandler.DecrUserMoney(ctx, in, out)
}

func (h *assetHandler) ListUserMoneyDetail(ctx context.Context, in *ListUserMoneyDetailReq, out *ListUserMoneyDetailResp) error {
	return h.AssetHandler.ListUserMoneyDetail(ctx, in, out)
}

func (h *assetHandler) GetUserCoin(ctx context.Context, in *GetUserCoinReq, out *GetUserCoinResp) error {
	return h.AssetHandler.GetUserCoin(ctx, in, out)
}

func (h *assetHandler) GetUserCoinMulti(ctx context.Context, in *GetUserCoinMultiReq, out *GetUserCoinMultiResp) error {
	return h.AssetHandler.GetUserCoinMulti(ctx, in, out)
}

func (h *assetHandler) IncrUserCoin(ctx context.Context, in *IncrUserCoinReq, out *ChangeUserCoinResp) error {
	return h.AssetHandler.IncrUserCoin(ctx, in, out)
}

func (h *assetHandler) DecrUserCoin(ctx context.Context, in *DecrUserCoinReq, out *ChangeUserCoinResp) error {
	return h.AssetHandler.DecrUserCoin(ctx, in, out)
}

func (h *assetHandler) ListUserCoinDetail(ctx context.Context, in *ListUserCoinDetailReq, out *ListUserCoinDetailResp) error {
	return h.AssetHandler.ListUserCoinDetail(ctx, in, out)
}

func (h *assetHandler) GetStreamerMoney(ctx context.Context, in *GetStreamerMoneyReq, out *GetStreamerMoneyResp) error {
	return h.AssetHandler.GetStreamerMoney(ctx, in, out)
}

func (h *assetHandler) GetStreamerMoneyMulti(ctx context.Context, in *GetStreamerMoneyMultiReq, out *GetStreamerMoneyMultiResp) error {
	return h.AssetHandler.GetStreamerMoneyMulti(ctx, in, out)
}

func (h *assetHandler) IncrStreamerMoney(ctx context.Context, in *IncrStreamerMoneyReq, out *ChangeStreamerMoneyResp) error {
	return h.AssetHandler.IncrStreamerMoney(ctx, in, out)
}

func (h *assetHandler) DecrStreamerMoney(ctx context.Context, in *DecrStreamerMoneyReq, out *ChangeStreamerMoneyResp) error {
	return h.AssetHandler.DecrStreamerMoney(ctx, in, out)
}

func (h *assetHandler) ListStreamerMoneyDetail(ctx context.Context, in *ListStreamerMoneyDetailReq, out *ListStreamerMoneyDetailResp) error {
	return h.AssetHandler.ListStreamerMoneyDetail(ctx, in, out)
}

func (h *assetHandler) GetStreamerCoin(ctx context.Context, in *GetStreamerCoinReq, out *GetStreamerCoinResp) error {
	return h.AssetHandler.GetStreamerCoin(ctx, in, out)
}

func (h *assetHandler) GetStreamerCoinMulti(ctx context.Context, in *GetStreamerCoinMultiReq, out *GetStreamerCoinMultiResp) error {
	return h.AssetHandler.GetStreamerCoinMulti(ctx, in, out)
}

func (h *assetHandler) IncrStreamerCoin(ctx context.Context, in *IncrStreamerCoinReq, out *ChangeStreamerCoinResp) error {
	return h.AssetHandler.IncrStreamerCoin(ctx, in, out)
}

func (h *assetHandler) DecrStreamerCoin(ctx context.Context, in *DecrStreamerCoinReq, out *ChangeStreamerCoinResp) error {
	return h.AssetHandler.DecrStreamerCoin(ctx, in, out)
}

func (h *assetHandler) ListStreamerCoinDetail(ctx context.Context, in *ListStreamerCoinDetailReq, out *ListStreamerCoinDetailResp) error {
	return h.AssetHandler.ListStreamerCoinDetail(ctx, in, out)
}

func (h *assetHandler) GetStreamerCoinTotal(ctx context.Context, in *GetStreamerCoinTotalReq, out *GetStreamerCoinTotalResp) error {
	return h.AssetHandler.GetStreamerCoinTotal(ctx, in, out)
}

func (h *assetHandler) GetUnionMoney(ctx context.Context, in *GetUnionMoneyReq, out *GetUnionMoneyResp) error {
	return h.AssetHandler.GetUnionMoney(ctx, in, out)
}

func (h *assetHandler) GetUnionMoneyMulti(ctx context.Context, in *GetUnionMoneyMultiReq, out *GetUnionMoneyMultiResp) error {
	return h.AssetHandler.GetUnionMoneyMulti(ctx, in, out)
}

func (h *assetHandler) IncrUnionMoney(ctx context.Context, in *IncrUnionMoneyReq, out *ChangeUnionMoneyResp) error {
	return h.AssetHandler.IncrUnionMoney(ctx, in, out)
}

func (h *assetHandler) DecrUnionMoney(ctx context.Context, in *DecrUnionMoneyReq, out *ChangeUnionMoneyResp) error {
	return h.AssetHandler.DecrUnionMoney(ctx, in, out)
}

func (h *assetHandler) ListUnionMoneyDetail(ctx context.Context, in *ListUnionMoneyDetailReq, out *ListUnionMoneyDetailResp) error {
	return h.AssetHandler.ListUnionMoneyDetail(ctx, in, out)
}

func (h *assetHandler) GetUnionCoin(ctx context.Context, in *GetUnionCoinReq, out *GetUnionCoinResp) error {
	return h.AssetHandler.GetUnionCoin(ctx, in, out)
}

func (h *assetHandler) GetUnionCoinMulti(ctx context.Context, in *GetUnionCoinMultiReq, out *GetUnionCoinMultiResp) error {
	return h.AssetHandler.GetUnionCoinMulti(ctx, in, out)
}

func (h *assetHandler) IncrUnionCoin(ctx context.Context, in *IncrUnionCoinReq, out *ChangeUnionCoinResp) error {
	return h.AssetHandler.IncrUnionCoin(ctx, in, out)
}

func (h *assetHandler) DecrUnionCoin(ctx context.Context, in *DecrUnionCoinReq, out *ChangeUnionCoinResp) error {
	return h.AssetHandler.DecrUnionCoin(ctx, in, out)
}

func (h *assetHandler) ListUnionCoinDetail(ctx context.Context, in *ListUnionCoinDetailReq, out *ListUnionCoinDetailResp) error {
	return h.AssetHandler.ListUnionCoinDetail(ctx, in, out)
}

func (h *assetHandler) GetMerchantCoin(ctx context.Context, in *GetMerchantCoinReq, out *GetMerchantCoinResp) error {
	return h.AssetHandler.GetMerchantCoin(ctx, in, out)
}

func (h *assetHandler) GetMerchantCoinMulti(ctx context.Context, in *GetMerchantCoinMultiReq, out *GetMerchantCoinMultiResp) error {
	return h.AssetHandler.GetMerchantCoinMulti(ctx, in, out)
}

func (h *assetHandler) IncrMerchantCoin(ctx context.Context, in *IncrMerchantCoinReq, out *ChangeMerchantCoinResp) error {
	return h.AssetHandler.IncrMerchantCoin(ctx, in, out)
}

func (h *assetHandler) DecrMerchantCoin(ctx context.Context, in *DecrMerchantCoinReq, out *ChangeMerchantCoinResp) error {
	return h.AssetHandler.DecrMerchantCoin(ctx, in, out)
}

func (h *assetHandler) ListMerchantCoinDetail(ctx context.Context, in *ListMerchantCoinDetailReq, out *ListMerchantCoinDetailResp) error {
	return h.AssetHandler.ListMerchantCoinDetail(ctx, in, out)
}
