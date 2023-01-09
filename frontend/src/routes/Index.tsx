import { useLoaderData } from "react-router-dom"

import Layout from "./shared/Layout"
import { getMessage } from "../lib/messages"

type LoaderData = {
  message: string
}

export async function loader<LoaderData>() {
  const { message } = await getMessage()
  return { message } as LoaderData
}

export default function Index() {
  const { message } = useLoaderData() as LoaderData
  return (
    <Layout>
      <h1>Index</h1>
      <p>{message}</p>
    </Layout>
  )
}
