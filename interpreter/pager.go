package interpreter

type page []uint8

type pager struct {
	pageMap  map[int]page
	pagesize int
}

// newPager instantiates a pager.
func newPager() *pager {
	return &pager{
		pageMap:  make(map[int]page),
		pagesize: 500,
	}
}

// increment increments the value of a memory cell by 1.
func (pager *pager) increment(pointer int) {
	page := pager.findPage(pointer)
	pageAddress := pageAddress(pointer, pager.pagesize)
	page[pageAddress]++
}

// decrement decrements the value of a memory cell by 1.
func (pager *pager) decrement(pointer int) {
	page := pager.findPage(pointer)
	pageAddress := pageAddress(pointer, pager.pagesize)
	page[pageAddress]--
}

// getValue returns the value of a memory cell.
func (pager *pager) getValue(pointer int) uint8 {
	page := pager.findPage(pointer)
	pageAddress := pageAddress(pointer, pager.pagesize)
	return page[pageAddress]
}

// setValue sets the value of a memory cell.
func (pager *pager) setValue(pointer int, value uint8) {
	page := pager.findPage(pointer)
	pageAddress := pageAddress(pointer, pager.pagesize)
	page[pageAddress] = value
}

// findPage retrieves or creates the page containing the memory cell.
func (pager *pager) findPage(pointer int) page {
	pageIndex := pointer % pager.pagesize
	myPage, ok := pager.pageMap[pageIndex]
	if ok {
		return myPage
	}
	newPage := page(make([]uint8, pager.pagesize, pager.pagesize))
	pager.pageMap[pageIndex] = newPage
	return newPage
}

// pageAddress converts a memory pointer value to the address in
// the page containing the memory cell
func pageAddress(pointer int, pagesize int) int {
	pageAddress := pointer % pagesize
	if pageAddress < 0 {
		return pagesize + pageAddress
	}
	return pageAddress
}
