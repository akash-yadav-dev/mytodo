import { forwardRef } from "react";
import Link from "next/link";
import { cn } from "@/lib/utils/cn";

type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement> & {
	variant?: "primary" | "secondary" | "ghost";
	asChild?: boolean;
	href?: string;
};

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
	({ className, variant = "primary", asChild, href, ...props }, ref) => {
		const baseStyles =
			"inline-flex items-center justify-center rounded-full px-4 py-2 text-sm font-medium transition hover:-translate-y-0.5";
		const variants = {
			primary: "bg-accent text-white shadow-soft hover:shadow-lg",
			secondary: "bg-card text-fg border border-border hover:border-fg/30",
			ghost: "bg-transparent text-fg/80 hover:text-fg",
		};

		if (asChild && href) {
			return (
				<Link className={cn(baseStyles, variants[variant], className)} href={href}>
					{props.children}
				</Link>
			);
		}

		return (
			<button
				ref={ref}
				className={cn(baseStyles, variants[variant], className)}
				{...props}
			/>
		);
	}
);

Button.displayName = "Button";
