package sheetsdata

type UserInfo struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	Name  string `json:"name"`
}

type PTEB struct {
	FO    float64 `json:"fo"`
	HK    float64 `json:"hk"`
	FB    float64 `json:"fb"`
	AG    float64 `json:"ag"`
	HRD   float64 `json:"hrd"`
	Sales float64 `json:"sales"`
	Pomec float64 `json:"pomec"`
}

type COS struct {
	Breakfast float64 `json:"breakfast"`
	RestoF    float64 `json:"restoF"`
	RestoB    float64 `json:"restoB"`
	TapasF    float64 `json:"tapasF"`
	TapasB    float64 `json:"tapasB"`
	RsF       float64 `json:"rsF"`
	RsB       float64 `json:"rsB"`
	Bqt       float64 `json:"bqt"`
}

type FOFixed struct {
	Uniforms         float64 `json:"uniforms"`
	Telephone        float64 `json:"telephone"`
	OtherExpense     float64 `json:"otherExpense"`
	TravelAgencyComm float64 `json:"travelAgencyComm"`
	CableSatellite   float64 `json:"cableSatellite"`
	SystemsInternet  float64 `json:"systemsInternet"`
	OfficerCheck     float64 `json:"officerCheck"`
	Entertainment    float64 `json:"entertainment"`
}

type FOVar struct {
	CleaningSupplies    float64 `json:"cleaningSupplies"`
	GuestSupplies       float64 `json:"guestSupplies"`
	OtherSupplies       float64 `json:"otherSupplies"`
	Transportation      float64 `json:"transportation"`
	PrintingStationary  float64 `json:"printingStationary"`
	GuestTransportation float64 `json:"guestTransportation"`
	CompWelcomeDrink    float64 `json:"compWelcomeDrink"`
	CompFruitBasket     float64 `json:"compFruitBasket"`
}

type HKFixed struct {
	PestControl     float64 `json:"pestControl"`
	Telephone       float64 `json:"telephone"`
	OtherExpense    float64 `json:"otherExpense"`
	ContractHygiene float64 `json:"contractHygiene"`
	Entertainment   float64 `json:"entertainment"`
	OfficerCheck    float64 `json:"officerCheck"`
	Uniforms        float64 `json:"uniforms"`
	LicensesAndFees float64 `json:"licensesAndFees"`
	EquipmentRental float64 `json:"equipmentRental"`
}

type HKVar struct {
	Linen              float64 `json:"linen"`
	CleaningSupplies   float64 `json:"cleaningSupplies"`
	GuestSupplies      float64 `json:"guestSupplies"`
	OtherSupplies      float64 `json:"otherSupplies"`
	OtherExpense       float64 `json:"otherExpense"`
	LaundryDryCleaning float64 `json:"laundryDryCleaning"`
}

type FBFixed struct {
	Menus              float64 `json:"menus"`
	Licenses           float64 `json:"licenses"`
	OtherExpense       float64 `json:"otherExpense"`
	BanquetExpense     float64 `json:"banquetExpense"`
	ContractService    float64 `json:"contractService"`
	Entertainment      float64 `json:"entertainment"`
	OfficerCheck       float64 `json:"officerCheck"`
	Uniforms           float64 `json:"uniforms"`
	Training           float64 `json:"training"`
	EquipmentRental    float64 `json:"equipmentRental"`
	MusicEntertainment float64 `json:"musicEntertainment"`
}

type FBVar struct {
	CleaningSupplies   float64 `json:"cleaningSupplies"`
	GuestSupplies      float64 `json:"guestSupplies"`
	KitchenFuel        float64 `json:"kitchenFuel"`
	KitchenSupplies    float64 `json:"kitchenSupplies"`
	PrintingStationary float64 `json:"printingStationary"`
	MusicEntertainment float64 `json:"musicEntertainment"`
	Utensil            float64 `json:"utensil"`
}

type AGFixed struct {
	Uniforms         float64 `json:"uniforms"`
	Telephone        float64 `json:"telephone"`
	LicensesAndFees  float64 `json:"licensesAndFees"`
	ItSoftware       float64 `json:"itSoftware"`
	ItHardware       float64 `json:"itHardware"`
	CashierOverShort float64 `json:"cashierOverShort"`
	BankCharges      float64 `json:"bankCharges"`
	PostageExpress   float64 `json:"postageExpress"`
}

type AGVar struct {
	Transportation        float64 `json:"transportation"`
	PrintingStationary    float64 `json:"printingStationary"`
	CreditCardCommissions float64 `json:"creditCardCommissions"`
}

type HRDFixed struct {
	Telephone        float64 `json:"telephone"`
	Transportation   float64 `json:"transportation"`
	Training         float64 `json:"training"`
	OtherExpense     float64 `json:"otherExpense"`
	SportSocial      float64 `json:"sportSocial"`
	EmployeeRelation float64 `json:"employeeRelation"`
	Outsource        float64 `json:"outsource"`
	License          float64 `json:"license"`
	Entertainment    float64 `json:"entertainment"`
	OfficerCheck     float64 `json:"officerCheck"`
}

type HRDVar struct {
	PrintingStationary float64 `json:"printingStationary"`
}

type SalesFixed struct {
	PartnershipSponsors float64 `json:"partnershipSponsors"`
	Photography         float64 `json:"photography"`
	Merchandise         float64 `json:"merchandise"`
	Collaterals         float64 `json:"collaterals"`
	Entertainment       float64 `json:"entertainment"`
	Telephone           float64 `json:"telephone"`
	Miscellaneous       float64 `json:"miscellaneous"`
	PostageExpress      float64 `json:"postageExpress"`
}

type SalesVar struct {
	LocalSalesCall     float64 `json:"localSalesCall"`
	PrintingStationary float64 `json:"printingStationary"`
}

type PomecFixed struct {
	PestControl     float64 `json:"pestControl"`
	Telephone       float64 `json:"telephone"`
	OtherExpense    float64 `json:"otherExpense"`
	OfficerCheck    float64 `json:"officerCheck"`
	Uniforms        float64 `json:"uniforms"`
	LicensesAndFees float64 `json:"licensesAndFees"`
}

type PomecVar struct {
	Uniforms           float64 `json:"uniforms"`
	OtherSupplies      float64 `json:"otherSupplies"`
	PrintingStationary float64 `json:"printingStationary"`
	OtherExpense       float64 `json:"otherExpense"`
	CleaningSupplies   float64 `json:"cleaningSupplies"`
	AirconVentilation  float64 `json:"airconVentilation"`
	Building           float64 `json:"building"`
	Electrical         float64 `json:"electrical"`
	ElectricBulbs      float64 `json:"electricBulbs"`
	ElevEscalators     float64 `json:"elevEscalators"`
	Furniture          float64 `json:"furniture"`
	FbKitchenRefrig    float64 `json:"fbKitchenRefrig"`
	VehicleMaintenance float64 `json:"vehicleMaintenance"`
	PaintingDecoration float64 `json:"paintingDecoration"`
	PlumbingHeating    float64 `json:"plumbingHeating"`
	OfficeEquipment    float64 `json:"officeEquipment"`
	RemovalWasteMatter float64 `json:"removalWasteMatter"`
	LocksKeys          float64 `json:"locksKeys"`
	WaterTreatment     float64 `json:"waterTreatment"`
	EnergyCost         float64 `json:"energyCost"`
}

type Setup struct {
	Pteb       PTEB       `json:"pteb"`
	Cos        COS        `json:"cos"`
	FoFixed    FOFixed    `json:"foFixed"`
	FoVar      FOVar      `json:"foVar"`
	HkFixed    HKFixed    `json:"hkFixed"`
	HkVar      HKVar      `json:"hkVar"`
	FbFixed    FBFixed    `json:"fbFixed"`
	FbVar      FBVar      `json:"fbVar"`
	AgFixed    AGFixed    `json:"agFixed"`
	AgVar      AGVar      `json:"agVar"`
	HrdFixed   HRDFixed   `json:"hrdFixed"`
	HrdVar     HRDVar     `json:"hrdVar"`
	SalesFixed SalesFixed `json:"salesFixed"`
	SalesVar   SalesVar   `json:"salesVar"`
	PomecFixed PomecFixed `json:"pomecFixed"`
	PomecVar   PomecVar   `json:"pomecVar"`
}

type Revenue struct {
	Room        float64 `json:"room"`
	Breakfast   float64 `json:"breakfast"`
	RestoF      float64 `json:"restoF"`
	RestoB      float64 `json:"restoB"`
	TapasF      float64 `json:"tapasF"`
	TapasB      float64 `json:"tapasB"`
	RsF         float64 `json:"rsF"`
	RsB         float64 `json:"rsB"`
	Bqt         float64 `json:"bqt"`
	Spa         float64 `json:"spa"`
	Laundry     float64 `json:"laundry"`
	OtherIncome float64 `json:"otherIncome"`
}

type SavedRevenue struct {
	SavedBy string  `json:"savedBy"`
	SavedAt string  `json:"savedAt"`
	Rev     Revenue `json:"rev"`
}
