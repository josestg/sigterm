package sigterm

import (
	"syscall"
	"testing"
)

func TestSignal_String(t *testing.T) {
	tests := []struct {
		name string
		s    Signal
		want string
	}{
		{name: "Interrupt signal", s: SIGINT, want: "SIGINT"},
		{name: "Hangup signal", s: SIGHUP, want: "SIGHUP"},
		{name: "Terminate signal", s: SIGTERM, want: "SIGTERM"},
		{name: "Quit signal", s: SIGQUIT, want: "SIGQUIT"},
		{name: "Kill signal", s: SIGKILL, want: "SIGKILL"},
		{name: "Unknown signal", s: -1, want: "sigterm.Signal(-1): unknown signal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignal_Unwrap(t *testing.T) {
	if SIGTERM.Unwrap() != syscall.SIGTERM {
		t.Errorf("Unwrap() = %v, want %v", SIGTERM.Unwrap(), syscall.SIGTERM)
	}
}

func TestSignal_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		text    []byte
		wantErr bool
	}{
		{name: "Interrupt signal", text: []byte("SIGINT"), wantErr: false},
		{name: "Hangup signal", text: []byte("SIGHUP"), wantErr: false},
		{name: "Terminate signal", text: []byte("SIGTERM"), wantErr: false},
		{name: "Quit signal", text: []byte("SIGQUIT"), wantErr: false},
		{name: "Kill signal", text: []byte("SIGKILL"), wantErr: false},
		{name: "Empty signal", text: []byte(""), wantErr: true},
		{name: "Invalid signal", text: []byte("SIG_UNKNOWN"), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Signal
			if err := s.UnmarshalText(tt.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignal_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		s       Signal
		wantErr bool
	}{
		{name: "Interrupt signal", s: SIGINT, wantErr: false},
		{name: "Hangup signal", s: SIGHUP, wantErr: false},
		{name: "Terminate signal", s: SIGTERM, wantErr: false},
		{name: "Quit signal", s: SIGQUIT, wantErr: false},
		{name: "Kill signal", s: SIGKILL, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsTermination(t *testing.T) {
	t.Run("sigterm.Signal", func(t *testing.T) {
		tests := []struct {
			name string
			sig  Signal
			want bool
		}{
			{name: "Interrupt signal", sig: SIGINT, want: true},
			{name: "Hangup signal", sig: SIGHUP, want: true},
			{name: "Terminate signal", sig: SIGTERM, want: true},
			{name: "Quit signal", sig: SIGQUIT, want: true},
			{name: "Kill signal", sig: SIGKILL, want: true},
			{name: "Unknown signal", sig: -1, want: false},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := IsTermination(tt.sig); got != tt.want {
					t.Errorf("IsTermination() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("syscall.Signal", func(t *testing.T) {
		tests := []struct {
			name string
			sig  syscall.Signal
			want bool
		}{
			{name: "Interrupt signal", sig: syscall.SIGINT, want: true},
			{name: "Hangup signal", sig: syscall.SIGHUP, want: true},
			{name: "Terminate signal", sig: syscall.SIGTERM, want: true},
			{name: "Quit signal", sig: syscall.SIGQUIT, want: true},
			{name: "Kill signal", sig: syscall.SIGKILL, want: true},
			{name: "Unknown signal", sig: syscall.SIGSEGV, want: false},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := IsTermination(tt.sig); got != tt.want {
					t.Errorf("IsTermination() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
