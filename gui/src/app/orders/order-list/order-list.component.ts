import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material';
import { OrderAPIService } from '../../api/order.api.service';
import { Order } from '../order';

@Component({
  selector: 'app-order-list',
  templateUrl: './order-list.component.html',
  styleUrls: ['./order-list.component.scss']
})

export class OrderListComponent implements OnInit {

  OrdersDataSource: MatTableDataSource<Order>;
  displayedColumns = ['id', 'customerid', 'createdat', 'updatedat', 'items'];

  constructor(private apiService: OrderAPIService) { }

  ngOnInit() {
    this.getOrders();
  }

  public getOrders() {
    this.apiService.getOrders()
      .subscribe((orderData: Array<Order>) => {
        this.OrdersDataSource = new MatTableDataSource();
        this.OrdersDataSource.data = orderData;
        console.log(orderData);
      });
  }
}
