package example

type Server struct {
	done    chan struct{}
	stopped bool
}

func NewServer() *Server {
	return &Server{
		done: make(chan struct{}),
	}
}

func (s *Server) Stop() {
	// clean up resources, etc
	close(s.done)
	s.stopped = true // TODO: use a sync.RWMutex
}

func (s *Server) Stopped() bool {
	return s.stopped
}

func (s *Server) Done() <-chan struct{} {
	return s.done // TODO: use a sync.RWMutex
}

//-

func (s *Server) Multiply(width, height int64) int64 {
	return width * height
}
