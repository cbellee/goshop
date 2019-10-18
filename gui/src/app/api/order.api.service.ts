import { Injectable, Inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Order } from '../orders/order';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})

export class OrderAPIService {
  apiUrl = environment.orderApiUrl;
  constructor(private httpClient: HttpClient) { }

  getOrders(): Observable<Order[]> {
    return this.httpClient.get<Order[]>(`${this.apiUrl}/order`);
  }

  getOrder(id: number): Observable<Order> {
    return this.httpClient.get<Order>(`${this.apiUrl}/order/${id}`);
  }
}
