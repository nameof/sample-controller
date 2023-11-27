package main

import (
	"context"
	"fmt"
	"github.com/nameof/sample-controller/cmd/operator"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
	"github.com/nameof/sample-controller/pkg/client/clientset/versioned"
	"github.com/nameof/sample-controller/pkg/client/informers/externalversions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type OperationFunc func(clientset *versioned.Clientset)

func main() {
	clientset := createClient()

	factory := externalversions.NewSharedInformerFactory(clientset, 0)
	factory.Nameof().V1().GithubInfos().Informer().AddEventHandler(&operator.PrintHandler{})
	factory.WaitForCacheSync(nil)
	factory.Start(nil)

	printall(clientset)

	createOne(clientset)

	printall(clientset)

	var funcs []OperationFunc = []OperationFunc{printall, createOne}
	for _, i := range funcs {
		i(clientset)
	}
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

func createClient() *versioned.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}

func printall(clientset *versioned.Clientset) {
	list, err := clientset.NameofV1().GithubInfos(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for index, item := range list.Items {
		fmt.Printf("%d: GithubInfo(%s)\n", index+1, item.GetName())
	}
}

func count(clientset *versioned.Clientset) int {
	list, err := clientset.NameofV1().GithubInfos(metav1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return len(list.Items)
}
