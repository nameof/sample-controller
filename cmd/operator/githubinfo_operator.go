package operator

import v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"

type GithubInfoOperator interface {
	Create(info *v1.GithubInfo) error
	PrintAll()
	GetByName(name string) (*v1.GithubInfo, error)
}
