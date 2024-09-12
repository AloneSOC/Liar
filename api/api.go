package api

import (
   "fmt"
   "strings"
   "net/http"
   "io"
   "encoding/json"

   "github.com/SECSpell/Disillusion/config"
)

// Choice 结构体定义了API响应中的单个选择
// Choice struct defines a single choice in the API response
type Choice struct {
   Message struct {
       Content string `json:"content"`
   } `json:"message"`
}

// Response 结构体定义了整个API响应
// Response struct defines the entire API response
type Response struct {
   Choices []Choice `json:"choices"`
}

func ChatGPTApi() {
   // 获取环境信息，使用 execargs 作为 payload
   // Get environment info, use execargs as payload
   _, _, _, execargs := config.GetEnvInfo()

   // 设置请求URL和方法
   // Set request URL and method
   url := "https://api.secspell.com/disillusion"
   method := "POST"

   // 使用 execargs 作为请求体
   // Use execargs as request body
   payload := strings.NewReader(execargs)

   // 创建HTTP客户端
   // Create HTTP client
   client := &http.Client{}

   // 创建新的HTTP请求
   // Create new HTTP request
   req, err := http.NewRequest(method, url, payload)
   if err != nil {
      fmt.Println("Failed to create request:", err)
      return
   }

   // 设置请求头
   // Set request headers
   req.Header.Add("Content-Type", "text/plain")

   // 发送请求并获取响应
   // Send request and get response
   res, err := client.Do(req)
   if err != nil {
      fmt.Println("Failed to send request:", err)
      return
   }
   defer res.Body.Close()

   // 读取响应体
   // Read response body
   body, err := io.ReadAll(res.Body)
   if err != nil {
      fmt.Println("Failed to read response:", err)
      return
   }

   // 解析JSON响应
   var response struct {
      Output []string `json:"output"`
   }
   err = json.Unmarshal(body, &response)
   if err != nil {
      // 解析错误时不打印任何内容
      return
   }

   // 只有在成功解析JSON后才打印输出
   for _, item := range response.Output {
      fmt.Println(item)
   }
}