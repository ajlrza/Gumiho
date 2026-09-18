package test
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

func MockResource() {

	ResourceData := Resource {
		GPUModel: "NVIDIA",
		MemoryPerc: 29,
		Status: "UP",
	}

	fmt.Println(ResourceData)
}

func MockPayload() {

	HPPayloadData := HPPayload{
		Checksum: "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
		Header: Header{ 
			Name: "Content-Type",
			Value: "application/json",
		},
		Data: [256]HPData{
			23.5, 24.1, 22.8, 25.0, 24.6,
			23.9, 26.2, 25.8, 24.0, 23.4,
			22.1, 24.7, 25.3, 26.0, 24.9,
		},
		Trailer: "Maybe inspired by networking",
	}

	MPPayloadData := MPPayload{
		Checksum: "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6",
		Header: Header{ 
			Name: "Content-Type",
			Value: "application/json",
		},
		Data: [256]MPData{
			23.5, 24.1, 22.8, 25.0, 24.6,
			23.9, 26.2, 25.8, 24.0, 23.4,
			22.1, 24.7, 25.3, 26.0, 24.9,
		},
		Trailer: "Maybe inspired by networking",
	}

	fmt.Println(HPPayloadData, MPPayloadData)
}