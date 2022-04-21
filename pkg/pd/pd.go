package pd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/imroc/req"
)

const (
	Name             = "PixelDrain.com"
	BaseURL          = "https://pixeldrain.com/"
	APIURL           = BaseURL + "api"
	DefaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36"
	// errors
	ErrMissingPathToFile = "file path is required"
	ErrMissingFileID     = "file id is required"
)

type ClientOptions struct {
	Debug             bool
	ProxyURL          string
	EnableCookies     bool
	EnableInsecureTLS bool
	Timeout           time.Duration
}

type Client struct {
	Header  req.Header
	Request *req.Req
}

type PixelDrainClient struct {
	Client *Client
	Debug  bool
}

// New - create a new PixelDrainClient
func New(opt *ClientOptions, c *Client) *PixelDrainClient {
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
		c = &Client{
			Header: req.Header{
				"User-Agent": DefaultUserAgent,
			},
			Request: req.New(),
		}
	}

	// set the request options
	c.Request.EnableCookie(opt.EnableCookies)
	c.Request.EnableInsecureTLS(opt.EnableInsecureTLS)
	c.Request.SetTimeout(opt.Timeout)
	if opt.ProxyURL != "" {
		_ = c.Request.SetProxyUrl(opt.ProxyURL)
	}

	pdc := &PixelDrainClient{
		Client: c,
		Debug:  opt.Debug,
	}

	return pdc
}

// UploadPOST POST /api/file
// curl -X POST -i -H "Authorization: Basic <TOKEN>" -F "file=@cat.jpg" https://pixeldrain.com/api/file
func (pd *PixelDrainClient) UploadPOST(r *RequestUpload) (*ResponseUpload, error) {
	if r.PathToFile == "" {
		return nil, errors.New(ErrMissingPathToFile)
	}

	if r.URL == "" {
		r.URL = fmt.Sprint(APIURL + "/file")
	}

	file, err := os.Open(r.PathToFile)
	if err != nil {
		return nil, err
	}

	reqFileUpload := req.FileUpload{
		FileName:  r.GetFileName(),
		FieldName: "file",
		File:      file,
	}

	reqParams := req.Param{
		"anonymous": r.Anonymous,
	}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() && !r.Anonymous {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Post(r.URL, pd.Client.Header, reqFileUpload, reqParams)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	uploadRsp := &ResponseUpload{}
	uploadRsp.StatusCode = rsp.Response().StatusCode
	err = rsp.ToJSON(uploadRsp)
	if err != nil {
		return nil, err
	}

	return uploadRsp, nil
}

// UploadPUT PUT /api/file/{name}
// curl -X PUT -i -H "Authorization: Basic <TOKEN>" --upload-file cat.jpg https://pixeldrain.com/api/file/test_cat.jpg
func (pd *PixelDrainClient) UploadPUT(r *RequestUpload) (*ResponseUpload, error) {
	if r.PathToFile == "" {
		return nil, errors.New(ErrMissingPathToFile)
	}

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/file/%s", r.GetFileName())
	}

	file, err := os.Open(r.PathToFile)
	if err != nil {
		return nil, err
	}

	// we dont send this paramter due a bug of pixeldrain side
	//reqParams := req.Param{
	//	"anonymous": r.Anonymous,
	//}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() && !r.Anonymous {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Put(r.URL, pd.Client.Header, file)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	uploadRsp := &ResponseUpload{}
	uploadRsp.StatusCode = rsp.Response().StatusCode
	if uploadRsp.StatusCode == http.StatusCreated {
		uploadRsp.Success = true
	}
	err = rsp.ToJSON(uploadRsp)
	if err != nil {
		return nil, err
	}

	return uploadRsp, nil
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

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Get(r.URL, pd.Client.Header)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	err = rsp.ToFile(r.PathToSave)
	if err != nil {
		return nil, err
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
			StatusCode: rsp.Response().StatusCode,
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

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Get(r.URL, pd.Client.Header)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	fileInfoRsp := &ResponseFileInfo{}
	fileInfoRsp.StatusCode = rsp.Response().StatusCode
	if fileInfoRsp.StatusCode == http.StatusOK {
		fileInfoRsp.Success = true
	}
	err = rsp.ToJSON(fileInfoRsp)
	if err != nil {
		return nil, err
	}

	return fileInfoRsp, nil
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

	queryParams := req.QueryParam{}
	if r.Width != "" {
		queryParams["width"] = r.Width
	}
	if r.Height != "" {
		queryParams["height"] = r.Height
	}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Get(r.URL, pd.Client.Header, queryParams)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	err = rsp.ToFile(r.PathToSave)
	if err != nil {
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
			StatusCode: rsp.Response().StatusCode,
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

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Delete(r.URL, pd.Client.Header)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	rspStruct := &ResponseDelete{}
	err = rsp.ToJSON(rspStruct)
	if err != nil {
		return nil, err
	}

	rspStruct.StatusCode = rsp.Response().StatusCode

	return rspStruct, nil
}

// CreateList POST /api/list
func (pd *PixelDrainClient) CreateList(r *RequestCreateList) (*ResponseCreateList, error) {
	if r.URL == "" {
		r.URL = APIURL + "/list"
	}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() && !r.Anonymous {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	data, err := json.Marshal(r)

	rsp, err := pd.Client.Request.Post(r.URL, pd.Client.Header, data)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	rspStruct := &ResponseCreateList{}
	err = rsp.ToJSON(rspStruct)
	if err != nil {
		return nil, err
	}

	rspStruct.StatusCode = rsp.Response().StatusCode

	return rspStruct, nil
}

// GetList GET /api/list/{id}
func (pd *PixelDrainClient) GetList(r *RequestGetList) (*ResponseGetList, error) {
	if r.ID == "" {
		return nil, errors.New(ErrMissingFileID)
	}

	if r.URL == "" {
		r.URL = fmt.Sprintf(APIURL+"/list/%s", r.ID)
	}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Get(r.URL, pd.Client.Header)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	rspStruct := &ResponseGetList{}
	err = rsp.ToJSON(rspStruct)
	if err != nil {
		return nil, err
	}

	rspStruct.StatusCode = rsp.Response().StatusCode

	return rspStruct, nil
}

// GetUserFiles GET /api/user/files
func (pd *PixelDrainClient) GetUserFiles(r *RequestGetUserFiles) (*ResponseGetUserFiles, error) {
	if r.URL == "" {
		r.URL = APIURL + "/user/files"
	}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Get(r.URL, pd.Client.Header)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	rspStruct := &ResponseGetUserFiles{}
	err = rsp.ToJSON(rspStruct)
	if err != nil {
		return nil, err
	}

	status := false
	if rsp.Response().StatusCode == http.StatusOK {
		status = true
	}

	rspStruct.Success = status
	rspStruct.StatusCode = rsp.Response().StatusCode

	return rspStruct, nil
}

// GetUserLists GET /api/user/lists
func (pd *PixelDrainClient) GetUserLists(r *RequestGetUserLists) (*ResponseGetUserLists, error) {
	if r.URL == "" {
		r.URL = APIURL + "/user/lists"
	}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.IsAuthAvailable() {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Get(r.URL, pd.Client.Header)
	if pd.Debug {
		log.Println(rsp.Dump())
	}
	if err != nil {
		return nil, err
	}

	rspStruct := &ResponseGetUserLists{}
	err = rsp.ToJSON(rspStruct)
	if err != nil {
		return nil, err
	}

	status := false
	if rsp.Response().StatusCode == http.StatusOK {
		status = true
	}

	rspStruct.Success = status
	rspStruct.StatusCode = rsp.Response().StatusCode

	return rspStruct, nil
}

// pixeldrain want an empty username and the APIKey as password
// addBasicAuthHeader create a http basic auth header from username and password
func addBasicAuthHeader(h req.Header, u string, p string) *req.Header {
	h["Authorization"] = "Basic " + generateBasicAuthToken(u, p)
	return &h
}

// generateBasicAuthToken generate string for basic auth header
func generateBasicAuthToken(u string, p string) string {
	auth := u + ":" + p
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
