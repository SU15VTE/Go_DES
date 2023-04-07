# Go_DES
## GO_DES是用Go语言编写的一个DES加密解密的Cli程序。同时也是本学期《计算机网安安全原理》的实验一的源码。
## 如何使用
文件下包含了一个我在Windows10_x64平台下编译的可执行程序，你可以这样使用它：
``` bash
//加密
.\godes.exe encrypt --key "SU15VTE!" --text "Hello World!"
You key : SU15VTE!
You Text: Hello World!
Encrypt result:1c3568382f24163e7368342104541010
// 解密
.\godes.exe decrypt --key "SU15VTE!" --text "1c3568382f24163e7368342104541010"  
You key :       SU15VTE!
You Text:       1c3568382f24163e7368342104541010
Decrypt result: Hello World!
```
或是自己在自己的平台下编译文件然后再使用。
