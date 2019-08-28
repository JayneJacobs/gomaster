package hydraconfigurator

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ConfigFields map[string]reflect.Value

// Add function takes name and value strings detects the type from
// the reflect package and populates a map .
func (f ConfigFields) Add(name, value, t string) error {
	switch t {
	case "STRING":
		f[name] = reflect.ValueOf(value)
	case "INTEGER":
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(i)
	case "FLOAT":
		fl, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(fl)
	case "BOOL":
		b, err := strconv.ParseBool(value)

		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(b)
	}
	return nil
}

// MarshalCustomConfig takes a value from a reflect.Value type and a filename
// it uses these to populate a ConfigFields map and returns an error.
func MarshallCustomConfig(v reflect.Value, filename string) error {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("A panic orrured", r)
		}
	}()
	if !v.CanSet() {
		return errors.New("Value passed not settable")
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fields := make(ConfigFields) // make map[string].reflect.Value
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Processing line", line)
		if strings.Count(line, "|") != 1 || strings.Count(line, ";") != 1 {
			continue //ends the loop
		}
		args := strings.Split(line, "|")
		valuetype := strings.Split(args[1], ";")
		name, value, vtype := strings.TrimSpace(args[0]), strings.TrimSpace(valuetype[0]), strings.ToUpper(strings.TrimSpace(valuetype[1]))
		fields.Add(name, value, vtype)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	vt := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldType := vt.Field(i)
		fieldValue := v.Field(i)

		name := fieldType.Tag.Get("name")
		if name == "" {
			name = fieldType.Name
		}

		if v, ok := fields[name]; ok {
			fieldValue.Set(v)
		}
	}
	return nil
}
