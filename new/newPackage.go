package new

import (
	"fmt"
	"os"
	"os/exec"
)

type newPackage struct{}

func (np *newPackage) Run(name string) {
	np.CreateWorkSpace()
	np.createProject(name)
	np.addMod(name)
	np.addWork(name)
}

// create folder/ , main.go , .git/
func (np *newPackage) createProject(name string) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		panic("create folder err")
	}
	f, err := os.Create(name + "/main.go")
	if err != nil {
		panic("create file error")
	}
	content := `package main
import "fmt"
func main(){
	fmt.Println("Ciallo～(∠・ω< )⌒★")
}
	`
	_, err = f.WriteString(content)
	if err != nil {
		panic("write error")
	}
	os.Chdir("./" + name)
	cmd := exec.Command("bash", "-c", "git init")
	err = cmd.Run()
	if err != nil {
		panic("git init error")
	}
}

// go work use **
func (np *newPackage) addWork(name string) {
	os.Chdir("../")
	cmd := exec.Command("bash", "-c", "go work use ./"+name)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		panic("go work use error")
	}
}

// go mod init
func (np *newPackage) addMod(name string) {

	os.Chdir(name)
	cmd := exec.Command("bash", "-c", "go mod init "+name)
	err := cmd.Run()
	if err != nil {
		panic("go mod init error")
	}
}

// go work init
func (np *newPackage) CreateWorkSpace() {
	_, err := os.Stat("go.work")
	if os.IsNotExist(err) {
		cmd := exec.Command("bash", "-c", "go work init")
		err := cmd.Run()
		if err != nil {
			panic("go work init error")
		}
	}
}
