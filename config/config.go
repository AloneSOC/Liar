package config

import (
	"bufio"
	"net"
	"os"
	"strings"
)

func GetEnvInfo() (string, string, string, string) { 
   // zsh默认环境变量 PS1='[%n@%m %c]%$'
   // zsh default environment variables PS1='[%n@%m %c]%$'
	// sh&bash默认环境变量 PS1='[\u@\h \W]\$'
	// sh&bash default environment variables PS1='[\u@\h \W]\$'
	var EnvUser = os.Getenv("USER") 
	var EnvPWD  = os.Getenv("PWD")  
	var EnvHost, _ = os.Hostname()  
   var ExecArgs = strings.Join(os.Args, " ") 
   
	return EnvUser, EnvPWD, EnvHost, ExecArgs 
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
       // 检查IP地址以确定它是否为回环地址
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