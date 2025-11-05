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
		if msg == "" {
			msg = "Process service request successfully."
		}
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

func MapPushUSSDToPSP(resp *kacha.PushUSSDResponse, success bool) kacha.PSPResponse {
	status := "FAILURE"
	message := "Failed to process service request."
	if success && resp.Status == "PENDING" {
		status = "SUCCESS"
		message = "Process service request successfully."
	} else if success {
		status = "SUCCESS"
		message = resp.Message
	}

	pspDataXML := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Body>
        <api:Result xmlns:api="http://cps.huawei.com/cpsinterface/api_resultmgr" xmlns:res="http://cps.huawei.com/cpsinterface/result">
            <res:Header>
                <res:Version>1.0</res:Version>
                <res:OriginatorConversationID>%s</res:OriginatorConversationID>
            </res:Header>
            <res:Body>
                <res:ResultType>0</res:ResultType>
                <res:ResultCode>0</res:ResultCode>
                <res:ResultDesc>%s</res:ResultDesc>
                <res:QueryTransactionStatusResult>
                    <res:TransactionStatus>%s</res:TransactionStatus>
                </res:QueryTransactionStatusResult>
            </res:Body>
        </api:Result>
    </soapenv:Body>
</soapenv:Envelope>`,
		resp.TraceNumber,
		message,
		resp.Status,
	)

	return kacha.PSPResponse{
		ReferenceID: resp.TraceNumber,
		Status:      status,
		Message:     message,
		PSPTxID:     resp.TraceNumber,
		PSPData:     pspDataXML,
		Signature:   GenerateSignature(resp.TraceNumber, message, status),
	}
}

func MapTransferToPSP(resp *kacha.TransferResponse, success bool) kacha.PSPResponse {
	data, _ := json.Marshal(resp)
	status := "FAILURE"
	message := resp.Message
	if success {
		status = "SUCCESS"
		if message == "" {
			message = "Process service request successfully."
		}
	}

	return kacha.PSPResponse{
		ReferenceID: resp.Reference,
		Status:      status,
		Message:     message,
		PSPTxID:     resp.TransactionID,
		PSPData:     string(data),
		Signature:   GenerateSignature(resp.Reference, message, status),
	}
}

func GenerateSignature(reference, message, status string) string {
	raw := fmt.Sprintf("%s|%s|%s", reference, message, status)
	hash := sha256.Sum256([]byte(raw))
	return base64.StdEncoding.EncodeToString(hash[:])
}
