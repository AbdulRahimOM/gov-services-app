package gateway

import (
	logs "github.com/AbdulRahimOM/gov-services-app/internal/logger"
	"github.com/sirupsen/logrus"
)

var gatewayLogger *logrus.Entry = logs.NewLoggerWithServiceName("gateway")
var userLogger *logrus.Entry = logs.NewLoggerWithServiceName("user-api-gateway")
var adminLogger *logrus.Entry = logs.NewLoggerWithServiceName("admin-api-gateway")