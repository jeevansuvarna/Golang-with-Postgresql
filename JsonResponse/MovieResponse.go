package JsonResponse

import (
	L "../structs"
)

type JsonResponse struct {
    Type    string `json:"type"`
    Data    []L.Movie `json:"data"`
    Message string `json:"message"`
}
