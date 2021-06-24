package httpsrv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asasmoyo/demo-golang-buku-tamu/model"
	"github.com/go-chi/chi/v5"
)

func (s *Server) listTamu(w http.ResponseWriter, r *http.Request) {
	var tamus []model.Tamu
	if err := s.DB.Find(&tamus).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tamus); err != nil {
		fmt.Println(err)
	}
}

func (s *Server) createTamu(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("input is not valid"))
		return
	}

	name := r.PostForm.Get("name")
	keperluan := r.PostForm.Get("keperluan")

	tamu := model.Tamu{
		Name:      name,
		Keperluan: keperluan,
	}
	if err := s.DB.Create(&tamu).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(tamu); err != nil {
		fmt.Println(err)
	}
}

func (s *Server) deleteTamu(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("id '%s' is not valid", idStr)))
		return
	}

	if err := s.DB.Delete(model.Tamu{}, id).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("tamu is deleted"))
}
