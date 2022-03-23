package scheduling

import (
	"github.com/grussorusso/serverledge/internal/node"
	"log"
)

//EdgePolicy supports only Edge-Edge offloading
type EdgePolicy struct{}

func (p *EdgePolicy) Init() {
	InitDropManager()
}

func (p *EdgePolicy) OnCompletion(r *scheduledRequest) {

}

func (p *EdgePolicy) OnArrival(r *scheduledRequest) {
	containerID, err := node.AcquireWarmContainer(r.Fun)
	if err == nil {
		log.Printf("Using a warm container for: %v", r)
		execLocally(r, containerID, true)
	} else if handleColdStart(r) {
		return
	} else {
		url := handleEdgeOffloading(r)
		if url != "" {
			handleOffload(r, url)
			return
		}
	}
	dropRequest(r)
}