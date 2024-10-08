package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	models "github.com/LeonardoGrigolettoDev/fly-esp-server-go/models/device"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer decode do json, $v", err)
		http.Error(w, http.StatusText((http.StatusInternalServerError), httpp.StatusInternalServerError))
		return
	}

	var device models.Device
	err := json.NewDecoder(r.Body).Decode($device)
	if err != nil {
		log.Printf("Erro ao fazer decode do json, $v", err)
		http.Error(w, http.StatusText((http.StatusInternalServerError), httpp.StatusInternalServerError))
		return
	}
	row, arr := models.Update(int64(id), device)
	if err != nil {
		log.Printf("Erro ao atualizar registro, $v", err)
		http.Error(w, http.StatusText((http.StatusInternalServerError), httpp.StatusInternalServerError))
		return
	}
	if rows >1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}
	resp := map[string]any{
		"Error": false,
		"Message": "Dados atualizados com sucesso."
	}
}
