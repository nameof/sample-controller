package util

import (
	"github.com/nameof/sample-controller/pkg/client/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateClientset() *versioned.Clientset {
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
