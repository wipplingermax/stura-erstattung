import '../styles/globals.css';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html>
      <head>
        <title>Formular Erstattung Semesterbeitrag</title>
      </head>
      <body className="flex bg-gradient-to-r from-white to-gray-200 h-screen text-gray-700">
        {children}
      </body>
    </html>
  )
}
