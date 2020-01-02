package contacts

import (
	configcat "gopkg.in/configcat/go-sdk.v1"
	"net/http"
	"os"
)

func List(w http.ResponseWriter, r *http.Request) {
	configCatAPIKey := os.Getenv("CONFIGCAT_APIKEY")
	client := configcat.NewClient(configCatAPIKey)

	isEnabled, _ := client.GetValue("enabled", false).(bool)
	//isVersion1, _ := client.GetValue("version1", false).(bool)
	isVersion2, _ := client.GetValue("version2", false).(bool)

	if !isEnabled {
		w.Write([]byte("Not Implemented"))
		w.WriteHeader(501)
		return
	}

	if isVersion2 {
		w = responseV2(w)
		w.WriteHeader(200)
		return
	}

	w = responseV1(w)
	w.WriteHeader(200)
	return
}

func responseV1(w http.ResponseWriter) http.ResponseWriter {
	w.Write([]byte("Enabled"))
	return w
}

func responseV2(w http.ResponseWriter) http.ResponseWriter {
	w.Write([]byte("{\"accounts\": [\"A\", \"B\"]}"))
	return w
}
