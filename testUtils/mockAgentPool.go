package testUtils

import (
	"github.com/jarcoal/httpmock"
	"github.com/kitwalker12/go-instagram/utils"
)

func MockAgentPool(capacity int) (*utils.SuperAgentPool, error) {
	pool, _ := utils.NewSuperAgentPool(capacity)

	for i := 0; i < pool.Len(); i++ {
		agent := pool.Get()
		httpmock.ActivateNonDefault(agent.Client)
		pool.Put(agent)
	}

	return pool, nil
}
