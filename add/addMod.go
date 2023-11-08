package add

import "os"

type addMod struct{}

func (am *addMod) addFile(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		panic("create folder err")
	}
	f, err := os.Create(name + "/mod.go")
	if err != nil {
		panic("create file error")
	}
	content := `package ` + name
	_, err = f.WriteString(content)
	if err != nil {
		panic("write error")
	}
}
func (am *addMod) Run(name string) {
	am.addFile(name)
}
