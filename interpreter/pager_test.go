package interpreter

import (
	"reflect"
	"testing"
)

func TestFindPageCreatesMissingPages(t *testing.T) {
	mp := 1000000
	pager := newPager()
	page := pager.findPage(mp)
	pageIndex := mp % pager.pagesize
	p, _ := pager.pageMap[pageIndex]
	if !reflect.DeepEqual(p, page) {
		t.Errorf("Wrong page")
	}
}

func TestDecrementMemoryValue(t *testing.T) {
	mp := 1000000
	pager := newPager()
	pager.decrement(mp)
	val := pager.getValue(mp)
	expected := uint8(255)
	if val != expected {
		t.Errorf("Got %v, expected %v", val, expected)
	}
}

func TestIncrementMemoryValue(t *testing.T) {
	mp := 1000000
	pager := newPager()
	pager.increment(mp)
	val := pager.getValue(mp)
	expected := uint8(1)
	if val != expected {
		t.Errorf("Got %v, expected %v", val, expected)
	}
}

func TestSetGetMemoryValue(t *testing.T) {
	mp := 1000000
	pager := newPager()
	expected := uint8(55)
	pager.setValue(mp, expected)
	val := pager.getValue(mp)
	if val != expected {
		t.Errorf("Got %v, expected %v", val, expected)
	}

}

func TestCalcPageAddressPageBoundary(t *testing.T) {
	pagesize := 5
	address := pageAddress(5, pagesize)
	expected := 0
	if address != expected {
		t.Errorf("Got %v, expected %v", address, expected)
	}
}

func TestCalcPageAddressPositivePointer(t *testing.T) {
	pagesize := 5
	address := pageAddress(6, pagesize)
	expected := 1
	if address != expected {
		t.Errorf("Got %v, expected %v", address, expected)
	}
}

func TestCalcPageAddressZeroPointer(t *testing.T) {
	pagesize := 5
	address := pageAddress(0, pagesize)
	expected := 0
	if address != expected {
		t.Errorf("Got %v, expected %v", address, expected)
	}
}

func TestCalcPageAddressNegativePointer(t *testing.T) {
	pagesize := 5
	address := pageAddress(-3, pagesize)
	expected := 2
	if address != expected {
		t.Errorf("Got %v, expected %v", address, expected)
	}
}
