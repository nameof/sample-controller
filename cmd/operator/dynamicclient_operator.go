package operator

import (
	"context"
	"github.com/nameof/sample-controller/cmd/util"
	v1 "github.com/nameof/sample-controller/pkg/apis/nameof.github.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type DynamicClientOperator struct {
	client        *dynamic.DynamicClient
	gvr           schema.GroupVersionResource
	converterFunc ConverterFuncType
}

func NewDynamicClientOperator() *DynamicClientOperator {
	return NewDynamicClientOperator2(DefaultConverter)
}

func NewDynamicClientOperator2(convert ConverterFuncType) *DynamicClientOperator {
	if convert == nil {
		convert = DefaultConverter
	}
	dynamicClient, err := dynamic.NewForConfig(util.GetConfig())
	if err != nil {
		panic(err)
	}
	return &DynamicClientOperator{client: dynamicClient, gvr: schema.GroupVersionResource{
		Group:    v1.Group,
		Version:  v1.Version,
		Resource: v1.ResourceName,
	},
		converterFunc: convert,
	}
}

func (d *DynamicClientOperator) Create(info *v1.GithubInfo) error {
	data := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "nameof.github.com/v1",
			"kind":       "GithubInfo",
			"metadata": map[string]interface{}{
				"name": info.GetName(),
			},
			"spec": map[string]interface{}{
				"username":  info.Spec.Username,
				"link":      info.Spec.Link,
				"repoCount": info.Spec.RepoCount,
			},
		},
	}
	_, err := d.client.Resource(d.gvr).Namespace(metav1.NamespaceDefault).Create(context.TODO(), data, metav1.CreateOptions{})
	return err
}

func (d *DynamicClientOperator) PrintAll() {
	list, err := d.client.Resource(d.gvr).Namespace(metav1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for index, item := range list.Items {
		Print(index, &item)
	}
}

func (d *DynamicClientOperator) GetByName(name string) (*v1.GithubInfo, error) {
	get, err := d.client.Resource(d.gvr).Namespace(metav1.NamespaceDefault).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	info := v1.GithubInfo{}
	err = d.converterFunc(get, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
