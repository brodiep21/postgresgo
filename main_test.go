package main_test

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize()
}
