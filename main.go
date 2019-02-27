package main

import (
	"net/http"
	"time"

	"github.com/fwiedmann/airgab/pkg/metrics"

	"github.com/fwiedmann/airgab/pkg/opts"
	"github.com/fwiedmann/airgab/pkg/rsync"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	opts := opts.New()

	r := rsync.InitRsync(*opts)
	r.CheckKey()
	r.CheckConnection()

	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics.AirgabCounter)
	registry.MustRegister(metrics.AirgabTimeStamp)
	registry.MustRegister(metrics.AirgabBackupSize)
	registry.MustRegister(metrics.AirgabBackupDuration)

	go func() {
		for {
			metrics.AirgabBackupDuration.Set(r.RunSync())
			metrics.AirgabCounter.Inc()
			metrics.AirgabTimeStamp.SetToCurrentTime()
			metrics.AirgabBackupSize.Set(r.GetBackupSize())

			time.Sleep(opts.Interval)
		}
	}()
	log.Info("Starting to handle connections under /metrics")
	log.Infof("Start to back up %q", opts.Source)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.ListenAndServe(":9100", nil)
}
