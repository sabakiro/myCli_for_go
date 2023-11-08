package watch

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/howeyc/fsnotify"
)

type watcher struct {
	tmp chan struct{}
}

// go mod tidy && go run main.go
func (w *watcher) check() {
	fmt.Println("1111")
	cmd := exec.Command("bash", "-c", "go mod tidy")
	err := cmd.Run()
	if err != nil {
		fmt.Println("go mod err")
	}
	cmd = exec.Command("bash", "-c", "go run main.go run")
	err = cmd.Run()
	if err != nil {
		fmt.Println("go run err")
	}
}

func (w *watcher) startWatch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("fail to create new watcher")
		return
	}
	go func() {
		fmt.Println("开始监听")
		for {
			select {
			case <-watcher.Event:
				fmt.Println("1111")
				// w.check()
			case err := <-watcher.Error:
				fmt.Printf(" %s\n", err.Error())
			}
		}
	}()
	//
	err = watcher.Watch("./")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	<-w.tmp
}

func (w *watcher) Run() {

	w.startWatch()
}
