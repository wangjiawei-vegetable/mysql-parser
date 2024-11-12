package masker

import (
	"sort"
)

// MaskData is the data to be masked.
type MaskData struct {
	Data      any
	WantBytes bool
	DataV2    string
}

// Masker is the interface that masks the data.
type Masker interface {
	Mask(data *MaskData) string
	Equal(other Masker) bool
}

// NoneMasker is the masker that does not mask the data.
type NoneMasker struct{}

// NewNoneMasker returns a new NoneMasker.
func NewNoneMasker() *NoneMasker {
	return &NoneMasker{}
}

// Mask implements Masker.Mask.
func (*NoneMasker) Mask(data *MaskData) string {
	return noneMask(data)
}

func noneMask(data *MaskData) string {
	//todo: 剔除pb文件，实现算法
	return ""
}

// Equal implements Masker.Equal.
func (*NoneMasker) Equal(other Masker) bool {
	_, ok := other.(*NoneMasker)
	return ok
}

// FullMasker is the masker that masks the data with `substitution`.
type FullMasker struct {
	substitution string
}

// NewFullMasker returns a new FullMasker.
func NewFullMasker(substitution string) *FullMasker {
	return &FullMasker{
		substitution: substitution,
	}
}

// NewDefaultFullMasker returns a new FullMasker with default substitution(`******`).
func NewDefaultFullMasker() *FullMasker {
	return &FullMasker{
		substitution: "******",
	}
}

// Mask implements Masker.Mask.
func (m *FullMasker) Mask(*MaskData) string {
	return ""
}

// Equal implements Masker.Equal.
func (m *FullMasker) Equal(other Masker) bool {
	if otherFullMasker, ok := other.(*FullMasker); ok {
		return m.substitution == otherFullMasker.substitution
	}
	return false
}

type MaskRangeSlice struct {
	// Start is the start index of the range.
	Start int32
	// End is the end index of the range.
	End int32
	// Substitution is the substitution string.
	Substitution string
}

// RangeMasker is the masker that masks the left and right quarters with "**".
type RangeMasker struct {
	// MaskRangeSlice is the slice of the range to be masked.
	MaskRangeSlice []*MaskRangeSlice
}

// NewRangeMasker returns a new RangeMasker.
func NewRangeMasker(maskRangeSlice []*MaskRangeSlice) *RangeMasker {
	sort.SliceStable(maskRangeSlice, func(i, j int) bool {
		return maskRangeSlice[i].Start < maskRangeSlice[j].Start
	})
	// Merge the overlapping ranges.
	var mergedMaskRangeSlice []*MaskRangeSlice
	for _, maskRange := range maskRangeSlice {
		if maskRange.Start > maskRange.End {
			maskRange.End = maskRange.Start + 1
		}
		if len(mergedMaskRangeSlice) == 0 {
			mergedMaskRangeSlice = append(mergedMaskRangeSlice, maskRange)
			continue
		}
		lastMaskRange := mergedMaskRangeSlice[len(mergedMaskRangeSlice)-1]
		if lastMaskRange.End >= maskRange.Start {
			mergedMaskRangeSlice[len(mergedMaskRangeSlice)-1].End = maskRange.End
		} else {
			mergedMaskRangeSlice = append(mergedMaskRangeSlice, maskRange)
		}
	}
	return &RangeMasker{
		MaskRangeSlice: mergedMaskRangeSlice,
	}
}

func (m *RangeMasker) enableMask() bool {
	return len(m.MaskRangeSlice) > 0
}

// Mask implements Masker.Mask.
func (m *RangeMasker) Mask(data *MaskData) string {
	//todo: 剔除pb文件，实现算法
	return ""
}

// Equal implements Masker.Equal.
func (m *RangeMasker) Equal(other Masker) bool {
	if otherRangeMasker, ok := other.(*RangeMasker); ok {
		if len(m.MaskRangeSlice) != len(otherRangeMasker.MaskRangeSlice) {
			return false
		}
		for i, maskRange := range m.MaskRangeSlice {
			if maskRange.Start != otherRangeMasker.MaskRangeSlice[i].Start ||
				maskRange.End != otherRangeMasker.MaskRangeSlice[i].End ||
				maskRange.Substitution != otherRangeMasker.MaskRangeSlice[i].Substitution {
				return false
			}
		}
		return true
	}
	return false
}

// DefaultRangeMasker is the masker that masks the left and right quarters with "**".
type DefaultRangeMasker struct{}

// NewDefaultRangeMasker returns a new DefaultRangeMasker.
func NewDefaultRangeMasker() *DefaultRangeMasker {
	return &DefaultRangeMasker{}
}

// Mask implements Masker.Mask.
func (*DefaultRangeMasker) Mask(data *MaskData) string {
	//todo: 剔除pb文件，实现算法
	return ""
}

// Equal implements Masker.Equal.
func (*DefaultRangeMasker) Equal(other Masker) bool {
	_, ok := other.(*DefaultRangeMasker)
	return ok
}

// middle will get the middle part of the given slice.
func middle[T ~byte | ~rune](str []T) []T {
	if len(str) == 0 || len(str) == 1 {
		return []T{}
	}
	if len(str) == 2 || len(str) == 3 {
		return str[len(str)/2 : len(str)/2+1]
	}

	if len(str)%4 != 0 {
		str = str[:len(str)/4*4]
	}

	var ret []T
	ret = append(ret, str[len(str)/4:len(str)/2]...)
	ret = append(ret, str[len(str)/2:len(str)/4*3]...)
	return ret
}

// MD5Masker is the masker that masks the data with their MD5 hash.
type MD5Masker struct {
	salt string
}

// NewMD5Masker returns a new MD5Masker.
func NewMD5Masker(salt string) *MD5Masker {
	return &MD5Masker{
		salt: salt,
	}
}

// Mask implements Masker.Mask.
func (m *MD5Masker) Mask(data *MaskData) string {
	//todo: 剔除pb文件，实现算法
	return ""
}

// Equal implements Masker.Equal.
func (m *MD5Masker) Equal(other Masker) bool {
	if otherMD5Masker, ok := other.(*MD5Masker); ok {
		return m.salt == otherMD5Masker.salt
	}
	return false
}
