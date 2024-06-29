import "@/styles/globals.css";
import { Metadata, Viewport } from "next";
import clsx from "clsx";
import { UserProvider } from "@auth0/nextjs-auth0/client";

import { Providers, ThemeProviderProps } from "./providers"; // Assuming ThemeProviderProps is imported from providers.tsx

import { siteConfig } from "@/config/site";
import { fontSans } from "@/config/fonts";
import Navbar from "@/components/layout/Navbar";
import Container from "@/components/common/Container";

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s - ${siteConfig.name}`,
  },
  description: siteConfig.description,
  icons: {
    icon: "/favicon.ico",
  },
};

export const viewport: Viewport = {
  themeColor: [
    { media: "(prefers-color-scheme: light)", color: "white" },
    { media: "(prefers-color-scheme: dark)", color: "black" },
  ],
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html suppressHydrationWarning lang="en">
      <head />
      <UserProvider>
        <body
          className={clsx(
            "min-h-screen bg-background font-sans antialiased",
            fontSans.variable,
          )}
        >
          <Providers themeProps={{ attribute: "class", defaultTheme: "dark", children }}>
            <Container className="relative flex flex-col h-screen">
              <Navbar />
              {children}
            </Container>
          </Providers>
        </body>
      </UserProvider>
    </html>
  );
}
