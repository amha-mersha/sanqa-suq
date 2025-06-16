export interface Product {
  product_id: number;
  categroy_id: number; // Note: Keeping the typo as it comes from the API
  brand_id: number;
  name: string;
  description: string;
  price: number;
  stock_quantity: number;
  created_at: string;
  image?: string; // Making it optional since it might not always be present
}

export interface Category {
  category_id: number;
  category_name: string;
}

export interface Brand {
  brand_id: number;
  name: string;
}

export interface ProductsResponse {
  data: {
    products: Product[];
  };
  message: string;
}

export interface CategoriesResponse {
  data: {
    categories: Category[];
  };
  message: string;
}

export interface BrandsResponse {
  data: {
    brands: Brand[];
  };
  message: string;
} 