package builder

type Server struct {
	port int
	host string
}

func NewServer(port int, host string) *Server {
	return &Server{
		port: port,
		host: host,
	}
}
