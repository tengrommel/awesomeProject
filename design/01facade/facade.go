package facade

import "fmt"

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type aModuleImpl struct {
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type bModuleImpl struct {
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type API interface {
	Test() string
}
