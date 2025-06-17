"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import { ShoppingCart, Package, CreditCard, Banknote, Plus, Minus } from "lucide-react"
import Image from "next/image"
import Link from "next/link"
import { useGetProductsQuery } from "@/lib/api"
import type { Product } from "@/lib/types"

interface CartItem {
  product_id: number
  quantity: number
}

// Add the random image function
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

export default function CartPage() {
  const [cartItems, setCartItems] = useState<CartItem[]>([])
  const [paymentMethod, setPaymentMethod] = useState<"credit" | "debit">("credit")
  const { data: productsData } = useGetProductsQuery()
  const products = productsData?.data.products || []

  useEffect(() => {
    // Get cart items from localStorage
    const savedCart = localStorage.getItem("cart")
    if (savedCart) {
      setCartItems(JSON.parse(savedCart))
    }
  }, [])

  const handleQuantityChange = (productId: number, newQuantity: number) => {
    if (newQuantity < 1) return

    setCartItems(prev => {
      const updated = prev.map(item =>
        item.product_id === productId ? { ...item, quantity: newQuantity } : item
      )
      localStorage.setItem("cart", JSON.stringify(updated))
      return updated
    })
  }

  const handleRemoveItem = (productId: number) => {
    setCartItems(prev => {
      const updated = prev.filter(item => item.product_id !== productId)
      localStorage.setItem("cart", JSON.stringify(updated))
      return updated
    })
  }

  const getCartProducts = () => {
    return cartItems.map(item => {
      const product = products.find(p => p.product_id === item.product_id)
      return product ? {
        ...product,
        quantity: item.quantity
      } : null
    }).filter(Boolean) as (Product & { quantity: number })[]
  }

  const cartProducts = getCartProducts()

  const subtotal = cartProducts.reduce((sum, item) => sum + (item.price * item.quantity), 0)
  const shipping = subtotal > 0 ? 10 : 0
  const total = subtotal + shipping

  return (
    <div className="container mx-auto p-6">
      <div className="flex items-center space-x-2 mb-6">
        <ShoppingCart className="h-6 w-6 text-indigo-600" />
        <h1 className="text-2xl font-bold">Shopping Cart</h1>
      </div>

      {cartProducts.length === 0 ? (
        <Card>
          <CardContent className="p-6 text-center">
            <Package className="h-12 w-12 text-gray-400 mx-auto mb-4" />
            <h2 className="text-xl font-semibold mb-2">Your cart is empty</h2>
            <p className="text-gray-600 mb-4">Add some products to your cart to see them here</p>
            <Link href="/home">
              <Button>Continue Shopping</Button>
            </Link>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Cart Items */}
          <div className="lg:col-span-2 space-y-4">
            {cartProducts.map((item) => (
              <Card key={item.product_id}>
                <CardContent className="p-4">
                  <div className="flex items-center space-x-4">
                    <div className="relative h-24 w-24 bg-gray-100 rounded-lg">
                      <Image
                        src={getRandomProductImage()}
                        alt={item.name}
                        fill
                        className="object-cover rounded-lg"
                      />
                    </div>
                    <div className="flex-1">
                      <h3 className="font-semibold">{item.name}</h3>
                      <p className="text-sm text-gray-600 mb-2">Stock: {item.stock_quantity}</p>
                      <div className="flex items-center justify-between">
                        <div className="flex items-center space-x-2">
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => handleQuantityChange(item.product_id, item.quantity - 1)}
                            disabled={item.quantity <= 1}
                          >
                            <Minus className="h-4 w-4" />
                          </Button>
                          <span className="w-8 text-center">{item.quantity}</span>
                          <Button
                            variant="outline"
                            size="sm"
                            onClick={() => handleQuantityChange(item.product_id, item.quantity + 1)}
                            disabled={item.quantity >= item.stock_quantity}
                          >
                            <Plus className="h-4 w-4" />
                          </Button>
                        </div>
                        <div className="text-right">
                          <p className="font-semibold">${(item.price * item.quantity).toFixed(2)}</p>
                          <p className="text-sm text-gray-600">${item.price.toFixed(2)} each</p>
                        </div>
                      </div>
                    </div>
                    <Button
                      variant="ghost"
                      size="sm"
                      onClick={() => handleRemoveItem(item.product_id)}
                    >
                      Remove
                    </Button>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>

          {/* Order Summary */}
          <div className="space-y-4">
            <Card>
              <CardHeader>
                <CardTitle>Order Summary</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="space-y-2">
                  <div className="flex justify-between">
                    <span>Subtotal</span>
                    <span>${subtotal.toFixed(2)}</span>
                  </div>
                  <div className="flex justify-between">
                    <span>Shipping</span>
                    <span>${shipping.toFixed(2)}</span>
                  </div>
                  <div className="border-t pt-2">
                    <div className="flex justify-between font-semibold">
                      <span>Total</span>
                      <span>${total.toFixed(2)}</span>
                    </div>
                  </div>
                </div>

                <div className="space-y-4">
                  <Label>Payment Method</Label>
                  <RadioGroup
                    value={paymentMethod}
                    onValueChange={(value) => setPaymentMethod(value as "credit" | "debit")}
                    className="space-y-2"
                  >
                    <div className="flex items-center space-x-2">
                      <RadioGroupItem value="credit" id="credit" />
                      <Label htmlFor="credit" className="flex items-center space-x-2">
                        <CreditCard className="h-4 w-4" />
                        <span>Credit Card</span>
                      </Label>
                    </div>
                    <div className="flex items-center space-x-2">
                      <RadioGroupItem value="debit" id="debit" />
                      <Label htmlFor="debit" className="flex items-center space-x-2">
                        <Banknote className="h-4 w-4" />
                        <span>Debit Card</span>
                      </Label>
                    </div>
                  </RadioGroup>
                </div>

                <Button className="w-full" size="lg">
                  Proceed to Checkout
                </Button>
              </CardContent>
            </Card>
          </div>
        </div>
      )}
    </div>
  )
}
