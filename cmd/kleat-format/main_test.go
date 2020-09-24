package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_format(t *testing.T) {
	type args struct {
		data string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				data: `providers:
  kubernetes:
    enabled: true
    accounts:
    - name: k8s
      kubeconfigFile: /var/secrets/k8s-kubeconfig
    primaryAccount: k8s
`,
			},
			want: `providers:
  kubernetes:
    accounts:
    - kubeconfigFile: /var/secrets/k8s-kubeconfig
      name: k8s
    enabled: true
    primaryAccount: k8s
`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := format([]byte(tt.args.data))
			if (err != nil) != tt.wantErr {
				t.Errorf("format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, string(got)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
