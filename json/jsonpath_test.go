package json

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNewJSONPath(t *testing.T) {
	type JS struct {
		Int          int                    `json:"int" json_path:"int"`
		Uint         uint                   `json:"uint" json_path:"uint"`
		Float        float64                `json:"float" json_path:"float"`
		String       string                 `json:"string" json_path:"string"`
		Bool         bool                   `json:"bool" json_path:"bool"`
		IntLevel2    int                    `json:"int_level_2" json_path:"level2.int"`
		UintLevel2   uint                   `json:"uint_level_2" json_path:"level2.uint"`
		FloatLevel2  float64                `json:"float_level_2" json_path:"level2.float"`
		StringLevel2 string                 `json:"string_level_2" json_path:"level2.string"`
		BoolLevel2   bool                   `json:"bool_level_2" json_path:"level2.bool"`
		Map          map[string]interface{} `json:"map" json_path:"map"`
		Array        []interface{}          `json:"array" json_path:"array"`
		Array2       [][]interface{}        `json:"array2" json_path:"array2"`
		Array2int    [][]int                `json:"array2int" json_path:"array2"`
		Array3       [][][]interface{}      `json:"array3" json_path:"array3"`
		Array3int    [][][]int              `json:"array3int" json_path:"array3"`
		StringArray  []string               `json:"string_array" json_path:"string_array"`
		StructArray  []struct {
			A int     `json:"a"`
			B string  `json:"b"`
			C float64 `json:"c"`
		} `json:"struct_array" json_path:"arraystruct"`
		Struct0 struct {
			Struct0A int     `json_path:"arraystruct.0.a"`
			Struct0B string  `json_path:"arraystruct.0.b"`
			Struct0C float64 `json_path:"arraystruct.0.c"`
		} `json:"struct0"`
	}
	jp, err := NewJSONPath(
		[]byte(`
{
    "int": 1,
    "uint": 1,
    "float": 1.11,
    "string": "jsonpath",
    "bool": true,
    "map": {
        "a": 1,
        "b": "2",
        "c": 3
    },
    "array": [
        1,
        "2",
        3
    ],
    "string_array": [
        "asdf",
        "ghjk",
        "zxcv"
    ],
    "string_array_null": [
        "abc",
        null,
        "efg"
    ],
    "arraywithmap": [
        {
            "map11": 1
        },
        {
            "map21": 1,
            "map22": 2
        }
    ],
    "array2": [
        [
            1,
            2
        ],
        [
            2,
            3
        ],
        [
            3,
            4
        ]
    ],
    "array3": [
        [
            [
                1,
                2
            ]
        ],
        [
            [
                2,
                3
            ]
        ],
        [
            [
                3,
                4
            ]
        ]
    ],
    "arraystruct": [
        {
            "a": 1,
            "b": "1",
            "c": 1.11
        },
        {
            "a": 2,
            "b": "2",
            "c": 2.22
        }
    ],
    "level2": {
        "int": 2,
        "uint": 2,
        "float": 2.22,
        "string": "jsonpath2",
        "bool": false
    }
}`),
	)
	js := JS{}

	assert.NotEqual(t, nil, jp)
	assert.Equal(t, nil, err)

	var ok bool
	_, ok = jp.Get2("int")
	assert.Equal(t, true, ok)

	_, ok = jp.Get2("missing_key")
	assert.Equal(t, false, ok)

	awm := jp.Get("arraywithmap")
	assert.NotEqual(t, nil, awm)
	var awsval int
	awsval, _ = awm.Get(0).Get("map11").Int()
	assert.Equal(t, 1, awsval)
	awsval, _ = awm.Get(1).Get("map21").Int()
	assert.Equal(t, 1, awsval)
	awsval, _ = awm.Get(1).Get("map22").Int()
	assert.Equal(t, 2, awsval)

	i, _ := jp.Get("int").Int()
	assert.Equal(t, 1, i)

	f, _ := jp.Get("float").Float64()
	assert.Equal(t, 1.11, f)

	s, _ := jp.Get("string").String()
	assert.Equal(t, "jsonpath", s)

	b, _ := jp.Get("bool").Bool()
	assert.Equal(t, true, b)

	strs, err := jp.Get("string_array").StringArray()
	assert.Equal(t, nil, err)
	assert.Equal(t, strs[0], "asdf")
	assert.Equal(t, strs[1], "ghjk")
	assert.Equal(t, strs[2], "zxcv")

	strs2, err := jp.Get("string_array_null").StringArray()
	assert.Equal(t, nil, err)
	assert.Equal(t, strs2[0], "abc")
	assert.Equal(t, strs2[1], "")
	assert.Equal(t, strs2[2], "efg")

	gp, _ := jp.GetPath("level2", "string").String()
	assert.Equal(t, "jsonpath2", gp)

	gp2, _ := jp.GetPath("level2", "int").Int()
	assert.Equal(t, 2, gp2)

	err = jp.ParseWithJSONPath(&js)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, js.Int)
	assert.Equal(t, uint(1), js.Uint)
	assert.Equal(t, 1.11, js.Float)
	assert.Equal(t, "jsonpath", js.String)
	assert.Equal(t, true, js.Bool)
	assert.Equal(t, 2, js.IntLevel2)
	assert.Equal(t, uint(2), js.UintLevel2)
	assert.Equal(t, 2.22, js.FloatLevel2)
	assert.Equal(t, "jsonpath2", js.StringLevel2)
	assert.Equal(t, false, js.BoolLevel2)
	assert.Equal(t, map[string]interface{}{"a": Number("1"), "b": "2", "c": Number("3")}, js.Map)
	assert.Equal(t, []interface{}{Number("1"), "2", Number("3")}, js.Array)
	assert.Equal(
		t,
		[][]interface{}{{Number("1"), Number("2")}, {Number("2"), Number("3")}, {Number("3"), Number("4")}},
		js.Array2,
	)
	assert.Equal(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, js.Array2int)
	assert.Equal(
		t,
		[][][]interface{}{{{Number("1"), Number("2")}}, {{Number("2"), Number("3")}}, {{Number("3"), Number("4")}}},
		js.Array3,
	)
	assert.Equal(t, [][][]int{{{1, 2}}, {{2, 3}}, {{3, 4}}}, js.Array3int)
	assert.Equal(t, []string{"asdf", "ghjk", "zxcv"}, js.StringArray)
	assert.Equal(
		t, []struct {
			A int     `json:"a"`
			B string  `json:"b"`
			C float64 `json:"c"`
		}{
			{
				A: 1,
				B: "1",
				C: 1.11,
			}, {
				A: 2,
				B: "2",
				C: 2.22,
			},
		}, js.StructArray,
	)
	assert.Equal(
		t, struct {
			Struct0A int     `json_path:"arraystruct.0.a"`
			Struct0B string  `json_path:"arraystruct.0.b"`
			Struct0C float64 `json_path:"arraystruct.0.c"`
		}{
			Struct0A: 1,
			Struct0B: "1",
			Struct0C: 1.11,
		}, js.Struct0,
	)
}
