package permissions

import (
	"backend/internal/core/api/middleware"
	"strings"

	"github.com/iancoleman/strcase"

	"fmt"
)

func Gen(authPermissions middleware.AuthPermissions) {
	permissionsList := authPermissions.PermissionsList()
	GenCsv(permissionsList)
}

func GenCsv(permissionsList []string) {
	//p,USR,activity,find-by-id
	for _, data := range permissionsList {
		name := data
		sep := strings.Split(name, ":")

		fmt.Println(`p,ADM,` + sep[0] + `,` + sep[1])
	}
}

func GenConstants(permissionsList []string) {
	for _, data := range permissionsList {
		name := data
		name = strings.Replace(name, ":", "_", -1)
		name = strings.Replace(name, "-", "_", -1)
		fmt.Println(strcase.ToCamel(name) + " = " + `"` + data + `"`)
	}
}
