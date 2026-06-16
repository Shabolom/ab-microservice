package di

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (d *DI) Healthcheck() {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		if err := d.pgConn.Ping(ctx); err != nil {
			http.Error(w, "postgres unavailable", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	mux.HandleFunc("/healthz/kafka", func(w http.ResponseWriter, r *http.Request) {
		if err := d.kafka.IsInitialized(); err != nil {
			http.Error(w, "kafka is not initialized", http.StatusServiceUnavailable)
			return
		}

		if err := d.kafka.HealthCheck(); err != nil {
			http.Error(w, "kafka is not healthy", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("kafka ok"))
	})

	d.Logger().Info(
		"healthcheck server started",
		zap.String("addr", ":"+d.Config().HealthcheckPort),
	)

	go func() {
		if err := http.ListenAndServe(":"+d.Config().HealthcheckPort, mux); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {

			d.Logger().Fatal(
				"healthcheck server failed",
				zap.String("addr", ":"+d.Config().HealthcheckPort),
				zap.Error(err),
			)
		}
	}()
}
