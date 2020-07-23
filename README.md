# Moon_Trace
A tool for web safe

一个用于web安全的综合工具集。

# 使用
## 环境
需要go语言环境。第三方库使用go mod管理。
## 功能
### 子域扫描
支持：DNS数据集与ce证书透明度查询

语法： `go run moon_trace.go -u xxx.com -sub`

# 待完成功能 1.0
1. 子域扫描
- DNS数据集
    - dns.bufferover.run  
    - 百度观测云 x√
    - 站长之家
    - site.ip138.com
    - api.hackertarget.com
    - netcraft
    - sitedossier
    - threatcrowd
2. 端口扫描
3. 结果输出

## 暂时无法解决的问题
interface to []byte