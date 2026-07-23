package export

import "flashpnl/internal/sheetsdata"

// PLData mirrors the object returned by calcPL() in index.html, sent
// verbatim as the request body's plData field for /api/export-excel.
type PLData struct {
	Rev          sheetsdata.Revenue `json:"rev"`
	TotalRevenue float64            `json:"totalRevenue"`
	FbRevenue    float64            `json:"fbRevenue"`

	Pteb      sheetsdata.PTEB `json:"pteb"`
	TotalPTEB float64         `json:"totalPTEB"`

	Cos      sheetsdata.COS `json:"cos"`
	TotalCOS float64        `json:"totalCOS"`

	FoFixed      sheetsdata.FOFixed `json:"foFixed"`
	FoFixedTotal float64            `json:"foFixedTotal"`
	FoVar        sheetsdata.FOVar   `json:"foVar"`
	FoVarTotal   float64            `json:"foVarTotal"`
	FoOtherExp   float64            `json:"foOtherExp"`

	HkFixed      sheetsdata.HKFixed `json:"hkFixed"`
	HkFixedTotal float64            `json:"hkFixedTotal"`
	HkVar        sheetsdata.HKVar   `json:"hkVar"`
	HkVarTotal   float64            `json:"hkVarTotal"`
	HkOtherExp   float64            `json:"hkOtherExp"`

	FbFixed      sheetsdata.FBFixed `json:"fbFixed"`
	FbFixedTotal float64            `json:"fbFixedTotal"`
	FbVar        sheetsdata.FBVar   `json:"fbVar"`
	FbVarTotal   float64            `json:"fbVarTotal"`
	FbOtherExp   float64            `json:"fbOtherExp"`

	AgFixed      sheetsdata.AGFixed `json:"agFixed"`
	AgFixedTotal float64            `json:"agFixedTotal"`
	AgVar        sheetsdata.AGVar   `json:"agVar"`
	AgVarTotal   float64            `json:"agVarTotal"`
	AgOtherExp   float64            `json:"agOtherExp"`

	HrdFixed      sheetsdata.HRDFixed `json:"hrdFixed"`
	HrdFixedTotal float64             `json:"hrdFixedTotal"`
	HrdVar        sheetsdata.HRDVar   `json:"hrdVar"`
	HrdVarTotal   float64             `json:"hrdVarTotal"`
	HrdOtherExp   float64             `json:"hrdOtherExp"`

	SalesFixed      sheetsdata.SalesFixed `json:"salesFixed"`
	SalesFixedTotal float64               `json:"salesFixedTotal"`
	SalesVar        sheetsdata.SalesVar   `json:"salesVar"`
	SalesVarTotal   float64               `json:"salesVarTotal"`
	SalesOtherExp   float64               `json:"salesOtherExp"`

	PomecFixed      sheetsdata.PomecFixed `json:"pomecFixed"`
	PomecFixedTotal float64               `json:"pomecFixedTotal"`
	PomecVar        sheetsdata.PomecVar   `json:"pomecVar"`
	PomecVarTotal   float64               `json:"pomecVarTotal"`
	PomecOtherExp   float64               `json:"pomecOtherExp"`

	TotalOtherExp float64 `json:"totalOtherExp"`
	ManagementFee float64 `json:"managementFee"`
	Gop           float64 `json:"gop"`
	GopMargin     float64 `json:"gopMargin"`
}
