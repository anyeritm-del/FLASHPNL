// Package export ports exportFlashPL() from Code.gs: it builds a formatted
// Flash P&L report as a new Google Sheet and shares it with the requesting
// user so they can open the .xlsx export link (the service account, not the
// user, owns the sheet, so this explicit share step replaces what used to be
// implicit under Session.getActiveUser()).
package export

import (
	"strconv"
	"strings"

	"google.golang.org/api/sheets/v4"
)

const (
	colorNavy   = "#0d1f4e"
	colorBlueH  = "#1b4fd8"
	colorBlueL  = "#e8effe"
	colorGrayBg = "#f1f5f9"
)

type styleOp struct {
	row, colStart, colEnd int
	bg, fontColor         *sheets.Color
	bold                  bool
	fontSize              int64
	numberFormat          string
}

type mergeOp struct {
	row, colStart, colEnd int
}

type item struct {
	Label string
	Val   float64
}

type builder struct {
	values [][]interface{}
	styles []styleOp
	merges []mergeOp
}

func hexColor(hex string) *sheets.Color {
	hex = strings.TrimPrefix(hex, "#")
	var r, g, b int64
	if len(hex) == 6 {
		r, _ = strconv.ParseInt(hex[0:2], 16, 64)
		g, _ = strconv.ParseInt(hex[2:4], 16, 64)
		b, _ = strconv.ParseInt(hex[4:6], 16, 64)
	}
	return &sheets.Color{Red: float64(r) / 255, Green: float64(g) / 255, Blue: float64(b) / 255}
}

func (b *builder) addValues(vals []interface{}) int {
	row := make([]interface{}, 3)
	copy(row, vals)
	b.values = append(b.values, row)
	return len(b.values) - 1
}

func (b *builder) style(row, colStart, colEnd int, bg, fontColor string, bold bool, fontSize int) {
	s := styleOp{row: row, colStart: colStart, colEnd: colEnd, bold: bold}
	if bg != "" {
		s.bg = hexColor(bg)
	}
	if fontColor != "" {
		s.fontColor = hexColor(fontColor)
	}
	if fontSize > 0 {
		s.fontSize = int64(fontSize)
	}
	b.styles = append(b.styles, s)
}

func (b *builder) numberFormat(row, col int, pattern string) {
	b.styles = append(b.styles, styleOp{row: row, colStart: col, colEnd: col + 1, numberFormat: pattern})
}

func (b *builder) merge(row, numCols int) {
	b.merges = append(b.merges, mergeOp{row: row, colStart: 0, colEnd: numCols})
}

func (b *builder) blank() {
	b.addValues([]interface{}{nil, nil, nil})
}

// wr mirrors the JS `wr()` closure: write a 3-column row and format it.
func (b *builder) wr(vals []interface{}, bg string, bold bool, fontColor string) int {
	idx := b.addValues(vals)
	if bg != "" || bold || fontColor != "" {
		b.style(idx, 0, 3, bg, fontColor, bold, 0)
	}
	if len(vals) > 1 {
		if _, ok := vals[1].(float64); ok {
			b.numberFormat(idx, 1, "#,##0")
		}
	}
	if len(vals) > 2 {
		if _, ok := vals[2].(float64); ok {
			b.numberFormat(idx, 2, "0.0%")
		}
	}
	return idx
}

func pctOf(val, tot float64) interface{} {
	if tot > 0 {
		return val / tot
	}
	return float64(0)
}

func (b *builder) sec(label string) {
	idx := b.addValues([]interface{}{label, nil, nil})
	b.style(idx, 0, 3, colorGrayBg, "#475569", true, 0)
	b.merge(idx, 3)
}

func (b *builder) row(label string, val, tot float64) {
	b.wr([]interface{}{"  " + label, val, pctOf(val, tot)}, "", false, "")
}

func (b *builder) subtotal(label string, val, tot float64) {
	b.wr([]interface{}{label, val, pctOf(val, tot)}, colorBlueL, true, "#1e3a6e")
}

func (b *builder) deptRow(label string, total, tot float64) {
	b.wr([]interface{}{"  " + label, total, pctOf(total, tot)}, "#fafafa", true, "#334155")
}

func (b *builder) subSectionLabel(label string) {
	idx := b.addValues([]interface{}{"      " + label, nil, nil})
	b.style(idx, 0, 3, "", "#94a3b8", false, 9)
	b.merge(idx, 3)
}

func (b *builder) subRow(label string, val, tot float64) {
	idx := b.wr([]interface{}{"        " + label, val, pctOf(val, tot)}, "#fafafa", false, "#64748b")
	b.style(idx, 0, 2, "", "#64748b", false, 11)
	b.style(idx, 2, 3, "", "#64748b", false, 9)
}

func numLabel(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func renderDeptBreakdown(b *builder, deptLabel string, fixed, vars []item, varBase float64, varBaseLabel string, deptTotal, tot float64) {
	b.deptRow(deptLabel, deptTotal, tot)
	b.subSectionLabel("Fixed Cost")
	for _, it := range fixed {
		if it.Val > 0 {
			b.subRow(it.Label, it.Val, tot)
		}
	}
	b.subSectionLabel("Variable Cost (" + varBaseLabel + ")")
	for _, it := range vars {
		if it.Val > 0 {
			vv := varBase * it.Val / 100
			b.subRow(it.Label+" ("+numLabel(it.Val)+"%)", vv, tot)
		}
	}
}
