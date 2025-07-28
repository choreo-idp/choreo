// Copyright 2025 The OpenChoreo Authors
// SPDX-License-Identifier: Apache-2.0

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/openchoreo/openchoreo/internal/openchoreo-api/models"
)

// writeSuccessResponse writes a successful API response
func writeSuccessResponse[T any](w http.ResponseWriter, statusCode int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.SuccessResponse(data)
	_ = json.NewEncoder(w).Encode(response) // Ignore encoding errors for response
}

// writeErrorResponse writes an error API response
func writeErrorResponse(w http.ResponseWriter, statusCode int, message, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := models.ErrorResponse(message, code)
	_ = json.NewEncoder(w).Encode(response) // Ignore encoding errors for response
}

// writeListResponse writes a paginated list response
func writeListResponse[T any](w http.ResponseWriter, items []T, total, page, pageSize int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := models.ListSuccessResponse(items, total, page, pageSize)
	_ = json.NewEncoder(w).Encode(response) // Ignore encoding errors for response
}
