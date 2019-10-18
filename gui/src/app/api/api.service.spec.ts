import { TestBed } from '@angular/core/testing';
import { ProductAPIService } from './product.api.service';
import { CustomerAPIService } from './customer.api.service';
import { OrderAPIService } from './order.api.service';

describe('ProductAPIService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: ProductAPIService = TestBed.get(ProductAPIService);
    expect(service).toBeTruthy();
  });
});

describe('CustomerAPIService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: CustomerAPIService = TestBed.get(CustomerAPIService);
    expect(service).toBeTruthy();
  });
});

describe('OrderAPIService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: OrderAPIService = TestBed.get(OrderAPIService);
    expect(service).toBeTruthy();
  });
});
