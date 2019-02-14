package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	AirgabCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "airgab_counter",
		Help: "Counter of succeeded airgab runs"})
	AirgabTimeStamp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "airgab_last_success",
		Help: "Last succeeded run of airgab"})
	AirgabBackupSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "airgab_current_backup_size_megabyte",
		Help: "Size of last backup in megabyte"})
)
