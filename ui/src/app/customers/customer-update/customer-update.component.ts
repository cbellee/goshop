import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { CustomerAPIService } from '../../api/customer.api.service';
import { Customer } from '../../customers/customer';

@Component({
  selector: 'app-customer-update',
  templateUrl: './customer-update.component.html',
  styleUrls: ['./customer-update.component.scss']
})

export class CustomerUpdateComponent implements OnInit {

  constructor(
    private customerApiService: CustomerAPIService,
    private route: ActivatedRoute
  ) { }

  customer: Customer;
  customerId = null;

  ngOnInit() {
    this.getData();
  }

  getData() {
    this.customerId = this.route.snapshot.paramMap.get('id');

    console.log(this.customerId);

    this.customerApiService.getCustomer(this.customerId)
      .subscribe(customer => {
        this.customer = customer;
        console.log(this.customer);
      });
  }
}
