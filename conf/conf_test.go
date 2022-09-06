package conf

import (
	"reflect"
	"testing"
)

func TestGetConf(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConf()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
