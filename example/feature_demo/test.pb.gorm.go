// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/feature_demo/test.proto

package example

import gorm1 "github.com/jinzhu/gorm"
import context "golang.org/x/net/context"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TestTypesORM ...  test_id_and_stuff is a message that serves as an example
type TestTypesORM struct {
	// Skipping field: ApiOnlyString
	Numbers        []int32
	OptionalString *string
	BecomesInt     int32
	// Empty type has no ORM equivalency
}

func (TestTypesORM) TableName() string {
	return "smorgasbord"
}

// ConvertTestTypesToORM takes a pb object and returns an orm object
func ConvertTestTypesToORM(from TestTypes) TestTypesORM {
	to := TestTypesORM{}
	// Skipping field: ApiOnlyString
	for _, v := range from.Numbers {
		to.Numbers = append(to.Numbers, v)
	}
	if from.OptionalString != nil {
		v := from.OptionalString.Value
		to.OptionalString = &v
	}
	to.BecomesInt = int32(from.BecomesInt)
	return to
}

// ConvertTestTypesFromORM takes an orm object and returns a pb object
func ConvertTestTypesFromORM(from TestTypesORM) TestTypes {
	to := TestTypes{}
	// Skipping field: ApiOnlyString
	for _, v := range from.Numbers {
		to.Numbers = append(to.Numbers, v)
	}
	if from.OptionalString != nil {
		to.OptionalString = &google_protobuf1.StringValue{Value: *from.OptionalString}
	}
	to.BecomesInt = TestTypesStatus(from.BecomesInt)
	return to
}

// TypeWithIDORM no comment was provided for message type
type TypeWithIDORM struct {
	UUID          int32  `gorm:"primary_key"`
	IP            string `gorm:"ip_addr"`
	Things        []*TestTypesORM
	ANestedObject *TestTypesORM
}

func (TypeWithIDORM) TableName() string {
	return "type_with_ids"
}

// ConvertTypeWithIDToORM takes a pb object and returns an orm object
func ConvertTypeWithIDToORM(from TypeWithId) TypeWithIDORM {
	to := TypeWithIDORM{}
	to.IP = from.Ip
	for _, v := range from.Things {
		if from.Things != nil {
			tempThings := ConvertTestTypesToORM(*v)
			to.Things = append(to.Things, &tempThings)
		} else {
			to.Things = append(to.Things, nil)
		}
	}
	if from.ANestedObject != nil {
		tempANestedObject := ConvertTestTypesToORM(*from.ANestedObject)
		to.ANestedObject = &tempANestedObject
	}
	return to
}

// ConvertTypeWithIDFromORM takes an orm object and returns a pb object
func ConvertTypeWithIDFromORM(from TypeWithIDORM) TypeWithId {
	to := TypeWithId{}
	to.Ip = from.IP
	for _, v := range from.Things {
		if from.Things != nil {
			tempThings := ConvertTestTypesFromORM(*v)
			to.Things = append(to.Things, &tempThings)
		} else {
			to.Things = append(to.Things, nil)
		}
	}
	if from.ANestedObject != nil {
		tempANestedObject := ConvertTestTypesFromORM(*from.ANestedObject)
		to.ANestedObject = &tempANestedObject
	}
	return to
}

// TypeBecomesEmptyORM no comment was provided for message type
type TypeBecomesEmptyORM struct {
	// Can't work with type *ApiOnlyType, not tagged as ormable
}

func (TypeBecomesEmptyORM) TableName() string {
	return "type_becomes_empties"
}

// ConvertTypeBecomesEmptyToORM takes a pb object and returns an orm object
func ConvertTypeBecomesEmptyToORM(from TypeBecomesEmpty) TypeBecomesEmptyORM {
	to := TypeBecomesEmptyORM{}
	return to
}

// ConvertTypeBecomesEmptyFromORM takes an orm object and returns a pb object
func ConvertTypeBecomesEmptyFromORM(from TypeBecomesEmptyORM) TypeBecomesEmpty {
	to := TypeBecomesEmpty{}
	return to
}

// DefaultCreateTestTypes executes a basic gorm create call
func DefaultCreateTestTypes(ctx context.Context, in *TestTypes, db gorm1.DB) (*TestTypes, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCreateTestTypes")
	}
	ormObj := ConvertTestTypesToORM(*in)
	db.Create(&ormObj)
	pbResponse := ConvertTestTypesFromORM(ormObj)
	return &pbResponse, nil
}

// DefaultReadTestTypes executes a basic gorm read call
func DefaultReadTestTypes(ctx context.Context, in *TestTypes, db gorm1.DB) (*TestTypes, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultReadTestTypes")
	}
	ormParams := ConvertTestTypesToORM(*in)
	ormResponse := TestTypesORM{}
	db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse)
	pbResponse := ConvertTestTypesFromORM(ormResponse)
	return &pbResponse, nil
}

// DefaultUpdateTestTypes executes a basic gorm update call
func DefaultUpdateTestTypes(ctx context.Context, in *TestTypes, db gorm1.DB) (*TestTypes, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultUpdateTestTypes")
	}
	ormObj := ConvertTestTypesToORM(*in)
	db.Save(&ormObj)
	pbResponse := ConvertTestTypesFromORM(ormObj)
	return &pbResponse, nil
}

// DefaultDeleteTestTypes executes a basic gorm delete call
func DefaultDeleteTestTypes(ctx context.Context, in *TestTypes, db gorm1.DB) error {
	if in == nil {
		return fmt.Errorf("Nil argument to DefaultDeleteTestTypes")
	}
	ormObj := ConvertTestTypesToORM(*in)
	db.Where(&ormObj).Delete(&TestTypesORM{})
	return nil
}

// DefaultCreateTypeWithID executes a basic gorm create call
func DefaultCreateTypeWithID(ctx context.Context, in *TypeWithId, db gorm1.DB) (*TypeWithId, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCreateTypeWithID")
	}
	ormObj := ConvertTypeWithIDToORM(*in)
	db.Create(&ormObj)
	pbResponse := ConvertTypeWithIDFromORM(ormObj)
	return &pbResponse, nil
}

// DefaultReadTypeWithID executes a basic gorm read call
func DefaultReadTypeWithID(ctx context.Context, in *TypeWithId, db gorm1.DB) (*TypeWithId, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultReadTypeWithID")
	}
	ormParams := ConvertTypeWithIDToORM(*in)
	ormResponse := TypeWithIDORM{}
	db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse)
	pbResponse := ConvertTypeWithIDFromORM(ormResponse)
	return &pbResponse, nil
}

// DefaultUpdateTypeWithID executes a basic gorm update call
func DefaultUpdateTypeWithID(ctx context.Context, in *TypeWithId, db gorm1.DB) (*TypeWithId, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultUpdateTypeWithID")
	}
	ormObj := ConvertTypeWithIDToORM(*in)
	db.Save(&ormObj)
	pbResponse := ConvertTypeWithIDFromORM(ormObj)
	return &pbResponse, nil
}

// DefaultDeleteTypeWithID executes a basic gorm delete call
func DefaultDeleteTypeWithID(ctx context.Context, in *TypeWithId, db gorm1.DB) error {
	if in == nil {
		return fmt.Errorf("Nil argument to DefaultDeleteTypeWithID")
	}
	ormObj := ConvertTypeWithIDToORM(*in)
	db.Where(&ormObj).Delete(&TypeWithIDORM{})
	return nil
}

// DefaultCreateTypeBecomesEmpty executes a basic gorm create call
func DefaultCreateTypeBecomesEmpty(ctx context.Context, in *TypeBecomesEmpty, db gorm1.DB) (*TypeBecomesEmpty, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCreateTypeBecomesEmpty")
	}
	ormObj := ConvertTypeBecomesEmptyToORM(*in)
	db.Create(&ormObj)
	pbResponse := ConvertTypeBecomesEmptyFromORM(ormObj)
	return &pbResponse, nil
}

// DefaultReadTypeBecomesEmpty executes a basic gorm read call
func DefaultReadTypeBecomesEmpty(ctx context.Context, in *TypeBecomesEmpty, db gorm1.DB) (*TypeBecomesEmpty, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultReadTypeBecomesEmpty")
	}
	ormParams := ConvertTypeBecomesEmptyToORM(*in)
	ormResponse := TypeBecomesEmptyORM{}
	db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse)
	pbResponse := ConvertTypeBecomesEmptyFromORM(ormResponse)
	return &pbResponse, nil
}

// DefaultUpdateTypeBecomesEmpty executes a basic gorm update call
func DefaultUpdateTypeBecomesEmpty(ctx context.Context, in *TypeBecomesEmpty, db gorm1.DB) (*TypeBecomesEmpty, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultUpdateTypeBecomesEmpty")
	}
	ormObj := ConvertTypeBecomesEmptyToORM(*in)
	db.Save(&ormObj)
	pbResponse := ConvertTypeBecomesEmptyFromORM(ormObj)
	return &pbResponse, nil
}

// DefaultDeleteTypeBecomesEmpty executes a basic gorm delete call
func DefaultDeleteTypeBecomesEmpty(ctx context.Context, in *TypeBecomesEmpty, db gorm1.DB) error {
	if in == nil {
		return fmt.Errorf("Nil argument to DefaultDeleteTypeBecomesEmpty")
	}
	ormObj := ConvertTypeBecomesEmptyToORM(*in)
	db.Where(&ormObj).Delete(&TypeBecomesEmptyORM{})
	return nil
}
