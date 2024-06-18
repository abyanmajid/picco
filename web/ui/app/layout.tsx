import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

import { Providers } from "@/lib/chakra/providers";
import Navbar from "@/components/layout/Navbar";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "codemore.io",
  description: "Learn effectively by simply writing more code.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className} style={{ backgroundColor: '#262525' }}>
        <Providers>
          <Navbar />
          {children}
        </Providers>
      </body>
    </html>
  );
}