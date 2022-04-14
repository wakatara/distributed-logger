package log

import (
    “fmt"
    "io"
    "io/ioutil"
    "os"
    "path"
    "sort"
    "strconv"
    "strings"
    "sync"

    api "github.com/travisjeffery/proglog/api/v1"
)

type Log struct {
    mu             sync.RWMutex
    Dir            string
    Config         Config

    activeSegment  *segment
    segments       []*segment
}

func NewLog(dir string, c Config) (*log, error) {
    if c.Segment.MaxStoreBytes == 0 {
        c.Segment.MaxStoreBytes = 1024
    }
  
    if c.Segment.MaxIndexBytes == 0 {
        c.Segment.MaxIndexBytes = 1024
    }
    l := &Log{
        Dir: dir,
        Config: c
    }
  
    return l, l.setup()
}


