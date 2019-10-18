export class Price {
    cd: number;
    lp: number;
    mp3: number;
}

export class Product {
    id: number;
    imageurl: string;
    artist: string;
    album: string;
    releaseyear: string;
    description: number;
    price: Price;
    createdat: Date;
    updatedat: Date;
}


