import { cn } from "@/lib/utils/cn";

type CardProps = React.HTMLAttributes<HTMLDivElement>;

export function Card({ className, ...props }: CardProps) {
	return (
		<div
			className={cn("rounded-3xl border border-border bg-card p-6 shadow-soft", className)}
			{...props}
		/>
	);
}
