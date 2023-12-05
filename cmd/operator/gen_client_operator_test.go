package operator

import (
	"fmt"
	"github.com/nameof/sample-controller/cmd/util"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
	"testing"
	"time"
)

func createClient() *GenClientOperator {
	return NewGenClientOperator(util.CreateClientset())
}

func Test_Create_Ok(t *testing.T) {
	operator := createClient()
	info := buildInfo()
	operator.Create(info)

	get, _ := operator.GetByName(info.GetName())
	if get.GetName() != info.GetName() {
		t.Errorf("create %s, get %s", info.GetName(), get.GetName())
	}
	reflect.DeepEqual(get, info)
}

func Test_Create_Fail(t *testing.T) {
	operator := createClient()
	info := buildInfo()
	operator.Create(info)

	err := operator.Create(info)
	if !errors.IsAlreadyExists(err) {
		t.Errorf("error recreate %s", err)
	}
}

func buildInfo() *v1.GithubInfo {
	return &v1.GithubInfo{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s%d", "nameof-in-github", time.Now().UnixMilli()),
		},
		Spec: v1.GithubInfoSpec{
			Username:  "nameof",
			Link:      "https://github.com/nameof",
			RepoCount: 10,
		},
	}
}
