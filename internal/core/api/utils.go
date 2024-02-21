package api

import (
	apicommons "backend/internal/core/api/commons"
	"backend/internal/core/auth/jwt"
	coreconstants "backend/internal/core/constants"
	coremodels "backend/internal/core/models"
	"backend/internal/core/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func FormFile(c *fiber.Ctx, fileKey string) *coremodels.FileData {
	file, _ := formFile(c, fileKey, false)
	return file
}

func formFile(c *fiber.Ctx, fileKey string, errorIfNotFound bool) (*coremodels.FileData, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files, ok := form.File[fileKey]
	if !ok {
		if errorIfNotFound {
			return nil, errors.New("fileKey not found:" + fileKey)
		}
		return nil, nil
	}
	file := files[0]
	return &coremodels.FileData{
		FileData:        utils.FileHeaderToBytes(file),
		FileSize:        file.Size,
		FileContentType: file.Header["Content-Type"][0],
		FileName:        file.Filename,
	}, nil
}

func FormFiles(c *fiber.Ctx, fileKey string) *[]coremodels.FileData {
	files, _ := formFiles(c, fileKey, false)
	return files
}

func formFiles(c *fiber.Ctx, fileKey string, errorIfNotFound bool) (*[]coremodels.FileData, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files, ok := form.File[fileKey]
	if !ok {
		if errorIfNotFound {
			return nil, errors.New("fileKey not found:" + fileKey)
		}
		return nil, nil
	}

	var uploadFiles []coremodels.FileData
	for _, file := range files {
		uploadFiles = append(uploadFiles, coremodels.FileData{
			FileData:        utils.FileHeaderToBytes(file),
			FileSize:        file.Size,
			FileContentType: file.Header["Content-Type"][0],
			FileName:        file.Filename,
		})
	}
	return &uploadFiles, nil
}

func getUserRequestToken(c *fiber.Ctx) string {
	return jwt.ExtractToken(c.Get(coreconstants.AuthorizationHeaderName))
}

func getUserRequestInfo(c *fiber.Ctx, jwtToken jwt.JwtToken) *coremodels.UserRequestInfo {
	token := getUserRequestToken(c)
	tokenData, _ := jwtToken.GetTokenData(token)
	return initUserRequestInfo(c, tokenData)
}

func initUserRequestInfo(c *fiber.Ctx, tokenData *jwt.TokenData) *coremodels.UserRequestInfo {
	userRequestInfo := &coremodels.UserRequestInfo{}
	userRequestInfo.UserRequestIP = apicommons.GetClientIP(c)
	if tokenData != nil {
		userRequestInfo.HasUserRequest = true
		userRequestInfo.ID = tokenData.ID
		userRequestInfo.Scope = tokenData.Scope
		userRequestInfo.Org = tokenData.Org
		userRequestInfo.OrgIssuer = tokenData.OrgIssuer
		userRequestInfo.Info = tokenData.Info
		userRequestInfo.UserType = tokenData.UserType
	}
	return userRequestInfo
}
