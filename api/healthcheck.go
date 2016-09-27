// podium
// https://github.com/topfreegames/podium
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2016 Top Free Games <backend@tfgco.com>
// Forked from
// https://github.com/dayvson/go-leaderboard
// Copyright © 2013 Maxwell Dayvson da Silva

package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// HealthCheckHandler is the handler responsible for validating that the app is still up
func HealthCheckHandler(app *App) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Set("route", "Healthcheck")
		workingString := app.Config.GetString("healthcheck.workingText")
		res, err := app.RedisClient.Client.Ping().Result()
		if err != nil || res != "PONG" {
			return FailWith(http.StatusInternalServerError, fmt.Sprintf("Error connecting to redis: %s", err), c)
		}

		workingString = strings.TrimSpace(workingString)
		return c.String(http.StatusOK, workingString)
	}
}
