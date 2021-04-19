package main

type ControlServer struct {
}

func NewControlServer(config string) *ControlServer {
	s := &ControlServer{}
	return s
}

func (s *ControlServer) Run() error {
	return nil
}
