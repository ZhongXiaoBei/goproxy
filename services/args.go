package services

// tcp := app.Command("tcp", "proxy on tcp mode")
// t := tcp.Flag("tcp-timeout", "tcp timeout milliseconds when connect to real server or parent proxy").Default("2000").Int()

const (
	TYPE_TCP  = "tcp"
	TYPE_UDP  = "udp"
	TYPE_HTTP = "http"
	TYPE_TLS  = "tls"
)

type Args struct {
	Local               *string
	Parent              *string
	CertBytes           []byte
	KeyBytes            []byte
	PoolSize            *int
	CheckParentInterval *int
}
type TunnelArgs struct {
	Args
	Timeout *int
}
type TCPArgs struct {
	Args
	Timeout    *int
	ParentType *string
	IsTLS      *bool
}

type HTTPArgs struct {
	Args
	Always      *bool
	HTTPTimeout *int
	Timeout     *int
	Interval    *int
	Blocked     *string
	Direct      *string
	AuthFile    *string
	Auth        *[]string
	ParentType  *string
	LocalType   *string
}
type UDPArgs struct {
	Args
	ParentType *string
	Timeout    *int
}

func (a *TCPArgs) Protocol() string {
	if *a.IsTLS {
		return "tls"
	}
	return "tcp"
}