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
    signup: builder.mutation<any, {
      first_name: string;
      last_name: string;
      email: string;
      password: string;
      phone: string;
      role: "customer" | "seller";
      provider: string;
      provider_id: string;
    }>({
      query: (userData) => ({
        url: 'user/signup',
        method: 'POST',
        body: userData,
      }),
    }),
    login: builder.mutation<any, {
      email: string;
      password: string;
    }>({
      query: (credentials) => ({
        url: 'user/login',
        method: 'POST',
        body: credentials,
      }),
    }),
  }),
});

export const {
  useGetProductsQuery,
  useGetProductsByCategoryQuery,
  useGetCategoriesQuery,
  useGetBrandsQuery,
  useAddProductMutation,
  useSignupMutation,
  useLoginMutation,
} = api; 