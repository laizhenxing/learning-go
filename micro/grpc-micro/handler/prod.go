package handler

import (
	"context"
	"strconv"
	"time"

	"grpc-micro/proto/prod"
)

type ProdService struct {
}

// 获取单个商品信息
func (p ProdService) GetProdDetail(ctx context.Context, req *prod.ProdRequest, rsp *prod.ProdDetailResponse) error {
	time.Sleep(time.Second * 3)

	prod := NewProd(req.ProdId, "测试商品信息")
	rsp.Data = prod
	return nil
}

// 获取商品列表
func (p ProdService) GetProdList(ctx context.Context, req *prod.ProdRequest, rsp *prod.ProdResponse) error {
	time.Sleep(time.Second * 3)

	rsp.Data = generateProdList(req.Size)
	return nil
}

func NewProd(id int32, name string) *prod.ProdModel {
	return &prod.ProdModel{
		ProdID:   id,
		ProdName: name,
	}
}

func generateProdList(n int32) (data []*prod.ProdModel) {
	var i int32
	for i = 0; i < n; i++ {
		data = append(data, NewProd(i, "product"+strconv.Itoa(int(i))))
	}

	return
}

func DefaultProds() (rsp *prod.ProdResponse) {
	rsp = &prod.ProdResponse{}
	data := make([]*prod.ProdModel, 0)

	var i int32
	for i = 0; i < 5; i++ {
		data = append(data, NewProd(i, "Default product"+strconv.Itoa(999+int(i))))
	}

	rsp.Data = data

	return rsp
}

func DefaultProdsWithRsp(rsp interface{}) {
	res := rsp.(*prod.ProdResponse)
	res.Data = DefaultProds().Data
}

// 通用降级方法
func DefaultCommonProds(rsp interface{})  {
	switch t := rsp.(type) {
	case *prod.ProdResponse:
		t.Data = DefaultProds().Data
	case *prod.ProdDetailResponse:
		t.Data = NewProd(999, "单个-降级商品信息")
	}
}