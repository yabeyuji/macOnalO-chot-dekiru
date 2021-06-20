package entity

import (
	"context"
	"time"
)

type (
	entity struct{}

	// ToEntity ...
	ToEntity interface {
		ParseOrder(ctx context.Context, order *Order) *Assemble
		CookHamburgers(ctx context.Context, hamburgers []Hamburger) error
	}
)

type (
	// framework_driver から来る用
	// Order ...
	Order struct {
		OrderInfo OrderInfo
		Product   Product
	}

	// OrderInfo ...
	OrderInfo struct {
		OrderNumber string
		OrderType   string
		OrderTime   time.Time
		PassTime    time.Time
	}

	// Product ...
	Product struct {
		Combos     []Combo     `json:"combos"`
		Hamburgers []Hamburger `json:"hamburgers"`
		SideMenus  []SideMenu  `json:"side_menus"`
		Drinks     []Drink     `json:"drinks"`
	}

	// Combo ...
	Combo struct {
		Hamburger *Hamburger `json:"hamburger"`
		SideMenu  *SideMenu  `json:"side_menu"`
		Drink     *Drink     `json:"drink"`
	}

	// Hamburger ...
	Hamburger struct {
		// bans
		Top    int `json:"top"`
		Middle int `json:"middle"`
		Bottom int `json:"bottom"`
		// patty
		Beef    int `json:"beef"`
		Chicken int `json:"chicken"`
		Fish    int `json:"fish"`
		//vegetable
		Lettuce int `json:"lettuce"`
		Tomato  int `json:"tomato"`
		Onion   int `json:"onion"`
		//ingredient
		Cheese  int `json:"cheese"`
		Pickles int `json:"pickles"`
	}

	// SideMenu ...
	SideMenu struct {
		Name string `json:"name"`
	}

	// Drink ...
	Drink struct {
		Name string `json:"name"`
	}
)

// framework_driver へ行く用
type (
	// Assemble ...
	Assemble struct {
		Bans        map[string]int
		Patties     map[string]int
		Vegetables  map[string]int
		Ingredients map[string]int
	}
)
