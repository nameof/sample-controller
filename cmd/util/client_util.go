package util

import (
	"fmt"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
	"github.com/nameof/sample-controller/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"time"
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

func BuildInfo() *v1.GithubInfo {
	return &v1.GithubInfo{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "nameof.github.com/v1",
			Kind:       "GithubInfo",
		},
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
