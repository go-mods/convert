package convert

import (
	"encoding/json"
	"fmt"
	"time"
)

// MapStringStringConverter est un type de fonction pour la conversion personnalisée de map[string]string.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]string si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringStringConverter := func(value interface{}) *map[string]string {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringString()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringString(someValue, customMapStringStringConverter)
type MapStringStringConverter func(value interface{}) *map[string]string

// MapStringSliceStringConverter est un type de fonction pour la conversion personnalisée de map[string][]string.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string][]string si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringSliceStringConverter := func(value interface{}) *map[string][]string {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringSliceString()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringSliceString(someValue, customMapStringSliceStringConverter)
type MapStringSliceStringConverter func(value interface{}) *map[string][]string

// MapStringBoolConverter est un type de fonction pour la conversion personnalisée de map[string]bool.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]bool si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringBoolConverter := func(value interface{}) *map[string]bool {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringBool()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringBool(someValue, customMapStringBoolConverter)
type MapStringBoolConverter func(value interface{}) *map[string]bool

// MapStringIntConverter est un type de fonction pour la conversion personnalisée de map[string]int.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]int si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringIntConverter := func(value interface{}) *map[string]int {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringInt(someValue, customMapStringIntConverter)
type MapStringIntConverter func(value interface{}) *map[string]int

// MapStringInt64Converter est un type de fonction pour la conversion personnalisée de map[string]int64.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]int64 si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringInt64Converter := func(value interface{}) *map[string]int64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringInt64(someValue, customMapStringInt64Converter)
type MapStringInt64Converter func(value interface{}) *map[string]int64

// MapStringFloat32Converter est un type de fonction pour la conversion personnalisée de map[string]float32.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]float32 si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringFloat32Converter := func(value interface{}) *map[string]float32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat32()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringFloat32(someValue, customMapStringFloat32Converter)
type MapStringFloat32Converter func(value interface{}) *map[string]float32

// MapStringFloat64Converter est un type de fonction pour la conversion personnalisée de map[string]float64.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]float64 si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringFloat64Converter := func(value interface{}) *map[string]float64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringFloat64(someValue, customMapStringFloat64Converter)
type MapStringFloat64Converter func(value interface{}) *map[string]float64

// MapStringTimeConverter est un type de fonction pour la conversion personnalisée de map[string]time.Time.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]time.Time si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringTimeConverter := func(value interface{}) *map[string]time.Time {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringTime()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringTime(someValue, customMapStringTimeConverter)
type MapStringTimeConverter func(value interface{}) *map[string]time.Time

// MapStringDurationConverter est un type de fonction pour la conversion personnalisée de map[string]time.Duration.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]time.Duration si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringDurationConverter := func(value interface{}) *map[string]time.Duration {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringDuration()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringDuration(someValue, customMapStringDurationConverter)
type MapStringDurationConverter func(value interface{}) *map[string]time.Duration

// MapStringInterfaceConverter est un type de fonction pour la conversion personnalisée de map[string]interface{}.
// Elle prend n'importe quelle valeur et retourne un pointeur vers une map[string]interface{} si la conversion réussit, ou nil si elle échoue.
// Cela permet une logique de conversion flexible et définie par l'utilisateur.
//
// Exemple d'utilisation :
//
//	customMapStringInterfaceConverter := func(value interface{}) *map[string]interface{} {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInterface()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringInterface(someValue, customMapStringInterfaceConverter)
type MapStringInterfaceConverter func(value interface{}) *map[string]interface{}

// ToMapStringString convertit n'importe quel type de valeur en map[string]string.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]string {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringString()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringString(someValue, customConverter)
func ToMapStringString(value interface{}, converters ...MapStringStringConverter) map[string]string {
	res, _ := ToMapStringStringE(value, converters...)
	return res
}

// ToMapStringStringOrDefault convertit n'importe quel type de valeur en map[string]string ou retourne une valeur par défaut.
// Elle prend une valeur de n'importe quel type, une valeur par défaut de type map[string]string, et un nombre variable de convertisseurs personnalisés.
// Si la valeur d'entrée est nil ou si la conversion échoue, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]string {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringString()
//			return &result
//		}
//		return nil
//	}
//
//	defaultValue := map[string]string{"default": "value"}
//	convertedValue := ToMapStringStringOrDefault(someValue, defaultValue, customConverter)
func ToMapStringStringOrDefault(value interface{}, defaultValue map[string]string, converters ...MapStringStringConverter) map[string]string {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringStringE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringStringE converts any type of value to map[string]string or returns an error.
// It handles various types including:
//   - string, []byte
//   - map[string]string
//   - map[string]interface{}
//   - map[interface{}]string
//   - map[interface{}]interface{}
//   - map[string]map[string]map[string]string
//   - map[string]map[string]map[string]interface{}
//   - map[string]map[string]map[string]map[string]string
//   - map[string]map[string]map[string]map[string]interface{}
//
// It uses the following rules to convert:
//   - string, []byte: map[string]string{value: value}
//   - map[string]string: value
//   - map[string]interface{}, map[interface{}]string, map[interface{}]interface{}: convert all keys and values to string
func ToMapStringStringE(value interface{}, converters ...MapStringStringConverter) (map[string]string, error) {
	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case string:
		var res map[string]string
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]string
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case map[string]string:
		return v, nil

	case map[string]interface{}, map[interface{}]string, map[interface{}]interface{}:
		res := make(map[string]string)
		for key, val := range i.(map[interface{}]interface{}) {
			res[ToString(key)] = ToString(val)
		}
		return res, nil

	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// ToMapStringSliceString convertit une valeur en map[string][]string.
// Il prend une valeur de n'importe quel type et des convertisseurs personnalisés optionnels.
// Cette fonction renvoie le résultat de la conversion sans l'erreur.
//
// Exemple d'utilisation :
//
//	convertisseurPersonnalise := func(valeur interface{}) *map[string][]string {
//		if valeurPersonnalisee, ok := valeur.(TypePersonnalise); ok {
//			resultat := valeurPersonnalisee.VersMapStringSliceString()
//			return &resultat
//		}
//		return nil
//	}
//
//	valeurConvertie := ToMapStringSliceString(uneValeur, convertisseurPersonnalise)
func ToMapStringSliceString(value interface{}, converters ...MapStringSliceStringConverter) map[string][]string {
	res, _ := ToMapStringSliceStringE(value, converters...)
	return res
}

// ToMapStringSliceStringOrDefault convertit une valeur en map[string][]string avec une valeur par défaut.
// Il prend une valeur de n'importe quel type, une valeur par défaut et des convertisseurs personnalisés optionnels.
// Si la conversion échoue ou si la valeur d'entrée est nil, la valeur par défaut est renvoyée.
//
// Exemple d'utilisation :
//
//	convertisseurPersonnalise := func(valeur interface{}) *map[string][]string {
//		if valeurPersonnalisee, ok := valeur.(TypePersonnalise); ok {
//			resultat := valeurPersonnalisee.VersMapStringSliceString()
//			return &resultat
//		}
//		return nil
//	}
//
//	valeurParDefaut := map[string][]string{"cle": {"valeur1", "valeur2"}}
//	valeurConvertie := ToMapStringSliceStringOrDefault(uneValeur, valeurParDefaut, convertisseurPersonnalise)
func ToMapStringSliceStringOrDefault(value interface{}, defaultValue map[string][]string, converters ...MapStringSliceStringConverter) map[string][]string {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringSliceStringE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringSliceStringE convertit une valeur en map[string][]string et retourne une erreur si la conversion échoue.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion réussit, elle retourne la map[string][]string résultante et une erreur nil.
// Si la conversion échoue, elle retourne nil et une erreur décrivant le problème.
//
// Exemple d'utilisation :
//
//	convertisseurPersonnalise := func(valeur interface{}) *map[string][]string {
//		if valeurPersonnalisee, ok := valeur.(TypePersonnalise); ok {
//			resultat := valeurPersonnalisee.VersMapStringSliceString()
//			return &resultat
//		}
//		return nil
//	}
//
//	valeurConvertie, err := ToMapStringSliceStringE(uneValeur, convertisseurPersonnalise)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringSliceStringE(value interface{}, converters ...MapStringSliceStringConverter) (map[string][]string, error) {
	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case string:
		var res map[string][]string
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string][]string
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case map[string][]string:
		return v, nil

	case map[string]interface{}, map[interface{}]string, map[interface{}]interface{}:
		res := make(map[string][]string)
		for key, val := range i.(map[interface{}]interface{}) {
			strKey := ToString(key)
			switch sliceVal := val.(type) {
			case []string:
				res[strKey] = sliceVal
			case []interface{}:
				strSlice := make([]string, len(sliceVal))
				for i, v := range sliceVal {
					strSlice[i] = ToString(v)
				}
				res[strKey] = strSlice
			default:
				res[strKey] = []string{ToString(val)}
			}
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringBool convertit n'importe quel type de valeur en map[string]bool.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]bool {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringBool()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringBool(someValue, customConverter)
func ToMapStringBool(value interface{}, converters ...MapStringBoolConverter) map[string]bool {
	res, _ := ToMapStringBoolE(value, converters...)
	return res
}

// ToMapStringBoolOrDefault convertit n'importe quel type de valeur en map[string]bool ou retourne une valeur par défaut.
// Elle prend une valeur de n'importe quel type, une valeur par défaut de type map[string]bool, et un nombre variable de convertisseurs personnalisés.
// Si la valeur d'entrée est nil ou si la conversion échoue, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]bool {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringBool()
//			return &result
//		}
//		return nil
//	}
//
//	defaultValue := map[string]bool{"default": true}
//	convertedValue := ToMapStringBoolOrDefault(someValue, defaultValue, customConverter)
func ToMapStringBoolOrDefault(value interface{}, defaultValue map[string]bool, converters ...MapStringBoolConverter) map[string]bool {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringBoolE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringBoolE convertit n'importe quel type de valeur en map[string]bool.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion réussit, elle retourne la map[string]bool résultante et une erreur nil.
// Si la conversion échoue, elle retourne nil et une erreur décrivant le problème.
//
// Les types pris en charge sont :
//   - map[string]bool
//   - map[string]interface{}
//   - string (JSON)
//   - []byte (JSON)
//
// Pour les autres types, elle retourne une erreur.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]bool {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringBool()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue, err := ToMapStringBoolE(someValue, customConverter)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringBoolE(value interface{}, converters ...MapStringBoolConverter) (map[string]bool, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]bool:
		return v, nil

	case map[string]interface{}:
		res := make(map[string]bool)
		for key, val := range v {
			boolValue, err := ToBoolE(val)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en bool : %v", key, err)
			}
			res[key] = boolValue
		}
		return res, nil

	case string:
		var res map[string]bool
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]bool
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringInt convertit n'importe quel type de valeur en map[string]int.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]int {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringInt(someValue, customConverter)
func ToMapStringInt(value interface{}, converters ...MapStringIntConverter) map[string]int {
	res, _ := ToMapStringIntE(value, converters...)
	return res
}

// ToMapStringIntOrDefault convertit n'importe quel type de valeur en map[string]int ou retourne une valeur par défaut.
// Elle prend une valeur de n'importe quel type, une valeur par défaut de type map[string]int, et un nombre variable de convertisseurs personnalisés.
// Si la valeur d'entrée est nil ou si la conversion échoue, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]int {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt()
//			return &result
//		}
//		return nil
//	}
//
//	defaultValue := map[string]int{"default": 0}
//	convertedValue := ToMapStringIntOrDefault(someValue, defaultValue, customConverter)
func ToMapStringIntOrDefault(value interface{}, defaultValue map[string]int, converters ...MapStringIntConverter) map[string]int {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringIntE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringIntE convertit n'importe quel type de valeur en map[string]int.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion réussit, elle retourne la map[string]int convertie et nil pour l'erreur.
// Si la conversion échoue, elle retourne nil et une erreur décrivant le problème.
//
// Les types pris en charge pour la conversion sont :
// - map[string]int
// - map[interface{}]interface{}
// - map[string]interface{}
// - string (doit être un JSON valide)
// - []byte (doit être un JSON valide)
//
// Pour les types non pris en charge, une erreur sera retournée.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]int {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue, err := ToMapStringIntE(someValue, customConverter)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringIntE(value interface{}, converters ...MapStringIntConverter) (map[string]int, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]int:
		return v, nil

	case map[interface{}]interface{}:
		res := make(map[string]int)
		for k, v := range v {
			key, err := ToStringE(k)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la clé en string : %v", err)
			}
			intValue, err := ToIntE(v)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en int : %v", key, err)
			}
			res[key] = intValue
		}
		return res, nil

	case map[string]interface{}:
		res := make(map[string]int)
		for k, v := range v {
			intValue, err := ToIntE(v)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en int : %v", k, err)
			}
			res[k] = intValue
		}
		return res, nil

	case string:
		var res map[string]int
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]int
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringInt64 convertit n'importe quel type de valeur en map[string]int64.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]int64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringInt64(someValue, customConverter)
func ToMapStringInt64(value interface{}, converters ...MapStringInt64Converter) map[string]int64 {
	res, _ := ToMapStringInt64E(value, converters...)
	return res
}

// ToMapStringInt64OrDefault convertit n'importe quel type de valeur en map[string]int64 ou retourne une valeur par défaut.
// Elle prend une valeur de n'importe quel type, une valeur par défaut de type map[string]int64, et un nombre variable de convertisseurs personnalisés.
// Si la valeur d'entrée est nil ou si la conversion échoue, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]int64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt64()
//			return &result
//		}
//		return nil
//	}
//
//	defaultValue := map[string]int64{"default": 42}
//	convertedValue := ToMapStringInt64OrDefault(someValue, defaultValue, customConverter)
func ToMapStringInt64OrDefault(value interface{}, defaultValue map[string]int64, converters ...MapStringInt64Converter) map[string]int64 {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringInt64E(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringInt64E convertit n'importe quel type de valeur en map[string]int64 et retourne une erreur si la conversion échoue.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion réussit, elle retourne la map[string]int64 résultante et une erreur nil.
// Si la conversion échoue, elle retourne nil et une erreur décrivant le problème.
//
// Les types pris en charge pour la conversion sont :
// - map[string]int64 (retourné tel quel)
// - map[string]interface{} (chaque valeur est convertie en int64)
// - string (supposée être une représentation JSON valide d'une map[string]int64)
// - []byte (supposé être une représentation JSON valide d'une map[string]int64)
//
// Pour les autres types, une erreur "type non pris en charge" est retournée.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]int64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringInt64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue, err := ToMapStringInt64E(someValue, customConverter)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringInt64E(value interface{}, converters ...MapStringInt64Converter) (map[string]int64, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]int64:
		return v, nil

	case map[string]interface{}:
		res := make(map[string]int64)
		for k, v := range v {
			int64Value, err := ToInt64E(v)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en int64 : %v", k, err)
			}
			res[k] = int64Value
		}
		return res, nil

	case string:
		var res map[string]int64
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]int64
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringFloat32 convertit n'importe quel type de valeur en map[string]float32.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]float32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat32()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringFloat32(someValue, customConverter)
func ToMapStringFloat32(value interface{}, converters ...MapStringFloat32Converter) map[string]float32 {
	res, _ := ToMapStringFloat32E(value, converters...)
	return res
}

// ToMapStringFloat32OrDefault convertit n'importe quel type de valeur en map[string]float32 ou retourne une valeur par défaut.
// Elle prend une valeur de n'importe quel type, une valeur par défaut de type map[string]float32, et un nombre variable de convertisseurs personnalisés.
// Si la valeur d'entrée est nil ou si la conversion échoue, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]float32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat32()
//			return &result
//		}
//		return nil
//	}
//
//	defaultValue := map[string]float32{"default": 0.0}
//	convertedValue := ToMapStringFloat32OrDefault(someValue, defaultValue, customConverter)
func ToMapStringFloat32OrDefault(value interface{}, defaultValue map[string]float32, converters ...MapStringFloat32Converter) map[string]float32 {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringFloat32E(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringFloat32E convertit n'importe quel type de valeur en map[string]float32.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion réussit, elle retourne la map[string]float32 convertie et nil pour l'erreur.
// Si la conversion échoue, elle retourne nil et une erreur décrivant le problème.
//
// Les types pris en charge sont :
//   - map[string]float32
//   - map[string]interface{}
//   - string (doit être un JSON valide)
//   - []byte (doit être un JSON valide)
//
// Pour les types non pris en charge, une erreur sera retournée.
//
// Les convertisseurs personnalisés sont appelés dans l'ordre où ils sont fournis.
// Si un convertisseur réussit, sa valeur de retour est utilisée et les autres convertisseurs ne sont pas appelés.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]float32 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat32()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue, err := ToMapStringFloat32E(someValue, customConverter)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringFloat32E(value interface{}, converters ...MapStringFloat32Converter) (map[string]float32, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]float32:
		return v, nil

	case map[string]interface{}:
		res := make(map[string]float32)
		for k, v := range v {
			float32Value, err := ToFloat32E(v)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en float32 : %v", k, err)
			}
			res[k] = float32Value
		}
		return res, nil

	case string:
		var res map[string]float32
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]float32
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringFloat64 convertit n'importe quel type de valeur en map[string]float64.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]float64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringFloat64(someValue, customConverter)
func ToMapStringFloat64(value interface{}, converters ...MapStringFloat64Converter) map[string]float64 {
	res, _ := ToMapStringFloat64E(value, converters...)
	return res
}

// ToMapStringFloat64OrDefault convertit n'importe quel type de valeur en map[string]float64 ou retourne une valeur par défaut.
// Elle prend une valeur de n'importe quel type, une valeur par défaut de type map[string]float64, et un nombre variable de convertisseurs personnalisés.
// Si la valeur d'entrée est nil ou si la conversion échoue, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]float64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat64()
//			return &result
//		}
//		return nil
//	}
//
//	defaultValue := map[string]float64{"default": 0.0}
//	convertedValue := ToMapStringFloat64OrDefault(someValue, defaultValue, customConverter)
func ToMapStringFloat64OrDefault(value interface{}, defaultValue map[string]float64, converters ...MapStringFloat64Converter) map[string]float64 {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringFloat64E(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringFloat64E convertit n'importe quel type de valeur en map[string]float64.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion réussit, elle retourne la map[string]float64 résultante et une erreur nil.
// Si la conversion échoue, elle retourne nil et une erreur décrivant le problème.
//
// Les types pris en charge sont :
//   - map[string]float64 (retourné tel quel)
//   - map[string]interface{} (chaque valeur est convertie en float64)
//   - string (interprétée comme JSON)
//   - []byte (interprété comme JSON)
//
// Pour les types non pris en charge, une erreur est retournée.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]float64 {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringFloat64()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue, err := ToMapStringFloat64E(someValue, customConverter)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringFloat64E(value interface{}, converters ...MapStringFloat64Converter) (map[string]float64, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]float64:
		return v, nil

	case map[string]interface{}:
		res := make(map[string]float64)
		for k, v := range v {
			float64Value, err := ToFloat64E(v)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en float64 : %v", k, err)
			}
			res[k] = float64Value
		}
		return res, nil

	case string:
		var res map[string]float64
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]float64
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringTime convertit n'importe quel type de valeur en map[string]time.Time.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]time.Time {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringTime()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringTime(someValue, customConverter)
func ToMapStringTime(value interface{}, converters ...MapStringTimeConverter) map[string]time.Time {
	res, _ := ToMapStringTimeE(value, converters...)
	return res
}

// ToMapStringTimeOrDefault convertit n'importe quel type de valeur en map[string]time.Time ou retourne une valeur par défaut.
// Elle prend une valeur de n'importe quel type, une valeur par défaut de type map[string]time.Time, et un nombre variable de convertisseurs personnalisés.
// Si la valeur d'entrée est nil ou si la conversion échoue, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]time.Time {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringTime()
//			return &result
//		}
//		return nil
//	}
//
//	defaultValue := map[string]time.Time{"default": time.Now()}
//	convertedValue := ToMapStringTimeOrDefault(someValue, defaultValue, customConverter)
func ToMapStringTimeOrDefault(value interface{}, defaultValue map[string]time.Time, converters ...MapStringTimeConverter) map[string]time.Time {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringTimeE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringTimeE convertit n'importe quel type de valeur en map[string]time.Time.
// Elle prend une valeur de n'importe quel type et un nombre variable de convertisseurs personnalisés.
// Si la conversion réussit, elle retourne la map[string]time.Time résultante et une erreur nil.
// Si la conversion échoue, elle retourne nil et une erreur décrivant le problème.
//
// Les types pris en charge pour la conversion sont :
// - map[string]time.Time : retourné tel quel
// - map[string]interface{} : chaque valeur est convertie en time.Time
// - string : interprété comme JSON et désérialisé en map[string]time.Time
// - []byte : interprété comme JSON et désérialisé en map[string]time.Time
//
// Pour les types non pris en charge, une erreur est retournée.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]time.Time {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringTime()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue, err := ToMapStringTimeE(someValue, customConverter)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringTimeE(value interface{}, converters ...MapStringTimeConverter) (map[string]time.Time, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]time.Time:
		return v, nil

	case map[string]interface{}:
		res := make(map[string]time.Time)
		for k, v := range v {
			timeValue, err := ToTimeE(v)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en time.Time : %v", k, err)
			}
			res[k] = timeValue
		}
		return res, nil

	case string:
		var res map[string]time.Time
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]time.Time
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringDuration convertit une valeur en map[string]time.Duration.
// Elle prend n'importe quelle valeur et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue, elle retourne une map vide.
//
// Exemple d'utilisation :
//
//	customConverter := func(value interface{}) *map[string]time.Duration {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringDuration()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringDuration(someValue, customConverter)
func ToMapStringDuration(value interface{}, converters ...MapStringDurationConverter) map[string]time.Duration {
	res, _ := ToMapStringDurationE(value, converters...)
	return res
}

// ToMapStringDurationOrDefault convertit une valeur en map[string]time.Duration avec une valeur par défaut.
// Elle prend une valeur, une valeur par défaut, et un nombre variable de convertisseurs personnalisés.
// Si la conversion échoue ou si la valeur d'entrée est nil, elle retourne la valeur par défaut.
//
// Exemple d'utilisation :
//
//	defaultValue := map[string]time.Duration{"key": 5 * time.Second}
//	customConverter := func(value interface{}) *map[string]time.Duration {
//		if customVal, ok := value.(CustomType); ok {
//			result := customVal.ToMapStringDuration()
//			return &result
//		}
//		return nil
//	}
//
//	convertedValue := ToMapStringDurationOrDefault(someValue, defaultValue, customConverter)
func ToMapStringDurationOrDefault(value interface{}, defaultValue map[string]time.Duration, converters ...MapStringDurationConverter) map[string]time.Duration {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringDurationE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringDurationE convertit une valeur en map[string]time.Duration.
// Elle prend n'importe quelle valeur et un nombre variable de convertisseurs personnalisés.
// Elle retourne la map[string]time.Duration convertie et une erreur si la conversion échoue.
func ToMapStringDurationE(value interface{}, converters ...MapStringDurationConverter) (map[string]time.Duration, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]time.Duration:
		return v, nil

	case map[string]interface{}:
		res := make(map[string]time.Duration)
		for k, v := range v {
			durationValue, err := ToDurationE(v)
			if err != nil {
				return nil, fmt.Errorf("impossible de convertir la valeur pour la clé '%s' en time.Duration : %v", k, err)
			}
			res[k] = durationValue
		}
		return res, nil

	case string:
		var res map[string]time.Duration
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]time.Duration
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}

// ToMapStringInterface est une fonction de conversion pour map[string]interface{}.
// Elle prend n'importe quelle valeur et un nombre variable de convertisseurs personnalisés.
// Elle retourne la map[string]interface{} convertie.
//
// Exemple d'utilisation :
//
//	convertisseurPersonnalise := func(valeur interface{}) *map[string]interface{} {
//		if valeurPersonnalisee, ok := valeur.(TypePersonnalise); ok {
//			resultat := valeurPersonnalisee.VersMapStringInterface()
//			return &resultat
//		}
//		return nil
//	}
//
//	valeurConvertie := ToMapStringInterface(uneValeur, convertisseurPersonnalise)
func ToMapStringInterface(value interface{}, converters ...MapStringInterfaceConverter) map[string]interface{} {
	res, _ := ToMapStringInterfaceE(value, converters...)
	return res
}

// ToMapStringInterfaceOrDefault est une fonction de conversion pour map[string]interface{} avec une valeur par défaut.
// Elle prend une valeur quelconque, une valeur par défaut de type map[string]interface{}, et un nombre variable de convertisseurs personnalisés.
// Elle retourne la map[string]interface{} convertie ou la valeur par défaut si la conversion échoue.
//
// Exemple d'utilisation :
//
//	valeurParDefaut := map[string]interface{}{"cle": "valeur"}
//	convertisseurPersonnalise := func(valeur interface{}) *map[string]interface{} {
//		if valeurPersonnalisee, ok := valeur.(TypePersonnalise); ok {
//			resultat := valeurPersonnalisee.VersMapStringInterface()
//			return &resultat
//		}
//		return nil
//	}
//
//	valeurConvertie := ToMapStringInterfaceOrDefault(uneValeur, valeurParDefaut, convertisseurPersonnalise)
func ToMapStringInterfaceOrDefault(value interface{}, defaultValue map[string]interface{}, converters ...MapStringInterfaceConverter) map[string]interface{} {
	if value == nil {
		return defaultValue
	}
	res, _ := ToMapStringInterfaceE(value, converters...)
	if res == nil {
		return defaultValue
	}
	return res
}

// ToMapStringInterfaceE est une fonction de conversion pour map[string]interface{} avec gestion d'erreur.
// Elle prend n'importe quelle valeur et un nombre variable de convertisseurs personnalisés.
// Elle retourne la map[string]interface{} convertie et une erreur si la conversion échoue.
//
// Exemple d'utilisation :
//
//	convertisseurPersonnalise := func(valeur interface{}) *map[string]interface{} {
//		if valeurPersonnalisee, ok := valeur.(TypePersonnalise); ok {
//			resultat := valeurPersonnalisee.VersMapStringInterface()
//			return &resultat
//		}
//		return nil
//	}
//
//	valeurConvertie, err := ToMapStringInterfaceE(uneValeur, convertisseurPersonnalise)
//	if err != nil {
//		// Gérer l'erreur
//	}
func ToMapStringInterfaceE(value interface{}, converters ...MapStringInterfaceConverter) (map[string]interface{}, error) {
	if value == nil {
		return nil, nil
	}

	i := Indirect(value)

	for _, converter := range converters {
		if result := converter(value); result != nil {
			return *result, nil
		}
	}

	switch v := i.(type) {
	case map[string]interface{}:
		return v, nil

	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, v := range v {
			res[ToString(k)] = v
		}
		return res, nil

	case string:
		var res map[string]interface{}
		err := json.Unmarshal([]byte(v), &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	case []byte:
		var res map[string]interface{}
		err := json.Unmarshal(v, &res)
		if err != nil {
			return nil, err
		}
		return res, nil

	default:
		return nil, fmt.Errorf("type non pris en charge : %T", value)
	}
}
