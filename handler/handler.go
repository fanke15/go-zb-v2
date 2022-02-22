package handler

//true: 关闭tcp连接  false: 不关闭tcp连接
type Handler func([]byte) bool
