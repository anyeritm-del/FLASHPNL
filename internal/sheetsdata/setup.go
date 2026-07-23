package sheetsdata

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

const setupSheetName = "ForecastSetup"

var setupHeader = []interface{}{
	"Year", "Month",
	"PTEB_FO", "PTEB_HK", "PTEB_FB", "PTEB_AG", "PTEB_HRD", "PTEB_Sales", "PTEB_POMEC",
	"COS_Breakfast", "COS_RestoF", "COS_RestoB", "COS_TapasF", "COS_TapasB", "COS_RSF", "COS_RSB", "COS_BQT",
	"FOFixed_Uniforms", "FOFixed_Telephone", "FOFixed_OtherExpense", "FOFixed_TravelAgencyComm",
	"FOFixed_CableSatellite", "FOFixed_SystemsInternet", "FOFixed_OfficerCheck", "FOFixed_Entertainment",
	"FOVar_CleaningSupplies", "FOVar_GuestSupplies", "FOVar_OtherSupplies", "FOVar_Transportation",
	"FOVar_PrintingStationary", "FOVar_GuestTransportation", "FOVar_CompWelcomeDrink", "FOVar_CompFruitBasket",
	"HKFixed_PestControl", "HKFixed_Telephone", "HKFixed_OtherExpense", "HKFixed_ContractHygiene",
	"HKFixed_Entertainment", "HKFixed_OfficerCheck", "HKFixed_Uniforms", "HKFixed_LicensesAndFees", "HKFixed_EquipmentRental",
	"HKVar_Linen", "HKVar_CleaningSupplies", "HKVar_GuestSupplies", "HKVar_OtherSupplies", "HKVar_OtherExpense", "HKVar_LaundryDryCleaning",
	"FBFixed_Menus", "FBFixed_Licenses", "FBFixed_OtherExpense", "FBFixed_BanquetExpense", "FBFixed_ContractService",
	"FBFixed_Entertainment", "FBFixed_OfficerCheck", "FBFixed_Uniforms", "FBFixed_Training", "FBFixed_EquipmentRental", "FBFixed_MusicEntertainment",
	"FBVar_CleaningSupplies", "FBVar_GuestSupplies", "FBVar_KitchenFuel", "FBVar_KitchenSupplies",
	"FBVar_PrintingStationary", "FBVar_MusicEntertainment", "FBVar_Utensil",
	"AGFixed_Uniforms", "AGFixed_Telephone", "AGFixed_LicensesAndFees", "AGFixed_ItSoftware",
	"AGFixed_ItHardware", "AGFixed_CashierOverShort", "AGFixed_BankCharges", "AGFixed_PostageExpress",
	"AGVar_Transportation", "AGVar_PrintingStationary", "AGVar_CreditCardCommissions",
	"HRDFixed_Telephone", "HRDFixed_Transportation", "HRDFixed_Training", "HRDFixed_OtherExpense",
	"HRDFixed_SportSocial", "HRDFixed_EmployeeRelation", "HRDFixed_Outsource", "HRDFixed_License",
	"HRDFixed_Entertainment", "HRDFixed_OfficerCheck",
	"HRDVar_PrintingStationary",
	"SalesFixed_PartnershipSponsors", "SalesFixed_Photography", "SalesFixed_Merchandise", "SalesFixed_Collaterals",
	"SalesFixed_Entertainment", "SalesFixed_Telephone", "SalesFixed_Miscellaneous", "SalesFixed_PostageExpress",
	"SalesVar_LocalSalesCall", "SalesVar_PrintingStationary",
	"POMECFixed_PestControl", "POMECFixed_Telephone", "POMECFixed_OtherExpense", "POMECFixed_OfficerCheck",
	"POMECFixed_Uniforms", "POMECFixed_LicensesAndFees",
	"POMECVar_Uniforms", "POMECVar_OtherSupplies", "POMECVar_PrintingStationary", "POMECVar_OtherExpense",
	"POMECVar_CleaningSupplies", "POMECVar_AirconVentilation", "POMECVar_Building", "POMECVar_Electrical",
	"POMECVar_ElectricBulbs", "POMECVar_ElevEscalators", "POMECVar_Furniture", "POMECVar_FbKitchenRefrig",
	"POMECVar_VehicleMaintenance", "POMECVar_PaintingDecoration", "POMECVar_PlumbingHeating",
	"POMECVar_OfficeEquipment", "POMECVar_RemovalWasteMatter", "POMECVar_LocksKeys",
	"POMECVar_WaterTreatment", "POMECVar_EnergyCost",
}

func rowToSetup(row []interface{}) *Setup {
	return &Setup{
		Pteb: PTEB{
			FO: toFloat(at(row, 2)), HK: toFloat(at(row, 3)), FB: toFloat(at(row, 4)),
			AG: toFloat(at(row, 5)), HRD: toFloat(at(row, 6)), Sales: toFloat(at(row, 7)), Pomec: toFloat(at(row, 8)),
		},
		Cos: COS{
			Breakfast: toFloat(at(row, 9)), RestoF: toFloat(at(row, 10)), RestoB: toFloat(at(row, 11)),
			TapasF: toFloat(at(row, 12)), TapasB: toFloat(at(row, 13)), RsF: toFloat(at(row, 14)),
			RsB: toFloat(at(row, 15)), Bqt: toFloat(at(row, 16)),
		},
		FoFixed: FOFixed{
			Uniforms: toFloat(at(row, 17)), Telephone: toFloat(at(row, 18)),
			OtherExpense: toFloat(at(row, 19)), TravelAgencyComm: toFloat(at(row, 20)),
			CableSatellite: toFloat(at(row, 21)), SystemsInternet: toFloat(at(row, 22)),
			OfficerCheck: toFloat(at(row, 23)), Entertainment: toFloat(at(row, 24)),
		},
		FoVar: FOVar{
			CleaningSupplies: toFloat(at(row, 25)), GuestSupplies: toFloat(at(row, 26)),
			OtherSupplies: toFloat(at(row, 27)), Transportation: toFloat(at(row, 28)),
			PrintingStationary: toFloat(at(row, 29)), GuestTransportation: toFloat(at(row, 30)),
			CompWelcomeDrink: toFloat(at(row, 31)), CompFruitBasket: toFloat(at(row, 32)),
		},
		HkFixed: HKFixed{
			PestControl: toFloat(at(row, 33)), Telephone: toFloat(at(row, 34)),
			OtherExpense: toFloat(at(row, 35)), ContractHygiene: toFloat(at(row, 36)),
			Entertainment: toFloat(at(row, 37)), OfficerCheck: toFloat(at(row, 38)),
			Uniforms: toFloat(at(row, 39)), LicensesAndFees: toFloat(at(row, 40)),
			EquipmentRental: toFloat(at(row, 41)),
		},
		HkVar: HKVar{
			Linen: toFloat(at(row, 42)), CleaningSupplies: toFloat(at(row, 43)),
			GuestSupplies: toFloat(at(row, 44)), OtherSupplies: toFloat(at(row, 45)),
			OtherExpense: toFloat(at(row, 46)), LaundryDryCleaning: toFloat(at(row, 47)),
		},
		FbFixed: FBFixed{
			Menus: toFloat(at(row, 48)), Licenses: toFloat(at(row, 49)),
			OtherExpense: toFloat(at(row, 50)), BanquetExpense: toFloat(at(row, 51)),
			ContractService: toFloat(at(row, 52)), Entertainment: toFloat(at(row, 53)),
			OfficerCheck: toFloat(at(row, 54)), Uniforms: toFloat(at(row, 55)),
			Training: toFloat(at(row, 56)), EquipmentRental: toFloat(at(row, 57)),
			MusicEntertainment: toFloat(at(row, 58)),
		},
		FbVar: FBVar{
			CleaningSupplies: toFloat(at(row, 59)), GuestSupplies: toFloat(at(row, 60)),
			KitchenFuel: toFloat(at(row, 61)), KitchenSupplies: toFloat(at(row, 62)),
			PrintingStationary: toFloat(at(row, 63)), MusicEntertainment: toFloat(at(row, 64)),
			Utensil: toFloat(at(row, 65)),
		},
		AgFixed: AGFixed{
			Uniforms: toFloat(at(row, 66)), Telephone: toFloat(at(row, 67)),
			LicensesAndFees: toFloat(at(row, 68)), ItSoftware: toFloat(at(row, 69)),
			ItHardware: toFloat(at(row, 70)), CashierOverShort: toFloat(at(row, 71)),
			BankCharges: toFloat(at(row, 72)), PostageExpress: toFloat(at(row, 73)),
		},
		AgVar: AGVar{
			Transportation:        toFloat(at(row, 74)),
			PrintingStationary:    toFloat(at(row, 75)),
			CreditCardCommissions: toFloat(at(row, 76)),
		},
		HrdFixed: HRDFixed{
			Telephone: toFloat(at(row, 77)), Transportation: toFloat(at(row, 78)),
			Training: toFloat(at(row, 79)), OtherExpense: toFloat(at(row, 80)),
			SportSocial: toFloat(at(row, 81)), EmployeeRelation: toFloat(at(row, 82)),
			Outsource: toFloat(at(row, 83)), License: toFloat(at(row, 84)),
			Entertainment: toFloat(at(row, 85)), OfficerCheck: toFloat(at(row, 86)),
		},
		HrdVar: HRDVar{
			PrintingStationary: toFloat(at(row, 87)),
		},
		SalesFixed: SalesFixed{
			PartnershipSponsors: toFloat(at(row, 88)), Photography: toFloat(at(row, 89)),
			Merchandise: toFloat(at(row, 90)), Collaterals: toFloat(at(row, 91)),
			Entertainment: toFloat(at(row, 92)), Telephone: toFloat(at(row, 93)),
			Miscellaneous: toFloat(at(row, 94)), PostageExpress: toFloat(at(row, 95)),
		},
		SalesVar: SalesVar{
			LocalSalesCall:     toFloat(at(row, 96)),
			PrintingStationary: toFloat(at(row, 97)),
		},
		PomecFixed: PomecFixed{
			PestControl: toFloat(at(row, 98)), Telephone: toFloat(at(row, 99)),
			OtherExpense: toFloat(at(row, 100)), OfficerCheck: toFloat(at(row, 101)),
			Uniforms: toFloat(at(row, 102)), LicensesAndFees: toFloat(at(row, 103)),
		},
		PomecVar: PomecVar{
			Uniforms: toFloat(at(row, 104)), OtherSupplies: toFloat(at(row, 105)),
			PrintingStationary: toFloat(at(row, 106)), OtherExpense: toFloat(at(row, 107)),
			CleaningSupplies: toFloat(at(row, 108)), AirconVentilation: toFloat(at(row, 109)),
			Building: toFloat(at(row, 110)), Electrical: toFloat(at(row, 111)),
			ElectricBulbs: toFloat(at(row, 112)), ElevEscalators: toFloat(at(row, 113)),
			Furniture: toFloat(at(row, 114)), FbKitchenRefrig: toFloat(at(row, 115)),
			VehicleMaintenance: toFloat(at(row, 116)), PaintingDecoration: toFloat(at(row, 117)),
			PlumbingHeating: toFloat(at(row, 118)), OfficeEquipment: toFloat(at(row, 119)),
			RemovalWasteMatter: toFloat(at(row, 120)), LocksKeys: toFloat(at(row, 121)),
			WaterTreatment: toFloat(at(row, 122)), EnergyCost: toFloat(at(row, 123)),
		},
	}
}

func setupToRow(year, month int, s Setup) []interface{} {
	p, c := s.Pteb, s.Cos
	f, v := s.FoFixed, s.FoVar
	hkf, hkv := s.HkFixed, s.HkVar
	fbf, fbv := s.FbFixed, s.FbVar
	agf, agv := s.AgFixed, s.AgVar
	hrdf, hrdv := s.HrdFixed, s.HrdVar
	sf, sv := s.SalesFixed, s.SalesVar
	pmf, pmv := s.PomecFixed, s.PomecVar

	return []interface{}{
		year, month,
		p.FO, p.HK, p.FB, p.AG, p.HRD, p.Sales, p.Pomec,
		c.Breakfast, c.RestoF, c.RestoB, c.TapasF, c.TapasB, c.RsF, c.RsB, c.Bqt,
		f.Uniforms, f.Telephone, f.OtherExpense, f.TravelAgencyComm,
		f.CableSatellite, f.SystemsInternet, f.OfficerCheck, f.Entertainment,
		v.CleaningSupplies, v.GuestSupplies, v.OtherSupplies, v.Transportation,
		v.PrintingStationary, v.GuestTransportation, v.CompWelcomeDrink, v.CompFruitBasket,
		hkf.PestControl, hkf.Telephone, hkf.OtherExpense, hkf.ContractHygiene,
		hkf.Entertainment, hkf.OfficerCheck, hkf.Uniforms, hkf.LicensesAndFees, hkf.EquipmentRental,
		hkv.Linen, hkv.CleaningSupplies, hkv.GuestSupplies, hkv.OtherSupplies,
		hkv.OtherExpense, hkv.LaundryDryCleaning,
		fbf.Menus, fbf.Licenses, fbf.OtherExpense, fbf.BanquetExpense, fbf.ContractService,
		fbf.Entertainment, fbf.OfficerCheck, fbf.Uniforms, fbf.Training,
		fbf.EquipmentRental, fbf.MusicEntertainment,
		fbv.CleaningSupplies, fbv.GuestSupplies, fbv.KitchenFuel, fbv.KitchenSupplies,
		fbv.PrintingStationary, fbv.MusicEntertainment, fbv.Utensil,
		agf.Uniforms, agf.Telephone, agf.LicensesAndFees, agf.ItSoftware,
		agf.ItHardware, agf.CashierOverShort, agf.BankCharges, agf.PostageExpress,
		agv.Transportation, agv.PrintingStationary, agv.CreditCardCommissions,
		hrdf.Telephone, hrdf.Transportation, hrdf.Training, hrdf.OtherExpense,
		hrdf.SportSocial, hrdf.EmployeeRelation, hrdf.Outsource, hrdf.License,
		hrdf.Entertainment, hrdf.OfficerCheck,
		hrdv.PrintingStationary,
		sf.PartnershipSponsors, sf.Photography, sf.Merchandise, sf.Collaterals,
		sf.Entertainment, sf.Telephone, sf.Miscellaneous, sf.PostageExpress,
		sv.LocalSalesCall, sv.PrintingStationary,
		pmf.PestControl, pmf.Telephone, pmf.OtherExpense, pmf.OfficerCheck,
		pmf.Uniforms, pmf.LicensesAndFees,
		pmv.Uniforms, pmv.OtherSupplies, pmv.PrintingStationary, pmv.OtherExpense,
		pmv.CleaningSupplies, pmv.AirconVentilation, pmv.Building, pmv.Electrical,
		pmv.ElectricBulbs, pmv.ElevEscalators, pmv.Furniture, pmv.FbKitchenRefrig,
		pmv.VehicleMaintenance, pmv.PaintingDecoration, pmv.PlumbingHeating,
		pmv.OfficeEquipment, pmv.RemovalWasteMatter, pmv.LocksKeys,
		pmv.WaterTreatment, pmv.EnergyCost,
	}
}

// GetSetup mirrors getSetup_(): returns nil (no error) if the sheet doesn't
// exist yet or no row matches the given period.
func GetSetup(ctx context.Context, svc *sheets.Service, sheetID string, month, year int) (*Setup, error) {
	resp, err := svc.Spreadsheets.Values.Get(sheetID, setupSheetName).
		ValueRenderOption("UNFORMATTED_VALUE").Context(ctx).Do()
	if err != nil {
		if isMissingSheetErr(err) {
			return nil, nil
		}
		return nil, err
	}
	for i, row := range resp.Values {
		if i == 0 {
			continue
		}
		if toInt(at(row, 0)) == year && toInt(at(row, 1)) == month {
			return rowToSetup(row), nil
		}
	}
	return nil, nil
}

// SaveSetup mirrors saveSetup(): upserts the row for (year, month), creating
// the sheet with its header the first time it's used.
func SaveSetup(ctx context.Context, svc *sheets.Service, sheetID string, month, year int, setup Setup) error {
	if err := ensureSheetWithHeader(ctx, svc, sheetID, setupSheetName, setupHeader); err != nil {
		return err
	}

	resp, err := svc.Spreadsheets.Values.Get(sheetID, setupSheetName).
		ValueRenderOption("UNFORMATTED_VALUE").Context(ctx).Do()
	if err != nil {
		return err
	}

	newRow := setupToRow(year, month, setup)
	for i, row := range resp.Values {
		if i == 0 {
			continue
		}
		if toInt(at(row, 0)) == year && toInt(at(row, 1)) == month {
			rangeA1 := sheetRowRange(setupSheetName, i+1, len(newRow))
			_, err := svc.Spreadsheets.Values.Update(sheetID, rangeA1, &sheets.ValueRange{Values: [][]interface{}{newRow}}).
				ValueInputOption("RAW").Context(ctx).Do()
			return err
		}
	}

	_, err = svc.Spreadsheets.Values.Append(sheetID, setupSheetName, &sheets.ValueRange{Values: [][]interface{}{newRow}}).
		ValueInputOption("RAW").Context(ctx).Do()
	return err
}
