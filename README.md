# 生成人脸128D特征值

## 使用方法：
1. 下载本程序的模型文件夹`testdata`：https://www.aliyundrive.com/s/4JNfjeoBJs4
2. 把同一人的多张照片都放到同一个文件夹下，比如放到了`face-images`文件夹里
3. 从本项目releases处下载编译好的二进制程序generate-face-128.exe，github访问较慢也可从此处下载：
https://www.aliyundrive.com/s/VQTwUeysrU3 运行如下命令:
  ```bash
  generate-face-128  -i  face-images  -m  testdata
  #-i 指定存放目标人脸图片的文件夹，-m 指定人脸识别的模型存放文件夹，模型放在了tesdata目录里。
 ```
  程序会自动计算出每张照片中人脸的128D数据, 然后计算出这些数据的平均值并输出, 分隔符是空格, 记得用的时候替换成逗号。

  每张照片中只能有一个人脸，这些人脸还必须都是同一个人, 不然算出来没啥意义。