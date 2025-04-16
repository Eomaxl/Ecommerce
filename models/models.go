package models

type User struct {
	ID
	First_Name
	Last_Name
	Password
	Email
	Phone
	Token
	Refresh_Token
	Created_At
	Updated_At
	User_ID
	UserCart
	Address_Details
	Order_Status
}

type Product struct {
	Product_ID
	Product_Name
	Price
	Rating
	Image
}

type ProductUser struct {
}

type Address struct {
}

type Order struct {
}

type Payment struct {
}
