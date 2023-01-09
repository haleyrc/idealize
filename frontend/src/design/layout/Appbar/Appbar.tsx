type Props = {
  children: React.ReactNode
}

export default function Appbar({ children }: Props) {
  return (
    <div className="flex items-center justify-between p-4 bg-gray-300">
      {children}
    </div>
  )
}
