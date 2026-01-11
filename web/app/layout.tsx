import type { Metadata } from "next";
import "./globals.css";
import { Navbar } from "@/components/navbar";
import { Footer } from "@/components/footer";
import { config } from "@/lib/config";

export const metadata: Metadata = {
  title: {
    default: config.websiteName,
    template: `%s | ${config.websiteName}`,
  },
  description: config.websiteDescription,
  keywords: [
    "internet censorship",
    "India",
    "ISP blocking",
    "website blocking",
    "internet freedom",
    "net neutrality",
  ],
  authors: [{ name: "Internet Chowkidar Team" }],
  openGraph: {
    type: "website",
    locale: "en_US",
    url: config.websiteBaseUrl,
    siteName: config.websiteName,
    title: config.websiteName,
    description: config.websiteDescription,
  },
  twitter: {
    card: "summary_large_image",
    title: config.websiteName,
    description: config.websiteDescription,
  },
  robots: {
    index: true,
    follow: true,
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body className="min-h-screen bg-background antialiased noise-bg">
        <Navbar />
        <main className="pt-16 min-h-screen">{children}</main>
        <Footer />
      </body>
    </html>
  );
}
