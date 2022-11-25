import '../styles/globals.css';
import {QueryClient, QueryClientProvider} from "react-query"

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html className="h-full scroll-smooth">
      <body className="min-h-screen flex flex-col bg-gradient-to-r from-white to-gray-200 text-gray-700">
        {children}
      </body>
    </html>
  )
}
