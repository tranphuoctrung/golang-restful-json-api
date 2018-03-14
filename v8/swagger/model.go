package swagger

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response validationError
type SwaggValidationError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// Success response
// swagger:response ok
type SwaggScsResp struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Code int `json:"code"`
	}
}

// Error Bad Request
// swagger:response badReq
type SwaggErrBadReq struct {
	// in:body
	Body struct {
		// HTTP status code 400 - Status Bad Request
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Forbidden
// swagger:response forbidden
type SwaggErrForbidden struct {
	// in:body
	Body struct {
		// HTTP status code 403 - Forbidden
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Not Found
// swagger:response notFound
type SwaggErrNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 - Not Found
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Conflict
// swagger:response conflict
type SwaggErrConflict struct {
	// in:body
	Body struct {
		// HTTP status code 409 - Conflict
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Interval Server
// swagger:response internal
type SwaggErrInternal struct {
	// in:body
	Body struct {
		// HTTP status code 500 - Internal server error
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}
