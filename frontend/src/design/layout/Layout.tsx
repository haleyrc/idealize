type Props = {
  children: React.ReactNode
}

export default function Layout(props: Props) {
  return <div className="flex flex-col h-full" {...props} />
}
