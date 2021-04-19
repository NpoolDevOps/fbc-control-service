package types

import (
	"github.com/google/uuid"
)

type CreateDeployInput struct {
	AuthCode string `json:"auth_code"`
}

type CreateDeployOutput struct {
	DeployId uuid.UUID `json:"id"`
}

type QueryDeployInput struct {
	AuthCode string    `json:"auth_code"`
	DeployId uuid.UUID `json:"id"`
}

type DeleteDeployInput QueryDeployInput
