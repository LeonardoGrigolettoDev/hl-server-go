package services

import (
	"fmt"
	"net/http"

	models "github.com/LeonardoGrigolettoDev/fly-esp-server-go/models/device"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	err := json.NewDecoder(r.Body).Decode(device)
	if err != nill {
		log.Printf("Erro ao fazer decode do json, $v", err)
		http.Error(w, http.StatusText((http.StatusInternalServerError), httpp.StatusInternalServerError))
		return
	}

	id, err := models.Insert(device)
	var res map[string]any
	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf(("Inserido com sucesso! ID: %d", id)),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	
	json.NewEnconder(w).Encode(resp)
}