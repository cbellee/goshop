import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material';
import { ProductAPIService } from '../../api/product.api.service';
import { Product } from '../product';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})

export class ProductListComponent implements OnInit {

  ProductsDataSource: MatTableDataSource<Product>;
  displayedColumns = ['image', 'album', 'artist', 'releaseyear', 'price'];

  constructor(private apiService: ProductAPIService) { }

  ngOnInit() {
    this.getProducts();
  }

  public getProducts() {
    this.apiService.getProducts()
      .subscribe((productData: Array<Product>) => {
        this.ProductsDataSource = new MatTableDataSource();
        this.ProductsDataSource.data = productData;
        console.log(productData);
      });
  }
}
