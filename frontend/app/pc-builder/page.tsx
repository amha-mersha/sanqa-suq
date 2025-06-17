"use client"

import { useState, useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Progress } from "@/components/ui/progress"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import {
  Cpu,
  HardDrive,
  MemoryStick,
  Monitor,
  Zap,
  Box,
  AlertTriangle,
  CheckCircle,
  Save,
  Share2,
  ArrowLeft,
} from "lucide-react"
import Image from "next/image"
import Link from "next/link"

// Mock component data with compatibility information
const components = {
  cpu: [
    {
      id: 1,
      name: "Intel Core i9-13900K",
      price: 589.99,
      socket: "LGA1700",
      tdp: 125,
      cores: 24,
      threads: 32,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { socket: "LGA1700", memoryType: "DDR4/DDR5" },
    },
    {
      id: 2,
      name: "AMD Ryzen 9 7900X",
      price: 549.99,
      socket: "AM5",
      tdp: 170,
      cores: 12,
      threads: 24,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { socket: "AM5", memoryType: "DDR5" },
    },
    {
      id: 3,
      name: "Intel Core i7-13700K",
      price: 409.99,
      socket: "LGA1700",
      tdp: 125,
      cores: 16,
      threads: 24,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { socket: "LGA1700", memoryType: "DDR4/DDR5" },
    },
  ],
  motherboard: [
    {
      id: 1,
      name: "ASUS ROG STRIX Z790-E",
      price: 449.99,
      socket: "LGA1700",
      formFactor: "ATX",
      memorySlots: 4,
      maxMemory: 128,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { socket: "LGA1700", memoryType: "DDR5", formFactor: "ATX" },
    },
    {
      id: 2,
      name: "MSI MAG X670E TOMAHAWK",
      price: 389.99,
      socket: "AM5",
      formFactor: "ATX",
      memorySlots: 4,
      maxMemory: 128,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { socket: "AM5", memoryType: "DDR5", formFactor: "ATX" },
    },
    {
      id: 3,
      name: "ASUS PRIME Z790-P",
      price: 229.99,
      socket: "LGA1700",
      formFactor: "ATX",
      memorySlots: 4,
      maxMemory: 128,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { socket: "LGA1700", memoryType: "DDR4/DDR5", formFactor: "ATX" },
    },
  ],
  memory: [
    {
      id: 1,
      name: "Corsair Vengeance DDR5-5600 32GB",
      price: 179.99,
      type: "DDR5",
      speed: 5600,
      capacity: 32,
      modules: 2,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { memoryType: "DDR5" },
    },
    {
      id: 2,
      name: "G.Skill Trident Z5 DDR5-6000 32GB",
      price: 199.99,
      type: "DDR5",
      speed: 6000,
      capacity: 32,
      modules: 2,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { memoryType: "DDR5" },
    },
    {
      id: 3,
      name: "Corsair Vengeance LPX DDR4-3200 32GB",
      price: 129.99,
      type: "DDR4",
      speed: 3200,
      capacity: 32,
      modules: 2,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { memoryType: "DDR4" },
    },
  ],
  gpu: [
    {
      id: 1,
      name: "NVIDIA RTX 4080 SUPER",
      price: 999.99,
      memory: 16,
      powerRequirement: 320,
      length: 310,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { minPowerSupply: 750, pciSlots: 3 },
    },
    {
      id: 2,
      name: "NVIDIA RTX 4070 SUPER",
      price: 599.99,
      memory: 12,
      powerRequirement: 220,
      length: 285,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { minPowerSupply: 650, pciSlots: 2.5 },
    },
    {
      id: 3,
      name: "AMD RX 7800 XT",
      price: 499.99,
      memory: 16,
      powerRequirement: 263,
      length: 295,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { minPowerSupply: 700, pciSlots: 2.5 },
    },
  ],
  storage: [
    {
      id: 1,
      name: "Samsung 980 PRO 2TB NVMe",
      price: 199.99,
      type: "NVMe SSD",
      capacity: 2000,
      interface: "PCIe 4.0",
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { interface: "M.2" },
    },
    {
      id: 2,
      name: "WD Black SN850X 1TB NVMe",
      price: 129.99,
      type: "NVMe SSD",
      capacity: 1000,
      interface: "PCIe 4.0",
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { interface: "M.2" },
    },
    {
      id: 3,
      name: "Seagate Barracuda 2TB HDD",
      price: 59.99,
      type: "HDD",
      capacity: 2000,
      interface: "SATA",
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { interface: "SATA" },
    },
  ],
  psu: [
    {
      id: 1,
      name: "Corsair RM850x 850W 80+ Gold",
      price: 149.99,
      wattage: 850,
      efficiency: "80+ Gold",
      modular: true,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { wattage: 850 },
    },
    {
      id: 2,
      name: "EVGA SuperNOVA 750W 80+ Gold",
      price: 119.99,
      wattage: 750,
      efficiency: "80+ Gold",
      modular: true,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { wattage: 750 },
    },
    {
      id: 3,
      name: "Seasonic Focus GX-650 650W",
      price: 99.99,
      wattage: 650,
      efficiency: "80+ Gold",
      modular: true,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { wattage: 650 },
    },
  ],
  case: [
    {
      id: 1,
      name: "Fractal Design Define 7",
      price: 169.99,
      formFactor: "ATX",
      maxGpuLength: 315,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { formFactor: "ATX", maxGpuLength: 315 },
    },
    {
      id: 2,
      name: "NZXT H7 Flow",
      price: 129.99,
      formFactor: "ATX",
      maxGpuLength: 365,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { formFactor: "ATX", maxGpuLength: 365 },
    },
    {
      id: 3,
      name: "Corsair 4000D Airflow",
      price: 104.99,
      formFactor: "ATX",
      maxGpuLength: 360,
      image: "/placeholder.svg?height=100&width=100",
      compatibility: { formFactor: "ATX", maxGpuLength: 360 },
    },
  ],
}

interface SelectedComponents {
  cpu: any
  motherboard: any
  memory: any
  gpu: any
  storage: any
  psu: any
  case: any
}

interface CompatibilityIssue {
  type: "error" | "warning"
  message: string
  components: string[]
}

export default function PCBuilderPage() {
  const [selectedComponents, setSelectedComponents] = useState<SelectedComponents>({
    cpu: null,
    motherboard: null,
    memory: null,
    gpu: null,
    storage: null,
    psu: null,
    case: null,
  })

  const [compatibilityIssues, setCompatibilityIssues] = useState<CompatibilityIssue[]>([])
  const [totalPrice, setTotalPrice] = useState(0)
  const [estimatedWattage, setEstimatedWattage] = useState(0)

  // Real-time compatibility validation
  useEffect(() => {
    const issues: CompatibilityIssue[] = []
    let wattage = 0

    // Calculate total price
    const price = Object.values(selectedComponents).reduce((sum, component) => {
      return sum + (component?.price || 0)
    }, 0)
    setTotalPrice(price)

    // Calculate estimated wattage
    if (selectedComponents.cpu) wattage += selectedComponents.cpu.tdp
    if (selectedComponents.gpu) wattage += selectedComponents.gpu.powerRequirement
    if (selectedComponents.memory) wattage += 10 // Approximate RAM power consumption
    wattage += 50 // Base system power (motherboard, storage, fans, etc.)
    setEstimatedWattage(wattage)

    // CPU and Motherboard socket compatibility
    if (selectedComponents.cpu && selectedComponents.motherboard) {
      if (selectedComponents.cpu.socket !== selectedComponents.motherboard.socket) {
        issues.push({
          type: "error",
          message: `CPU socket ${selectedComponents.cpu.socket} is not compatible with motherboard socket ${selectedComponents.motherboard.socket}`,
          components: ["cpu", "motherboard"],
        })
      }
    }

    // Memory type compatibility
    if (selectedComponents.cpu && selectedComponents.motherboard && selectedComponents.memory) {
      const cpuMemoryTypes = selectedComponents.cpu.compatibility.memoryType.split("/")
      const motherboardMemoryTypes = selectedComponents.motherboard.compatibility.memoryType.split("/")
      const memoryType = selectedComponents.memory.compatibility.memoryType

      const isCompatible = cpuMemoryTypes.includes(memoryType) && motherboardMemoryTypes.includes(memoryType)

      if (!isCompatible) {
        issues.push({
          type: "error",
          message: `Memory type ${memoryType} is not compatible with selected CPU and motherboard`,
          components: ["cpu", "motherboard", "memory"],
        })
      }
    }

    // Power supply wattage check
    if (selectedComponents.psu && estimatedWattage > 0) {
      const psuWattage = selectedComponents.psu.wattage
      const recommendedWattage = estimatedWattage * 1.2 // 20% headroom

      if (psuWattage < recommendedWattage) {
        issues.push({
          type: "error",
          message: `Power supply (${psuWattage}W) is insufficient. Recommended: ${Math.ceil(recommendedWattage)}W or higher`,
          components: ["psu"],
        })
      } else if (psuWattage < estimatedWattage * 1.1) {
        issues.push({
          type: "warning",
          message: `Power supply wattage is close to system requirements. Consider a higher wattage PSU for better efficiency`,
          components: ["psu"],
        })
      }
    }

    // GPU and Case clearance
    if (selectedComponents.gpu && selectedComponents.case) {
      if (selectedComponents.gpu.length > selectedComponents.case.maxGpuLength) {
        issues.push({
          type: "error",
          message: `Graphics card (${selectedComponents.gpu.length}mm) is too long for the selected case (max: ${selectedComponents.case.maxGpuLength}mm)`,
          components: ["gpu", "case"],
        })
      }
    }

    // Motherboard and Case form factor compatibility
    if (selectedComponents.motherboard && selectedComponents.case) {
      if (selectedComponents.motherboard.formFactor !== selectedComponents.case.formFactor) {
        issues.push({
          type: "error",
          message: `Motherboard form factor (${selectedComponents.motherboard.formFactor}) is not compatible with case (${selectedComponents.case.formFactor})`,
          components: ["motherboard", "case"],
        })
      }
    }

    setCompatibilityIssues(issues)
  }, [selectedComponents, estimatedWattage])

  const handleComponentSelect = (category: keyof SelectedComponents, component: any) => {
    setSelectedComponents((prev) => ({
      ...prev,
      [category]: component,
    }))
  }

  const getCompatibleComponents = (category: keyof typeof components) => {
    return components[category].filter((component) => {
      // Filter based on current selections for better UX
      if (category === "motherboard" && selectedComponents.cpu) {
        return component.socket === selectedComponents.cpu.socket
      }
      if (category === "memory" && selectedComponents.motherboard) {
        const motherboardMemoryTypes = selectedComponents.motherboard.compatibility.memoryType.split("/")
        return motherboardMemoryTypes.includes(component.compatibility.memoryType)
      }
      if (category === "psu" && estimatedWattage > 0) {
        return component.wattage >= estimatedWattage * 1.1
      }
      return true
    })
  }

  const ComponentSelector = ({
    category,
    icon: Icon,
    title,
  }: { category: keyof SelectedComponents; icon: any; title: string }) => {
    const compatibleComponents = getCompatibleComponents(category as keyof typeof components)
    const allComponents = components[category as keyof typeof components]
    const selectedComponent = selectedComponents[category]

    return (
      <Card className="h-full border-0 shadow-lg">
        <CardHeader className="bg-gradient-to-r from-indigo-50 to-purple-50">
          <CardTitle className="flex items-center text-indigo-700">
            <Icon className="h-5 w-5 mr-2" />
            {title}
          </CardTitle>
          {selectedComponent && (
            <CardDescription className="text-indigo-600">
              Selected: {selectedComponent.name} - ${selectedComponent.price}
            </CardDescription>
          )}
        </CardHeader>
        <CardContent className="p-6">
          <Select
            value={selectedComponent?.id?.toString() || ""}
            onValueChange={(value) => {
              const component = allComponents.find((c) => c.id.toString() === value)
              const isCompatible = compatibleComponents.some((c) => c.id.toString() === value)
              if (isCompatible) {
                handleComponentSelect(category, component)
              }
            }}
          >
            <SelectTrigger className="border-indigo-200 focus:border-indigo-500">
              <SelectValue placeholder={`Choose ${title}`} />
            </SelectTrigger>
            <SelectContent>
              {allComponents.map((component) => {
                const isCompatible = compatibleComponents.some((c) => c.id === component.id)
                return (
                  <SelectItem
                    key={component.id}
                    value={component.id.toString()}
                    disabled={!isCompatible}
                    className={!isCompatible ? "opacity-50 line-through" : ""}
                  >
                    <div className="flex items-center justify-between w-full">
                      <span className={!isCompatible ? "text-gray-400" : ""}>{component.name}</span>
                      <span className={`ml-2 font-semibold ${!isCompatible ? "text-gray-400" : "text-green-600"}`}>
                        ${component.price}
                      </span>
                      {!isCompatible && (
                        <Badge variant="destructive" className="ml-2 text-xs">
                          Incompatible
                        </Badge>
                      )}
                    </div>
                  </SelectItem>
                )
              })}
            </SelectContent>
          </Select>

          {selectedComponent && (
            <div className="mt-4 p-4 bg-gradient-to-r from-blue-50 to-indigo-50 rounded-lg border border-indigo-100">
              <div className="flex items-center mb-3">
                <Image
                  src={selectedComponent.image || "/placeholder.svg"}
                  alt={selectedComponent.name}
                  width={60}
                  height={60}
                  className="rounded mr-3 border border-indigo-200"
                />
                <div>
                  <h4 className="font-semibold text-indigo-900">{selectedComponent.name}</h4>
                  <p className="text-lg font-bold text-green-600">${selectedComponent.price}</p>
                </div>
              </div>

              {/* Component-specific details with better styling */}
              <div className="text-sm text-indigo-700 space-y-1 bg-white/50 p-3 rounded border border-indigo-100">
                {category === "cpu" && (
                  <>
                    <p>
                      <span className="font-medium">Socket:</span> {selectedComponent.socket}
                    </p>
                    <p>
                      <span className="font-medium">Cores/Threads:</span> {selectedComponent.cores}/
                      {selectedComponent.threads}
                    </p>
                    <p>
                      <span className="font-medium">TDP:</span> {selectedComponent.tdp}W
                    </p>
                  </>
                )}
                {category === "gpu" && (
                  <>
                    <p>
                      <span className="font-medium">VRAM:</span> {selectedComponent.memory}GB
                    </p>
                    <p>
                      <span className="font-medium">Power:</span> {selectedComponent.powerRequirement}W
                    </p>
                    <p>
                      <span className="font-medium">Length:</span> {selectedComponent.length}mm
                    </p>
                  </>
                )}
                {category === "memory" && (
                  <>
                    <p>
                      <span className="font-medium">Type:</span> {selectedComponent.type}
                    </p>
                    <p>
                      <span className="font-medium">Speed:</span> {selectedComponent.speed} MHz
                    </p>
                    <p>
                      <span className="font-medium">Capacity:</span> {selectedComponent.capacity}GB
                    </p>
                  </>
                )}
                {category === "psu" && (
                  <>
                    <p>
                      <span className="font-medium">Wattage:</span> {selectedComponent.wattage}W
                    </p>
                    <p>
                      <span className="font-medium">Efficiency:</span> {selectedComponent.efficiency}
                    </p>
                    <p>
                      <span className="font-medium">Modular:</span> {selectedComponent.modular ? "Yes" : "No"}
                    </p>
                  </>
                )}
              </div>
            </div>
          )}
        </CardContent>
      </Card>
    )
  }

  const hasErrors = compatibilityIssues.some((issue) => issue.type === "error")
  const completionPercentage = (Object.values(selectedComponents).filter(Boolean).length / 7) * 100

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-lg border-b border-indigo-100">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-4">
              <Link href="/home">
                <Button variant="outline" size="sm" className="border-indigo-200 text-indigo-600 hover:bg-indigo-50">
                  <ArrowLeft className="h-4 w-4 mr-2" />
                  Back to Shop
                </Button>
              </Link>
              <h1 className="text-2xl font-bold bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent">
                Custom PC Builder
              </h1>
            </div>
            <div className="flex items-center space-x-4">
              <Button variant="outline" className="border-indigo-200 text-indigo-600 hover:bg-indigo-50">
                <Save className="h-4 w-4 mr-2" />
                Save Build
              </Button>
              <Button variant="outline" className="border-indigo-200 text-indigo-600 hover:bg-indigo-50">
                <Share2 className="h-4 w-4 mr-2" />
                Share
              </Button>
            </div>
          </div>
        </div>
      </header>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="grid grid-cols-1 lg:grid-cols-4 gap-8">
          {/* Component Selection */}
          <div className="lg:col-span-3">
            <div className="mb-8">
              <div className="flex items-center justify-between mb-4">
                <h2 className="text-xl font-semibold bg-gradient-to-r from-indigo-600 to-purple-600 bg-clip-text text-transparent">
                  Build Progress
                </h2>
                <span className="text-sm text-gray-600 bg-indigo-50 px-3 py-1 rounded-full">
                  {Math.round(completionPercentage)}% Complete
                </span>
              </div>
              <Progress value={completionPercentage} className="h-3 bg-gradient-to-r from-indigo-100 to-purple-100" />
            </div>

            <Tabs defaultValue="components" className="w-full">
              <TabsList className="grid w-full grid-cols-2">
                <TabsTrigger value="components">Select Components</TabsTrigger>
                <TabsTrigger value="compatibility">Compatibility Check</TabsTrigger>
              </TabsList>

              <TabsContent value="components" className="space-y-6">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <ComponentSelector category="cpu" icon={Cpu} title="Processor (CPU)" />
                  <ComponentSelector category="motherboard" icon={Monitor} title="Motherboard" />
                  <ComponentSelector category="memory" icon={MemoryStick} title="Memory (RAM)" />
                  <ComponentSelector category="gpu" icon={Monitor} title="Graphics Card" />
                  <ComponentSelector category="storage" icon={HardDrive} title="Storage" />
                  <ComponentSelector category="psu" icon={Zap} title="Power Supply" />
                  <ComponentSelector category="case" icon={Box} title="Case" />
                </div>
              </TabsContent>

              <TabsContent value="compatibility" className="space-y-4">
                <div className="space-y-4">
                  {compatibilityIssues.length === 0 ? (
                    <Alert>
                      <CheckCircle className="h-4 w-4" />
                      <AlertDescription>
                        All selected components are compatible! Your build looks great.
                      </AlertDescription>
                    </Alert>
                  ) : (
                    compatibilityIssues.map((issue, index) => (
                      <Alert key={index} variant={issue.type === "error" ? "destructive" : "default"}>
                        <AlertTriangle className="h-4 w-4" />
                        <AlertDescription>{issue.message}</AlertDescription>
                      </Alert>
                    ))
                  )}
                </div>

                {/* Detailed compatibility matrix */}
                <Card>
                  <CardHeader>
                    <CardTitle>Compatibility Matrix</CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="space-y-4">
                      <div className="grid grid-cols-2 gap-4 text-sm">
                        <div>
                          <h4 className="font-semibold mb-2">Socket Compatibility</h4>
                          <p>CPU: {selectedComponents.cpu?.socket || "Not selected"}</p>
                          <p>Motherboard: {selectedComponents.motherboard?.socket || "Not selected"}</p>
                          <Badge
                            variant={
                              selectedComponents.cpu &&
                              selectedComponents.motherboard &&
                              selectedComponents.cpu.socket === selectedComponents.motherboard.socket
                                ? "default"
                                : "destructive"
                            }
                          >
                            {selectedComponents.cpu && selectedComponents.motherboard
                              ? selectedComponents.cpu.socket === selectedComponents.motherboard.socket
                                ? "Compatible"
                                : "Incompatible"
                              : "Incomplete"}
                          </Badge>
                        </div>
                        <div>
                          <h4 className="font-semibold mb-2">Power Requirements</h4>
                          <p>Estimated Usage: {estimatedWattage}W</p>
                          <p>PSU Capacity: {selectedComponents.psu?.wattage || 0}W</p>
                          <Badge
                            variant={
                              selectedComponents.psu && selectedComponents.psu.wattage >= estimatedWattage * 1.1
                                ? "default"
                                : "destructive"
                            }
                          >
                            {selectedComponents.psu
                              ? selectedComponents.psu.wattage >= estimatedWattage * 1.1
                                ? "Sufficient"
                                : "Insufficient"
                              : "No PSU Selected"}
                          </Badge>
                        </div>
                      </div>
                    </div>
                  </CardContent>
                </Card>
              </TabsContent>
            </Tabs>
          </div>

          {/* Build Summary */}
          <div className="lg:col-span-1">
            <Card className="sticky top-8 border-0 shadow-xl bg-gradient-to-br from-white to-indigo-50">
              <CardHeader className="bg-gradient-to-r from-indigo-600 to-purple-600 text-white rounded-t-lg">
                <CardTitle>Build Summary</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4 p-6">
                <div className="space-y-3">
                  <div className="flex justify-between items-center">
                    <span className="text-gray-600">Total Price:</span>
                    <span className="font-bold text-2xl bg-gradient-to-r from-green-600 to-emerald-600 bg-clip-text text-transparent">
                      ${totalPrice.toFixed(2)}
                    </span>
                  </div>
                  <div className="flex justify-between text-sm text-gray-600 bg-white p-2 rounded border border-indigo-100">
                    <span>Est. Wattage:</span>
                    <span className="font-medium text-indigo-600">{estimatedWattage}W</span>
                  </div>
                  <div className="flex justify-between text-sm text-gray-600 bg-white p-2 rounded border border-indigo-100">
                    <span>Components:</span>
                    <span className="font-medium text-indigo-600">
                      {Object.values(selectedComponents).filter(Boolean).length}/7
                    </span>
                  </div>
                </div>

                {compatibilityIssues.length > 0 && (
                  <div className="space-y-2 p-3 bg-red-50 rounded-lg border border-red-200">
                    <h4 className="font-semibold text-sm text-red-800">Issues Found:</h4>
                    <div className="space-y-1">
                      {compatibilityIssues.map((issue, index) => (
                        <Badge
                          key={index}
                          variant={issue.type === "error" ? "destructive" : "secondary"}
                          className="text-xs mr-1"
                        >
                          {issue.type === "error" ? "❌ Error" : "⚠️ Warning"}
                        </Badge>
                      ))}
                    </div>
                  </div>
                )}

                <div className="space-y-2 pt-4 border-t border-indigo-200">
                  <Button
                    className="w-full bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-700 hover:to-emerald-700 text-white shadow-lg"
                    disabled={hasErrors || Object.values(selectedComponents).filter(Boolean).length < 7}
                  >
                    Add to Cart - ${totalPrice.toFixed(2)}
                  </Button>
                  <Button variant="outline" className="w-full border-indigo-200 text-indigo-600 hover:bg-indigo-50">
                    Save Build
                  </Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  )
}
