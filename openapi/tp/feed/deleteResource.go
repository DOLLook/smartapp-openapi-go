package feed

import (
	"github.com/baidu/smartapp-openapi-go/utils"
)

// DeleteResourceRequest 请求结构体
type DeleteResourceRequest struct {
	AccessToken string      // 授权小程序的接口调用凭据
	AppID       interface{} // 小程序 Id
	Path        string      // 智能小程序内页链接
}

// 响应结构体

type DeleteResourceResponse struct {
	Data      bool   `json:"data"`       // 响应参数
	Errno     int64  `json:"errno"`      // 状态码
	ErrMsg    string `json:"msg"`        // 错误信息
	ErrorCode int64  `json:"error_code"` // openapi 错误码
	ErrorMsg  string `json:"error_msg"`  // openapi 错误信息
}

// DeleteResource
func DeleteResource(params *DeleteResourceRequest) (bool, error) {
	var (
		err        error
		defaultRet bool
	)
	respData := &DeleteResourceResponse{}

	client := utils.NewHTTPClient().
		SetContentType(utils.ContentTypeForm).
		SetConverterType(utils.ConverterTypeJSON).
		SetMethod("POST").
		SetScheme(utils.SCHEME).
		SetHost(utils.OPENAPIHOST).
		SetPath("/rest/2.0/smartapp/access/resource/delete")
	client.AddGetParam("access_token", params.AccessToken)
	client.AddGetParam("sp_sdk_ver", utils.SDKVERSION)
	client.AddGetParam("sp_sdk_lang", utils.SDKLANG)
	client.AddPostParam("app_id", params.AppID)
	client.AddPostParam("path", params.Path)

	err = client.Do()
	if err != nil {
		return defaultRet, err
	}
	err = client.Convert(respData)
	if err != nil {
		return defaultRet, err
	}
	if respData.ErrorCode != 0 {
		return defaultRet, &utils.OpenAPIError{respData.ErrorCode, respData.ErrorMsg, respData}
	}

	if respData.Errno != 0 {
		return defaultRet, &utils.APIError{respData.Errno, respData.ErrMsg, respData}
	}
	return respData.Data, nil
}
