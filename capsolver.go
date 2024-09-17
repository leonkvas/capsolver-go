package capsolver_go

import (
	"errors"
	"time"
)

func (c CapSolver) Solve(task map[string]any) (*capSolverResponse, error) {
	//
	capRes, err := request(CREATE_TASK_URI, &capSolverRequest{Task: &task, ClientKey: c.getApiKey(), AppId: "16EF5A46-0889-4B03-BC61-4D900B387202"})
	if err != nil {
		return nil, err
	}
	if capRes.ErrorId == 1 {
		return nil, errors.New(capRes.ErrorDescription)
	}
	if capRes.Status == STATUS_READY {
		return capRes, nil
	}
	for i := 0; i < TASK_TIMEOUT; i++ {
		capRes, err = request(GET_TASK_URI, &capSolverRequest{ClientKey: c.getApiKey(), TaskId: capRes.TaskId, AppId: "16EF5A46-0889-4B03-BC61-4D900B387202"})
		time.Sleep(time.Second * 1)
		if err != nil {
			return nil, err
		}
		if capRes.ErrorId == 1 {
			return nil, errors.New(capRes.ErrorDescription)
		}
		if capRes.Status == STATUS_READY {
			break
		}
	}
	return capRes, err
}

func (c CapSolver) Balance() (*capSolverResponse, error) {
	capRes, err := request(BALANCE_URI, &capSolverRequest{ClientKey: c.getApiKey()})
	if err != nil {
		return nil, err
	}
	return capRes, nil

}

func (c *CapSolver) getApiKey() string {
	if c.ApiKey != "" {
		return c.ApiKey
	}
	return ApiKey
}
