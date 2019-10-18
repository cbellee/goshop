import { Component, OnInit } from "@angular/core";
import { ActivatedRoute, Router } from "@angular/router";
import { OrderAPIService } from "../../api/order.api.service";
import { CustomerAPIService } from "../../api/customer.api.service";
import { ProductAPIService } from "../../api/product.api.service";
import { Order } from "../order";
import { Customer } from "../../customers/customer";
import { Product } from "../../products/product";

@Component({
  selector: "app-order-update",
  templateUrl: "./order-update.component.html",
  styleUrls: ["./order-update.component.scss"]
})
export class OrderUpdateComponent implements OnInit {
  constructor(
    private orderApiService: OrderAPIService,
    private customerApiService: CustomerAPIService,
    private productApiService: ProductAPIService,
    private route: ActivatedRoute
  ) {}

  order: Order;
  customer: Customer;
  product: Product;
  loadedOrder: {};
  orderId: any = null;
  customerId: number = null;

  ngOnInit() {
    this.getData();
  }

  getData() {
    this.orderId = this.route.snapshot.paramMap.get("id");

    console.log("order id:" + this.orderId);
    this.orderApiService.getOrder(this.orderId).subscribe(order => {
      this.customerApiService
        .getCustomer(order.customerid)
        .subscribe(customer => {
          this.customer = customer;
          this.order = order;
          this.order.total = 0;
          this.order.items.forEach(item => {
              item.subTotal = item.quantity * item.unitprice;
              this.order.total += item.subTotal;
          });
          console.log("order total: " + this.order.total);
        });
    });
  }
}
