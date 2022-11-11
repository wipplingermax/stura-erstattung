import '../styles/globals.css';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html>
      <head></head>
      <body className="bg-body-background text-white">
        {children}
      </body>
    </html>
  )
}
