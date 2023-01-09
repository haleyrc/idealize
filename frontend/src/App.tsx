import { Outlet } from "react-router-dom"

import Appbar, { Brand, Nav } from "./design/layout/Appbar"
import Content from "./design/layout/Content"
import Layout from "./design/layout/Layout"

export default function App() {
  const links = [
    { href: "/", label: "Home" },
    { href: "/contact", label: "Contact" },
  ]

  return (
    <Layout>
      <Appbar>
        <Brand>idealize</Brand>
        <Nav links={links} />
      </Appbar>
      <Content>
        <Outlet />
      </Content>
    </Layout>
  )
}
