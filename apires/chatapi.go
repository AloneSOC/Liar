package apires

import (
   "fmt"
   "strings"
   "net/http"
   "io"
   "encoding/json"

   "github.com/AloneSOC/Liar/config" 

)

type Choice struct {
   Message struct {
       Content string `json:"content"`
   } `json:"message"`
}

type Response struct {
   Choices []Choice `json:"choices"`
}

func ChatGPTApi() {

   chaturl, apikey, prompt := config.GetConfig()
   _, _, _, execargs := config.GetEnvInfo() 
   method := "POST"

   payload := fmt.Sprintf(`{
      "model": "gpt-3.5-turbo",
      "messages": [{"role": "user", "content": "%s%s"}]
   }`, prompt, execargs)

   //Test use
   // fmt.Println(payload)

   client := &http.Client {}
   req, err := http.NewRequest(method, chaturl, strings.NewReader(payload))
   
   //Test use  
   // fmt.Println(req)

   if err != nil {
      return
   }
   req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apikey))
   req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
   req.Header.Add("Content-Type", "application/json")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := io.ReadAll(res.Body)
   if err != nil {
      return
   }
   //Test use  
   // fmt.Println(body)

   var promptRes Response
   json.Unmarshal([]byte(body), &promptRes)

   for _, choice := range promptRes.Choices {
       fmt.Println(choice.Message.Content)
   }


}