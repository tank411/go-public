// +build wireinject

package main

import "github.com/google/wire"

func InitializeShop() shop {
	wire.Build(NewA, NewS, NewB)
	return shop{}
}
