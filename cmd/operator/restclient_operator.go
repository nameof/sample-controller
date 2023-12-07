package operator

import (
	"context"
	"fmt"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	crd = "githubinfos"
)

type RestClientOperator struct {
	client *rest.RESTClient
}

func NewRestClientOperator() *RestClientOperator {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "nameof.github.com", Version: "v1"}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.UnversionedRESTClientFor(&crdConfig)
	if err != nil {
		panic(err)
	}

	return &RestClientOperator{client: client}
}

func (r *RestClientOperator) Create(info *v1.GithubInfo) error {
	result := v1.GithubInfo{}
	return r.client.Post().Namespace(v12.NamespaceDefault).Resource(crd).Body(info).Do(context.TODO()).Into(&result)
}

func (r *RestClientOperator) PrintAll() {
	list := v1.GithubInfoList{}
	err := r.client.Get().Namespace(v12.NamespaceDefault).Resource(crd).Do(context.TODO()).Into(&list)
	if err != nil {
		panic(err)
	}

	for index, item := range list.Items {
		fmt.Printf("%d: GithubInfo(%s)\n", index+1, item.GetName())
	}
}

func (r *RestClientOperator) GetByName(name string) (*v1.GithubInfo, error) {
	info := v1.GithubInfo{}
	err := r.client.Get().Namespace(v12.NamespaceDefault).Resource(crd).Name(name).Do(context.TODO()).Into(&info)
	return &info, err
}
