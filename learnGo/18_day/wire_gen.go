// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeShop() shop {
	mainApple := NewA()
	mainBanana := NewB()
	mainShop := NewS(mainApple, mainBanana)
	return mainShop
}
