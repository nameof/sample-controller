package operator

import (
	"github.com/nameof/sample-controller/cmd/util"
	"k8s.io/apimachinery/pkg/api/errors"
	"testing"
)

func Test_Create_OK(t *testing.T) {
	tests := prepare()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := util.BuildInfo()
			err := tt.operator.Create(info)
			if err != nil {
				t.Errorf("create error : %s", err)
				return
			}

			get, _ := tt.operator.GetByName(info.GetName())
			if get.GetName() != info.GetName() {
				t.Errorf("create %s, get %s", info.GetName(), get.GetName())
			}
		})
	}
}

func Test_Create_Fail(t *testing.T) {
	tests := prepare()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info := util.BuildInfo()
			err := tt.operator.Create(info)
			if err != nil {
				t.Errorf("create error : %s", err)
				return
			}

			err = tt.operator.Create(info)
			if !errors.IsAlreadyExists(err) {
				t.Errorf("error recreate %s", err)
			}
		})
	}
}

func Test_PrintAll(t *testing.T) {
	tests := prepare()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operator.PrintAll()
		})
	}
}

func prepare() []struct {
	name     string
	operator GithubInfoOperator
} {
	tests := []struct {
		name     string
		operator GithubInfoOperator
	}{
		{"GenClientOperator", NewGenClientOperator()},
		{"RestClientOperator", NewRestClientOperator()},
		{"DynamicClientOperator", NewDynamicClientOperator()},
	}
	return tests
}
