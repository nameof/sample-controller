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

func Test_Count(t *testing.T) {
	tests := prepare()
	counts := make([]int, len(tests))
	for i, tt := range tests {
		counts[i] = tt.operator.Count()
	}

	for i := 1; i < len(counts); i++ {
		if counts[i] != counts[0] {
			t.Errorf("count error %d %d", counts[i], counts[0])
		}
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
