package operator

import (
	"context"
	"fmt"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
	"github.com/nameof/sample-controller/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GenClientOperator struct {
	client *versioned.Clientset
}

func NewGenClientOperator(client *versioned.Clientset) *GenClientOperator {
	return &GenClientOperator{client: client}
}

func (g *GenClientOperator) Create(info *v1.GithubInfo) error {
	_, err := g.client.NameofV1().GithubInfos(metav1.NamespaceDefault).Create(context.Background(), info, metav1.CreateOptions{})
	return err
}

func (g *GenClientOperator) PrintAll() {
	list, err := g.client.NameofV1().GithubInfos(metav1.NamespaceDefault).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for index, item := range list.Items {
		fmt.Printf("%d: GithubInfo(%s)\n", index+1, item.GetName())
	}
}

func (g *GenClientOperator) GetByName(name string) (*v1.GithubInfo, error) {
	return g.client.NameofV1().GithubInfos(metav1.NamespaceDefault).Get(context.Background(), name, metav1.GetOptions{})
}
