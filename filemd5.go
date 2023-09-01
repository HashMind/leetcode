package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main5() {
	imageFile, _ := os.Open("/Users/badboy/Downloads/scenario/cetc54all_v1.2_arm64_scene.tar")
	imageFileMd5 := md5.New()
	io.Copy(imageFileMd5, imageFile)
	sum := imageFileMd5.Sum(nil)
	fmt.Println(hex.EncodeToString(sum))

}
