package users

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"crmata-go/internal/adapters/http/shared"
	usersservice "crmata-go/internal/application/service/users"
	domain "crmata-go/internal/domain/user"
	userrepository "crmata-go/internal/domain/user/repository"
	"crmata-go/internal/helpers"
)

type Handler struct {
	service usersservice.Service
}

func New(service usersservice.Service) Handler {
	return Handler{service: service}
}

func (h Handler) UsersCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createUser(w, r)
	case http.MethodGet:
		h.listUsers(w, r)
	default:
		writeMethodNotAllowed(w, "GET, POST")
	}
}

func (h Handler) UsersResource(w http.ResponseWriter, r *http.Request) {
	userID, action, ok := parseUserRoute(r.URL.Path)
	if !ok {
		writeNotFound(w)
		return
	}

	switch action {
	case "resource":
		switch r.Method {
		case http.MethodGet:
			h.getUserByID(w, r, userID)
		case http.MethodPut:
			h.updateUser(w, r, userID)
		case http.MethodDelete:
			h.deleteUser(w, r, userID)
		default:
			writeMethodNotAllowed(w, "GET, PUT, DELETE")
		}
	case "group":
		if r.Method != http.MethodPatch {
			writeMethodNotAllowed(w, "PATCH")
			return
		}
		h.patchUserGroup(w, r, userID)
	default:
		writeNotFound(w)
	}
}

func (h Handler) handleUserError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidID),
		errors.Is(err, domain.ErrInvalidName),
		errors.Is(err, domain.ErrInvalidEmail),
		errors.Is(err, domain.ErrInvalidPassword),
		errors.Is(err, domain.ErrInvalidGroupID),
		errors.Is(err, domain.ErrInvalidActive):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, userrepository.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, userrepository.ErrEmailAlreadyExists):
		writeError(w, http.StatusConflict, err.Error())
	case errors.Is(err, userrepository.ErrGroupNotFound):
		writeError(w, http.StatusBadRequest, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, helpers.MsgInternalServerError)
	}
}

func parseUserRoute(path string) (int64, string, bool) {
	trimmed := strings.TrimPrefix(path, "/v1/users/")
	trimmed = strings.Trim(trimmed, "/")
	if trimmed == "" {
		return 0, "", false
	}

	parts := strings.Split(trimmed, "/")
	if len(parts) == 1 {
		id, ok := parsePositiveInt64(parts[0])
		return id, "resource", ok
	}

	if len(parts) == 2 && parts[1] == "group" {
		id, ok := parsePositiveInt64(parts[0])
		return id, "group", ok
	}

	return 0, "", false
}

func parsePositiveInt64(raw string) (int64, bool) {
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}

func writeSuccess(w http.ResponseWriter, status int, message string, data any) {
	shared.WriteSuccess(w, status, message, data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	shared.WriteError(w, status, message)
}

func writeMethodNotAllowed(w http.ResponseWriter, allowed string) {
	shared.WriteMethodNotAllowed(w, allowed)
}

func writeNotFound(w http.ResponseWriter) {
	shared.WriteNotFound(w)
}
