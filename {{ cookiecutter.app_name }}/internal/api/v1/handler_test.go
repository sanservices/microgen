package v1

import (
	"github.com/sanservices/apicore/validator"
	_ "{{ cookiecutter.module_name }}/internal/api/v1/swagger" // statik file

	"github.com/rakyll/statik/fs"
	"github.com/stretchr/testify/require"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/repository/mock"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.main_domain }}/service"
	"{{ cookiecutter.module_name }}/settings"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	cfg := &settings.Settings{}
	svc := service.New(mock.NewWithExpectations())
	vld := validator.NewValidator()

	statikFS, err := fs.New()
	// Try to get swagger from statik
	if err != nil {
		panic(err)
	}

	type args struct {
		cfg      *settings.Settings
		svc      {{ cookiecutter.main_domain }}.Service
		statikFS http.FileSystem
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "Basic",
			args: args{
				cfg:      cfg,
				svc:      svc,
				statikFS: statikFS,
			},
			want: &Handler{
				cfg:       cfg,
				service:   svc,
				validator: vld,
				statikFS:  statikFS,
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
