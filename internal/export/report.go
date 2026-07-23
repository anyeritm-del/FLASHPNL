package export

import "fmt"

var monthNames = [...]string{
	"", "Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

// buildReport ports exportFlashPL()'s row-by-row report construction.
func buildReport(month, year int, ssName, generatedAt string, pl PLData) *builder {
	b := &builder{}
	tot := pl.TotalRevenue
	fbRev := pl.FbRevenue

	titleIdx := b.addValues([]interface{}{fmt.Sprintf("Flash P&L Forecast — %s %d", monthNames[month], year), nil, nil})
	b.style(titleIdx, 0, 3, colorNavy, "#ffffff", true, 14)
	b.merge(titleIdx, 3)

	subIdx := b.addValues([]interface{}{fmt.Sprintf("%s  |  Dibuat: %s", ssName, generatedAt), nil, nil})
	b.style(subIdx, 0, 3, "", "#64748b", false, 9)
	b.merge(subIdx, 3)
	b.blank()

	b.wr([]interface{}{"Keterangan", "Nominal (Rp)", "% Rev"}, colorBlueH, true, "#ffffff")

	// A. Revenue
	b.sec("A. Revenue")
	for _, it := range []item{
		{"Rooms", pl.Rev.Room}, {"Breakfast", pl.Rev.Breakfast}, {"Resto Food", pl.Rev.RestoF}, {"Resto Beverage", pl.Rev.RestoB},
		{"Tapas Food", pl.Rev.TapasF}, {"Tapas Beverage", pl.Rev.TapasB}, {"RS Food", pl.Rev.RsF}, {"RS Beverage", pl.Rev.RsB},
		{"BQT + Wedding", pl.Rev.Bqt}, {"Spa", pl.Rev.Spa}, {"Laundry", pl.Rev.Laundry}, {"Other Income", pl.Rev.OtherIncome},
	} {
		b.row(it.Label, it.Val, tot)
	}
	b.subtotal("TOTAL REVENUE", tot, tot)
	b.blank()

	// B. PTEB
	b.sec("B. Payroll & Related Expenses (PTEB)")
	for _, it := range []item{
		{"Front Office", pl.Pteb.FO}, {"Housekeeping", pl.Pteb.HK}, {"Food & Beverage", pl.Pteb.FB},
		{"Accounting / GA", pl.Pteb.AG}, {"HRD", pl.Pteb.HRD}, {"Sales & Marketing", pl.Pteb.Sales}, {"POMEC", pl.Pteb.Pomec},
	} {
		b.row(it.Label, it.Val, tot)
	}
	b.subtotal("TOTAL PTEB", pl.TotalPTEB, tot)
	b.blank()

	// C. COS
	b.sec("C. Cost of Sales (COS)")
	for _, it := range []item{
		{"Breakfast", pl.Cos.Breakfast}, {"Resto Food", pl.Cos.RestoF}, {"Resto Bev", pl.Cos.RestoB}, {"Tapas Food", pl.Cos.TapasF},
		{"Tapas Bev", pl.Cos.TapasB}, {"RS Food", pl.Cos.RsF}, {"RS Bev", pl.Cos.RsB}, {"BQT+Wedding", pl.Cos.Bqt},
	} {
		b.row(it.Label, it.Val, tot)
	}
	b.subtotal("TOTAL COS", pl.TotalCOS, tot)
	b.blank()

	// D. Other Departmental Expenses
	b.sec("D. Other Departmental Expenses")

	renderDeptBreakdown(b, "Front Office",
		[]item{
			{"Uniforms", pl.FoFixed.Uniforms}, {"Telephone & Fax", pl.FoFixed.Telephone}, {"Other Expense", pl.FoFixed.OtherExpense},
			{"Travel Agency Comm.", pl.FoFixed.TravelAgencyComm}, {"Cable & TV Satellite", pl.FoFixed.CableSatellite},
			{"Systems / Internet", pl.FoFixed.SystemsInternet}, {"Officer Check", pl.FoFixed.OfficerCheck}, {"Entertainment", pl.FoFixed.Entertainment},
		},
		[]item{
			{"Cleaning Supplies", pl.FoVar.CleaningSupplies}, {"Guest Supplies", pl.FoVar.GuestSupplies}, {"Other Supplies", pl.FoVar.OtherSupplies},
			{"Transportation", pl.FoVar.Transportation}, {"Printing & Stationary", pl.FoVar.PrintingStationary},
			{"Guest Transportation", pl.FoVar.GuestTransportation}, {"Comp. Welcome Drink", pl.FoVar.CompWelcomeDrink}, {"Comp. Fruit Basket", pl.FoVar.CompFruitBasket},
		},
		pl.Rev.Room, "% Room Rev", pl.FoOtherExp, tot)

	renderDeptBreakdown(b, "Housekeeping",
		[]item{
			{"Pest Control", pl.HkFixed.PestControl}, {"Telephone & Fax", pl.HkFixed.Telephone}, {"Other Expense", pl.HkFixed.OtherExpense},
			{"Contract Hygiene", pl.HkFixed.ContractHygiene}, {"Entertainment", pl.HkFixed.Entertainment}, {"Officer Check", pl.HkFixed.OfficerCheck},
			{"Uniforms", pl.HkFixed.Uniforms}, {"Licenses & Fees", pl.HkFixed.LicensesAndFees}, {"Equipment Rental", pl.HkFixed.EquipmentRental},
		},
		[]item{
			{"Linen", pl.HkVar.Linen}, {"Cleaning Supplies", pl.HkVar.CleaningSupplies}, {"Guest Supplies", pl.HkVar.GuestSupplies},
			{"Other Supplies", pl.HkVar.OtherSupplies}, {"Other Expense", pl.HkVar.OtherExpense}, {"Laundry & Dry Cleaning", pl.HkVar.LaundryDryCleaning},
		},
		pl.Rev.Room, "% Room Rev", pl.HkOtherExp, tot)

	renderDeptBreakdown(b, "Food & Beverage",
		[]item{
			{"Menus", pl.FbFixed.Menus}, {"Licenses", pl.FbFixed.Licenses}, {"Other Expense", pl.FbFixed.OtherExpense},
			{"Banquet Expense", pl.FbFixed.BanquetExpense}, {"Contract Service", pl.FbFixed.ContractService}, {"Entertainment", pl.FbFixed.Entertainment},
			{"Officer Check", pl.FbFixed.OfficerCheck}, {"Uniforms", pl.FbFixed.Uniforms}, {"Training", pl.FbFixed.Training},
			{"Equipment Rental", pl.FbFixed.EquipmentRental}, {"Music Entertainment", pl.FbFixed.MusicEntertainment},
		},
		[]item{
			{"Cleaning Supplies", pl.FbVar.CleaningSupplies}, {"Guest Supplies", pl.FbVar.GuestSupplies}, {"Kitchen Fuel", pl.FbVar.KitchenFuel},
			{"Kitchen Supplies", pl.FbVar.KitchenSupplies}, {"Printing & Stationary", pl.FbVar.PrintingStationary},
			{"Music Entertainment", pl.FbVar.MusicEntertainment}, {"Utensil", pl.FbVar.Utensil},
		},
		fbRev, "% F&B Rev", pl.FbOtherExp, tot)

	renderDeptBreakdown(b, "Accounting / GA",
		[]item{
			{"Uniforms", pl.AgFixed.Uniforms}, {"Telephone & Fax", pl.AgFixed.Telephone}, {"Licenses & Fees", pl.AgFixed.LicensesAndFees},
			{"IT Software", pl.AgFixed.ItSoftware}, {"IT Hardware", pl.AgFixed.ItHardware}, {"Cashier Over/Short", pl.AgFixed.CashierOverShort},
			{"Bank Charges", pl.AgFixed.BankCharges}, {"Postage & Express", pl.AgFixed.PostageExpress},
		},
		[]item{
			{"Transportation", pl.AgVar.Transportation}, {"Printing & Stationary", pl.AgVar.PrintingStationary},
			{"Credit Card Commissions", pl.AgVar.CreditCardCommissions},
		},
		tot, "% Total Rev", pl.AgOtherExp, tot)

	renderDeptBreakdown(b, "HRD",
		[]item{
			{"Telephone & Fax", pl.HrdFixed.Telephone}, {"Transportation", pl.HrdFixed.Transportation}, {"Training", pl.HrdFixed.Training},
			{"Other Expense", pl.HrdFixed.OtherExpense}, {"Sport & Social", pl.HrdFixed.SportSocial}, {"Employee Relation", pl.HrdFixed.EmployeeRelation},
			{"Outsource", pl.HrdFixed.Outsource}, {"License", pl.HrdFixed.License}, {"Entertainment", pl.HrdFixed.Entertainment}, {"Officer Check", pl.HrdFixed.OfficerCheck},
		},
		[]item{
			{"Printing & Stationary", pl.HrdVar.PrintingStationary},
		},
		tot, "% Total Rev", pl.HrdOtherExp, tot)

	renderDeptBreakdown(b, "Sales & Marketing",
		[]item{
			{"Partnership & Sponsors", pl.SalesFixed.PartnershipSponsors}, {"Photography", pl.SalesFixed.Photography},
			{"Merchandise", pl.SalesFixed.Merchandise}, {"Collaterals", pl.SalesFixed.Collaterals}, {"Entertainment", pl.SalesFixed.Entertainment},
			{"Telephone & Fax", pl.SalesFixed.Telephone}, {"Miscellaneous", pl.SalesFixed.Miscellaneous}, {"Postage & Express", pl.SalesFixed.PostageExpress},
		},
		[]item{
			{"Local Sales Call", pl.SalesVar.LocalSalesCall}, {"Printing & Stationary", pl.SalesVar.PrintingStationary},
		},
		tot, "% Total Rev", pl.SalesOtherExp, tot)

	renderDeptBreakdown(b, "POMEC",
		[]item{
			{"Pest Control", pl.PomecFixed.PestControl}, {"Telephone & Fax", pl.PomecFixed.Telephone}, {"Other Expense", pl.PomecFixed.OtherExpense},
			{"Officer Check", pl.PomecFixed.OfficerCheck}, {"Uniforms", pl.PomecFixed.Uniforms}, {"Licenses & Fees", pl.PomecFixed.LicensesAndFees},
		},
		[]item{
			{"Uniforms", pl.PomecVar.Uniforms}, {"Other Supplies", pl.PomecVar.OtherSupplies}, {"Printing & Stationary", pl.PomecVar.PrintingStationary},
			{"Other Expense", pl.PomecVar.OtherExpense}, {"Cleaning Supplies", pl.PomecVar.CleaningSupplies}, {"Aircon & Ventilation", pl.PomecVar.AirconVentilation},
			{"Building", pl.PomecVar.Building}, {"Electrical", pl.PomecVar.Electrical}, {"Electric Bulbs", pl.PomecVar.ElectricBulbs},
			{"Elev. & Escalators", pl.PomecVar.ElevEscalators}, {"Furniture", pl.PomecVar.Furniture}, {"F&B Kitchen & Refrig.", pl.PomecVar.FbKitchenRefrig},
			{"Vehicle Maintenance", pl.PomecVar.VehicleMaintenance}, {"Painting & Decoration", pl.PomecVar.PaintingDecoration},
			{"Plumbing & Heating", pl.PomecVar.PlumbingHeating}, {"Office Equipment", pl.PomecVar.OfficeEquipment},
			{"Removal & Waste Matter", pl.PomecVar.RemovalWasteMatter}, {"Locks & Keys", pl.PomecVar.LocksKeys},
			{"Water Treatment", pl.PomecVar.WaterTreatment}, {"Energy Cost", pl.PomecVar.EnergyCost},
		},
		tot, "% Total Rev", pl.PomecOtherExp, tot)

	b.subtotal("TOTAL OTHER EXPENSES", pl.TotalOtherExp, tot)
	b.blank()

	// E. Management Fee
	b.sec("E. Management Fee")
	b.row("Management Fee (1,11% × Room Revenue)", pl.ManagementFee, tot)
	b.blank()

	// GOP
	gopIdx := b.wr([]interface{}{"GROSS OPERATING PROFIT (GOP)", pl.Gop, pctOf(pl.Gop, tot)}, colorNavy, true, "#ffffff")
	b.style(gopIdx, 1, 2, "", "#4ade80", false, 0)
	b.style(gopIdx, 2, 3, "", "#4ade80", false, 0)
	b.numberFormat(gopIdx, 2, "0.0%")

	return b
}
