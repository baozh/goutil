package callrecord

import (
	"os"
	"strings"
	"time"
	"fmt"

	co "github.com/baozh/goutil/concurrent"
	"github.com/baozh/goutil/system"
)


type Counter struct {
	callCounter 	co.AtomicUint64
	lastCallCounter uint64
	lineCount	uint32
	pid 		int
	file   		*os.File
	filePath	string
}

const (
	MaxLineCountPerFile = 200000
)

func NewCounter(path string, basename string) (*Counter, error) {
	c := &Counter{}
	c.callCounter.Set(0)
	c.lastCallCounter = 0
	c.lineCount = 0
	c.pid = os.Getpid()
	c.file = nil
	if (strings.TrimSpace(path) == "") {
		path = "."
	}
	if (strings.TrimSpace(basename) == "") {
		basename = time.Now().Format("2006-01-02 15:04:05")
	}
	c.filePath = path + "/" + basename + ".cr"
	return c , nil
}

func (c *Counter) init() bool {
	if system.IsExist(c.filePath) == true && system.IsDir(c.filePath) == false {
		fmt.Printf("[CallCounter] delete cr file, file:%s\n", c.filePath)
		if err := os.Remove(c.filePath); err != nil {
			fmt.Printf("[CallCounter] delete cr file failed! file:%s, err:%s\n",
				c.filePath, err.Error())
		}
	}

	if (system.IsDir(c.filePath) == true) {
		return false
	}

	if file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err == nil {
		c.file = file
	} else {
		fmt.Printf("[CallCounter] openfile failed! file:%s, err:%s", c.filePath, err.Error())
		return false
	}
	return true
}

func (c *Counter) IncreBy(num uint64) {
	c.callCounter.AddAndGet(num)
}

func (c *Counter) setCounter(num uint64) {
	c.callCounter.Set(num)
}

func (c *Counter) sync() {
	if c.file == nil {
		return
	}

	tmpCount := c.callCounter.Get()
	if (c.lastCallCounter == tmpCount) {
		return
	}

	c.lineCount++
	if (c.lineCount > MaxLineCountPerFile) {
		err := c.file.Close()
		if err != nil {
			fmt.Printf("[CallCounter] close file failed! file:%s, err:%s\n",
					c.file.Name(), err.Error())
		}

		derr := os.Remove(c.filePath)
		if derr != nil {
			fmt.Printf("[CallCounter] remove file failed! file:%s, err:%s\n",
					c.filePath, derr.Error())
		}

		c.lineCount = 0;

		if file, oerr := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err == nil {
			c.file = file
		} else {
			fmt.Printf("[CallCounter] openfile failed! file:%s, err:%s", c.filePath, oerr.Error())
			c.file = nil
			return
		}
	}

	c.file.WriteString(fmt.Sprintf("%d %s %d\n", c.pid, time.Now().Format("2006-01-02 15:04:05"), tmpCount))
	c.lastCallCounter = tmpCount
}

