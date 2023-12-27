package operator

import (
	"fmt"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
)

type GithubInfoOperator interface {
	Create(info *v1.GithubInfo) error
	PrintAll()
	GetByName(name string) (*v1.GithubInfo, error)
	Count() int
}

type ResourceNameGetter interface {
	GetName() string
}

func Print(index int, getter ResourceNameGetter) {
	fmt.Printf("%d: GithubInfo(%s)\n", index+1, getter.GetName())
}
