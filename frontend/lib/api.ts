import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { ProductsResponse, CategoriesResponse, BrandsResponse } from './types';

export const api = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({ baseUrl: 'http://localhost:8080/api/v1' }),
  tagTypes: ['Product', 'Category', 'Brand'],
  endpoints: (builder) => ({
    getProducts: builder.query<ProductsResponse, void>({
      query: () => 'product',
      providesTags: ['Product'],
    }),
    getProductsByCategory: builder.query<ProductsResponse, number>({
      query: (categoryId) => `product/get-by-category/${categoryId}`,
      providesTags: ['Product'],
    }),
    getCategories: builder.query<CategoriesResponse, void>({
      query: () => 'categories',
      providesTags: ['Category'],
    }),
    getBrands: builder.query<BrandsResponse, void>({
      query: () => 'brand',
      providesTags: ['Brand'],
    }),
    addProduct: builder.mutation<any, {
      category_id: number;
      brand_id: number;
      name: string;
      description: string;
      price: number;
      stock_quantity: number;
    }>({
      query: (product) => ({
        url: 'product/add',
        method: 'POST',
        body: product,
      }),
      invalidatesTags: ['Product'],
    }),
  }),
});

export const {
  useGetProductsQuery,
  useGetProductsByCategoryQuery,
  useGetCategoriesQuery,
  useGetBrandsQuery,
  useAddProductMutation,
} = api; 