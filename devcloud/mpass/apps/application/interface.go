package application

import "context"

type Service interface {
	CreateApplication(context.Context)
}
