package file

import (
	"github.com/seehuhn/mt19937"
	"time"
	"encoding/binary"
	"strings"
	"encoding/hex"
	"os"
	"fmt"
	"mime/multipart"
	"net/http"
)

/*
首先在usr目录创建图片目录，用chmod -R 777命令设置权限，然后用ln命令链接到go项目的目录，
ln -s /usr/imgTest/ /home/memory/goworkspace/src/ImageServer/imgTest，这样就
可以把/usr/imgTest目录映射到项目目录
 */

var imgType = [4]string{"image/png", "image/jpg", "image/gif", "image/jpeg"}

var spilePath = "/home/memory/goworkspace/src/ImageServer/"

var basePath string = spilePath + "imgTest"
//var basePath string = "/usr/imgSrc"

//创建一个不重复的ImgId
func MakeImgId() string {
	mt := mt19937.New()
	mt.Seed(time.Now().UnixNano())
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, mt.Uint64())
	return strings.ToUpper(hex.EncodeToString(buf))
}

//检查文件名是否存在
func FileExist(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		return false
	} else {
		return true
	}

}

func ImageID2Path(imageid string, fileType string)string{
	temp := strings.Split(fileType, "/")
	ext := temp[len(temp) -1]
	fmt.Println(ext)
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s.%s", basePath,imageid[0:2],imageid[2:4],imageid[4:6],imageid[6:8],imageid[8:10],imageid[10:12],imageid[12:14],imageid[14:16], ext)
}

//检查文件类型是否为指定的类型
func CheckFileType(file multipart.File) (string, bool) {
	//检测文件类型
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	fileType := http.DetectContentType(buff)
	fmt.Println(fileType)
	for _,item := range imgType {
		if fileType == item {
			return fileType, true
		}

	}
	return "", false
}

//创建目录树
func BuildTree(imageId string) error {
	path := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", basePath,imageId[0:2],imageId[2:4],imageId[4:6],imageId[6:8],imageId[8:10],imageId[10:12],imageId[12:14])
	fmt.Println("要生成的path", path)
	return os.MkdirAll(path, os.ModePerm)
}

//截取文件的路径
func SplitPath(path string) string {
	strs := strings.Split(path, spilePath)
	fmt.Println(strs)
	return strs[len(strs) - 1]
}
