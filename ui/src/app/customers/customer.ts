export class Phone {
    home: string;
    mobile: string;
    work: string;
}

export class Address {
    isshippingaddress: boolean;
    streetaddress: string;
    suburb: string;
    postcode: number;
    city: string;
    state: string;
    country: string;
    countycode: string;
}

export class Customer {
    id: number;
    email: string;
    firstname: string;
    lastname: string;
    middlename: string;
    title: string;
    phone: Phone;
    addresses: Address[];
    createdat: Date;
    updatedat: Date;
}
