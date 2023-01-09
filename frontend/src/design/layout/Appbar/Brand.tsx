import { Link } from "react-router-dom"

type Props = {
  children: string
}

export default function Brand(props: Props) {
  return <Link className="text-xl font-bold text-gray-700" to="/" {...props} />
}
