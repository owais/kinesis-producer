package producer

type Hooks interface {
	OnDrain(size, length int64)
	OnPutRecords(numRecords, numTotalRecords, putLatencyMS int64, reason string)
	OnPutErr(errCode string)
}

type noopHooks struct{}

func (h *noopHooks) OnDrain(size, length int64) {}

func (h *noopHooks) OnPutRecords(batches, records, putLatencyMS int64, reason string) {}

func (h *noopHooks) OnPutErr(errCode string) {}
