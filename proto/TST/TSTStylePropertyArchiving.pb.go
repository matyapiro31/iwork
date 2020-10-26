// Code generated by protoc-gen-go.
// source: TSTStylePropertyArchiving.proto
// DO NOT EDIT!

/*
Package TST is a generated protocol buffer package.

It is generated from these files:
	TSTStylePropertyArchiving.proto

It has these top-level messages:
	Deprecated_TableStrokeArchive
	CellStylePropertiesArchive
	Deprecated_StrokePresetDataArchive
	StrokePresetDataArchive
	StrokePresetListArchive
	TableStylePropertiesArchive
	TableStylePresetArchive
	TableStrokePresetArchive
	ThemePresetsArchive
*/
package TST

import proto "github.com/golang/protobuf/proto"
import math "math"
import "github.com/dunhamsteve/iwork/proto/TSP"
import "github.com/dunhamsteve/iwork/proto/TSD"

// discarding unused import TSK "TSKArchives.pb"
import "github.com/dunhamsteve/iwork/proto/TSS"
import "github.com/dunhamsteve/iwork/proto/TSWP"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Deprecated_TableStrokeArchive struct {
	Stroke           *TSD.StrokeArchive `protobuf:"bytes,1,opt,name=stroke" json:"stroke,omitempty"`
	Background       *bool              `protobuf:"varint,2,opt,name=background" json:"background,omitempty"`
	Opacity          *float32           `protobuf:"fixed32,3,opt,name=opacity" json:"opacity,omitempty"`
	Empty            *bool              `protobuf:"varint,4,opt,name=empty" json:"empty,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Deprecated_TableStrokeArchive) Reset()         { *m = Deprecated_TableStrokeArchive{} }
func (m *Deprecated_TableStrokeArchive) String() string { return proto.CompactTextString(m) }
func (*Deprecated_TableStrokeArchive) ProtoMessage()    {}

func (m *Deprecated_TableStrokeArchive) GetStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.Stroke
	}
	return nil
}

func (m *Deprecated_TableStrokeArchive) GetBackground() bool {
	if m != nil && m.Background != nil {
		return *m.Background
	}
	return false
}

func (m *Deprecated_TableStrokeArchive) GetOpacity() float32 {
	if m != nil && m.Opacity != nil {
		return *m.Opacity
	}
	return 0
}

func (m *Deprecated_TableStrokeArchive) GetEmpty() bool {
	if m != nil && m.Empty != nil {
		return *m.Empty
	}
	return false
}

type CellStylePropertiesArchive struct {
	CellFill               *TSD.FillArchive               `protobuf:"bytes,1,opt,name=cell_fill" json:"cell_fill,omitempty"`
	TextWrap               *bool                          `protobuf:"varint,3,opt,name=text_wrap" json:"text_wrap,omitempty"`
	DeprecatedTopStroke    *Deprecated_TableStrokeArchive `protobuf:"bytes,4,opt,name=deprecated_top_stroke" json:"deprecated_top_stroke,omitempty"`
	DeprecatedRightStroke  *Deprecated_TableStrokeArchive `protobuf:"bytes,5,opt,name=deprecated_right_stroke" json:"deprecated_right_stroke,omitempty"`
	DeprecatedBottomStroke *Deprecated_TableStrokeArchive `protobuf:"bytes,6,opt,name=deprecated_bottom_stroke" json:"deprecated_bottom_stroke,omitempty"`
	DeprecatedLeftStroke   *Deprecated_TableStrokeArchive `protobuf:"bytes,7,opt,name=deprecated_left_stroke" json:"deprecated_left_stroke,omitempty"`
	VerticalAlignment      *int32                         `protobuf:"varint,8,opt,name=vertical_alignment" json:"vertical_alignment,omitempty"`
	Padding                *TSWP.PaddingArchive           `protobuf:"bytes,9,opt,name=padding" json:"padding,omitempty"`
	TopStroke              *TSD.StrokeArchive             `protobuf:"bytes,10,opt,name=top_stroke" json:"top_stroke,omitempty"`
	RightStroke            *TSD.StrokeArchive             `protobuf:"bytes,11,opt,name=right_stroke" json:"right_stroke,omitempty"`
	BottomStroke           *TSD.StrokeArchive             `protobuf:"bytes,12,opt,name=bottom_stroke" json:"bottom_stroke,omitempty"`
	LeftStroke             *TSD.StrokeArchive             `protobuf:"bytes,13,opt,name=left_stroke" json:"left_stroke,omitempty"`
	XXX_unrecognized       []byte                         `json:"-"`
}

func (m *CellStylePropertiesArchive) Reset()         { *m = CellStylePropertiesArchive{} }
func (m *CellStylePropertiesArchive) String() string { return proto.CompactTextString(m) }
func (*CellStylePropertiesArchive) ProtoMessage()    {}

func (m *CellStylePropertiesArchive) GetCellFill() *TSD.FillArchive {
	if m != nil {
		return m.CellFill
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetTextWrap() bool {
	if m != nil && m.TextWrap != nil {
		return *m.TextWrap
	}
	return false
}

func (m *CellStylePropertiesArchive) GetDeprecatedTopStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedTopStroke
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetDeprecatedRightStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedRightStroke
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetDeprecatedBottomStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedBottomStroke
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetDeprecatedLeftStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedLeftStroke
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetVerticalAlignment() int32 {
	if m != nil && m.VerticalAlignment != nil {
		return *m.VerticalAlignment
	}
	return 0
}

func (m *CellStylePropertiesArchive) GetPadding() *TSWP.PaddingArchive {
	if m != nil {
		return m.Padding
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetTopStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.TopStroke
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetRightStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.RightStroke
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetBottomStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.BottomStroke
	}
	return nil
}

func (m *CellStylePropertiesArchive) GetLeftStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.LeftStroke
	}
	return nil
}

type Deprecated_StrokePresetDataArchive struct {
	DeprecatedHorizontalStroke *Deprecated_TableStrokeArchive `protobuf:"bytes,2,req,name=deprecated_horizontal_stroke" json:"deprecated_horizontal_stroke,omitempty"`
	DeprecatedVerticalStroke   *Deprecated_TableStrokeArchive `protobuf:"bytes,1,req,name=deprecated_vertical_stroke" json:"deprecated_vertical_stroke,omitempty"`
	DeprecatedExteriorStroke   *Deprecated_TableStrokeArchive `protobuf:"bytes,3,req,name=deprecated_exterior_stroke" json:"deprecated_exterior_stroke,omitempty"`
	DeprecatedVisibleMask      *int32                         `protobuf:"varint,5,req,name=deprecated_visible_mask" json:"deprecated_visible_mask,omitempty"`
	XXX_unrecognized           []byte                         `json:"-"`
}

func (m *Deprecated_StrokePresetDataArchive) Reset()         { *m = Deprecated_StrokePresetDataArchive{} }
func (m *Deprecated_StrokePresetDataArchive) String() string { return proto.CompactTextString(m) }
func (*Deprecated_StrokePresetDataArchive) ProtoMessage()    {}

func (m *Deprecated_StrokePresetDataArchive) GetDeprecatedHorizontalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHorizontalStroke
	}
	return nil
}

func (m *Deprecated_StrokePresetDataArchive) GetDeprecatedVerticalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedVerticalStroke
	}
	return nil
}

func (m *Deprecated_StrokePresetDataArchive) GetDeprecatedExteriorStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedExteriorStroke
	}
	return nil
}

func (m *Deprecated_StrokePresetDataArchive) GetDeprecatedVisibleMask() int32 {
	if m != nil && m.DeprecatedVisibleMask != nil {
		return *m.DeprecatedVisibleMask
	}
	return 0
}

type StrokePresetDataArchive struct {
	HorizontalStroke *TSD.StrokeArchive `protobuf:"bytes,1,opt,name=horizontal_stroke" json:"horizontal_stroke,omitempty"`
	VerticalStroke   *TSD.StrokeArchive `protobuf:"bytes,2,opt,name=vertical_stroke" json:"vertical_stroke,omitempty"`
	ExteriorStroke   *TSD.StrokeArchive `protobuf:"bytes,3,opt,name=exterior_stroke" json:"exterior_stroke,omitempty"`
	VisibleMask      *int32             `protobuf:"varint,4,opt,name=visible_mask" json:"visible_mask,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *StrokePresetDataArchive) Reset()         { *m = StrokePresetDataArchive{} }
func (m *StrokePresetDataArchive) String() string { return proto.CompactTextString(m) }
func (*StrokePresetDataArchive) ProtoMessage()    {}

func (m *StrokePresetDataArchive) GetHorizontalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HorizontalStroke
	}
	return nil
}

func (m *StrokePresetDataArchive) GetVerticalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.VerticalStroke
	}
	return nil
}

func (m *StrokePresetDataArchive) GetExteriorStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.ExteriorStroke
	}
	return nil
}

func (m *StrokePresetDataArchive) GetVisibleMask() int32 {
	if m != nil && m.VisibleMask != nil {
		return *m.VisibleMask
	}
	return 0
}

type StrokePresetListArchive struct {
	Count            *int32                                `protobuf:"varint,1,req,name=count" json:"count,omitempty"`
	DeprecatedPreset []*Deprecated_StrokePresetDataArchive `protobuf:"bytes,2,rep,name=deprecated_preset" json:"deprecated_preset,omitempty"`
	Preset           []*StrokePresetDataArchive            `protobuf:"bytes,3,rep,name=preset" json:"preset,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *StrokePresetListArchive) Reset()         { *m = StrokePresetListArchive{} }
func (m *StrokePresetListArchive) String() string { return proto.CompactTextString(m) }
func (*StrokePresetListArchive) ProtoMessage()    {}

func (m *StrokePresetListArchive) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *StrokePresetListArchive) GetDeprecatedPreset() []*Deprecated_StrokePresetDataArchive {
	if m != nil {
		return m.DeprecatedPreset
	}
	return nil
}

func (m *StrokePresetListArchive) GetPreset() []*StrokePresetDataArchive {
	if m != nil {
		return m.Preset
	}
	return nil
}

type TableStylePropertiesArchive struct {
	BandedRows                                *bool                          `protobuf:"varint,1,opt,name=banded_rows" json:"banded_rows,omitempty"`
	BandedFill                                *TSD.FillArchive               `protobuf:"bytes,2,opt,name=banded_fill" json:"banded_fill,omitempty"`
	BehavesLikeSpreadsheet                    *bool                          `protobuf:"varint,21,opt,name=behaves_like_spreadsheet" json:"behaves_like_spreadsheet,omitempty"`
	AutoResize                                *bool                          `protobuf:"varint,22,opt,name=auto_resize" json:"auto_resize,omitempty"`
	DeprecatedHeaderRowSeparatorStroke        *Deprecated_TableStrokeArchive `protobuf:"bytes,4,opt,name=deprecated_header_row_separator_stroke" json:"deprecated_header_row_separator_stroke,omitempty"`
	DeprecatedHeaderRowBorderStroke           *Deprecated_TableStrokeArchive `protobuf:"bytes,5,opt,name=deprecated_header_row_border_stroke" json:"deprecated_header_row_border_stroke,omitempty"`
	DeprecatedHeaderRowHorizontalStroke       *Deprecated_TableStrokeArchive `protobuf:"bytes,23,opt,name=deprecated_header_row_horizontal_stroke" json:"deprecated_header_row_horizontal_stroke,omitempty"`
	DeprecatedHeaderRowVerticalStroke         *Deprecated_TableStrokeArchive `protobuf:"bytes,24,opt,name=deprecated_header_row_vertical_stroke" json:"deprecated_header_row_vertical_stroke,omitempty"`
	DeprecatedHeaderColumnBorderStroke        *Deprecated_TableStrokeArchive `protobuf:"bytes,7,opt,name=deprecated_header_column_border_stroke" json:"deprecated_header_column_border_stroke,omitempty"`
	DeprecatedHeaderColumnSeparatorStroke     *Deprecated_TableStrokeArchive `protobuf:"bytes,8,opt,name=deprecated_header_column_separator_stroke" json:"deprecated_header_column_separator_stroke,omitempty"`
	DeprecatedHeaderColumnHorizontalStroke    *Deprecated_TableStrokeArchive `protobuf:"bytes,25,opt,name=deprecated_header_column_horizontal_stroke" json:"deprecated_header_column_horizontal_stroke,omitempty"`
	DeprecatedHeaderColumnVerticalStroke      *Deprecated_TableStrokeArchive `protobuf:"bytes,26,opt,name=deprecated_header_column_vertical_stroke" json:"deprecated_header_column_vertical_stroke,omitempty"`
	DeprecatedFooterRowSeparatorStroke        *Deprecated_TableStrokeArchive `protobuf:"bytes,10,opt,name=deprecated_footer_row_separator_stroke" json:"deprecated_footer_row_separator_stroke,omitempty"`
	DeprecatedFooterRowBorderStroke           *Deprecated_TableStrokeArchive `protobuf:"bytes,11,opt,name=deprecated_footer_row_border_stroke" json:"deprecated_footer_row_border_stroke,omitempty"`
	DeprecatedFooterRowHorizontalStroke       *Deprecated_TableStrokeArchive `protobuf:"bytes,27,opt,name=deprecated_footer_row_horizontal_stroke" json:"deprecated_footer_row_horizontal_stroke,omitempty"`
	DeprecatedFooterRowVerticalStroke         *Deprecated_TableStrokeArchive `protobuf:"bytes,28,opt,name=deprecated_footer_row_vertical_stroke" json:"deprecated_footer_row_vertical_stroke,omitempty"`
	DeprecatedTableBodyHorizontalBorderStroke *Deprecated_TableStrokeArchive `protobuf:"bytes,12,opt,name=deprecated_table_body_horizontal_border_stroke" json:"deprecated_table_body_horizontal_border_stroke,omitempty"`
	DeprecatedTableBodyVerticalBorderStroke   *Deprecated_TableStrokeArchive `protobuf:"bytes,29,opt,name=deprecated_table_body_vertical_border_stroke" json:"deprecated_table_body_vertical_border_stroke,omitempty"`
	DeprecatedTableBodyHorizontalStroke       *Deprecated_TableStrokeArchive `protobuf:"bytes,30,opt,name=deprecated_table_body_horizontal_stroke" json:"deprecated_table_body_horizontal_stroke,omitempty"`
	DeprecatedTableBodyVerticalStroke         *Deprecated_TableStrokeArchive `protobuf:"bytes,31,opt,name=deprecated_table_body_vertical_stroke" json:"deprecated_table_body_vertical_stroke,omitempty"`
	StrokePresetList                          *StrokePresetListArchive       `protobuf:"bytes,32,opt,name=stroke_preset_list" json:"stroke_preset_list,omitempty"`
	VStrokesVisible                           *bool                          `protobuf:"varint,33,opt,name=v_strokes_visible" json:"v_strokes_visible,omitempty"`
	HStrokesVisible                           *bool                          `protobuf:"varint,34,opt,name=h_strokes_visible" json:"h_strokes_visible,omitempty"`
	HrSeparatorVisible                        *bool                          `protobuf:"varint,35,opt,name=hr_separator_visible" json:"hr_separator_visible,omitempty"`
	HcSeparatorVisible                        *bool                          `protobuf:"varint,36,opt,name=hc_separator_visible" json:"hc_separator_visible,omitempty"`
	FooterSeparatorVisible                    *bool                          `protobuf:"varint,37,opt,name=footer_separator_visible" json:"footer_separator_visible,omitempty"`
	TableBorderVisible                        *bool                          `protobuf:"varint,38,opt,name=table_border_visible" json:"table_border_visible,omitempty"`
	TableHeaderBorderVisible                  *bool                          `protobuf:"varint,39,opt,name=table_header_border_visible" json:"table_header_border_visible,omitempty"`
	TableHcDividerVisible                     *bool                          `protobuf:"varint,42,opt,name=table_hc_divider_visible" json:"table_hc_divider_visible,omitempty"`
	TableHrDividerVisible                     *bool                          `protobuf:"varint,43,opt,name=table_hr_divider_visible" json:"table_hr_divider_visible,omitempty"`
	TableFooterDividerVisible                 *bool                          `protobuf:"varint,44,opt,name=table_footer_divider_visible" json:"table_footer_divider_visible,omitempty"`
	OBSOLETEMasterFontSize                    *int32                         `protobuf:"varint,40,opt,name=OBSOLETE_master_font_size" json:"OBSOLETE_master_font_size,omitempty"`
	MasterFontFamily                          *string                        `protobuf:"bytes,41,opt,name=master_font_family" json:"master_font_family,omitempty"`
	WritingDirection                          *TSWP.WritingDirectionType     `protobuf:"varint,45,opt,name=writing_direction,enum=TSWP.WritingDirectionType" json:"writing_direction,omitempty"`
	HeaderRowSeparatorStroke                  *TSD.StrokeArchive             `protobuf:"bytes,46,opt,name=header_row_separator_stroke" json:"header_row_separator_stroke,omitempty"`
	HeaderRowBorderStroke                     *TSD.StrokeArchive             `protobuf:"bytes,47,opt,name=header_row_border_stroke" json:"header_row_border_stroke,omitempty"`
	HeaderRowHorizontalStroke                 *TSD.StrokeArchive             `protobuf:"bytes,48,opt,name=header_row_horizontal_stroke" json:"header_row_horizontal_stroke,omitempty"`
	HeaderRowVerticalStroke                   *TSD.StrokeArchive             `protobuf:"bytes,49,opt,name=header_row_vertical_stroke" json:"header_row_vertical_stroke,omitempty"`
	HeaderColumnBorderStroke                  *TSD.StrokeArchive             `protobuf:"bytes,50,opt,name=header_column_border_stroke" json:"header_column_border_stroke,omitempty"`
	HeaderColumnSeparatorStroke               *TSD.StrokeArchive             `protobuf:"bytes,51,opt,name=header_column_separator_stroke" json:"header_column_separator_stroke,omitempty"`
	HeaderColumnHorizontalStroke              *TSD.StrokeArchive             `protobuf:"bytes,52,opt,name=header_column_horizontal_stroke" json:"header_column_horizontal_stroke,omitempty"`
	HeaderColumnVerticalStroke                *TSD.StrokeArchive             `protobuf:"bytes,53,opt,name=header_column_vertical_stroke" json:"header_column_vertical_stroke,omitempty"`
	FooterRowSeparatorStroke                  *TSD.StrokeArchive             `protobuf:"bytes,54,opt,name=footer_row_separator_stroke" json:"footer_row_separator_stroke,omitempty"`
	FooterRowBorderStroke                     *TSD.StrokeArchive             `protobuf:"bytes,55,opt,name=footer_row_border_stroke" json:"footer_row_border_stroke,omitempty"`
	FooterRowHorizontalStroke                 *TSD.StrokeArchive             `protobuf:"bytes,56,opt,name=footer_row_horizontal_stroke" json:"footer_row_horizontal_stroke,omitempty"`
	FooterRowVerticalStroke                   *TSD.StrokeArchive             `protobuf:"bytes,57,opt,name=footer_row_vertical_stroke" json:"footer_row_vertical_stroke,omitempty"`
	TableBodyHorizontalBorderStroke           *TSD.StrokeArchive             `protobuf:"bytes,58,opt,name=table_body_horizontal_border_stroke" json:"table_body_horizontal_border_stroke,omitempty"`
	TableBodyVerticalBorderStroke             *TSD.StrokeArchive             `protobuf:"bytes,59,opt,name=table_body_vertical_border_stroke" json:"table_body_vertical_border_stroke,omitempty"`
	TableBodyHorizontalStroke                 *TSD.StrokeArchive             `protobuf:"bytes,60,opt,name=table_body_horizontal_stroke" json:"table_body_horizontal_stroke,omitempty"`
	TableBodyVerticalStroke                   *TSD.StrokeArchive             `protobuf:"bytes,61,opt,name=table_body_vertical_stroke" json:"table_body_vertical_stroke,omitempty"`
	XXX_unrecognized                          []byte                         `json:"-"`
}

func (m *TableStylePropertiesArchive) Reset()         { *m = TableStylePropertiesArchive{} }
func (m *TableStylePropertiesArchive) String() string { return proto.CompactTextString(m) }
func (*TableStylePropertiesArchive) ProtoMessage()    {}

func (m *TableStylePropertiesArchive) GetBandedRows() bool {
	if m != nil && m.BandedRows != nil {
		return *m.BandedRows
	}
	return false
}

func (m *TableStylePropertiesArchive) GetBandedFill() *TSD.FillArchive {
	if m != nil {
		return m.BandedFill
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetBehavesLikeSpreadsheet() bool {
	if m != nil && m.BehavesLikeSpreadsheet != nil {
		return *m.BehavesLikeSpreadsheet
	}
	return false
}

func (m *TableStylePropertiesArchive) GetAutoResize() bool {
	if m != nil && m.AutoResize != nil {
		return *m.AutoResize
	}
	return false
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderRowSeparatorStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderRowSeparatorStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderRowBorderStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderRowBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderRowHorizontalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderRowHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderRowVerticalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderRowVerticalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderColumnBorderStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderColumnBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderColumnSeparatorStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderColumnSeparatorStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderColumnHorizontalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderColumnHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedHeaderColumnVerticalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedHeaderColumnVerticalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedFooterRowSeparatorStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedFooterRowSeparatorStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedFooterRowBorderStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedFooterRowBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedFooterRowHorizontalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedFooterRowHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedFooterRowVerticalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedFooterRowVerticalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedTableBodyHorizontalBorderStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedTableBodyHorizontalBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedTableBodyVerticalBorderStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedTableBodyVerticalBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedTableBodyHorizontalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedTableBodyHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetDeprecatedTableBodyVerticalStroke() *Deprecated_TableStrokeArchive {
	if m != nil {
		return m.DeprecatedTableBodyVerticalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetStrokePresetList() *StrokePresetListArchive {
	if m != nil {
		return m.StrokePresetList
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetVStrokesVisible() bool {
	if m != nil && m.VStrokesVisible != nil {
		return *m.VStrokesVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetHStrokesVisible() bool {
	if m != nil && m.HStrokesVisible != nil {
		return *m.HStrokesVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetHrSeparatorVisible() bool {
	if m != nil && m.HrSeparatorVisible != nil {
		return *m.HrSeparatorVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetHcSeparatorVisible() bool {
	if m != nil && m.HcSeparatorVisible != nil {
		return *m.HcSeparatorVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetFooterSeparatorVisible() bool {
	if m != nil && m.FooterSeparatorVisible != nil {
		return *m.FooterSeparatorVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetTableBorderVisible() bool {
	if m != nil && m.TableBorderVisible != nil {
		return *m.TableBorderVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetTableHeaderBorderVisible() bool {
	if m != nil && m.TableHeaderBorderVisible != nil {
		return *m.TableHeaderBorderVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetTableHcDividerVisible() bool {
	if m != nil && m.TableHcDividerVisible != nil {
		return *m.TableHcDividerVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetTableHrDividerVisible() bool {
	if m != nil && m.TableHrDividerVisible != nil {
		return *m.TableHrDividerVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetTableFooterDividerVisible() bool {
	if m != nil && m.TableFooterDividerVisible != nil {
		return *m.TableFooterDividerVisible
	}
	return false
}

func (m *TableStylePropertiesArchive) GetOBSOLETEMasterFontSize() int32 {
	if m != nil && m.OBSOLETEMasterFontSize != nil {
		return *m.OBSOLETEMasterFontSize
	}
	return 0
}

func (m *TableStylePropertiesArchive) GetMasterFontFamily() string {
	if m != nil && m.MasterFontFamily != nil {
		return *m.MasterFontFamily
	}
	return ""
}

func (m *TableStylePropertiesArchive) GetWritingDirection() TSWP.WritingDirectionType {
	if m != nil && m.WritingDirection != nil {
		return *m.WritingDirection
	}
	return TSWP.WritingDirectionType_kWritingDirectionNatural
}

func (m *TableStylePropertiesArchive) GetHeaderRowSeparatorStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderRowSeparatorStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetHeaderRowBorderStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderRowBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetHeaderRowHorizontalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderRowHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetHeaderRowVerticalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderRowVerticalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetHeaderColumnBorderStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderColumnBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetHeaderColumnSeparatorStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderColumnSeparatorStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetHeaderColumnHorizontalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderColumnHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetHeaderColumnVerticalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.HeaderColumnVerticalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetFooterRowSeparatorStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.FooterRowSeparatorStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetFooterRowBorderStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.FooterRowBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetFooterRowHorizontalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.FooterRowHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetFooterRowVerticalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.FooterRowVerticalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetTableBodyHorizontalBorderStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.TableBodyHorizontalBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetTableBodyVerticalBorderStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.TableBodyVerticalBorderStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetTableBodyHorizontalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.TableBodyHorizontalStroke
	}
	return nil
}

func (m *TableStylePropertiesArchive) GetTableBodyVerticalStroke() *TSD.StrokeArchive {
	if m != nil {
		return m.TableBodyVerticalStroke
	}
	return nil
}

type TableStylePresetArchive struct {
	Index            *int32         `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	Image            *TSP.Reference `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	StyleNetwork     *TSP.Reference `protobuf:"bytes,3,opt,name=style_network" json:"style_network,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *TableStylePresetArchive) Reset()         { *m = TableStylePresetArchive{} }
func (m *TableStylePresetArchive) String() string { return proto.CompactTextString(m) }
func (*TableStylePresetArchive) ProtoMessage()    {}

func (m *TableStylePresetArchive) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *TableStylePresetArchive) GetImage() *TSP.Reference {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *TableStylePresetArchive) GetStyleNetwork() *TSP.Reference {
	if m != nil {
		return m.StyleNetwork
	}
	return nil
}

type TableStrokePresetArchive struct {
	Index            *int32 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TableStrokePresetArchive) Reset()         { *m = TableStrokePresetArchive{} }
func (m *TableStrokePresetArchive) String() string { return proto.CompactTextString(m) }
func (*TableStrokePresetArchive) ProtoMessage()    {}

func (m *TableStrokePresetArchive) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

type ThemePresetsArchive struct {
	TableStylePresets      []*TSP.Reference `protobuf:"bytes,1,rep,name=table_style_presets" json:"table_style_presets,omitempty"`
	TableCellStrokePresets []*TSP.Reference `protobuf:"bytes,2,rep,name=table_cell_stroke_presets" json:"table_cell_stroke_presets,omitempty"`
	XXX_unrecognized       []byte           `json:"-"`
}

func (m *ThemePresetsArchive) Reset()         { *m = ThemePresetsArchive{} }
func (m *ThemePresetsArchive) String() string { return proto.CompactTextString(m) }
func (*ThemePresetsArchive) ProtoMessage()    {}

func (m *ThemePresetsArchive) GetTableStylePresets() []*TSP.Reference {
	if m != nil {
		return m.TableStylePresets
	}
	return nil
}

func (m *ThemePresetsArchive) GetTableCellStrokePresets() []*TSP.Reference {
	if m != nil {
		return m.TableCellStrokePresets
	}
	return nil
}

var E_ThemePresetsArchive_Extension = &proto.ExtensionDesc{
	ExtendedType:  (*TSS.ThemeArchive)(nil),
	ExtensionType: (*ThemePresetsArchive)(nil),
	Field:         200,
	Name:          "TST.ThemePresetsArchive.extension",
	Tag:           "bytes,200,opt,name=extension",
}

func init() {
	proto.RegisterExtension(E_ThemePresetsArchive_Extension)
}
