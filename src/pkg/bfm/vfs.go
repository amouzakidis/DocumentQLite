package bfm

import (
	"os"
)

//BlockFile is a representation of a BlockFile
type BlockFile struct {
	fileName  string
	fp        *os.File
	blockSize int64
	err       error
}

//CreateFile creates an empty BlockFile at he disc
func (blockFile BlockFile) CreateFile() {
	blockFile.fp, blockFile.err = os.Create(blockFile.fileName)
}

//OpenFile open a file, if it does not exist it will be created
func (blockFile BlockFile) OpenFile() {
	_, err := os.Stat("file-exists.go")
	if os.IsExist(err) {
		blockFile.fp, blockFile.err = os.Open(blockFile.fileName)
	} else {
		blockFile.CreateFile()
	}
}

//WriteBytes writes a byte slice at the index i * blockSize
func (blockFile BlockFile) WriteBytes(data []byte, index int64) {
	blockFile.fp.Seek(blockFile.blockSize*index, 0)
	blockFile.fp.Write(data)
}

//ReadBytes reads a byte slice from the index i * blockSize
func (blockFile BlockFile) ReadBytes(data []byte, index int64) int {
	blockFile.fp.Seek(blockFile.blockSize*index, 0)
	var n int
	n, blockFile.err = blockFile.fp.Read(data)
	return n
}
