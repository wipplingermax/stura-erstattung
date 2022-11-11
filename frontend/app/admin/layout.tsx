import '../../styles/globals.css';
import Header from './header';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html>
      <head></head>
      <body className="bg-body-background text-white">
        <Header />
        {children}</body>
    </html>
  )
}
