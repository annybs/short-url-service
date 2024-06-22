package api

import (
	"io"
	"net/http"

	"github.com/annybs/ezdb"
	"github.com/annybs/go/rest"
	"github.com/annybs/go/validate"
	"github.com/rs/zerolog"
)

type API struct {
	DB  ezdb.Collection[[]byte]
	Log zerolog.Logger

	Token string
}

func (api *API) Delete(w http.ResponseWriter, req *http.Request) {
	if !rest.IsAuthenticated(req, api.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte{})
		return
	}

	path := req.URL.Path

	if exist, _ := api.DB.Has(path); !exist {
		w.Write([]byte{})
		return
	}

	if err := api.DB.Delete(path); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte{})
	}
}

func (api *API) Get(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	dest, _ := api.DB.Get(path)
	if len(dest) > 0 {
		w.Header().Set("Location", string(dest))
		w.WriteHeader(http.StatusPermanentRedirect)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write([]byte{})
}

func (api *API) Put(w http.ResponseWriter, req *http.Request) {
	if !rest.IsAuthenticated(req, api.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte{})
		return
	}

	path := req.URL.Path

	dest, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err := validate.URL(string(dest)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := api.DB.Put(path, dest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte{})
	}
}

func (api *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	api.Log.Info().Msgf("%s %s", req.Method, req.URL.Path)

	switch req.Method {
	case http.MethodGet:
		api.Get(w, req)
	case http.MethodPut:
		api.Put(w, req)
	case http.MethodDelete:
		api.Delete(w, req)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unsupported method"))
	}
}
