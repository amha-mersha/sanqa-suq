import type React from "react"
import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "./globals.css"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "SanqaSuq - Your Ultimate Electronics Destination",
  description:
    "Discover cutting-edge laptops, smartphones, and electronics. Build your dream PC with our Custom PC Builder featuring real-time compatibility validation.",
  keywords: "electronics, laptops, smartphones, PC builder, gaming, components",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        {children}
      </body>
    </html>
  )
}
