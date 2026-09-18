package app
import "fmt"

type Header struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type HPData struct {
	Tensor float64
}

type MPData struct {
	Tensor float32
}

type HPPayload struct {
	Checksum string `json:"checksum"`
	Header Header `json:"header"`
	Data [256]HPData `json:"data"`
	Trailer string `json:"trailer"`
}

type MPPayload struct {
	Checksum string `json:"checksum"`
	Header Header `json:"header"`
	Data [256]MPData `json:"data"`
	Trailer string `json:"trailer"`
}


type Resource struct {
	GPUModel string `json:"model"`
	MemoryPerc int `json:"perc"`
	Status string `json:"status"`
}

type DockerCall struct {

}

func app() {
	fmt.Println("App")
}