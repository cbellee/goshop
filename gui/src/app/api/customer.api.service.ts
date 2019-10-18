import { Injectable, Inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Customer } from '../customers/customer';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})

export class CustomerAPIService {
  apiUrl = environment.customerApiUrl;
  constructor(private httpClient: HttpClient) { }

  getCustomers(): Observable<Customer[]> {
    return this.httpClient.get<Customer[]>(`${this.apiUrl}/customer`);
  }

  getCustomer(id: number): Observable<Customer> {
    return this.httpClient.get<Customer>(`${this.apiUrl}/customer/${id}`);
  }
}
