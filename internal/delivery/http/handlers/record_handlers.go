package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/cutecarti/db-ops/internal/delivery/http/dto"
	"github.com/cutecarti/db-ops/internal/delivery/http/request"
	"github.com/cutecarti/db-ops/internal/delivery/http/response"
	"github.com/cutecarti/db-ops/internal/models"
	"github.com/cutecarti/db-ops/internal/repository"
	"github.com/cutecarti/db-ops/internal/service"
)

type RecordHandler struct {
	service *service.DBService
}

func NewRecordHandler(service *service.DBService) *RecordHandler {
	return &RecordHandler{service: service}
}

// Create godoc
// @Summary Create record
// @Description Create new record
// @Tags records
// @Accept json
// @Produce json
// @Param request body dto.CreateRecordRequest true "record data"
// @Success 201 {object} dto.RecordResponse
// @Failure 400 {string} string "invalid request body"
// @Failure 500 {string} string "failed to create record"
// @Router /records [post]
func (h *RecordHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, err := request.DecodeJSON[dto.CreateRecordRequest](r)
	if err != nil {
		response.Error(w, 400, "invalid request body")
		return
	}

	if req.Name == "" {
		response.Error(w, 400, "name is required")
		return
	}

	record := models.Record{
		Name: req.Name,
	}

	id, err := h.service.CreateRecord(ctx, record)
	if err != nil {
		response.Error(w, 500, err.Error())
	}

	response.JSON(w, 201, dto.RecordResponse{
		ID:   id,
		Name: req.Name,
	})
}

func (h *RecordHandler) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, 200, "Home page")
}

// GetRecord godoc
// @Summary Get record by ID
// @Description Get record
// @Tags records
// @Produce json
// @Param id path int true "Record ID"
// @Success 200 {object} dto.RecordResponse
// @Failure 400 {string} string "invalid id"
// @Failure 404 {string} string "not found"
// @Router /records/{id} [get]
func (h *RecordHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	record, err := h.service.GetRecord(ctx, id)
	if err != nil {
		if errors.Is(repository.ErrRecordNotFound, err) {
			response.Error(w, 404, err.Error())
		}
		response.Error(w, 500, err.Error())
	}

	response.JSON(w, 200, dto.RecordResponse{
		ID:   record.ID,
		Name: record.Name,
	})

}

// DeleteRecord godoc
// @Summary Delete record
// @Description Delete record by ID
// @Tags records
// @Param id path int true "Record ID"
// @Success 204
// @Failure 400 {string} string "invalid id"
// @Failure 500 {string} string "error"
// @Router /records/{id} [delete]
func (h *RecordHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.Error(w, 400, "Id is required")
		return
	}

	err = h.service.DeleteRecord(ctx, id)
	if err != nil {
		if errors.Is(repository.ErrRecordNotFound, err) {
			response.Error(w, 404, err.Error())
		}
		response.Error(w, 500, err.Error())
		return
	}

	response.JSON(w, 204, "")
}

// UpdateRecord godoc
// @Summary Update record
// @Description Update record by ID
// @Tags records
// @Accept json
// @Produce json
// @Param id path int true "Record ID"
// @Param request body dto.UpdateRecordRequest true "record data"
// @Success 200
// @Failure 400 {string} string "invalid request body"
// @Failure 500 {string} string "failed to update record"
// @Router /records/{id} [put]
func (h *RecordHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response.Error(w, 400, "invalid id")
		return
	}
	req, err := request.DecodeJSON[dto.UpdateRecordRequest](r)
	if err != nil {
		response.Error(w, 400, "invalid request body")
		return
	}

	if req.Name == "" {
		response.Error(w, 400, "name is required")
		return
	}

	err = h.service.UpdateRecord(ctx, id, req.Name)
	if err != nil {
		if errors.Is(repository.ErrRecordNotFound, err) {
			response.Error(w, 404, err.Error())
		}
		response.Error(w, 500, err.Error())
		return
	}

}
