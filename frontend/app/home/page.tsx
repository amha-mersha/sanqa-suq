"use client"

import { useState } from "react"
import { useGetProductsQuery, useGetCategoriesQuery, useGetBrandsQuery } from "@/lib/api"
import { Product, Category, Brand } from "@/lib/types"
import { ShoppingCart, Search, PanelRightClose, PanelLeftClose } from "lucide-react"
import Link from "next/link"

export default function HomePage() {
  const [showSidebar, setShowSidebar] = useState(false)
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedCategory, setSelectedCategory] = useState<number | null>(null)
  const [selectedBrand, setSelectedBrand] = useState<number | null>(null)

  const { data: productsData, isLoading: isLoadingProducts } = useGetProductsQuery()
  const { data: categoriesData, isLoading: isLoadingCategories } = useGetCategoriesQuery()
  const { data: brandsData, isLoading: isLoadingBrands } = useGetBrandsQuery()

  const products = productsData?.data.products || []
  const categories = categoriesData?.data.categories || []
  const brands = brandsData?.data.brands || []

  const filteredProducts = products.filter((product: Product) => {
    const matchesSearch = product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                         product.description.toLowerCase().includes(searchQuery.toLowerCase())
    const matchesCategory = selectedCategory ? product.categroy_id === selectedCategory : true
    const matchesBrand = selectedBrand ? product.brand_id === selectedBrand : true
    return matchesSearch && matchesCategory && matchesBrand
  })

  return (
    <div className="min-h-screen bg-gray-100">
      {/* Header */}
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <button
                onClick={() => setShowSidebar(!showSidebar)}
                className="p-2 rounded-md hover:bg-gray-100"
              >
                {showSidebar ? <PanelLeftClose className="h-6 w-6" /> : <PanelRightClose className="h-6 w-6" />}
              </button>
              <Link href="/" className="text-2xl font-bold text-gray-900">
                SanqaSuq
              </Link>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative">
                <input
                  type="text"
                  placeholder="Search products..."
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                  className="pl-10 pr-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
                <Search className="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
              </div>
              <Link href="/cart" className="relative p-2">
                <ShoppingCart className="h-6 w-6" />
                <span className="absolute -top-1 -right-1 bg-blue-500 text-white rounded-full w-5 h-5 flex items-center justify-center text-xs">
                  0
                </span>
              </Link>
            </div>
          </div>
        </div>
      </header>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex">
          {/* Sidebar */}
          {showSidebar && (
            <div className="w-64 pr-8">
              <div className="bg-white rounded-lg shadow p-4">
                <h2 className="text-lg font-semibold mb-4">Categories</h2>
                <div className="space-y-2">
                  <button
                    onClick={() => setSelectedCategory(null)}
                    className={`block w-full text-left px-2 py-1 rounded ${
                      selectedCategory === null ? 'bg-blue-100 text-blue-700' : 'hover:bg-gray-100'
                    }`}
                  >
                    All Categories
                  </button>
                  {categories.map((category: Category) => (
                    <button
                      key={category.category_id}
                      onClick={() => setSelectedCategory(category.category_id)}
                      className={`block w-full text-left px-2 py-1 rounded ${
                        selectedCategory === category.category_id ? 'bg-blue-100 text-blue-700' : 'hover:bg-gray-100'
                      }`}
                    >
                      {category.category_name}
                    </button>
                  ))}
                </div>

                <h2 className="text-lg font-semibold mt-6 mb-4">Brands</h2>
                <div className="space-y-2">
                  <button
                    onClick={() => setSelectedBrand(null)}
                    className={`block w-full text-left px-2 py-1 rounded ${
                      selectedBrand === null ? 'bg-blue-100 text-blue-700' : 'hover:bg-gray-100'
                    }`}
                  >
                    All Brands
                  </button>
                  {brands.map((brand: Brand) => (
                    <button
                      key={brand.brand_id}
                      onClick={() => setSelectedBrand(brand.brand_id)}
                      className={`block w-full text-left px-2 py-1 rounded ${
                        selectedBrand === brand.brand_id ? 'bg-blue-100 text-blue-700' : 'hover:bg-gray-100'
                      }`}
                    >
                      {brand.name}
                    </button>
                  ))}
                </div>
              </div>
            </div>
          )}

          {/* Main Content */}
          <div className={`flex-1 ${showSidebar ? 'ml-8' : ''}`}>
            {isLoadingProducts ? (
              <div className="text-center py-8">Loading products...</div>
            ) : (
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {filteredProducts.map((product: Product) => (
                  <div key={product.product_id} className="bg-white rounded-lg shadow overflow-hidden">
                    <div className="aspect-w-1 aspect-h-1">
                      <img
                        src={product.image || '/pc1.jpg'}
                        alt={product.name}
                        className="w-full h-48 object-cover"
                      />
                    </div>
                    <div className="p-4">
                      <h3 className="text-lg font-semibold text-gray-900">{product.name}</h3>
                      <p className="mt-1 text-sm text-gray-500 line-clamp-2">{product.description}</p>
                      <div className="mt-4 flex items-center justify-between">
                        <span className="text-lg font-bold text-gray-900">${product.price.toFixed(2)}</span>
                        <button
                          className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 transition-colors"
                        >
                          Add to Cart
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
