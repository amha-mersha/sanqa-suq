"use client"

import { useState } from "react"
import { useGetProductsQuery, useGetCategoriesQuery, useGetBrandsQuery } from "@/lib/api"
import type { Product, Category, Brand } from "@/lib/types"
import { ShoppingCart, Search, PanelRightClose, PanelLeftClose, ChevronRight, ArrowLeft } from "lucide-react"
import Link from "next/link"
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from "@/components/ui/select"
import { Card, CardContent } from "@/components/ui/card"
import Image from "next/image"
import { Button } from "@/components/ui/button"

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

function CategorySidebar() {
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
      // This is a final category - trigger search
      onCategorySelect(subcategory.name)
    }
  }

  const handleThirdLevelClick = (item: any) => {
    // This is a final category - trigger search
    onCategorySelect(item.name)
  }

  const handleBackClick = () => {
    setShowThirdLevel(false)
    setSelectedSubcategory(null)
  }

  return (
    <div className="h-full">
      <div className="flex h-full">
        {/* First Level - Always visible */}
        <div
          className={`w-72 bg-white shadow-xl border-r transition-all duration-300 ${showThirdLevel ? "opacity-50" : ""}`}
        >
          <div className="p-4 border-b bg-gradient-to-r from-indigo-600 to-purple-600">
            <h3 className="font-semibold text-white">Categories</h3>
          </div>
          <div className="p-2 max-h-screen overflow-y-auto">
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

        {/* Second Level - Shows when category selected */}
        {selectedCategory && (
          <div
            className={`w-72 bg-gradient-to-b from-gray-50 to-white shadow-xl border-r transition-all duration-300 ${showThirdLevel ? "-ml-72 z-10 relative" : ""}`}
          >
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
            <div className="p-2 max-h-screen overflow-y-auto">
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

        {/* Third Level - Shows when subcategory with children selected */}
        {showThirdLevel && selectedSubcategory && (
          <div className="w-72 bg-white shadow-xl border-r">
            <div className="p-4 border-b bg-gradient-to-r from-pink-500 to-red-500">
              <h4 className="font-semibold text-white">{selectedSubcategory.name}</h4>
            </div>
            <div className="p-2 max-h-screen overflow-y-auto">
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

function ProductCard({ product }: { product: Product }) {
  return (
    <Link href={`/products/${product.product_id}`}>
      <Card className="h-full transition-all hover:shadow-lg">
        <CardContent className="p-4">
          <div className="relative aspect-square mb-4">
            <Image
              src={product.image || '/pc1.jpg'}
              alt={product.name}
              fill
              className="object-contain rounded-lg"
            />
          </div>
          <h3 className="font-semibold mb-2 line-clamp-2">{product.name}</h3>
          <p className="text-sm text-muted-foreground mb-4 line-clamp-2">
            {product.description}
          </p>
          <div className="flex items-center justify-between">
            <p className="font-bold text-primary">${product.price}</p>
            <Button size="sm" disabled={product.stock_quantity === 0}>
              <ShoppingCart className="h-4 w-4 mr-2" />
              Add to Cart
            </Button>
          </div>
        </CardContent>
      </Card>
    </Link>
  )
}

export default function HomePage() {
  const [showSidebar, setShowSidebar] = useState(false)
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedBrand, setSelectedBrand] = useState<number | null>(null)
  const [priceRange, setPriceRange] = useState<string>("all")
  const [sortBy, setSortBy] = useState<string>("featured")
  const [categorySearchQuery, setCategorySearchQuery] = useState<string>("")

  const { data: productsData, isLoading: isLoadingProducts } = useGetProductsQuery()
  const { data: categoriesData, isLoading: isLoadingCategories } = useGetCategoriesQuery()
  const { data: brandsData, isLoading: isLoadingBrands } = useGetBrandsQuery()
  
  const products = productsData?.data.products || []
  const categories = categoriesData?.data.categories || []
  const brands = brandsData?.data.brands || []

  const filteredProducts = products
    .filter((product: Product) => {
      const matchesSearch = product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                           product.description.toLowerCase().includes(searchQuery.toLowerCase())
      const matchesBrand = selectedBrand ? product.brand_id === selectedBrand : true
      const matchesCategorySearch = categorySearchQuery
        ? product.name.toLowerCase().includes(categorySearchQuery.toLowerCase()) ||
          product.description.toLowerCase().includes(categorySearchQuery.toLowerCase())
        : true
      
      // Price range filter
      if (priceRange !== "all") {
        const [min, max] = priceRange.split("-").map(Number)
        if (max) {
          if (product.price < min || product.price > max) return false
        } else {
          if (product.price < min) return false
        }
      }
      
      return matchesSearch && matchesBrand && matchesCategorySearch
    })
    .sort((a: Product, b: Product) => {
      switch (sortBy) {
        case "price-low":
          return a.price - b.price
        case "price-high":
          return b.price - a.price
        case "name-asc":
          return a.name.localeCompare(b.name)
        case "name-desc":
          return b.name.localeCompare(a.name)
        default:
          return 0
      }
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
                <span className="absolute -top-1 -right-1 bg-gradient-to-r from-red-500 to-pink-500 text-white rounded-full w-5 h-5 flex items-center justify-center text-xs">
                  0
                </span>
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
              <div className="absolute inset-y-0 left-0">
                <CategorySidebar />
              </div>
            </div>
          )}

          {/* Main Content */}
          <div>
            {/* Filters */}
            <div className="mb-6 flex flex-wrap gap-4">
              {/* Brand Filter */}
              <Select
                value={selectedBrand?.toString() || "all"}
                onValueChange={(value) => setSelectedBrand(value === "all" ? null : parseInt(value))}
              >
                <SelectTrigger className="w-[200px]">
                  <SelectValue placeholder="Filter by brand" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All Brands</SelectItem>
                  {brands.map((brand: Brand) => (
                    <SelectItem key={brand.brand_id} value={brand.brand_id.toString()}>
                      {brand.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              {/* Price Range Filter */}
              <Select
                value={priceRange}
                onValueChange={setPriceRange}
              >
                <SelectTrigger className="w-[200px]">
                  <SelectValue placeholder="Price range" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All Prices</SelectItem>
                  <SelectItem value="0-50">Under $50</SelectItem>
                  <SelectItem value="50-100">$50 - $100</SelectItem>
                  <SelectItem value="100-200">$100 - $200</SelectItem>
                  <SelectItem value="200-500">$200 - $500</SelectItem>
                  <SelectItem value="500">$500 & Above</SelectItem>
                </SelectContent>
              </Select>

              {/* Sort Filter */}
              <Select
                value={sortBy}
                onValueChange={setSortBy}
              >
                <SelectTrigger className="w-[200px]">
                  <SelectValue placeholder="Sort by" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="featured">Featured</SelectItem>
                  <SelectItem value="price-low">Price: Low to High</SelectItem>
                  <SelectItem value="price-high">Price: High to Low</SelectItem>
                  <SelectItem value="name-asc">Name: A to Z</SelectItem>
                  <SelectItem value="name-desc">Name: Z to A</SelectItem>
                </SelectContent>
              </Select>
            </div>

            {isLoadingProducts ? (
              <div className="text-center py-8">Loading products...</div>
            ) : (
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {filteredProducts.map((product: Product) => (
                  <ProductCard key={product.product_id} product={product} />
                ))}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
