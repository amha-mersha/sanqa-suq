"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Plus, Edit, Trash2, Package, DollarSign, Tag } from "lucide-react"
import Image from "next/image"

interface Product {
  id: string
  name: string
  description: string
  price: number
  category: string
  image: string
  stock: number
}

const CATEGORIES = [
  "Laptops",
  "Smartphones",
  "PC Components",
  "Accessories",
  "Gaming",
  "Audio",
  "Networking",
  "Storage",
  "Monitors",
  "Other"
]

export default function SellerPage() {
  const [products, setProducts] = useState<Product[]>([
    {
      id: "1",
      name: "Gaming Laptop",
      description: "High-performance gaming laptop with RTX 3080",
      price: 1499.99,
      category: "Laptops",
      image: "/images/laptop.jpg",
      stock: 10,
    },
    {
      id: "2",
      name: "iPhone 15 Pro Max",
      description: "Latest Apple flagship smartphone with advanced camera system",
      price: 1199.99,
      category: "Smartphones",
      image: "/images/phone.jpg",
      stock: 15,
    },
    {
      id: "3",
      name: "RTX 4080 Graphics Card",
      description: "Top-tier graphics card for gaming and creative work",
      price: 899.99,
      category: "PC Components",
      image: "/images/gpu.jpg",
      stock: 5,
    },
    {
      id: "4",
      name: "Sony WH-1000XM5 Headphones",
      description: "Industry-leading noise cancelling headphones",
      price: 349.99,
      category: "Accessories",
      image: "/images/headphones.jpg",
      stock: 20,
    },
  ])

  const [isAddingProduct, setIsAddingProduct] = useState(false)
  const [newProduct, setNewProduct] = useState<Partial<Product>>({
    name: "",
    description: "",
    price: 0,
    category: "",
    image: "/placeholder.svg",
    stock: 0,
  })

  const handleAddProduct = () => {
    if (newProduct.name && newProduct.price) {
      const product: Product = {
        id: Date.now().toString(),
        name: newProduct.name,
        description: newProduct.description || "",
        price: newProduct.price,
        category: newProduct.category || "Uncategorized",
        image: newProduct.image || "/placeholder.svg",
        stock: newProduct.stock || 0,
      }
      setProducts([...products, product])
      setNewProduct({
        name: "",
        description: "",
        price: 0,
        category: "",
        image: "/placeholder.svg",
        stock: 0,
      })
      setIsAddingProduct(false)
    }
  }

  const handleDeleteProduct = (id: string) => {
    setProducts(products.filter(product => product.id !== id))
  }

  return (
    <div className="container mx-auto p-6">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Seller Dashboard</h1>
        <Button onClick={() => setIsAddingProduct(true)}>
          <Plus className="mr-2 h-4 w-4" />
          Add Product
        </Button>
      </div>

      <Tabs defaultValue="products" className="space-y-4">
        <TabsList>
          <TabsTrigger value="products">Products</TabsTrigger>
          <TabsTrigger value="orders">Orders</TabsTrigger>
          <TabsTrigger value="analytics">Analytics</TabsTrigger>
        </TabsList>

        <TabsContent value="products" className="space-y-4">
          {isAddingProduct ? (
            <Card>
              <CardHeader>
                <CardTitle>Add New Product</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">Product Name</label>
                    <Input
                      value={newProduct.name}
                      onChange={(e) => setNewProduct({ ...newProduct, name: e.target.value })}
                      placeholder="Enter product name"
                    />
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">Price</label>
                    <Input
                      type="number"
                      value={newProduct.price}
                      onChange={(e) => setNewProduct({ ...newProduct, price: parseFloat(e.target.value) })}
                      placeholder="Enter price"
                    />
                  </div>
                </div>
                <div className="space-y-2">
                  <label className="text-sm font-medium">Description</label>
                  <Textarea
                    value={newProduct.description}
                    onChange={(e) => setNewProduct({ ...newProduct, description: e.target.value })}
                    placeholder="Enter product description"
                  />
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">Category</label>
                    <Select
                      value={newProduct.category}
                      onValueChange={(value) => setNewProduct({ ...newProduct, category: value })}
                    >
                      <SelectTrigger>
                        <SelectValue placeholder="Select a category" />
                      </SelectTrigger>
                      <SelectContent>
                        {CATEGORIES.map((category) => (
                          <SelectItem key={category} value={category}>
                            {category}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">Stock</label>
                    <Input
                      type="number"
                      value={newProduct.stock}
                      onChange={(e) => setNewProduct({ ...newProduct, stock: parseInt(e.target.value) })}
                      placeholder="Enter stock quantity"
                    />
                  </div>
                </div>
                <div className="flex justify-end space-x-2">
                  <Button variant="outline" onClick={() => setIsAddingProduct(false)}>
                    Cancel
                  </Button>
                  <Button onClick={handleAddProduct}>Add Product</Button>
                </div>
              </CardContent>
            </Card>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {products.map((product) => (
                <Card key={product.id} className="hover:shadow-lg transition-shadow">
                  <CardContent className="p-4">
                    <div className="relative h-48 mb-4 bg-gray-100 rounded-lg">
                      <Image
                        src={product.image}
                        alt={product.name}
                        fill
                        className="object-contain rounded-lg"
                        onError={(e: React.SyntheticEvent<HTMLImageElement, Event>) => {
                          const target = e.target as HTMLImageElement;
                          target.src = "/placeholder.svg";
                        }}
                      />
                    </div>
                    <div className="space-y-2">
                      <h3 className="font-semibold text-lg">{product.name}</h3>
                      <p className="text-sm text-muted-foreground line-clamp-2">
                        {product.description}
                      </p>
                      <div className="flex items-center justify-between">
                        <span className="font-bold">${product.price}</span>
                        <span className="text-sm text-muted-foreground">
                          Stock: {product.stock}
                        </span>
                      </div>
                      <div className="flex justify-end space-x-2">
                        <Button variant="outline" size="sm">
                          <Edit className="h-4 w-4" />
                        </Button>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => handleDeleteProduct(product.id)}
                        >
                          <Trash2 className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          )}
        </TabsContent>

        <TabsContent value="orders">
          <Card>
            <CardHeader>
              <CardTitle>Recent Orders</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-muted-foreground">No orders yet</p>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="analytics">
          <Card>
            <CardHeader>
              <CardTitle>Sales Analytics</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-muted-foreground">No analytics data available</p>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  )
} 