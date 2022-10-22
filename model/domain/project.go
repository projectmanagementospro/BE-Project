package domain

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	KodeSortname          string  `json:"kode_sortname" gorm:"type:TEXT;not null"`
	JumlahStakeholder     uint64  `json:"jumlah_stakeholder" gorm:"type:uint;not null"`
	Nama                  string  `json:"nama" gorm:"type:varchar(255);not null, unique"`
	MulaiProyek           string  `json:"mulai_proyek" gorm:"type:timestamp;not null"`
	AkhirProyek           string  `json:"akhir_proyek" gorm:"type:timestamp;not null"`
	AreaKerja             string  `json:"area_kerja" gorm:"type:varchar(200);not null"`
	LokasiKantor          string  `json:"lokasi_kantor" gorm:"type:TEXT;not null"`
	RencanaBiaya          string  `json:"rencana_biaya" gorm:"type:varchar(255);not null"`
	BiayaActual           string  `json:"biaya_actual" gorm:"type:varchar(255);not null"`
	Company               string  `json:"company" gorm:"type:varchar(255);not null"`
	PmId                  uint64  `json:"pm_id" gorm:"type:uint;not null"`
	TypeProyekId          uint64  `json:"type_proyek_id" gorm:"type:uint;not null"`
	PersentaseProgress    float32 `json:"persentase_progress" gorm:"type:double;not null"`
	Keterangan            string  `json:"keterangan" gorm:"type:TEXT;not null"`
	DurasiProyek          uint    `json:"durasi_proyek" gorm:"type:uint;not null"`
	CreatedBy             string  `json:"created_by" gorm:"type:varchar(80);not null"`
	UpdatedBy             string  `json:"updated_by" gorm:"type:varchar(80);not null"`
	DeletedBy             string  `json:"deleted_by" gorm:"type:varchar(80)"`
	ProgressByWorklog     float32 `json:"progress_by_worklog" gorm:"type:boolean;not null"`
	Status                string  `json:"status" gorm:"type:varchar(50);not null"`
	ProjectObjectives     string  `json:"project_objectives" gorm:"type:TEXT;not null"`
	ConsideredSuccessWhen string  `json:"considered_success_when" gorm:"type:TEXT;not null"`
	PotentialRisk         string  `json:"potential_risk" gorm:"type:TEXT;not null"`
	CestingEnvironment    string  `json:"testing_environment" gorm:"type:TEXT;not null"`
	CurrencySymbol        string  `json:"currency_symbol" gorm:"type:varchar(100);not null"`
	CurrencyCode          string  `json:"currency_code" gorm:"type:varchar(100);not null"`
	CurrencyName          string  `json:"currency_name" gorm:"type:varchar(100);not null"`
	PhaseId               uint64  `json:"phase_id" gorm:"type:uint;not null"`
	BudgetHealth          uint64  `json:"budget_health" gorm:"type:varchar(10);not null"`
}
