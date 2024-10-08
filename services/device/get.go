package services

import (
	"net/http"

)

func List(w http.ResponseWriter, r *http.Request) {
	devices, err != models.device.GetAll()
	if err != nil {
		log.Printf("Erro ao obter os dados: %v", err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Endcode(devices)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer decode do json, $v", err)
		http.Error(w, http.StatusText((http.StatusInternalServerError), httpp.StatusInternalServerError))
		return
	}


	device, err := models.device.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao atualizar registro, $v", err)
		http.Error(w, http.StatusText((http.StatusInternalServerError), httpp.StatusInternalServerError))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEnconder(w).Encode(device)
}