package sav

import (
	"fmt"
	"log"
	"reflect"
	"services/io/spss"
	"strconv"
)

type Writer interface {
	Write(rows interface{}) error
}

// Example buffered implementation
type BufferOutput struct {
	inputType string
}

// Example buffered implementation
func (b BufferOutput) Write(rows interface{}) error {
	return nil
}

type FileOutput struct {
	inputType string
}

func (f FileOutput) Write(rows interface{}) error {

	inValue, inType := spss.GetConcreteReflectValueAndType(rows) // Get the concrete type (not pointer) (Slice<?> or Array<?>)
	if err := ensureInType(inType); err != nil {
		return err
	}

	inInnerWasPointer, inInnerType := spss.GetConcreteContainerInnerType(inType) // Get the concrete inner type (not pointer) (Container<"?">)
	if err := ensureInInnerType(inInnerType); err != nil {
		return err
	}

	inInnerStructInfo := spss.GetStructInfo(inInnerType) // Get the inner struct info to get SPSS annotations
	var header []Header
	var data []DataItem

	for _, fieldInfo := range inInnerStructInfo.Fields { // Used to write metadata rows SPSS

		var spssType spss.ColumnType = 0

		switch fieldInfo.FieldType {
		case reflect.String:
			spssType = spss.ReadstatTypeString
		case reflect.Int8, reflect.Uint8:
			spssType = spss.ReadstatTypeInt8
		case reflect.Int, reflect.Int32, reflect.Uint32:
			spssType = spss.ReadstatTypeInt32
		case reflect.Float32:
			spssType = spss.ReadstatTypeFloat
		case reflect.Float64:
			spssType = spss.ReadstatTypeDouble
		default:
			return fmt.Errorf("cannot convert type for struct variable %s into SPSS type", fieldInfo.Keys[0])
		}

		header = append(header, Header{spssType, fieldInfo.Keys[0], fieldInfo.Keys[0]})
	}

	if inValue.Kind() != reflect.Slice {
		panic("You need to pass a slice of interface{} to save to an SPSS SAV file")
	}

	inLen := inValue.Len()
	for i := 0; i < inLen; i++ { // Iterate over container rows
		var dataItem []interface{}
		for j, fieldInfo := range inInnerStructInfo.Fields {
			header[j].Label = ""
			inInnerFieldValue, err := getInnerField(inValue.Index(i), inInnerWasPointer, fieldInfo.IndexChain) // Get the correct field header <-> position
			if err != nil {
				return err
			}
			// convert to correct type
			var spssType interface{}

			switch fieldInfo.FieldType {
			case reflect.String:
				spssType = inInnerFieldValue
			case reflect.Int8, reflect.Uint8:
				spssType, _ = strconv.Atoi(inInnerFieldValue)
			case reflect.Int, reflect.Int32, reflect.Uint32:
				spssType, _ = strconv.Atoi(inInnerFieldValue)
			case reflect.Float32:
				spssType, _ = strconv.ParseFloat(inInnerFieldValue, 32)
			case reflect.Float64:
				spssType, _ = strconv.ParseFloat(inInnerFieldValue, 64)
			default:
				return fmt.Errorf("cannot convert value for struct variable %s into SPSS type", fieldInfo.Keys[0])
			}

			dataItem = append(dataItem, spssType)
		}
		data = append(data, DataItem{dataItem})

	}

	val := Export(f.inputType, "SAV from GO", header, data)

	if val != 0 {
		return fmt.Errorf("cannot open or write to file: %s", f.inputType)
	}

	log.Printf("Finished writing to: %s, return value: %d", f.inputType, val)

	return nil
}

// Check if the inType is an array or a slice
func ensureInType(outType reflect.Type) error {
	switch outType.Kind() {
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		return nil
	}
	return fmt.Errorf("cannot use " + outType.String() + ", only slice or array supported")
}

// Check if the inInnerType is of type struct
func ensureInInnerType(outInnerType reflect.Type) error {
	switch outInnerType.Kind() {
	case reflect.Struct:
		return nil
	}
	return fmt.Errorf("cannot use " + outInnerType.String() + ", only struct supported")
}

func getInnerField(outInner reflect.Value, outInnerWasPointer bool, index []int) (string, error) {
	oi := outInner
	if outInnerWasPointer {
		if oi.IsNil() {
			return "", nil
		}
		oi = outInner.Elem()
	}
	// because pointers can be nil need to recurse one index at a time and perform nil check
	if len(index) > 1 {
		nextField := oi.Field(index[0])
		return getInnerField(nextField, nextField.Kind() == reflect.Ptr, index[1:])
	}
	return spss.GetFieldAsString(oi.FieldByIndex(index))
}
