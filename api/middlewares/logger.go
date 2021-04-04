package middlewares

import (
	"github.com/abhay676/auth/api/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type LogMiddleware struct {
	logger *log.Logger
}

func NewLogMiddleware(logger *log.Logger) *LogMiddleware {
	return &LogMiddleware{logger: logger}
}

func (m *LogMiddleware) Func() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			logRespWriter := utils.NewLogResponseWriter(r, w)
			next.ServeHTTP(logRespWriter, r)

			m.logger.Printf(
				"duration=%s status=%d params = %s body=%s \n",
				time.Since(startTime).String(),
				logRespWriter.StatusCode,
				logRespWriter.Request,
				logRespWriter.Buf.String())

		})
	}
}
