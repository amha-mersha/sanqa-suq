"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Star, ShoppingCart, ChevronLeft } from "lucide-react"
import Image from "next/image"
import Link from "next/link"
import { useParams } from "next/navigation"

interface Review {
  review_id: string
  user_id: string
  rating: number
  comment: string
  review_date: string
  user_name: string
}

interface Specification {
  spec_name: string
  spec_value: string
}

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
  specifications: Specification[]
  reviews: Review[]
  average_rating: number
  total_reviews: number
}

export default function ProductDetail() {
  const params = useParams()
  const [product, setProduct] = useState<Product | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    // TODO: Replace with actual API call
    // This is mock data for demonstration
    setProduct({
      product_id: parseInt(params.id as string),
      category_id: 1,
      brand_id: 1,
      name: "Gaming Laptop Pro X",
      description: "High-performance gaming laptop with the latest RTX 4080 graphics card, 32GB RAM, and 2TB SSD storage.",
      price: 2499.99,
      stock_quantity: 10,
      created_at: "2024-03-15T10:30:00Z",
      image: "/path/to/image.jpg",
      brand: {
        name: "GamingTech",
        description: "Premium gaming hardware manufacturer"
      },
      category: {
        category_name: "Laptops"
      },
      specifications: [
        { spec_name: "Processor", spec_value: "Intel Core i9-13900K" },
        { spec_name: "Graphics", spec_value: "NVIDIA RTX 4080 16GB" },
        { spec_name: "Memory", spec_value: "32GB DDR5" },
        { spec_name: "Storage", spec_value: "2TB NVMe SSD" },
        { spec_name: "Display", spec_value: "17.3\" 4K 144Hz" },
        { spec_name: "Operating System", spec_value: "Windows 11 Pro" },
        { spec_name: "Battery", spec_value: "90Wh" },
        { spec_name: "Weight", spec_value: "2.5 kg" },
      ],
      reviews: [
        {
          review_id: "1",
          user_id: "1",
          rating: 5,
          comment: "Amazing laptop! The performance is incredible and the display is stunning.",
          review_date: "2024-03-15T10:30:00Z",
          user_name: "John Doe"
        },
        {
          review_id: "2",
          user_id: "2",
          rating: 4,
          comment: "Great laptop but a bit heavy. The battery life could be better.",
          review_date: "2024-03-14T15:45:00Z",
          user_name: "Jane Smith"
        }
      ],
      average_rating: 4.5,
      total_reviews: 2
    })
    setLoading(false)
  }, [params.id])

  if (loading) {
    return <div className="container mx-auto p-6">Loading...</div>
  }

  if (!product) {
    return <div className="container mx-auto p-6">Product not found</div>
  }

  return (
    <div className="container mx-auto p-6">
      <Link href="/products" className="inline-flex items-center text-sm text-muted-foreground hover:text-foreground mb-6">
        <ChevronLeft className="h-4 w-4 mr-1" />
        Back to Products
      </Link>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
        {/* Product Image */}
        <div className="relative aspect-square">
          <Image
            src={product.image}
            alt={product.name}
            fill
            className="object-contain rounded-lg"
          />
        </div>

        {/* Product Info */}
        <div className="space-y-6">
          <div>
            <div className="flex items-center space-x-2 mb-2">
              <span className="text-sm text-muted-foreground">{product.brand.name}</span>
              <span className="text-muted-foreground">•</span>
              <span className="text-sm text-muted-foreground">{product.category.category_name}</span>
            </div>
            <h1 className="text-3xl font-bold mb-2">{product.name}</h1>
            <div className="flex items-center space-x-2 mb-4">
              <div className="flex items-center">
                {[...Array(5)].map((_, i) => (
                  <Star
                    key={i}
                    className={`h-5 w-5 ${
                      i < Math.floor(product.average_rating)
                        ? "text-yellow-400 fill-current"
                        : "text-gray-300"
                    }`}
                  />
                ))}
              </div>
              <span className="text-sm text-muted-foreground">
                ({product.total_reviews} reviews)
              </span>
            </div>
            <p className="text-2xl font-bold text-primary">${product.price}</p>
            <p className="text-sm text-muted-foreground mt-2">
              {product.stock_quantity > 0 ? (
                <span className="text-green-600">In Stock ({product.stock_quantity} available)</span>
              ) : (
                <span className="text-red-600">Out of Stock</span>
              )}
            </p>
          </div>

          <p className="text-muted-foreground">{product.description}</p>

          <Button size="lg" className="w-full" disabled={product.stock_quantity === 0}>
            <ShoppingCart className="h-5 w-5 mr-2" />
            Add to Cart
          </Button>
        </div>
      </div>

      {/* Tabs for Specifications and Reviews */}
      <Tabs defaultValue="specifications" className="space-y-6">
        <TabsList>
          <TabsTrigger value="specifications">Specifications</TabsTrigger>
          <TabsTrigger value="reviews">Reviews</TabsTrigger>
        </TabsList>

        <TabsContent value="specifications">
          <Card>
            <CardContent className="p-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                {product.specifications.map((spec, index) => (
                  <div
                    key={index}
                    className={`p-4 ${
                      index % 2 === 0 ? "bg-muted/50" : "bg-background"
                    } rounded-lg`}
                  >
                    <div className="font-semibold mb-1">{spec.spec_name}</div>
                    <div className="text-muted-foreground">{spec.spec_value}</div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="reviews">
          <div className="space-y-6">
            {product.reviews.map((review) => (
              <Card key={review.review_id}>
                <CardContent className="p-6">
                  <div className="flex items-center justify-between mb-4">
                    <div className="flex items-center space-x-2">
                      <div className="flex items-center">
                        {[...Array(5)].map((_, i) => (
                          <Star
                            key={i}
                            className={`h-4 w-4 ${
                              i < review.rating
                                ? "text-yellow-400 fill-current"
                                : "text-gray-300"
                            }`}
                          />
                        ))}
                      </div>
                      <span className="font-semibold">{review.user_name}</span>
                    </div>
                    <span className="text-sm text-muted-foreground">
                      {new Date(review.review_date).toLocaleDateString()}
                    </span>
                  </div>
                  <p className="text-muted-foreground">{review.comment}</p>
                </CardContent>
              </Card>
            ))}
          </div>
        </TabsContent>
      </Tabs>
    </div>
  )
} 