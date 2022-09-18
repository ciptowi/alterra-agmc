package restrict

import "github.com/labstack/echo/v4/middleware"

var CostumLogger = middleware.LoggerConfig{
	Format: "time=${time_rfc3339_nano},s method=${method}, uri=${uri}, status=${status}\n",
}
