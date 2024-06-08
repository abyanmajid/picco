import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

import { Providers } from "@/lib/chakra/providers";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "PICCO",
  description: "Learn effectively by simply writing more code.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
