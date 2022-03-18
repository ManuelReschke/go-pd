package pd

import (
	"encoding/base64"
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
	if r.Auth.APIKey != "" && !r.Anonymous {
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

	reqParams := req.Param{
		"anonymous": r.Anonymous,
	}

	// pixeldrain want an empty username and the APIKey as password
	if r.Auth.APIKey != "" && !r.Anonymous {
		addBasicAuthHeader(pd.Client.Header, "", r.Auth.APIKey)
	}

	rsp, err := pd.Client.Request.Put(r.URL, pd.Client.Header, file, reqParams)
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

// GetFile GET /api/file/{id}
func (pd *PixelDrainClient) GetFile() (*ResponseUpload, error) {
	// todo
	return nil, nil
}

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
