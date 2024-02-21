package utils

import (
	"io"
	"net/url"
	"os"
	"strings"

	coremodels "backend/internal/core/models"
)

func GetDownloadFileNane(fileName string) string {
	return url.PathEscape(fileName)
}

func GetDownloadFileHeaderValue(fileName string) string {
	var builder strings.Builder
	builder.WriteString("attachment;")
	builder.WriteString("filename=")
	builder.WriteString(GetDownloadFileNane(fileName))
	builder.WriteString(";")
	builder.WriteString("filename*=")
	builder.WriteString("utf-8''")
	builder.WriteString(GetDownloadFileNane(fileName))
	return builder.String()
}

func GetDownloadFileHeader(fileName string, safariFileName string) string {
	var builder strings.Builder
	builder.WriteString("attachment;")
	builder.WriteString("filename=")
	builder.WriteString(GetDownloadFileNane(fileName))
	builder.WriteString(";")
	builder.WriteString("filename*=")
	builder.WriteString("utf-8''")
	builder.WriteString(GetDownloadFileNane(safariFileName))
	return builder.String()
}

func GetDownloadFileHeaderValueByMaxLength(fileName string, maxLength int, safariMaxLength int) string {
	extension := GetFileExtension(fileName)
	baseName := GetBaseName(fileName)
	newFileName := GetValueByMaxLength(baseName, maxLength)
	newSafariFileName := GetValueByMaxLength(baseName, safariMaxLength)
	return GetDownloadFileHeader(newFileName+extension, newSafariFileName+extension)
}

func DownloadFileAsBytes(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Download(url string) (*coremodels.FileData, error) {
	data, err := DownloadFileAsBytes(url)
	if err != nil {
		return nil, err
	}
	return &coremodels.FileData{
		FileData:        data,
		FileName:        FilenameFromUrl(url),
		FileContentType: GetContentTypeByBytes(data),
		FileSize:        FileSizeByBytes(data),
	}, nil
}

func DownloadFileToPath(filepath string, url string) error {
	// Get the data
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
