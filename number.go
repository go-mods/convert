package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToInt converts any type of values to int
func ToInt(value interface{}) (res int, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
				if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to int failed", value.(string))
						} else {
							res, err = ToInt(resBool)
						}
					} else {
						res, err = ToInt(resF64)
					}
				} else {
					res, err = ToInt(resU64)
				}

			} else {
				res = int(res64)
			}
		}
	case int:
		{
			res = value.(int)
		}
	case int8:
		{
			res = int(value.(int8))
		}
	case int16:
		{
			res = int(value.(int16))
		}
	case int32:
		{
			res = int(value.(int32))
		}
	case int64:
		{
			res = int(value.(int64))
		}
	case uint:
		{
			res = int(value.(uint))
		}
	case uint8:
		{
			res = int(value.(uint8))
		}
	case uint16:
		{
			res = int(value.(uint16))
		}
	case uint32:
		{
			res = int(value.(uint32))
		}
	case uint64:
		{
			res = int(value.(uint64))
		}
	case uintptr:
		{
			res = int(value.(uintptr))
		}
	case float32:
		{
			res = int(value.(float32))
		}
	case float64:
		{
			res = int(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToInt(valueStr)
		}
	}
	return
}

// ToInt8 converts any type of values to int8
func ToInt8(value interface{}) (res int8, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
				if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to int8 failed", value.(string))
						} else {
							res, err = ToInt8(resBool)
						}
					} else {
						res, err = ToInt8(resF64)
					}
				} else {
					res, err = ToInt8(resU64)
				}

			} else {
				res = int8(res64)
			}
		}
	case int:
		{
			res = int8(value.(int))
		}
	case int8:
		{
			res = value.(int8)
		}
	case int16:
		{
			res = int8(value.(int16))
		}
	case int32:
		{
			res = int8(value.(int32))
		}
	case int64:
		{
			res = int8(value.(int64))
		}
	case uint:
		{
			res = int8(value.(uint))
		}
	case uint8:
		{
			res = int8(value.(uint8))
		}
	case uint16:
		{
			res = int8(value.(uint16))
		}
	case uint32:
		{
			res = int8(value.(uint32))
		}
	case uint64:
		{
			res = int8(value.(uint64))
		}
	case uintptr:
		{
			res = int8(value.(uintptr))
		}
	case float32:
		{
			res = int8(value.(float32))
		}
	case float64:
		{
			res = int8(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToInt8(valueStr)
		}
	}
	return
}

// ToInt16 converts any type of values to int16
func ToInt16(value interface{}) (res int16, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
				if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to int16 failed", value.(string))
						} else {
							res, err = ToInt16(resBool)
						}
					} else {
						res, err = ToInt16(resF64)
					}
				} else {
					res, err = ToInt16(resU64)
				}

			} else {
				res = int16(res64)
			}
		}
	case int:
		{
			res = int16(value.(int))
		}
	case int8:
		{
			res = int16(value.(int8))
		}
	case int16:
		{
			res = value.(int16)
		}
	case int32:
		{
			res = int16(value.(int32))
		}
	case int64:
		{
			res = int16(value.(int64))
		}
	case uint:
		{
			res = int16(value.(uint))
		}
	case uint8:
		{
			res = int16(value.(uint8))
		}
	case uint16:
		{
			res = int16(value.(uint16))
		}
	case uint32:
		{
			res = int16(value.(uint32))
		}
	case uint64:
		{
			res = int16(value.(uint64))
		}
	case uintptr:
		{
			res = int16(value.(uintptr))
		}
	case float32:
		{
			res = int16(value.(float32))
		}
	case float64:
		{
			res = int16(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToInt16(valueStr)
		}
	}
	return
}

// ToInt32 converts any type of values to int32
func ToInt32(value interface{}) (res int32, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
				if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to int32 failed", value.(string))
						} else {
							res, err = ToInt32(resBool)
						}
					} else {
						res, err = ToInt32(resF64)
					}
				} else {
					res, err = ToInt32(resU64)
				}

			} else {
				res = int32(res64)
			}
		}
	case int:
		{
			res = int32(value.(int))
		}
	case int8:
		{
			res = int32(value.(int8))
		}
	case int16:
		{
			res = int32(value.(int16))
		}
	case int32:
		{
			res = value.(int32)
		}
	case int64:
		{
			res = int32(value.(int64))
		}
	case uint:
		{
			res = int32(value.(uint))
		}
	case uint8:
		{
			res = int32(value.(uint8))
		}
	case uint16:
		{
			res = int32(value.(uint16))
		}
	case uint32:
		{
			res = int32(value.(uint32))
		}
	case uint64:
		{
			res = int32(value.(uint64))
		}
	case uintptr:
		{
			res = int32(value.(uintptr))
		}
	case float32:
		{
			res = int32(value.(float32))
		}
	case float64:
		{
			res = int32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToInt32(valueStr)
		}
	}
	return
}

// ToInt64 converts any type of values to int64
func ToInt64(value interface{}) (res int64, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
				if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to int64 failed", value.(string))
						} else {
							res, err = ToInt64(resBool)
						}
					} else {
						res, err = ToInt64(resF64)
					}
				} else {
					res, err = ToInt64(resU64)
				}

			} else {
				res = res64
			}
		}
	case int:
		{
			res = int64(value.(int))
		}
	case int8:
		{
			res = int64(value.(int8))
		}
	case int16:
		{
			res = int64(value.(int16))
		}
	case int32:
		{
			res = value.(int64)
		}
	case uint:
		{
			res = int64(value.(uint))
		}
	case uint8:
		{
			res = int64(value.(uint8))
		}
	case uint16:
		{
			res = int64(value.(uint16))
		}
	case uint32:
		{
			res = int64(value.(uint32))
		}
	case uint64:
		{
			res = int64(value.(uint64))
		}
	case uintptr:
		{
			res = int64(value.(uintptr))
		}
	case float32:
		{
			res = int64(value.(float32))
		}
	case float64:
		{
			res = int64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToInt64(valueStr)
		}
	}
	return
}

// ToUint converts any type of values to uint
func ToUint(value interface{}) (res uint, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
				if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to uint failed", value.(string))
						} else {
							res, err = ToUint(resBool)
						}
					} else {
						res, err = ToUint(resF64)
					}
				} else {
					res, err = ToUint(res64)
				}
			} else {
				res = uint(resU64)
			}
		}
	case int:
		{
			res = uint(value.(int))
		}
	case int8:
		{
			res = uint(value.(int8))
		}
	case int16:
		{
			res = uint(value.(int16))
		}
	case int32:
		{
			res = uint(value.(int32))
		}
	case int64:
		{
			res = uint(value.(int64))
		}
	case uint:
		{
			res = value.(uint)
		}
	case uint8:
		{
			res = uint(value.(uint8))
		}
	case uint16:
		{
			res = uint(value.(uint16))
		}
	case uint32:
		{
			res = uint(value.(uint32))
		}
	case uint64:
		{
			res = uint(value.(uint64))
		}
	case uintptr:
		{
			res = uint(value.(uintptr))
		}
	case float32:
		{
			res = uint(value.(float32))
		}
	case float64:
		{
			res = uint(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToUint(valueStr)
		}
	}
	return
}

// ToUint8 converts any type of values to uint8
func ToUint8(value interface{}) (res uint8, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
				if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to uint8 failed", value.(string))
						} else {
							res, err = ToUint8(resBool)
						}
					} else {
						res, err = ToUint8(resF64)
					}
				} else {
					res, err = ToUint8(res64)
				}
			} else {
				res = uint8(resU64)
			}
		}
	case int:
		{
			res = uint8(value.(int))
		}
	case int8:
		{
			res = uint8(value.(int8))
		}
	case int16:
		{
			res = uint8(value.(int16))
		}
	case int32:
		{
			res = uint8(value.(int32))
		}
	case int64:
		{
			res = uint8(value.(int64))
		}
	case uint:
		{
			res = uint8(value.(uint))
		}
	case uint8:
		{
			res = value.(uint8)
		}
	case uint16:
		{
			res = uint8(value.(uint16))
		}
	case uint32:
		{
			res = uint8(value.(uint32))
		}
	case uint64:
		{
			res = uint8(value.(uint64))
		}
	case uintptr:
		{
			res = uint8(value.(uintptr))
		}
	case float32:
		{
			res = uint8(value.(float32))
		}
	case float64:
		{
			res = uint8(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToUint8(valueStr)
		}
	}
	return
}

// ToUint16 converts any type of values to uint16
func ToUint16(value interface{}) (res uint16, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
				if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to uint16 failed", value.(string))
						} else {
							res, err = ToUint16(resBool)
						}
					} else {
						res, err = ToUint16(resF64)
					}
				} else {
					res, err = ToUint16(res64)
				}
			} else {
				res = uint16(resU64)
			}
		}
	case int:
		{
			res = uint16(value.(int))
		}
	case int8:
		{
			res = uint16(value.(int8))
		}
	case int16:
		{
			res = uint16(value.(int16))
		}
	case int32:
		{
			res = uint16(value.(int32))
		}
	case int64:
		{
			res = uint16(value.(int64))
		}
	case uint:
		{
			res = uint16(value.(uint))
		}
	case uint8:
		{
			res = uint16(value.(uint8))
		}
	case uint16:
		{
			res = value.(uint16)
		}
	case uint32:
		{
			res = uint16(value.(uint32))
		}
	case uint64:
		{
			res = uint16(value.(uint64))
		}
	case uintptr:
		{
			res = uint16(value.(uintptr))
		}
	case float32:
		{
			res = uint16(value.(float32))
		}
	case float64:
		{
			res = uint16(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToUint16(valueStr)
		}
	}
	return
}

// ToUint32 converts any type of values to uint32
func ToUint32(value interface{}) (res uint32, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
				if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to uint32 failed", value.(string))
						} else {
							res, err = ToUint32(resBool)
						}
					} else {
						res, err = ToUint32(resF64)
					}
				} else {
					res, err = ToUint32(res64)
				}
			} else {
				res = uint32(resU64)
			}
		}
	case int:
		{
			res = uint32(value.(int))
		}
	case int8:
		{
			res = uint32(value.(int8))
		}
	case int16:
		{
			res = uint32(value.(int16))
		}
	case int32:
		{
			res = uint32(value.(int32))
		}
	case int64:
		{
			res = uint32(value.(int64))
		}
	case uint:
		{
			res = uint32(value.(uint))
		}
	case uint8:
		{
			res = uint32(value.(uint8))
		}
	case uint16:
		{
			res = uint32(value.(uint16))
		}
	case uint32:
		{
			res = value.(uint32)
		}
	case uint64:
		{
			res = uint32(value.(uint64))
		}
	case uintptr:
		{
			res = uint32(value.(uintptr))
		}
	case float32:
		{
			res = uint32(value.(float32))
		}
	case float64:
		{
			res = uint32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToUint32(valueStr)
		}
	}
	return
}

// ToUint64 converts any type of values to uint64
func ToUint64(value interface{}) (res uint64, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
				if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {
					if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to uint64 failed", value.(string))
						} else {
							res, err = ToUint64(resBool)
						}
					} else {
						res, err = ToUint64(resF64)
					}
				} else {
					res, err = ToUint64(res64)
				}
			} else {
				res = resU64
			}
		}
	case int:
		{
			res = uint64(value.(int))
		}
	case int8:
		{
			res = uint64(value.(int8))
		}
	case int16:
		{
			res = uint64(value.(int16))
		}
	case int32:
		{
			res = uint64(value.(int32))
		}
	case int64:
		{
			res = uint64(value.(int64))
		}
	case uint:
		{
			res = uint64(value.(uint))
		}
	case uint8:
		{
			res = uint64(value.(uint8))
		}
	case uint16:
		{
			res = uint64(value.(uint16))
		}
	case uint32:
		{
			res = uint64(value.(uint32))
		}
	case uint64:
		{
			res = value.(uint64)
		}
	case uintptr:
		{
			res = uint64(value.(uintptr))
		}
	case float32:
		{
			res = uint64(value.(float32))
		}
	case float64:
		{
			res = uint64(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToUint64(valueStr)
		}
	}
	return
}

// ToFloat32 converts any type of values to float32
func ToFloat32(value interface{}) (res float32, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
				if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
					if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {

						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to float32 failed", value.(string))
						} else {
							res, err = ToFloat32(resBool)
						}
					} else {
						res, err = ToFloat32(res64)
					}
				} else {
					res, err = ToFloat32(resU64)
				}
			} else {
				res, err = ToFloat32(resF64)
			}
		}
	case int:
		{
			res = float32(value.(int))
		}
	case int8:
		{
			res = float32(value.(int8))
		}
	case int16:
		{
			res = float32(value.(int16))
		}
	case int32:
		{
			res = float32(value.(int32))
		}
	case int64:
		{
			res = float32(value.(int64))
		}
	case uint:
		{
			res = float32(value.(uint))
		}
	case uint8:
		{
			res = float32(value.(uint8))
		}
	case uint16:
		{
			res = float32(value.(uint16))
		}
	case uint32:
		{
			res = float32(value.(uint32))
		}
	case uint64:
		{
			res = float32(value.(uint64))
		}
	case uintptr:
		{
			res = float32(value.(uintptr))
		}
	case float32:
		{
			res = value.(float32)
		}
	case float64:
		{
			res = float32(value.(float64))
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToFloat32(valueStr)
		}
	}
	return
}

// ToFloat64 converts any type of values to float64
func ToFloat64(value interface{}) (res float64, err error) {
	switch v := value.(type) {
	case string:
		{
			valueString := value.(string)
			if len(valueString) == 0 {
				return 0, nil
			}
			if resF64, e := strconv.ParseFloat(valueString, 64); e != nil {
				if resU64, e := strconv.ParseUint(valueString, 0, 64); e != nil {
					if res64, e := strconv.ParseInt(valueString, 0, 64); e != nil {

						if resBool, e := strconv.ParseBool(valueString); e != nil {
							err = fmt.Errorf("convert: string \"%s\" to float64 failed", value.(string))
						} else {
							res, err = ToFloat64(resBool)
						}
					} else {
						res, err = ToFloat64(res64)
					}
				} else {
					res, err = ToFloat64(resU64)
				}
			} else {
				res = resF64
			}
		}
	case int:
		{
			res = float64(value.(int))
		}
	case int8:
		{
			res = float64(value.(int8))
		}
	case int16:
		{
			res = float64(value.(int16))
		}
	case int32:
		{
			res = float64(value.(int32))
		}
	case int64:
		{
			res = float64(value.(int64))
		}
	case uint:
		{
			res = float64(value.(uint))
		}
	case uint8:
		{
			res = float64(value.(uint8))
		}
	case uint16:
		{
			res = float64(value.(uint16))
		}
	case uint32:
		{
			res = float64(value.(uint32))
		}
	case uint64:
		{
			res = float64(value.(uint64))
		}
	case uintptr:
		{
			res = float64(value.(uintptr))
		}
	case float32:
		{
			res = float64(value.(float32))
		}
	case float64:
		{
			res = value.(float64)
		}
	case bool:
		{
			if value.(bool) {
				res = 1
			} else {
				res = 0
			}
		}
	default:
		{
			valueStr := ToValidString(v)
			res, err = ToFloat64(valueStr)
		}
	}
	return
}

func ToIntDef(value interface{}, def int) int {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToInt(value)
	if err != nil {
		return def
	}
	return res
}

func ToInt8Def(value interface{}, def int8) int8 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToInt8(value)
	if err != nil {
		return def
	}
	return res
}

func ToInt16Def(value interface{}, def int16) int16 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToInt16(value)
	if err != nil {
		return def
	}
	return res
}

func ToInt32Def(value interface{}, def int32) int32 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToInt32(value)
	if err != nil {
		return def
	}
	return res
}

func ToInt64Def(value interface{}, def int64) int64 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToInt64(value)
	if err != nil {
		return def
	}
	return res
}

func ToUintDef(value interface{}, def uint) uint {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToUint(value)
	if err != nil {
		return def
	}
	return res
}

func ToUint8Def(value interface{}, def uint8) uint8 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToUint8(value)
	if err != nil {
		return def
	}
	return res
}

func ToUint16Def(value interface{}, def uint16) uint16 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToUint16(value)
	if err != nil {
		return def
	}
	return res
}

func ToUint32Def(value interface{}, def uint32) uint32 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToUint32(value)
	if err != nil {
		return def
	}
	return res
}

func ToUint64Def(value interface{}, def uint64) uint64 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToUint64(value)
	if err != nil {
		return def
	}
	return res
}

func ToFloat32Def(value interface{}, def float32) float32 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToFloat32(value)
	if err != nil {
		return def
	}
	return res
}

func ToFloat64Def(value interface{}, def float64) float64 {
	// return th default value if value is nil
	if value == nil {
		return def
	}

	// return th default value if value is of type string and its length is zero
	if v, ok := value.(string); ok {
		if len(v) == 0 {
			return def
		}
	}

	// return the default value if value is of type []byte and its length is zero
	if v, ok := value.([]byte); ok {
		if len(v) == 0 {
			return def
		}
	}

	res, err := ToFloat64(value)
	if err != nil {
		return def
	}
	return res
}

// ToIntOrPanic converts any type of value to int or call the panic function
func ToIntOrPanic(value interface{}) (res int) {
	if res, err := ToInt(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt8OrPanic converts any type of value to int8 or call the panic function
func ToInt8OrPanic(value interface{}) (res int8) {
	if res, err := ToInt8(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt16OrPanic converts any type of value to int16 or call the panic function
func ToInt16OrPanic(value interface{}) (res int16) {
	if res, err := ToInt16(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt32OrPanic converts any type of value to int32 or call the panic function
func ToInt32OrPanic(value interface{}) (res int32) {
	if res, err := ToInt32(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt64OrPanic converts any type of value to int64 or call the panic function
func ToInt64OrPanic(value interface{}) (res int64) {
	if res, err := ToInt64(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUintOrPanic converts any type of value to uint or call the panic function
func ToUintOrPanic(value interface{}) (res uint) {
	if res, err := ToUint(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUint8OrPanic converts any type of value to uint8 or call the panic function
func ToUint8OrPanic(value interface{}) (res uint8) {
	if res, err := ToUint8(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUint16OrPanic converts any type of value to uint16 or call the panic function
func ToUint16OrPanic(value interface{}) (res uint16) {
	if res, err := ToUint16(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUin32OrPanic converts any type of value to uint32 or call the panic function
func ToUin32OrPanic(value interface{}) (res uint32) {
	if res, err := ToUint32(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUint64OrPanic converts any type of value to uint64 or call the panic function
func ToUint64OrPanic(value interface{}) (res uint64) {
	if res, err := ToUint64(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToFloat64OrPanic converts any type of value to float64 or call the panic function
func ToFloat64OrPanic(value interface{}) (res float64) {
	if res, err := ToFloat64(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToIntArray converts any type of values to array of int
func ToIntArray(value interface{}) (resArray []int, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int, v.Len())
			for index := range resArray {
				resArray[index], err = ToInt(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToInt8Array converts any type of values to array of int8
func ToInt8Array(value interface{}) (resArray []int8, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int8, v.Len())
			for index := range resArray {
				resArray[index], err = ToInt8(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToInt16Array converts any type of values to array of int16
func ToInt16Array(value interface{}) (resArray []int16, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int16, v.Len())
			for index := range resArray {
				resArray[index], err = ToInt16(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToInt32Array converts any type of values to array of int32
func ToInt32Array(value interface{}) (resArray []int32, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int32, v.Len())
			for index := range resArray {
				resArray[index], err = ToInt32(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)

}

// ToInt64Array converts any type of values to array of int64
func ToInt64Array(value interface{}) (resArray []int64, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]int64, v.Len())
			for index := range resArray {
				resArray[index], err = ToInt64(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)

}

// ToUintArray converts any type of values to array of uint
func ToUintArray(value interface{}) (resArray []uint, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint, v.Len())
			for index := range resArray {
				resArray[index], err = ToUint(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToUint8Array converts any type of values to array of uint8
func ToUint8Array(value interface{}) (resArray []uint8, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint8, v.Len())
			for index := range resArray {
				resArray[index], err = ToUint8(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToUint16Array converts any type of values to array of uint16
func ToUint16Array(value interface{}) (resArray []uint16, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint16, v.Len())
			for index := range resArray {
				resArray[index], err = ToUint16(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToUint32Array converts any type of values to array of uint32
func ToUint32Array(value interface{}) (resArray []uint32, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint32, v.Len())
			for index := range resArray {
				resArray[index], err = ToUint32(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToUint64Array converts any type of values to array of uint64
func ToUint64Array(value interface{}) (resArray []uint64, err error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Array, reflect.Slice:
		{
			v := reflect.ValueOf(value)
			resArray = make([]uint64, v.Len())
			for index := range resArray {
				resArray[index], err = ToUint64(v.Index(index).Interface())
				if err != nil {
					return nil, fmt.Errorf("convert: can not convert %v at index %d", v.Index(index).Interface(), index)
				}
			}
			return resArray, nil
		}
	}
	return nil, fmt.Errorf("convert: %T is not an array or slice", value)
}

// ToIntArrayOrPanic converts any type of values to a valid array of int or call the panic function
func ToIntArrayOrPanic(value interface{}) (resArray []int) {
	if res, err := ToIntArray(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt8ArrayOrPanic converts any type of values to a valid array of int8 or call the panic function
func ToInt8ArrayOrPanic(value interface{}) (resArray []int8) {
	if res, err := ToInt8Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt16ArrayOrPanic converts any type of values to a valid array of int16 or call the panic function
func ToInt16ArrayOrPanic(value interface{}) (resArray []int16) {
	if res, err := ToInt16Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt32ArrayOrPanic converts any type of values to a valid array of int32 or call the panic function
func ToInt32ArrayOrPanic(value interface{}) (resArray []int32) {
	if res, err := ToInt32Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToInt64ArrayOrPanic converts any type of values to a valid array of int64 or call the panic function
func ToInt64ArrayOrPanic(value interface{}) (resArray []int64) {
	if res, err := ToInt64Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUintArrayOrPanic converts any type of values to a valid array of uint or call the panic function
func ToUintArrayOrPanic(value interface{}) (resArray []uint) {
	if res, err := ToUintArray(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUint8ArrayOrPanic converts any type of values to a valid array of uint8 or call the panic function
func ToUint8ArrayOrPanic(value interface{}) (resArray []uint8) {
	if res, err := ToUint8Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUint16ArrayOrPanic converts any type of values to a valid array of uint16 or call the panic function
func ToUint16ArrayOrPanic(value interface{}) (resArray []uint16) {
	if res, err := ToUint16Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUint32ArrayOrPanic converts any type of values to a valid array of uint32 or call the panic function
func ToUint32ArrayOrPanic(value interface{}) (resArray []uint32) {
	if res, err := ToUint32Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}

// ToUint64ArrayOrPanic converts any type of values to a valid array of uint64 or call the panic function
func ToUint64ArrayOrPanic(value interface{}) (resArray []uint64) {
	if res, err := ToUint64Array(value); err == nil {
		return res
	} else {
		panic(err)
	}
}
