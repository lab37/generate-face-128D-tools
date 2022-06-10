# generate-face-128D-tools
生成人脸128D特征值
- 使用方法：
```bash
cd ~/go/src/
git clone  https://github.com/lab37/generate-face-128D-tools.git
cd generate-face-128D-tools/
把照片都放到这个目录下的face-images文件夹里, 然后在generate-face-128D-tools目录运行:
go get
go run *.go

程序会自动计算出每张照片中人脸的128D数据, 然后计算出这些数据的平均值并输出, 分隔符是空格, 记得用的时候替换成逗号。
每张照片中只能有一个人脸，这些人脸还必须都是同一个人, 不然算出来没啥意义。
```