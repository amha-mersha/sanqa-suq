"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Sidebar, SidebarContent, SidebarProvider } from "@/components/ui/sidebar"
import { Star, ShoppingCart, Search, ChevronRight, ArrowLeft, User, Heart, Menu, X, Home, Cpu, PanelRightClose, PanelLeftClose } from "lucide-react"
import Image from "next/image"
import Link from "next/link"
import { ThemeToggle } from "@/components/theme-toggle"
import { Sheet, SheetContent, SheetClose } from "@/components/ui/sheet"

// Mock data for products
const topRatedProducts = [
  {
    product_id: 1,
    category_id: 1,
    brand_id: 1,
    name: "Gaming Laptop Pro X",
    description: "High-performance gaming laptop with the latest RTX 4080 graphics card, 32GB RAM, and 2TB SSD storage.",
    price: 2499.99,
    stock_quantity: 10,
    created_at: "2024-03-15T10:30:00Z",
    image: "/pc2.jpg",
    brand: {
      name: "GamingTech",
      description: "Premium gaming hardware manufacturer"
    },
    category: {
      category_name: "Laptops"
    },
    average_rating: 4.8,
    total_reviews: 156
  },
  {
    product_id: 2,
    category_id: 2,
    brand_id: 2,
    name: "Wireless Gaming Mouse",
    description: "Ergonomic wireless gaming mouse with 20,000 DPI sensor and customizable RGB lighting.",
    price: 79.99,
    stock_quantity: 25,
    created_at: "2024-03-14T15:45:00Z",
    image: "/pc1.jpg",
    brand: {
      name: "TechGear",
      description: "Gaming peripherals manufacturer"
    },
    category: {
      category_name: "Accessories"
    },
    average_rating: 4.6,
    total_reviews: 89
  },
  {
    product_id: 3,
    category_id: 2,
    brand_id: 3,
    name: "Mechanical Keyboard",
    description: "RGB mechanical keyboard with Cherry MX switches and aluminum frame.",
    price: 149.99,
    stock_quantity: 15,
    created_at: "2024-03-13T09:15:00Z",
    image: "/pc1.jpg",
    brand: {
      name: "KeyMaster",
      description: "Premium keyboard manufacturer"
    },
    category: {
      category_name: "Accessories"
    },
    average_rating: 4.7,
    total_reviews: 234
  },
  {
    product_id: 4,
    category_id: 2,
    brand_id: 4,
    name: "Gaming Headset",
    description: "7.1 surround sound gaming headset with noise-cancelling microphone and memory foam ear cushions.",
    price: 129.99,
    stock_quantity: 20,
    created_at: "2024-03-12T14:30:00Z",
    image: "/pc2.jpg",
    brand: {
      name: "SoundPro",
      description: "Audio equipment manufacturer"
    },
    category: {
      category_name: "Accessories"
    },
    average_rating: 4.5,
    total_reviews: 178
  }
]

const allProducts = [
  ...topRatedProducts,
  {
    product_id: 5,
    category_id: 3,
    brand_id: 5,
    name: "Gaming Monitor",
    description: "27-inch 4K gaming monitor with 144Hz refresh rate and HDR support.",
    price: 499.99,
    stock_quantity: 8,
    created_at: "2024-03-11T11:20:00Z",
    image: "/pc1.jpg",
    brand: {
      name: "DisplayTech",
      description: "Display technology manufacturer"
    },
    category: {
      category_name: "Monitors"
    },
    average_rating: 4.4,
    total_reviews: 92
  },
  {
    product_id: 6,
    category_id: 4,
    brand_id: 6,
    name: "Gaming Chair",
    description: "Ergonomic gaming chair with adjustable lumbar support and 4D armrests.",
    price: 299.99,
    stock_quantity: 12,
    created_at: "2024-03-10T16:45:00Z",
    image: "/pc2.jpg",
    brand: {
      name: "ComfortPro",
      description: "Ergonomic furniture manufacturer"
    },
    category: {
      category_name: "Furniture"
    },
    average_rating: 4.3,
    total_reviews: 67
  }
]

// Category structure with nested levels
const categories = [
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
      // Navigate to products page with this subcategory
      window.location.href = `/products?category=${encodeURIComponent(subcategory.name)}`
    }
  }

  const handleThirdLevelClick = (item: any) => {
    // Navigate to products page with this specific category
    window.location.href = `/products?category=${encodeURIComponent(item.name)}`
  }

  const handleBackClick = () => {
    setShowThirdLevel(false)
    setSelectedSubcategory(null)
  }

  return (
    <div className="fixed inset-0 z-50 bg-black bg-opacity-50 lg:relative lg:bg-transparent lg:inset-auto">
      <div className="flex h-full lg:h-auto">
        {/* First Level - Always visible */}
        <div
          className={`w-72 bg-white shadow-xl lg:shadow-none border-r transition-all duration-300 ${showThirdLevel ? "opacity-50" : ""}`}
        >
          <div className="p-4 border-b bg-gradient-to-r from-indigo-600 to-purple-600">
            <h3 className="font-semibold text-white">Categories</h3>
          </div>
          <div className="p-2 max-h-screen overflow-y-auto">
            {categories.map((category) => (
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

export default function HomePage() {
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedCategory, setSelectedCategory] = useState<string | undefined>(undefined)
  const [sortBy, setSortBy] = useState("featured")
  const [filteredProducts, setFilteredProducts] = useState<Product[]>(allProducts)
  const [cartItems, setCartItems] = useState<CartItem[]>([])
  const [isSidebarOpen, setIsSidebarOpen] = useState(false)
  const [isCartOpen, setIsCartOpen] = useState(false)
  const [showSidebar, setShowSidebar] = useState(true)

  useEffect(() => {
    let filtered = allProducts

    if (selectedCategory) {
      filtered = filtered.filter(
        (product) => product.category.category_name === selectedCategory
      )
    }

    if (searchQuery) {
      filtered = filtered.filter(
        (product) =>
          product.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          product.category.category_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          product.brand.name.toLowerCase().includes(searchQuery.toLowerCase())
      )
    }

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
  }, [searchQuery, selectedCategory, sortBy])

  const handleAddToCart = (product: Product) => {
    setCartItems((prev) => {
      const existing = prev.find((item) => item.product_id === product.product_id)
      if (existing) {
        return prev.map((item) => (item.product_id === product.product_id ? { ...item, quantity: item.quantity + 1 } : item))
      }
      return [...prev, { ...product, quantity: 1 }]
    })
  }

  return (
    <SidebarProvider>
      <div className="min-h-screen bg-background">
        <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
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
            {/* Top Rated Products Section */}
            <section className="mb-12">
              <div className="flex items-center justify-between mb-6">
                <h2 className="text-2xl font-bold text-gray-900">Top Rated Products</h2>
                <Link href="/products/top-rated">
                  <Button variant="outline">View All</Button>
                </Link>
              </div>

              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                {topRatedProducts.map((product) => (
                  <ProductCard key={product.product_id} product={product} />
                ))}
              </div>
            </section>

            {/* Filters and All Products */}
            <section>
              <div className="flex items-center justify-between mb-6">
                <h2 className="text-2xl font-bold text-gray-900">All Products</h2>
                <div className="flex items-center space-x-4">
                  <Select value={selectedCategory} onValueChange={setSelectedCategory}>
                    <SelectTrigger className="w-48">
                      <SelectValue placeholder="Category" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value="All">All Categories</SelectItem>
                      <SelectItem value="Laptops">Laptops</SelectItem>
                      <SelectItem value="Smartphones">Smartphones</SelectItem>
                      <SelectItem value="PC Components">PC Components</SelectItem>
                      <SelectItem value="Audio">Audio</SelectItem>
                    </SelectContent>
                  </Select>

                  <Select value={sortBy} onValueChange={setSortBy}>
                    <SelectTrigger className="w-48">
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
                </div>
              </div>

              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {filteredProducts.map((product) => (
                  <ProductCard key={product.product_id} product={product} />
                ))}
              </div>

              {filteredProducts.length === 0 && (
                <div className="text-center py-12">
                  <p className="text-gray-500 text-lg">No products found matching your criteria.</p>
                </div>
              )}
            </section>
          </main>
        </div>

        {/* Sidebar */}
        <Sheet open={isSidebarOpen} onOpenChange={setIsSidebarOpen}>
          <SheetContent side="left" className="w-[300px] sm:w-[400px] p-0">
            <div className="flex flex-col h-full">
              <div className="flex items-center justify-between p-4 border-b">
                <h2 className="text-lg font-semibold">Menu</h2>
                <SheetClose asChild>
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-8 w-8"
                  >
                    <X className="h-4 w-4" />
                    <span className="sr-only">Close menu</span>
                  </Button>
                </SheetClose>
              </div>
              <div className="flex-1 overflow-y-auto p-4">
                <nav className="space-y-4">
                  <SheetClose asChild>
                    <Link
                      href="/"
                      className="flex items-center space-x-2 text-foreground hover:text-primary transition-colors"
                    >
                      <Home className="h-4 w-4" />
                      <span>Home</span>
                    </Link>
                  </SheetClose>
                  <SheetClose asChild>
                    <Link
                      href="/pc-builder"
                      className="flex items-center space-x-2 text-foreground hover:text-primary transition-colors"
                    >
                      <Cpu className="h-4 w-4" />
                      <span>PC Builder</span>
                    </Link>
                  </SheetClose>
                  <SheetClose asChild>
                    <Link
                      href="/cart"
                      className="flex items-center space-x-2 text-foreground hover:text-primary transition-colors"
                    >
                      <ShoppingCart className="h-4 w-4" />
                      <span>Cart</span>
                    </Link>
                  </SheetClose>
                  <SheetClose asChild>
                    <Link
                      href="/account"
                      className="flex items-center space-x-2 text-foreground hover:text-primary transition-colors"
                    >
                      <User className="h-4 w-4" />
                      <span>Account</span>
                    </Link>
                  </SheetClose>
                </nav>
              </div>
            </div>
          </SheetContent>
        </Sheet>
      </div>
    </SidebarProvider>
  )
}
