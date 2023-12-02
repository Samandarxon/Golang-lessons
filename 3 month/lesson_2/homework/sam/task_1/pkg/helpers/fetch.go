package helpers

import (
	"app/modules"

	"github.com/go-zoox/fetch"
)

type Fetch struct {
	Obj modules.Monies
	Url string
}

func (f *Fetch) GetRequest() {
	response, err := fetch.Get(f.Url)
	if err != nil {
		panic(err)
	}
	err = response.UnmarshalJSON(&f.Obj.Objs)
	if err != nil {
		panic(err)
	}
}
