import Header from "../header"

export default function RequestLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <>
      <Header />   
      {children}
    </>
  )
}
