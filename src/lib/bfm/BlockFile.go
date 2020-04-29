package bfm

import (
	"os"
)

//BlockFile is a representation of a BlockFile
type BlockFile struct {
	FileName  string
	Fp        *os.File
	BlockSize int64
	Err       error
}

//CreateFile creates an empty BlockFile at he disc
func (blockFile BlockFile) CreateFile() {
	blockFile.Fp, blockFile.Err = os.Create(blockFile.FileName)
}

//OpenFile open a file, if it does not exist it will be created
func (blockFile BlockFile) OpenFile() {
	_, err := os.Stat("file-exists.go")
	if os.IsExist(err) {
		blockFile.Fp, blockFile.Err = os.Open(blockFile.FileName)
	} else {
		blockFile.CreateFile()
	}
}

//WriteBytes writes a byte slice at the index i * blockSize
func (blockFile BlockFile) WriteBytes(data []byte, index int64) {
	blockFile.Fp.Seek(blockFile.BlockSize*index, 0)
	blockFile.Fp.Write(data)
}

//ReadBytes reads a byte slice from the index i * blockSize
func (blockFile BlockFile) ReadBytes(data []byte, index int64) int {
	blockFile.Fp.Seek(blockFile.BlockSize*index, 0)
	var n int
	n, blockFile.Err = blockFile.Fp.Read(data)
	return n
}
