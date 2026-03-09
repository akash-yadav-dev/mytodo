import { cn } from "@/lib/utils/cn";

type BadgeProps = React.HTMLAttributes<HTMLSpanElement> & {
	tone?: "default" | "accent";
};

export function Badge({ className, tone = "default", ...props }: BadgeProps) {
	const tones = {
		default: "bg-card text-fg/80 border border-border",
		accent: "bg-accentMuted text-fg border border-accent/30",
	};

	return (
		<span
			className={cn(
				"inline-flex items-center rounded-full px-3 py-1 text-xs font-medium",
				tones[tone],
				className
			)}
			{...props}
		/>
	);
}
