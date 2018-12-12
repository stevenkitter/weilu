package wxcrypter

import (
	"errors"
)

//WechatErrorCode WechatErrorCode
type WechatErrorCode int

const (
	//ValidateSignatureErrorCode ValidateSignatureErrorCode
	ValidateSignatureErrorCode WechatErrorCode = -40001 - iota
	//ParseXMLErrorCode ParseXMLErrorCode
	ParseXMLErrorCode
	//ComputeSignatureErrorCode ComputeSignatureErrorCode
	ComputeSignatureErrorCode
	//IllegalAesKeyCode IllegalAesKeyCode
	IllegalAesKeyCode
	//ValidateAppIDErrorCode ValidateAppIDErrorCode
	ValidateAppIDErrorCode
	//EncryptAESErrorCode EncryptAESErrorCode
	EncryptAESErrorCode
	//DecryptAESErrorCode DecryptAESErrorCode
	DecryptAESErrorCode
	//IllegalBufferCode IllegalBufferCode
	IllegalBufferCode
	//EncodeBase64ErrorCode EncodeBase64ErrorCode
	EncodeBase64ErrorCode
	//DecodeBase64ErrorCode DecodeBase64ErrorCode
	DecodeBase64ErrorCode
	//GenReturnXMLErrorCode GenReturnXMLErrorCode
	GenReturnXMLErrorCode
)

var (
	//ErrorValidateSignature ErrorValidateSignature
	ErrorValidateSignature = errors.New("ErrorValidateSignature")
	//ErrorParseXML ErrorParseXML
	ErrorParseXML = errors.New("ErrorParseXML")
	//ErrorComputeSignature ErrorComputeSignature
	ErrorComputeSignature = errors.New("ErrorComputeSignature")
	//ErrorIllegalAesKey ErrorIllegalAesKey
	ErrorIllegalAesKey = errors.New("ErrorIllegalAesKey")
	//ErrorValidateAppID ErrorValidateAppID
	ErrorValidateAppID = errors.New("ErrorValidateAppID")
	//ErrorEncryptAES ErrorEncryptAES
	ErrorEncryptAES = errors.New("ErrorEncryptAES")
	//ErrorDecryptAES ErrorDecryptAES
	ErrorDecryptAES = errors.New("ErrorDecryptAES")
	//ErrorIllegalBuffer ErrorIllegalBuffer
	ErrorIllegalBuffer = errors.New("ErrorIllegalBuffer")
	//ErrorEncodeBase64 ErrorEncodeBase64
	ErrorEncodeBase64 = errors.New("ErrorEncodeBase64")
	//ErrorDecodeBase64 ErrorDecodeBase64
	ErrorDecodeBase64 = errors.New("ErrorDecodeBase64")
	//ErrorGenReturnXML ErrorGenReturnXML
	ErrorGenReturnXML = errors.New("ErrorGenReturnXML")
)

func errorToCode(err1 error) (code WechatErrorCode, err error) {
	switch err1 {
	case ErrorValidateSignature:
		code = ValidateSignatureErrorCode
	case ErrorParseXML:
		code = ParseXMLErrorCode
	case ErrorComputeSignature:
		code = ComputeSignatureErrorCode
	case ErrorIllegalAesKey:
		code = IllegalAesKeyCode
	case ErrorValidateAppID:
		code = ValidateAppIDErrorCode
	case ErrorEncryptAES:
		code = EncryptAESErrorCode
	case ErrorDecryptAES:
		code = DecryptAESErrorCode
	case ErrorIllegalBuffer:
		code = IllegalBufferCode
	case ErrorEncodeBase64:
		code = EncodeBase64ErrorCode
	case ErrorDecodeBase64:
		code = DecodeBase64ErrorCode
	case ErrorGenReturnXML:
		code = GenReturnXMLErrorCode
	default:
		err = errors.New("Not found Code")
	}
	return
}
