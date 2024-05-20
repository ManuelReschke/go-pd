package pd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/imroc/req/v3"
)

const (
	Name             = "PixelDrain.com"
	BaseURL          = "https://pixeldrain.com/"
	APIURL           = BaseURL + "api"
	DefaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36"
	// errors
	ErrMissingPathToFile = "file path or file reader is required"
	ErrMissingFileID     = "file id is required"
	ErrMissingFilename   = "if you use ReadCloser you need to specify the filename"
)

type ErrorMessage struct {
	ErrorCode    int    `json:"error_code" xml:"ErrorCode"`
	ErrorMessage string `json:"error_message" xml:"ErrorMessage"`
}

type ClientOptions struct {
	Debug             bool
	ProxyURL          string
	EnableCookies     bool
	EnableInsecureTLS bool
	Timeout           time.Duration
}

type ClientWrapper struct {
	Client         *req.Client
	UploadCallback func(info req.UploadInfo)
}

type PixelDrainClient struct {
	Client *ClientWrapper
	Debug  bool
}

// New - create a new PixelDrainClient
func New(opt *ClientOptions, c *ClientWrapper) *PixelDrainClient {
	// set default values if no other options available
	if opt == nil {
		opt = &ClientOptions{
			Debug:             false,
			ProxyURL:          "",
			EnableCookies:     true,
			EnableInsecureTLS: true,
			Timeout:           1 * time.Hour,
		}
	}

	// build default client if not available
	if c == nil {
		client := req.C()
		client.SetUserAgent(DefaultUserAgent)

		c = &ClientWrapper{
			Client:         client,
			UploadCallback: nil,
		}
	}

	if opt.EnableInsecureTLS {
		c.Client.EnableInsecureSkipVerify()
	}

	if opt.Timeout != 0 {
		c.Client.SetTimeout(opt.Timeout)
	}

	if opt.ProxyURL != "" {
		_ = c.Client.SetProxyURL(opt.ProxyURL)
	}
	// cookie is available by default in v3
	if opt.EnableCookies == false {
		c.Client.SetCookieJar(nil)
	}

	pdc := &PixelDrainClient{
		Client: c,
		Debug:  opt.Debug,
	}

	return pdc
}

func (pd *PixelDrainClient) SetUploadCallback(callback func(info req.UploadInfo)) {
	pd.Client.UploadCallback = callback
}

// UploadPOST POST /api/file
// curl -X POST -i -H "Authorization: Basic <TOKEN>" -F "file=@cat.jpg" https://pixeldrain.com/api/file
func (pd *PixelDrainClient) UploadPOST(r *RequestUpload) (*ResponseUpload, error) {
	if r.PathToFile == "" && r.File == nil {
		return nil, errors.New(ErrMissingPathToFile)
	}

	if r.URL == "" {
		r.URL = fmt.Sprint(APIURL + "/file")
	}

	if r.File != nil {
		if r.FileName == "" {
			return nil, errors.New(ErrMissingFilename)
		}
	} else {
		file, err := os.Open(r.PathToFile)
		defer file.Close()
		if err != nil {
			return nil, err
		}
	}

	headers := map[string]string{}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() && !r.Anonymous {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var uploadRsp ResponseUpload
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetQueryParam("anonymous", strconv.FormatBool(r.Anonymous)).
		SetSuccessResult(&uploadRsp).
		SetErrorResult(&errMsg).
		// currently a bug - https://github.com/imroc/req/issues/352
		// SetFileReader("file", r.GetFileName(), r.File).
		SetFile("file", r.PathToFile).
		SetUploadCallback(pd.Client.UploadCallback).
		Post(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	uploadRsp.StatusCode = rsp.GetStatusCode()

	return &uploadRsp, nil
}

// UploadPUT PUT /api/file/{name}
// curl -X PUT -i -H "Authorization: Basic <TOKEN>" --upload-file cat.jpg https://pixeldrain.com/api/file/test_cat.jpg
func (pd *PixelDrainClient) UploadPUT(r *RequestUpload) (*ResponseUpload, error) {
	if r.PathToFile == "" && r.File == nil {
		return nil, errors.New(ErrMissingPathToFile)
	}

	if r.File == nil && r.FileName == "" {
		return nil, errors.New(ErrMissingFilename)
	}

	file, err := os.Open(r.PathToFile)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	r.File = file

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/file/%s", r.GetFileName())
	}

	headers := map[string]string{}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() && !r.Anonymous {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var uploadRsp ResponseUpload
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		// SetQueryParam("anonymous", strconv.FormatBool(r.Anonymous)).
		SetSuccessResult(&uploadRsp).
		SetErrorResult(&errMsg).
		SetBody(r.File).
		SetUploadCallback(pd.Client.UploadCallback).
		Put(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	uploadRsp.StatusCode = rsp.GetStatusCode()

	return &uploadRsp, nil
}

// Download GET /api/file/{id}
func (pd *PixelDrainClient) Download(r *RequestDownload) (*ResponseDownload, error) {
	if r.PathToSave == "" {
		return nil, errors.New(ErrMissingPathToFile)
	}

	if r.ID == "" {
		return nil, errors.New(ErrMissingFileID)
	}

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/file/%s", r.ID)
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetErrorResult(&errMsg).
		SetOutputFile(r.PathToSave).
		Get(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	if rsp.GetStatusCode() != 200 {
		defaultRsp := &ResponseDefault{}
		defaultRsp.StatusCode = rsp.GetStatusCode()
		defaultRsp.Success = false
		downloadRsp := &ResponseDownload{
			ResponseDefault: *defaultRsp,
		}

		return downloadRsp, nil
	}

	fInfo, err := os.Stat(r.PathToSave)
	if err != nil {
		return nil, err
	}

	downloadRsp := &ResponseDownload{
		FilePath: r.PathToSave,
		FileName: fInfo.Name(),
		FileSize: fInfo.Size(),
		ResponseDefault: ResponseDefault{
			StatusCode: rsp.GetStatusCode(),
			Success:    true,
		},
	}

	return downloadRsp, nil
}

// GetFileInfo GET /api/file/{id}/info
func (pd *PixelDrainClient) GetFileInfo(r *RequestFileInfo) (*ResponseFileInfo, error) {
	if r.ID == "" {
		return nil, errors.New(ErrMissingFileID)
	}

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/file/%s/info", r.ID)
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var fileInfoRsp ResponseFileInfo
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetSuccessResult(&fileInfoRsp).
		SetErrorResult(&errMsg).
		Get(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	fileInfoRsp.StatusCode = rsp.GetStatusCode()
	if fileInfoRsp.StatusCode == http.StatusOK {
		fileInfoRsp.Success = true
	}

	return &fileInfoRsp, nil
}

// DownloadThumbnail GET /api/file/{id}/thumbnail?width=x&height=x
func (pd *PixelDrainClient) DownloadThumbnail(r *RequestThumbnail) (*ResponseThumbnail, error) {
	if r.PathToSave == "" {
		return nil, errors.New(ErrMissingPathToFile)
	}

	if r.ID == "" {
		return nil, errors.New(ErrMissingFileID)
	}

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/file/%s/thumbnail", r.ID)
	}

	queryParams := map[string]string{}
	if r.Width != "" {
		queryParams["width"] = r.Width
	}
	if r.Height != "" {
		queryParams["height"] = r.Height
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetQueryParams(queryParams).
		SetHeaders(headers).
		SetErrorResult(&errMsg).
		SetOutputFile(r.PathToSave).
		Get(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	fInfo, err := os.Stat(r.PathToSave)
	if err != nil {
		return nil, err
	}

	rspStruct := &ResponseThumbnail{
		FilePath: r.PathToSave,
		FileName: fInfo.Name(),
		FileSize: fInfo.Size(),
		ResponseDefault: ResponseDefault{
			StatusCode: rsp.GetStatusCode(),
			Success:    true,
		},
	}

	return rspStruct, nil
}

// Delete DELETE /api/file/{id}
func (pd *PixelDrainClient) Delete(r *RequestDelete) (*ResponseDelete, error) {
	if r.ID == "" {
		return nil, errors.New(ErrMissingFileID)
	}

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/file/%s", r.ID)
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var rspStruct ResponseDelete
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetSuccessResult(&rspStruct).
		SetErrorResult(&errMsg).
		Delete(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	rspStruct.StatusCode = rsp.GetStatusCode()

	return &rspStruct, nil
}

// CreateList POST /api/list
func (pd *PixelDrainClient) CreateList(r *RequestCreateList) (*ResponseCreateList, error) {
	if r.URL == "" {
		r.URL = APIURL + "/list"
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	var rspStruct ResponseCreateList
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetBodyJsonBytes(data).
		SetSuccessResult(&rspStruct).
		SetErrorResult(&errMsg).
		Post(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	rspStruct.StatusCode = rsp.GetStatusCode()

	return &rspStruct, nil
}

// GetList GET /api/list/{id}
func (pd *PixelDrainClient) GetList(r *RequestGetList) (*ResponseGetList, error) {
	if r.ID == "" {
		return nil, errors.New(ErrMissingFileID)
	}

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/list/%s", r.ID)
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var rspStruct ResponseGetList
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetSuccessResult(&rspStruct).
		SetErrorResult(&errMsg).
		Get(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	rspStruct.StatusCode = rsp.GetStatusCode()

	return &rspStruct, nil
}

// GetUser GET /api/user
func (pd *PixelDrainClient) GetUser(r *RequestGetUser) (*ResponseGetUser, error) {
	if r.URL == "" {
		r.URL = APIURL + "/user"
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var rspStruct ResponseGetUser
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetSuccessResult(&rspStruct).
		SetErrorResult(&errMsg).
		Get(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	status := false
	if rsp.GetStatusCode() == http.StatusOK {
		status = true
	}

	rspStruct.Success = status
	rspStruct.StatusCode = rsp.GetStatusCode()

	return &rspStruct, nil
}

// GetUserFiles GET /api/user/files
func (pd *PixelDrainClient) GetUserFiles(r *RequestGetUserFiles) (*ResponseGetUserFiles, error) {
	if r.URL == "" {
		r.URL = APIURL + "/user/files"
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var rspStruct ResponseGetUserFiles
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetSuccessResult(&rspStruct).
		SetErrorResult(&errMsg).
		Get(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	status := false
	if rsp.GetStatusCode() == http.StatusOK {
		status = true
	}

	rspStruct.Success = status
	rspStruct.StatusCode = rsp.GetStatusCode()

	return &rspStruct, nil
}

// GetUserLists GET /api/user/lists
func (pd *PixelDrainClient) GetUserLists(r *RequestGetUserLists) (*ResponseGetUserLists, error) {
	if r.URL == "" {
		r.URL = APIURL + "/user/lists"
	}

	headers := map[string]string{}
	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(headers, "", r.Auth.APIKey)
	}

	var rspStruct ResponseGetUserLists
	var errMsg ErrorMessage
	rsp, err := pd.Client.Client.R().
		SetHeaders(headers).
		SetSuccessResult(&rspStruct).
		SetErrorResult(&errMsg).
		Get(r.URL)
	if pd.Debug && rsp != nil {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	if rsp != nil && rsp.IsErrorState() { // Status code >= 400.
		err = errors.New(errMsg.ErrorMessage)
		return nil, err
	}

	status := false
	if rsp.GetStatusCode() == http.StatusOK {
		status = true
	}

	rspStruct.Success = status
	rspStruct.StatusCode = rsp.GetStatusCode()

	return &rspStruct, nil
}

// pixeldrain want an empty username and the APIKey as password
// addBasicAuthHeader create a http basic auth header from username and password
func addBasicAuthHeader(h map[string]string, u string, p string) map[string]string {
	h["Authorization"] = "Basic " + generateBasicAuthToken(u, p)

	return h
}

// generateBasicAuthToken generate string for basic auth header
func generateBasicAuthToken(u string, p string) string {
	auth := u + ":" + p

	return base64.StdEncoding.EncodeToString([]byte(auth))
}
