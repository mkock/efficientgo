package efficientgo

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Project represents some kind of IT project.
type Project struct {
	Name   string    `json:"name"`
	Author string    `json:"author"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}

func marshalDirect(out io.Writer) {
	p := Project{"Efficient Go", "Martin Kock", time.Now(), time.Now()}
	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(out, j)
}

func marshalWriter(out io.Writer) {
	p := Project{"Efficient Go", "Martin Kock", time.Now(), time.Now()}
	e := json.NewEncoder(out)
	err := e.Encode(p)
	if err != nil {
		panic(err)
	}
}
