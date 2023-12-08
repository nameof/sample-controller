package operator

import (
	"github.com/nameof/sample-controller/cmd/util"
	"k8s.io/apimachinery/pkg/api/errors"
	"reflect"
	"testing"
)

func TestRestClientOperator_Create_Ok(t *testing.T) {
	operator := NewRestClientOperator()
	info := util.BuildInfo()
	err := operator.Create(info)
	if err != nil {
		t.Errorf("create error : %s", err)
		return
	}

	get, err := operator.GetByName(info.GetName())
	if err != nil {
		t.Errorf("get error : %s", err)
		return
	}

	if get.GetName() != info.GetName() {
		t.Errorf("create %s, get %s", info.GetName(), get.GetName())
	}
	reflect.DeepEqual(get, info)
}

func TestRestClientOperator_Create_Fail(t *testing.T) {
	operator := NewRestClientOperator()
	info := util.BuildInfo()
	err := operator.Create(info)
	if err != nil {
		t.Errorf("create error : %s", err)
		return
	}

	err = operator.Create(info)
	if !errors.IsAlreadyExists(err) {
		t.Errorf("error recreate %s", err)
	}
}
