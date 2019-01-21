
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import mock "github.com/stretchr/testify/mock"

//
type Query struct {
	mock.Mock
}

//Done为给定字段提供模拟函数：
func (_m *Query) Done() {
	_m.Called()
}

//GetState提供了一个具有给定字段的模拟函数：命名空间、键
func (_m *Query) GetState(namespace string, key string) ([]byte, error) {
	ret := _m.Called(namespace, key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(namespace, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}