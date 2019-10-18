import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { ProductListComponent } from './products/product-list/product-list.component';
import { ProductCreateComponent } from './products/product-create/product-create.component';
import { ProductUpdateComponent } from './products/product-update/product-update.component';

import { CustomerListComponent } from './customers/customer-list/customer-list.component';
import { CustomerCreateComponent } from './customers/customer-create/customer-create.component';
import { CustomerUpdateComponent } from './customers/customer-update/customer-update.component';

import { OrderListComponent } from './orders/order-list/order-list.component';
import { OrderCreateComponent } from './orders/order-create/order-create.component';
import { OrderUpdateComponent } from './orders/order-update/order-update.component';

const routes: Routes = [
  {
    path: 'products',
    component: ProductListComponent
  },
  {
    path: 'product-create',
    component: ProductCreateComponent
  },
  {
    path: 'products/:id',
    component: ProductUpdateComponent
  },
  {
    path: 'customers',
    component: CustomerListComponent
  },
  {
    path: 'customer-create',
    component: CustomerCreateComponent
  },
  {
    path: 'customers/:id',
    component: CustomerUpdateComponent
  },
  {
    path: 'orders',
    component: OrderListComponent
  },
  {
    path: 'order-create',
    component: OrderCreateComponent
  },
  {
    path: 'orders/:id',
    component: OrderUpdateComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
