"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { ChevronRight, X } from "lucide-react"
import { cn } from "@/lib/utils"

const categories = [
  {
    name: "Laptops",
    subcategories: ["Gaming Laptops", "Business Laptops", "Student Laptops"],
  },
  {
    name: "Smartphones",
    subcategories: ["Android", "iOS", "Budget Phones"],
  },
  {
    name: "PC Components",
    subcategories: ["Processors", "Graphics Cards", "Motherboards", "Memory", "Storage"],
  },
  {
    name: "Accessories",
    subcategories: ["Keyboards", "Mice", "Headsets", "Monitors"],
  },
]

export function CategorySidebar() {
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null)

  return (
    <div className="flex h-full">
      {/* Main Categories */}
      <div className="w-48 border-r bg-card">
        <div className="flex items-center justify-between p-4 border-b">
          <h2 className="text-lg font-semibold text-card-foreground">Categories</h2>
          <Button
            variant="ghost"
            size="icon"
            className="h-8 w-8"
            onClick={() => {
              const event = new CustomEvent("closeSidebar")
              window.dispatchEvent(event)
            }}
          >
            <X className="h-4 w-4" />
            <span className="sr-only">Close sidebar</span>
          </Button>
        </div>
        <div className="p-2">
          {categories.map((category) => (
            <button
              key={category.name}
              onClick={() => setSelectedCategory(category.name)}
              className={cn(
                "w-full text-left px-3 py-2 rounded-md text-sm transition-colors",
                selectedCategory === category.name
                  ? "bg-primary text-primary-foreground"
                  : "text-card-foreground hover:bg-accent hover:text-accent-foreground"
              )}
            >
              {category.name}
            </button>
          ))}
        </div>
      </div>

      {/* Subcategories */}
      {selectedCategory && (
        <div className="flex-1 bg-background">
          <div className="p-4 border-b">
            <h3 className="text-lg font-semibold text-foreground">{selectedCategory}</h3>
          </div>
          <div className="p-2">
            {categories
              .find((cat) => cat.name === selectedCategory)
              ?.subcategories.map((subcategory) => (
                <button
                  key={subcategory}
                  className="w-full text-left px-3 py-2 rounded-md text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors"
                >
                  {subcategory}
                </button>
              ))}
          </div>
        </div>
      )}
    </div>
  )
} 