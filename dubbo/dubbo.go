package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/dubbogo/triple/pkg/triple"
	"github.com/fastmall/cart/service"
	customerApi "github.com/fastmall/customer/api"
	goodsApi "github.com/fastmall/goods/api"
	orderApi "github.com/fastmall/order/api"
)

var GoodsService = &goodsApi.GoodsServiceClientImpl{}
var CustomerService = &customerApi.CustomerServiceClientImpl{}
var OrderService = &orderApi.OrderServiceClientImpl{}

func StartDubbo() {
	config.SetConsumerService(GoodsService)
	config.SetConsumerService(CustomerService)
	config.SetConsumerService(OrderService)
	config.SetProviderService(&service.CartService{})
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
