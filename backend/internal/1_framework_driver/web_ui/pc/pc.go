package pc

import (
	"github.com/gin-gonic/gin"

	"backend/internal/2_interface_adapter/controller"
	"backend/internal/4_enterprise_business_rule/entity"
	"backend/pkg"
)

var (
	orderType = "pc"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "pc")
}

type (
	// PC ...
	PC struct {
		GinEngine  *gin.Engine
		Controller controller.ToController
	}
)

// NewPC ...
func NewPC(ctrl controller.ToController) *PC {
	pc := &PC{
		GinEngine:  NewGin(),
		Controller: ctrl,
	}

	return pc
}

// NewGin ...
func NewGin() *gin.Engine {
	return gin.Default()
}

// Start ...
func (pc *PC) Start() {
	pc.GinEngine.POST("/", pc.IndexPost)

	pc.GinEngine.Run(pkg.PCPort)
}

// IndexPost ...
func (pc *PC) IndexPost(c *gin.Context) {
	// 標準コンテキストを取得
	ctx := c.Request.Context()

	// web_uiのデータ型をControllerに持ち込まないようにproductに変換
	product := &entity.Product{}
	if err := c.Bind(product); err != nil {
		myErr.Logging(err)
		return
	}
	order := &entity.Order{Product: *product}

	pc.Controller.Reserve(ctx, order, orderType) // オーダー番号発行
	pc.Controller.Order(&ctx, order)             // オーダー
	c.JSON(200, order.OrderInfo.OrderNumber)     // オーダー番号返却
}
