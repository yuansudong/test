package codes

import (
	"io/fs"
	"log"
	"path/filepath"
)

// ZIP 一个压缩类
type ZIP struct {
}

// LoadFile 用于加载文件
func LoadFile(iPath string) {

	filepath.Walk(iPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		log.Println(info.Name(), info.Size(), info.Sys())
		if info.IsDir() {
			log.Println("此处是目录")
		}
		return nil
	})
	// ioutil.ReadDir(dirname string)

}
