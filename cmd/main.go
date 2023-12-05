package main

import (
	"context"
	"fmt"
	"github.com/nameof/sample-controller/cmd/operator"
	"github.com/nameof/sample-controller/cmd/util"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
	"github.com/nameof/sample-controller/pkg/client/clientset/versioned"
	"github.com/nameof/sample-controller/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type OperationFunc func(clientset *versioned.Clientset)

func main() {
	factory := externalversions.NewSharedInformerFactory(util.CreateClientset(), 0)
	factory.Nameof().V1().GithubInfos().Informer().AddEventHandler(&operator.PrintHandler{})
	factory.WaitForCacheSync(nil)
	factory.Start(nil)
}

func createOne(clientset *versioned.Clientset) {
	info := v1.GithubInfo{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s%d", "nameof-in-github", time.Now().UnixMilli()),
		},
		Spec: v1.GithubInfoSpec{
			Username:  "nameof",
			Link:      "https://github.com/nameof",
			RepoCount: 10,
		},
	}

	_, err := clientset.NameofV1().GithubInfos(metav1.NamespaceDefault).Create(context.Background(), &info, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("create success")
}

func count(clientset *versioned.Clientset) int {
	list, err := clientset.NameofV1().GithubInfos(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return len(list.Items)
}
