package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"kacha-psp/kacha"
)

func MapKachaToPSPResponse(kachaResp map[string]interface{}, success bool) kacha.PSPResponse {
    data, _ := json.Marshal(kachaResp)

    ref := fmt.Sprint(kachaResp["trace_number"])
    id := fmt.Sprint(kachaResp["id"])
    msg := fmt.Sprint(kachaResp["detail"])
    status := "FAILURE"
    if success {
        status = "SUCCESS"
    }

    return kacha.PSPResponse{
        ReferenceID: ref,
        Status:      status,
        Message:     msg,
        PSPTxID:     id,
        PSPData:     string(data),
        Signature:   GenerateSignature(ref, msg, status),
    }
}

func GenerateSignature(reference, message, status string) string {
	raw := fmt.Sprintf("%s|%s|%s", reference, message, status)
	hash := sha256.Sum256([]byte(raw))
	return base64.StdEncoding.EncodeToString(hash[:])
}
