package infras

import (
	"log"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/bep/debounce"
	"github.com/endaaman/api.endaaman.me/config"
	"github.com/radovskyb/watcher"
)

var watcher_mutex sync.Mutex
var ch = make(chan bool)

func notify() {
	logs.Info("Detect changes")
	ReadAllArticles()
	WaitIO()
	ch <- true
}

func AwaitNextChange() {
	logs.Info("Start awaiting next change and loading done")
	select {
	case <-ch:
		logs.Info("Load done by event triggered")
	case <-time.After(3 * time.Second):
	}
}

func StartWatching() {
	w := watcher.New()
	// w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Rename, watcher.Move, watcher.Write)
	// r := regexp.MustCompile("^abc$")
	// w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		debounced := debounce.New(time.Millisecond * 100)
		for {
			select {
			case <-w.Event:
				debounced(notify)
			case err := <-w.Error:
				logs.Error(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	articlesDir := config.GetArticlesDir()
	if err := w.AddRecursive(articlesDir); err != nil {
		log.Fatalln(err)
	}

	// for path, f := range w.WatchedFiles() {
	// 	fmt.Printf("%s: %s\n", path, f.Name())
	// }

	if err := w.Start(time.Millisecond * 300); err != nil {
		log.Fatalln(err)
	}
}
