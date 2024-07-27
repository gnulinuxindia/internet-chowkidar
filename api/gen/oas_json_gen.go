// Code generated by ogen, DO NOT EDIT.

package genapi

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *AbuseReport) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AbuseReport) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.SiteID.Set {
			e.FieldStart("site_id")
			s.SiteID.Encode(e)
		}
	}
	{
		if s.Reason.Set {
			e.FieldStart("reason")
			s.Reason.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
	{
		if s.ResolvedAt.Set {
			e.FieldStart("resolved_at")
			s.ResolvedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfAbuseReport = [7]string{
	0: "id",
	1: "site_id",
	2: "reason",
	3: "status",
	4: "resolved_at",
	5: "created_at",
	6: "updated_at",
}

// Decode decodes AbuseReport from json.
func (s *AbuseReport) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AbuseReport to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "site_id":
			if err := func() error {
				s.SiteID.Reset()
				if err := s.SiteID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"site_id\"")
			}
		case "reason":
			if err := func() error {
				s.Reason.Reset()
				if err := s.Reason.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reason\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "resolved_at":
			if err := func() error {
				s.ResolvedAt.Reset()
				if err := s.ResolvedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"resolved_at\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AbuseReport")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AbuseReport) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AbuseReport) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *AbuseReportInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AbuseReportInput) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("site_id")
		e.Int(s.SiteID)
	}
	{
		e.FieldStart("reason")
		e.Str(s.Reason)
	}
}

var jsonFieldsNameOfAbuseReportInput = [2]string{
	0: "site_id",
	1: "reason",
}

// Decode decodes AbuseReportInput from json.
func (s *AbuseReportInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AbuseReportInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "site_id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.SiteID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"site_id\"")
			}
		case "reason":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Reason = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reason\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AbuseReportInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfAbuseReportInput) {
					name = jsonFieldsNameOfAbuseReportInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AbuseReportInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AbuseReportInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AbuseReportStatus as json.
func (s AbuseReportStatus) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes AbuseReportStatus from json.
func (s *AbuseReportStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AbuseReportStatus to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch AbuseReportStatus(v) {
	case AbuseReportStatusPending:
		*s = AbuseReportStatusPending
	case AbuseReportStatusResolved:
		*s = AbuseReportStatusResolved
	default:
		*s = AbuseReportStatus(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s AbuseReportStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AbuseReportStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Block) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Block) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.SiteID.Set {
			e.FieldStart("site_id")
			s.SiteID.Encode(e)
		}
	}
	{
		if s.IspID.Set {
			e.FieldStart("isp_id")
			s.IspID.Encode(e)
		}
	}
	{
		if s.LastReportedAt.Set {
			e.FieldStart("last_reported_at")
			s.LastReportedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.BlockReports.Set {
			e.FieldStart("block_reports")
			s.BlockReports.Encode(e)
		}
	}
	{
		if s.UnblockReports.Set {
			e.FieldStart("unblock_reports")
			s.UnblockReports.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfBlock = [8]string{
	0: "id",
	1: "site_id",
	2: "isp_id",
	3: "last_reported_at",
	4: "block_reports",
	5: "unblock_reports",
	6: "created_at",
	7: "updated_at",
}

// Decode decodes Block from json.
func (s *Block) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Block to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "site_id":
			if err := func() error {
				s.SiteID.Reset()
				if err := s.SiteID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"site_id\"")
			}
		case "isp_id":
			if err := func() error {
				s.IspID.Reset()
				if err := s.IspID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"isp_id\"")
			}
		case "last_reported_at":
			if err := func() error {
				s.LastReportedAt.Reset()
				if err := s.LastReportedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"last_reported_at\"")
			}
		case "block_reports":
			if err := func() error {
				s.BlockReports.Reset()
				if err := s.BlockReports.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"block_reports\"")
			}
		case "unblock_reports":
			if err := func() error {
				s.UnblockReports.Reset()
				if err := s.UnblockReports.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"unblock_reports\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Block")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Block) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Block) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *BlockInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *BlockInput) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("site_id")
		e.Int(s.SiteID)
	}
	{
		e.FieldStart("isp_id")
		e.Int(s.IspID)
	}
}

var jsonFieldsNameOfBlockInput = [2]string{
	0: "site_id",
	1: "isp_id",
}

// Decode decodes BlockInput from json.
func (s *BlockInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode BlockInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "site_id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.SiteID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"site_id\"")
			}
		case "isp_id":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.IspID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"isp_id\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode BlockInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfBlockInput) {
					name = jsonFieldsNameOfBlockInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *BlockInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *BlockInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ISP) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ISP) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Latitude.Set {
			e.FieldStart("latitude")
			s.Latitude.Encode(e)
		}
	}
	{
		if s.Longitude.Set {
			e.FieldStart("longitude")
			s.Longitude.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfISP = [6]string{
	0: "id",
	1: "latitude",
	2: "longitude",
	3: "name",
	4: "created_at",
	5: "updated_at",
}

// Decode decodes ISP from json.
func (s *ISP) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ISP to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "latitude":
			if err := func() error {
				s.Latitude.Reset()
				if err := s.Latitude.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latitude\"")
			}
		case "longitude":
			if err := func() error {
				s.Longitude.Reset()
				if err := s.Longitude.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"longitude\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ISP")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ISP) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ISP) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ISPInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ISPInput) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("latitude")
		e.Float32(s.Latitude)
	}
	{
		e.FieldStart("longitude")
		e.Float32(s.Longitude)
	}
	{
		e.FieldStart("name")
		e.Str(s.Name)
	}
}

var jsonFieldsNameOfISPInput = [3]string{
	0: "latitude",
	1: "longitude",
	2: "name",
}

// Decode decodes ISPInput from json.
func (s *ISPInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ISPInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "latitude":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Float32()
				s.Latitude = float32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"latitude\"")
			}
		case "longitude":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Float32()
				s.Longitude = float32(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"longitude\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Str()
				s.Name = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ISPInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfISPInput) {
					name = jsonFieldsNameOfISPInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ISPInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ISPInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes AbuseReportStatus as json.
func (o OptAbuseReportStatus) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes AbuseReportStatus from json.
func (o *OptAbuseReportStatus) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptAbuseReportStatus to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptAbuseReportStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptAbuseReportStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

// Encode encodes float32 as json.
func (o OptFloat32) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Float32(float32(o.Value))
}

// Decode decodes float32 from json.
func (o *OptFloat32) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptFloat32 to nil")
	}
	o.Set = true
	v, err := d.Float32()
	if err != nil {
		return err
	}
	o.Value = float32(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptFloat32) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptFloat32) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SiteSuggestionStatus as json.
func (o OptSiteSuggestionStatus) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes SiteSuggestionStatus from json.
func (o *OptSiteSuggestionStatus) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptSiteSuggestionStatus to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptSiteSuggestionStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptSiteSuggestionStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Site) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Site) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Domain.Set {
			e.FieldStart("domain")
			s.Domain.Encode(e)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfSite = [4]string{
	0: "id",
	1: "domain",
	2: "created_at",
	3: "updated_at",
}

// Decode decodes Site from json.
func (s *Site) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Site to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "domain":
			if err := func() error {
				s.Domain.Reset()
				if err := s.Domain.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"domain\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Site")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Site) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Site) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SiteInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SiteInput) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("domain")
		e.Str(s.Domain)
	}
}

var jsonFieldsNameOfSiteInput = [1]string{
	0: "domain",
}

// Decode decodes SiteInput from json.
func (s *SiteInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SiteInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "domain":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Domain = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"domain\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SiteInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSiteInput) {
					name = jsonFieldsNameOfSiteInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SiteInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SiteInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SiteSuggestion) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SiteSuggestion) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.SiteID.Set {
			e.FieldStart("site_id")
			s.SiteID.Encode(e)
		}
	}
	{
		if s.Reason.Set {
			e.FieldStart("reason")
			s.Reason.Encode(e)
		}
	}
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
	{
		if s.ResolvedAt.Set {
			e.FieldStart("resolved_at")
			s.ResolvedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.CreatedAt.Set {
			e.FieldStart("created_at")
			s.CreatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		if s.UpdatedAt.Set {
			e.FieldStart("updated_at")
			s.UpdatedAt.Encode(e, json.EncodeDateTime)
		}
	}
}

var jsonFieldsNameOfSiteSuggestion = [7]string{
	0: "id",
	1: "site_id",
	2: "reason",
	3: "status",
	4: "resolved_at",
	5: "created_at",
	6: "updated_at",
}

// Decode decodes SiteSuggestion from json.
func (s *SiteSuggestion) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SiteSuggestion to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "site_id":
			if err := func() error {
				s.SiteID.Reset()
				if err := s.SiteID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"site_id\"")
			}
		case "reason":
			if err := func() error {
				s.Reason.Reset()
				if err := s.Reason.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reason\"")
			}
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		case "resolved_at":
			if err := func() error {
				s.ResolvedAt.Reset()
				if err := s.ResolvedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"resolved_at\"")
			}
		case "created_at":
			if err := func() error {
				s.CreatedAt.Reset()
				if err := s.CreatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"created_at\"")
			}
		case "updated_at":
			if err := func() error {
				s.UpdatedAt.Reset()
				if err := s.UpdatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"updated_at\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SiteSuggestion")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SiteSuggestion) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SiteSuggestion) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SiteSuggestionInput) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SiteSuggestionInput) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("site_id")
		e.Int(s.SiteID)
	}
	{
		e.FieldStart("reason")
		e.Str(s.Reason)
	}
}

var jsonFieldsNameOfSiteSuggestionInput = [2]string{
	0: "site_id",
	1: "reason",
}

// Decode decodes SiteSuggestionInput from json.
func (s *SiteSuggestionInput) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SiteSuggestionInput to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "site_id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.SiteID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"site_id\"")
			}
		case "reason":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Reason = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reason\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SiteSuggestionInput")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSiteSuggestionInput) {
					name = jsonFieldsNameOfSiteSuggestionInput[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SiteSuggestionInput) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SiteSuggestionInput) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SiteSuggestionStatus as json.
func (s SiteSuggestionStatus) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes SiteSuggestionStatus from json.
func (s *SiteSuggestionStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SiteSuggestionStatus to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch SiteSuggestionStatus(v) {
	case SiteSuggestionStatusPending:
		*s = SiteSuggestionStatusPending
	case SiteSuggestionStatusResolved:
		*s = SiteSuggestionStatusResolved
	default:
		*s = SiteSuggestionStatus(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SiteSuggestionStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SiteSuggestionStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
