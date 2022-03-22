package metric

type metricType interface{
	Increment()
	Update(value uint64)
	GetName() string
	GetType() string
	GetFullMetric() string
}