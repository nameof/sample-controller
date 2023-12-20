package operator

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
)

type ConverterFuncType func(data *unstructured.Unstructured, result any) error

var DefaultConverter ConverterFuncType = ConvertToGithubInfo

func ConvertToGithubInfo(data *unstructured.Unstructured, result any) error {
	return runtime.DefaultUnstructuredConverter.FromUnstructured(data.UnstructuredContent(), result)
}

func MyConverter(data *unstructured.Unstructured, result any) error {
	t := reflect.TypeOf(result)
	value := reflect.ValueOf(result)
	if t.Kind() != reflect.Pointer || value.IsNil() {
		panic(fmt.Errorf("requires a non-nil pointer to an object, got %v", t))
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(marshal, result)
}
