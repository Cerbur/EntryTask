package img

import (
	"entrytask/http/config"
	"entrytask/http/controller"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GETHandler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("img")
	sprint := fmt.Sprint(config.Path, s)
	fileBytes, err := ioutil.ReadFile(sprint)
	if err != nil {
		controller.HTTPError(w, -403, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(fileBytes)
	return
}
