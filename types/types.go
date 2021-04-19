package types

import (
	"github.com/google/uuid"
)

type DeployTarget struct {
	Ip       string `json:"ip"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type FilecoinDeployParams struct {
	Fullnodes []DeployTarget `json:"fullnodes"`
	Miners    []DeployTarget `json:"miners"`
	Workers   []DeployTarget `json:"workers"`
	Storages  []DeployTarget `json:"storages"`
}

type FilecashDeployParams FilecoinDeployParams

type CreateDeployInput struct {
	AuthCode      string      `json:"auth_code"`
	DeployType    string      `json:"deploy_type"`
	TargetGateway uuid.UUID   `json:"target_gateway"`
	DeployParams  interface{} `json:"params"`
}

type CreateDeployOutput struct {
	DeployId uuid.UUID `json:"id"`
}

type QueryDeployInput struct {
	AuthCode string    `json:"auth_code"`
	DeployId uuid.UUID `json:"id"`
}

type DeleteDeployInput QueryDeployInput
