package watch

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/howeyc/fsnotify"
)

type watcher struct {
	tmp     chan struct{}
	watcher *fsnotify.Watcher
}

// go mod tidy && go run main.go
func (w *watcher) check() {
	w.clear()
	w.modTidy()
	w.runMain()
}
func (w *watcher) modTidy() {
	cmd := exec.Command("bash", "-c", "go mod tidy")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("go mod err")
	}
	fmt.Println(string(out))
}
func (w *watcher) runMain() {
	cmd := exec.Command("bash", "-c", "go run main.go run")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("go run err")
	}
	fmt.Println(string(out))
}
func (w *watcher) clear() {
	cmd := exec.Command("bash", "-c", "clear")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("clear err")
	}
	fmt.Println(string(out))
}

// add folder
func (w *watcher) addFolder(path string) {
	w.watcher.Watch(path)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		filePath := path + "/" + f.Name()
		if f.IsDir() {
			w.addFolder(filePath)
		}
	}
}
func (w *watcher) startWatch() {
	watcher, err := fsnotify.NewWatcher()
	w.watcher = watcher
	if err != nil {
		fmt.Println("fail to create new watcher")
		return
	}
	go func() {
		fmt.Println("开始监听")
		for {
			select {
			case e := <-watcher.Event:
				// fmt.Println("1111")
				if e.IsCreate() {
					w.addFolder("./")
					fmt.Println("2222222")
				}
				w.check()
			case err := <-watcher.Error:
				fmt.Printf(" %s\n", err.Error())
			}
		}
	}()
	//
	w.addFolder("./")
	<-w.tmp
}

func (w *watcher) Run() {
	w.runMain()
	w.startWatch()
}
