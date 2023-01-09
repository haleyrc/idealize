import App from "./App"
import Contact from "./routes/Contact"
import Index, { loader as indexLoader } from "./routes/Index"

const routes = [
  {
    path: "/",
    element: <App />,
    children: [
      { index: true, element: <Index />, loader: indexLoader },
      { path: "/contact", element: <Contact /> },
    ],
  },
]

export default routes
