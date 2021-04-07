package v1

import (
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}"
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}/repository/mock"
	"{{cookiecutter.module_name}}/internal/{{cookiecutter.main_domain}}/service"
	"{{cookiecutter.module_name}}/internal/api"
	"{{cookiecutter.module_name}}/settings"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHandler(t *testing.T) {
	cfg := &settings.Settings{}
	svc := service.New(mock.NewWithExpectations())
	vld := api.NewValidator()

	type args struct {
		cfg *settings.Settings
		svc {{cookiecutter.main_domain}}.Service
	}
	tests := []struct {
		name string
		args args
		want *handler
	}{
		{
			name: "Basic",
			args: args{
				cfg: cfg,
				svc: svc,
			},
			want: &handler{
				cfg:       cfg,
				service:   svc,
				validator: vld,
			},
		},
	}

	assert := require.New(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHandler(cfg, svc, vld)
			assert.Equal(tt.want, got)
		})
	}
}
