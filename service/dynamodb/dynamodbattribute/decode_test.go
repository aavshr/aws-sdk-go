package dynamodbattribute

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/aavshr/aws-sdk-go/aws"
	"github.com/aavshr/aws-sdk-go/aws/awserr"
	"github.com/aavshr/aws-sdk-go/service/dynamodb"
)

func TestUnmarshalErrorTypes(t *testing.T) {
	var _ awserr.Error = (*UnmarshalTypeError)(nil)
	var _ awserr.Error = (*InvalidUnmarshalError)(nil)
}

func TestUnmarshalShared(t *testing.T) {
	for i, c := range sharedTestCases {
		err := Unmarshal(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

func TestUnmarshal(t *testing.T) {
	cases := []struct {
		in               *dynamodb.AttributeValue
		actual, expected interface{}
		err              error
	}{
		//------------
		// Sets
		//------------
		{
			in: &dynamodb.AttributeValue{BS: [][]byte{
				{48, 49}, {50, 51},
			}},
			actual:   &[][]byte{},
			expected: [][]byte{{48, 49}, {50, 51}},
		},
		{
			in: &dynamodb.AttributeValue{NS: []*string{
				aws.String("123"), aws.String("321"),
			}},
			actual:   &[]int{},
			expected: []int{123, 321},
		},
		{
			in: &dynamodb.AttributeValue{NS: []*string{
				aws.String("123"), aws.String("321"),
			}},
			actual:   &[]interface{}{},
			expected: []interface{}{123., 321.},
		},
		{
			in: &dynamodb.AttributeValue{SS: []*string{
				aws.String("abc"), aws.String("123"),
			}},
			actual:   &[]string{},
			expected: &[]string{"abc", "123"},
		},
		{
			in: &dynamodb.AttributeValue{SS: []*string{
				aws.String("abc"), aws.String("123"),
			}},
			actual:   &[]*string{},
			expected: &[]*string{aws.String("abc"), aws.String("123")},
		},
		//------------
		// Interfaces
		//------------
		{
			in: &dynamodb.AttributeValue{B: []byte{48, 49}},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: []byte{48, 49},
		},
		{
			in: &dynamodb.AttributeValue{BS: [][]byte{
				{48, 49}, {50, 51},
			}},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: [][]byte{{48, 49}, {50, 51}},
		},
		{
			in: &dynamodb.AttributeValue{BOOL: aws.Bool(true)},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: bool(true),
		},
		{
			in: &dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{
				{S: aws.String("abc")}, {S: aws.String("123")},
			}},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: []interface{}{"abc", "123"},
		},
		{
			in: &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
				"123": {S: aws.String("abc")},
				"abc": {S: aws.String("123")},
			}},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: map[string]interface{}{"123": "abc", "abc": "123"},
		},
		{
			in: &dynamodb.AttributeValue{N: aws.String("123")},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: float64(123),
		},
		{
			in: &dynamodb.AttributeValue{NS: []*string{
				aws.String("123"), aws.String("321"),
			}},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: []float64{123., 321.},
		},
		{
			in: &dynamodb.AttributeValue{S: aws.String("123")},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: "123",
		},
		{
			in: &dynamodb.AttributeValue{SS: []*string{
				aws.String("123"), aws.String("321"),
			}},
			actual: func() interface{} {
				var v interface{}
				return &v
			}(),
			expected: []string{"123", "321"},
		},
		{
			in: &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
				"abc": {S: aws.String("123")},
				"Cba": {S: aws.String("321")},
			}},
			actual:   &struct{ Abc, Cba string }{},
			expected: struct{ Abc, Cba string }{Abc: "123", Cba: "321"},
		},
		{
			in:     &dynamodb.AttributeValue{N: aws.String("512")},
			actual: new(uint8),
			err: &UnmarshalTypeError{
				Value: fmt.Sprintf("number overflow, 512"),
				Type:  reflect.TypeOf(uint8(0)),
			},
		},
		// -------
		// Empty Values
		// -------
		{
			in:       &dynamodb.AttributeValue{B: []byte{}},
			actual:   &[]byte{},
			expected: []byte{},
		},
		{
			in:       &dynamodb.AttributeValue{BS: [][]byte{}},
			actual:   &[][]byte{},
			expected: [][]byte{},
		},
		{
			in:       &dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{}},
			actual:   &[]interface{}{},
			expected: []interface{}{},
		},
		{
			in:       &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{}},
			actual:   &map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			in:     &dynamodb.AttributeValue{N: aws.String("")},
			actual: new(int),
			err:    fmt.Errorf("invalid syntax"),
		},
		{
			in:       &dynamodb.AttributeValue{NS: []*string{}},
			actual:   &[]*string{},
			expected: []*string{},
		},
		{
			in:       &dynamodb.AttributeValue{S: aws.String("")},
			actual:   new(string),
			expected: "",
		},
		{
			in:       &dynamodb.AttributeValue{SS: []*string{}},
			actual:   &[]*string{},
			expected: []*string{},
		},
	}

	for i, c := range cases {
		err := Unmarshal(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

func TestInterfaceInput(t *testing.T) {
	var v interface{}
	expected := []interface{}{"abc", "123"}
	err := Unmarshal(&dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{
		{S: aws.String("abc")}, {S: aws.String("123")},
	}}, &v)
	assertConvertTest(t, 0, v, expected, err, nil)
}

func TestUnmarshalError(t *testing.T) {
	cases := []struct {
		in               *dynamodb.AttributeValue
		actual, expected interface{}
		err              error
	}{
		{
			in:       &dynamodb.AttributeValue{},
			actual:   int(0),
			expected: nil,
			err:      &InvalidUnmarshalError{Type: reflect.TypeOf(int(0))},
		},
	}

	for i, c := range cases {
		err := Unmarshal(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

func TestUnmarshalListShared(t *testing.T) {
	for i, c := range sharedListTestCases {
		err := UnmarshalList(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

func TestUnmarshalListError(t *testing.T) {
	cases := []struct {
		in               []*dynamodb.AttributeValue
		actual, expected interface{}
		err              error
	}{
		{
			in:       []*dynamodb.AttributeValue{},
			actual:   []interface{}{},
			expected: nil,
			err:      &InvalidUnmarshalError{Type: reflect.TypeOf([]interface{}{})},
		},
	}

	for i, c := range cases {
		err := UnmarshalList(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

func TestUnmarshalConvertToData(t *testing.T) {
	type T struct {
		Int       int
		Str       string
		ByteSlice []byte
		StrSlice  []string
	}

	exp := T{
		Int:       42,
		Str:       "foo",
		ByteSlice: []byte{42, 97, 83},
		StrSlice:  []string{"cat", "dog"},
	}
	av, err := ConvertToMap(exp)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	var act T
	err = UnmarshalMap(av, &act)
	assertConvertTest(t, 0, act, exp, err, nil)
}

func TestUnmarshalMapShared(t *testing.T) {
	for i, c := range sharedMapTestCases {
		err := UnmarshalMap(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

func TestUnmarshalMapError(t *testing.T) {
	cases := []struct {
		in               map[string]*dynamodb.AttributeValue
		actual, expected interface{}
		err              error
	}{
		{
			in:       map[string]*dynamodb.AttributeValue{},
			actual:   map[string]interface{}{},
			expected: nil,
			err:      &InvalidUnmarshalError{Type: reflect.TypeOf(map[string]interface{}{})},
		},
		{
			in: map[string]*dynamodb.AttributeValue{
				"BOOL": {BOOL: aws.Bool(true)},
			},
			actual:   &map[int]interface{}{},
			expected: nil,
			err:      &UnmarshalTypeError{Value: "map string key", Type: reflect.TypeOf(int(0))},
		},
	}

	for i, c := range cases {
		err := UnmarshalMap(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

func TestUnmarshalListOfMaps(t *testing.T) {
	type testItem struct {
		Value  string
		Value2 int
	}

	cases := []struct {
		in               []map[string]*dynamodb.AttributeValue
		actual, expected interface{}
		err              error
	}{
		{ // Simple map conversion.
			in: []map[string]*dynamodb.AttributeValue{
				{
					"Value": &dynamodb.AttributeValue{
						BOOL: aws.Bool(true),
					},
				},
			},
			actual: &[]map[string]interface{}{},
			expected: []map[string]interface{}{
				{
					"Value": true,
				},
			},
		},
		{ // attribute to struct.
			in: []map[string]*dynamodb.AttributeValue{
				{
					"Value": &dynamodb.AttributeValue{
						S: aws.String("abc"),
					},
					"Value2": &dynamodb.AttributeValue{
						N: aws.String("123"),
					},
				},
			},
			actual: &[]testItem{},
			expected: []testItem{
				{
					Value:  "abc",
					Value2: 123,
				},
			},
		},
	}

	for i, c := range cases {
		err := UnmarshalListOfMaps(c.in, c.actual)
		assertConvertTest(t, i, c.actual, c.expected, err, c.err)
	}
}

type unmarshalUnmarshaler struct {
	Value  string
	Value2 int
	Value3 bool
	Value4 time.Time
}

func (u *unmarshalUnmarshaler) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if av.M == nil {
		return fmt.Errorf("expected AttributeValue to be map")
	}

	if v, ok := av.M["abc"]; !ok {
		return fmt.Errorf("expected `abc` map key")
	} else if v.S == nil {
		return fmt.Errorf("expected `abc` map value string")
	} else {
		u.Value = *v.S
	}

	if v, ok := av.M["def"]; !ok {
		return fmt.Errorf("expected `def` map key")
	} else if v.N == nil {
		return fmt.Errorf("expected `def` map value number")
	} else {
		n, err := strconv.ParseInt(*v.N, 10, 64)
		if err != nil {
			return err
		}
		u.Value2 = int(n)
	}

	if v, ok := av.M["ghi"]; !ok {
		return fmt.Errorf("expected `ghi` map key")
	} else if v.BOOL == nil {
		return fmt.Errorf("expected `ghi` map value number")
	} else {
		u.Value3 = *v.BOOL
	}

	if v, ok := av.M["jkl"]; !ok {
		return fmt.Errorf("expected `jkl` map key")
	} else if v.S == nil {
		return fmt.Errorf("expected `jkl` map value string")
	} else {
		t, err := time.Parse(time.RFC3339, *v.S)
		if err != nil {
			return err
		}
		u.Value4 = t
	}

	return nil
}

func TestUnmarshalUnmashaler(t *testing.T) {
	u := &unmarshalUnmarshaler{}
	av := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"abc": {S: aws.String("value")},
			"def": {N: aws.String("123")},
			"ghi": {BOOL: aws.Bool(true)},
			"jkl": {S: aws.String("2016-05-03T17:06:26.209072Z")},
		},
	}

	err := Unmarshal(av, u)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	if e, a := "value", u.Value; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 123, u.Value2; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := true, u.Value3; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := testDate, u.Value4; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDecodeUseNumber(t *testing.T) {
	u := map[string]interface{}{}
	av := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"abc": {S: aws.String("value")},
			"def": {N: aws.String("123")},
			"ghi": {BOOL: aws.Bool(true)},
		},
	}

	decoder := NewDecoder(func(d *Decoder) {
		d.UseNumber = true
	})
	err := decoder.Decode(av, &u)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	if e, a := "value", u["abc"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	n := u["def"].(Number)
	if e, a := "123", n.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := true, u["ghi"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDecodeUseJsonNumber(t *testing.T) {
	u := map[string]interface{}{}
	av := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"abc": {S: aws.String("value")},
			"def": {N: aws.String("123")},
			"ghi": {BOOL: aws.Bool(true)},
		},
	}

	decoder := NewDecoder(func(d *Decoder) {
		d.UseJsonNumber = true
	})
	err := decoder.Decode(av, &u)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	if e, a := "value", u["abc"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	n := u["def"].(json.Number)
	if e, a := "123", n.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := true, u["ghi"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDecodeUseNumberNumberSet(t *testing.T) {
	u := map[string]interface{}{}
	av := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"ns": {
				NS: []*string{
					aws.String("123"), aws.String("321"),
				},
			},
		},
	}

	decoder := NewDecoder(func(d *Decoder) {
		d.UseNumber = true
	})
	err := decoder.Decode(av, &u)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	ns := u["ns"].([]Number)

	if e, a := "123", ns[0].String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "321", ns[1].String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDecodeEmbeddedPointerStruct(t *testing.T) {
	type B struct {
		Bint int
	}
	type C struct {
		Cint int
	}
	type A struct {
		Aint int
		*B
		*C
	}
	av := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"Aint": {
				N: aws.String("321"),
			},
			"Bint": {
				N: aws.String("123"),
			},
		},
	}
	decoder := NewDecoder()
	a := A{}
	err := decoder.Decode(av, &a)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := 321, a.Aint; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	// Embedded pointer struct can be created automatically.
	if e, a := 123, a.Bint; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	// But not for absent fields.
	if a.C != nil {
		t.Errorf("expect nil, got %v", a.C)
	}
}

func TestDecodeBooleanOverlay(t *testing.T) {
	type BooleanOverlay bool

	av := &dynamodb.AttributeValue{
		BOOL: aws.Bool(true),
	}

	decoder := NewDecoder()

	var v BooleanOverlay

	err := decoder.Decode(av, &v)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := BooleanOverlay(true), v; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDecodeUnixTime(t *testing.T) {
	type A struct {
		Normal time.Time
		Tagged time.Time `dynamodbav:",unixtime"`
		Typed  UnixTime
	}

	expect := A{
		Normal: time.Unix(123, 0).UTC(),
		Tagged: time.Unix(456, 0),
		Typed:  UnixTime(time.Unix(789, 0)),
	}

	input := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"Normal": {
				S: aws.String("1970-01-01T00:02:03Z"),
			},
			"Tagged": {
				N: aws.String("456"),
			},
			"Typed": {
				N: aws.String("789"),
			},
		},
	}
	actual := A{}

	err := Unmarshal(input, &actual)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := expect, actual; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDecodeAliasedUnixTime(t *testing.T) {
	type A struct {
		Normal AliasedTime
		Tagged AliasedTime `dynamodbav:",unixtime"`
	}

	expect := A{
		Normal: AliasedTime(time.Unix(123, 0).UTC()),
		Tagged: AliasedTime(time.Unix(456, 0)),
	}

	input := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"Normal": {
				S: aws.String("1970-01-01T00:02:03Z"),
			},
			"Tagged": {
				N: aws.String("456"),
			},
		},
	}
	actual := A{}

	err := Unmarshal(input, &actual)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if expect != actual {
		t.Errorf("expect %v, got %v", expect, actual)
	}
}

func TestDecoderFieldByIndex(t *testing.T) {
	type (
		Middle struct{ Inner int }
		Outer  struct{ *Middle }
	)
	var outer Outer

	outerType := reflect.TypeOf(outer)
	outerValue := reflect.ValueOf(&outer)
	outerFields := unionStructFields(outerType, MarshalOptions{})
	innerField, _ := outerFields.FieldByName("Inner")

	f := decoderFieldByIndex(outerValue.Elem(), innerField.Index)
	if outer.Middle == nil {
		t.Errorf("expected outer.Middle to be non-nil")
	}
	if f.Kind() != reflect.Int || f.Int() != int64(outer.Inner) {
		t.Error("expected f to be an int with value equal to outer.Inner")
	}
}

func TestDecodeAliasType(t *testing.T) {
	type Str string
	type Int int
	type Uint uint
	type TT struct {
		A Str
		B Int
		C Uint
		S Str
	}

	expect := TT{
		A: "12345",
		B: 12345,
		C: 12345,
		S: "string",
	}
	m := map[string]*dynamodb.AttributeValue{
		"A": {
			N: aws.String("12345"),
		},
		"B": {
			N: aws.String("12345"),
		},
		"C": {
			N: aws.String("12345"),
		},
		"S": {
			S: aws.String("string"),
		},
	}

	var actual TT
	err := UnmarshalMap(m, &actual)
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect:\n%v\nactual:\n%v", expect, actual)
	}
}
