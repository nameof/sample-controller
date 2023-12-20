package operator

import (
	"testing"
)

var name = "nameof-in-github1700745933429"

func TestDynamicClientOperator_GetByName_DefaultConverter(t *testing.T) {
	operator := NewDynamicClientOperator()
	get, err := operator.GetByName(name)
	if err != nil || get.GetName() != name {
		t.Errorf("error")
	}
}

func TestDynamicClientOperator_GetByName_MyConverter(t *testing.T) {
	operator := NewDynamicClientOperator2(MyConverter)
	get, err := operator.GetByName(name)
	if err != nil || get.GetName() != name {
		t.Errorf("error")
	}
}
