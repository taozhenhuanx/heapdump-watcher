package collectors

import "github.com/prometheus/client_golang/prometheus"

/*
   统计告警次数metrics接口
*/

type AlertCountCollector struct {
	// 定义指标, 自己实现Prometheus的2个接口, 指标在这里定义
	alertCount *prometheus.Desc
}

// 构造函数
func NewAlertCountCollector() *AlertCountCollector {
	return &AlertCountCollector{
		alertCount: prometheus.NewDesc(
			"OOM Alert Count",
			"OOM Alert Count",
			nil,
			nil,
		),
	}
}

func (c *AlertCountCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.alertCount
}

func (c *AlertCountCollector) Collect(metrics chan<- prometheus.Metric) {
	// 获取OOM的次数
	alertCount := 1.0

	// 写入指标的值
	metrics <- prometheus.MustNewConstMetric(c.alertCount, prometheus.GaugeValue, alertCount)
}
