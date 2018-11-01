# json-go-grpc
1. 将json转为grpc的定义，目前只实现json转为message的定义
2. 该项目基于我的另一个项目 https://github.com/GrayOxygen/json-to-grpc
</b>
TODO 后续考虑如何生成完整的grpc(方法的定义等)

## Example
input
## JSON 
```
{
  "items": [
    {
      "order_number": "614318762004012957",
      "order_source": "po",
      "sender_name": "jd",
      "order_status": "processing",
      "currency_code": "CNY",
      "total_qty_ordered": 1.0000,
      "total_item_count": "1",
      "total_weight": 100,
      "discount_amount": "0.00",
      "shipping_amount": "0.00",
      "subtotal": "79.9900",
      "grand_total": "79.9900",
      "created_at": "2015-12-11 22:51:53",
      "updated_at": "2015-12-22 20:14:18",
      "paid_at": "2015-12-22 20:14:18",
      "payment_type": "alipay_payment",
      "shipping_address_info": {
        "email": "416757228@qq.com",
        "name": "常璐",
        "telephone": "18687079066",
        "province": "北京市",
        "city": "北京市",
        "county": "海淀区",
        "street": "西土城路25号中国政法大学研究生院",
        "postcode": "100088",
        "id_card": "431102199603656899",
        "order_items": [
          {
            "sku": "LANCOSC73978802",
            "weight": "100.0000",
            "price": "79.9900",
            "qty_ordered": "1.0000",
            "name": "Lanc?me Génifique Advanced Youth Activating Concentrate 75ml",
            "item_discount_amount": "0.00"
          }
        ]
      },
      "order_items": [
        {
          "sku": "LANCOSC73978802",
          "weight": "100.0000",
          "price": "79.9900",
          "qty_ordered": "1",
          "name": "Lanc?me Génifique Advanced Youth Activating Concentrate 75ml",
          "item_discount_amount": "0.00"
        }
      ]
    },
    {
      "order_number": "614318762004012951",
      "order_source": "po",
      "sender_name": "jd",
      "order_status": "processing",
      "currency_code": "CNY",
      "total_qty_ordered": 1,
      "total_item_count": "1",
      "total_weight": 100,
      "discount_amount": "0.00",
      "shipping_amount": "0.00",
      "subtotal": "79.9900",
      "grand_total": "79.9900",
      "created_at": "2015-12-11 22:51:53",
      "updated_at": "2015-12-22 20:14:18",
      "paid_at": "2015-12-22 20:14:18",
      "payment_type": "vt_payment",
      "shipping_address_info": {
        "email": "416757228@qq.com",
        "name": "常璐",
        "telephone": "18687079066",
        "province": "北京市",
        "city": "北京市",
        "county": "海淀区",
        "street": "西土城路25号中国政法大学研究生院",
        "postcode": "100088",
        "id_card": "431102199603656899"
      },
      "order_items": [
        {
          "sku": "LANCOSC73978802",
          "weight": "100.0000",
          "price": "79.9900",
          "qty_ordered": "1.0000",
          "name": "Lanc?me Génifique Advanced Youth Activating Concentrate 75ml",
          "item_discount_amount": "0.00"
        }
      ]
    }
  ]
}
```
output
## GRPC Message
``` 
syntax = "proto3";
package protobuf;

message  Object   {
	repeated Items items  =1;
}
message Items   {
	float totalQtyOrdered=1;
	string totalItemCount=2;
	string discountAmount=3;
	string grandTotal=4;
	string paidAt=5;
	string orderNumber=6;
	string orderStatus=7;
	string currencyCode=8;
	string updatedAt=9;
	string paymentType=10;
	ShippingAddressInfo shippingAddressInfo  =11;
	string senderName=12;
	string shippingAmount=13;
	string createdAt=14;
	string subtotal=15;
	repeated ItemsOrderItems itemsOrderItems  =16;
	string orderSource=17;
	int32 totalWeight=18;
}
message ShippingAddressInfo   {
	string name=1;
	string province=2;
	string street=3;
	string postcode=4;
	repeated OrderItems orderItems  =5;
	string email=6;
	string telephone=7;
	string city=8;
	string county=9;
	string iDCard=10;
}
message OrderItems   {
	string sku=1;
	string weight=2;
	string price=3;
	string qtyOrdered=4;
	string name=5;
	string itemDiscountAmount=6;
}
message ItemsOrderItems   {
	string sku=1;
	string weight=2;
	string price=3;
	string qtyOrdered=4;
	string name=5;
	string itemDiscountAmount=6;
}
	

```


> I used https://github.com/mholt/json-to-go to get nested golang struct , nice project ,  much appreciated!!!
