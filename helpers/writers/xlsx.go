package writers

import (
	"bytes"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

type IXlsxHelper interface {
	CreateFile() ([]byte, error)
}

type XlsxHelper struct {
	file        *excelize.File           `json:"file,omitempty"`
	keys        []string                 `json:"keys,omitempty"`
	data        []map[string]interface{} `json:"data,omitempty"`
	HeaderCells []string                 `json:"header_cells,omitempty"`

	Data      []string `bun:"-" json:"data,omitempty"`
	Err       error    `json:"err,omitempty"`
	Cursor    int      `json:"cursor,omitempty"`
	StrCursor int      `json:"str_cursor,omitempty"`
}

func (x *XlsxHelper) CreateFile() {
	x.file = excelize.NewFile()
	//x.initXlsxData(keys, data)
	//Err := x.WriteHeader()
	//if Err != nil {
	// return nil, Err
	//}
	//Err := x.writeData()
	//if Err != nil {
	// return nil, Err
	//}
	//return x.WriteFile()
}

func (x *XlsxHelper) initXlsxData(keys []string, data []map[string]interface{}) {
	x.keys = keys
	x.data = data
}

func NewXlsxHelper() *XlsxHelper {
	return &XlsxHelper{}
}

func (x *XlsxHelper) stylingHeader() error {
	style, err := x.file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#D0D0D0"}, Pattern: 1},
	})
	if err != nil {
		return err
	}
	err = x.file.SetColWidth("Sheet1", "A", "BA", 30)
	if err != nil {
		return err
	}
	err = x.file.SetCellStyle("Sheet1", "A1", "BA1", style)
	if err != nil {
		return err
	}
	return err
}

// func (x *XlsxHelper)
func (x *XlsxHelper) SetError(err error) {
	if x.Err != nil {
		return
	}
	x.Err = err
}

func (x *XlsxHelper) WriteString(strNum int, colNum int, data interface{}) {
	col := x.GetCol(colNum)
	startCell := fmt.Sprintf("%s%d", col, strNum)
	x.SetError(x.file.SetSheetRow("Sheet1", startCell, data))
}

func (x *XlsxHelper) WriteCell(strNum int, colNum int, data interface{}) {
	col := x.GetCol(colNum)
	cell := fmt.Sprintf("%s%d", col, strNum)
	var err error
	switch d := data.(type) {
	case nil:
		err = x.file.SetCellStr("Sheet1", cell, "")
	case string:
		err = x.file.SetCellStr("Sheet1", cell, d)
	case int, uint:
		err = x.file.SetCellInt("Sheet1", cell, d.(int))
	case float64:
		err = x.file.SetCellFloat("Sheet1", cell, d, 2, 32)
	case float32:
		err = x.file.SetCellFloat("Sheet1", cell, float64(d), 2, 64)
	case *time.Time:
		// err = x.file.SetCellStr("Sheet1", cell, d.Format("02.01.2006"))
		style, err := x.file.NewStyle(&excelize.Style{NumFmt: 22})
		if err != nil {
			fmt.Println(err)
		}
		err = x.file.SetCellStyle("Sheet1", cell, cell, style)
		if err != nil {
			fmt.Println(err)
		}

		// err = x.file.SetCellValue("Sheet1", cell, d.Format("02.01.2006"))
		err = x.file.SetCellValue("Sheet1", cell, *d)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(d)
	case DataWithStyle:
		style, err := x.file.NewStyle(&excelize.Style{NumFmt: 40})
		if err != nil {
			fmt.Println(err)
		}
		err = x.file.SetCellStyle("Sheet1", cell, cell, style)
		if err != nil {
			fmt.Println(err)
		}

		// err = x.file.SetCellValue("Sheet1", cell, d.Format("02.01.2006"))
		err = x.file.SetCellValue("Sheet1", cell, d)
		if err != nil {
			fmt.Println(err)
		}
		// err = x.file.SetCellValue("Sheet1", cell, GetNormalizedData(d.Data).(string))
		// err = x.file.SetCellValue("Sheet1", cell, d.Data)
		// if err != nil {
		// 	break
		// }
		color := d.Style.Color
		x.SetCellColor(cell, color)
	}

	x.SetError(err)
}

func (x *XlsxHelper) MergeArea(start string, end string) {
	x.SetError(x.file.MergeCell("Sheet1", start, end))
}

var cols = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ", "CA", "CB", "CC", "CD", "CE", "CF", "CG", "CH", "CI", "CJ", "CK", "CL", "CM", "CN", "CO", "CP", "CQ", "CR", "CS", "CT", "CU", "CV", "CW", "CX", "CY", "CZ", "DA", "DB", "DC", "DD", "DE", "DF", "DG", "DH", "DI", "DJ", "DK", "DL", "DM", "DN", "DO", "DP", "DQ", "DR", "DS", "DT", "DU", "DV", "DW", "DX", "DY", "DZ", "EA", "EB", "EC", "ED", "EE", "EF", "EG", "EH", "EI", "EJ", "EK", "EL", "EM", "EN", "EO", "EP", "EQ", "ER", "ES", "ET", "EU", "EV", "EW", "EX", "EY", "EZ", "FA", "FB", "FC", "FD", "FE", "FF", "FG", "FH", "FI", "FJ", "FK", "FL", "FM", "FN", "FO", "FP", "FQ", "FR", "FS", "FT", "FU", "FV", "FW", "FX", "FY", "FZ", "GA", "GB", "GC", "GD", "GE", "GF", "GG", "GH", "GI", "GJ", "GK", "GL", "GM", "GN", "GO", "GP", "GQ", "GR", "GS", "GT", "GU", "GV", "GW", "GX", "GY", "GZ", "HA", "HB", "HC", "HD", "HE", "HF", "HG", "HH", "HI", "HJ", "HK", "HL", "HM", "HN", "HO", "HP", "HQ", "HR", "HS", "HT", "HU", "HV", "HW", "HX", "HY", "HZ", "IA", "IB", "IC", "ID", "IE", "IF", "IG", "IH", "II", "IJ", "IK", "IL", "IM", "IN", "IO", "IP", "IQ", "IR", "IS", "IT", "IU", "IV", "IW", "IX", "IY", "IZ", "JA", "JB", "JC", "JD", "JE", "JF", "JG", "JH", "JI", "JJ", "JK", "JL", "JM", "JN", "JO", "JP", "JQ", "JR", "JS", "JT", "JU", "JV", "JW", "JX", "JY", "JZ", "KA", "KB", "KC", "KD", "KE", "KF", "KG", "KH", "KI", "KJ", "KK", "KL", "KM", "KN", "KO", "KP", "KQ", "KR", "KS", "KT", "KU", "KV", "KW", "KX", "KY", "KZ", "LA", "LB", "LC", "LD", "LE", "LF", "LG", "LH", "LI", "LJ", "LK", "LL", "LM", "LN", "LO", "LP", "LQ", "LR", "LS", "LT", "LU", "LV", "LW", "LX", "LY", "LZ", "MA", "MB", "MC", "MD", "ME", "MF", "MG", "MH", "MI", "MJ", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NB", "NC", "ND", "NE", "NF", "NG", "NH", "NI", "NJ", "NK", "NL", "NM", "NN", "NO", "NP", "NQ", "NR", "NS", "NT", "NU", "NV", "NW", "NX", "NY", "NZ", "OA", "OB", "OC", "OD", "OE", "OF", "OG", "OH", "OI", "OJ", "OK", "OL", "OM", "ON", "OO", "OP", "OQ", "OR", "OS", "OT", "OU", "OV", "OW", "OX", "OY", "OZ", "PA", "PB", "PC", "PD", "PE", "PF", "PG", "PH", "PI", "PJ", "PK", "PL", "PM", "PN", "PO", "PP", "PQ", "PR", "PS", "PT", "PU", "PV", "PW", "PX", "PY", "PZ", "QA", "QB", "QC", "QD", "QE", "QF", "QG", "QH", "QI", "QJ", "QK", "QL", "QM", "QN", "QO", "QP", "QQ", "QR", "QS", "QT", "QU", "QV", "QW", "QX", "QY", "QZ", "RA", "RB", "RC", "RD", "RE", "RF", "RG", "RH", "RI", "RJ", "RK", "RL", "RM", "RN", "RO", "RP", "RQ", "RR", "RS", "RT", "RU", "RV", "RW", "RX", "RY", "RZ", "SA", "SB", "SC", "SD", "SE", "SF", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SP", "SQ", "SR", "SS", "ST", "SU", "SV", "SW", "SX", "SY", "SZ", "TA", "TB", "TC", "TD", "TE", "TF", "TG", "TH", "TI", "TJ", "TK", "TL", "TM", "TN", "TO", "TP", "TQ", "TR", "TS", "TT", "TU", "TV", "TW", "TX", "TY", "TZ", "UA", "UB", "UC", "UD", "UE", "UF", "UG", "UH", "UI", "UJ", "UK", "UL", "UM", "UN", "UO", "UP", "UQ", "UR", "US", "UT", "UU", "UV", "UW", "UX", "UY", "UZ", "VA", "VB", "VC", "VD", "VE", "VF", "VG", "VH", "VI", "VJ", "VK", "VL", "VM", "VN", "VO", "VP", "VQ", "VR", "VS", "VT", "VU", "VV", "VW", "VX", "VY", "VZ", "WA", "WB", "WC", "WD", "WE", "WF", "WG", "WH", "WI", "WJ", "WK", "WL", "WM", "WN", "WO", "WP", "WQ", "WR", "WS", "WT", "WU", "WV", "WW", "WX", "WY", "WZ", "XA", "XB", "XC", "XD", "XE", "XF", "XG", "XH", "XI", "XJ", "XK", "XL", "XM", "XN", "XO", "XP", "XQ", "XR", "XS", "XT", "XU", "XV", "XW", "XX", "XY", "XZ", "YA", "YB", "YC", "YD", "YE", "YF", "YG", "YH", "YI", "YJ", "YK", "YL", "YM", "YN", "YO", "YP", "YQ", "YR", "YS", "YT", "YU", "YV", "YW", "YX", "YY", "YZ", "ZA", "ZB", "ZC", "ZD", "ZE", "ZF", "ZG", "ZH", "ZI", "ZJ", "ZK", "ZL", "ZM", "ZN", "ZO", "ZP", "ZQ", "ZR", "ZS", "ZT", "ZU", "ZV", "ZW", "ZX", "ZY", "ZZ"}

func (x *XlsxHelper) GetColsArea(i int) string {
	return cols[i]
}

func (x *XlsxHelper) GetCol(i int) string {
	return cols[i]
}

func (x *XlsxHelper) WriteHeader() error {
	err := x.file.SetSheetRow("Sheet1", "A1", &x.keys)
	if err != nil {
		return err
	}
	err = x.stylingHeader()
	if err != nil {
		return err
	}
	return err
}

func (x *XlsxHelper) writeData() error {
	if len(x.data) == 0 {
		return nil
	}
	for i, mapa := range x.data {
		values := make([]interface{}, 0)
		for _, k := range x.keys {
			values = append(values, mapa[k])
		}
		row := fmt.Sprintf("A%d", i+2)
		// fmt.Println(row)
		err := x.file.SetSheetRow("Sheet1", row, &values)
		if err != nil {
			return err
		}
	}
	return nil
}

// Write func
func (x *XlsxHelper) WriteFile() ([]byte, error) {
	var b bytes.Buffer
	err := x.file.Write(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (x *XlsxHelper) AutofitAllColumns() {
	cols, _ := x.file.GetCols("Sheet1")
	//if err != nil {
	//	return x
	//}
	for _, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		//name, _ := excelize.ColumnNumberToName(idx + 1)
		//if err != nil {
		//	return err
		//}
		//x.file.SetColWidth("Sheet1", name, name, float64(largestWidth))
	}
}

func (x *XlsxHelper) SetBorder(cell string) {
	style, err := x.file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
	})
	if err != nil {
		x.Err = err
	}
	err = x.file.SetCellStyle("Sheet1", cell, cell, style)
	if err != nil {
		x.Err = err
	}
}

func (x *XlsxHelper) SetCellColor(cell string, color string) {
	style, err := x.file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Color: color,
		},
	})
	if err != nil {
		x.Err = err
	}
	err = x.file.SetCellStyle("Sheet1", cell, cell, style)
	if err != nil {
		x.Err = err
	}
}
