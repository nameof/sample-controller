package operator

import (
	"github.com/nameof/sample-controller/cmd/util"
	"k8s.io/apimachinery/pkg/api/errors"
	"reflect"
	"testing"
)

func createClient() *GenClientOperator {
	return NewGenClientOperator(util.CreateClientset())
}

func Test_Create_Ok(t *testing.T) {
	operator := createClient()
	info := util.BuildInfo()
	operator.Create(info)

	get, _ := operator.GetByName(info.GetName())
	if get.GetName() != info.GetName() {
		t.Errorf("create %s, get %s", info.GetName(), get.GetName())
	}
	reflect.DeepEqual(get, info)
}

func Test_Create_Fail(t *testing.T) {
	operator := createClient()
	info := util.BuildInfo()
	operator.Create(info)

	err := operator.Create(info)
	if !errors.IsAlreadyExists(err) {
		t.Errorf("error recreate %s", err)
	}
}
