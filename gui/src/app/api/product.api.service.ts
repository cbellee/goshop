import { Injectable, Inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Product } from '../products/product';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})

export class ProductAPIService {
  apiUrl = environment.productApiUrl;
  constructor(private httpClient: HttpClient) { }

  getProducts(): Observable<Product[]> {
    return this.httpClient.get<Product[]>(`${this.apiUrl}/product`);
  }

  getProduct(id: number): Observable<Product> {
    return this.httpClient.get<Product>(`${this.apiUrl}/product/${id}`);
  }
}
