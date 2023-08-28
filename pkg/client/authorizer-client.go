package client

import (
	"fmt"
	"net/http"

	"github.com/marcosvscampos/mt-transfer-worker/pkg/utils"
)

const authorizerUrl string = "https://run.mocky.io/v3/8fafdd68-a090-496f-8c9a-3442cf30dae6"

func CallAuthorizer() {
	resp, err := http.Get(authorizerUrl)

	if err != nil {
		fmt.Println("An error has occurred during the call of [", authorizerUrl, "] -", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		utils.PrintJsonResponse(*resp)
	}

}
