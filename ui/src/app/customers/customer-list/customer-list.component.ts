import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material';
import { CustomerAPIService } from '../../api/customer.api.service';
import { Customer } from '../customer';

@Component({
  selector: 'app-customer-list',
  templateUrl: './customer-list.component.html',
  styleUrls: ['./customer-list.component.scss']
})

export class CustomerListComponent implements OnInit {

  CustomersDataSource: any;
  displayedColumns = ['firstName', 'lastName', 'email'];
  constructor(private apiService: CustomerAPIService) { }

  ngOnInit() {
    this.getCustomers();
  }

  public getCustomers() {
    this.apiService.getCustomers().subscribe((customerData: Array<Customer>) => {
      this.CustomersDataSource = new MatTableDataSource();
      this.CustomersDataSource.data = customerData;
      console.log(customerData);
    });
  }
}
