package v1

import (
	"net/http"
	"strconv"

	"goproposal/internal/api/v1/dtos"

	"github.com/gorilla/mux"
	"github.com/san-services/apicore"
	"github.com/san-services/apilogger"
)

func (h Handler) addPerson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := dtos.AddPersonInput{}
	logger := apilogger.New(ctx, "")

	// Parse body parameters.
	err := apicore.DecodeBody(r.Context(), r.Body, &params.Body)
	if err != nil {
		logger.Error(apilogger.LogCatUnmarshalReq, err)
		apicore.RespondError(ctx, w, http.StatusBadRequest, err)
		return
	}

	// Check for payload validation rules.
	err = h.validate.Struct(params)
	if err != nil {
		logger.Error(apilogger.LogCatInputValidation, err)
		apicore.RespondError(ctx, w, http.StatusBadRequest, err)
		return
	}

	err = h.service.AddPerson(params.Body.Name, params.Body.Age)
	if err != nil {
		logger.Error(apilogger.LogCatServiceOutput, err)
		apicore.RespondError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	apicore.RespondSuccess(ctx, w, "person added")
}

func (h Handler) persons(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := apilogger.New(ctx, "")

	pp, err := h.service.Persons()
	if err != nil {
		logger.Error(apilogger.LogCatServiceOutput, err)
		apicore.RespondError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	apicore.RespondSuccess(ctx, w, pp)
}

func (h Handler) person(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := apilogger.New(ctx, "")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error(apilogger.LogCatUnmarshalReq, err)
		apicore.RespondError(ctx, w, http.StatusBadRequest, err)
		return
	}

	p, err := h.service.Person(int64(id))
	if err != nil {
		logger.Error(apilogger.LogCatServiceOutput, err)
		apicore.RespondError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	apicore.RespondSuccess(ctx, w, p)
}
