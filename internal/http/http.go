package http

import (
	"gstat/internal/protocol"
)
import "net/http"

func Check(url string) protocol.Result {
	res := protocol.Result{
		Target:   url,
		Protocol: protocol.HTTP,
	}

	response, err := http.Get(url)
	if err != nil {
		res.Reachable = false
		res.Message = err.Error()
		return res
	}

	defer response.Body.Close()
	res.Reachable = true
	//body, err := io.ReadAll(response.Body)
	//if err == nil {
	//	res.Message = string(body)
	//}

	return res
}
