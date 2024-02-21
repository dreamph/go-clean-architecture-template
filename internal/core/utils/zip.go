package utils

import (
	"archive/zip"
	"bytes"
	"io"
	"os"

	"backend/internal/core/models"
)

func ZipFilesByURL(filename string, downloads []models.DownloadFile) error {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	for _, download := range downloads {
		body, err := DownloadFileAsBytes(download.URL)
		if err != nil {
			return err
		}
		zipFile, err := zipWriter.Create(download.FileName)
		if err != nil {
			return err
		}
		_, err = zipFile.Write(body)
		if err != nil {
			return err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, buf.Bytes(), 0777)
	if err != nil {
		return err
	}
	return nil
}

func ZipFilesByBytes(filename string, fileDataList []models.FileData) error {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	for _, dataFile := range fileDataList {
		zipFile, err := zipWriter.Create(dataFile.FileName)
		if err != nil {
			return err
		}
		_, err = zipFile.Write(dataFile.FileData)
		if err != nil {
			return err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return err
	}
	//err = os.WriteFile(filename, buf.Bytes(), 0777)
	err = WriteFile(filename, buf.Bytes())
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/internal/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return err
	}
	return nil
}

func ZipBytesByBytes(fileDataList []models.FileData) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)
	for _, dataFile := range fileDataList {
		zipFile, err := zipWriter.Create(dataFile.FileName)
		if err != nil {
			return nil, err
		}
		_, err = zipFile.Write(dataFile.FileData)
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
