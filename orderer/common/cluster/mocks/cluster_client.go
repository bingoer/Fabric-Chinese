
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import orderer "github.com/hyperledger/fabric/protos/orderer"

//clusterClient是为clusterClient类型自动生成的模拟类型
type ClusterClient struct {
	mock.Mock
}

//步骤为给定字段提供模拟函数：ctx、in、opts
func (_m *ClusterClient) Step(ctx context.Context, in *orderer.StepRequest, opts ...grpc.CallOption) (*orderer.StepResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *orderer.StepResponse
	if rf, ok := ret.Get(0).(func(context.Context, *orderer.StepRequest, ...grpc.CallOption) *orderer.StepResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orderer.StepResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *orderer.StepRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//Submit提供了一个具有给定字段的模拟函数：ctx、opts
func (_m *ClusterClient) Submit(ctx context.Context, opts ...grpc.CallOption) (orderer.Cluster_SubmitClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orderer.Cluster_SubmitClient
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) orderer.Cluster_SubmitClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orderer.Cluster_SubmitClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}