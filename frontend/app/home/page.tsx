"use client"

import { useState, useMemo } from "react"
import { useGetProductsQuery, useGetCategoriesQuery, useGetBrandsQuery, useGetProductsByCategoryQuery } from "@/lib/api"
import type { Product, Category, Brand } from "@/lib/types"
import { ShoppingCart, Search, PanelRightClose, PanelLeftClose, ChevronRight, ArrowLeft, Package } from "lucide-react"
import Link from "next/link"
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from "@/components/ui/select"
import Image from "next/image"

function NestedCategorySidebar({ onCategorySelect }: { onCategorySelect: (categoryId: number) => void }) {
  const [selectedCategory, setSelectedCategory] = useState<any>(null)
  const [selectedSubcategory, setSelectedSubcategory] = useState<any>(null)
  const [showThirdLevel, setShowThirdLevel] = useState(false)

  const { data: categoriesData } = useGetCategoriesQuery()
  const categories = categoriesData?.data || []
  console.log("categories", categoriesData)

  // Helper function to get random icon
  const getRandomIcon = () => {
    const icons = ["💻", "📱", "🔧", "🎧", "⌚", "📷", "🎮", "🖥️", "⌨️", "🖱️"]
    return icons[Math.floor(Math.random() * icons.length)]
  }

  // Transform flat categories into nested structure
  const buildCategoryTree = () => {
    const categoryMap = new Map()
    const rootCategories: any[] = []

    // First pass: Create all category objects
    categories.forEach(category => {
      categoryMap.set(category.category_id, {
        id: category.category_id,
        name: category.name,
        icon: getRandomIcon(),
        subcategories: []
      })
    })

    // Second pass: Build the tree structure
    categories.forEach(category => {
      const categoryObj = categoryMap.get(category.category_id)
      if (!category.parent_category_id) {
        rootCategories.push(categoryObj)
      } else {
        const parent = categoryMap.get(category.parent_category_id)
        if (parent) {
          parent.subcategories.push(categoryObj)
        }
      }
    })

    return rootCategories
  }

  const nestedCategories = buildCategoryTree()

  const handleCategoryClick = (category: any) => {
    setSelectedCategory(category)
    setSelectedSubcategory(null)
    setShowThirdLevel(false)
  }

  const handleSubcategoryClick = (subcategory: any) => {
    if (subcategory.subcategories && subcategory.subcategories.length > 0) {
      setSelectedSubcategory(subcategory)
      setShowThirdLevel(true)
    } else {
      onCategorySelect(subcategory.id)
    }
  }

  const handleThirdLevelClick = (item: any) => {
    onCategorySelect(item.id)
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
                  {subcategory.subcategories && subcategory.subcategories.length > 0 ? (
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

// Add this function before the HomePage component
const getRandomProductImage = () => {
  const images = [
    "/joshua-ng-1sSfrozgiFk-unsplash.jpg",
    "/kevin-bhagat-548zkUvqmlw-unsplash.jpg",
    "/itadaki-YWY_AASpmEI-unsplash.jpg",
    "/olivier-collet-JMwCe3w7qKk-unsplash.jpg",
    "/onur-binay-auf3GwpVaOM-unsplash.jpg",
    "/alienware-Hpaq-kBcYHk-unsplash.jpg",
    "/pc2.jpg",
    "/pc1.jpg"
  ]
  return images[Math.floor(Math.random() * images.length)]
}

export default function HomePage() {
  const [showSidebar, setShowSidebar] = useState(false)
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedBrand, setSelectedBrand] = useState<string>("all")
  const [priceRange, setPriceRange] = useState<string>("all")
  const [sortBy, setSortBy] = useState<string>("featured")
  const [categorySearchQuery, setCategorySearchQuery] = useState("")
  const [selectedCategoryId, setSelectedCategoryId] = useState<number | null>(null)

  const [cartItems, setCartItems] = useState<any[]>([])

  const handleAddToCart = (product: Product) => {
    setCartItems((prev) => {
      const existing = prev.find((item) => item.product_id === product.product_id)
      let updated;
      if (existing) {
        updated = prev.map((item) =>
          item.product_id === product.product_id ? { ...item, quantity: item.quantity + 1 } : item
        )
      } else {
        updated = [...prev, { product_id: product.product_id, quantity: 1 }]
      }
      // Save to localStorage
      localStorage.setItem("cart", JSON.stringify(updated))
      return updated
    })
  }

  const getTotalCartItems = () => {
    return cartItems.reduce((sum, item) => sum + item.quantity, 0)
  }

  const { data: productsData, isLoading: isLoadingProducts } = useGetProductsQuery()
  const { data: categoriesData } = useGetCategoriesQuery()
  const { data: brandsData } = useGetBrandsQuery()
  const { data: categoryProductsData, isLoading: isLoadingCategoryProducts } = useGetProductsByCategoryQuery(selectedCategoryId || 0, {
    skip: !selectedCategoryId
  })

  const products = selectedCategoryId ? categoryProductsData?.data.products || [] : productsData?.data.products || []
  const categories = categoriesData?.data.categories || []
  const brands = brandsData?.data.brands || []

  const handleCategorySelect = (categoryId: number) => {
    setSelectedCategoryId(categoryId)
    setCategorySearchQuery("")
  }

  const filteredProducts = useMemo(() => {
    return products
      .filter((product) => {
        const matchesSearch = product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          product.description.toLowerCase().includes(searchQuery.toLowerCase())
        const matchesBrand = selectedBrand === "all" || product.brand_id === parseInt(selectedBrand)
        const matchesCategory = !categorySearchQuery || 
          categories.find(c => c.category_id === product.categroy_id)?.category_name.toLowerCase().includes(categorySearchQuery.toLowerCase())

        // Price range filtering
        let matchesPriceRange = true
        if (priceRange !== "all") {
          const [min, max] = priceRange.split("-").map(Number)
          if (max) {
            matchesPriceRange = product.price >= min && product.price <= max
          } else {
            matchesPriceRange = product.price >= min
          }
        }

        return matchesSearch && matchesBrand && matchesCategory && matchesPriceRange
      })
      .sort((a, b) => {
        switch (sortBy) {
          case "price-low-high":
            return a.price - b.price
          case "price-high-low":
            return b.price - a.price
          case "name-a-z":
            return a.name.localeCompare(b.name)
          case "name-z-a":
            return b.name.localeCompare(a.name)
          default:
            return 0
        }
      })
  }, [products, searchQuery, selectedBrand, categorySearchQuery, categories, priceRange, sortBy])

  const isLoading = isLoadingProducts || (selectedCategoryId && isLoadingCategoryProducts)

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
            {(categorySearchQuery || selectedCategoryId) && (
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
                {selectedCategoryId && (
                  <span className="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-gradient-to-r from-blue-100 to-indigo-100 text-blue-700">
                    {categories.find((c) => c.category_id === selectedCategoryId)?.category_name}
                    <button
                      onClick={() => setSelectedCategoryId(null)}
                      className="ml-2 text-blue-500 hover:text-blue-700"
                    >
                      ×
                    </button>
                  </span>
                )}
              </div>
            )}

            {/* Filters */}
            <div className="mb-6 flex flex-wrap gap-4">
              <Select value={selectedBrand} onValueChange={setSelectedBrand}>
                <SelectTrigger className="w-[180px]">
                  <SelectValue placeholder="Select Brand" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All Brands</SelectItem>
                  {brands.map((brand) => (
                    <SelectItem key={brand.brand_id} value={brand.brand_id.toString()}>
                      {brand.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              <Select value={priceRange} onValueChange={setPriceRange}>
                <SelectTrigger className="w-[180px]">
                  <SelectValue placeholder="Price Range" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All Prices</SelectItem>
                  <SelectItem value="0-50">Under $50</SelectItem>
                  <SelectItem value="50-100">$50 - $100</SelectItem>
                  <SelectItem value="100-200">$100 - $200</SelectItem>
                  <SelectItem value="200-500">$200 - $500</SelectItem>
                  <SelectItem value="500-">$500 & Above</SelectItem>
                </SelectContent>
              </Select>

              <Select value={sortBy} onValueChange={setSortBy}>
                <SelectTrigger className="w-[180px]">
                  <SelectValue placeholder="Sort By" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="featured">Featured</SelectItem>
                  <SelectItem value="price-low-high">Price: Low to High</SelectItem>
                  <SelectItem value="price-high-low">Price: High to Low</SelectItem>
                  <SelectItem value="name-a-z">Name: A to Z</SelectItem>
                  <SelectItem value="name-z-a">Name: Z to A</SelectItem>
                </SelectContent>
              </Select>
            </div>

            {/* Products Grid */}
            {isLoading ? (
              <div className="text-center py-16">
                <div className="w-24 h-24 bg-gradient-to-r from-indigo-100 to-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
                </div>
                <h2 className="text-2xl font-semibold text-gray-900 mb-2">Loading products...</h2>
              </div>
            ) : filteredProducts.length === 0 ? (
              <div className="text-center py-16">
                <div className="w-24 h-24 bg-gradient-to-r from-indigo-100 to-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
                  <Search className="h-12 w-12 text-indigo-400" />
                </div>
                <h2 className="text-2xl font-semibold text-gray-900 mb-2">No products found</h2>
                <p className="text-gray-600 mb-8">Try adjusting your search or filters</p>
                <button
                  onClick={() => {
                    setSearchQuery("")
                    setSelectedCategoryId(null)
                    setCategorySearchQuery("")
                    setSelectedBrand("all")
                    setPriceRange("all")
                    setSortBy("featured")
                  }}
                  className="bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 text-white px-6 py-2 rounded-md transition-all duration-200"
                >
                  Clear All Filters
                </button>
              </div>
            ) : (
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {filteredProducts.map((product) => (
                  <div key={product.product_id} className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-200">
                    <div className="relative h-48 bg-gray-100">
                      <Image
                        src={getRandomProductImage()}
                        alt={product.name}
                        fill
                        className="object-cover"
                      />
                    </div>
                    <div className="p-4">
                      <div className="flex items-center justify-between mb-2">
                        <span className="text-sm text-gray-500">
                          {brands.find(b => b.brand_id === product.brand_id)?.name}
                        </span>
                        <span className="text-sm text-gray-500">
                          {categories.find(c => c.category_id === product.categroy_id)?.category_name}
                        </span>
                      </div>
                      <h3 className="text-lg font-semibold text-gray-900 mb-1">{product.name}</h3>
                      <p className="text-sm text-gray-600 mb-3 line-clamp-2">{product.description}</p>
                      <div className="flex items-center justify-between">
                        <span className="text-lg font-bold text-indigo-600">${product.price.toFixed(2)}</span>
                        <button
                          onClick={() => handleAddToCart(product)}
                          className="bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 text-white px-4 py-2 rounded-md transition-all duration-200"
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
