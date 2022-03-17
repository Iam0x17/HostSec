package log

type RecordData struct {
	AttackType   string `json:"attack_type"`
	VectorCnName string `json:"vector_cn_name"`
	RawData      string `json:"raw_data"`
	BackupData   string `json:"backup_data"`
	RecoveryType string `json:"recovery_type"`
}

type RecordSrv struct {
	Record map[string]*RecordData
}

type Record interface {
	Set(string, *RecordData)
	Get(string) (*RecordData, bool)
	Remove(string)
}

func NewRecord() Record {
	return &RecordSrv{
		Record: make(map[string]*RecordData, 0),
	}
}

func (r RecordSrv) Set(k string, record *RecordData) {
	r.Record[k] = record
}

func (r RecordSrv) Get(k string) (*RecordData, bool) {
	val, found := r.Record[k]
	return val, found
}

func (r RecordSrv) Remove(k string) {
	delete(r.Record, k)
}
