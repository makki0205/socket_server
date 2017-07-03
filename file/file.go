package file

import (
	"io/ioutil"
	"os"
	"fmt"
)

type fileRepository struct {
	BasePath string
	files []os.FileInfo
	FileCount int
}
// コンストラクタ
func NewfileRepository(basePath string)fileRepository{
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		panic("実行ファイルと同階層にassetsフォルダーを作ってください")
	}
	return fileRepository{
		BasePath:basePath,
		files:files,
		FileCount:len(files),
	}
}
// ファイルをprintする関数
func(self *fileRepository)Print() {
	index := 0
	for _, finfo := range self.files {
		fmt.Printf("%2d : %s\n", index, finfo.Name())
		index++
	}
}
func(self *fileRepository)GetFileString(i int)string{
	file, err := os.Open(self.BasePath + "/"+self.files[i].Name())
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, 10240)
	file.Read(buf)
	return string(buf)
}