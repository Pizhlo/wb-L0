package errs

import "errors"

// Database
var NotFound = errors.New("not found")

// Cache
var NilOrder = errors.New("cache: order is nil")
var NilDelivery = errors.New("cache: delivery is nil")
var NilPayment = errors.New("cache: payment is nil")

var UnableConvertOrder = errors.New("cache: unable to convert order (any) to type models.Order")
var UnableConvertDelivery = errors.New("cache: unable to convert delivery (any) to type models.Delivery")
var UnableConvertPayment = errors.New("cache: unable to convert payment (any) to type models.Payment")
