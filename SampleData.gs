// ============================================================
// Flash P&L Forecast — Sample Data Generator
// Jalankan createSampleData() sekali dari Apps Script Editor
// Menu: Run > createSampleData
// ============================================================

function createSampleData() {
  var ss = SpreadsheetApp.getActiveSpreadsheet();
  createUsersSheet_(ss);
  createForecastSetupSheet_(ss);
  createForecastRevenueSheet_(ss);
  SpreadsheetApp.getUi().alert(
    '✅ Sample data berhasil dibuat!\n\n' +
    'Sheet yang dibuat:\n' +
    '- Users (accounting + management)\n' +
    '- ForecastSetup (Juli 2026 — PTEB, COS%, breakdown semua dept)\n' +
    '- ForecastRevenue (Juli 2026 — revenue aktual dari PDF)\n\n' +
    'Ganti email di sheet Users dengan email Google Anda.'
  );
}

// ── USERS ─────────────────────────────────────────────────────────
function createUsersSheet_(ss) {
  var sheet = ss.getSheetByName('Users');
  if (sheet) sheet.clearContents(); else sheet = ss.insertSheet('Users');

  sheet.getRange(1, 1, 1, 3).setValues([['Email', 'Role', 'Name']])
    .setFontWeight('bold').setBackground('#0d1f4e').setFontColor('#ffffff');

  var data = [
    ['accounting@archipelagohotels.com', 'accounting', 'Dewi Susanti'],
    ['gm@archipelagohotels.com',         'management', 'Budi Santoso'],
  ];
  sheet.getRange(2, 1, data.length, 3).setValues(data);
  sheet.setColumnWidth(1, 280);
  sheet.setColumnWidth(2, 120);
  sheet.setColumnWidth(3, 200);
}

// ── FORECAST SETUP ────────────────────────────────────────────────
function createForecastSetupSheet_(ss) {
  var sheet = ss.getSheetByName('ForecastSetup');
  if (sheet) sheet.clearContents(); else sheet = ss.insertSheet('ForecastSetup');

  // 124-column schema
  var header = [[
    'Year','Month',
    // PTEB (col 3–9)
    'PTEB_FO','PTEB_HK','PTEB_FB','PTEB_AG','PTEB_HRD','PTEB_Sales','PTEB_POMEC',
    // COS% (col 10–17)
    'COS_Breakfast','COS_RestoF','COS_RestoB','COS_TapasF','COS_TapasB','COS_RSF','COS_RSB','COS_BQT',
    // FO Fixed (col 18–25)
    'FOFixed_Uniforms','FOFixed_Telephone','FOFixed_OtherExpense','FOFixed_TravelAgencyComm',
    'FOFixed_CableSatellite','FOFixed_SystemsInternet','FOFixed_OfficerCheck','FOFixed_Entertainment',
    // FO Var% of Room Rev (col 26–33)
    'FOVar_CleaningSupplies','FOVar_GuestSupplies','FOVar_OtherSupplies','FOVar_Transportation',
    'FOVar_PrintingStationary','FOVar_GuestTransportation','FOVar_CompWelcomeDrink','FOVar_CompFruitBasket',
    // HK Fixed (col 34–42)
    'HKFixed_PestControl','HKFixed_Telephone','HKFixed_OtherExpense','HKFixed_ContractHygiene',
    'HKFixed_Entertainment','HKFixed_OfficerCheck','HKFixed_Uniforms','HKFixed_LicensesAndFees','HKFixed_EquipmentRental',
    // HK Var% of Room Rev (col 43–48)
    'HKVar_Linen','HKVar_CleaningSupplies','HKVar_GuestSupplies','HKVar_OtherSupplies','HKVar_OtherExpense','HKVar_LaundryDryCleaning',
    // FB Fixed (col 49–59)
    'FBFixed_Menus','FBFixed_Licenses','FBFixed_OtherExpense','FBFixed_BanquetExpense','FBFixed_ContractService',
    'FBFixed_Entertainment','FBFixed_OfficerCheck','FBFixed_Uniforms','FBFixed_Training','FBFixed_EquipmentRental','FBFixed_MusicEntertainment',
    // FB Var% of F&B Rev (col 60–66)
    'FBVar_CleaningSupplies','FBVar_GuestSupplies','FBVar_KitchenFuel','FBVar_KitchenSupplies',
    'FBVar_PrintingStationary','FBVar_MusicEntertainment','FBVar_Utensil',
    // AG Fixed (col 67–74)
    'AGFixed_Uniforms','AGFixed_Telephone','AGFixed_LicensesAndFees','AGFixed_ItSoftware',
    'AGFixed_ItHardware','AGFixed_CashierOverShort','AGFixed_BankCharges','AGFixed_PostageExpress',
    // AG Var% of Total Rev (col 75–77)
    'AGVar_Transportation','AGVar_PrintingStationary','AGVar_CreditCardCommissions',
    // HRD Fixed (col 78–87)
    'HRDFixed_Telephone','HRDFixed_Transportation','HRDFixed_Training','HRDFixed_OtherExpense',
    'HRDFixed_SportSocial','HRDFixed_EmployeeRelation','HRDFixed_Outsource','HRDFixed_License',
    'HRDFixed_Entertainment','HRDFixed_OfficerCheck',
    // HRD Var% of Total Rev (col 88)
    'HRDVar_PrintingStationary',
    // Sales Fixed (col 89–96)
    'SalesFixed_PartnershipSponsors','SalesFixed_Photography','SalesFixed_Merchandise','SalesFixed_Collaterals',
    'SalesFixed_Entertainment','SalesFixed_Telephone','SalesFixed_Miscellaneous','SalesFixed_PostageExpress',
    // Sales Var% of Total Rev (col 97–98)
    'SalesVar_LocalSalesCall','SalesVar_PrintingStationary',
    // POMEC Fixed (col 99–104)
    'POMECFixed_PestControl','POMECFixed_Telephone','POMECFixed_OtherExpense','POMECFixed_OfficerCheck',
    'POMECFixed_Uniforms','POMECFixed_LicensesAndFees',
    // POMEC Var% of Total Rev (col 105–124)
    'POMECVar_Uniforms','POMECVar_OtherSupplies','POMECVar_PrintingStationary','POMECVar_OtherExpense',
    'POMECVar_CleaningSupplies','POMECVar_AirconVentilation','POMECVar_Building','POMECVar_Electrical',
    'POMECVar_ElectricBulbs','POMECVar_ElevEscalators','POMECVar_Furniture','POMECVar_FbKitchenRefrig',
    'POMECVar_VehicleMaintenance','POMECVar_PaintingDecoration','POMECVar_PlumbingHeating',
    'POMECVar_OfficeEquipment','POMECVar_RemovalWasteMatter','POMECVar_LocksKeys',
    'POMECVar_WaterTreatment','POMECVar_EnergyCost'
  ]];
  sheet.getRange(1, 1, 1, 124).setValues(header)
    .setFontWeight('bold').setBackground('#0d1f4e').setFontColor('#ffffff');

  // ── Juli 2026 — data aktual dari PDF ──────────────────────────
  //
  // Revenue base untuk variable costs:
  //   Room Revenue:  Rp 1.897.356.748
  //   F&B Revenue:   Rp   769.234.769  (sum 8 F&B streams)
  //   Total Revenue: Rp 2.737.286.541
  //
  // HK Fixed total:    Rp  14.000.000
  // HK Var (2.67% Room Rev): ≈ Rp 50.659.000 → HK total ≈ Rp 64.7M
  //
  // FB Fixed total:    Rp  15.500.000
  // FB Var (5.00% F&B Rev):  ≈ Rp 38.462.000 → FB total ≈ Rp 53.9M
  //
  // AG Fixed total:    Rp  30.500.000
  // AG Var (1.24% Total Rev): ≈ Rp 33.942.000 → AG total ≈ Rp 64.4M
  //
  // HRD Fixed total:   Rp  22.500.000
  // HRD Var (0.04% Total Rev): ≈ Rp 1.095.000 → HRD total ≈ Rp 23.6M
  //
  // Sales Fixed total: Rp  30.500.000
  // Sales Var (0.25% Total Rev): ≈ Rp 6.843.000 → Sales total ≈ Rp 37.3M
  //
  // POMEC Fixed total: Rp  15.000.000
  // POMEC Var (5.16% Total Rev): ≈ Rp 141.244.000 → POMEC total ≈ Rp 156.2M

  var data = [
    [2026, 7,
      // PTEB (Rp): FO, HK, FB, AG, HRD, Sales, POMEC
      155000000, 190000000, 245000000, 108000000, 55000000, 65000000, 58000000,
      // COS% (whole number, e.g. 40 = 40%): Breakfast,RestoF,RestoB,TapasF,TapasB,RSF,RSB,BQT
      40, 33, 28, 34, 34, 33, 28, 29,
      // FO Fixed (Rp): Uniforms,Telephone,OtherExpense,TravelAgencyComm,CableSatellite,SystemsInternet,OfficerCheck,Entertainment
      500000, 5000000, 1000000, 15000000, 3000000, 5000000, 3000000, 0,
      // FO Var% (% of Room Rev): CleaningSupplies,GuestSupplies,OtherSupplies,Transportation,PrintingStationary,GuestTransportation,CompWelcomeDrink,CompFruitBasket
      0.01, 0.08, 0.04, 0.15, 0.12, 0.30, 0.02, 0.04,
      // HK Fixed (Rp): PestControl,Telephone,OtherExpense,ContractHygiene,Entertainment,OfficerCheck,Uniforms,LicensesAndFees,EquipmentRental
      1500000, 1000000, 1000000, 5000000, 0, 2000000, 1000000, 500000, 2000000,
      // HK Var% (% of Room Rev): Linen,CleaningSupplies,GuestSupplies,OtherSupplies,OtherExpense,LaundryDryCleaning
      0.50, 0.40, 1.70, 0.01, 0.02, 0.04,
      // FB Fixed (Rp): Menus,Licenses,OtherExpense,BanquetExpense,ContractService,Entertainment,OfficerCheck,Uniforms,Training,EquipmentRental,MusicEntertainment
      500000, 1000000, 1000000, 2000000, 3000000, 0, 2000000, 1000000, 1000000, 1000000, 3000000,
      // FB Var% (% of F&B Rev): CleaningSupplies,GuestSupplies,KitchenFuel,KitchenSupplies,PrintingStationary,MusicEntertainment,Utensil
      0.53, 1.57, 1.83, 0.02, 0.82, 0.03, 0.20,
      // AG Fixed (Rp): Uniforms,Telephone,LicensesAndFees,ItSoftware,ItHardware,CashierOverShort,BankCharges,PostageExpress
      1000000, 3000000, 3000000, 12000000, 5000000, 1000000, 5000000, 500000,
      // AG Var% (% of Total Rev): Transportation,PrintingStationary,CreditCardCommissions
      0.15, 0.09, 1.00,
      // HRD Fixed (Rp): Telephone,Transportation,Training,OtherExpense,SportSocial,EmployeeRelation,Outsource,License,Entertainment,OfficerCheck
      1000000, 2000000, 5000000, 1000000, 3000000, 3000000, 5000000, 500000, 0, 2000000,
      // HRD Var% (% of Total Rev): PrintingStationary
      0.04,
      // Sales Fixed (Rp): PartnershipSponsors,Photography,Merchandise,Collaterals,Entertainment,Telephone,Miscellaneous,PostageExpress
      10000000, 3000000, 5000000, 5000000, 3000000, 2000000, 2000000, 500000,
      // Sales Var% (% of Total Rev): LocalSalesCall,PrintingStationary
      0.20, 0.05,
      // POMEC Fixed (Rp): PestControl,Telephone,OtherExpense,OfficerCheck,Uniforms,LicensesAndFees
      5000000, 2000000, 1000000, 3000000, 1000000, 3000000,
      // POMEC Var% (% of Total Rev): 20 items
      0.03, 0.15, 0.04, 0.01, 0.02, 0.04, 0.05, 0.03, 0.05, 0.20,
      0.02, 0.04, 0.05, 0.13, 0.01, 0.08, 0.18, 0.01, 0.02, 4.00
    ]
  ];

  sheet.getRange(2, 1, 1, 124).setValues(data);

  // Number formats
  sheet.getRange(2, 3,  1, 7).setNumberFormat('#,##0');     // PTEB
  sheet.getRange(2, 10, 1, 8).setNumberFormat('0.00"%"');   // COS%
  sheet.getRange(2, 18, 1, 8).setNumberFormat('#,##0');     // FO Fixed
  sheet.getRange(2, 26, 1, 8).setNumberFormat('0.00"%"');   // FO Var%
  sheet.getRange(2, 34, 1, 9).setNumberFormat('#,##0');     // HK Fixed
  sheet.getRange(2, 43, 1, 6).setNumberFormat('0.00"%"');   // HK Var%
  sheet.getRange(2, 49, 1, 11).setNumberFormat('#,##0');    // FB Fixed
  sheet.getRange(2, 60, 1, 7).setNumberFormat('0.00"%"');   // FB Var%
  sheet.getRange(2, 67, 1, 8).setNumberFormat('#,##0');     // AG Fixed
  sheet.getRange(2, 75, 1, 3).setNumberFormat('0.00"%"');   // AG Var%
  sheet.getRange(2, 78, 1, 10).setNumberFormat('#,##0');    // HRD Fixed
  sheet.getRange(2, 88, 1, 1).setNumberFormat('0.00"%"');   // HRD Var%
  sheet.getRange(2, 89, 1, 8).setNumberFormat('#,##0');     // Sales Fixed
  sheet.getRange(2, 97, 1, 2).setNumberFormat('0.00"%"');   // Sales Var%
  sheet.getRange(2, 99, 1, 6).setNumberFormat('#,##0');     // POMEC Fixed
  sheet.getRange(2, 105, 1, 20).setNumberFormat('0.00"%"'); // POMEC Var%

  for (var c = 1; c <= 124; c++) sheet.autoResizeColumn(c);
}

// ── FORECAST REVENUE ─────────────────────────────────────────────
function createForecastRevenueSheet_(ss) {
  var sheet = ss.getSheetByName('ForecastRevenue');
  if (sheet) sheet.clearContents(); else sheet = ss.insertSheet('ForecastRevenue');

  sheet.getRange(1, 1, 1, 16).setValues([[
    'Year','Month','SavedBy','SavedAt',
    'Room','Breakfast','RestoFood','RestoBev',
    'TapasFood','TapasBev','RSFood','RSBev',
    'BQT','Spa','Laundry','OtherIncome'
  ]]).setFontWeight('bold').setBackground('#0d1f4e').setFontColor('#ffffff');

  // Revenue Juli 2026 dari PDF — Total Rp 2.737.286.541
  var data = [
    [2026, 7, 'accounting@archipelagohotels.com', '23 Jun 2026 09:00',
      1897356748,  // Room
      300000000,   // Breakfast
      46000000,    // Resto Food
      36000000,    // Resto Beverage
      21470149,    // Tapas Food
      27606108,    // Tapas Beverage
      33402435,    // RS Food
      6781683,     // RS Beverage
      297984874,   // BQT + Wedding
      21963812,    // Spa
      28846352,    // Laundry
      19874380     // Other Income
    ]
  ];

  sheet.getRange(2, 1, 1, 16).setValues(data);
  sheet.setColumnWidth(3, 260);
  sheet.setColumnWidth(4, 160);
  sheet.getRange(2, 5, 1, 12).setNumberFormat('#,##0');
  for (var c = 1; c <= 4; c++) sheet.autoResizeColumn(c);
}
