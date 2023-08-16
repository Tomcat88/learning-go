package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	timeoutServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(11 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))

	defer slowServer.Close()
	defer fastServer.Close()
	defer timeoutServer.Close()

	type args struct {
		website1 string
		website2 string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "case 1",
			args: args{
				website1: slowServer.URL,
				website2: fastServer.URL,
			},
			want:    fastServer.URL,
			wantErr: nil,
		},
		{
			name: "case 2",
			args: args{
				website1: timeoutServer.URL,
				website2: timeoutServer.URL,
			},
			want:    "",
			wantErr: ErrTimeout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConfigurableRacer(tt.args.website1, tt.args.website2, 50 * time.Millisecond)
			if tt.wantErr != err {
				t.Errorf("WebsiteRacer() = %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("WebsiteRacer() = %v, want %v", got, tt.want)

			}
		})
	}
}
