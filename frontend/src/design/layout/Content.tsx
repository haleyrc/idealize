type ContentProps = {
  children: React.ReactNode
}

export default function Content(props: ContentProps) {
  return <div className="flex-grow p-4 bg-gray-100 min-h-max" {...props} />
}
