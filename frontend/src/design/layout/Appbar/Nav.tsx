import { Link as RouterLink } from "react-router-dom"

type Link = {
  href: string
  label: string
}

type NavProps = {
  links: Link[]
}

export default function Nav({ links }: NavProps) {
  return (
    <nav role="navigation" aria-label="Main">
      {links.map((link) => (
        <RouterLink
          key={link.href}
          className="mr-4 last:mr-0 hover:underline"
          to={link.href}
        >
          {link.label}
        </RouterLink>
      ))}
    </nav>
  )
}
