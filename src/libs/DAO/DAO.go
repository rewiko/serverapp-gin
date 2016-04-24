package DAO

import "fmt"

type DAO interface {
	Select(interface{}) error
	Delete(interface{}) error
	Insert(interface{}, interface{}) error
	Update(interface{}, interface{}) error
}

type ManageDataMongoDb struct {
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func maintest() {
	var i I = T{"hello"}
	i.M()
}
