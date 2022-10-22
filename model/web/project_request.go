package web

type ProjectCreateRequest struct {
	KodeSortname          string  `json:"kode_sortname" binding:"required"`
	JumlahStakeholder     uint64  `json:"jumlah_stakeholder" binding:"required"`
	Nama                  string  `json:"nama" binding:"required"`
	MulaiProyek           string  `json:"mulai_proyek" binding:"required"`
	AkhirProyek           string  `json:"akhir_proyek" binding:"required"`
	AreaKerja             string  `json:"area_kerja" binding:"required"`
	LokasiKantor          string  `json:"lokasi_kantor" binding:"required"`
	RencanaBiaya          string  `json:"rencana_biaya" binding:"required"`
	BiayaActual           string  `json:"biaya_actual" binding:"required"`
	Company               string  `json:"company" binding:"required"`
	PmId                  uint64  `json:"pm_id" binding:"required"`
	TypeProyekId          uint64  `json:"type_proyek_id" binding:"required"`
	PersentaseProgress    float32 `json:"persentase_progress" binding:"required"`
	Keterangan            string  `json:"keterangan" binding:"required"`
	DurasiProyek          uint    `json:"durasi_proyek" binding:"required"`
	CreatedBy             string  `json:"created_by" binding:"required"`
	ProgressByWorklog     float32 `json:"progress_by_worklog" binding:"required"`
	Status                string  `json:"status" binding:"required"`
	ProjectObjectives     string  `json:"project_objectives" binding:"required"`
	ConsideredSuccessWhen string  `json:"considered_success_when" binding:"required"`
	PotentialRisk         string  `json:"potential_risk" binding:"required"`
	CestingEnvironment    string  `json:"testing_environment" binding:"required"`
	CurrencySymbol        string  `json:"currency_symbol" binding:"required"`
	CurrencyCode          string  `json:"currency_code" binding:"required"`
	CurrencyName          string  `json:"currency_name" binding:"required"`
	PhaseId               uint64  `json:"phase_id" binding:"required"`
	BudgetHealth          uint64  `json:"budget_health" binding:"required"`
}

type ProjectUpdateRequest struct {
	ID                    uint64
	KodeSortname          string  `json:"kode_sortname" binding:"required"`
	JumlahStakeholder     uint64  `json:"jumlah_stakeholder" binding:"required"`
	Nama                  string  `json:"nama" binding:"required"`
	MulaiProyek           string  `json:"mulai_proyek" binding:"required"`
	AkhirProyek           string  `json:"akhir_proyek" binding:"required"`
	AreaKerja             string  `json:"area_kerja" binding:"required"`
	LokasiKantor          string  `json:"lokasi_kantor" binding:"required"`
	RencanaBiaya          string  `json:"rencana_biaya" binding:"required"`
	BiayaActual           string  `json:"biaya_actual" binding:"required"`
	Company               string  `json:"company" binding:"required"`
	PmId                  uint64  `json:"pm_id" binding:"required"`
	TypeProyekId          uint64  `json:"type_proyek_id" binding:"required"`
	PersentaseProgress    float32 `json:"persentase_progress" binding:"required"`
	Keterangan            string  `json:"keterangan" binding:"required"`
	DurasiProyek          uint    `json:"durasi_proyek" binding:"required"`
	UpdatedBy             string  `json:"updated_by" binding:"required"`
	ProgressByWorklog     float32 `json:"progress_by_worklog" binding:"required"`
	Status                string  `json:"status" binding:"required"`
	ProjectObjectives     string  `json:"project_objectives" binding:"required"`
	ConsideredSuccessWhen string  `json:"considered_success_when" binding:"required"`
	PotentialRisk         string  `json:"potential_risk" binding:"required"`
	CestingEnvironment    string  `json:"testing_environment" binding:"required"`
	CurrencySymbol        string  `json:"currency_symbol" binding:"required"`
	CurrencyCode          string  `json:"currency_code" binding:"required"`
	CurrencyName          string  `json:"currency_name" binding:"required"`
	PhaseId               uint64  `json:"phase_id" binding:"required"`
	BudgetHealth          uint64  `json:"budget_health" binding:"required"`
}
