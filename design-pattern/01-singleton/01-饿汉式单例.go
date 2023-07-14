package _1_singleton

//必须是私有的，否则外部可以直接创建这个对象，单例就无从谈起
type singleton struct {
}

//提供一个唯一实例,但是这个指针永远不能改变
var instance *singleton = new(singleton)

//对外提供一个函数获取到单例对象
func GetInstance() *singleton {
	return instance
}
