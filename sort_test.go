package mapsort

import (
	"testing"
)

func TestSort(t * testing.T) {
	s := []int{2,1,4,6,5,7,3}
	e := []int{1,2,3,4,5,6,7}
	Sort(s)
	for k,_ := range s{
		if s[k] != e[k]{
			t.Errorf("got:%v\nwant:%v\n",s,e)
			break
		}
	}
}