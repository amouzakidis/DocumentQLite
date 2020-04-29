package bfm

//PageManager contains chunks of data in memory
type PageManager struct {
	File               *BlockFile
	PageSize           int64
	NumOfPagesInMemory uint64
	pages              []Page
}

//Page describe a page of data in memory
type Page struct {
	data       []byte
	pageNumber uint64
}
