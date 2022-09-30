package handlers

import (
	"encoding/json"
	"net/http"
	dto "technicaltest/dto/result"
	"technicaltest/models"
	"technicaltest/repositories"
)

type dataHandler struct {
	DataRepository repositories.DataRepository
}

func HandleData(DataRepository repositories.DataRepository) *dataHandler {
	return &dataHandler{DataRepository}
}

func (h *dataHandler) CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// request := datadto.CreateDataRequest{}

	// validation := validator.New()
	// err := validation.Struct(request)

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	product := models.Data{
		Image: filename,
	}

	data, err := h.DataRepository.CreateData(product)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response := dto.ErrorResult{Code: http.StatusForbidden, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}
