package entity

import (
	"context"
	"time"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("enterprise_business_rule", "entity")
}

// NewEntity ...
func NewEntity() ToEntity {
	return &entity{}
}

// ParseOrder ...
func (entt *entity) ParseOrder(ctx context.Context, order *Order) *Assemble {
	assemble := &Assemble{
		Bans:        map[string]int{},
		Patties:     map[string]int{},
		Vegetables:  map[string]int{},
		Ingredients: map[string]int{},
	}

	// オーダー内容から必要な食材の数を計算
	if len(order.Product.Hamburgers) != 0 {
		entt.countAssembleHamburger(ctx, assemble, order.Product.Hamburgers)
	}

	return assemble
}

func (entt *entity) countAssembleHamburger(ctx context.Context, assemble *Assemble, hamburgers []Hamburger) {

	for _, hamburger := range hamburgers {
		// bans
		assemble.Bans["top"] += hamburger.Top
		assemble.Bans["middle"] += hamburger.Middle
		assemble.Bans["bottom"] += hamburger.Bottom

		// patty
		assemble.Patties["beef"] += hamburger.Beef
		assemble.Patties["chicken"] += hamburger.Chicken
		assemble.Patties["fish"] += hamburger.Fish

		//vegetable
		assemble.Vegetables["lettuce"] += hamburger.Lettuce
		assemble.Vegetables["tomato"] += hamburger.Tomato
		assemble.Vegetables["onion"] += hamburger.Onion

		//ingredient
		assemble.Ingredients["cheese"] += hamburger.Cheese
		assemble.Ingredients["pickles"] += hamburger.Pickles
	}
}

// CookHamburgers ...
func (entt *entity) CookHamburgers(ctx context.Context, hamburgers []Hamburger) error {
	// ハンバーガーの調理
	for _, hamburger := range hamburgers {
		entt.cutVegetables(ctx, hamburger)
		entt.grillPatties(ctx, hamburger)
		entt.assembleHamburger(ctx, hamburger)
	}

	return nil
}

func (entt *entity) cutVegetables(ctx context.Context, hamburger Hamburger) {
	// 調理時間を擬似的に再現
	if hamburger.Lettuce > 0 {
		time.Sleep(1 * time.Second)
	}
	if hamburger.Onion > 0 {
		time.Sleep(1 * time.Second)
	}
	if hamburger.Pickles > 0 {
		time.Sleep(1 * time.Second)
	}
}

func (entt *entity) assembleHamburger(ctx context.Context, hamburger Hamburger) {
	// 調理時間を擬似的に再現
	time.Sleep(1 * time.Second)
}

func (entt *entity) grillPatties(ctx context.Context, hamburger Hamburger) {
	// 調理時間を擬似的に再現
	if hamburger.Beef > 0 {
		time.Sleep(time.Duration(hamburger.Beef*1) * time.Second)
	}
	if hamburger.Chicken > 0 {
		time.Sleep(time.Duration(hamburger.Chicken*1) * time.Second)
	}
	if hamburger.Fish > 0 {
		time.Sleep(time.Duration(hamburger.Fish*1) * time.Second)
	}
}
