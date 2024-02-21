package swagger

import (
	"backend/internal/core/api/middleware"
	"backend/internal/core/utils"
	"os"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/swag"
)

const (
	HeaderSwaggerApiKey = "x-swagger-api-key"
)

func RegisterSwaggerRouter(router fiber.Router, contextPath string, prodMode bool, swaggerInfo *swag.Spec) {
	conf := swagger.ConfigDefault
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if utils.IsNotEmpty(swaggerHost) {
		swaggerInfo.Host = swaggerHost
	}
	if utils.IsNotEmpty(contextPath) {
		swaggerInfo.BasePath = contextPath
	}
	swaggerBasePath := os.Getenv("BASE_PATH")
	if utils.IsNotEmpty(swaggerBasePath) {
		swaggerInfo.BasePath = swaggerBasePath
		conf.URL = swaggerInfo.BasePath + "/swagger/doc.json"
	}

	conf.PersistAuthorization = !prodMode
	conf.DisplayRequestDuration = true
	conf.Filter = swagger.FilterConfig{
		Enabled: true,
		//Expression: "app",
	}
	handler := swagger.New(conf)

	headerKey := os.Getenv("SWAGGER_HEADER_NAME_API_KEY")
	if utils.IsEmpty(headerKey) {
		headerKey = HeaderSwaggerApiKey
	}
	apiKey := os.Getenv("SWAGGER_API_KEY")
	apiKeyAuthEnable := utils.IsNotEmpty(apiKey)
	apiKeyAuth := middleware.NewApiKeyAuth(&middleware.ApiKeyOption{
		Key:       apiKey,
		HeaderKey: headerKey,
		Enable:    apiKeyAuthEnable,
	})

	router.Get("/swagger", apiKeyAuth.Auth, handler)
	router.Get("/swagger/*", apiKeyAuth.Auth, handler)
}
