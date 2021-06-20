package freezer

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/internal/2_interface_adapter/gateway"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "freezer")
}

type (
	// Freezer ...
	Freezer struct {
		Conn *gorm.DB
	}

	// Vegetable ...
	Vegetable struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToFreezer ...
func NewToFreezer() gateway.ToFreezer {
	conn, err := open(30)
	if err != nil {
		myErr.Logging(err)
		panic(err)
	}

	s := new(Freezer)
	s.Conn = conn
	return s
}

func open(count uint) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(pkg.MySQLDSN), &gorm.Config{})
	if err != nil {
		if count == 0 {
			myErr.Logging(err)
			return nil, fmt.Errorf("Retry count over")
		}
		time.Sleep(time.Second)
		// カウントダウンさせるようにする
		count--
		return open(count)
	}

	return db, nil
}

// UpdatePatties ...
func (s *Freezer) UpdatePatties(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := s.Conn.
			Table("patties").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			myErr.Logging(res.Error)
			return res.Error
		}

		// 作業時間を擬似的に再現
		time.Sleep(1 * time.Second)
	}
	return nil
}
