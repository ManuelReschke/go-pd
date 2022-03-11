package pd

import (
	"errors"
	"fmt"
	"log"
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

type PixelDrainClient struct {
	Client *Client
	Debug  bool
}

type Client struct {
	Header   req.Header
	Request  *req.Req
	ProxyURL string
}

type ClientOptions struct {
	Debug             bool
	ProxyURL          string
	EnableCookies     bool
	EnableInsecureTLS bool
	Timeout           time.Duration
}

func New(opt *ClientOptions, c *Client) *PixelDrainClient {
	if c == nil {
		header := req.Header{
			"User-Agent": DefaultUserAgent,
		}

		request := req.New()
		request.EnableCookie(opt.EnableCookies)
		request.EnableInsecureTLS(opt.EnableInsecureTLS)
		request.SetTimeout(opt.Timeout)
		if opt.ProxyURL != "" {
			_ = request.SetProxyUrl(opt.ProxyURL)
		}

		c = &Client{
			Header:   header,
			Request:  request,
			ProxyURL: opt.ProxyURL,
		}
	}

	pdc := &PixelDrainClient{
		Client: c,
		Debug:  opt.Debug,
	}

	return pdc
}

// POST /api/file
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

	rsp, err := pd.Client.Request.Post(r.URL, pd.Client.Header, reqFileUpload)
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

// PUT /api/file/{name}
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
		"name":      r.GetFileName(),
		"anonymous": r.Anonymous,
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
	err = rsp.ToJSON(uploadRsp)
	if err != nil {
		return nil, err
	}

	return uploadRsp, nil
}

// GET /api/file/{id}
func (pd *PixelDrainClient) GetFile() (*ResponseUpload, error) {
	// todo
	return nil, nil
}
