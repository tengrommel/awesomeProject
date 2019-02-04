package main

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test03()
}

// 重复定义
//type CInterface interface {
//	AInterface
//	BInterface
//}

func main() {

}
