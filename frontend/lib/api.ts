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
    getCategories: builder.query<CategoriesResponse, void>({
      query: () => 'category',
      providesTags: ['Category'],
    }),
    getBrands: builder.query<BrandsResponse, void>({
      query: () => 'brand',
      providesTags: ['Brand'],
    }),
  }),
});

export const {
  useGetProductsQuery,
  useGetCategoriesQuery,
  useGetBrandsQuery,
} = api; 