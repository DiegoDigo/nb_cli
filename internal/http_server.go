package internal

import (
	"encoding/json"
	"fmt"
	"nb_client/config"
	"nb_client/internal/command/response"
	"net/http"
)

func HttpServer() {
	var resp []response.Revenda
	conf := config.GetFirebase()

	get, err := http.Get(conf.Url)
	if err != nil {
		return
	}

	defer get.Body.Close()

	dec := json.NewDecoder(get.Body)
	err = dec.Decode(&resp)
	if err != nil {
		fmt.Println("erro ao fazer a conversao")
	}

	fmt.Println(resp)

}
