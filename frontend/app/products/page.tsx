"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Star, ShoppingCart, Search, ArrowLeft, User, Heart, Grid, List, PanelRightClose, PanelLeftClose, X, Menu } from "lucide-react"
import Image from "next/image"
import Link from "next/link"
import { useSearchParams } from "next/navigation"
import { Sidebar, SidebarContent } from "@/components/ui/sidebar"

interface Product {
  product_id: number
  category_id: number
  brand_id: number
  name: string
  description: string
  price: number
  stock_quantity: number
  created_at: string
  image: string
  brand: {
    name: string
    description: string
  }
  category: {
    category_name: string
  }
  average_rating: number
  total_reviews: number
}

interface CartItem extends Product {
  quantity: number
}

function ProductCard({ product }: { product: Product }) {
  return (
    <Link href={`/products/${product.product_id}`}>
      <Card className="h-full transition-all hover:shadow-lg">
        <CardContent className="p-4">
          <div className="relative aspect-square mb-4">
            <Image
              src={product.image}
              alt={product.name}
              fill
              className="object-contain rounded-lg"
            />
          </div>
          <div className="flex items-center space-x-2 mb-2">
            <span className="text-sm text-muted-foreground">{product.brand.name}</span>
            <span className="text-muted-foreground">•</span>
            <span className="text-sm text-muted-foreground">{product.category.category_name}</span>
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

export default function ProductsPage() {
  const searchParams = useSearchParams()
  const categoryParam = searchParams.get("category")?.toLowerCase()
  const [searchQuery, setSearchQuery] = useState("")
  const [sortBy, setSortBy] = useState("featured")
  const [filteredProducts, setFilteredProducts] = useState<Product[]>([])
  const [cartItems, setCartItems] = useState<CartItem[]>([])
  const [viewMode, setViewMode] = useState<"grid" | "list">("grid")
  const [priceRange, setPriceRange] = useState<string>("all")
  const [showSidebar, setShowSidebar] = useState(false)
  const [isSidebarOpen, setIsSidebarOpen] = useState(false)

  useEffect(() => {
    let filtered: Product[] = []

    // Filter by category from URL parameter
    if (categoryParam) {
      filtered = filtered.filter(
        (product) =>
          product.category.category_name.toLowerCase().includes(categoryParam) ||
          product.brand.name.toLowerCase().includes(categoryParam),
      )
    }

    // Filter by search query
    if (searchQuery) {
      filtered = filtered.filter(
        (product) =>
          product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          product.category.category_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          product.brand.name.toLowerCase().includes(searchQuery.toLowerCase()),
      )
    }

    // Filter by price range
    if (priceRange !== "all") {
      const [min, max] = priceRange.split("-").map(Number)
      filtered = filtered.filter((product) => {
        if (max) {
          return product.price >= min && product.price <= max
        } else {
          return product.price >= min
        }
      })
    }

    // Sort products
    switch (sortBy) {
      case "price-low":
        filtered.sort((a, b) => a.price - b.price)
        break
      case "price-high":
        filtered.sort((a, b) => b.price - a.price)
        break
      case "rating":
        filtered.sort((a, b) => b.average_rating - a.average_rating)
        break
      case "reviews":
        filtered.sort((a, b) => b.total_reviews - a.total_reviews)
        break
    }

    setFilteredProducts(filtered)
  }, [searchQuery, sortBy, categoryParam, priceRange])

  const handleAddToCart = (product: Product) => {
    setCartItems((prev) => {
      const existing = prev.find((item) => item.product_id === product.product_id)
      if (existing) {
        return prev.map((item) => (item.product_id === product.product_id ? { ...item, quantity: item.quantity + 1 } : item))
      }
      return [...prev, { ...product, quantity: 1 }]
    })
  }

  function CategorySidebar() {
    const categories = [
      { id: 1, name: "Laptops" },
      { id: 2, name: "Accessories" },
      { id: 3, name: "Monitors" },
      { id: 4, name: "Furniture" },
    ]

    return (
      <div className="space-y-4">
        <h3 className="font-semibold">Categories</h3>
        <div className="space-y-2">
          {categories.map((category) => (
            <Link
              key={category.id}
              href={`/products?category=${category.name}`}
              className={`block px-4 py-2 rounded-lg ${
                categoryParam === category.name.toLowerCase()
                  ? "bg-primary text-primary-foreground"
                  : "hover:bg-muted"
              }`}
            >
              {category.name}
            </Link>
          ))}
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-blue-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b sticky top-0 z-50">
        <div className="container flex h-16 items-center justify-between px-4">
          <div className="flex items-center gap-6">
            <Button
              variant="ghost"
              size="icon"
              className="md:hidden"
              onClick={() => setIsSidebarOpen(true)}
            >
              <Menu className="h-6 w-6" />
            </Button>
            <Button
              variant="ghost"
              size="icon"
              className="hidden lg:flex"
              onClick={() => setShowSidebar(!showSidebar)}
            >
              {showSidebar ? <PanelRightClose className="h-6 w-6" /> : <PanelLeftClose className="h-6 w-6" />}
            </Button>
            <Link href="/" className="flex items-center space-x-2">
              <span className="text-xl font-bold">SanqaSuq</span>
            </Link>
          </div>

          <div className="flex items-center gap-4">
            <div className="relative">
              <Link href="/cart">
                <Button variant="ghost" size="icon">
                  <ShoppingCart className="h-6 w-6" />
                  {cartItems.length > 0 && (
                    <span className="absolute -top-1 -right-1 flex h-5 w-5 items-center justify-center rounded-full bg-primary text-xs text-primary-foreground">
                      {cartItems.length}
                    </span>
                  )}
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </header>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Filters and Controls */}
        <div className="flex flex-col lg:flex-row items-start lg:items-center justify-between mb-8 space-y-4 lg:space-y-0">
          <div className="flex items-center space-x-4">
            <Badge variant="secondary" className="bg-gradient-to-r from-indigo-100 to-purple-100 text-indigo-700">
              {filteredProducts.length} Products Found
            </Badge>
            {categoryParam && (
              <Badge className="bg-gradient-to-r from-blue-500 to-indigo-500 text-white">
                Category: {categoryParam}
              </Badge>
            )}
          </div>

          <div className="flex items-center space-x-4">
            <Select value={priceRange} onValueChange={setPriceRange}>
              <SelectTrigger className="w-48 border-indigo-200">
                <SelectValue placeholder="Price Range" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">All Prices</SelectItem>
                <SelectItem value="0-100">Under $100</SelectItem>
                <SelectItem value="100-500">$100 - $500</SelectItem>
                <SelectItem value="500-1000">$500 - $1000</SelectItem>
                <SelectItem value="1000-2000">$1000 - $2000</SelectItem>
                <SelectItem value="2000">Over $2000</SelectItem>
              </SelectContent>
            </Select>

            <Select value={sortBy} onValueChange={setSortBy}>
              <SelectTrigger className="w-48 border-indigo-200">
                <SelectValue placeholder="Sort by" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="featured">Featured</SelectItem>
                <SelectItem value="price-low">Price: Low to High</SelectItem>
                <SelectItem value="price-high">Price: High to Low</SelectItem>
                <SelectItem value="rating">Highest Rated</SelectItem>
                <SelectItem value="reviews">Most Reviews</SelectItem>
              </SelectContent>
            </Select>

            <div className="flex items-center border border-indigo-200 rounded-lg">
              <Button
                variant={viewMode === "grid" ? "default" : "ghost"}
                size="sm"
                onClick={() => setViewMode("grid")}
                className={viewMode === "grid" ? "bg-gradient-to-r from-indigo-600 to-purple-600" : ""}
              >
                <Grid className="h-4 w-4" />
              </Button>
              <Button
                variant={viewMode === "list" ? "default" : "ghost"}
                size="sm"
                onClick={() => setViewMode("list")}
                className={viewMode === "list" ? "bg-gradient-to-r from-indigo-600 to-purple-600" : ""}
              >
                <List className="h-4 w-4" />
              </Button>
            </div>
          </div>
        </div>

        {/* Products Grid */}
        {filteredProducts.length === 0 ? (
          <div className="text-center py-16">
            <div className="w-24 h-24 bg-gradient-to-r from-indigo-100 to-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <Search className="h-12 w-12 text-indigo-400" />
            </div>
            <h2 className="text-2xl font-semibold text-gray-900 mb-2">No products found</h2>
            <p className="text-gray-600 mb-8">Try adjusting your search or filters</p>
            <Link href="/home">
              <Button className="bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700">
                Browse All Products
              </Button>
            </Link>
          </div>
        ) : (
          <div
            className={`grid gap-6 ${
              viewMode === "grid" ? "grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4" : "grid-cols-1"
            }`}
          >
            {filteredProducts.map((product) => (
              <ProductCard key={product.product_id} product={product} />
            ))}
          </div>
        )}
      </div>

      <div className="flex">
        {/* Mobile Sidebar - Overlay */}
        {isSidebarOpen && (
          <div className="fixed inset-0 bg-black/50 z-40 lg:hidden" onClick={() => setIsSidebarOpen(false)}>
            <div className="fixed inset-y-0 left-0 w-72 bg-white shadow-lg z-50" onClick={e => e.stopPropagation()}>
              <CategorySidebar />
            </div>
          </div>
        )}

        {/* Desktop Sidebar */}
        {showSidebar && (
          <Sidebar className="hidden lg:block fixed left-0 top-16 h-[calc(100vh-4rem)] z-40">
            <SidebarContent className="p-0">
              <CategorySidebar />
            </SidebarContent>
          </Sidebar>
        )}

        {/* Main Content */}
        <main className={`flex-1 p-6 ${showSidebar ? 'lg:ml-72' : 'lg:mx-auto max-w-7xl'}`}>
          {/* Rest of the component content */}
        </main>
      </div>
    </div>
  )
}
