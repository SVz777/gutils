/**
 * @file    multitask_test.go
 * @author  903943711@qq.com
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * @date    2021/8/21
 * @desc
 */
package multitask_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bmizerany/assert"

	"github.com/SVz777/gutils/multitask"
)

func TestMultiTask(t *testing.T) {
	type fields struct {
		fs     map[string]multitask.Do
		result interface{}
		err    error
	}
	tests := []struct {
		name    string
		fields  fields
		wantRet interface{}
	}{
		{
			name: "",
			fields: fields{
				fs: map[string]multitask.Do{
					"k1": func(ctx context.Context) (interface{}, error) {
						return "k1", nil
					},
					"k2": func(ctx context.Context) (interface{}, error) {
						time.Sleep(1 * time.Second)
						return "k2", fmt.Errorf("k2 err")
					},
					"k3": func(ctx context.Context) (interface{}, error) {
						time.Sleep(2 * time.Second)
						return "k3", fmt.Errorf("k3 err")
					},
					"k4": func(ctx context.Context) (interface{}, error) {
						time.Sleep(3 * time.Second)
						return "k4", fmt.Errorf("k4 err")
					},
				},

				result: nil,
				err:    nil,
			},
			wantRet: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
			mt := multitask.NewTaskManager(ctx)
			for k, v := range tt.fields.fs {
				mt.Add(k, v)
			}
			_ = mt.Do()
			assert.Equal(t, nil, mt.GetTaskErr("k1"))
			assert.Equal(t, "k1", mt.GetTaskResult("k1"))
			assert.NotEqual(t, nil, mt.GetTaskErr("k2"))
			assert.Equal(t, nil, mt.GetTaskResult("k2"))
			assert.NotEqual(t, nil, mt.GetTaskErr("k3"))
			assert.Equal(t, nil, mt.GetTaskResult("k3"))
			assert.NotEqual(t, nil, mt.GetTaskErr("k4"))
			assert.Equal(t, nil, mt.GetTaskResult("k4"))
		})
	}
}