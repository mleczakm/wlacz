package tuya

import (
	"fmt"
	"github.com/tuya/tuya-cloud-sdk-go/api/common"
	"github.com/tuya/tuya-cloud-sdk-go/api/token"
)

func main() {
	// Use the SDK to make API calls
	client := common.NewClient("<your_access_id>", "<your_access_key>")
	resp, err := token.PostToken(client, "<your_username>", "<your_password>")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Access token:", resp.Result.AccessToken)
}
