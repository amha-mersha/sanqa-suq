"use client"

import { useState, useMemo } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Plus, Edit, Trash2, Package, DollarSign, Tag, CreditCard } from "lucide-react"
import Image from "next/image"
import { useGetProductsQuery, useGetCategoriesQuery, useGetBrandsQuery, useAddProductMutation } from "@/lib/api"
import { Product, Category, Brand } from "@/lib/types"

// Dummy order data
const dummyOrders = [
  {
    id: 1,
    user_id: "123e4567-e89b-12d3-a456-426614174000",
    address_id: 1,
    total_amount: 299.99,
    payment_method: "credit_card",
    status: "pending",
    created_at: "2024-03-15T10:30:00Z",
    items: [
      { product_id: 1, quantity: 2, price: 149.99 }
    ]
  },
  {
    id: 2,
    user_id: "123e4567-e89b-12d3-a456-426614174000",
    address_id: 1,
    total_amount: 599.99,
    payment_method: "debit_card",
    status: "completed",
    created_at: "2024-03-14T15:45:00Z",
    items: [
      { product_id: 2, quantity: 1, price: 599.99 }
    ]
  }
]

// Remove the getRandomProductImage function and replace with a constant
const DEFAULT_PRODUCT_IMAGE = "/pc1.jpg"

export default function SellerPage() {
  const [isAddingProduct, setIsAddingProduct] = useState(false)
  const [newProduct, setNewProduct] = useState({
    category_id: "",
    brand_id: "",
    name: "",
    description: "",
    price: 0,
    stock_quantity: 0,
  })

  const { data: productsData, isLoading: isLoadingProducts } = useGetProductsQuery()
  const { data: categoriesData, isLoading: isLoadingCategories } = useGetCategoriesQuery()
  const { data: brandsData, isLoading: isLoadingBrands } = useGetBrandsQuery()
  const [addProduct, { isLoading: isAdding }] = useAddProductMutation()

  const products = productsData?.data.products || []
  const categories = categoriesData?.data || []
  const brands = brandsData?.data.brands || []

  // Get leaf-level categories (categories without children)
  const leafCategories = useMemo(() => {
    const categoryMap = new Map()
    const allCategories = categories

    // First pass: Create a map of all categories
    allCategories.forEach(category => {
      categoryMap.set(category.category_id, {
        ...category,
        hasChildren: false
      })
    })

    // Second pass: Mark categories that have children
    allCategories.forEach(category => {
      if (category.parent_category_id) {
        const parent = categoryMap.get(category.parent_category_id)
        if (parent) {
          parent.hasChildren = true
        }
      }
    })

    // Return only categories without children
    return Array.from(categoryMap.values()).filter(category => !category.hasChildren)
  }, [categories])

  const handleAddProduct = async () => {
    if (!newProduct.name || !newProduct.category_id || !newProduct.brand_id || !newProduct.price || !newProduct.stock_quantity) {
      alert("Please fill in all required fields")
      return
    }

    try {
      await addProduct({
        category_id: parseInt(newProduct.category_id),
        brand_id: parseInt(newProduct.brand_id),
        name: newProduct.name,
        description: newProduct.description,
        price: parseFloat(newProduct.price.toString()),
        stock_quantity: parseInt(newProduct.stock_quantity.toString()),
      }).unwrap()

      // Reset form
      setNewProduct({
        category_id: "",
        brand_id: "",
        name: "",
        description: "",
        price: 0,
        stock_quantity: 0,
      })
      setIsAddingProduct(false)
    } catch (error) {
      console.error("Error adding product:", error)
      alert("Failed to add product. Please try again.")
    }
  }

  const handleDeleteProduct = async (id: number) => {
    if (!confirm("Are you sure you want to delete this product?")) {
      return
    }

    try {
      const response = await fetch(`http://localhost:8080/api/v1/product/${id}`, {
        method: "DELETE",
      })

      if (!response.ok) {
        throw new Error("Failed to delete product")
      }
    } catch (error) {
      console.error("Error deleting product:", error)
      alert("Failed to delete product. Please try again.")
    }
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
              <CardContent>
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
                <div className="space-y-2 mt-4">
                  <label className="text-sm font-medium">Description</label>
                  <Textarea
                    value={newProduct.description}
                    onChange={(e) => setNewProduct({ ...newProduct, description: e.target.value })}
                    placeholder="Enter product description"
                  />
                </div>
                <div className="grid grid-cols-2 gap-4 mt-4">
                  <div className="space-y-2">
                    <label className="text-sm font-medium">Category</label>
                    <Select
                      value={newProduct.category_id}
                      onValueChange={(value) => setNewProduct({ ...newProduct, category_id: value })}
                    >
                      <SelectTrigger>
                        <SelectValue placeholder="Select a category" />
                      </SelectTrigger>
                      <SelectContent>
                        {leafCategories.map((category) => (
                          <SelectItem key={category.category_id} value={category.category_id.toString()}>
                            {category.name}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>
                  <div className="space-y-2">
                    <label className="text-sm font-medium">Brand</label>
                    <Select
                      value={newProduct.brand_id}
                      onValueChange={(value) => setNewProduct({ ...newProduct, brand_id: value })}
                    >
                      <SelectTrigger>
                        <SelectValue placeholder="Select a brand" />
                      </SelectTrigger>
                      <SelectContent>
                        {brands.map((brand: Brand) => (
                          <SelectItem key={brand.brand_id} value={brand.brand_id.toString()}>
                            {brand.name}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>
                </div>
                <div className="space-y-2 mt-4">
                  <label className="text-sm font-medium">Stock Quantity</label>
                  <Input
                    type="number"
                    value={newProduct.stock_quantity}
                    onChange={(e) => setNewProduct({ ...newProduct, stock_quantity: parseInt(e.target.value) })}
                    placeholder="Enter stock quantity"
                  />
                </div>
                <div className="flex justify-end space-x-2 mt-6">
                  <Button variant="outline" onClick={() => setIsAddingProduct(false)}>
                    Cancel
                  </Button>
                  <Button onClick={handleAddProduct} disabled={isAdding}>
                    {isAdding ? "Adding..." : "Add Product"}
                  </Button>
                </div>
              </CardContent>
            </Card>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {products.map((product: Product) => (
                <Card key={product.product_id} className="hover:shadow-lg transition-shadow">
                  <CardContent className="p-4">
                    <div className="relative h-48 mb-4 bg-gray-100 rounded-lg">
                      <Image
                        src={DEFAULT_PRODUCT_IMAGE}
                        alt={product.name}
                        fill
                        className="object-contain rounded-lg"
                        onError={(e: React.SyntheticEvent<HTMLImageElement, Event>) => {
                          const target = e.target as HTMLImageElement;
                          target.src = "/placeholder.svg";
                        }}
                      />
                    </div>
                    <div>
                      <h3 className="font-semibold text-lg">{product.name}</h3>
                      <p className="text-sm text-muted-foreground line-clamp-2">
                        {product.description}
                      </p>
                      <div className="flex items-center justify-between mt-2">
                        <span className="font-bold">${product.price}</span>
                        <span className="text-sm text-muted-foreground">
                          Stock: {product.stock_quantity}
                        </span>
                      </div>
                      <div className="flex justify-end space-x-2 mt-4">
                        <Button variant="outline" size="sm">
                          <Edit className="h-4 w-4" />
                        </Button>
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => handleDeleteProduct(product.product_id)}
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
              <div className="grid grid-cols-1 gap-6">
                {dummyOrders.map((order) => (
                  <Card key={order.id}>
                    <CardHeader>
                      <CardTitle className="flex justify-between items-center">
                        <span>Order #{order.id}</span>
                        <span className={`text-sm px-2 py-1 rounded ${
                          order.status === "completed" ? "bg-green-100 text-green-800" : "bg-yellow-100 text-yellow-800"
                        }`}>
                          {order.status}
                        </span>
                      </CardTitle>
                    </CardHeader>
                    <CardContent>
                      <div className="space-y-4">
                        <div className="flex justify-between items-center">
                          <div className="flex items-center">
                            <DollarSign className="h-4 w-4 mr-2 text-gray-500" />
                            <span className="font-semibold">${order.total_amount}</span>
                          </div>
                          <div className="flex items-center">
                            <CreditCard className="h-4 w-4 mr-2 text-gray-500" />
                            <span className="text-gray-600">{order.payment_method}</span>
                          </div>
                        </div>
                        <div className="text-sm text-gray-500">
                          Ordered on: {new Date(order.created_at).toLocaleDateString()}
                        </div>
                        <div className="border-t pt-4">
                          <h4 className="font-medium mb-2">Order Items:</h4>
                          {order.items.map((item, index) => (
                            <div key={index} className="flex justify-between text-sm">
                              <span>Product #{item.product_id}</span>
                              <span>Qty: {item.quantity}</span>
                              <span>${item.price}</span>
                            </div>
                          ))}
                        </div>
                      </div>
                    </CardContent>
                  </Card>
                ))}
              </div>
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