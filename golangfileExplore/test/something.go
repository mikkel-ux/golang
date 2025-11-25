package test

type SampleTest struct{}

func NewSampleTest() *SampleTest {
	return &SampleTest{}
}

func (st *SampleTest) DoSomething(name string) string {
	return "Hello, " + name + "!"
}
