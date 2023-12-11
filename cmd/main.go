package main

import (
	"context"
	"fmt"
	"github.com/nameof/sample-controller/cmd/operator"
	"github.com/nameof/sample-controller/cmd/util"
	"github.com/nameof/sample-controller/pkg/client/clientset/versioned"
	"github.com/nameof/sample-controller/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OperationFunc func(clientset *versioned.Clientset)

func main() {
	factory := externalversions.NewSharedInformerFactory(util.CreateClientset(), 0)
	factory.Nameof().V1().GithubInfos().Informer().AddEventHandler(&operator.PrintHandler{})
	factory.WaitForCacheSync(nil)
	factory.Start(nil)
}

func createOne(clientset *versioned.Clientset) {
	info := util.BuildInfo()
	_, err := clientset.NameofV1().GithubInfos(metav1.NamespaceDefault).Create(context.Background(), info, metav1.CreateOptions{})
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
