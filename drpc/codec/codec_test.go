package codec_test

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/osgochina/dmicro/drpc/codec"
	"testing"
)

type TestCodec struct{}

const (
	NameTest = "test"
	IdTest   = 't'
)

func (TestCodec) ID() byte {
	return IdTest
}

func (TestCodec) Name() string {
	return NameTest
}

func (TestCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (TestCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func TestReg(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		reg := new(TestCodec)
		codec.Reg(reg)
		ret, err := codec.Get(IdTest)
		t.Assert(err, nil)
		t.Assert(ret, reg)
		ret, err = codec.GetByName(NameTest)
		t.Assert(err, nil)
		t.Assert(ret, reg)
	})
}

func TestGet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ret, err := codec.Get(codec.IdJson)
		t.Assert(err, nil)
		t.AssertNE(ret, nil)

		ret, err = codec.Get(codec.NilCodecID)
		t.AssertNE(err, nil)
		t.Assert(ret, nil)
	})
}

func TestGetByName(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ret, err := codec.GetByName(codec.NameJson)
		t.Assert(err, nil)
		t.AssertNE(ret, nil)

		ret, err = codec.GetByName(codec.NilCodecName)
		t.AssertNE(err, nil)
		t.Assert(ret, nil)
	})
}

func TestMarshal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := g.Map{
			"abc": "efg",
		}
		res, err := codec.Marshal(codec.IdJson, v)
		t.Assert(err, nil)
		t.Assert(res, gconv.Bytes("{\"abc\":\"efg\"}"))
	})
}

func TestUnmarshal(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v := g.Map{}
		err := codec.Unmarshal(codec.IdJson, gconv.Bytes("{\"abc\":\"efg\"}"), &v)
		t.Assert(err, nil)
		t.Assert(v["abc"], "efg")
	})
}
