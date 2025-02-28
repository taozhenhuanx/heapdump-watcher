package collectors

import (
	"heapdump_watcher/controller/watchFile"

	"github.com/prometheus/client_golang/prometheus"
)

/*
   统计告警次数metrics接口
*/

// AlertCountCollector 定义了一个用于收集内存溢出 (OOM) 警报计数的 Prometheus 收集器。
type AlertCountCollector struct {
	// alertCount 是一个 Prometheus 指标描述符，用于存储 OOM Alert 的计数。
	alertCount *prometheus.Desc
}

// NewAlertCountCollector 创建一个新的 AlertCountCollector 实例，并初始化其指标描述符。
// 返回值是一个指向 AlertCountCollector 的指针，已初始化并准备好使用。
func NewAlertCountCollector() *AlertCountCollector {
	return &AlertCountCollector{
		alertCount: prometheus.NewDesc(
			"oom_alert_count",            // 指标名称
			"Total number of OOM alerts", // 指标的帮助信息
			nil,                          // 标签列表（nil 表示不使用标签）
			nil,                          // const 标签（nil 表示不使用 const 标签）
		),
	}
}

// Describe 实现了 prometheus.Collector 接口的 Describe 方法，用于将指标的描述符发送到指定的通道。
// 这是 Prometheus 收集器的标准接口方法，用于告诉 Prometheus 可以收集哪些指标。
func (c *AlertCountCollector) Describe(descs chan<- *prometheus.Desc) {
	// 将 alertCount 指标的描述符发送到通道中。
	descs <- c.alertCount
}

// Collect 实现了 prometheus.Collector 接口的 Collect 方法，用于收集当前的 OOM Alert 计数值。
// 收集的值会被写入到指定的 prometheus.Metric 通道中。
func (c *AlertCountCollector) Collect(metrics chan<- prometheus.Metric) {
	// 获取真实的 OOM Alert 的次数
	alertCount := watchFile.AlertCount
	metrics <- prometheus.MustNewConstMetric(c.alertCount, prometheus.CounterValue, alertCount)
}

/*
	暂时不需要
	WriteCurrentCount 是一个帮助函数，用于将当前的 OOM Alert 计数值写入到指定的通道中。
	它将创建一个新的 ConstMetric 并将其发送到通道。
	这个函数可以被 Collect 方法调用，以避免代码重复。
	func WriteCurrentCount(desc *prometheus.Desc, count float64, metrics chan<- prometheus.Metric) {
		// 使用 prometheus.MustNewConstMetric 创建一个新的 ConstMetric。
		// 这里使用 CounterValue 类型，因为 OOM Alert 的次数是一个累计计数器。
		metric := prometheus.MustNewConstMetric(
			desc,                    // 指标描述符
			prometheus.CounterValue, // 指标类型为 Counter
			count,                   // 当前的 OOM Alert 计数值
		)

		// 将创建的 metric 发送到指定的通道。
		metrics <- metric
	}
*/
