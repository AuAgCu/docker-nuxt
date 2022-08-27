package container

import (
	"reflect"
)

type DiContainer struct {
	containerInfo map[interface{}]containerInfo
}

func NewDiContainer() *DiContainer {
	return &DiContainer{containerInfo: make(map[interface{}]containerInfo)}
}

func (d *DiContainer) Register(interfaceType interface{}, newFunc interface{}) {
	t := reflect.TypeOf(newFunc)
	if t.Kind() != reflect.Func {
		return
	}

	value := reflect.ValueOf(newFunc)
	args := getFuncParams(t)

	d.containerInfo[interfaceType] = containerInfo{
		newFunc:  value,
		args:     args,
		instance: nil,
	}
}

func (d *DiContainer) getInstance(interfaceType interface{}) reflect.Value {
	info := d.containerInfo[interfaceType]
	if info.instance != nil {
		return *info.instance
	}

	var args []reflect.Value
	for _, v := range info.args {
		args = append(args, d.getInstance(v))
	}

	return info.newFunc.Call(args)[0]
}

// 戻り値を返す
func getFuncOutPutFirstRes(t reflect.Type) reflect.Type {
	// TODO: 戻り値がない場合エラーを返すようにする
	return t.Out(0)
}

// 引数
func getFuncParams(t reflect.Type) []reflect.Type {
	len := t.NumIn()
	in := make([]reflect.Type, len)
	for i := 0; i < len; i++ {
		in[i] = t.In(i)
	}
	return in
}

type containerInfo struct {
	newFunc  reflect.Value
	args     []reflect.Type
	instance *reflect.Value
}
