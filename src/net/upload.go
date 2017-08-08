package net

import (
	"net/http"
	"ImageServer/src/file"
	"fmt"
	"os"
	"io"
	"io/ioutil"
	_"github.com/gorilla/mux"
	"strings"
)

//上传文件
func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	r.ParseMultipartForm(32 << 20)
	f, _, err := r.FormFile("uploadFile")
	if f == nil {
		w.Write([]byte("uploadFile参数为空"))
		return
	}
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error:upload error"))
		return
	}
	fileType, right := file.CheckFileType(f)
	if !right {
		w.Write([]byte("上传文件不合法，请上传jpg,png或gif格式的文件"))
		return
	}
	//随机生成一个不存在的fileId
	var imgId string
	var imgPath string
	for {
		imgId = file.MakeImgId()
		fmt.Println(imgId)
		imgPath = file.ImageID2Path(imgId, fileType)
		fmt.Println(imgPath)
		if !file.FileExist(imgPath) {
			break
		}
	}

	//回绕文件指针
	if _, err := f.Seek(0, 0); err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: exception"))
		return
	}

	if err = file.BuildTree(imgId); err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: exception"))
		return
	}

	imgFile, err := os.OpenFile(imgPath, os.O_WRONLY | os.O_CREATE, os.ModePerm | os.ModeTemporary)
	defer imgFile.Close()
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error: File created fail"))
		return
	}
	io.Copy(imgFile, f)
	resultUrl := file.SplitPath(imgPath)
	//w.Write([]byte(imgId))
	w.Write([]byte(resultUrl))
}

func LoadImg(w http.ResponseWriter, r *http.Request) {
	fmt.Println("图片")
	fmt.Println(r.URL.Path)
	rPath := r.URL.Path
	//path := r.FormValue("path")
	paths := strings.Split(rPath, "/img/")
	fmt.Println(paths)
	path := paths[len(paths) - 1]
	/*f, err := os.Open("imgTest/0A/82/32/A7/A5/EE/1D/61.png")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error:File open fail"))
		return
	}*/
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error:File open fail"))
		return
	}
	w.Write(f)
}

