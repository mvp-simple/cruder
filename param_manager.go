package cruder

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type paramManager struct {
	uri *url.URL
}

func NewParamManager(uri *url.URL) IParamManager {
	return &paramManager{uri: uri}
}

var errParamNotFound = errors.New("param not found")

func (r *paramManager) StrSlice(keyIn string) (sliceOut []string, errOut error) {
	sliceOut, ok := map[string][]string(r.uri.Query())[keyIn]
	if !ok {
		errOut = errParamNotFound
	}
	return
}

func (r *paramManager) Str(keyIn string) (out string, errOut error) {
	var strSliceOut []string
	strSliceOut, errOut = r.StrSlice(keyIn)
	if errOut != nil {
		return
	}
	return strings.Join(strSliceOut, ""), nil
}
func (r *paramManager) IntSlice(keyIn string) (sliceOut []int64, errOut error) {
	var strSliceOut []string
	strSliceOut, errOut = r.StrSlice(keyIn)
	if errOut != nil {
		return
	}
	var parseInt int64
	for _, str := range strSliceOut {
		parseInt, errOut = strconv.ParseInt(str, 10, 64)
		if errOut != nil {
			return
		}
		sliceOut = append(sliceOut, parseInt)
	}
	return
}

func (r *paramManager) Int(keyIn string) (out int64, errOut error) {
	str, err := r.Str(keyIn)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(str, 10, 64)
}

func (r *paramManager) BoolSlice(keyIn string) (sliceOut []bool, errOut error) {
	var strSliceOut []string
	strSliceOut, errOut = r.StrSlice(keyIn)
	if errOut != nil {
		return
	}
	var parseBool bool
	for _, str := range strSliceOut {
		parseBool, errOut = strconv.ParseBool(str)
		if errOut != nil {
			return
		}
		sliceOut = append(sliceOut, parseBool)
	}
	return
}

func (r *paramManager) Bool(keyIn string) (out bool, errOut error) {
	str, err := r.Str(keyIn)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(str)
}

func (r *paramManager) FloatSlice(keyIn string) (sliceOut []float64, errOut error) {
	var strSliceOut []string
	strSliceOut, errOut = r.StrSlice(keyIn)
	if errOut != nil {
		return
	}
	var parseInt float64
	for _, str := range strSliceOut {
		parseInt, errOut = strconv.ParseFloat(str, 64)
		if errOut != nil {
			return
		}
		sliceOut = append(sliceOut, parseInt)
	}
	return
}

func (r *paramManager) Float(keyIn string) (out float64, errOut error) {
	str, err := r.Str(keyIn)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(str, 64)
}
