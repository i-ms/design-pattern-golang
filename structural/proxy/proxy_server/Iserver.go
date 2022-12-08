package proxy_server

type Server interface {
	handleRequest(string, string) (int, string)
}
