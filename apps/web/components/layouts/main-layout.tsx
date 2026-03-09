"use client";

import Link from "next/link";
import { useAuthContext } from "@/components/auth/auth-provider";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Button } from "@/components/ui/button";
import { ROUTES } from "@/lib/constants/routes";

const navItems = [
	{ label: "Dashboard", href: ROUTES.dashboard },
	{ label: "Settings", href: ROUTES.settings },
];

export function MainLayout({ children }: { children: React.ReactNode }) {
	const { user, logout } = useAuthContext();

	return (
		<div className="min-h-screen bg-bg text-fg">
			<header className="border-b border-border bg-card/80 backdrop-blur">
				<div className="mx-auto flex max-w-6xl items-center justify-between px-6 py-4">
					<Link href={ROUTES.dashboard} className="text-lg font-semibold">
						MyTodo
					</Link>
					<nav className="hidden items-center gap-6 text-sm md:flex">
						{navItems.map((item) => (
							<Link key={item.href} href={item.href} className="text-fg/70 hover:text-fg">
								{item.label}
							</Link>
						))}
					</nav>
					<div className="flex items-center gap-3">
						<ThemeToggle />
						<Button variant="secondary" onClick={() => logout()}>
							{user?.name ?? "Sign out"}
						</Button>
					</div>
				</div>
			</header>
			<main className="mx-auto max-w-6xl px-6 py-10">{children}</main>
		</div>
	);
}
