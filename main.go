package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/Kagami/go-face"
)

func main() {
	var imgDir string
	var modelDir string
	flag.StringVar(&imgDir, "i", `face-images`, "图片文件夹路径")
	flag.StringVar(&modelDir, "m", `testdata`, "模型文件夹路径")
	flag.Parse()
	getAverageFaceDescription(modelDir, imgDir)
}

// 从同一个人的多张图片中识别人脸并计算出平均的人脸特征值，每个图片只能有一张脸。
func getAverageFaceDescription(modelDir string, faceImgDir string) {
	modelsPath := filepath.Join(modelDir, "models")

	rec, err := face.NewRecognizer(modelsPath)
	if err != nil {
		log.Println("Can't init face recognizer:", err)
		return
	}
	defer rec.Close()

	imgDir, err := os.ReadDir(faceImgDir)
	if err != nil {
		log.Fatal(err)
	}
	faceDescriptions := []face.Descriptor{}
	for _, cFile := range imgDir {
		if cFile.IsDir() { // 忽略子目录
			continue
		}
		absPath, err := filepath.Abs(faceImgDir + cFile.Name())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("now reading file:", absPath)
		faces, err := rec.RecognizeFile(absPath)
		if err != nil {
			log.Println("Can't recognize:", err)
			return
		}
		if len(faces) != 1 {
			log.Println("img:", cFile.Name(), "has wrong number of faces")
			return
		}
		faceDescriptions = append(faceDescriptions, faces[0].Descriptor)
	}

	num := len(faceDescriptions)
	avgFaceDescriptions := face.Descriptor{}
	for i := 0; i < 128; i++ {
		for j := 0; j < num; j++ {
			avgFaceDescriptions[i] = avgFaceDescriptions[i] + faceDescriptions[j][i]
		}
		avgFaceDescriptions[i] = avgFaceDescriptions[i] / float32(num)

	}

	log.Println(avgFaceDescriptions)
}
