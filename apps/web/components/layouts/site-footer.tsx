import Link from "next/link";
import { APP_NAME } from "@/lib/constants/config";

const footerLinks = {
  Product: [
    { label: "Dashboard", href: "/dashboard" },
    { label: "Issues", href: "/issues" },
    { label: "Projects", href: "/projects" },
    { label: "Boards", href: "/boards" },
  ],
  Company: [
    { label: "About", href: "/about" },
    { label: "Blog", href: "/blog" },
    { label: "Careers", href: "/careers" },
  ],
  Support: [
    { label: "Documentation", href: "/docs" },
    { label: "Changelog", href: "/changelog" },
    { label: "Status", href: "/status" },
  ],
  Legal: [
    { label: "Privacy", href: "/privacy" },
    { label: "Terms", href: "/terms" },
  ],
};

export function SiteFooter() {
  return (
    <footer className="border-t border-border bg-card/40 text-sm text-fg/70">
      <div className="mx-auto max-w-7xl px-6 py-12">
        <div className="grid grid-cols-2 gap-8 md:grid-cols-5">
          {/* Brand column */}
          <div className="col-span-2 md:col-span-1">
            <Link href="/" className="flex items-center gap-2 font-semibold text-fg">
              <span className="flex h-7 w-7 items-center justify-center rounded-lg bg-accent text-xs font-bold text-white">
                MT
              </span>
              <span>{APP_NAME}</span>
            </Link>
            <p className="mt-3 text-xs leading-relaxed text-fg/60">
              A clean, extensible workspace for your team&apos;s tickets and
              projects.
            </p>
          </div>

          {/* Link columns */}
          {Object.entries(footerLinks).map(([group, links]) => (
            <div key={group}>
              <h4 className="mb-3 text-xs font-semibold uppercase tracking-wider text-fg/50">
                {group}
              </h4>
              <ul className="space-y-2">
                {links.map((link) => (
                  <li key={link.href}>
                    <Link
                      href={link.href}
                      className="transition hover:text-fg"
                    >
                      {link.label}
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>

        <div className="mt-10 flex flex-col items-center justify-between gap-4 border-t border-border pt-6 text-xs text-fg/50 sm:flex-row">
          <span>
            &copy; {new Date().getFullYear()} {APP_NAME}. All rights reserved.
          </span>
          <span>Built with ❤ for productive teams.</span>
        </div>
      </div>
    </footer>
  );
}
