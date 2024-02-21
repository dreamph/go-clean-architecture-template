package utils

import (
	coremodels "backend/internal/core/models"
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	gomime "github.com/cubewise-code/go-mime"
	"github.com/pkg/errors"
	"github.com/valyala/bytebufferpool"
)

func OpenFile(filePath string) (*os.File, error) {
	return os.Open(filepath.Clean(filePath))
}

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func FileToBase64(file *os.File) (string, error) {
	reader := bufio.NewReader(file)
	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded, nil
}

func GenerateFileName(fileName string) string {
	fileExtension := filepath.Ext(fileName)
	return NewID() + fileExtension
}

func GenerateFileNameFromRequest(intendedFileName string, fileName string) string {
	fileExtension := filepath.Ext(fileName)
	return intendedFileName + fileExtension
}

func FileToBytes(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func FileHeaderToBytes(h *multipart.FileHeader) []byte {
	file, err := h.Open()
	if err != nil {
		return nil
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil
	}
	return data
}

func WriteFile(filePath string, bytes []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func CloseFile(file *os.File) {
	if file != nil {
		_ = file.Close()
	}
}

func RemoveFile(filePath string) {
	_ = os.Remove(filePath)
}

func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}

func GetBaseName(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func GetFileName(fileName string) string {
	return filepath.Base(fileName)
}

func HashFileByFile(reader io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, reader); err != nil {
		return "", err
	}
	hash := h.Sum(nil)
	return hex.EncodeToString(hash[:]), nil
}

func HashDataByBytes(data []byte) (string, error) {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

func HashFileByBytes(data []byte) (string, error) {
	return HashDataByBytes(data)
}

func ResetFile(file *os.File) (ret int64, err error) {
	return file.Seek(0, 0)
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func ReadFileUseByteBufferPool(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64ToBytes(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func FileSize(file *os.File) (int64, error) {
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func FileSizeByBytes(data []byte) int64 {
	return int64(len(data))
}

func FileSizeByUrl(requestUrl string) (int64, error) {
	resp, err := client.Head(requestUrl)
	if err != nil {
		return 0, err
	}
	return resp.ContentLength, nil
}

func FileSizeDesc(file *os.File) (string, error) {
	size, err := FileSize(file)
	if err != nil {
		return "", err
	}
	return ByteCountSI(size), nil
}

func FileSizeDescByBytes(data []byte) (string, error) {
	size := FileSizeByBytes(data)
	return ByteCountSI(size), nil
}

func GetContentTypeByFileName(fileName string) string {
	extension := filepath.Ext(fileName)
	return gomime.TypeByExtension(extension)
}

func GetContentTypeByBytes(data []byte) string {
	return http.DetectContentType(data)
}

func FilenameFromUrl(urlValue string) string {
	u, err := url.Parse(urlValue)
	if err != nil {
		return ""
	}
	path, err := url.QueryUnescape(u.EscapedPath())
	if err != nil {
		return ""
	}
	return filepath.Base(path)
}

func CopyFile(sourceFile string, destinationFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}
	return nil
}

func ByteToReader(data []byte) (io.Reader, error) {
	if len(data) == 0 {
		return nil, errors.Errorf("invalid bytes")
	}
	return bytes.NewReader(data), nil
}

func ToFileData(fileBytes []byte, fileName string) *coremodels.FileData {
	fileData := &coremodels.FileData{
		FileData:        fileBytes,
		FileName:        fileName,
		FileContentType: GetContentTypeByBytes(fileBytes),
		FileSize:        FileSizeByBytes(fileBytes),
	}
	return fileData
}
