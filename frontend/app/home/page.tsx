"use client"

import { useState } from "react"
import { useGetProductsQuery, useGetCategoriesQuery, useGetBrandsQuery } from "@/lib/api"
import type { Product, Category, Brand } from "@/lib/types"
import { ShoppingCart, Search, PanelRightClose, PanelLeftClose, ChevronRight, ArrowLeft } from "lucide-react"
import Link from "next/link"

// Mock nested categories structure - you can replace this with your API data
const nestedCategories = [
  {
    id: 1,
    name: "Laptops & Computers",
    icon: "💻",
    subcategories: [
      {
        id: 11,
        name: "Gaming Laptops",
        subcategories: [
          { id: 111, name: "RTX 4070 Series" },
          { id: 112, name: "RTX 4080 Series" },
          { id: 113, name: "RTX 4090 Series" },
        ],
      },
      {
        id: 12,
        name: "Business Laptops",
        subcategories: [
          { id: 121, name: "ThinkPad Series" },
          { id: 122, name: "Dell Latitude" },
          { id: 123, name: "HP EliteBook" },
        ],
      },
      {
        id: 13,
        name: "Desktop PCs",
        subcategories: [
          { id: 131, name: "Gaming PCs" },
          { id: 132, name: "Workstations" },
          { id: 133, name: "Mini PCs" },
        ],
      },
    ],
  },
  {
    id: 2,
    name: "Smartphones",
    icon: "📱",
    subcategories: [
      {
        id: 21,
        name: "iPhone",
        subcategories: [
          { id: 211, name: "iPhone 15 Series" },
          { id: 212, name: "iPhone 14 Series" },
          { id: 213, name: "iPhone 13 Series" },
        ],
      },
      {
        id: 22,
        name: "Samsung Galaxy",
        subcategories: [
          { id: 221, name: "Galaxy S Series" },
          { id: 222, name: "Galaxy Note Series" },
          { id: 223, name: "Galaxy A Series" },
        ],
      },
      {
        id: 23,
        name: "Google Pixel",
        subcategories: [
          { id: 231, name: "Pixel 8 Series" },
          { id: 232, name: "Pixel 7 Series" },
        ],
      },
    ],
  },
  {
    id: 3,
    name: "PC Components",
    icon: "🔧",
    subcategories: [
      {
        id: 31,
        name: "Graphics Cards",
        subcategories: [
          { id: 311, name: "NVIDIA RTX 40 Series" },
          { id: 312, name: "NVIDIA RTX 30 Series" },
          { id: 313, name: "AMD RX 7000 Series" },
        ],
      },
      {
        id: 32,
        name: "Processors",
        subcategories: [
          { id: 321, name: "Intel Core i9" },
          { id: 322, name: "Intel Core i7" },
          { id: 323, name: "AMD Ryzen 9" },
        ],
      },
      {
        id: 33,
        name: "Motherboards",
        subcategories: [
          { id: 331, name: "Intel Z790" },
          { id: 332, name: "AMD X670" },
          { id: 333, name: "Budget Boards" },
        ],
      },
    ],
  },
  {
    id: 4,
    name: "Audio & Accessories",
    icon: "🎧",
    subcategories: [
      {
        id: 41,
        name: "Headphones",
        subcategories: [
          { id: 411, name: "Wireless" },
          { id: 412, name: "Gaming" },
          { id: 413, name: "Studio" },
        ],
      },
      {
        id: 42,
        name: "Speakers",
        subcategories: [
          { id: 421, name: "Bluetooth" },
          { id: 422, name: "PC Speakers" },
          { id: 423, name: "Soundbars" },
        ],
      },
    ],
  },
]

function NestedCategorySidebar({ onCategorySelect }: { onCategorySelect: (categoryName: string) => void }) {
  const [selectedCategory, setSelectedCategory] = useState<any>(null)
  const [selectedSubcategory, setSelectedSubcategory] = useState<any>(null)
  const [showThirdLevel, setShowThirdLevel] = useState(false)

  const handleCategoryClick = (category: any) => {
    setSelectedCategory(category)
    setSelectedSubcategory(null)
    setShowThirdLevel(false)
  }

  const handleSubcategoryClick = (subcategory: any) => {
    if (subcategory.subcategories) {
      setSelectedSubcategory(subcategory)
      setShowThirdLevel(true)
    } else {
      onCategorySelect(subcategory.name)
    }
  }

  const handleThirdLevelClick = (item: any) => {
    onCategorySelect(item.name)
  }

  const handleBackClick = () => {
    setShowThirdLevel(false)
    setSelectedSubcategory(null)
  }

  return (
    <div className="h-screen bg-white">
      <div className="flex h-full">
        {/* First Level */}
        <div className={`w-72 bg-white shadow-xl border-r transition-all duration-300 ${showThirdLevel ? "opacity-50" : ""}`}>
          <div className="p-4 border-b bg-gradient-to-r from-indigo-600 to-purple-600">
            <h3 className="font-semibold text-white">Categories</h3>
          </div>
          <div className="p-2 h-[calc(100vh-4rem)] overflow-y-auto">
            {nestedCategories.map((category) => (
              <button
                key={category.id}
                onClick={() => handleCategoryClick(category)}
                className={`w-full text-left p-4 rounded-lg hover:bg-gradient-to-r hover:from-blue-50 hover:to-indigo-50 flex items-center justify-between transition-all duration-200 ${
                  selectedCategory?.id === category.id
                    ? "bg-gradient-to-r from-blue-100 to-indigo-100 text-indigo-700 shadow-md"
                    : ""
                }`}
              >
                <div className="flex items-center">
                  <span className="mr-3 text-2xl">{category.icon}</span>
                  <span className="text-sm font-medium">{category.name}</span>
                </div>
                <ChevronRight className="h-4 w-4 text-indigo-500" />
              </button>
            ))}
          </div>
        </div>

        {/* Second Level */}
        {selectedCategory && (
          <div className={`w-72 bg-gradient-to-b from-gray-50 to-white shadow-xl border-r transition-all duration-300 ${showThirdLevel ? "-ml-72 z-10 relative" : ""}`}>
            {showThirdLevel && (
              <div className="p-4 border-b bg-gradient-to-r from-purple-600 to-pink-600">
                <button
                  onClick={handleBackClick}
                  className="flex items-center text-white hover:text-purple-100 transition-colors"
                >
                  <ArrowLeft className="h-4 w-4 mr-2" />
                  Back to Categories
                </button>
              </div>
            )}
            <div className="p-4 border-b bg-gradient-to-r from-indigo-500 to-purple-500">
              <h4 className="font-semibold text-white">{selectedCategory.name}</h4>
            </div>
            <div className="p-2 h-[calc(100vh-4rem)] overflow-y-auto">
              {selectedCategory.subcategories?.map((subcategory: any) => (
                <button
                  key={subcategory.id}
                  onClick={() => handleSubcategoryClick(subcategory)}
                  className={`w-full text-left p-4 rounded-lg hover:bg-gradient-to-r hover:from-purple-50 hover:to-pink-50 flex items-center justify-between transition-all duration-200 group ${
                    selectedSubcategory?.id === subcategory.id
                      ? "bg-gradient-to-r from-purple-100 to-pink-100 text-purple-700 shadow-md"
                      : ""
                  }`}
                >
                  <span className="text-sm font-medium">{subcategory.name}</span>
                  {subcategory.subcategories ? (
                    <ChevronRight className="h-4 w-4 text-purple-500 group-hover:text-purple-700" />
                  ) : (
                    <div className="w-2 h-2 bg-green-500 rounded-full"></div>
                  )}
                </button>
              ))}
            </div>
          </div>
        )}

        {/* Third Level */}
        {showThirdLevel && selectedSubcategory && (
          <div className="w-72 bg-white shadow-xl border-r">
            <div className="p-4 border-b bg-gradient-to-r from-pink-500 to-red-500">
              <h4 className="font-semibold text-white">{selectedSubcategory.name}</h4>
            </div>
            <div className="p-2 h-[calc(100vh-4rem)] overflow-y-auto">
              {selectedSubcategory.subcategories?.map((item: any) => (
                <button
                  key={item.id}
                  onClick={() => handleThirdLevelClick(item)}
                  className="w-full text-left p-4 rounded-lg hover:bg-gradient-to-r hover:from-red-50 hover:to-pink-50 transition-all duration-200 flex items-center justify-between group"
                >
                  <span className="text-sm font-medium">{item.name}</span>
                  <div className="w-2 h-2 bg-red-500 rounded-full group-hover:bg-red-600"></div>
                </button>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

export default function HomePage() {
  const [showSidebar, setShowSidebar] = useState(false)
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedCategory, setSelectedCategory] = useState<number | null>(null)
  const [selectedBrand, setSelectedBrand] = useState<number | null>(null)
  const [categorySearchQuery, setCategorySearchQuery] = useState<string>("")

  const [cartItems, setCartItems] = useState<any[]>([])

  const handleAddToCart = (product: Product) => {
    setCartItems((prev) => {
      const existing = prev.find((item) => item.product_id === product.product_id)
      if (existing) {
        return prev.map((item) =>
          item.product_id === product.product_id ? { ...item, quantity: item.quantity + 1 } : item,
        )
      }
      return [...prev, { ...product, quantity: 1 }]
    })
  }

  const getTotalCartItems = () => {
    return cartItems.reduce((sum, item) => sum + item.quantity, 0)
  }

  const { data: productsData, isLoading: isLoadingProducts } = useGetProductsQuery()
  const { data: categoriesData, isLoading: isLoadingCategories } = useGetCategoriesQuery()
  const { data: brandsData, isLoading: isLoadingBrands } = useGetBrandsQuery()

  console.log("productsData", productsData)
  const products = productsData?.data.products || []
  const categories = categoriesData?.data.categories || []
  const brands = brandsData?.data.brands || []

  const handleCategorySelect = async (categoryName: string) => {
    setCategorySearchQuery(categoryName)
    setShowSidebar(false) // Close sidebar after selection

    // Here you can make an API call to search for products in this category
    // For example:
    // try {
    //   const response = await fetch(`/api/products/search?category=${encodeURIComponent(categoryName)}`)
    //   const data = await response.json()
    //   // Handle the response data
    // } catch (error) {
    //   console.error('Error searching products:', error)
    // }

    console.log(`Searching for products in category: ${categoryName}`)
  }

  const filteredProducts = products.filter((product: Product) => {
    const matchesSearch =
      product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      product.description.toLowerCase().includes(searchQuery.toLowerCase())
    const matchesCategory = selectedCategory ? product.categroy_id === selectedCategory : true
    const matchesBrand = selectedBrand ? product.brand_id === selectedBrand : true
    const matchesCategorySearch = categorySearchQuery
      ? product.name.toLowerCase().includes(categorySearchQuery.toLowerCase()) ||
        product.description.toLowerCase().includes(categorySearchQuery.toLowerCase())
      : true

    return matchesSearch && matchesCategory && matchesBrand && matchesCategorySearch
  })

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-blue-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b sticky top-0 z-40">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <button
                onClick={() => setShowSidebar(!showSidebar)}
                className="p-2 rounded-md hover:bg-gradient-to-r hover:from-indigo-50 hover:to-purple-50 transition-all duration-200"
              >
                {showSidebar ? (
                  <PanelLeftClose className="h-6 w-6 text-indigo-600" />
                ) : (
                  <PanelRightClose className="h-6 w-6 text-indigo-600" />
                )}
              </button>
              <Link
                href="/"
                className="text-2xl font-bold bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent"
              >
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
                  className="pl-10 pr-4 py-2 border border-indigo-200 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                />
                <Search className="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
              </div>
              <Link href="/cart" className="relative p-2 hover:bg-indigo-50 rounded-md transition-colors">
                <ShoppingCart className="h-6 w-6 text-indigo-600" />
                {getTotalCartItems() > 0 && (
                  <span className="absolute -top-1 -right-1 bg-gradient-to-r from-red-500 to-pink-500 text-white rounded-full w-5 h-5 flex items-center justify-center text-xs">
                    {getTotalCartItems()}
                  </span>
                )}
              </Link>
            </div>
          </div>
        </div>
      </header>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="relative">
          {/* Sidebar Overlay */}
          {showSidebar && (
            <div className="fixed inset-0 z-50">
              <div className="absolute inset-0 bg-black bg-opacity-50" onClick={() => setShowSidebar(false)} />
              <div className="absolute inset-y-0 left-0 h-screen">
                <NestedCategorySidebar onCategorySelect={handleCategorySelect} />
              </div>
            </div>
          )}

          {/* Main Content */}
          <div>
            {/* Active Filters Display */}
            {(categorySearchQuery || selectedCategory || selectedBrand) && (
              <div className="mb-6 flex flex-wrap gap-2">
                <span className="text-sm text-gray-600">Active filters:</span>
                {categorySearchQuery && (
                  <span className="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-gradient-to-r from-indigo-100 to-purple-100 text-indigo-700">
                    Category: {categorySearchQuery}
                    <button
                      onClick={() => setCategorySearchQuery("")}
                      className="ml-2 text-indigo-500 hover:text-indigo-700"
                    >
                      ×
                    </button>
                  </span>
                )}
                {selectedCategory && (
                  <span className="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-gradient-to-r from-blue-100 to-indigo-100 text-blue-700">
                    {categories.find((c) => c.category_id === selectedCategory)?.category_name}
                    <button
                      onClick={() => setSelectedCategory(null)}
                      className="ml-2 text-blue-500 hover:text-blue-700"
                    >
                      ×
                    </button>
                  </span>
                )}
                {selectedBrand && (
                  <span className="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-gradient-to-r from-green-100 to-emerald-100 text-green-700">
                    {brands.find((b) => b.brand_id === selectedBrand)?.name}
                    <button onClick={() => setSelectedBrand(null)} className="ml-2 text-green-500 hover:text-green-700">
                      ×
                    </button>
                  </span>
                )}
              </div>
            )}

            {isLoadingProducts ? (
              <div className="text-center py-8">
                <div className="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
                <p className="mt-2 text-gray-600">Loading products...</p>
              </div>
            ) : (
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {filteredProducts.map((product: Product) => (
                  <div
                    key={product.product_id}
                    onClick={() => (window.location.href = `/products/${product.product_id}`)}
                    className="bg-white rounded-lg shadow-lg overflow-hidden hover:shadow-xl transition-all duration-300 border border-gray-100 hover:scale-105 cursor-pointer"
                  >
                    <div className="aspect-w-1 aspect-h-1">
                      <img src={product.image || "/pc1.jpg"} alt={product.name} className="w-full h-48 object-cover" />
                    </div>
                    <div className="p-4">
                      <h3 className="text-lg font-semibold text-gray-900 line-clamp-2 min-h-[3rem]">{product.name}</h3>
                      <p className="mt-1 text-sm text-gray-500 line-clamp-2">{product.description}</p>
                      <div className="mt-4 flex items-center justify-between">
                        <span className="text-xl font-bold bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent">
                          ${product.price.toFixed(2)}
                        </span>
                        <button
                          onClick={(e) => {
                            e.stopPropagation() // Prevent card click
                            handleAddToCart(product)
                          }}
                          className="bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 text-white px-4 py-2 rounded-md transition-all duration-200 shadow-md hover:shadow-lg"
                        >
                          Add to Cart
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}

            {filteredProducts.length === 0 && !isLoadingProducts && (
              <div className="text-center py-16">
                <div className="w-24 h-24 bg-gradient-to-r from-indigo-100 to-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Search className="h-12 w-12 text-indigo-400" />
                </div>
                <h2 className="text-2xl font-semibold text-gray-900 mb-2">No products found</h2>
                <p className="text-gray-600 mb-8">Try adjusting your search or filters</p>
                <button
                  onClick={() => {
                    setSearchQuery("")
                    setSelectedCategory(null)
                    setSelectedBrand(null)
                    setCategorySearchQuery("")
                  }}
                  className="bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 text-white px-6 py-2 rounded-md transition-all duration-200"
                >
                  Clear All Filters
                </button>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
