package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	TransactionAuthorizedMetric    = prometheus.NewCounter(prometheus.CounterOpts{Name: "transaction_authorized"})
	TransactionNotAuthorizedMetric = prometheus.NewCounter(prometheus.CounterOpts{Name: "transaction_not_authorized"})
)

func InitTransactionsMetrics() {
	prometheus.MustRegister(TransactionAuthorizedMetric)
	prometheus.MustRegister(TransactionNotAuthorizedMetric)
}
