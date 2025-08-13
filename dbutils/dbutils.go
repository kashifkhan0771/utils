package dbutils

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
)

func Get(db *sql.DB, dest interface{}, query string, args ...interface{}) error {
	val := reflect.ValueOf(dest)

	if val.Kind() != reflect.Pointer || val.Elem().Kind() != reflect.Struct {
		return errors.New("destination must be a pointer to a struct")
	}

	val = val.Elem()
	t := val.Type()

	fieldMap := make(map[string]interface{})
	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		// Check if the field is unexported
		if field.PkgPath != "" {
			continue
		}

		tag := field.Tag.Get("json")

		if tag != "" {
			// Handle cases like json:"name, omitempty"
			tagName := strings.Split(tag, ",")[0]
			fieldMap[tagName] = val.Field(i).Addr().Interface()
		}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			return nil
		}

		scanArgs := make([]interface{}, len(columns))
		for i, colName := range columns {
			if destField, ok := fieldMap[colName]; ok {
				scanArgs[i] = destField
			} else {
				var dummy interface{}
				scanArgs[i] = &dummy
			}
		}

		return rows.Scan(scanArgs...)
	}

	return sql.ErrNoRows
}
