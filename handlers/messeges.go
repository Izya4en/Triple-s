package handlers

import (
	"encoding/xml"
	"log"
	"net/http"
	"triple-s/arch"
)

func WriteXMLResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(statusCode)

	if err := xml.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func ErrXMLResponse(w http.ResponseWriter, code int, message string) {
	WriteXMLResponse(w, code, arch.ErrResp{Code: code, Message: message})
}
