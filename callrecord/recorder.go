package callrecord

import (
	"sync"
	"strings"
	"github.com/baozh/goutil/system"
	"os"
	"fmt"
	"time"
)

var (
	DefaultRecorder = NewRecorder()
)

func init() {
	DefaultRecorder.Init()
}

func GetCounter(tag string) *Counter{
	return DefaultRecorder.GetCounter(tag)
}

type Recorder struct {
	filePath 	string
	counters	map[string]*Counter
	mu		sync.Mutex
}

func NewRecorder() *Recorder {
	return &Recorder {
		filePath: ".",
		counters: make(map[string]*Counter),
	}
}

//SetFilePath set the file path of recorder. It need called before Init().
func (r *Recorder) SetFilePath(path string) {
	r.filePath = path
}

func (r *Recorder) GetFilePath() string {
	return r.filePath
}

func (r *Recorder) Init() bool {
	if strings.TrimSpace(r.filePath) == "" {
		r.filePath = "."
	}

	if system.IsExist(r.filePath) == false {
		if err := os.Mkdir(r.filePath, os.ModeDir | os.ModePerm); err != nil {
			fmt.Printf("[CallRecorder] init(mkdir) failed! filePath:%s, err:%s",
				r.filePath, err.Error())
			return false
		}
	}

	//start timer
	go func () {
		t := time.NewTicker(time.Second)
		defer t.Stop()

		for {
			select {
			case <- t.C:
				r.mu.Lock()
				for _,v := range r.counters{
					v.sync()
				}
				r.mu.Unlock()
			}
		}
	}()
	return true
}

func (r *Recorder) GetCounter(tag string) *Counter{
	var cc *Counter = nil
	r.mu.Lock()

	tmpCC, ok := r.counters[tag]
	if (ok) {
		cc = tmpCC
		r.mu.Unlock()
	} else {
		cc, _ = NewCounter(r.filePath, tag)
		cc.init()
		r.counters[tag] = cc
		r.mu.Unlock()
	}
	return cc
}

