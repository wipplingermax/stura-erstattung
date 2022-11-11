import '../styles/globals.css';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html>
      <head></head>
      {/* <body className="bg-body-background text-white"> */}
      <body className="bg-gradient-to-r from-white to-gray-200">
    
        {children}
      </body>
    </html>
  )
}
