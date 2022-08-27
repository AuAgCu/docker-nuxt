package container

import (
	"log"
	"reflect"
)

type DiContainer struct {
	containerInfo map[reflect.Type]containerInfo
}

func NewDiContainer() *DiContainer {
	return &DiContainer{containerInfo: make(map[reflect.Type]containerInfo)}
}

func (d *DiContainer) Register(newFunc interface{}) {
	t := reflect.TypeOf(newFunc)
	if t.Kind() != reflect.Func {
		return
	}

	value := reflect.ValueOf(newFunc)
	args := getFuncParams(t)
	out := getFuncOutPutFirstRes(t)

	d.containerInfo[out] = containerInfo{
		newFunc:  value,
		args:     args,
		instance: nil,
	}
}

type InterfaceFunc interface{}

func (d *DiContainer) GetInstance(f InterfaceFunc) reflect.Value {
	t := d.createInterfaceReflectType(f)
	return d.getInstance(t)
}

// ちょっとトリッキーだけど、funcの引数を使ってinterface型のreflectTypeを作るためのメソッド
// HACK: もっといい方法ない？
func (d *DiContainer) createInterfaceReflectType(f InterfaceFunc) reflect.Type {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return nil
	}

	out := getFuncParams(t)[0]

	return out
}

func (d *DiContainer) getInstance(t reflect.Type) reflect.Value {
	info := d.containerInfo[t]
	log.Println("t: ", t.Kind())

	if info.instance != nil {
		return *info.instance
	}

	var args []reflect.Value
	for _, v := range info.args {
		log.Println(v)
		log.Println(v.Kind())
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
