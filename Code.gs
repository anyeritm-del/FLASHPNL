
// ============================================================
// Flash P&L Forecast — Google Apps Script Backend
// ============================================================
// Sheet "Users":           Email | Role (accounting/management) | Name
// Sheet "ForecastSetup":   124 columns:
//   Col   1-2:   Year, Month
//   Col   3-9:   PTEB×7 (FO,HK,FB,AG,HRD,Sales,POMEC)
//   Col  10-17:  COS%×8 (Breakfast,RestoF,RestoB,TapasF,TapasB,RSF,RSB,BQT)
//   Col  18-25:  FOFixed×8  (Uniforms,Telephone,OtherExpense,TravelAgencyComm,
//                             CableSatellite,SystemsInternet,OfficerCheck,Entertainment)
//   Col  26-33:  FOVar%×8   (CleaningSupplies,GuestSupplies,OtherSupplies,Transportation,
//                             PrintingStationary,GuestTransportation,CompWelcomeDrink,CompFruitBasket)
//   Col  34-42:  HKFixed×9  (PestControl,Telephone,OtherExpense,ContractHygiene,
//                             Entertainment,OfficerCheck,Uniforms,LicensesAndFees,EquipmentRental)
//   Col  43-48:  HKVar%×6   (Linen,CleaningSupplies,GuestSupplies,OtherSupplies,OtherExpense,LaundryDryCleaning)
//   Col  49-59:  FBFixed×11 (Menus,Licenses,OtherExpense,BanquetExpense,ContractService,
//                             Entertainment,OfficerCheck,Uniforms,Training,EquipmentRental,MusicEntertainment)
//   Col  60-66:  FBVar%×7   (CleaningSupplies,GuestSupplies,KitchenFuel,KitchenSupplies,
//                             PrintingStationary,MusicEntertainment,Utensil)
//   Col  67-74:  AGFixed×8  (Uniforms,Telephone,LicensesAndFees,ItSoftware,
//                             ItHardware,CashierOverShort,BankCharges,PostageExpress)
//   Col  75-77:  AGVar%×3   (Transportation,PrintingStationary,CreditCardCommissions)
//   Col  78-87:  HRDFixed×10(Telephone,Transportation,Training,OtherExpense,
//                             SportSocial,EmployeeRelation,Outsource,License,Entertainment,OfficerCheck)
//   Col  88:     HRDVar%×1  (PrintingStationary)
//   Col  89-96:  SalesFixed×8(PartnershipSponsors,Photography,Merchandise,Collaterals,
//                              Entertainment,Telephone,Miscellaneous,PostageExpress)
//   Col  97-98:  SalesVar%×2 (LocalSalesCall,PrintingStationary)
//   Col  99-104: POMECFixed×6(PestControl,Telephone,OtherExpense,OfficerCheck,Uniforms,LicensesAndFees)
//   Col 105-124: POMECVar%×20(Uniforms,OtherSupplies,PrintingStationary,OtherExpense,CleaningSupplies,
//                              AirconVentilation,Building,Electrical,ElectricBulbs,ElevEscalators,
//                              Furniture,FbKitchenRefrig,VehicleMaintenance,PaintingDecoration,
//                              PlumbingHeating,OfficeEquipment,RemovalWasteMatter,LocksKeys,
//                              WaterTreatment,EnergyCost)
// Sheet "ForecastRevenue": Year|Month|SavedBy|SavedAt|12 revenue streams (16 cols)
// ============================================================

function doGet() {
  return HtmlService.createHtmlOutputFromFile('index')
    .setTitle('Flash P&L Forecast')
    .setXFrameOptionsMode(HtmlService.XFrameOptionsMode.ALLOWALL);
}

function getUserRole_() {
  var email = Session.getActiveUser().getEmail().toLowerCase();
  var ss = SpreadsheetApp.getActiveSpreadsheet();
  var sheet = ss.getSheetByName('Users');
  if (!sheet) throw new Error('Sheet "Users" tidak ditemukan. Jalankan createSampleData() terlebih dahulu.');
  var rows = sheet.getDataRange().getValues();
  for (var i = 1; i < rows.length; i++) {
    if (rows[i][0] && rows[i][0].toString().toLowerCase() === email) {
      return { email: email, role: rows[i][1].toString().toLowerCase(), name: rows[i][2] ? rows[i][2].toString() : email };
    }
  }
  return { email: email, role: 'unknown', name: email };
}

function getSetup_(ss, month, year) {
  var sheet = ss.getSheetByName('ForecastSetup');
  if (!sheet) return null;
  var rows = sheet.getDataRange().getValues();
  for (var i = 1; i < rows.length; i++) {
    if (+rows[i][0] === year && +rows[i][1] === month) {
      return {
        pteb: {
          fo:    +rows[i][2]||0,  hk:    +rows[i][3]||0,  fb:  +rows[i][4]||0,
          ag:    +rows[i][5]||0,  hrd:   +rows[i][6]||0,  sales:+rows[i][7]||0, pomec:+rows[i][8]||0
        },
        cos: {
          breakfast:+rows[i][9]||0,  restoF:+rows[i][10]||0, restoB:+rows[i][11]||0,
          tapasF:   +rows[i][12]||0, tapasB:+rows[i][13]||0, rsF:   +rows[i][14]||0,
          rsB:      +rows[i][15]||0, bqt:   +rows[i][16]||0
        },
        foFixed: {
          uniforms:         +rows[i][17]||0, telephone:        +rows[i][18]||0,
          otherExpense:     +rows[i][19]||0, travelAgencyComm: +rows[i][20]||0,
          cableSatellite:   +rows[i][21]||0, systemsInternet:  +rows[i][22]||0,
          officerCheck:     +rows[i][23]||0, entertainment:    +rows[i][24]||0
        },
        foVar: {
          cleaningSupplies:   +rows[i][25]||0, guestSupplies:       +rows[i][26]||0,
          otherSupplies:      +rows[i][27]||0, transportation:      +rows[i][28]||0,
          printingStationary: +rows[i][29]||0, guestTransportation: +rows[i][30]||0,
          compWelcomeDrink:   +rows[i][31]||0, compFruitBasket:     +rows[i][32]||0
        },
        hkFixed: {
          pestControl:    +rows[i][33]||0, telephone:       +rows[i][34]||0,
          otherExpense:   +rows[i][35]||0, contractHygiene: +rows[i][36]||0,
          entertainment:  +rows[i][37]||0, officerCheck:    +rows[i][38]||0,
          uniforms:       +rows[i][39]||0, licensesAndFees: +rows[i][40]||0,
          equipmentRental:+rows[i][41]||0
        },
        hkVar: {
          linen:             +rows[i][42]||0, cleaningSupplies:  +rows[i][43]||0,
          guestSupplies:     +rows[i][44]||0, otherSupplies:     +rows[i][45]||0,
          otherExpense:      +rows[i][46]||0, laundryDryCleaning:+rows[i][47]||0
        },
        fbFixed: {
          menus:             +rows[i][48]||0, licenses:          +rows[i][49]||0,
          otherExpense:      +rows[i][50]||0, banquetExpense:    +rows[i][51]||0,
          contractService:   +rows[i][52]||0, entertainment:     +rows[i][53]||0,
          officerCheck:      +rows[i][54]||0, uniforms:          +rows[i][55]||0,
          training:          +rows[i][56]||0, equipmentRental:   +rows[i][57]||0,
          musicEntertainment:+rows[i][58]||0
        },
        fbVar: {
          cleaningSupplies:  +rows[i][59]||0, guestSupplies:     +rows[i][60]||0,
          kitchenFuel:       +rows[i][61]||0, kitchenSupplies:   +rows[i][62]||0,
          printingStationary:+rows[i][63]||0, musicEntertainment:+rows[i][64]||0,
          utensil:           +rows[i][65]||0
        },
        agFixed: {
          uniforms:         +rows[i][66]||0, telephone:       +rows[i][67]||0,
          licensesAndFees:  +rows[i][68]||0, itSoftware:      +rows[i][69]||0,
          itHardware:       +rows[i][70]||0, cashierOverShort:+rows[i][71]||0,
          bankCharges:      +rows[i][72]||0, postageExpress:  +rows[i][73]||0
        },
        agVar: {
          transportation:       +rows[i][74]||0,
          printingStationary:   +rows[i][75]||0,
          creditCardCommissions:+rows[i][76]||0
        },
        hrdFixed: {
          telephone:       +rows[i][77]||0, transportation:  +rows[i][78]||0,
          training:        +rows[i][79]||0, otherExpense:    +rows[i][80]||0,
          sportSocial:     +rows[i][81]||0, employeeRelation:+rows[i][82]||0,
          outsource:       +rows[i][83]||0, license:         +rows[i][84]||0,
          entertainment:   +rows[i][85]||0, officerCheck:    +rows[i][86]||0
        },
        hrdVar: {
          printingStationary:+rows[i][87]||0
        },
        salesFixed: {
          partnershipSponsors:+rows[i][88]||0, photography:   +rows[i][89]||0,
          merchandise:        +rows[i][90]||0, collaterals:   +rows[i][91]||0,
          entertainment:      +rows[i][92]||0, telephone:     +rows[i][93]||0,
          miscellaneous:      +rows[i][94]||0, postageExpress:+rows[i][95]||0
        },
        salesVar: {
          localSalesCall:    +rows[i][96]||0,
          printingStationary:+rows[i][97]||0
        },
        pomecFixed: {
          pestControl:    +rows[i][98]||0,  telephone:      +rows[i][99]||0,
          otherExpense:   +rows[i][100]||0, officerCheck:   +rows[i][101]||0,
          uniforms:       +rows[i][102]||0, licensesAndFees:+rows[i][103]||0
        },
        pomecVar: {
          uniforms:          +rows[i][104]||0, otherSupplies:      +rows[i][105]||0,
          printingStationary:+rows[i][106]||0, otherExpense:       +rows[i][107]||0,
          cleaningSupplies:  +rows[i][108]||0, airconVentilation:  +rows[i][109]||0,
          building:          +rows[i][110]||0, electrical:         +rows[i][111]||0,
          electricBulbs:     +rows[i][112]||0, elevEscalators:     +rows[i][113]||0,
          furniture:         +rows[i][114]||0, fbKitchenRefrig:    +rows[i][115]||0,
          vehicleMaintenance:+rows[i][116]||0, paintingDecoration: +rows[i][117]||0,
          plumbingHeating:   +rows[i][118]||0, officeEquipment:    +rows[i][119]||0,
          removalWasteMatter:+rows[i][120]||0, locksKeys:          +rows[i][121]||0,
          waterTreatment:    +rows[i][122]||0, energyCost:         +rows[i][123]||0
        }
      };
    }
  }
  return null;
}

function getRevenue_(ss, month, year) {
  var sheet = ss.getSheetByName('ForecastRevenue');
  if (!sheet) return null;
  var rows = sheet.getDataRange().getValues();
  for (var i = 1; i < rows.length; i++) {
    if (+rows[i][0] === year && +rows[i][1] === month) {
      return {
        savedBy: rows[i][2] ? rows[i][2].toString() : '',
        savedAt: rows[i][3] ? rows[i][3].toString() : '',
        rev: {
          room:       +rows[i][4]||0,  breakfast:  +rows[i][5]||0,
          restoF:     +rows[i][6]||0,  restoB:     +rows[i][7]||0,
          tapasF:     +rows[i][8]||0,  tapasB:     +rows[i][9]||0,
          rsF:        +rows[i][10]||0, rsB:        +rows[i][11]||0,
          bqt:        +rows[i][12]||0, spa:        +rows[i][13]||0,
          laundry:    +rows[i][14]||0, otherIncome:+rows[i][15]||0
        }
      };
    }
  }
  return null;
}

function initApp(month, year) {
  var userInfo = getUserRole_();
  var ss = SpreadsheetApp.getActiveSpreadsheet();
  return {
    role:         userInfo.role,
    user:         userInfo,
    hotelName:    ss.getName(),
    setup:        getSetup_(ss, month, year),
    savedRevenue: getRevenue_(ss, month, year)
  };
}

function loadPeriodData(month, year) {
  var ss = SpreadsheetApp.getActiveSpreadsheet();
  return {
    setup:        getSetup_(ss, month, year),
    savedRevenue: getRevenue_(ss, month, year)
  };
}

function saveRevenueForecast(month, year, revData) {
  var userInfo = getUserRole_();
  var ss       = SpreadsheetApp.getActiveSpreadsheet();
  var sheet    = ss.getSheetByName('ForecastRevenue');
  if (!sheet) {
    sheet = ss.insertSheet('ForecastRevenue');
    sheet.getRange(1, 1, 1, 16).setValues([[
      'Year','Month','SavedBy','SavedAt',
      'Room','Breakfast','RestoFood','RestoBev',
      'TapasFood','TapasBev','RSFood','RSBev',
      'BQT','Spa','Laundry','OtherIncome'
    ]]).setFontWeight('bold').setBackground('#0d1f4e').setFontColor('#ffffff');
  }

  var timestamp = Utilities.formatDate(new Date(), Session.getScriptTimeZone(), 'dd MMM yyyy HH:mm');
  var newRow = [
    year, month, userInfo.email, timestamp,
    revData.room||0, revData.breakfast||0, revData.restoF||0, revData.restoB||0,
    revData.tapasF||0, revData.tapasB||0, revData.rsF||0, revData.rsB||0,
    revData.bqt||0, revData.spa||0, revData.laundry||0, revData.otherIncome||0
  ];

  var rows = sheet.getDataRange().getValues();
  for (var i = 1; i < rows.length; i++) {
    if (+rows[i][0] === year && +rows[i][1] === month) {
      sheet.getRange(i + 1, 1, 1, 16).setValues([newRow]);
      return timestamp;
    }
  }
  sheet.appendRow(newRow);
  return timestamp;
}

function saveSetup(month, year, setupData) {
  var userInfo = getUserRole_();
  if (userInfo.role !== 'accounting') throw new Error('Akses ditolak: hanya role Accounting yang dapat mengubah setup.');

  var ss    = SpreadsheetApp.getActiveSpreadsheet();
  var sheet = ss.getSheetByName('ForecastSetup');
  if (!sheet) {
    sheet = ss.insertSheet('ForecastSetup');
    sheet.getRange(1, 1, 1, 124).setValues([[
      'Year','Month',
      'PTEB_FO','PTEB_HK','PTEB_FB','PTEB_AG','PTEB_HRD','PTEB_Sales','PTEB_POMEC',
      'COS_Breakfast','COS_RestoF','COS_RestoB','COS_TapasF','COS_TapasB','COS_RSF','COS_RSB','COS_BQT',
      'FOFixed_Uniforms','FOFixed_Telephone','FOFixed_OtherExpense','FOFixed_TravelAgencyComm',
      'FOFixed_CableSatellite','FOFixed_SystemsInternet','FOFixed_OfficerCheck','FOFixed_Entertainment',
      'FOVar_CleaningSupplies','FOVar_GuestSupplies','FOVar_OtherSupplies','FOVar_Transportation',
      'FOVar_PrintingStationary','FOVar_GuestTransportation','FOVar_CompWelcomeDrink','FOVar_CompFruitBasket',
      'HKFixed_PestControl','HKFixed_Telephone','HKFixed_OtherExpense','HKFixed_ContractHygiene',
      'HKFixed_Entertainment','HKFixed_OfficerCheck','HKFixed_Uniforms','HKFixed_LicensesAndFees','HKFixed_EquipmentRental',
      'HKVar_Linen','HKVar_CleaningSupplies','HKVar_GuestSupplies','HKVar_OtherSupplies','HKVar_OtherExpense','HKVar_LaundryDryCleaning',
      'FBFixed_Menus','FBFixed_Licenses','FBFixed_OtherExpense','FBFixed_BanquetExpense','FBFixed_ContractService',
      'FBFixed_Entertainment','FBFixed_OfficerCheck','FBFixed_Uniforms','FBFixed_Training','FBFixed_EquipmentRental','FBFixed_MusicEntertainment',
      'FBVar_CleaningSupplies','FBVar_GuestSupplies','FBVar_KitchenFuel','FBVar_KitchenSupplies',
      'FBVar_PrintingStationary','FBVar_MusicEntertainment','FBVar_Utensil',
      'AGFixed_Uniforms','AGFixed_Telephone','AGFixed_LicensesAndFees','AGFixed_ItSoftware',
      'AGFixed_ItHardware','AGFixed_CashierOverShort','AGFixed_BankCharges','AGFixed_PostageExpress',
      'AGVar_Transportation','AGVar_PrintingStationary','AGVar_CreditCardCommissions',
      'HRDFixed_Telephone','HRDFixed_Transportation','HRDFixed_Training','HRDFixed_OtherExpense',
      'HRDFixed_SportSocial','HRDFixed_EmployeeRelation','HRDFixed_Outsource','HRDFixed_License',
      'HRDFixed_Entertainment','HRDFixed_OfficerCheck',
      'HRDVar_PrintingStationary',
      'SalesFixed_PartnershipSponsors','SalesFixed_Photography','SalesFixed_Merchandise','SalesFixed_Collaterals',
      'SalesFixed_Entertainment','SalesFixed_Telephone','SalesFixed_Miscellaneous','SalesFixed_PostageExpress',
      'SalesVar_LocalSalesCall','SalesVar_PrintingStationary',
      'POMECFixed_PestControl','POMECFixed_Telephone','POMECFixed_OtherExpense','POMECFixed_OfficerCheck',
      'POMECFixed_Uniforms','POMECFixed_LicensesAndFees',
      'POMECVar_Uniforms','POMECVar_OtherSupplies','POMECVar_PrintingStationary','POMECVar_OtherExpense',
      'POMECVar_CleaningSupplies','POMECVar_AirconVentilation','POMECVar_Building','POMECVar_Electrical',
      'POMECVar_ElectricBulbs','POMECVar_ElevEscalators','POMECVar_Furniture','POMECVar_FbKitchenRefrig',
      'POMECVar_VehicleMaintenance','POMECVar_PaintingDecoration','POMECVar_PlumbingHeating',
      'POMECVar_OfficeEquipment','POMECVar_RemovalWasteMatter','POMECVar_LocksKeys',
      'POMECVar_WaterTreatment','POMECVar_EnergyCost'
    ]]).setFontWeight('bold').setBackground('#0d1f4e').setFontColor('#ffffff');
  }

  var p = setupData.pteb    || {}, c  = setupData.cos      || {};
  var f = setupData.foFixed || {}, v  = setupData.foVar    || {};
  var hkf = setupData.hkFixed || {}, hkv = setupData.hkVar || {};
  var fbf = setupData.fbFixed || {}, fbv = setupData.fbVar || {};
  var agf = setupData.agFixed || {}, agv = setupData.agVar || {};
  var hrdf = setupData.hrdFixed || {}, hrdv = setupData.hrdVar || {};
  var sf  = setupData.salesFixed  || {}, sv  = setupData.salesVar  || {};
  var pmf = setupData.pomecFixed  || {}, pmv = setupData.pomecVar  || {};

  var newRow = [
    year, month,
    p.fo||0, p.hk||0, p.fb||0, p.ag||0, p.hrd||0, p.sales||0, p.pomec||0,
    c.breakfast||0, c.restoF||0, c.restoB||0, c.tapasF||0, c.tapasB||0, c.rsF||0, c.rsB||0, c.bqt||0,
    f.uniforms||0, f.telephone||0, f.otherExpense||0, f.travelAgencyComm||0,
    f.cableSatellite||0, f.systemsInternet||0, f.officerCheck||0, f.entertainment||0,
    v.cleaningSupplies||0, v.guestSupplies||0, v.otherSupplies||0, v.transportation||0,
    v.printingStationary||0, v.guestTransportation||0, v.compWelcomeDrink||0, v.compFruitBasket||0,
    hkf.pestControl||0, hkf.telephone||0, hkf.otherExpense||0, hkf.contractHygiene||0,
    hkf.entertainment||0, hkf.officerCheck||0, hkf.uniforms||0, hkf.licensesAndFees||0, hkf.equipmentRental||0,
    hkv.linen||0, hkv.cleaningSupplies||0, hkv.guestSupplies||0, hkv.otherSupplies||0,
    hkv.otherExpense||0, hkv.laundryDryCleaning||0,
    fbf.menus||0, fbf.licenses||0, fbf.otherExpense||0, fbf.banquetExpense||0, fbf.contractService||0,
    fbf.entertainment||0, fbf.officerCheck||0, fbf.uniforms||0, fbf.training||0,
    fbf.equipmentRental||0, fbf.musicEntertainment||0,
    fbv.cleaningSupplies||0, fbv.guestSupplies||0, fbv.kitchenFuel||0, fbv.kitchenSupplies||0,
    fbv.printingStationary||0, fbv.musicEntertainment||0, fbv.utensil||0,
    agf.uniforms||0, agf.telephone||0, agf.licensesAndFees||0, agf.itSoftware||0,
    agf.itHardware||0, agf.cashierOverShort||0, agf.bankCharges||0, agf.postageExpress||0,
    agv.transportation||0, agv.printingStationary||0, agv.creditCardCommissions||0,
    hrdf.telephone||0, hrdf.transportation||0, hrdf.training||0, hrdf.otherExpense||0,
    hrdf.sportSocial||0, hrdf.employeeRelation||0, hrdf.outsource||0, hrdf.license||0,
    hrdf.entertainment||0, hrdf.officerCheck||0,
    hrdv.printingStationary||0,
    sf.partnershipSponsors||0, sf.photography||0, sf.merchandise||0, sf.collaterals||0,
    sf.entertainment||0, sf.telephone||0, sf.miscellaneous||0, sf.postageExpress||0,
    sv.localSalesCall||0, sv.printingStationary||0,
    pmf.pestControl||0, pmf.telephone||0, pmf.otherExpense||0, pmf.officerCheck||0,
    pmf.uniforms||0, pmf.licensesAndFees||0,
    pmv.uniforms||0, pmv.otherSupplies||0, pmv.printingStationary||0, pmv.otherExpense||0,
    pmv.cleaningSupplies||0, pmv.airconVentilation||0, pmv.building||0, pmv.electrical||0,
    pmv.electricBulbs||0, pmv.elevEscalators||0, pmv.furniture||0, pmv.fbKitchenRefrig||0,
    pmv.vehicleMaintenance||0, pmv.paintingDecoration||0, pmv.plumbingHeating||0,
    pmv.officeEquipment||0, pmv.removalWasteMatter||0, pmv.locksKeys||0,
    pmv.waterTreatment||0, pmv.energyCost||0
  ];

  var rows = sheet.getDataRange().getValues();
  for (var i = 1; i < rows.length; i++) {
    if (+rows[i][0] === year && +rows[i][1] === month) {
      sheet.getRange(i + 1, 1, 1, 124).setValues([newRow]);
      return true;
    }
  }
  sheet.appendRow(newRow);
  return true;
}

function exportFlashPL(month, year, plData) {
  var monthNames = ['','Januari','Februari','Maret','April','Mei','Juni',
                    'Juli','Agustus','September','Oktober','November','Desember'];

  var ss     = SpreadsheetApp.getActiveSpreadsheet();
  var tempSS = SpreadsheetApp.create('Flash P&L - ' + monthNames[month] + ' ' + year + ' - ' + ss.getName());
  var sheet  = tempSS.getSheets()[0];
  sheet.setName('Flash P&L');
  sheet.setColumnWidth(1, 340);
  sheet.setColumnWidth(2, 180);
  sheet.setColumnWidth(3, 80);

  var rowNum = 1;
  var navy = '#0d1f4e', blueH = '#1b4fd8', blueL = '#e8effe', grayBg = '#f1f5f9';

  function wr(values, bg, bold, fc) {
    var rng = sheet.getRange(rowNum, 1, 1, values.length);
    rng.setValues([values]);
    if (bg)   rng.setBackground(bg);
    if (bold) rng.setFontWeight('bold');
    if (fc)   rng.setFontColor(fc);
    if (typeof values[1] === 'number') sheet.getRange(rowNum, 2).setNumberFormat('#,##0');
    if (typeof values[2] === 'number') sheet.getRange(rowNum, 3).setNumberFormat('0.0%');
    rowNum++;
  }

  function sec(label) {
    sheet.getRange(rowNum, 1, 1, 3).merge().setValue(label)
      .setBackground(grayBg).setFontWeight('bold').setFontColor('#475569');
    rowNum++;
  }

  function row(label, val, tot) {
    wr(['  ' + label, val, tot > 0 ? val / tot : 0], null, false, null);
  }

  function subtotal(label, val, tot) {
    wr([label, val, tot > 0 ? val / tot : 0], blueL, true, '#1e3a6e');
  }

  function deptRow(label, total, tot) {
    wr(['  ' + label, total, tot > 0 ? total / tot : 0], '#fafafa', true, '#334155');
  }

  function subSectionLabel(label) {
    sheet.getRange(rowNum, 1, 1, 3).merge().setValue('      ' + label)
      .setFontSize(9).setFontColor('#94a3b8');
    rowNum++;
  }

  function subRow(label, val, tot) {
    wr(['        ' + label, val, tot > 0 ? val / tot : 0], '#fafafa', false, '#64748b');
    sheet.getRange(rowNum - 1, 1).setFontColor('#64748b').setFontSize(11);
    sheet.getRange(rowNum - 1, 2).setFontColor('#64748b').setFontSize(11);
    sheet.getRange(rowNum - 1, 3).setFontColor('#64748b').setFontSize(9);
  }

  var rev = plData.rev;
  var tot = plData.totalRevenue;
  var fbRev = plData.fbRevenue || 0;
  var ts  = Utilities.formatDate(new Date(), Session.getScriptTimeZone(), 'dd MMM yyyy HH:mm');

  sheet.getRange(rowNum, 1, 1, 3).merge()
    .setValue('Flash P&L Forecast — ' + monthNames[month] + ' ' + year)
    .setFontSize(14).setFontWeight('bold').setBackground(navy).setFontColor('#ffffff');
  rowNum++;
  sheet.getRange(rowNum, 1, 1, 3).merge()
    .setValue(ss.getName() + '  |  Dibuat: ' + ts)
    .setFontSize(9).setFontColor('#64748b');
  rowNum += 2;

  wr(['Keterangan', 'Nominal (Rp)', '% Rev'], blueH, true, '#ffffff');

  // A. Revenue
  sec('A. Revenue');
  var revLabels = ['Rooms','Breakfast','Resto Food','Resto Beverage','Tapas Food','Tapas Beverage',
                   'RS Food','RS Beverage','BQT + Wedding','Spa','Laundry','Other Income'];
  var revKeys   = ['room','breakfast','restoF','restoB','tapasF','tapasB','rsF','rsB','bqt','spa','laundry','otherIncome'];
  for (var ri = 0; ri < revKeys.length; ri++) row(revLabels[ri], rev[revKeys[ri]]||0, tot);
  subtotal('TOTAL REVENUE', tot, tot);
  rowNum++;

  // B. PTEB
  var ptebLabels = ['Front Office','Housekeeping','Food & Beverage','Accounting / GA','HRD','Sales & Marketing','POMEC'];
  var ptebKeys   = ['fo','hk','fb','ag','hrd','sales','pomec'];
  sec('B. Payroll & Related Expenses (PTEB)');
  for (var pi = 0; pi < ptebKeys.length; pi++) row(ptebLabels[pi], plData.pteb[ptebKeys[pi]]||0, tot);
  subtotal('TOTAL PTEB', plData.totalPTEB, tot);
  rowNum++;

  // C. COS
  var cosLabels = ['Breakfast','Resto Food','Resto Bev','Tapas Food','Tapas Bev','RS Food','RS Bev','BQT+Wedding'];
  var cosKeys   = ['breakfast','restoF','restoB','tapasF','tapasB','rsF','rsB','bqt'];
  sec('C. Cost of Sales (COS)');
  for (var ci = 0; ci < cosKeys.length; ci++) row(cosLabels[ci], plData.cos[cosKeys[ci]]||0, tot);
  subtotal('TOTAL COS', plData.totalCOS, tot);
  rowNum++;

  // D. Other Departmental Expenses
  sec('D. Other Departmental Expenses');

  // Helper to render a dept breakdown
  function renderDeptBreakdown(deptLabel, fixedObj, fixedLabels, fixedKeys, varObj, varLabels, varKeys, varBase, varBaseLabel, deptTotal) {
    deptRow(deptLabel, deptTotal, tot);
    subSectionLabel('Fixed Cost');
    for (var i = 0; i < fixedKeys.length; i++) {
      var v = fixedObj[fixedKeys[i]] || 0;
      if (v > 0) subRow(fixedLabels[i], v, tot);
    }
    subSectionLabel('Variable Cost (' + varBaseLabel + ')');
    for (var j = 0; j < varKeys.length; j++) {
      var pct = varObj[varKeys[j]] || 0;
      var vv  = varBase * pct / 100;
      if (pct > 0) subRow(varLabels[j] + ' (' + pct + '%)', vv, tot);
    }
  }

  // FO
  renderDeptBreakdown(
    'Front Office',
    plData.foFixed,
    ['Uniforms','Telephone & Fax','Other Expense','Travel Agency Comm.','Cable & TV Satellite','Systems / Internet','Officer Check','Entertainment'],
    ['uniforms','telephone','otherExpense','travelAgencyComm','cableSatellite','systemsInternet','officerCheck','entertainment'],
    plData.foVar,
    ['Cleaning Supplies','Guest Supplies','Other Supplies','Transportation','Printing & Stationary','Guest Transportation','Comp. Welcome Drink','Comp. Fruit Basket'],
    ['cleaningSupplies','guestSupplies','otherSupplies','transportation','printingStationary','guestTransportation','compWelcomeDrink','compFruitBasket'],
    rev.room || 0, '% Room Rev',
    plData.foOtherExp || 0
  );

  // HK
  renderDeptBreakdown(
    'Housekeeping',
    plData.hkFixed,
    ['Pest Control','Telephone & Fax','Other Expense','Contract Hygiene','Entertainment','Officer Check','Uniforms','Licenses & Fees','Equipment Rental'],
    ['pestControl','telephone','otherExpense','contractHygiene','entertainment','officerCheck','uniforms','licensesAndFees','equipmentRental'],
    plData.hkVar,
    ['Linen','Cleaning Supplies','Guest Supplies','Other Supplies','Other Expense','Laundry & Dry Cleaning'],
    ['linen','cleaningSupplies','guestSupplies','otherSupplies','otherExpense','laundryDryCleaning'],
    rev.room || 0, '% Room Rev',
    plData.hkOtherExp || 0
  );

  // F&B
  renderDeptBreakdown(
    'Food & Beverage',
    plData.fbFixed,
    ['Menus','Licenses','Other Expense','Banquet Expense','Contract Service','Entertainment','Officer Check','Uniforms','Training','Equipment Rental','Music Entertainment'],
    ['menus','licenses','otherExpense','banquetExpense','contractService','entertainment','officerCheck','uniforms','training','equipmentRental','musicEntertainment'],
    plData.fbVar,
    ['Cleaning Supplies','Guest Supplies','Kitchen Fuel','Kitchen Supplies','Printing & Stationary','Music Entertainment','Utensil'],
    ['cleaningSupplies','guestSupplies','kitchenFuel','kitchenSupplies','printingStationary','musicEntertainment','utensil'],
    fbRev, '% F&B Rev',
    plData.fbOtherExp || 0
  );

  // A&G
  renderDeptBreakdown(
    'Accounting / GA',
    plData.agFixed,
    ['Uniforms','Telephone & Fax','Licenses & Fees','IT Software','IT Hardware','Cashier Over/Short','Bank Charges','Postage & Express'],
    ['uniforms','telephone','licensesAndFees','itSoftware','itHardware','cashierOverShort','bankCharges','postageExpress'],
    plData.agVar,
    ['Transportation','Printing & Stationary','Credit Card Commissions'],
    ['transportation','printingStationary','creditCardCommissions'],
    tot, '% Total Rev',
    plData.agOtherExp || 0
  );

  // HRD
  renderDeptBreakdown(
    'HRD',
    plData.hrdFixed,
    ['Telephone & Fax','Transportation','Training','Other Expense','Sport & Social','Employee Relation','Outsource','License','Entertainment','Officer Check'],
    ['telephone','transportation','training','otherExpense','sportSocial','employeeRelation','outsource','license','entertainment','officerCheck'],
    plData.hrdVar,
    ['Printing & Stationary'],
    ['printingStationary'],
    tot, '% Total Rev',
    plData.hrdOtherExp || 0
  );

  // Sales
  renderDeptBreakdown(
    'Sales & Marketing',
    plData.salesFixed,
    ['Partnership & Sponsors','Photography','Merchandise','Collaterals','Entertainment','Telephone & Fax','Miscellaneous','Postage & Express'],
    ['partnershipSponsors','photography','merchandise','collaterals','entertainment','telephone','miscellaneous','postageExpress'],
    plData.salesVar,
    ['Local Sales Call','Printing & Stationary'],
    ['localSalesCall','printingStationary'],
    tot, '% Total Rev',
    plData.salesOtherExp || 0
  );

  // POMEC
  renderDeptBreakdown(
    'POMEC',
    plData.pomecFixed,
    ['Pest Control','Telephone & Fax','Other Expense','Officer Check','Uniforms','Licenses & Fees'],
    ['pestControl','telephone','otherExpense','officerCheck','uniforms','licensesAndFees'],
    plData.pomecVar,
    ['Uniforms','Other Supplies','Printing & Stationary','Other Expense','Cleaning Supplies',
     'Aircon & Ventilation','Building','Electrical','Electric Bulbs','Elev. & Escalators',
     'Furniture','F&B Kitchen & Refrig.','Vehicle Maintenance','Painting & Decoration',
     'Plumbing & Heating','Office Equipment','Removal & Waste Matter','Locks & Keys',
     'Water Treatment','Energy Cost'],
    ['uniforms','otherSupplies','printingStationary','otherExpense','cleaningSupplies',
     'airconVentilation','building','electrical','electricBulbs','elevEscalators',
     'furniture','fbKitchenRefrig','vehicleMaintenance','paintingDecoration',
     'plumbingHeating','officeEquipment','removalWasteMatter','locksKeys',
     'waterTreatment','energyCost'],
    tot, '% Total Rev',
    plData.pomecOtherExp || 0
  );

  subtotal('TOTAL OTHER EXPENSES', plData.totalOtherExp, tot);
  rowNum++;

  // E. Management Fee
  sec('E. Management Fee');
  row('Management Fee (1,11% × Room Revenue)', plData.managementFee, tot);
  rowNum++;

  // GOP
  wr(['GROSS OPERATING PROFIT (GOP)', plData.gop, tot > 0 ? plData.gop / tot : 0], navy, true, '#ffffff');
  sheet.getRange(rowNum - 1, 2).setFontColor('#4ade80');
  sheet.getRange(rowNum - 1, 3).setFontColor('#4ade80').setNumberFormat('0.0%');

  sheet.autoResizeColumn(1);
  return 'https://docs.google.com/spreadsheets/d/' + tempSS.getId() + '/export?format=xlsx';
}
