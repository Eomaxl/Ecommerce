package controllers

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("Can't find product")
	ErrCantDecodeProducts = errors.New("Can't decode products")
	ErrUserIdIsNotValid   = errors.New("UserId is not valid")
	ErrCantUpdateUser     = errors.New("Can't update user")
	ErrCantRemoveItemCart = errors.New("Can't remove item cart")
	ErrCantGetItem        = errors.New("Can't get item")
	ErrCantBuyCartItem    = errors.New("Can't buy cart item")
)

func AddProductToCart() {}

func RemoveCartItem() {}

func BuyItemFromCart() {}

func InstantBuyer() {}
