export class OrderItem {
    id: number;
    quantity: number;
    format: string;
    unitprice: number;
    isbackordered: boolean;
    subTotal: number;
}

export class Order {
    id: number;
    customerid: number;
    items: OrderItem[];
    createdat: Date;
    updatedat: Date;
    total: number;
}
