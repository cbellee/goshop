import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ProductAPIService } from '../../api/product.api.service';
import { Product } from '../product';

@Component({
  selector: 'app-product-update',
  templateUrl: './product-update.component.html',
  styleUrls: ['./product-update.component.scss']
})

export class ProductUpdateComponent implements OnInit {

  constructor(
    private apiService: ProductAPIService,
    private route: ActivatedRoute
  ) { }

  product: Product;
  id = null;

  ngOnInit() {
    console.log(this.route.snapshot.paramMap.keys);
    this.id = this.route.snapshot.paramMap.get('id');

    this.apiService.getProduct(this.id)
      .subscribe(product => this.product = product);
  }
}
