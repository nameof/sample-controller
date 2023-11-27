package operator

import (
	"fmt"
	"k8s.io/client-go/tools/cache"
)

type PrintHandler struct {
}

func (p PrintHandler) OnAdd(obj interface{}, isInInitialList bool) {
	key, _ := cache.MetaNamespaceKeyFunc(obj)
	fmt.Printf("OnAdd: %s\n", key)
}

func (p PrintHandler) OnUpdate(obj, newObj interface{}) {
	key, _ := cache.MetaNamespaceKeyFunc(obj)
	fmt.Printf("OnUpdate: %s\n", key)
}

func (p PrintHandler) OnDelete(obj interface{}) {
	key, _ := cache.MetaNamespaceKeyFunc(obj)
	fmt.Printf("OnDelete: %s\n", key)
}
