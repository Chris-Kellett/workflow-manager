package ping

import (
	"fmt"
	"io"
	"net/http"

	logger "github.com/Chris-Kellett/workflow-manager/Logger"
)

var responseBody string = "Ping!\n"

func Init() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		logger.Error("PING", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Ping: got / request\n")
	_, err := io.WriteString(w, responseBody)
	if err != nil {
		logger.Error("PING", err)
	}
}
