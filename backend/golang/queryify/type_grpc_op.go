package queryify

import (
	gRPCCommon "github.com/WM1rr0rB8/contractsTest/gen/go/common/filter/v1"
	"github.com/WM1rr0rB8/librariesTest/backend/golang/errors"
	"github.com/WM1rr0rB8/librariesTest/backend/golang/sfqb"
)

const (
	badOperator = "bad operator: "
)

var ErrUnknownOperator = errors.New("unknown operator")

type Operator interface {
	String() string
}

// MapOperator takes an operator of any type and maps it to a corresponding method.
// It supports various types of operators including boolean, string, integer, float,
// and array. If the operator type is not recognized, it returns an error.
func MapOperator(op Operator) (sfqb.Method, error) {
	switch o := op.(type) {
	case gRPCCommon.BoolFieldFilter_Operator:
		return mapBoolOperator(o)
	case gRPCCommon.StringFieldFilter_Operator:
		return mapStringOperator(o)
	case gRPCCommon.StringNumberFieldFilter_Operator:
		return mapStringNumberOperator(o)
	case gRPCCommon.IntFieldFilter_Operator:
		return mapIntOperator(o)
	case gRPCCommon.FloatFieldFilter_Operator:
		return mapFloatOperator(o)
	case gRPCCommon.ArrayStringFieldFilter_Operator:
		return mapArrayStringOperator(o)
	case gRPCCommon.ArrayIntFieldFilter_Operator:
		return mapArrayIntOperator(o)
	case gRPCCommon.ArrayFloatFieldFilter_Operator:
		return mapArrayFloatOperator(o)
	case gRPCCommon.RangeIntFieldFilter_Operator:
		return mapRangeIntOperator(o)
	case gRPCCommon.RangeStringFieldFilter_Operator:
		return mapRangeStringOperator(o)
	default:
		return "", ErrUnknownOperator
	}
}

// mapBoolOperator takes a boolean operator and maps it to a corresponding method.
// It supports unspecified and equality operators. If the operator is not recognized,
// it returns an error.
func mapBoolOperator(op gRPCCommon.BoolFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.BoolFieldFilter_OPERATOR_EQ:
		return sfqb.EQ, nil
	case gRPCCommon.BoolFieldFilter_OPERATOR_NEQ:
		return sfqb.NE, nil
	case gRPCCommon.BoolFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapStringOperator takes a string operator and maps it to a corresponding method.
// It supports unspecified, equality, non-equality, and like operators. If the operator
// is not recognized, it returns an error.
func mapStringOperator(op gRPCCommon.StringFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.StringFieldFilter_OPERATOR_EQ:
		return sfqb.EQ, nil
	case gRPCCommon.StringFieldFilter_OPERATOR_NEQ:
		return sfqb.NE, nil
	case gRPCCommon.StringFieldFilter_OPERATOR_LIKE:
		return sfqb.LIKE, nil
	case gRPCCommon.StringFieldFilter_OPERATOR_ILIKE:
		return sfqb.ILIKE, nil
	case gRPCCommon.StringFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapStringNumberOperator takes a string operator and maps it to a corresponding method.
// It supports unspecified, equality, non-equality,
// If the operator is not recognized, it returns an error.
func mapStringNumberOperator(op gRPCCommon.StringNumberFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.StringNumberFieldFilter_OPERATOR_EQ:
		return sfqb.EQ, nil
	case gRPCCommon.StringNumberFieldFilter_OPERATOR_NEQ:
		return sfqb.NE, nil
	case gRPCCommon.StringNumberFieldFilter_OPERATOR_LT:
		return sfqb.LT, nil
	case gRPCCommon.StringNumberFieldFilter_OPERATOR_LTE:
		return sfqb.LTE, nil
	case gRPCCommon.StringNumberFieldFilter_OPERATOR_GT:
		return sfqb.GT, nil
	case gRPCCommon.StringNumberFieldFilter_OPERATOR_GTE:
		return sfqb.GTE, nil
	case gRPCCommon.StringNumberFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapIntOperator takes an integer operator and maps it to a corresponding method.
// It supports unspecified, equality, non-equality, less than, less than or equal,
// greater than, and greater than or equal operators. If the operator is not recognized,
// it returns an error.
func mapIntOperator(op gRPCCommon.IntFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.IntFieldFilter_OPERATOR_EQ:
		return sfqb.EQ, nil
	case gRPCCommon.IntFieldFilter_OPERATOR_NEQ:
		return sfqb.NE, nil
	case gRPCCommon.IntFieldFilter_OPERATOR_LT:
		return sfqb.LT, nil
	case gRPCCommon.IntFieldFilter_OPERATOR_LTE:
		return sfqb.LTE, nil
	case gRPCCommon.IntFieldFilter_OPERATOR_GT:
		return sfqb.GT, nil
	case gRPCCommon.IntFieldFilter_OPERATOR_GTE:
		return sfqb.GTE, nil
	case gRPCCommon.IntFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapFloatOperator takes a float operator and maps it to a corresponding method.
// It supports unspecified, equality, non-equality, less than, less than or equal,
// greater than, and greater than or equal operators. If the operator is not recognized,
// it returns an error.
func mapFloatOperator(op gRPCCommon.FloatFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.FloatFieldFilter_OPERATOR_EQ:
		return sfqb.EQ, nil
	case gRPCCommon.FloatFieldFilter_OPERATOR_NEQ:
		return sfqb.NE, nil
	case gRPCCommon.FloatFieldFilter_OPERATOR_LT:
		return sfqb.LT, nil
	case gRPCCommon.FloatFieldFilter_OPERATOR_LTE:
		return sfqb.LTE, nil
	case gRPCCommon.FloatFieldFilter_OPERATOR_GT:
		return sfqb.GT, nil
	case gRPCCommon.FloatFieldFilter_OPERATOR_GTE:
		return sfqb.GTE, nil
	case gRPCCommon.FloatFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapArrayStringOperator takes an array string operator and maps it to a corresponding method.
// It supports unspecified, in, and not in operators. If the operator is not recognized,
// it returns an error.
func mapArrayStringOperator(op gRPCCommon.ArrayStringFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.ArrayStringFieldFilter_OPERATOR_IN:
		return sfqb.IN, nil
	case gRPCCommon.ArrayStringFieldFilter_OPERATOR_NIN:
		return sfqb.NIN, nil
	case gRPCCommon.ArrayStringFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapArrayIntOperator takes an array integer operator and maps it to a corresponding method.
// It supports unspecified, in, and not in operators. If the operator is not recognized,
// it returns an error.
func mapArrayIntOperator(op gRPCCommon.ArrayIntFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.ArrayIntFieldFilter_OPERATOR_IN:
		return sfqb.IN, nil
	case gRPCCommon.ArrayIntFieldFilter_OPERATOR_NIN:
		return sfqb.NIN, nil
	case gRPCCommon.ArrayIntFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapArrayFloatOperator takes an array float operator and maps it to a corresponding method.
// It supports unspecified, in, and not in operators. If the operator is not recognized,
// it returns an error.
func mapArrayFloatOperator(op gRPCCommon.ArrayFloatFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.ArrayFloatFieldFilter_OPERATOR_IN:
		return sfqb.IN, nil
	case gRPCCommon.ArrayFloatFieldFilter_OPERATOR_NIN:
		return sfqb.NIN, nil
	case gRPCCommon.ArrayFloatFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

// mapRangeIntOperator takes a range integer operator and maps it to a corresponding method.
func mapRangeIntOperator(op gRPCCommon.RangeIntFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.RangeIntFieldFilter_OPERATOR_RANGE:
		return sfqb.Range, nil
	case gRPCCommon.RangeIntFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}

func mapRangeStringOperator(op gRPCCommon.RangeStringFieldFilter_Operator) (sfqb.Method, error) {
	switch op {
	case gRPCCommon.RangeStringFieldFilter_OPERATOR_RANGE:
		return sfqb.Range, nil
	case gRPCCommon.RangeStringFieldFilter_OPERATOR_UNSPECIFIED:
		fallthrough
	default:
		return "", errors.New(badOperator + op.String())
	}
}
