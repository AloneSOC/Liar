package config

import (
	// "fmt"
	"bufio"
	"net"
	"os"
	"strings"
)


func GetEnvInfo() (string, string, string, string) { 
   //zsh default environment variables PS1='[%n@%m %c]%$'
	//sh&bash default environment variables PS1='[\u@\h \W]\$'
	var EnvUser = os.Getenv("USER") 
	var EnvPWD  = os.Getenv("PWD")  
	var EnvHost, _ = os.Hostname()  
   var ExecArgs = strings.Join(os.Args, " ") 
   
	return EnvUser, EnvPWD, EnvHost, ExecArgs 
}

func GetConfig() (string, string, string) { 
   // file, err := os.Open("/opt/Liar/config.ini")
   file, err := os.Open("config.ini")

   if err != nil {
      return "", "", "" 
   }
   defer file.Close()

   var ChatURL, ApiKey, Prompt string 

   scanner := bufio.NewScanner(file)

   for scanner.Scan() {
      line := scanner.Text()
      if strings.Contains(line, "chaturl") {
         startIndex := strings.Index(line, "=") + 1
         if startIndex > 0 {
            ChatURL = strings.TrimSpace(line[startIndex:])
         }
      }
      if strings.Contains(line, "apikey") {
         startIndex := strings.Index(line, "=") + 1
         if startIndex > 0 {
            ApiKey = strings.TrimSpace(line[startIndex:])
            }
      }
      if strings.Contains(line, "prompt") {
         startIndex := strings.Index(line, "=") + 1
         if startIndex > 0 {
            Prompt = strings.TrimSpace(line[startIndex:])      
            }
      }
   }
   return ChatURL, ApiKey, Prompt 
}

func GetBot() (string, string) { 
   file, err := os.Open("config.ini")
   if err != nil {
      return "", "" 
   }
   defer file.Close()

   var BotURL, BotJSON string 

   scanner := bufio.NewScanner(file)

   for scanner.Scan() {
      line := scanner.Text()
      if strings.Contains(line, "boturl") {
         startIndex := strings.Index(line, "=") + 1
         if startIndex > 0 {
            BotURL = strings.TrimSpace(line[startIndex:])
         }
      }
      if strings.Contains(line, "botjson") {
         startIndex := strings.Index(line, "=") + 1
         if startIndex > 0 {
            BotJSON = strings.TrimSpace(line[startIndex:])
         }
      }
   }
   return BotURL, BotJSON
}

func GetLocalIPs() []string {
   var LocalIPs []string 
   addrs, err := net.InterfaceAddrs()
   if err != nil {
       return LocalIPs
   }
   for _, address := range addrs {
       // Check the IP address to determine whether it is a loopback address
       if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
           if ipnet.IP.To4() != nil {
               LocalIPs = append(LocalIPs, ipnet.IP.String()) 
           }
       }
   }
   return LocalIPs 
}

func GetLocalIPsAsString() string {
   localIPs := GetLocalIPs() 
   return strings.Join(localIPs, ",") 
}