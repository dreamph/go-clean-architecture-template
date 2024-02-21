package models

import "os"

type FileDownloadByFileResponse struct {
	FilePath         string `json:"filePath"`
	DownloadFileName string `json:"downloadFileName"`
}

type FileDownloadByBytesResponse struct {
	FileData         []byte `json:"fileData"`
	DownloadFileName string `json:"downloadFileName"`
}

type DownloadFile struct {
	URL      string `json:"url"`
	FileName string `json:"fileName"`
}

type FileData struct {
	FileData        []byte `json:"fileData"`
	FileName        string `json:"fileName"`
	FileSize        int64  `json:"fileSize"`
	FileContentType string `json:"fileContentType"`
}

type File struct {
	File            *os.File `json:"file"`
	FileName        string   `json:"fileName"`
	FileSize        int64    `json:"fileSize"`
	FileContentType string   `json:"fileContentType"`
}

type RawResponse struct {
	Data       []byte `json:"data"`
	HttpStatus int    `json:"200"`
}
