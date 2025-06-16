"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Input } from "@/components/ui/input"
import { Separator } from "@/components/ui/separator"
import { Minus, Plus, Trash2, ArrowLeft, ShoppingBag, Tag } from "lucide-react"
import Image from "next/image"
import Link from "next/link"

// Mock cart data
const initialCartItems = [
  {
    id: 1,
    name: "Gaming Laptop RTX 4070",
    price: 1299.99,
    quantity: 1,
    image: "/placeholder.svg?height=100&width=100",
    category: "Laptops",
    inStock: true,
    maxQuantity: 5,
  },
  {
    id: 2,
    name: "RTX 4080 Graphics Card",
    price: 899.99,
    quantity: 1,
    image: "/placeholder.svg?height=100&width=100",
    category: "PC Components",
    inStock: true,
    maxQuantity: 3,
  },
  {
    id: 3,
    name: "Sony WH-1000XM5",
    price: 349.99,
    quantity: 2,
    image: "/placeholder.svg?height=100&width=100",
    category: "Audio",
    inStock: false,
    maxQuantity: 10,
  },
]

export default function CartPage() {
  const [cartItems, setCartItems] = useState(initialCartItems)
  const [promoCode, setPromoCode] = useState("")
  const [discount, setDiscount] = useState(0)
  const [isCheckingOut, setIsCheckingOut] = useState(false)

  const updateQuantity = (id: number, newQuantity: number) => {
    if (newQuantity < 1) return

    setCartItems((items) =>
      items.map((item) => {
        if (item.id === id) {
          return { ...item, quantity: Math.min(newQuantity, item.maxQuantity) }
        }
        return item
      }),
    )
  }

  const removeItem = (id: number) => {
    setCartItems((items) => items.filter((item) => item.id !== id))
  }

  const applyPromoCode = () => {
    if (promoCode.toLowerCase() === "save10") {
      setDiscount(0.1) // 10% discount
    } else if (promoCode.toLowerCase() === "welcome20") {
      setDiscount(0.2) // 20% discount
    } else {
      setDiscount(0)
    }
  }

  const subtotal = cartItems.reduce((sum, item) => sum + item.price * item.quantity, 0)
  const discountAmount = subtotal * discount
  const shipping = subtotal > 100 ? 0 : 15.99
  const tax = (subtotal - discountAmount) * 0.08 // 8% tax
  const total = subtotal - discountAmount + shipping + tax

  const handleCheckout = async () => {
    setIsCheckingOut(true)
    // Simulate checkout process
    await new Promise((resolve) => setTimeout(resolve, 2000))
    // Redirect to checkout page
    window.location.href = "/checkout"
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-4">
              <Link href="/home">
                <Button variant="outline" size="sm">
                  <ArrowLeft className="h-4 w-4 mr-2" />
                  Continue Shopping
                </Button>
              </Link>
              <h1 className="text-2xl font-bold text-gray-900">Shopping Cart</h1>
            </div>
            <Badge variant="secondary">{cartItems.reduce((sum, item) => sum + item.quantity, 0)} Items</Badge>
          </div>
        </div>
      </header>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {cartItems.length === 0 ? (
          <div className="text-center py-16">
            <ShoppingBag className="h-24 w-24 text-gray-300 mx-auto mb-4" />
            <h2 className="text-2xl font-semibold text-gray-900 mb-2">Your cart is empty</h2>
            <p className="text-gray-600 mb-8">Add some products to get started</p>
            <Link href="/home">
              <Button size="lg">Start Shopping</Button>
            </Link>
          </div>
        ) : (
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            {/* Cart Items */}
            <div className="lg:col-span-2 space-y-4">
              {cartItems.map((item) => (
                <Card key={item.id}>
                  <CardContent className="p-6">
                    <div className="flex items-center space-x-4">
                      <Image
                        src={item.image || "/placeholder.svg"}
                        alt={item.name}
                        width={100}
                        height={100}
                        className="rounded-lg"
                      />

                      <div className="flex-1">
                        <h3 className="font-semibold text-lg">{item.name}</h3>
                        <p className="text-gray-600">{item.category}</p>
                        <div className="flex items-center mt-2">
                          {item.inStock ? (
                            <Badge variant="default">In Stock</Badge>
                          ) : (
                            <Badge variant="destructive">Out of Stock</Badge>
                          )}
                        </div>
                      </div>

                      <div className="flex items-center space-x-3">
                        <Button
                          variant="outline"
                          size="icon"
                          onClick={() => updateQuantity(item.id, item.quantity - 1)}
                          disabled={item.quantity <= 1}
                        >
                          <Minus className="h-4 w-4" />
                        </Button>
                        <span className="w-12 text-center font-medium">{item.quantity}</span>
                        <Button
                          variant="outline"
                          size="icon"
                          onClick={() => updateQuantity(item.id, item.quantity + 1)}
                          disabled={item.quantity >= item.maxQuantity}
                        >
                          <Plus className="h-4 w-4" />
                        </Button>
                      </div>

                      <div className="text-right">
                        <p className="text-xl font-bold">${(item.price * item.quantity).toFixed(2)}</p>
                        <p className="text-sm text-gray-600">${item.price.toFixed(2)} each</p>
                      </div>

                      <Button
                        variant="outline"
                        size="icon"
                        onClick={() => removeItem(item.id)}
                        className="text-red-600 hover:text-red-800"
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    </div>
                  </CardContent>
                </Card>
              ))}

              {/* Promo Code */}
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center">
                    <Tag className="h-5 w-5 mr-2" />
                    Promo Code
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <div className="flex space-x-2">
                    <Input
                      placeholder="Enter promo code"
                      value={promoCode}
                      onChange={(e) => setPromoCode(e.target.value)}
                    />
                    <Button onClick={applyPromoCode} variant="outline">
                      Apply
                    </Button>
                  </div>
                  {discount > 0 && (
                    <p className="text-green-600 text-sm mt-2">
                      Promo code applied! {(discount * 100).toFixed(0)}% discount
                    </p>
                  )}
                  <div className="mt-4 text-sm text-gray-600">
                    <p>Try these codes:</p>
                    <p>• SAVE10 - 10% off</p>
                    <p>• WELCOME20 - 20% off</p>
                  </div>
                </CardContent>
              </Card>
            </div>

            {/* Order Summary */}
            <div className="lg:col-span-1">
              <Card className="sticky top-8">
                <CardHeader>
                  <CardTitle>Order Summary</CardTitle>
                </CardHeader>
                <CardContent className="space-y-4">
                  <div className="space-y-2">
                    <div className="flex justify-between">
                      <span>Subtotal:</span>
                      <span>${subtotal.toFixed(2)}</span>
                    </div>

                    {discount > 0 && (
                      <div className="flex justify-between text-green-600">
                        <span>Discount ({(discount * 100).toFixed(0)}%):</span>
                        <span>-${discountAmount.toFixed(2)}</span>
                      </div>
                    )}

                    <div className="flex justify-between">
                      <span>Shipping:</span>
                      <span>{shipping === 0 ? "Free" : `$${shipping.toFixed(2)}`}</span>
                    </div>

                    <div className="flex justify-between">
                      <span>Tax:</span>
                      <span>${tax.toFixed(2)}</span>
                    </div>

                    <Separator />

                    <div className="flex justify-between font-bold text-lg">
                      <span>Total:</span>
                      <span>${total.toFixed(2)}</span>
                    </div>
                  </div>

                  {shipping > 0 && (
                    <div className="bg-blue-50 p-3 rounded-lg text-sm text-blue-800">
                      Add ${(100 - subtotal).toFixed(2)} more for free shipping!
                    </div>
                  )}

                  <Button
                    className="w-full"
                    size="lg"
                    onClick={handleCheckout}
                    disabled={isCheckingOut || cartItems.some((item) => !item.inStock)}
                  >
                    {isCheckingOut ? "Processing..." : `Checkout - $${total.toFixed(2)}`}
                  </Button>

                  {cartItems.some((item) => !item.inStock) && (
                    <p className="text-red-600 text-sm text-center">
                      Some items are out of stock. Please remove them to continue.
                    </p>
                  )}

                  <div className="text-center text-sm text-gray-600">
                    <p>Secure checkout with SSL encryption</p>
                  </div>
                </CardContent>
              </Card>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}
