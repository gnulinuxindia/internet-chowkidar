// Code generated by ogen, DO NOT EDIT.

package genapi

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetSiteParams is parameters of getSite operation.
type GetSiteParams struct {
	ID int
}

func unpackGetSiteParams(packed middleware.Parameters) (params GetSiteParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(int)
	}
	return params
}

func decodeGetSiteParams(args [1]string, argsEscaped bool, r *http.Request) (params GetSiteParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListISPsParams is parameters of listISPs operation.
type ListISPsParams struct {
	// Number of ISPs to return.
	Limit OptInt
	// Number of ISPs to skip.
	Offset OptInt
	// Sort ISPs by field.
	Sort OptString
	// Sort order.
	Order OptListISPsOrder
}

func unpackListISPsParams(packed middleware.Parameters) (params ListISPsParams) {
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "offset",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Offset = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sort",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Sort = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "order",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Order = v.(OptListISPsOrder)
		}
	}
	return params
}

func decodeListISPsParams(args [0]string, argsEscaped bool, r *http.Request) (params ListISPsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: limit.
	{
		val := int(50)
		params.Limit.SetTo(val)
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: offset.
	{
		val := int(0)
		params.Offset.SetTo(val)
	}
	// Decode query: offset.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOffsetVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotOffsetVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Offset.SetTo(paramsDotOffsetVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "offset",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: sort.
	{
		val := string("id")
		params.Sort.SetTo(val)
	}
	// Decode query: sort.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSortVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Sort.SetTo(paramsDotSortVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sort",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: order.
	{
		val := ListISPsOrder("asc")
		params.Order.SetTo(val)
	}
	// Decode query: order.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderVal ListISPsOrder
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderVal = ListISPsOrder(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Order.SetTo(paramsDotOrderVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Order.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "order",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListSitesParams is parameters of listSites operation.
type ListSitesParams struct {
	// Filter sites by category, separated by commas.
	Category OptString
	// Number of sites to return.
	Limit OptInt
	// Number of sites to skip.
	Offset OptInt
	// Sort sites by field.
	Sort OptString
	// Sort order.
	Order OptListSitesOrder
}

func unpackListSitesParams(packed middleware.Parameters) (params ListSitesParams) {
	{
		key := middleware.ParameterKey{
			Name: "category",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Category = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "offset",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Offset = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sort",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Sort = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "order",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Order = v.(OptListSitesOrder)
		}
	}
	return params
}

func decodeListSitesParams(args [0]string, argsEscaped bool, r *http.Request) (params ListSitesParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: category.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "category",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCategoryVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCategoryVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Category.SetTo(paramsDotCategoryVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "category",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: limit.
	{
		val := int(50)
		params.Limit.SetTo(val)
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: offset.
	{
		val := int(0)
		params.Offset.SetTo(val)
	}
	// Decode query: offset.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOffsetVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotOffsetVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Offset.SetTo(paramsDotOffsetVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "offset",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: sort.
	{
		val := string("id")
		params.Sort.SetTo(val)
	}
	// Decode query: sort.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSortVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Sort.SetTo(paramsDotSortVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sort",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: order.
	{
		val := ListSitesOrder("asc")
		params.Order.SetTo(val)
	}
	// Decode query: order.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "order",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOrderVal ListSitesOrder
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotOrderVal = ListSitesOrder(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Order.SetTo(paramsDotOrderVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Order.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "order",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
